SQL_TEMPLATE_DIR := entropy/tmpl
SQL_TEMPLATES := $(foreach dir, $(SQL_TEMPLATE_DIR), $(wildcard $(dir)/*))

MIGRATION_DIR := entropy/db
MIGRATION_SCRIPTS := $(foreach dir, $(MIGRATION_DIR), $(wildcard $(dir)/*))

test: entropy/versions.db

entropy/sql_tmpl.go: $(SQL_TEMPLATES)
	go-bindata -pkg=entropy -prefix=entropy/tmpl -o=$@ entropy/tmpl

entropy/migration:
	mkdir -p $@

entropy/migration/steps.go: entropy/migration $(MIGRATION_SCRIPTS)
	go-bindata -pkg=migration -o=$@ entropy/db

entropy/versions.db: entropy/migration/steps.go
	go run entropy/migrate_db.go
