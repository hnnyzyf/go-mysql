package meta

import "strings"

type SqlType int

const (
	_ SqlType = iota
	type_digit_start
	type_tinyint
	type_smallint
	type_mediumint
	type_int
	type_bigint
	type_float
	type_double
	type_decimal
	type_digit_end

	type_date_start
	type_date
	type_time
	type_year
	type_datetime
	type_timestamp
	type_date_end

	type_string_start
	type_char
	type_varchar
	type_tinyblob
	type_tinytext
	type_blob
	type_text
	type_mediumblob
	type_mediumtext
	type_longblob
	type_longtext
	type_string_end
)

var SqlDict = map[string]SqlType{
	"tinyint":    type_tinyint,
	"smallint":   type_smallint,
	"mediumint":  type_mediumint,
	"int":        type_int,
	"bigint":     type_bigint,
	"float":      type_float,
	"double":     type_double,
	"decimal":    type_decimal,
	"date":       type_date,
	"time":       type_time,
	"year":       type_year,
	"datetime":   type_datetime,
	"timestamp":  type_timestamp,
	"char":       type_char,
	"varchar":    type_varchar,
	"tinyblob":   type_tinyblob,
	"tinytext":   type_tinytext,
	"blob":       type_blob,
	"text":       type_text,
	"mediumblob": type_mediumblob,
	"mediumtext": type_mediumtext,
	"longblob":   type_longblob,
	"longtext":   type_longtext,
}

var ReSqlDict = map[SqlType]string{
	type_tinyint:    "tinyint",
	type_smallint:   "smallint",
	type_mediumint:  "mediumint",
	type_int:        "int",
	type_bigint:     "bigint",
	type_float:      "float",
	type_double:     "double",
	type_decimal:    "decimal",
	type_date:       "date",
	type_time:       "time",
	type_year:       "year",
	type_datetime:   "datetime",
	type_timestamp:  "timestamp",
	type_char:       "char",
	type_varchar:    "varchar",
	type_tinyblob:   "tinyblob",
	type_tinytext:   "tinytext",
	type_blob:       "blob",
	type_text:       "text",
	type_mediumblob: "mediumblob",
	type_mediumtext: "mediumtext",
	type_longblob:   "longblob",
	type_longtext:   "longtext",
}

func getsqltype(t string) SqlType {
	s := strings.Split(t, "(")
	return SqlDict[s[0]]
}
