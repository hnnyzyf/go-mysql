package parser

import "testing"
import "reflect"

const BuiltinCharacterIdent = 57346
const BuiltinFucTimeAddIdent = 57347
const BuiltinFucTimeSubIdent = 57348
const BuiltinTimeUnitIdent = 57349
const IDENT = 57350
const DOUBLE_QUOTA_STRING = 57351
const PARAM_MARKER = 57352
const SINGLE_QUOTA_STRING = 57353
const STRING = 57354
const BIT_NUMBER = 57355
const HEX_NUMBER = 57356
const NUMBER = 57357
const ALL = 57358
const ANY = 57359
const AS = 57360
const ASC = 57361
const AVG = 57362
const BY = 57363
const CAST = 57364
const CHARACTER = 57365
const COMMENT = 57366
const CONVERT = 57367
const COUNT = 57368
const DESC = 57369
const DISTINCT = 57370
const DISTINCTROW = 57371
const FALSE = 57372
const FOR = 57373
const FROM = 57374
const GROUP = 57375
const HAVING = 57376
const IF = 57377
const INTO = 57378
const LIMIT = 57379
const LOCK = 57380
const MAX = 57381
const MIN = 57382
const MODE = 57383
const NULL = 57384
const OFFSET = 57385
const ORDER = 57386
const POSITION = 57387
const QUARTER = 57388
const SELECT = 57389
const SET = 57390
const SHARE = 57391
const SOME = 57392
const SOUNDS = 57393
const SUM = 57394
const TO = 57395
const TRUE = 57396
const UPDATE = 57397
const WHERE = 57398
const EXISTS = 57399
const OUTER = 57400
const LOW = 57401
const UNION = 57402
const EXCEPT = 57403
const INTERSECT = 57404
const JOIN = 57405
const STRAIGHT_JOIN = 57406
const NATURAL = 57407
const LEFT = 57408
const RIGHT = 57409
const INNER = 57410
const CROSS = 57411
const USE = 57412
const TABLE_REF_PRIORITY = 57413
const ON = 57414
const USING = 57415
const OR = 57416
const OO = 57417
const XOR = 57418
const AND = 57419
const AA = 57420
const BETWEEN = 57421
const CASE = 57422
const WHEN = 57423
const THEN = 57424
const ELSE = 57425
const LIKE = 57426
const REGEXP = 57427
const IN = 57428
const GE = 57429
const LE = 57430
const NE = 57431
const LG = 57432
const IE = 57433
const LNE = 57434
const LEG = 57435
const IS = 57436
const SL = 57437
const SR = 57438
const DIV = 57439
const MOD = 57440
const NOT = 57441
const OP = 57442
const COLLATE = 57443
const INTERVAL = 57444
const UMINUS = 57445
const END = 57446
const INTEGER=57443
const FLOAT=57444

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
		"XY abc",
		" 000000000x 000000",
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
		"XY",
		"000000000x",
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

}

func Test_number(t *testing.T) {
	sql := []string{
		" X'123456' ",
		" X'12ae' ",
		" 0x12345 ",
		" 0x123 45 ",
		"0x000000010000345 ",

	}
	res := []string{
		"X'123456'",
		"X'12ae'",
		"0x12345",
		"0x123",
		"0x000000010000345",
	}

	for i, v := range sql {
		mtoken := NewTokener(v)
		tp, s := mtoken.Scan()
		if tp == HEX_NUMBER && s == res[i] {
			t.Log("编号:", i, " HEX_NUMBER测试通过")
		} else {
			t.Error("编号:", i, " HEX_NUMBER测试失败,返回:", s, "实际:", res[i])
		}
	}

	sqlb := []string{
		" B'00001001' ",
		" 0b000101010 ",
		
	}
	resb := []string{
		"B'00001001'",
		"0b000101010",
	}

	for i, v := range sqlb {
		mtoken := NewTokener(v)
		tp, s := mtoken.Scan()
		if tp == BIT_NUMBER && s == resb[i] {
			t.Log("编号:", i, " BIT_NUMBER测试通过")
		} else {
			t.Error("编号:", i, " BIT_NUMBER测试失败,返回:", s, "实际:", resb[i])
		}
	}

	sql2 := " X'12xyz '"
	mtoken := NewTokener(sql2)
	tp, s := mtoken.Scan()
	tp, s = mtoken.Scan()
	if tp == SINGLE_QUOTA_STRING && s == "'12xyz '" {
		t.Log("编号:1000 HEX_NUMBER测试通过")
	} else {
		t.Error("编号:1000 HEX_NUMBER测试失败,返回:", s, "实际:'12xyz '")
	}

	sql3 := []string{
		" 0.123avc ",
		" 1234.2323ab ",
		"  .123456abv ",
		"  .123456 ",
		" 1234.23 23ab ",
		"00000.0000",
		"01234.23234",
	}
	res3 := []string{
		"0.123",
		"1234.2323",
		".123456",
		".123456",
		"1234.23",
		"00000.0000",
		"01234.23234",
	}

	for i, v := range sql3 {
		mtoken := NewTokener(v)
		tp, s := mtoken.Scan()
		if tp == FLOAT && s == res3[i] {
			t.Log("编号:", i, " NUMBER测试通过")
		} else {
			t.Error("编号:", i, " NUMBER测试失败,返回:", s, "实际:", res3[i])
		}
	}

	sql4 := []string{
		" 0.a ",
		" 0 ",
		" 123456678 ",
		" 1234.a ",
		" 1234. ",
		" 12345 6678 ",
		"0000000000",
	}
	res4 := []string{
		"0",
		"0",
		"123456678",
		"1234",
		"1234",
		"12345",
		"0000000000",
	}

	for i, v := range sql4 {
		mtoken := NewTokener(v)
		tp, s := mtoken.Scan()
		if tp == INTEGER && s == res4[i] {
			t.Log("编号:", i, " NUMBER测试通过")
		} else {
			t.Error("编号:", i, " NUMBER测试失败,返回:", s, "实际:", res4[i])
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
	sql := "select 123.123"
	//sql:="SELECT a.userIp as userIp FROM (SELECT userIp FROM loginfo_login WHERE source_from = 'CRM' AND subsource_from = 'LOGIN' AND (user_type = 'MemberLogin' OR user_type = 'MemberLoginMask') AND field5 = 'P' AND userip IS NOT NULL AND userip !=0 AND userip != 2130706433 AND ip_range_1 != 0x0A000000 AND user_agent_md5 IS NOT NULL AND request_time >= '2018-01-23 04:55:01' AND request_time < '2018-01-23 16:55:01' GROUP BY userIp HAVING COUNT(userIp) >= 30 AND sum(respond_type = 'T') >= COUNT(Respond_Type) * 0.7 AND COUNT(DISTINCT user_agent_md5) = 1 AND COUNT(DISTINCT uid) >= COUNT(uid) * 0.5 ) AS a  INNER JOIN (SELECT userIp FROM loginfo_login WHERE source_from = 'CRM' AND subsource_from = 'LOGIN' AND (user_type = 'MemberLogin' OR user_type = 'MemberLoginMask')) as f"
	mtoken := NewTokener(sql)
	i := 0

	for mtoken.EOF == false {
		id, str := mtoken.Scan()
		t.Log(i, id, str)
		i = i + 1
	}

}
