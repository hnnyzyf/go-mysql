package parser

var Keywords = map[string]int{
	"all":                ALL,
	"and":                AND,
	"any":                ANY,
	"as":                 AS,
	"asc":                ASC,
	"avg":                AVG,
	"between":            BETWEEN,
	"by":                 BY,
	"case":               CASE,
	"character":          CHARACTER,
	"collate":            COLLATE,
	"count":              COUNT,
	"cross":              CROSS,
	"day":                DAY,
	"day_hour":           DAY_HOUR,
	"day_microsecond":    DAY_MICROSECOND,
	"day_minute":         DAY_MINUTE,
	"day_second":         DAY_SECOND,
	"desc":               DESC,
	"distinct":           DISTINCT,
	"distinctrow":        DISTINCTROW,
	"div":                DIV,
	"else":               ELSE,
	"end":                END,
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
	"max":                MAX,
	"microsecond":        MICROSECOND,
	"min":                MIN,
	"minute":             MINUTE,
	"minute_microsecond": MINUTE_MICROSECOND,
	"minute_second":      MINUTE_SECOND,
	"mod":                MOD,
	"mode":               MODE,
	"month":              MONTH,
	"natural":            NATURAL,
	"not":                NOT,
	"null":               NULL,
	"offset":             OFFSET,
	"on":                 ON,
	"or":                 OR,
	"order":              ORDER,
	"outer":              OUTER,
	"quarter":            QUARTER,
	"regexp":             REGEXP,
	"right":              RIGHT,
	"second":             SECOND,
	"second_microsecond": SECOND_MICROSECOND,
	"select":             SELECT,
	"set":                SET,
	"share":              SHARE,
	"some":               SOME,
	"sounds":             SOUNDS,
	"sum":                SUM,
	"straight_join":      STRAIGHT_JOIN,
	"then":               THEN,
	"to":                 TO,
	"true":               TRUE,
	"union":              UNION,
	"update":             UPDATE,
	"use":                USE,
	"using":              USING,
	"week":               WEEK,
	"when":               WHEN,
	"where":              WHERE,
	"xor":                XOR,
	"year":               YEAR,
	"year_month":         YEAR_MONTH,
}

var Character = map[string]int{
	"armscii8_general_ci": armscii8_general_ci,
	"ascii_general_ci":    ascii_general_ci,
	"big5_chinese_ci":     big5_chinese_ci,
	"binary":              binary,
	"cp1250_general_ci":   cp1250_general_ci,
	"cp1251_general_ci":   cp1251_general_ci,
	"cp1256_general_ci":   cp1256_general_ci,
	"cp1257_general_ci":   cp1257_general_ci,
	"cp850_general_ci":    cp850_general_ci,
	"cp852_general_ci":    cp852_general_ci,
	"cp866_general_ci":    cp866_general_ci,
	"cp932_japanese_ci":   cp932_japanese_ci,
	"dec8_swedish_ci":     dec8_swedish_ci,
	"eucjpms_japanese_ci": eucjpms_japanese_ci,
	"euckr_korean_ci":     euckr_korean_ci,
	"gb18030_chinese_ci":  gb18030_chinese_ci,
	"gb2312_chinese_ci":   gb2312_chinese_ci,
	"gbk_chinese_ci":      gbk_chinese_ci,
	"geostd8_general_ci":  geostd8_general_ci,
	"greek_general_ci":    greek_general_ci,
	"hebrew_general_ci":   hebrew_general_ci,
	"hp8_english_ci":      hp8_english_ci,
	"keybcs2_general_ci":  keybcs2_general_ci,
	"koi8r_general_ci":    koi8r_general_ci,
	"koi8u_general_ci":    koi8u_general_ci,
	"latin1_swedish_ci":   latin1_swedish_ci,
	"latin2_general_ci":   latin2_general_ci,
	"latin5_turkish_ci":   latin5_turkish_ci,
	"latin7_general_ci":   latin7_general_ci,
	"macce_general_ci":    macce_general_ci,
	"macroman_general_ci": macroman_general_ci,
	"sjis_japanese_ci":    sjis_japanese_ci,
	"swe7_swedish_ci":     swe7_swedish_ci,
	"tis620_thai_ci":      tis620_thai_ci,
	"ucs2_general_ci":     ucs2_general_ci,
	"ujis_japanese_ci":    ujis_japanese_ci,
	"utf16_general_ci":    utf16_general_ci,
	"utf16le_general_ci":  utf16le_general_ci,
	"utf32_general_ci":    utf32_general_ci,
	"utf8_general_ci":     utf8_general_ci,
	"utf8mb4_general_ci":  utf8mb4_general_ci,
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

func IsAt(ch rune) bool {
	if ch == '@' {
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
