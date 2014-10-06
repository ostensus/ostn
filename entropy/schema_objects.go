// THIS FILE WAS AUTOGENERATED - ANY EDITS TO THIS WILL BE LOST WHEN IT IS REGENERATED
package entropy

import (
	"github.com/relops/sqlc/sqlc"
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
	t.alias = a
	return t
}

func (t *repositories) Alias() string {
	return t.alias
}

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
	t.alias = a
	return t
}

func (t *unique_partition_names) Alias() string {
	return t.alias
}

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
	t.alias = a
	return t
}

func (t *range_partitions) Alias() string {
	return t.alias
}

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
	t.alias = a
	return t
}

func (t *set_partitions) Alias() string {
	return t.alias
}

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
	t.alias = a
	return t
}

func (t *x) Alias() string {
	return t.alias
}

func (t *x) Fields() []sqlc.Field {
	return []sqlc.Field{
		
		sqlc.Int(X, "ID"),
		
		sqlc.String(X, "VERSION"),
		
		sqlc.Time(X, "TS"),
		
	}
}



type v_1 struct {
	
	ID sqlc.StringField
	
	VERSION sqlc.StringField
	
	_TSTAMP sqlc.TimeField
	
	TS_COL sqlc.TimeField
	
	X sqlc.StringField
	
	alias string
}

func (t *v_1) IsSelectable() {}

func (t *v_1) Name() string {
	return "v_1"
}

func (t *v_1) As(a string) sqlc.Selectable {
	t.alias = a
	return t
}

func (t *v_1) Alias() string {
	return t.alias
}

func (t *v_1) Fields() []sqlc.Field {
	return []sqlc.Field{
		
		sqlc.String(V_1, "ID"),
		
		sqlc.String(V_1, "VERSION"),
		
		sqlc.Time(V_1, "_TSTAMP"),
		
		sqlc.Time(V_1, "TS_COL"),
		
		sqlc.String(V_1, "X"),
		
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



var __v_1 = &v_1{}
var V_1 = &v_1 {
	

	ID: sqlc.String(__v_1, "ID"),

	VERSION: sqlc.String(__v_1, "VERSION"),

	_TSTAMP: sqlc.Time(__v_1, "_TSTAMP"),

	TS_COL: sqlc.Time(__v_1, "TS_COL"),

	X: sqlc.String(__v_1, "X"),


}


