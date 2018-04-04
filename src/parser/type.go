package parser


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
