package myparser

var Keywords = map[string]int{
	"ident":         IDENT,
	"string":        STRING,
	"number":        NUMBER,
	"as":            AS,
	"any":           ANY,
	"all":           ALL,
	"avg":           AVG,
	"asc":           ASC,
	"by":            BY,
	"count":         COUNT,
	"distinct":      DISTINCT,
	"desc":          DESC,
	"exists":        EXISTS,
	"from":          FROM,
	"for":           FOR,
	"false": 		 FALSE,
	"group":         GROUP,
	"having":        HAVING,
	"into":          INTO,
	"if":            IF,
	"limit":         LIMIT,
	"lock":          LOCK,
	"null":          NULL,
	"order":         ORDER,
	"offset":        OFFSET,
	"select":        SELECT,
	"sum": 			 SUM,
	"to":            TO,
	"true":          TRUE,
	"update":        UPDATE,
	"where":         WHERE,
	"union":         UNION,
	"using": 		 USING,
	"except":        EXCEPT,
	"intersect":     INTERSECT,
	"join":          JOIN,
	"straight_join": STRAIGHT_JOIN,
	"share":		 SHARE,
	"some":          SOME,
	"left":          LEFT,
	"right":         RIGHT,
	"inner":         INNER,
	"outer":         OUTER,
	"cross":         CROSS,
	"natural":       NATURAL,
	"use":           USE,
	"on":            ON,
	"or":            OR,
	"and":           AND,
	"not":           NOT,
	"between":       BETWEEN,
	"case":          CASE,
	"when":          WHEN,
	"then":          THEN,
	"else":          ELSE,
	"is":            IS,
	"like":          LIKE,
	"in":            IN,
}

func IsDigit(ch rune) bool {
	if ch >= '0' && ch <= '9' {
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

func IsSymbol(ch rune) bool {
	if ch == '@' || ch == '$' || ch == '^' || ch == '#' || ch == '&' || ch == '?' || ch == '_' || ch == '~' {
		return true
	} else {
		return false
	}
}
func IsBlank(ch rune) bool {
	if ch == ' ' || ch == '\n' || ch == '\r' || ch == '\t' {
		return true
	} else {
		return false
	}
}

func IsDot(ch rune) bool {
	if ch == '.' {
		return true
	} else {
		return false
	}
}

func IsBinaryOP(ch rune) bool {
	if ch == '+' || ch == '-' || ch == '*' || ch == '\\' || ch == '%' {
		return true
	} else {
		return false
	}
}

func IsLogicalOP(ch rune) bool {
	if ch == '!' || ch == '<' || ch == '=' || ch == '>' {
		return true
	} else {
		return false
	}
}

func IsQuota(ch rune) bool {
	if ch == '\'' || ch == '"' || ch == '`' {
		return true
	} else {
		return false
	}
}

func IsSeparator(ch rune) bool {
	if ch == ',' || ch == ';' || ch == '(' || ch == ')' {
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
