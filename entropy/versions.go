package entropy

// #include <sqlite3.h>
// #cgo LDFLAGS: -lsqlite3
import "C"
import (
	"bytes"
	"database/sql"
	"errors"
	log "github.com/cihub/seelog"
	_ "github.com/mattn/go-sqlite3"
	"github.com/ostensus/ostn/entropy/migration"
	"github.com/relops/sqlc/sqlc"
	"text/template"
)

type VersionStore struct {
	db *sql.DB
}

var repoTmpl, digestTmpl, upsertTmpl *template.Template

func init() {
	m := template.FuncMap{
		"columnType": columnType,
	}

	repoBin, _ := repo_tmpl()
	repoTmpl = template.Must(template.New("repo.tmpl").Funcs(m).Parse(string(repoBin)))

	digestBin, _ := digest_tmpl()
	digestTmpl = template.Must(template.New("digest.tmpl").Funcs(m).Parse(string(digestBin)))

	upsertBin, _ := upsert_tmpl()
	upsertTmpl = template.Must(template.New("upsert.tmpl").Funcs(m).Parse(string(upsertBin)))
}

func OpenStore(path string) (*VersionStore, error) {

	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}

	v := &VersionStore{
		db: db,
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

	//sqlc.Select(...)

	rows, err := v.db.Query(metadata, repo)

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

	sql := renderSQL(repo, parts, digestTmpl)
	log.Infof("Parts: %+v", parts)
	log.Infof("Query: %s", sql)

	st, err := v.db.Prepare(sql)
	if err != nil {
		return nil, err
	}
	defer st.Close()

	rows, err = st.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	digests := make(map[string]string)

	for rows.Next() {
		var id string
		var digest string
		rows.Scan(&id, &digest)
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

const metadata = `
	SELECT u.name, s.value
	FROM unique_partition_names u
	LEFT OUTER JOIN range_partitions r ON u.repository = r.repository AND u.name = r.name
	LEFT OUTER JOIN set_partitions s ON u.repository = s.repository AND u.name = s.name
	WHERE u.repository = ?;
`

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
