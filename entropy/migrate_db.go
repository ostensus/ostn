// +build ignore

package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/ostensus/ostn/entropy/migration"
	"github.com/relops/sqlc/sqlc"
	"log"
	"os"
)

type opts struct {
	driver  string
	url     string
	dialect sqlc.Dialect
}

// TODO paramterize this path
var dbFile = "entropy/versions.db"

var sqlite = opts{
	driver:  "sqlite3",
	url:     dbFile,
	dialect: sqlc.Sqlite,
}

func main() {
	migrate(sqlite)
}

func migrate(o opts) {

	db, err := sql.Open(o.driver, o.url)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	steps := sqlc.LoadBindata(migration.AssetNames(), migration.Asset)
	err = sqlc.Migrate(db, o.dialect, steps)

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
