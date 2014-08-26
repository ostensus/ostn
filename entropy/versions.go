package entropy

// #include <sqlite3.h>
// #cgo LDFLAGS: -lsqlite3
import "C"
import (
	"database/sql"
	"errors"
	_ "github.com/mattn/go-sqlite3"
)

type VersionStore struct {
	db *sql.DB
}

func OpenStore() (*VersionStore, error) {

	db, err := sql.Open("sqlite3", "x.db")
	if err != nil {
		return nil, err
	}

	v := &VersionStore{
		db: db,
	}

	if err := autoload(); err != nil {
		return nil, err
	}

	_, err = db.Exec(create)
	if err != nil {
		return nil, err
	}

	return v, err
}

func (v *VersionStore) SliceThreshold() int {
	return 128
}

func (v *VersionStore) Accept(ev ChangeEvent) error {

	parted, ok := ev.(PartitionedEvent)
	if ok {
		st, err := v.db.Prepare(upsert)
		if err != nil {
			return err
		}
		_, err = st.Exec(parted.Id(), parted.Version())
		return err
	}

	return errors.New("Bogus event")
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