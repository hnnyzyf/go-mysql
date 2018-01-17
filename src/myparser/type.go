package myparser

var Keywords = map[string]int{
	"all": ALL,
	"and": AND,
	"any": ANY,
	"as":  AS,
	"asc": ASC,
	//"avg":           		AVG,
	"between": BETWEEN,
	"by":      BY,
	"case":    CASE,
	//"count":         		COUNT,
	"collate":            COLLATE,
	"cross":              CROSS,
	"character":          CHARACTER,
	"day":                DAY,
	"day_hour":           DAY_HOUR,
	"day_microsecond":    DAY_MICROSECOND,
	"day_minute":         DAY_MINUTE,
	"day_second":         DAY_SECOND,
	"desc":               DESC,
	"distinct":           DISTINCT,
	"distinctrow":        DISTINCTROW,
	"end":                END,
	"else":               ELSE,
	"except":             EXCEPT,
	"exists":             EXISTS,
	"false":              FALSE,
	"for":                FOR,
	"from":               FROM,
	"group":              GROUP,
	"having":             HAVING,
	"hour":               HOUR,
	"hour_microsecond":   HOUR_MICROSECOND,
	"hour_minute":        HOUR_MINUTE,
	"hour_second":        HOUR_SECOND,
	"ident":              IDENT,
	"if":                 IF,
	"in":                 IN,
	"inner":              INNER,
	"intersect":          INTERSECT,
	"interval":           INTERVAL,
	"into":               INTO,
	"is":                 IS,
	"join":               JOIN,
	"left":               LEFT,
	"like":               LIKE,
	"limit":              LIMIT,
	"lock":               LOCK,
	"microsecond":        MICROSECOND,
	"minute":             MINUTE,
	"minute_microsecond": MINUTE_MICROSECOND,
	"minute_second":      MINUTE_SECOND,
	"mode":               MODE,
	"month":              MONTH,
	"natural":            NATURAL,
	"not":                NOT,
	"null":               NULL,
	"number":             NUMBER,
	"offset":             OFFSET,
	"on":                 ON,
	"or":                 OR,
	"order":              ORDER,
	"outer":              OUTER,
	"quarter":            QUARTER,
	"right":              RIGHT,
	"second":             SECOND,
	"second_microsecond": SECOND_MICROSECOND,
	"set":                SET,
	"select":             SELECT,
	"share":              SHARE,
	"some":               SOME,
	"straight_join":      STRAIGHT_JOIN,
	"string":             STRING,
	//"sum":           		SUM,
	"then":       THEN,
	"to":         TO,
	"true":       TRUE,
	"union":      UNION,
	"update":     UPDATE,
	"use":        USE,
	"using":      USING,
	"when":       WHEN,
	"where":      WHERE,
	"week":       WEEK,
	"year":       YEAR,
	"year_month": YEAR_MONTH,
}

//判断是否是String类型的前缀
func IsStringPrefix(ch rune) bool {
	if IsQuota(ch) || IsNationalSet(ch) || IsUnderLine(ch) {
		return true
	} else {
		return false
	}
}

func IsQuestion(ch rune) bool {
	if ch == '?' {
		return true
	} else {
		return false
	}
}

//判断引号
func IsQuota(ch rune) bool {
	if ch == '\'' || ch == '"' {
		return true
	} else {
		return false
	}
}

//判断national character set
func IsNationalSet(ch rune) bool {
	if ch == 'N' || ch == 'n' {
		return true
	} else {
		return false
	}
}

//判断_开头的
func IsUnderLine(ch rune) bool {
	if ch == '_' {
		return true
	} else {
		return false
	}
}

//判断是否是数字的前缀
func IsNumberPrefix(ch rune) bool {
	if IsDot(ch) || IsDigit(ch) || IsHex(ch) || IsBit(ch) {
		return true
	} else {
		return false
	}
}

//判断是否是带.开始的小数
func IsDot(ch rune) bool {
	if ch == '.' {
		return true
	} else {
		return false
	}
}

//判断是否是数字开头
func IsDigit(ch rune) bool {
	if ch >= '0' && ch <= '9' {
		return true
	} else {
		return false
	}
}

//判断是否是十六进制
func IsHex(ch rune) bool {
	if ch == 'X' || ch == 'x' {
		return true
	} else {
		return false
	}
}

func IsHexLetter(ch rune) bool {
	if IsDigit(ch) || (ch >= 'a' && ch <= 'f') || (ch >= 'A' && ch <= 'F') {
		return true
	} else {
		return false
	}
}

//判断是否是Bit
func IsBit(ch rune) bool {
	if ch == 'b' || ch == 'B' {
		return true
	} else {
		return false
	}
}

func IsBitLetter(ch rune) bool {
	if ch == '0' || ch == '1' {
		return true
	} else {
		return false
	}
}

//判断是否是0开头的
func IsZero(ch rune) bool {
	if ch == '0' {
		return true
	} else {
		return false
	}
}

//判断是否是保留字类型类的前缀

func IsIdentPrefix(ch rune) bool {
	if IsBackQuote(ch) || IsLetter(ch) {
		return true
	} else {
		return false
	}
}

func IsIdentLetter(ch rune) bool {
	if IsLetter(ch) || IsDigit(ch) || IsUnderLine(ch) {
		return true
	} else {
		return false
	}
}

func IsBackQuote(ch rune) bool {
	if ch == '`' {
		return true
	} else {
		return false
	}
}

func IsLetter(ch rune) bool {
	if (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') {
		return true
	} else {
		return false
	}
}

//判断是否是空格
func IsBlank(ch rune) bool {
	if ch == ' ' || ch == '\n' || ch == '\r' || ch == '\t' {
		return true
	} else {
		return false
	}
}

func IsOperator(ch rune) bool {
	if ch == '*' || ch == '%' || ch == '^' || ch == '~' || ch == '=' {
		return true
	} else {
		return false
	}
}

func IsLogicalOP(ch rune) bool {
	if ch == '!' || ch == '<' || ch == '>' || ch == '&' || ch == '|' || ch == '!' {
		return true
	} else {
		return false
	}
}

func IsComment(ch rune) bool {
	if ch == '/' {
		return true
	} else {
		return false
	}
}
