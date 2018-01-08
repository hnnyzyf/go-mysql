package myparser

import "strings"
import "bytes"

//实现mysql的此法解析器
//定义关键子

//需要实现Lex方法和error方法
type Tokener struct {
	Scanner  *strings.Reader
	Buf      *bytes.Buffer
	Position int //输出错误信息
	Lastchar rune
	EOF      bool //输出语句结尾
	AST      string
	LastError string
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
	if token.EOF == true {
		return -1, ""
	}
	if token.Position == 0 {
		token.next()
	}
	token.skip()
	switch ch := token.Lastchar; {
	//判断是否是数字
	case IsDigit(ch):
		return token.scanNumber()
	case IsLetter(ch):
		return token.scanIdent()
	case IsQuota(ch):
		return token.scanString()
	case IsBinaryOP(ch):
		defer token.next()
		return int(ch), string(ch)
	case IsLogicalOP(ch):
		return token.scanOpertator()
	case IsSeparator(ch):
		defer token.next()
		return int(ch), string(ch)
	case IsDot(ch):
		defer token.next()
		return int(ch), string(ch)
	case IsComment(ch):
		return token.scanComment()
	}
	return -1, ""
}

func (token *Tokener) scanNumber() (int, string) {
	buffer := token.Buf
	defer buffer.Reset()
	buffer.WriteRune(token.Lastchar)
	token.next()

	//第一次跳出
	for IsDigit(token.Lastchar) && token.EOF == false {
		buffer.WriteRune(token.Lastchar)
		token.next()
	}

	//判断是否为小数
	if IsDot(token.Lastchar) {
		buffer.WriteRune(token.Lastchar)
		token.next()
	} else {
		return NUMBER, buffer.String()
	}

	//继续
	for IsDigit(token.Lastchar) && token.EOF == false {
		buffer.WriteRune(token.Lastchar)
		token.next()
	}

	return NUMBER, buffer.String()
}

func (token *Tokener) scanIdent() (int, string) {
	buffer := token.Buf
	defer buffer.Reset()
	buffer.WriteRune(token.Lastchar)
	token.next()

	for (IsSymbol(token.Lastchar) || IsLetter(token.Lastchar) || IsDigit(token.Lastchar)) && token.EOF == false {
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

func (token *Tokener) scanString() (int, string) {
	buffer := token.Buf
	defer buffer.Reset()
	ch := token.Lastchar
	buffer.WriteRune(token.Lastchar)
	token.next()
	for token.Lastchar != ch && token.EOF == false {
		buffer.WriteRune(token.Lastchar)
		token.next()
	}
	//如果没有结束
	if token.EOF == false {
		buffer.WriteRune(token.Lastchar)
		token.next()
	}
	if ch=='`'{
		return IDENT, buffer.String()
	}else{
		return STRING,buffer.String()
	}
	
}

func (token *Tokener) scanOpertator() (int, string) {
	buffer := token.Buf
	buffer.WriteRune(token.Lastchar)
	defer buffer.Reset()
	switch token.Lastchar {
	case '!':
		token.next()
		if token.Lastchar == '=' {
			buffer.WriteRune(token.Lastchar)
			token.next()
			return LNE, buffer.String()
		} else {
			return STRING, buffer.String()
		}
	case '=':
		switch token.EOF {
		case true:
			return int('='), buffer.String()
		case false:
			token.next()
			if token.Lastchar == '=' && token.EOF == false {
				buffer.WriteRune(token.Lastchar)
				token.next()
				return IE, buffer.String()
			} else {
				return int('='), buffer.String()
			}

		}
	case '>':
		switch token.EOF {
		case true:
			return int('>'), buffer.String()
		case false:
			token.next()
			if token.Lastchar == '=' {
				buffer.WriteRune(token.Lastchar)
				token.next()
				return GE, buffer.String()
			} else {
				return int('>'), buffer.String()
			}

		}
	case '<':
		switch token.EOF {
		case true:
			return int('<'), buffer.String()
		case false:
			token.next()
			switch token.Lastchar {
			case '=':
				buffer.WriteRune(token.Lastchar)
				token.next()
				return LE, buffer.String()
			case '>':
				buffer.WriteRune(token.Lastchar)
				token.next()
				return NE, buffer.String()
			default:
				return int('<'), buffer.String()
			}

		}

	}
	return -1, buffer.String()
}

func (token *Tokener) scanComment() (int, string) {
	buffer := token.Buf
	defer buffer.Reset()
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

