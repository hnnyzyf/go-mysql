package myparser

import "strings"
import "bytes"
import AST "myast"

//实现mysql的此法解析器
//定义关键子

//需要实现Lex方法和error方法
type Tokener struct {
	Scanner   *strings.Reader
	Buf       *bytes.Buffer
	Position  int //输出错误信息
	Lastchar  rune
	EOF       bool //输出语句结尾
	LastError string
	ASTree 	AST.SelectNode
}

func NewTokener(sql string) *Tokener {
	return &Tokener{Scanner: strings.NewReader(sql),
		Buf: bytes.NewBuffer([]byte{})}
}

//读取当前字符串的下一个字符
func (token *Tokener) next() {
	ch, _, err := token.Scanner.ReadRune()
	if err != nil {
		token.EOF = true
	} else {
		token.Lastchar = ch
	}
	token.Position++
}

func (token *Tokener) skip() {
	for IsBlank(token.Lastchar) && token.EOF == false {
		token.next()
	}
}

//scan函数，每次scan获得一个类型，返回类型和字符串
//调用scan的时候,进入switch的前提是token.EOF==false
func (token *Tokener) Scan() (int, string) {
	defer token.Buf.Reset()
	token.skip()
	if token.EOF == true {
		return -1, ""
	}
	if token.Position == 0 {
		token.next()
	}
	switch ch := token.Lastchar; {

	//判断是否是String
	case IsQuota(ch):
		return token.ScanString()
	case IsUnderLine(ch):
		return token.ScanUnderline()
	//判断是否是数字
	case IsDot(ch):
		return token.ScanDecimal()
	//判断是否ident
	case IsBackQuote(ch):
		return token.ScanBackQuote()

	//进行多段判断判断是否是指定的String
	case IsLetter(ch):
		switch ch = ch; {
		case IsHex(ch):
			return token.ScanHexAndBit()
		case IsBit(ch):
			return token.ScanHexAndBit()
		case IsNationalSet(ch):
			return token.ScanNationalSet()
		default:
			return token.ScanIdent()
		}
	//进行多段判断，判断是否是数字
	case IsDigit(ch):
		switch ch = ch; {
		case IsZero(ch):
			return token.ScanHexOrBit()
		default:
			return token.ScanNumber()
		}
	case IsOperator(ch):
		token.next()
		return int(ch), string(ch)
	case IsLogicalOP(ch):
		return token.ScanLogicalOp()
	case IsComment(ch):
		return token.scanComment()
	case IsQuestion(ch):
		token.next()
		return STRING,"?"
	default:
		token.next()
		return int(ch), string(ch)
	}
}

func (token *Tokener) ScanString() (int, string) {
	buffer := token.Buf
	ch := token.Lastchar
	buffer.WriteRune(token.Lastchar)
	token.next()

	for token.Lastchar != ch && token.EOF == false {
		buffer.WriteRune(token.Lastchar)
		token.next()
	}

	if token.EOF == true {
		return STRING, buffer.String()
	} else {
		buffer.WriteRune(token.Lastchar)
		token.next()
		//如果没有结束
		return STRING, buffer.String()
	}

}

func (token *Tokener) ScanUnderline() (int, string) {
	buffer := token.Buf

	buffer.WriteRune(token.Lastchar)
	token.next()
	//查看字母至少为一个
	if IsLetter(token.Lastchar) {
		buffer.WriteRune(token.Lastchar)
		token.next()
	} else {
		return int('_'), buffer.String()
	}

	//至少存在一个字母，遍历直到结束或者遇到'或者“
	for IsLetter(token.Lastchar) && token.EOF == false {
		buffer.WriteRune(token.Lastchar)
		token.next()
	}

	if IsQuota(token.Lastchar) {
		return token.ScanString()
	} else {
		return STRING, buffer.String()
	}
}

func (token *Tokener) ScanDecimal() (int, string) {
	buffer := token.Buf
	buffer.WriteRune(token.Lastchar)
	token.next()

	//判断是否是数字
	if IsDigit(token.Lastchar) {
		buffer.WriteRune(token.Lastchar)
		token.next()
	}else if IsDot(token.Lastchar){
		return token.ScanDot()
	}else{
		return int('.'), buffer.String()
	}

	for IsDigit(token.Lastchar) && token.EOF == false {
		buffer.WriteRune(token.Lastchar)
		token.next()
	}

	return NUMBER, buffer.String()
}


func (token *Tokener) ScanDot() (int, string) {
	buffer := token.Buf
	buffer.WriteRune(token.Lastchar)
	token.next()

	//判断是否是数字

	for IsDot(token.Lastchar) && token.EOF == false {
		buffer.WriteRune(token.Lastchar)
		token.next()
	}

	return STRING, buffer.String()
}


func (token *Tokener) ScanBackQuote() (int, string) {
	buffer := token.Buf
	ch := token.Lastchar
	buffer.WriteRune(token.Lastchar)
	token.next()
	for token.Lastchar != ch && token.EOF == false {
		buffer.WriteRune(token.Lastchar)
		token.next()
	}

	if token.EOF == true {
		return IDENT, buffer.String()
	} else {
		buffer.WriteRune(token.Lastchar)
		token.next()
		//如果没有结束
		return IDENT, buffer.String()
	}
}

//NationalSet判断
func (token *Tokener) ScanNationalSet() (int, string) {
	buffer := token.Buf
	buffer.WriteRune(token.Lastchar)
	token.next()

	//判断下一个是否是''
	if !IsQuota(token.Lastchar) {
		return token.ScanIdent()
	}

	//进入字符串判断
	_, NationalSet := token.ScanString()

	return STRING, NationalSet
}

//十六进制和bi判断
func (token *Tokener) ScanHexAndBit() (int, string) {
	buffer := token.Buf
	ch := token.Lastchar
	buffer.WriteRune(token.Lastchar)
	token.next()
	var quote rune
	//判断下一个是否是'',不是的话判断是否是Ident
	if IsQuota(token.Lastchar) {
		quote = token.Lastchar
		buffer.WriteRune(token.Lastchar)
		token.next()
	} else if IsIdentLetter(token.Lastchar){
		return token.ScanIdent()
	}else{
		return IDENT,buffer.String()
	}

	StringType := NUMBER
	//进入string判断
	switch ch = ch; {
	case IsHex(ch):
		for token.EOF == false && token.Lastchar != quote {
			if !IsHexLetter(token.Lastchar) {
				StringType = IDENT
			}
			buffer.WriteRune(token.Lastchar)
			token.next()

		}
	case IsBit(ch):
		for token.EOF == false && token.Lastchar != quote {
			if IsBlank(token.Lastchar) {
				break
			}
			if !IsBitLetter(token.Lastchar) {
				StringType = IDENT
			}
			buffer.WriteRune(token.Lastchar)
			token.next()
		}
	}

	if token.EOF == true {
		return StringType, buffer.String()
	} else {
		buffer.WriteRune(token.Lastchar)
		token.next()
		//如果没有结束
		return StringType, buffer.String()
	}
}

//判断Hex和bit的另一种写法
func (token *Tokener) ScanHexOrBit() (int, string) {
	buffer := token.Buf
	buffer.WriteRune(token.Lastchar)
	token.next()
	//分两种两筐
	switch token.Lastchar {
	//进行16进制判断
	case 'x':
		buffer.WriteRune(token.Lastchar)
		token.next()
		for token.EOF == false && !IsBlank(token.Lastchar) {
			if !IsHexLetter(token.Lastchar) {
				return NUMBER, buffer.String()
			}
			buffer.WriteRune(token.Lastchar)
			token.next()

		}
	case 'b':
		buffer.WriteRune(token.Lastchar)
		token.next()
		for token.EOF == false && !IsBlank(token.Lastchar) {
			if !IsBitLetter(token.Lastchar) {
				return NUMBER, buffer.String()
			}
			buffer.WriteRune(token.Lastchar)
			token.next()
		}
	default:
		//判断是否是数字
		switch ch:=token.Lastchar;{
			case IsDigit(ch):
				return token.ScanNumber()
			default:
				return NUMBER,buffer.String()
		}
	}
	return NUMBER, buffer.String()

}


//判断Hex和bit的另一种写法
func (token *Tokener) ScanNumber() (int, string) {
	buffer := token.Buf
	buffer.WriteRune(token.Lastchar)
	token.next()

	//第一次循环
	for IsDigit(token.Lastchar) && token.EOF == false {
		buffer.WriteRune(token.Lastchar)
		token.next()
	}

	if !IsDot(token.Lastchar) {
		return NUMBER, buffer.String()
	} else {
		return token.ScanDecimal()
	}

}


//判断Hex和bit的另一种写法
func (token *Tokener) ScanLogicalOp() (int, string) {
	buffer := token.Buf
	ch := token.Lastchar
	token.next()

	if token.EOF == true {
		return int(ch), buffer.String()
	}

	switch ch {
	case '!':
		if token.Lastchar == '=' {
			token.next()
			return LNE, "!="
		} else {
			return int('!'), "!"
		}
	case '<':
		switch token.Lastchar {
		case '<':
			token.next()
			return SL, "<<"
		case '>':
			token.next()
			return NE, "<>"
		case '=':
			token.next()
			if token.EOF == true {
				return LE, "<="
			}
			if token.Lastchar == '>' {
				token.next()
				return LEG, "<=>"
			} else {
				return LE, "<="
			}
		default:
			return int('<'), "<"
		}
	case '>':
		switch token.Lastchar {
		case '>':
			token.next()
			return SR, ">>"
		case '=':
			token.next()
			return GE, ">="
		default:
			return int('>'), ">"
		}
	case '&':
		if token.Lastchar == '&' {
			token.next()
			return AA, "&&"
		} else {
			return int('&'), "&"
		}
	case '|':
		if token.Lastchar == '|' {
			token.next()
			return OO, "||"
		} else {
			return int('|'), "|"
		}
	default:
		return int(ch), string(ch)
	}
}

func (token *Tokener) ScanIdent() (int, string) {
	buffer := token.Buf
	buffer.WriteRune(token.Lastchar)
	token.next()

	for token.EOF == false && IsIdentLetter(token.Lastchar){
		buffer.WriteRune(token.Lastchar)
		token.next()
	}

	//判断是否是关键字
	ident := buffer.String()
	if _, ok := Keywords[strings.ToLower(ident)]; ok {
		return Keywords[strings.ToLower(ident)], strings.ToUpper(ident)
	} else {
		return IDENT, ident
	}
}



func (token *Tokener) scanComment() (int, string) {
	buffer := token.Buf
	buffer.WriteRune(token.Lastchar)
	token.next()
	//判断下一个字符是否是*或者已经遍历结束
	if token.Lastchar == '*' {
		buffer.WriteRune(token.Lastchar)
		token.next()
	} else {
		return int('/'), buffer.String()
	}

comment:
	//继续遍历直到遇到*或者结束
	for token.Lastchar != '*' && token.EOF == false {
		buffer.WriteRune(token.Lastchar)
		token.next()
	}

	//判断如果遇到的是*并且还没结束
	switch token.EOF {
	case true:
		buffer.WriteRune(token.Lastchar)
		return COMMENT, buffer.String()
	case false:
		buffer.WriteRune(token.Lastchar)
		token.next()
		if token.Lastchar == '/' || token.EOF == true {
			buffer.WriteRune(token.Lastchar)
			token.next()
			return COMMENT, buffer.String()
		} else {
			goto comment
		}
	}

	return COMMENT, buffer.String()

}