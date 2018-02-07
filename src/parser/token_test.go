package parser

import "testing"
import "reflect"

const SINGLE_QUOTA_STRING = 13245
const DOUBLE_QUOTA_STRING = 13246
const HEX_NUMBER = 132432
const BIT_NUMBER = 13243234
const DOUBLE_AT_IDENT = 2432
const SINGLE_AT_IDENT = 244123435
const PARAM_MARKER = 66187
const IDENT = 57346
const STRING = 57347
const NUMBER = 57348
const AS = 57349
const ASC = 57350
const ALL = 57351
const ANY = 57352
const BY = 57353
const COMMENT = 57354
const CHARACTER = 57355
const COLLATE = 57356
const DISTINCT = 57357
const DISTINCTROW = 57358
const DESC = 57359
const DAY = 57360
const DAY_HOUR = 57361
const DAY_MICROSECOND = 57362
const DAY_MINUTE = 57363
const DAY_SECOND = 57364
const FROM = 57365
const FOR = 57366
const FALSE = 57367
const GROUP = 57368
const HAVING = 57369
const HOUR = 57370
const HOUR_MICROSECOND = 57371
const HOUR_MINUTE = 57372
const HOUR_SECOND = 57373
const INTO = 57374
const IF = 57375
const INTERVAL = 57376
const LIMIT = 57377
const LOCK = 57378
const MINUTE = 57379
const MINUTE_SECOND = 57380
const MODE = 57381
const MONTH = 57382
const MICROSECOND = 57383
const MINUTE_MICROSECOND = 57384
const NULL = 57385
const ORDER = 57386
const OFFSET = 57387
const QUARTER = 57388
const SELECT = 57389
const SOME = 57390
const SHARE = 57391
const SECOND = 57392
const SECOND_MICROSECOND = 57393
const SET = 57394
const TO = 57395
const TRUE = 57396
const UPDATE = 57397
const WHERE = 57398
const WEEK = 57399
const YEAR = 57400
const YEAR_MONTH = 57401
const PRIORITY = 57402
const UNION = 57403
const EXCEPT = 57404
const INTERSECT = 57405
const JOIN = 57406
const STRAIGHT_JOIN = 57407
const LEFT = 57408
const RIGHT = 57409
const INNER = 57410
const OUTER = 57411
const CROSS = 57412
const NATURAL = 57413
const USE = 57414
const ON = 57415
const USING = 57416
const OR = 57417
const AND = 57418
const NOT = 57419
const BETWEEN = 57420
const CASE = 57421
const WHEN = 57422
const THEN = 57423
const ELSE = 57424
const IS = 57425
const LIKE = 57426
const IN = 57427
const GE = 57428
const LE = 57429
const NE = 57430
const IE = 57431
const LNE = 57432
const SL = 57433
const SR = 57434
const LEG = 57435
const AA = 57436
const OO = 57437
const XOR = 57438
const EXISTS = 57439
const OP = 57440
const UMINUS = 57441
const END = 57442

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

func Test_string(t *testing.T) {
	sql := []string{
		"    \"a1\" ",
		"  \"b2\"\"\" ",
		"\"c3  '''\"  ",
		"\"王'\"''\" ",
	}
	res := []string{
		"\"a1\"",
		"\"b2\"\"\"",
		"\"c3  '''\"",
		"\"王'\"",
	}

	for i, v := range sql {
		mtoken := NewTokener(v)
		tp, s := mtoken.Scan()
		if tp == DOUBLE_QUOTA_STRING && s == res[i] {
			t.Log("编号:", i, " DOUBLE_STRING测试通过")
		} else {
			t.Error("编号:", i, "DOUBLE_STRING测试失败,返回:", s, "实际:", res[i])
		}
	}

	sql2 := []string{
		"N'ab\"\"'''abcd",
		"  N'ab  \"\"'    '' ab  cd   ",
	}
	res2 := []string{
		"N'ab\"\"'''",
		"N'ab  \"\"'",
	}

	for i, v := range sql2 {
		mtoken := NewTokener(v)
		tp, s := mtoken.Scan()
		if tp == SINGLE_QUOTA_STRING && s == res2[i] {
			t.Log("编号:", i, "SINGLE_STRING测试通过")
		} else {
			t.Error("编号:", i, "SINGLE_STRING测试失败,返回:", s, "实际:", res2[i])
		}
	}

	sql3 := []string{
		"  .... 12345",
		"  ....12345",
	}
	res3 := []string{
		"....",
		"....",
	}

	for i, v := range sql3 {
		mtoken := NewTokener(v)
		tp, s := mtoken.Scan()
		if tp == STRING && s == res3[i] {
			t.Log("编号:", i, "STRING测试通过")
		} else {
			t.Error("编号:", i, "STRING测试失败,返回:", s, "实际:", res3[i])
		}
	}

	sql4 := []string{
		"  ???1234",
		"  ??? 1234",
		"  ??? ",
	}
	res4 := []string{
		"???",
		"???",
		"???",
	}

	for i, v := range sql4 {
		mtoken := NewTokener(v)
		tp, s := mtoken.Scan()
		if tp == PARAM_MARKER && s == res4[i] {
			t.Log("编号:", i, "PARAM_MARKER测试通过")
		} else {
			t.Error("编号:", i, "PARAM_MARKER测试失败,返回:", s, "实际:", res4[i])
		}
	}

}

func Test_ident(t *testing.T) {
	sql := []string{
		"N\"ab\"\"''\"abcd ",
		"N_sd23abcd@asdasd ",
		"N",
		"__asda234@#$%^sdad ",
		"  `asdadasd`dadas ",
		"   c ",
		" X'12xyz' ",
		" b ",
		"0x ",
		"0x123acfOG ",
		"0X123acfOG ",
		"1234abcde ",
		"0x123O  cfOG ",
		"  `asda dasd ` dadas ",
		"BY abc",
		"XY abc",
	}
	res := []string{
		"N",
		"N_sd23abcd",
		"N",
		"__asda234",
		"`asdadasd`",
		"c",
		"X",
		"b",
		"0x",
		"0x123acfOG",
		"0X123acfOG",
		"1234abcde",
		"0x123O",
		"`asda dasd `",
		"BY",
		"XY",
	}

	for i, v := range sql {
		mtoken := NewTokener(v)
		tp, s := mtoken.Scan()
		if tp == IDENT && s == res[i] {
			t.Log("编号:", i, " IDENT测试通过")
		} else {
			t.Error("编号:", i, " IDENT测试失败,返回:", s, "实际:", res[i])
		}
	}

	sql2 := []string{
		" @@ab1234?? ",
		" @@ ",
		"  @@ ?? ",
	}
	res2 := []string{
		"@@ab1234",
		"@@",
		"@@",
	}

	for i, v := range sql2 {
		mtoken := NewTokener(v)
		tp, s := mtoken.Scan()
		if tp == DOUBLE_AT_IDENT && s == res2[i] {
			t.Log("编号:", i, " DOUBLE_AT_IDENT测试通过")
		} else {
			t.Error("编号:", i, " DOUBLE_AT_IDENT测试失败,返回:", s, "实际:", res2[i])
		}
	}
}

func Test_number(t *testing.T) {
	sql := []string{
		" X'123456' ",
		" X'12ae' ",
		" 0x12345 ",
		" 0x123 45 ",
	}
	res := []string{
		"X'123456'",
		"X'12ae'",
		"0x12345",
		"0x123",
	}

	for i, v := range sql {
		mtoken := NewTokener(v)
		tp, s := mtoken.Scan()
		if tp == HEX_NUMBER && s == res[i] {
			t.Log("编号:", i, " NUMBER测试通过")
		} else {
			t.Error("编号:", i, " NUMBER测试失败,返回:", s, "实际:", res[i])
		}
	}

	sql2 := " X'12xyz '"
	mtoken := NewTokener(sql2)
	tp, s := mtoken.Scan()
	tp, s = mtoken.Scan()
	if tp == SINGLE_QUOTA_STRING && s == "'12xyz '" {
		t.Log("编号:1000 NUMBER测试通过")
	} else {
		t.Error("编号:1000 NUMBER测试失败,返回:", s, "实际:'12xyz '")
	}

	sql3 := []string{
		" 0.123avc ",
		" 0.a ",
		" 0 ",
		" 123456678 ",
		" 1234.2323ab ",
		" 1234.a ",
		" 1234. ",
		"  .123456abv ",
		"  .123456 ",
		" 12345 6678 ",
		" 1234.23 23ab ",
	}
	res3 := []string{
		"0.123",
		"0",
		"0",
		"123456678",
		"1234.2323",
		"1234",
		"1234",
		".123456",
		".123456",
		"12345",
		"1234.23",
	}

	for i, v := range sql3 {
		mtoken := NewTokener(v)
		tp, s := mtoken.Scan()
		if tp == NUMBER && s == res3[i] {
			t.Log("编号:", i, " NUMBER测试通过")
		} else {
			t.Error("编号:", i, " NUMBER测试失败,返回:", s, "实际:", res3[i])
		}
	}

}

func Test_symbol(t *testing.T) {
	sql := []rune{
		'_',
		'"',
		'\'',
		'`',
	}
	res := []string{
		"_",
		"\"",
		"'",
		"`",
	}

	for i, v := range sql {
		mtoken := NewTokener(string(v))
		tp, s := mtoken.Scan()
		if tp == int(v) && s == res[i] {
			t.Log("编号:", i, " SYMBOL:", string(v), "测试通过")
		} else {
			t.Error("编号:", i, " SYMBOL:", string(v), "测试失败,返回:", s, "实际:", res[i])
		}
	}
	sql2 := ".asdasd"
	res2 := "."
	mtoken := NewTokener(sql2)
	tp, s := mtoken.Scan()
	if tp == int('.') && s == res2 {
		t.Log("编号: SYMBOL:", string('.'), "测试通过")
	} else {
		t.Error("编号: SYMBOL:", string('.'), "测试失败,返回:", s, "实际:", res2)
	}
}

func Test_op(t *testing.T) {
	sql := []string{
		" !=0 ",
	}
	res := []int{
		LNE,
	}

	for i, v := range sql {
		mtoken := NewTokener(v)
		tp, s := mtoken.Scan()
		if tp == res[i] {
			t.Log("编号:", i, " OP:", "测试通过")
		} else {
			t.Error("编号:", i, " OP:", "测试失败,返回:", s, "实际:", res[i])
		}
	}
}

func Test_Scansql(t *testing.T) {
	//sql := "SELECT -1.234,classno,classname,avg(score) as `avg_score 12321` From sc,(SELECT *.a FROM class WHERE class.gno='%grade one%') as sub WHERE sc.sno in (SELECT cno FROM student WHERE student.classno=sub.classno) and sc.sno in (SELECT course.cno FROM course WHERE  course.cname='computer') GROUP BY classno,classname HAVING avg(score)>60 ORDER BY avg_score"
	//sql := "SELECT count(1) from sub_settlement_record_43 where id in (select id from sub_settlement_record_43 where settlement_date >= '2018-01-16' and settlement_date <= '2018-01-16') and oper_type = 'P' and id not in (select a.id from sub_settlement_record_43 a join atom_sale_void_settlement_record b on a.settlement_date >= '2018-01-16' and a.settlement_date <= '2018-01-16' and a.unique_id = b.related_unique_id ) and matching_state in (0,1)"
	//sql := "select 123.123"
	//sql:="SELECT a.userIp as userIp FROM (SELECT userIp FROM loginfo_login WHERE source_from = 'CRM' AND subsource_from = 'LOGIN' AND (user_type = 'MemberLogin' OR user_type = 'MemberLoginMask') AND field5 = 'P' AND userip IS NOT NULL AND userip !=0 AND userip != 2130706433 AND ip_range_1 != 0x0A000000 AND user_agent_md5 IS NOT NULL AND request_time >= '2018-01-23 04:55:01' AND request_time < '2018-01-23 16:55:01' GROUP BY userIp HAVING COUNT(userIp) >= 30 AND sum(respond_type = 'T') >= COUNT(Respond_Type) * 0.7 AND COUNT(DISTINCT user_agent_md5) = 1 AND COUNT(DISTINCT uid) >= COUNT(uid) * 0.5 ) AS a  INNER JOIN (SELECT userIp FROM loginfo_login WHERE source_from = 'CRM' AND subsource_from = 'LOGIN' AND (user_type = 'MemberLogin' OR user_type = 'MemberLoginMask')) as f"
	mtoken := NewTokener(sql)
	i := 0

	for mtoken.EOF == false {
		id, str := mtoken.Scan()
		t.Log(i, id, str)
		i = i + 1
	}

}
