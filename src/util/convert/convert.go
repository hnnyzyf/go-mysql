package convert

import "strings"

//判断是否是数字开头
func isDigit(ch rune) bool {
	if ch >= '0' && ch <= '9' {
		return true
	} else {
		return false
	}
}

//实现mysql的隐式转换

func String2Number(str string) string {
	r := []rune(str)
	for index := 1; index < len(r)-1; index++ {
		if !isDigit(r[index]) {
			return Replace0(string(r[1:index]), Raplace_prefix)
		}
	}
	return Replace0(string(r[1:len(r)-1]), Raplace_prefix)
}


func RemoveQuote(str string) string {
	r := []rune(str)
	return string(r[1:len(r)-1])
}


func Removebacktick(str string) string {
	if str == "" {
		return strings.ToLower(str)
	}

	r := []rune(str)

	if len(r) <= 2 {
		return strings.ToLower(str)
	}

	if r[0] == '`' && r[len(r)-1] == '`' {
		return strings.ToLower(string(r[1 : len(r)-1]))
	}

	return strings.ToLower(str)
}
