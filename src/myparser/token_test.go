package myparser

import "testing"
import "reflect"

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
	if mtoken.EOF == true {
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
	if ch == '\t' {
		t.Log("第2个测试通过了,字符是:", string(mtoken.Lastchar))
	} else {
		t.Error("第2个测试失败了,位置是:", string(mtoken.Lastchar))
		t.Error(reflect.TypeOf(mtoken.Lastchar))
	}
}

func Test_scan(t *testing.T) {
	sql := "'abcd我们试试节上最好的&……%&……3451235！@32545《》？；【】‘'asdasd"
	mtoken := NewTokener(sql)
	_, str := mtoken.Scan()
	if str == "'abcd我们试试节上最好的&……%&……3451235！@32545《》？；【】‘'" {
		t.Log("string测试通过了,字符是:", str)
	} else {
		t.Error("string测试失败了,位置是:", mtoken.Position, str)
	}

	sql = "_sda$%sd'12345'asdaf"
	mtoken = NewTokener(sql)
	_, str = mtoken.Scan()
	if str == "_sda" {
		t.Log("underline测试通过了,字符是:", str)
	} else {
		t.Error("underline测试失败了,位置是:", str)
	}

	sql = "_sdad\"12345\"asdaf"
	mtoken = NewTokener(sql)
	_, str = mtoken.Scan()
	if str == "_sdad\"12345\"" {
		t.Log("underline测试通过了,字符是:", str)
	} else {
		t.Error("underline测试失败了,位置是:", str)
	}

	sql = ". abcde"
	mtoken = NewTokener(sql)
	_, str = mtoken.Scan()
	if str == "." {
		t.Log("dot测试通过了,字符是:", str)
	} else {
		t.Error("dot测试失败了,位置是:", str)
	}

	sql = ".abcd"
	mtoken = NewTokener(sql)
	_, str = mtoken.Scan()
	if str == "." {
		t.Log("dot测试通过了,字符是:", str)
	} else {
		t.Error("dot测试失败了,位置是:", str)
	}

	sql = ".123$345a"
	mtoken = NewTokener(sql)
	_, str = mtoken.Scan()
	if str == ".123" {
		t.Log("dot测试通过了,字符是:", str)
	} else {
		t.Error("dot测试失败了,位置是:", str)
	}

	sql = "`dasd34235&*%^*^`"
	mtoken = NewTokener(sql)
	_, str = mtoken.Scan()
	if str == "`dasd34235&*%^*^`" {
		t.Log("`测试通过了,字符是:", str)
	} else {
		t.Error("`测试失败了,位置是:", str)
	}

	sql = "x'12345678'abcd"
	mtoken = NewTokener(sql)
	ty, str1 := mtoken.Scan()
	if str1 == "x'12345678'" && ty == NUMBER {
		t.Log("hex测试通过了,字符是:", str1, ty)
	} else {
		t.Error("hex测试失败了,位置是:", str1, ty)
	}

	sql = "x'21a4z46'abcd"
	mtoken = NewTokener(sql)
	ty, str1 = mtoken.Scan()
	if str1 == "x'21a4z46'" && ty == IDENT {
		t.Log("hex测试通过了,字符是:", str1, ty)
	} else {
		t.Error("hex测试失败了,位置是:", str1, ty)
	}

	sql = "B'0101011'abcd"
	mtoken = NewTokener(sql)
	ty, str1 = mtoken.Scan()
	if str1 == "B'0101011'" && ty == NUMBER {
		t.Log("BIT测试通过了,字符是:", str1, ty)
	} else {
		t.Error("BIT测试失败了,位置是:", str1, ty)
	}

	sql = "B'0101a11'abcd"
	mtoken = NewTokener(sql)
	ty, str1 = mtoken.Scan()
	if str1 == "B'0101a11'" && ty == IDENT {
		t.Log("BIT测试通过了,字符是:", str1, ty)
	} else {
		t.Error("BIT测试失败了,位置是:", str1, ty)
	}

	sql = "N'0101a11'abcd"
	mtoken = NewTokener(sql)
	ty, str1 = mtoken.Scan()
	if str1 == "N'0101a11'" && ty == STRING {
		t.Log("NATIONSET测试通过了,字符是:", str1, ty)
	} else {
		t.Error("NATIONSET测试失败了,位置是:", str1, ty)
	}

	sql = "ddsad45#$#asd dasdasdad"
	mtoken = NewTokener(sql)
	ty, str1 = mtoken.Scan()
	if str1 == "ddsad45#$#asd" && ty == IDENT {
		t.Log("IDENT测试通过了,字符是:", str1, ty)
	} else {
		t.Error("IDENT测试失败了,位置是:", str1, ty)
	}

	sql = "ddsad45#$#asd dasdasdad"
	mtoken = NewTokener(sql)
	ty, str1 = mtoken.Scan()
	if str1 == "ddsad45#$#asd" && ty == IDENT {
		t.Log("IDENT测试通过了,字符是:", str1, ty)
	} else {
		t.Error("IDENT测试失败了,位置是:", str1, ty)
	}

	sql = "0x13213123    asdas d"
	mtoken = NewTokener(sql)
	ty, str1 = mtoken.Scan()
	if str1 == "0x13213123" && ty == NUMBER {
		t.Log("NUMBER测试通过了,字符是:", str1, ty)
	} else {
		t.Error("NUMBER测试失败了,位置是:", []byte(str1))
		t.Error("NUMBER测试失败了,位置是:", []byte("0x1"))
		t.Error("NUMBER测试失败了,位置是:", str1)
		t.Error("NUMBER测试失败了,位置是:", ty)
	}

	sql = "0x1321zz3123    asdas d"
	mtoken = NewTokener(sql)
	ty, str1 = mtoken.Scan()
	if str1 == "0x1321" && ty == NUMBER {
		t.Log("NUMBER测试通过了,字符是:", str1, ty)
	} else {
		t.Error("NUMBER测试失败了,位置是:", []byte(str1))
		t.Error("NUMBER测试失败了,位置是:", []byte("0x1"))
		t.Error("NUMBER测试失败了,位置是:", str1)
		t.Error("NUMBER测试失败了,位置是:", ty)
	}

	sql = "01234zz    asdas d"
	mtoken = NewTokener(sql)
	ty, str1 = mtoken.Scan()
	if str1 == "01234" && ty == NUMBER {
		t.Log("NUMBER测试通过了,字符是:", str1, ty)
	} else {
		t.Error("NUMBER测试失败了,位置是:", []byte(str1))
		t.Error("NUMBER测试失败了,位置是:", []byte("0x1"))
		t.Error("NUMBER测试失败了,位置是:", str1)
		t.Error("NUMBER测试失败了,位置是:", ty)
	}

	sql = "  -1234  "
	mtoken = NewTokener(sql)
	ty, str1 = mtoken.Scan()
	if str1 == "-1234" && ty == NUMBER {
		t.Log("NUMBER测试通过了,字符是:", str1, ty)
	} else {
		t.Error("NUMBER测试失败了,位置是:", []byte(str1))
		t.Error("NUMBER测试失败了,位置是:", []byte("0x1"))
		t.Error("NUMBER测试失败了,位置是:", str1)
		t.Error("NUMBER测试失败了,位置是:", ty)
	}

	sql = "  -  1234  "
	mtoken = NewTokener(sql)
	ty, str1 = mtoken.Scan()
	if str1 == "-" {
		t.Log("op测试通过了,字符是:", str1, ty)
	} else {
		t.Error("op测试失败了,位置是:", []byte(str1))
		t.Error("op测试失败了,位置是:", []byte("0x1"))
		t.Error("op测试失败了,位置是:", str1)
		t.Error("op测试失败了,位置是:", ty)
	}

	sql = "  <=> "
	mtoken = NewTokener(sql)
	ty, str1 = mtoken.Scan()
	if str1 == "<=>" {
		t.Log("op测试通过了,字符是:", str1, ty)
	} else {
		t.Error("op测试失败了,位置是:", []byte(str1))
		t.Error("op测试失败了,位置是:", []byte("0x1"))
		t.Error("op测试失败了,位置是:", str1)
		t.Error("op测试失败了,位置是:", ty)
	}

	sql = "  <  => "
	mtoken = NewTokener(sql)
	ty, str1 = mtoken.Scan()
	if str1 == "<" {
		t.Log("op测试通过了,字符是:", str1, ty)
	} else {
		t.Error("op测试失败了,位置是:", []byte(str1))
		t.Error("op测试失败了,位置是:", []byte("0x1"))
		t.Error("op测试失败了,位置是:", str1)
		t.Error("op测试失败了,位置是:", ty)
	}

	sql = "  /*213123*/ "
	mtoken = NewTokener(sql)
	ty, str1 = mtoken.Scan()
	if str1 == "/*213123*/" {
		t.Log("comment测试通过了,字符是:", str1, ty)
	} else {
		t.Error("comment测试失败了,位置是:", []byte(str1))
		t.Error("comment测试失败了,位置是:", []byte("0x1"))
		t.Error("comment测试失败了,位置是:", str1)
		t.Error("comment测试失败了,位置是:", ty)
	}

	sql = "  / *213123*/ "
	mtoken = NewTokener(sql)
	ty, str1 = mtoken.Scan()
	if str1 == "/" {
		t.Log("comment测试通过了,字符是:", str1, ty)
	} else {
		t.Error("comment测试失败了,位置是:", []byte(str1))
		t.Error("comment测试失败了,位置是:", []byte("0x1"))
		t.Error("comment测试失败了,位置是:", str1)
		t.Error("comment测试失败了,位置是:", ty)
	}

	sql = " fsdgdfg "
	mtoken = NewTokener(sql)
	ty, str1 = mtoken.Scan()
	if str1 == "fsdgdfg" {
		t.Log("IDENT测试通过了,字符是:", str1, ty)
	} else {
		t.Error("IDENT测试失败了,位置是:", []byte(str1))
		t.Error("IDENT测试失败了,位置是:", []byte("0x1"))
		t.Error("IDENT测试失败了,位置是:", str1)
		t.Error("IDENT测试失败了,位置是:", ty)
	}

	sql = " ,fsdgdfg "
	mtoken = NewTokener(sql)
	ty, str1 = mtoken.Scan()
	if str1 == "," {
		t.Log("IDENT测试通过了,字符是:", str1, ty)
	} else {
		t.Error("IDENT测试失败了,位置是:", []byte(str1))
		t.Error("IDENT测试失败了,位置是:", []byte("0x1"))
		t.Error("IDENT测试失败了,位置是:", str1)
		t.Error("IDENT测试失败了,位置是:", ty)
	}

}

func Test_Scansql(t *testing.T) {
	//sql := "SELECT -1.234,classno,classname,avg(score) as `avg_score 12321` From sc,(SELECT *.a FROM class WHERE class.gno='%grade one%') as sub WHERE sc.sno in (SELECT cno FROM student WHERE student.classno=sub.classno) and sc.sno in (SELECT course.cno FROM course WHERE  course.cname='computer') GROUP BY classno,classname HAVING avg(score)>60 ORDER BY avg_score"
	sql := "SELECT count(1) from sub_settlement_record_43 where id in (select id from sub_settlement_record_43 where settlement_date >= '2018-01-16' and settlement_date <= '2018-01-16') and oper_type = 'P' and id not in (select a.id from sub_settlement_record_43 a join atom_sale_void_settlement_record b on a.settlement_date >= '2018-01-16' and a.settlement_date <= '2018-01-16' and a.unique_id = b.related_unique_id ) and matching_state in (0,1)"
	//sql := "select 123.123"
	mtoken := NewTokener(sql)
	i := 0

	for mtoken.EOF == false {
		id, str := mtoken.Scan()
		t.Log(i, id, str)
		i = i + 1
	}

}
