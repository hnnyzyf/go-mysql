package parser

import "strings"
import "bytes"
import "github.com/hnnyzyf/mysql-go/advisor/ast"
import "github.com/hnnyzyf/mysql-go/advisor/util/conv"

//实现mysql的此法解析器
//定义关键子
type Tokener struct {
	Scanner   *strings.Reader
	Buf       *bytes.Buffer
	Position  int //输出错误信息
	Lastchar  rune
	EOF       bool //输出语句结尾 false变为true表示已经结束
	LastError string
	Comment   []string
	step      stepFunc
	ParseStmt ast.Node
}

type stepFunc func(token *Tokener, ch rune) (int, string)

func NewTokener(sql string) *Tokener {
	return &Tokener{
		Scanner:   strings.NewReader(sql),
		Buf:       bytes.NewBuffer([]byte{}),
		Comment:   []string{},
		ParseStmt: nil,
	}
}

func (token *Tokener) Stmt() ast.Node {
	return token.ParseStmt
}

func (token *Tokener) Reset(sql string) {
	token.Scanner = strings.NewReader(sql)
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

func (token *Tokener) doubleRune(ch rune) bool {
	buffer := token.Buf
	buffer.WriteRune(token.Lastchar)
	token.next()
	//判断是否结束
	if token.EOF == true {
		return false
	}
	//判断是否一致
	if token.Lastchar == ch {
		buffer.WriteRune(ch)
		token.next()
		return true
	} else {
		return false
	}
}

func checkIdent(s string) (int, string) {
	ident := strings.ToLower(s)
	if _, ok := Keywords[ident]; ok {
		return Keywords[ident], strings.ToUpper(ident)
	} else if _, ok := Character[ident]; ok {
		return Character[ident], ident
	} else if _, ok := TimeUnit[ident]; ok {
		return TimeUnit[ident], strings.ToUpper(ident)
	} else if _, ok := FuncTime[ident]; ok {
		return FuncTime[ident], strings.ToUpper(ident)
	} else {
		return IDENT, s
	}
}

//scan函数，每次scan获得一个类型，返回类型和字符串
func (token *Tokener) Scan() (int, string) {
	if token.Position == 0 {
		token.next()
	}

	//跳过所有空格
	token.skip()

	//进入switch的前提是EOF=false
	if token.EOF == true {
		return -1, ""
	}

	//有step函数，进入step函数
	if token.step != nil {
		tp, s := token.step(token, '\'')
		token.step = nil
		return tp, s
	}

	token.Buf.Reset()
	//进入所有scan函数,需要判断是否已经到字符串结尾
	switch ch := token.Lastchar; {
	//判断是否是String
	case IsQuota(ch):
		return token.ScanString(ch)
	//判断是否是Ident
	case IsUnderLine(ch):
		return token.ScanUnderline()
	case IsBackQuote(ch):
		return token.ScanBackQuote()

	//起始是字母的，分情况讨论
	case IsLetter(ch):
		switch ch = ch; {
		case IsHex(ch) || IsBit(ch):
			return token.ScanHexAndBit()
		case IsNationalSet(ch):
			return token.ScanNationalSet()
		default:
			return token.ScanIdent()
		}
	case IsDigit(ch):
		switch ch = ch; {
		case IsZero(ch):
			return token.ScanHexOrBit()
		default:
			return token.ScanNumber()
		}
	case IsDot(ch):
		return token.ScanDot()
	case IsOperator(ch):
		token.next()
		return int(ch), string(ch)
	case IsLogicalOP(ch):
		return token.ScanLogicalOp()
	case IsComment(ch):
		return token.scanComment()
	case IsQuestion(ch):
		return token.ScanQuestion()
	default:
		token.next()
		return int(ch), string(ch)
	}
	return -1, ""
}

/*****************************************************************************
 *
 *	string
 *	'a',"a"
 *
 *****************************************************************************/
func (token *Tokener) ScanString(ch rune) (int, string) {
	buffer := token.Buf
	buffer.WriteRune(token.Lastchar)
	token.next()
	state := 0
	if token.EOF == true {
		state = -1
	}
	for state != -1 && state != -2 {
		switch state {
		case 0:
			if token.Lastchar != ch {
				buffer.WriteRune(token.Lastchar)
				token.next()
				state = 1
			} else {
				state = 2
			}
		case 1:
			if token.EOF == true {
				state = -2
			} else {
				state = 0
			}
		case 2:
			if token.doubleRune(ch) {
				state = 1
			} else {
				state = -2
			}
		}
	}

	if state == -1 {
		return int(ch), buffer.String()
	}

	if state == -2 {
		switch ch {
		case '\'':
			return SINGLE_QUOTA_STRING, conv.RemoveQuote(buffer.String())
		case '"':
			return DOUBLE_QUOTA_STRING, conv.RemoveQuote(buffer.String())
		}
	}
	return -1, ""

}

func stepScanString(token *Tokener, ch rune) (int, string) {
	return token.ScanString(ch)

}

func (token *Tokener) ScanNationalSet() (int, string) {
	buffer := token.Buf
	buffer.WriteRune(token.Lastchar)
	token.next()

	state := 0
	if token.EOF == true {
		state = -1
	}
	for state != -1 && state != -2 && state != -3 {
		switch state {
		case 0:
			if token.Lastchar == '\'' {
				state = -2
			} else if IsIdentLetter(token.Lastchar) {
				state = -3
			} else {
				state = -1
			}
		}
	}

	//字符串EOF或者遇到"或者非IDENTletter
	if state == -1 {
		return IDENT, buffer.String()
	}

	//遇到',开始变里
	if state == -2 {
		return token.ScanString('\'')
	}

	//按照ident遍历
	if state == -3 {
		return token.ScanIdent()
	}

	return -1, ""

}

/*****************************************************************************
 *
 *	IDENT
 *	123_abcd
 *	asdvxde
 *	`asdcdef#￥#￥Asd`
 *	1234adcdef
 *
 *****************************************************************************/

func (token *Tokener) ScanUnderline() (int, string) {
	buffer := token.Buf
	buffer.WriteRune(token.Lastchar)
	token.next()
	state := 0
	if token.EOF == true {
		state = -1
	}
	for state != -1 && state != -2 {
		switch state {
		case 0:
			if IsIdentLetter(token.Lastchar) {
				buffer.WriteRune(token.Lastchar)
				token.next()
				state = 1
			} else {
				state = -2
			}
		case 1:
			if token.EOF == true {
				state = -2
			} else {
				state = 0
			}
		}

	}

	if state == -1 {
		return int('_'), buffer.String()
	}

	if state == -2 {
		//判断是否是关键字
		return checkIdent(buffer.String())
	}

	return -1, ""

}

func (token *Tokener) ScanBackQuote() (int, string) {
	buffer := token.Buf
	ch := token.Lastchar
	buffer.WriteRune(token.Lastchar)
	token.next()
	state := 0
	if token.EOF == true {
		state = -1
	}
	for state != -1 && state != -2 {
		switch state {
		case 0:
			if token.Lastchar != ch {
				buffer.WriteRune(token.Lastchar)
				token.next()
				state = 1
			} else {
				state = 2
			}
		case 1:
			if token.EOF == true {
				state = -2
			} else {
				state = 0
			}
		case 2:
			if token.doubleRune(ch) {
				state = 1
			} else {
				state = -2
			}
		}
	}
	if state == -1 {
		return int('`'), buffer.String()
	}
	if state == -2 {
		return IDENT, buffer.String()
	}

	return -1, ""

}

func (token *Tokener) ScanIdent() (int, string) {
	buffer := token.Buf
	buffer.WriteRune(token.Lastchar)
	token.next()
	state := 0
	if token.EOF == true {
		state = -1
	}
	for state != -1 && state != -2 {
		switch state {
		case 0:
			if IsIdentLetter(token.Lastchar) {
				buffer.WriteRune(token.Lastchar)
				token.next()
				state = 1
			} else {
				state = -2
			}
		case 1:
			if token.EOF == true {
				state = -2
			} else {
				state = 0
			}

		}
	}

	if state == -1 || state == -2 {
		//判断是否是关键字
		return checkIdent(buffer.String())
	}
	return -1, ""
}

/*****************************************************************************
 *
 *	NUMBER
 *
 * 十六进制 X'01AF' X'01af' x'01AF' x'01af' 0x01AF 0x01af
 * 二级制 b'01' B'01' 0b01
 * 	数字 1, .2, 3.4, -5, -6.78, +9.10 -1.2E3
 *****************************************************************************/

//十六进制和bi判断
func (token *Tokener) ScanHexAndBit() (int, string) {
	buffer := bytes.NewBuffer([]byte{})
	ch := token.Lastchar
	token.Buf.WriteRune(token.Lastchar)
	token.next()
	state := 0
	if token.EOF == true {
		state = -1
	}
	for state != -1 && state != -2 && state != -3 && state != -4 {
		switch state {
		case 0:
			if token.Lastchar == '\'' {
				buffer.WriteRune(token.Lastchar)
				token.next()
				state = 1
			} else if IsIdentLetter(token.Lastchar) {
				token.Buf.WriteRune(token.Lastchar)
				token.next()
				state = 8
			} else {
				state = -1
			}
		case 1:
			//遇到结尾，返回B：IDENT
			//' “’”
			if token.EOF == true {
				state = -1
				//将读取状态设置为false,可以进入下一个switch，识别出'\''
				token.EOF = false
				//识别十六进制
			} else if IsHex(ch) {
				state = 2
				//识别二进制
			} else if IsBit(ch) {
				state = 3
				//拆开，识别
			} else {
				//不是二进制或者十六进制，应该拆开来看
				state = -3
			}
		case 2:
			if token.Lastchar == '\'' {
				buffer.WriteRune(token.Lastchar)
				token.next()
				state = -4
			} else {
				state = 4
			}
		case 3:
			if token.Lastchar == '\'' {
				buffer.WriteRune(token.Lastchar)
				token.next()
				state = -4
			} else {
				state = 5
			}
		case 4:
			if IsHexLetter(token.Lastchar) {
				buffer.WriteRune(token.Lastchar)
				token.next()
				state = 6
			} else {
				state = -3
			}
		case 5:
			if IsBitLetter(token.Lastchar) {
				buffer.WriteRune(token.Lastchar)
				token.next()
				state = 7
			} else {
				state = -3
			}
		case 6:
			if token.EOF == true {
				state = -4
			} else {
				state = 2
			}
		case 7:
			if token.EOF == true {
				state = -4
			} else {
				state = 3
			}
		case 8:
			if token.EOF == true {
				state = -1
			} else if IsIdentLetter(token.Lastchar) {
				state = -2
			} else {
				state = -1
			}

		}
	}

	if state == -1 {
		return checkIdent(token.Buf.String())
	}

	//遇到的是IDENT 继续遍历
	if state == -2 {
		return token.ScanIdent()
	}

	//遇到的是需要回退多个rune的
	if state == -3 {
		//先清除
		token.Buf.Reset()
		//将已经写入的填入
		token.Buf.WriteString(buffer.String())
		token.step = stepScanString
		return IDENT, string(ch)
	}

	if state == -4 {
		token.Buf.WriteString(buffer.String())
		if IsHex(ch) {
			return HEX_NUMBER, token.Buf.String()
		} else {
			return BIT_NUMBER, token.Buf.String()
		}
	}

	return -1, ""
}

//判断Hex和bit的另一种写法
func (token *Tokener) ScanHexOrBit() (int, string) {
	buffer := token.Buf
	buffer.WriteRune(token.Lastchar)
	token.next()
	state := 0
	if token.EOF == true {
		state = -1
	}

	for state != -1 && state != -2 && state != -3 && state != -4 && state != -5 && state != -6 && state != -7 {
		switch state {
		case 0:
			if token.Lastchar == 'x' {
				buffer.WriteRune(token.Lastchar)
				token.next()
				state = 1
			} else if token.Lastchar == 'b' {
				buffer.WriteRune(token.Lastchar)
				token.next()
				state = 4
			} else {
				state = 7
			}
		case 1:
			if token.EOF == true {
				state = -2
			} else if IsBlank(token.Lastchar) {
				state = -2
			} else {
				state = 2
			}
		case 2:
			if IsHexLetter(token.Lastchar) {
				buffer.WriteRune(token.Lastchar)
				token.next()
				state = 3
			} else if IsBlank(token.Lastchar) {
				state = -3
			} else {
				state = -4
			}
		case 3:
			if token.EOF == true {
				state = -3
			} else {
				state = 2
			}

		case 4:
			if token.EOF == true {
				state = -2
			} else if IsBlank(token.Lastchar) {
				state = -2
			} else {
				state = 5
			}
		case 5:
			if IsBitLetter(token.Lastchar) {
				buffer.WriteRune(token.Lastchar)
				token.next()
				state = 6
			} else if IsBlank(token.Lastchar) {
				state = -5
			} else {
				state = -4
			}
		case 6:
			if token.EOF == true {
				state = -5
			} else {
				state = 5
			}
		case 7:
			if IsDigit(token.Lastchar) {
				buffer.WriteRune(token.Lastchar)
				token.next()
				state = 8
			} else if IsDot(token.Lastchar) {
				state = 10
			} else if IsIdentLetter(token.Lastchar) {
				state = -4
			} else {
				state = -1
			}
		case 8:
			if token.EOF == true {
				state = -6
			} else {
				state = 9
			}
		case 9:
			if IsDigit(token.Lastchar) {
				buffer.WriteRune(token.Lastchar)
				token.next()
				state = 8
			} else if IsDot(token.Lastchar) {
				state = 10
			} else if IsIdentLetter(token.Lastchar) {
				state = -4
			} else {
				state = -6
			}
		case 10:
			token.next()
			if IsDigit(token.Lastchar) {
				buffer.WriteRune('.')
				buffer.WriteRune(token.Lastchar)
				token.next()
				state = 11
			} else {
				token.Scanner.UnreadRune()
				token.Lastchar = '.'
				token.Position--
				state = -1
			}
		case 11:
			if token.EOF == true {
				state = -7
			} else {
				state = 12
			}
		case 12:
			if IsDigit(token.Lastchar) {
				buffer.WriteRune(token.Lastchar)
				token.next()
				state = 11
			} else {
				state = -7
			}
		}
	}

	switch state {
	case -1:
		return INTEGER, "0"
	case -2:
		return IDENT, buffer.String()
	case -3:
		return HEX_NUMBER, buffer.String()
	case -4:
		return token.ScanIdent()
	case -5:
		return BIT_NUMBER, buffer.String()
	case -6:

		return INTEGER, buffer.String()
	case -7:
		return FLOAT, buffer.String()
	}

	return -1, ""

}

//判断Hex和bit的另一种写法
func (token *Tokener) ScanNumber() (int, string) {
	buffer := token.Buf
	buffer.WriteRune(token.Lastchar)
	token.next()

	state := 0
	if token.EOF == true {
		state = -1
	}
	for state != -1 && state != -2 && state != -3 {
		switch state {
		case 0:
			if IsDigit(token.Lastchar) {
				buffer.WriteRune(token.Lastchar)
				token.next()
				state = 1
			} else {
				state = 2
			}
		case 1:
			if token.EOF == true {
				state = -1
			} else {
				state = 0
			}
		case 2:
			if IsDot(token.Lastchar) {
				state = 3
			} else {
				state = 7
			}
		case 3:
			token.next()
			if token.EOF == true {
				token.EOF = false
				state = -1
			} else {
				state = 4
			}
		case 4:
			if IsDigit(token.Lastchar) {
				buffer.WriteRune('.')
				buffer.WriteRune(token.Lastchar)
				token.next()
				state = 5
			} else {
				state = -1
			}
		case 5:
			if token.EOF == true {
				state = -3
			} else {
				state = 6
			}
		case 6:
			if IsDigit(token.Lastchar) {
				buffer.WriteRune(token.Lastchar)
				token.next()
				state = 5
			} else {
				state = -3
			}
		case 7:
			if IsIdentLetter(token.Lastchar) {
				state = -2
			} else {
				state = -1
			}
		}
	}

	switch state {
	case -1:
		return INTEGER, buffer.String()
	case -2:
		return token.ScanIdent()
	case -3:
		return FLOAT, buffer.String()
	}

	return -1, ""

}

func (token *Tokener) ScanDot() (int, string) {
	buffer := token.Buf
	buffer.WriteRune(token.Lastchar)
	token.next()

	//判断是否是数字
	state := 0
	if token.EOF == true {
		state = -1
	}
	for state != -1 && state != -2 && state != -3 {
		switch state {
		case 0:
			if IsDot(token.Lastchar) {
				buffer.WriteRune(token.Lastchar)
				token.next()
				state = 1
			} else if IsDigit(token.Lastchar) {
				buffer.WriteRune(token.Lastchar)
				token.next()
				state = 3
			} else {
				state = -1
			}
		case 1:
			if token.EOF == true {
				state = -2
			} else {
				state = 2
			}
		case 2:
			if IsDot(token.Lastchar) {
				buffer.WriteRune(token.Lastchar)
				token.next()
				state = 1
			} else {
				state = -2
			}
		case 3:
			if token.EOF == true {
				state = -3
			} else {
				state = 4
			}
		case 4:
			if IsDigit(token.Lastchar) {
				buffer.WriteRune(token.Lastchar)
				token.next()
				state = 3
			} else {
				state = -3
			}
		}
	}

	if state == -1 {
		return int('.'), "."
	}

	if state == -2 {
		return STRING, buffer.String()
	}

	if state == -3 {
		return FLOAT, buffer.String()
	}

	return -1, ""
}

/*****************************************************************************
 *
 *	逻辑符号
 *
 *****************************************************************************/

//判断逻辑符号
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
			return LG, "<>"
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

/*****************************************************************************
 *
 *	Comment
 *
 *****************************************************************************/

func (token *Tokener) scanComment() (int, string) {
	buffer := token.Buf
	buffer.WriteRune(token.Lastchar)
	token.next()

	state := 0
	for state != -1 && state != -2 {
		switch state {
		//判断下一个字符是否是*或者已经遍历结束
		case 0:
			if token.Lastchar == '*' {
				buffer.WriteRune(token.Lastchar)
				token.next()
				state = 1
			} else {
				state = -1
			}
		case 1:
			if token.Lastchar != '*' && token.EOF == false {
				buffer.WriteRune(token.Lastchar)
				token.next()
			} else {
				state = 2
			}
		case 2:
			if token.EOF == true {
				buffer.WriteRune(token.Lastchar)
				state = -2
			} else {
				buffer.WriteRune(token.Lastchar)
				token.next()
				state = 3
			}
		case 3:
			if token.Lastchar == '/' || token.EOF == true {
				buffer.WriteRune(token.Lastchar)
				token.next()
				state = -2
			} else {
				state = 1
			}

		}
	}

	if state == -1 {
		return int('/'), buffer.String()
	}

	if state == -2 {
		return COMMENT, buffer.String()
	}
	return -1, ""

}

func (token *Tokener) ScanQuestion() (int, string) {
	buffer := token.Buf
	buffer.WriteRune(token.Lastchar)
	token.next()

	state := 0
	for state != -1 {
		switch state {
		case 0:
			if token.EOF == true {
				state = -1
			} else {
				state = 1
			}
		case 1:
			if IsQuestion(token.Lastchar) {
				buffer.WriteRune(token.Lastchar)
				token.next()
				state = 0
			} else {
				state = -1
			}
		}
	}

	if state == -1 {
		return PARAM_MARKER, buffer.String()
	}

	return -1, ""
}
