SQL_TEMPLATE_DIR := entropy/tmpl
SQL_TEMPLATES := $(foreach dir, $(SQL_TEMPLATE_DIR), $(wildcard $(dir)/*))

MIGRATION_DIR := entropy/db
MIGRATION_SCRIPTS := $(foreach dir, $(MIGRATION_DIR), $(wildcard $(dir)/*))

test: entropy/schema_objects.go
	go test ./... -v

entropy/sql_tmpl.go: $(SQL_TEMPLATES)
	go-bindata -pkg=entropy -prefix=entropy/tmpl -o=$@ entropy/tmpl

entropy/migration:
	mkdir -p $@

entropy/migration/steps.go: entropy/migration $(MIGRATION_SCRIPTS)
	go-bindata -pkg=migration -o=$@ entropy/db

entropy/versions.db: entropy/migration/steps.go
	go run entropy/migrate.go

entropy/schema_objects.go: entropy/versions.db
	sqlc -p entropy -o $@ -f entropy/versions.db -t sqlite