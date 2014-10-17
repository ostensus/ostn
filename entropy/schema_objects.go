// THIS FILE WAS AUTOGENERATED - ANY EDITS TO THIS WILL BE LOST WHEN IT IS REGENERATED
// AT 2014-10-18 00:23:31.827194689 +0100 BST USING sqlc VERSION 0.1.5



package entropy

import (
	"github.com/relops/sqlc/sqlc"
	"strings"
)



type repositories struct {
	
	ID sqlc.Int64Field
	
	SOURCE sqlc.StringField
	
	NAME sqlc.StringField
	
	alias string
}

func (t *repositories) IsSelectable() {}

func (t *repositories) Name() string {
	return "repositories"
}

func (t *repositories) As(a string) sqlc.Selectable {
	return &repositories{
		
		ID: t.ID,
		
		SOURCE: t.SOURCE,
		
		NAME: t.NAME,
		
		alias:a,
	}
}

func (t *repositories) Alias() string {
	return t.alias
}

func (t *repositories) MaybeAlias() string {
	if t.alias == "" {
		return "repositories"
	} else {
		return t.alias
	}
}

/////


func (t *repositories) StringField(name string) sqlc.StringField {
	return sqlc.String(REPOSITORIES, strings.ToUpper(name))
}

func (t *repositories) IntField(name string) sqlc.IntField {
	return sqlc.Int(REPOSITORIES, strings.ToUpper(name))
}

func (t *repositories) Int64Field(name string) sqlc.Int64Field {
	return sqlc.Int64(REPOSITORIES, strings.ToUpper(name))
}

func (t *repositories) TimeField(name string) sqlc.TimeField {
	return sqlc.Time(REPOSITORIES, strings.ToUpper(name))
}


/////

func (t *repositories) Fields() []sqlc.Field {
	return []sqlc.Field{
		
		sqlc.Int64(REPOSITORIES, "ID"),
		
		sqlc.String(REPOSITORIES, "SOURCE"),
		
		sqlc.String(REPOSITORIES, "NAME"),
		
	}
}



type unique_partition_names struct {
	
	REPOSITORY sqlc.Int64Field
	
	NAME sqlc.StringField
	
	alias string
}

func (t *unique_partition_names) IsSelectable() {}

func (t *unique_partition_names) Name() string {
	return "unique_partition_names"
}

func (t *unique_partition_names) As(a string) sqlc.Selectable {
	return &unique_partition_names{
		
		REPOSITORY: t.REPOSITORY,
		
		NAME: t.NAME,
		
		alias:a,
	}
}

func (t *unique_partition_names) Alias() string {
	return t.alias
}

func (t *unique_partition_names) MaybeAlias() string {
	if t.alias == "" {
		return "unique_partition_names"
	} else {
		return t.alias
	}
}

/////


func (t *unique_partition_names) StringField(name string) sqlc.StringField {
	return sqlc.String(UNIQUE_PARTITION_NAMES, strings.ToUpper(name))
}

func (t *unique_partition_names) IntField(name string) sqlc.IntField {
	return sqlc.Int(UNIQUE_PARTITION_NAMES, strings.ToUpper(name))
}

func (t *unique_partition_names) Int64Field(name string) sqlc.Int64Field {
	return sqlc.Int64(UNIQUE_PARTITION_NAMES, strings.ToUpper(name))
}

func (t *unique_partition_names) TimeField(name string) sqlc.TimeField {
	return sqlc.Time(UNIQUE_PARTITION_NAMES, strings.ToUpper(name))
}


/////

func (t *unique_partition_names) Fields() []sqlc.Field {
	return []sqlc.Field{
		
		sqlc.Int64(UNIQUE_PARTITION_NAMES, "REPOSITORY"),
		
		sqlc.String(UNIQUE_PARTITION_NAMES, "NAME"),
		
	}
}



type range_partitions struct {
	
	REPOSITORY sqlc.Int64Field
	
	NAME sqlc.StringField
	
	alias string
}

func (t *range_partitions) IsSelectable() {}

func (t *range_partitions) Name() string {
	return "range_partitions"
}

func (t *range_partitions) As(a string) sqlc.Selectable {
	return &range_partitions{
		
		REPOSITORY: t.REPOSITORY,
		
		NAME: t.NAME,
		
		alias:a,
	}
}

func (t *range_partitions) Alias() string {
	return t.alias
}

func (t *range_partitions) MaybeAlias() string {
	if t.alias == "" {
		return "range_partitions"
	} else {
		return t.alias
	}
}

/////


func (t *range_partitions) StringField(name string) sqlc.StringField {
	return sqlc.String(RANGE_PARTITIONS, strings.ToUpper(name))
}

func (t *range_partitions) IntField(name string) sqlc.IntField {
	return sqlc.Int(RANGE_PARTITIONS, strings.ToUpper(name))
}

func (t *range_partitions) Int64Field(name string) sqlc.Int64Field {
	return sqlc.Int64(RANGE_PARTITIONS, strings.ToUpper(name))
}

func (t *range_partitions) TimeField(name string) sqlc.TimeField {
	return sqlc.Time(RANGE_PARTITIONS, strings.ToUpper(name))
}


/////

func (t *range_partitions) Fields() []sqlc.Field {
	return []sqlc.Field{
		
		sqlc.Int64(RANGE_PARTITIONS, "REPOSITORY"),
		
		sqlc.String(RANGE_PARTITIONS, "NAME"),
		
	}
}



type set_partitions struct {
	
	REPOSITORY sqlc.Int64Field
	
	NAME sqlc.StringField
	
	VALUE sqlc.StringField
	
	alias string
}

func (t *set_partitions) IsSelectable() {}

func (t *set_partitions) Name() string {
	return "set_partitions"
}

func (t *set_partitions) As(a string) sqlc.Selectable {
	return &set_partitions{
		
		REPOSITORY: t.REPOSITORY,
		
		NAME: t.NAME,
		
		VALUE: t.VALUE,
		
		alias:a,
	}
}

func (t *set_partitions) Alias() string {
	return t.alias
}

func (t *set_partitions) MaybeAlias() string {
	if t.alias == "" {
		return "set_partitions"
	} else {
		return t.alias
	}
}

/////


func (t *set_partitions) StringField(name string) sqlc.StringField {
	return sqlc.String(SET_PARTITIONS, strings.ToUpper(name))
}

func (t *set_partitions) IntField(name string) sqlc.IntField {
	return sqlc.Int(SET_PARTITIONS, strings.ToUpper(name))
}

func (t *set_partitions) Int64Field(name string) sqlc.Int64Field {
	return sqlc.Int64(SET_PARTITIONS, strings.ToUpper(name))
}

func (t *set_partitions) TimeField(name string) sqlc.TimeField {
	return sqlc.Time(SET_PARTITIONS, strings.ToUpper(name))
}


/////

func (t *set_partitions) Fields() []sqlc.Field {
	return []sqlc.Field{
		
		sqlc.Int64(SET_PARTITIONS, "REPOSITORY"),
		
		sqlc.String(SET_PARTITIONS, "NAME"),
		
		sqlc.String(SET_PARTITIONS, "VALUE"),
		
	}
}



type x struct {
	
	ID sqlc.IntField
	
	VERSION sqlc.StringField
	
	TS sqlc.TimeField
	
	alias string
}

func (t *x) IsSelectable() {}

func (t *x) Name() string {
	return "x"
}

func (t *x) As(a string) sqlc.Selectable {
	return &x{
		
		ID: t.ID,
		
		VERSION: t.VERSION,
		
		TS: t.TS,
		
		alias:a,
	}
}

func (t *x) Alias() string {
	return t.alias
}

func (t *x) MaybeAlias() string {
	if t.alias == "" {
		return "x"
	} else {
		return t.alias
	}
}

/////


func (t *x) StringField(name string) sqlc.StringField {
	return sqlc.String(X, strings.ToUpper(name))
}

func (t *x) IntField(name string) sqlc.IntField {
	return sqlc.Int(X, strings.ToUpper(name))
}

func (t *x) Int64Field(name string) sqlc.Int64Field {
	return sqlc.Int64(X, strings.ToUpper(name))
}

func (t *x) TimeField(name string) sqlc.TimeField {
	return sqlc.Time(X, strings.ToUpper(name))
}


/////

func (t *x) Fields() []sqlc.Field {
	return []sqlc.Field{
		
		sqlc.Int(X, "ID"),
		
		sqlc.String(X, "VERSION"),
		
		sqlc.Time(X, "TS"),
		
	}
}




var __repositories = &repositories{}
var REPOSITORIES = &repositories {
	

	ID: sqlc.Int64(__repositories, "ID"),

	SOURCE: sqlc.String(__repositories, "SOURCE"),

	NAME: sqlc.String(__repositories, "NAME"),


}



var __unique_partition_names = &unique_partition_names{}
var UNIQUE_PARTITION_NAMES = &unique_partition_names {
	

	REPOSITORY: sqlc.Int64(__unique_partition_names, "REPOSITORY"),

	NAME: sqlc.String(__unique_partition_names, "NAME"),


}



var __range_partitions = &range_partitions{}
var RANGE_PARTITIONS = &range_partitions {
	

	REPOSITORY: sqlc.Int64(__range_partitions, "REPOSITORY"),

	NAME: sqlc.String(__range_partitions, "NAME"),


}



var __set_partitions = &set_partitions{}
var SET_PARTITIONS = &set_partitions {
	

	REPOSITORY: sqlc.Int64(__set_partitions, "REPOSITORY"),

	NAME: sqlc.String(__set_partitions, "NAME"),

	VALUE: sqlc.String(__set_partitions, "VALUE"),


}



var __x = &x{}
var X = &x {
	

	ID: sqlc.Int(__x, "ID"),

	VERSION: sqlc.String(__x, "VERSION"),

	TS: sqlc.Time(__x, "TS"),


}


