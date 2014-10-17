package entropy

// #include <sqlite3.h>
// #cgo LDFLAGS: -lsqlite3
import "C"
import (
	"bytes"
	"database/sql"
	"errors"
	"fmt"
	log "github.com/cihub/seelog"
	_ "github.com/mattn/go-sqlite3"
	"github.com/ostensus/ostn/entropy/migration"
	"github.com/relops/sqlc/sqlc"
	"text/template"
)

type VersionStore struct {
	db      *sql.DB
	dialect sqlc.Dialect
}

var repoTmpl, upsertTmpl *template.Template

func init() {
	m := template.FuncMap{
		"columnType": columnType,
	}

	repoBin, _ := repo_tmpl()
	repoTmpl = template.Must(template.New("repo.tmpl").Funcs(m).Parse(string(repoBin)))

	upsertBin, _ := upsert_tmpl()
	upsertTmpl = template.Must(template.New("upsert.tmpl").Funcs(m).Parse(string(upsertBin)))
}

func OpenStore(path string) (*VersionStore, error) {

	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}

	v := &VersionStore{
		db:      db,
		dialect: sqlc.Sqlite,
	}

	if err := autoload(); err != nil {
		return nil, err
	}

	if _, err := db.Exec(`PRAGMA foreign_keys = ON;`); err != nil {
		return nil, err
	}

	steps := sqlc.LoadBindata(migration.AssetNames(), migration.Asset)
	if err = sqlc.Migrate(db, sqlc.Sqlite, steps); err != nil {
		return nil, err
	}

	// TODO delete this hardcorded create
	_, err = db.Exec(create)
	if err != nil {
		return nil, err
	}

	return v, err
}

func (v *VersionStore) SliceThreshold() int {
	return 128
}

const newRepo = `INSERT INTO repositories (source, name) VALUES (?, ?);"`
const newPartitionName = `INSERT INTO unique_partition_names (repository, name) VALUES (?, ?);"`
const newRangePartition = `INSERT INTO range_partitions (repository, name) VALUES (?, ?);"`
const newSetPartition = `INSERT INTO set_partitions (repository, name, value) VALUES (?, ?, ?);"`

func (v *VersionStore) NewRepository(src, repoName string, parts map[string]PartitionDescriptor) (repo int64, err error) {
	tx, err := v.db.Begin()
	if err != nil {
		return repo, err
	}

	st, err := tx.Prepare(newRepo)
	if err != nil {
		tx.Rollback()
		return repo, err
	}
	res, err := st.Exec(src, repoName)
	if err != nil {
		tx.Rollback()
		return repo, err
	}

	repo, err = res.LastInsertId()
	if err != nil {
		tx.Rollback()
		return repo, err
	}

	for name, desc := range parts {

		st, err = tx.Prepare(newPartitionName)
		if err != nil {
			tx.Rollback()
			return repo, err
		}
		_, err = st.Exec(repo, name)
		if err != nil {
			tx.Rollback()
			return repo, err
		}

		switch d := desc.(type) {
		case *SetPartitionDescriptor:

			for _, value := range d.Values {
				st, err = tx.Prepare(newSetPartition)
				if err != nil {
					tx.Rollback()
					return repo, err
				}
				_, err = st.Exec(repo, name, value)
				if err != nil {
					tx.Rollback()
					return repo, err
				}
			}

		case *RangePartitionDescriptor:

			_ = d.DataType

			st, err = tx.Prepare(newRangePartition)
			if err != nil {
				tx.Rollback()
				return repo, err
			}
			_, err = st.Exec(repo, name)
			if err != nil {
				tx.Rollback()
				return repo, err
			}
		}

	}

	/////////////////////////////////////////////////////////////////

	sql := renderSQL(repo, parts, repoTmpl)

	log.Infof("New repo: \n\n____________\n%s\n____________\n", sql)

	_, err = tx.Exec(sql)
	if err != nil {
		tx.Rollback()
		return repo, err
	}

	/////////////////////////////////////////////////////////////////

	err = tx.Commit()
	if err != nil {
		return repo, err
	}

	return repo, err
}

func renderSQL(repo int64, parts map[string]PartitionDescriptor, t *template.Template) string {
	params := map[string]interface{}{
		"Postfix":    repo,
		"Partitions": parts,
	}

	var b bytes.Buffer
	t.Execute(&b, params)

	return b.String()
}

func columnType(d PartitionDescriptor) string {
	switch d.(type) {
	case *RangePartitionDescriptor:
		return "TIMESTAMP"
	default:
		return "TEXT"
	}
}

func (v *VersionStore) Accept(repo int64, ev ChangeEvent) error {

	// TODO validate the repo id

	parted, ok := ev.(PartitionedEvent)
	if ok {

		parts := make(map[string]PartitionDescriptor)
		n := len(parted.Attributes()) + 2
		args := make([]interface{}, n)
		i := 0
		for name, value := range parted.Attributes() {
			parts[name] = &RangePartitionDescriptor{}
			args[i] = value
			i++
		}

		args[n-2] = parted.Id()
		args[n-1] = parted.Version()

		// TODO need to validate the supplied attributes

		sql := renderSQL(repo, parts, upsertTmpl)

		st, err := v.db.Prepare(sql)
		if err != nil {
			return err
		}
		defer st.Close()
		_, err = st.Exec(args...)
		return err
	}

	return errors.New("Bogus event")
}

func (v *VersionStore) Digest(repo int64) (map[string]string, error) {

	parts := make(map[string]PartitionDescriptor)

	rows, err := sqlc.Select(UNIQUE_PARTITION_NAMES.NAME, SET_PARTITIONS.VALUE).
		From(UNIQUE_PARTITION_NAMES).
		LeftOuterJoin(RANGE_PARTITIONS).
		On(UNIQUE_PARTITION_NAMES.NAME.IsEq(RANGE_PARTITIONS.NAME),
		UNIQUE_PARTITION_NAMES.REPOSITORY.IsEq(RANGE_PARTITIONS.REPOSITORY)).
		LeftOuterJoin(SET_PARTITIONS).
		On(UNIQUE_PARTITION_NAMES.NAME.IsEq(SET_PARTITIONS.NAME),
		UNIQUE_PARTITION_NAMES.REPOSITORY.IsEq(SET_PARTITIONS.REPOSITORY)).
		Where(UNIQUE_PARTITION_NAMES.REPOSITORY.Eq(repo)).
		Query(v.dialect, v.db)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var name, value string
		_ = rows.Scan(&name, &value)

		log.Infof("Name(%s) -> Value(%s)", name, value)

		if value == "" {
			parts[name] = &RangePartitionDescriptor{}
		} else {
			parts[name] = &SetPartitionDescriptor{Values: []string{value}}
		}
	}

	// v_ID
	tableName := fmt.Sprintf("v_%d", repo)
	t := sqlc.Table(tableName)

	lhs := t.As("lhs")
	rhs := t.As("rhs")

	lhsIdField := lhs.StringField("id")
	rhsIdField := rhs.StringField("id")

	lhsVersion := sqlc.GroupConcat(lhs.StringField("version")).Separator("").Md5().Hex().Lower().As("version")

	sliceThreshold := 127
	bucket := sqlc.Count().Cast("REAL").Div(sliceThreshold).Cast("INT").As("bucket")

	q := sqlc.Select(lhsIdField.As("id"), lhsVersion, bucket).
		From(lhs).
		Join(rhs).On(lhsIdField.IsGe(rhsIdField)).
		GroupBy(lhsIdField)

	rows, err = q.Query(v.dialect, v.db)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	digests := make(map[string]string)

	for rows.Next() {
		var id string
		var digest string
		var bucket int // currently not used for much
		rows.Scan(&id, &digest, &bucket)
		digests[id] = digest
	}

	return digests, nil
}

func (v *VersionStore) Interview(repository int64, cons []Constraint, aggs []Aggregation, answer []Answer) *Iter {
	if len(aggs) != 1 {
		return &Iter{err: errors.New("Must supply 1 aggregation")}
	}

	aggregation, ok := aggs[0].(*DateAggregation)
	if !ok {
		return &Iter{err: errors.New("Must be a date aggregation")}
	}

	if aggregation.Granularity() == Yearly {
		slice := v.SliceThreshold()
		_ = slice
		return &Iter{}
	} else {
		return &Iter{err: errors.New("Query not supported")}
	}
}

const create = `
	CREATE TABLE IF NOT EXISTS x (
		id INT PRIMARY KEY,
		version TEXT,
		ts TIMESTAMP DATETIME DEFAULT(STRFTIME('%Y-%m-%d %H:%M:%f', 'NOW'))
	);
`

const upsert = `
	INSERT OR REPLACE INTO x (id, version, ts)
	SELECT 
	    new.id, 
	    new.version, 
	    CASE 
	        WHEN old.version <> new.version 
	        THEN STRFTIME('%Y-%m-%d %H:%M:%f', 'NOW') 
	        ELSE COALESCE(old.ts, STRFTIME('%Y-%m-%d %H:%M:%f', 'NOW')) 
	    END AS ts
	FROM ( SELECT
	     ? AS id,
	     ? AS version
	) AS new
	LEFT JOIN (
	    SELECT id, version, ts
	    FROM x
	) AS old ON new.id = old.id;
`
