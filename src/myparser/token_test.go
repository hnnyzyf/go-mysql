package myparser

import "testing"
import "reflect"


const IDENT = 57346
const UN_IDENT = 57347
const STRING = 57348
const NUMBER = 57349
const AS = 57350
const ASC = 57351
const BY = 57352
const COMMENT = 57353
const DISTINCT = 57354
const DESC = 57355
const EXISTS = 57356
const FROM = 57357
const FOR = 57358
const GROUP = 57359
const HAVING = 57360
const INTO = 57361
const IF = 57362
const LIMIT = 57363
const LOCK = 57364
const NULL = 57365
const ORDER = 57366
const OFFSET = 57367
const SELECT = 57368
const TO = 57369
const UPDATE = 57370
const WHERE = 57371
const UNION = 57372
const EXCEPT = 57373
const INTERSECT = 57374
const JOIN = 57375
const STRAIGHT_JOIN = 57376
const LEFT = 57377
const RIGHT = 57378
const INNER = 57379
const OUTER = 57380
const CROSS = 57381
const NATURAL = 57382
const USE = 57383
const ON = 57384
const OR = 57385
const AND = 57386
const NOT = 57387
const BETWEEN = 57388
const CASE = 57389
const WHEN = 57390
const THEN = 57391
const ELSE = 57392
const IS = 57393
const LIKE = 57394
const IN = 57395
const GE = 57396
const LE = 57397
const NE = 57398
const IE = 57399
const UMINUS = 57400
const END = 57401

func Test_NewTokener(t *testing.T) {
	token := NewTokener("")
	if token.Position == 0 {
		t.Log("NewTokener测试通过了", token.Position)
	} else {
		t.Error("NewTokener测试失败了")
	}
	if token.EOF == false {
		t.Log("NewTokener测试通过了", token.EOF)
	} else {
		t.Error("NewTokener测试失败了")
	}
}

func Test_next(t *testing.T) {
	sql := "s !'中国\\"
	mtoken := NewTokener(sql)
	mtoken.next()
	if mtoken.Lastchar == 's' {
		t.Log("第1个测试通过了,字符是:", string(mtoken.Lastchar))
	} else {
		t.Error("第1个测试失败了,字符是:", string(mtoken.Lastchar))
	}
	mtoken.next()
	if mtoken.Lastchar == ' ' {
		t.Log("第2个测试通过了,字符是:", string(mtoken.Lastchar))
	} else {
		t.Error("第2个测试失败了,字符是:", string(mtoken.Lastchar))
	}
	mtoken.next()
	if mtoken.Lastchar == '!' {
		t.Log("第3个测试通过了,字符是:", string(mtoken.Lastchar))
	} else {
		t.Error("第3个测试失败了,字符是:", string(mtoken.Lastchar))
	}
	mtoken.next()
	if mtoken.Lastchar == '\'' {
		t.Log("第4个测试通过了,字符是:", string(mtoken.Lastchar))
	} else {
		t.Error("第4个测试失败了,字符是:", string(mtoken.Lastchar))
	}
	mtoken.next()
	if mtoken.Lastchar == '中' {
		t.Log("第5个测试通过了,字符是:", string(mtoken.Lastchar))
	} else {
		t.Error("第5个测试失败了,字符是:", string(mtoken.Lastchar))
	}
	mtoken.next()
	if mtoken.Lastchar == '国' {
		t.Log("第6个测试通过了,字符是:", string(mtoken.Lastchar))
	} else {
		t.Error("第6个测试失败了,字符是:", string(mtoken.Lastchar))
	}
	mtoken.next()
	if mtoken.Lastchar == '\\' {
		t.Log("第7个测试通过了,字符是:", string(mtoken.Lastchar))
	} else {
		t.Error("第7个测试失败了,字符是:", string(mtoken.Lastchar))
	}
	mtoken.next()
	if mtoken.EOF==true {
		t.Log("第8个测试通过了,错误是:", true)
	} else {
		t.Error("第8个测试失败了")
	}
}

func Test_skip(t *testing.T) {
	sql := "      s"
	mtoken := NewTokener(sql)
	mtoken.next()
	mtoken.skip()
	ch := mtoken.Lastchar
	if ch == 's' {
		t.Log("第1个测试通过了,字符是:", string(mtoken.Lastchar))
	} else {
		t.Error("第1个测试失败了,位置是:", mtoken.Position)
	}
	mtoken.next()
	if ch == '\n' {
		t.Log("第2个测试通过了,字符是:", string(mtoken.Lastchar))
	} else {
		t.Error("第2个测试失败了,位置是:", string(mtoken.Lastchar))
		t.Error(reflect.TypeOf(mtoken.Lastchar))
	}
}

func Test_scanNumber(t *testing.T) {
	sql := "12356"
	mtoken := NewTokener(sql)
	mtoken.next()
	_, str := mtoken.scanNumber()
	if str == "12356" {
		t.Log("第1个测试通过了,字符是:", str)
	} else {
		t.Error("第1个测试失败了,字符是:", str)
	}

	sql = "12356.12345"
	mtoken = NewTokener(sql)
	mtoken.next()
	_, str = mtoken.scanNumber()
	if str == "12356.12345" {
		t.Log("第2个测试通过了,字符是:", str)
	} else {
		t.Error("第2个测试失败了,字符是:", str)
	}
	//
	sql = "12356."
	mtoken = NewTokener(sql)
	mtoken.next()
	res, str := mtoken.scanNumber()
	if str == "12356." && res == 0 {
		t.Log("第3个测试通过了,字符是:", str)
	} else {
		t.Error("第3个测试失败了,字符是:", str)
	}

	sql = "12356312312312312312312312312 ."
	mtoken = NewTokener(sql)
	mtoken.next()
	res, str = mtoken.scanNumber()
	if str == "12356312312312312312312312312" && res == 0 {
		t.Log("第4个测试通过了,字符是:", str)
	} else {
		t.Error("第4个测试失败了,字符是:", str)
	}

	sql = "12356a"
	mtoken = NewTokener(sql)
	mtoken.next()
	res, str = mtoken.scanNumber()
	if str == "12356" && res == 0 {
		t.Log("第5个测试通过了,字符是:", str)
	} else {
		t.Error("第5个测试失败了,字符是:", str)
	}

	sql = "12356.a"
	mtoken = NewTokener(sql)
	mtoken.next()
	res, str = mtoken.scanNumber()
	if str == "12356." && res == 0 {
		t.Log("第6个测试通过了,字符是:", str)
	} else {
		t.Error("第6个测试失败了,字符是:", str)
	}

	sql = "12356+"
	mtoken = NewTokener(sql)
	mtoken.next()
	res, str = mtoken.scanNumber()
	if str == "12356" && res == 0 {
		t.Log("第6个测试通过了,字符是:", str)
	} else {
		t.Error("第6个测试失败了,字符是:", str)
	}

	sql = "12356.2.3.4"
	mtoken = NewTokener(sql)
	mtoken.next()
	res, str = mtoken.scanNumber()
	if str == "12356.2" && res == 0 {
		t.Log("第7个测试通过了,字符是:", str)
	} else {
		t.Error("第7个测试失败了,字符是:", str)
	}

}

func Test_scanIdent(t *testing.T) {
	sql := "abc&de"
	mtoken := NewTokener(sql)
	mtoken.next()
	_, str := mtoken.scanIdent()
	if str == "abc&de" {
		t.Log("第1个测试通过了,字符是:", str)
	} else {
		t.Error("第1个测试失败了,字符是:", str)
	}

	sql = "ab+cde"
	mtoken = NewTokener(sql)
	mtoken.next()
	_, str = mtoken.scanIdent()
	if str == "ab" {
		t.Log("第2个测试通过了,字符是:", str)
	} else {
		t.Error("第2个测试失败了,字符是:", str)
	}

	sql = "ab  cde"
	mtoken = NewTokener(sql)
	mtoken.next()
	_, str = mtoken.scanIdent()
	if str == "ab" {
		t.Log("第3个测试通过了,字符是:", str)
	} else {
		t.Error("第3个测试失败了,字符是:", str)
	}

	sql = "abcde@@@"
	mtoken = NewTokener(sql)
	mtoken.next()
	_, str = mtoken.scanIdent()
	if str == "abcde@@@" {
		t.Log("第4个测试通过了,字符是:", str)
	} else {
		t.Error("第4个测试失败了,字符是:", str)
	}

	sql = "ab.adc"
	mtoken = NewTokener(sql)
	mtoken.next()
	_, str = mtoken.scanIdent()
	if str == "ab" {
		t.Log("第5个测试通过了,字符是:", str)
	} else {
		t.Error("第5个测试失败了,字符是:", str)
	}

	sql = "select/*1234*/adc"
	mtoken = NewTokener(sql)
	mtoken.next()
	_, str = mtoken.scanIdent()
	if str == "select" {
		t.Log("第6个测试通过了,字符是:", str)
	} else {
		t.Error("第6个测试失败了,字符是:", str)
	}

}

func Test_scanString(t *testing.T) {
	sql := "'abcdef'"
	mtoken := NewTokener(sql)
	mtoken.next()
	_, str := mtoken.scanString()
	if str == sql {
		t.Log("第1个测试通过了,字符是:", str)
	} else {
		t.Error("第1个测试失败了,str是:", []rune(str))
		t.Error("第1个测试失败了,sql是:", []rune(sql))
	}

	sql = "'\"\"\"\"abcdefghijklmn中国123@"
	mtoken = NewTokener(sql)
	mtoken.next()
	_, str = mtoken.scanString()
	if str == sql {
		t.Log("第2个测试通过了,字符是:", str)
	} else {
		t.Error("第2个测试失败了,str是:", []rune(str))
		t.Error("第2个测试失败了,sql是:", []rune(sql))
	}

	sql = "`abcdefghijklmn中国123@`  1234"
	mtoken = NewTokener(sql)
	mtoken.next()
	_, str = mtoken.scanString()
	if str == "`abcdefghijklmn中国123@`" {
		t.Log("第3个测试通过了,字符是:", str)
	} else {
		t.Error("第3个测试失败了,str是:", []rune(str))
		t.Error("第3个测试失败了,sql是:", []rune(sql))
	}

}

func Test_scanOpertator(t *testing.T) {
	sql := "=A"
	mtoken := NewTokener(sql)
	mtoken.next()
	_, str := mtoken.scanOpertator()
	if str == "=" {
		t.Log("第1个测试通过了,字符是:", str)
	} else {
		t.Error("第1个测试失败了,str是:", []rune(str))
		t.Error("第1个测试失败了,sql是:", []rune(sql))
	}

	sql = "=="
	mtoken = NewTokener(sql)
	mtoken.next()
	_, str = mtoken.scanOpertator()
	if str == "==" {
		t.Log("第2个测试通过了,字符是:", str)
	} else {
		t.Error("第2个测试失败了,str是:", []rune(str))
		t.Error("第2个测试失败了,sql是:", []rune(sql))
	}

	sql = "<=a"
	mtoken = NewTokener(sql)
	mtoken.next()
	_, str = mtoken.scanOpertator()
	if str == "<=" {
		t.Log("第3个测试通过了,字符是:", str)
	} else {
		t.Error("第3个测试失败了,str是:", []rune(str))
		t.Error("第3个测试失败了,sql是:", []rune(sql))
	}

	sql = "<>a"
	mtoken = NewTokener(sql)
	mtoken.next()
	_, str = mtoken.scanOpertator()
	if str == "<>" {
		t.Log("第4个测试通过了,字符是:", str)
	} else {
		t.Error("第4个测试失败了,str是:", []rune(str))
		t.Error("第4个测试失败了,sql是:", []rune(sql))
	}

	sql = "='"
	mtoken = NewTokener(sql)
	mtoken.next()
	_, str = mtoken.scanOpertator()
	if str == "=" {
		t.Log("第5个测试通过了,字符是:", str)
	} else {
		t.Error("第5个测试失败了,str是:", []rune(str))
		t.Error("第5个测试失败了,sql是:", []rune(sql))
	}
}

func Test_scanComment(t *testing.T) {
	sql := "/*1234567**/abcderf"
	mtoken := NewTokener(sql)
	mtoken.next()
	_, str := mtoken.scanComment()
	if str == "/*1234567**/" {
		t.Log("第1个测试通过了,字符是:", str)
	} else {
		t.Error("第1个测试失败了,str是:", str)
		t.Error("第1个测试失败了,str是:", []rune(str))
		t.Error("第1个测试失败了,sql是:", []rune(sql))
	}


}

func Test_Scan(t *testing.T) {
	sql := "/*12345*/SELECT classno,classname,avg(score) as `avg_score` From sc,(SELECT *.a FROM class WHERE class.gno='%grade one%') as sub WHERE sc.sno in (SELECT cno FROM student WHERE student.classno=sub.classno) and sc.sno in (SELECT course.cno FROM course WHERE  course.cname='computer') GROUP BY classno,classname HAVING avg(score)>60 ORDER BY avg_score; "
	//sql:="select 123.123"
	mtoken := NewTokener(sql)
	i:=0
	for mtoken.EOF==false {
		id, str := mtoken.Scan()
		t.Log(i,id, str)
		i=i+1
	}

}
