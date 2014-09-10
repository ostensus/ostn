SQL_TEMPLATE_DIR := entropy/tmpl
SQL_TEMPLATES := $(foreach dir, $(SQL_TEMPLATE_DIR), $(wildcard $(dir)/*))

entropy/sql_tmpl.go: $(SQL_TEMPLATES)
	go-bindata -pkg=entropy -prefix=entropy/tmpl -o=$@ entropy/tmpl
