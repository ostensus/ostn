package entropy

import (
	"database/sql"
	"errors"
	log "github.com/cihub/seelog"
	"time"
)

var s1 = []string{
	`
CREATE TABLE repositories ( 
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  name VARCHAR(100) UNIQUE 
);
`,
	`
CREATE TABLE unique_partition_names ( 
  repository INTEGER,
  name VARCHAR(100),
  PRIMARY KEY(repository, name),
  FOREIGN KEY(repository) REFERENCES repositories(id)
);
`,
	`
CREATE TABLE range_partitions ( 
  repository INTEGER,
  name VARCHAR(100),
  PRIMARY KEY(repository, name),
  FOREIGN KEY(repository, name) REFERENCES unique_partition_names(repository, name)
);
`,
}

var (
	ErrTableDoesNotExist = errors.New("table does not exist")
	ErrNoPreviousVersion = errors.New("no previous version found")
)

type MigrationRecord struct {
	VersionId int64
	TStamp    time.Time
	IsApplied bool
}

var steps = [][]string{
	s1,
}

func Migrate(db *sql.DB) error {

	current, err := EnsureDBVersion(db)
	if err != nil {
		return err
	}

	log.Infof("Current DB version: %d", current)

	for i, step := range steps {

		version := i + 1

		if version <= int(current) {
			continue
		}

		txn, err := db.Begin()
		if err != nil {
			return err
		}

		for _, stmt := range step {

			log.Infof("Step %d: Applying statement: %s", version, stmt)

			_, err = db.Exec(stmt)
			if err != nil {
				log.Error(err)
				return txn.Rollback()
			}
		}

		if _, err := txn.Exec(insertVersionSql(), version, true); err != nil {
			txn.Rollback()
			return err
		}

		err = txn.Commit()
		if err != nil {
			return err
		}

		log.Infof("Successfully migrated DB to version %d", version)

	}

	return nil
}

func EnsureDBVersion(db *sql.DB) (int64, error) {

	rows, err := dbVersionQuery(db)
	if err != nil {
		if err == ErrTableDoesNotExist {
			return 0, createVersionTable(db)
		}
		log.Error(err)
		return 0, err
	}
	defer rows.Close()

	// The most recent record for each migration specifies
	// whether it has been applied or rolled back.
	// The first version we find that has been applied is the current version.

	toSkip := make([]int64, 0)
	_ = toSkip

	for rows.Next() {
		var row MigrationRecord
		if err = rows.Scan(&row.VersionId, &row.IsApplied); err != nil {
			log.Criticalf("error scanning rows:", err)
		}

		// have we already marked this version to be skipped?
		skip := false
		for _, v := range toSkip {
			if v == row.VersionId {
				skip = true
				break
			}
		}

		// if version has been applied and not marked to be skipped, we're done
		if row.IsApplied && !skip {
			return row.VersionId, nil
		}

		// version is either not applied, or we've already seen a more
		// recent version of it that was not applied.
		if !skip {
			toSkip = append(toSkip, row.VersionId)
		}
	}

	panic("failure in EnsureDBVersion()")
}

func createVersionTable(db *sql.DB) error {
	txn, err := db.Begin()
	if err != nil {
		return err
	}

	if _, err := txn.Exec(createVersionTableSql()); err != nil {
		log.Errorf("failed here 2 %v", err)
		txn.Rollback()
		return err
	}

	version := 0
	applied := true
	if _, err := txn.Exec(insertVersionSql(), version, applied); err != nil {
		txn.Rollback()
		return err
	}

	return txn.Commit()
}

func createVersionTableSql() string {
	return `CREATE TABLE schema_version (
            	id INTEGER PRIMARY KEY AUTOINCREMENT,
                version_id BIGINT NOT NULL,
                is_applied BOOLEAN NOT NULL,
                tstamp TIMESTAMP DATETIME DEFAULT(STRFTIME('%Y-%m-%d %H:%M:%f', 'NOW'))
            );`
}

func insertVersionSql() string {
	return "INSERT INTO schema_version (version_id, is_applied) VALUES (?, ?);"
}

func dbVersionQuery(db *sql.DB) (*sql.Rows, error) {
	rows, err := db.Query("SELECT version_id, is_applied from schema_version ORDER BY id DESC")

	// For now, assume any error is because the table doesn't exist,
	// in which case we'll try to create it.
	if err != nil {
		return nil, ErrTableDoesNotExist
	}

	return rows, err
}
