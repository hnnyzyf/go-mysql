package myparser

import "testing"
import "errors"
import AST "myast"
import "reflect"

func Test_scan(t *testing.T) {
	//sql:="SELECT count(1) from sub_settlement_record_43 where id in (select id from sub_settlement_record_43 where settlement_date >= '2018-01-16' and settlement_date <= '2018-01-16') and oper_type = 'P' and id not in (select a.id from sub_settlement_record_43 a join atom_sale_void_settlement_record b on a.settlement_date >= '2018-01-16' and a.settlement_date <= '2018-01-16' and a.unique_id = b.related_unique_id ) and matching_state in (0,1)"
	//sql:="select ? ? from test where a in (.....)"
	sql := "SELECT date_format ( send_Date , ? ) DataDate , COUNT (?) ReplyCount FROM chat_question_reply_info WHERE send_Date >= ? AND send_Date <= ? AND product_line = ? AND message_type = ? GROUP BY date_format "
	mtoken := NewTokener(sql)
	i := 0

	for mtoken.EOF == false {
		id, str := mtoken.Scan()
		t.Log(i, id, str)
		i = i + 1
	}

}

func Test_parse(t *testing.T) {
	//sql := "select EXTRACT(HOUR_SECOND FROM '2003-01-02 10:30:00.000123'),+-+-   1.234,CAST(6/4 AS DECIMAL(3,1)),DATE_SUB(CURDATE(),INTERVAL 30 DAY),DATE_ADD('2006-05-00',INTERVAL 1 DAY)+1,* from (SELECT * FROM Letter WHERE Uid='2116310619' and messageType != '' and businessType != '' and MessageCagetory=4 and isdelete=0 union all ( SELECT * FROM Letter l1 WHERE uid='' and messageType != '' and businessType != '' and MessageCagetory=4 and not exists (select letterId from Letter l2 where l2.uniqueId=l1.uniqueId and uid='2116310619'))) as a order by letterid desc limit 29550,50"
	//sql:="SELECT CAST(6/4 AS DECIMAL(3,1) CHARACTER SET abcd) COLLATE dasd,1"
	//sql:="select 1,2,3,4,5,6,7,8,`asdad`,\"dasdasd\",c as a,d,a.c from test,c.c"
	sql := "select count(1) from test"
	//sql := "select b,c,d,e from x,y"
	mtoken := NewTokener(sql)
	if yyParse(mtoken) != 0 {
		t.Error(errors.New(mtoken.LastError))
	} else {
		t.Log(mtoken.ASTree)
	}
	query := mtoken.ASTree.(*AST.SimpleSelectStmt)

	t.Log(query.NType)
	t.Log(query.Op)
	for _, expr := range query.Target {
		t.Log(reflect.TypeOf(expr.(*AST.TargetExpr).Tuple))
		switch sh := expr.(*AST.TargetExpr).Tuple.(type) {
		case *AST.ColumnExpr:
			{
				t.Log("ColumnExpr", sh.Table, sh.Column)
			}
		}
	}

	from := query.From.(*AST.Fromclause).Expr
	for _, expr := range from {
		t.Log(reflect.TypeOf(expr))
		t.Log(expr.(*AST.SimpleTableExpr).Table)
		t.Log(expr.(*AST.SimpleTableExpr).Alias)
	}
}

func Test_append(t *testing.T) {
	a := []int{1, 2, 3, 4}
	b := []int{5, 7, 8}

	for _, v := range a {
		b = append(b, v)
	}
	t.Log(b)
}
