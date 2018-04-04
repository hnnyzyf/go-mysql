package convert

//去掉所有的Zero前缀或者后缀
const (
	Raplace_prefix = iota + 1
	Raplace_suffix
	Raplace_bit_prefix
	Raplace_hex_prefix
)

func Replace0(str string, tag int) string {
	switch tag {
	case Raplace_prefix:
		return replace0InPrefix(str)
	case Raplace_suffix:
		return replace0InSuffix(str)
	default:
		return replace0InHex0Bit(str)
	}

}

func replace0InPrefix(str string) string {
	r := []rune(str)
	for index := 0; index < len(r); index++ {
		if r[index] != '0' {
			return string(r[index:])
		}
	}
	return "0"
}

func replace0InSuffix(str string) string {
	r := []rune(str)
	for index := len(r) - 1; index >= 0; index-- {
		if r[index] != '0' {
			return string(r[:index+1])
		}
	}
	return "0"
}

func replace0InHex0Bit(str string) string {
	r := []rune(str)
	for index := 2; index < len(r); index++ {
		if r[index] != '0' && r[index] != '\'' {
			return string(append(r[:2], r[index:]...))
		}
	}

	if r[1] == '\'' {
		return string(append(r[:2], r[len(r)-2:]...))
	} else {
		return string(append(r[:2], r[len(r)-1:]...))
	}
}
