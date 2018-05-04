package meta

type SqlType int

const (
	_ SqlType = iota + 1
	Type_unknown

	Type_digit_start
	Type_tinyint
	Type_smallint
	Type_mediumint
	Type_bit
	Type_int
	Type_bigint
	Type_float
	Type_double
	Type_decimal
	Type_digit_end

	Type_date_start
	Type_date
	Type_time
	Type_year
	Type_datetime
	Type_timestamp
	Type_date_end

	Type_string_start
	Type_char
	Type_varchar
	Type_tinyblob
	Type_tinytext
	Type_blob
	Type_text
	Type_mediumblob
	Type_mediumtext
	Type_longblob
	Type_longtext
	Type_string_end
)

var SqlDict = map[string]SqlType{
	"unknown":    Type_unknown,
	"tinyint":    Type_tinyint,
	"smallint":   Type_smallint,
	"mediumint":  Type_mediumint,
	"bit":        Type_bit,
	"int":        Type_int,
	"bigint":     Type_bigint,
	"float":      Type_float,
	"double":     Type_double,
	"decimal":    Type_decimal,
	"date":       Type_date,
	"time":       Type_time,
	"year":       Type_year,
	"datetime":   Type_datetime,
	"timestamp":  Type_timestamp,
	"char":       Type_char,
	"varchar":    Type_varchar,
	"tinyblob":   Type_tinyblob,
	"tinytext":   Type_tinytext,
	"blob":       Type_blob,
	"text":       Type_text,
	"mediumblob": Type_mediumblob,
	"mediumtext": Type_mediumtext,
	"longblob":   Type_longblob,
	"longtext":   Type_longtext,
}

var ReSqlDict = map[SqlType]string{
	Type_unknown:    "unknown",
	Type_tinyint:    "tinyint",
	Type_smallint:   "smallint",
	Type_mediumint:  "mediumint",
	Type_bit:        "bit",
	Type_int:        "int",
	Type_bigint:     "bigint",
	Type_float:      "float",
	Type_double:     "double",
	Type_decimal:    "decimal",
	Type_date:       "date",
	Type_time:       "time",
	Type_year:       "year",
	Type_datetime:   "datetime",
	Type_timestamp:  "timestamp",
	Type_char:       "char",
	Type_varchar:    "varchar",
	Type_tinyblob:   "tinyblob",
	Type_tinytext:   "tinytext",
	Type_blob:       "blob",
	Type_text:       "text",
	Type_mediumblob: "mediumblob",
	Type_mediumtext: "mediumtext",
	Type_longblob:   "longblob",
	Type_longtext:   "longtext",
}
