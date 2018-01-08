package myparser

import "testing"
import "errors"

func Test_parse(t *testing.T){
	//sql := "select * from (SELECT * FROM Letter WHERE Uid='2116310619' and messageType != '' and businessType != '' and MessageCagetory=4 and isdelete=0 union all ( SELECT * FROM Letter l1 WHERE uid='' and messageType != '' and businessType != '' and MessageCagetory=4 and not exists (select letterId from Letter l2 where l2.uniqueId=l1.uniqueId and uid='2116310619'))) as a order by letterid desc limit 29550,50"
	sql:="SELECT film.title, film.film_id FROM film JOIN film_actor using(a) WHERE a is true"
	mtoken:=NewTokener(sql)
	if yyParse(mtoken) != 0 {
		t.Error(errors.New(mtoken.LastError))
	}else{
		if mtoken.AST=="SELECT`中国`FROM1"{
			t.Log("测试通过,结果是",mtoken.AST)
		}else{
			t.Error("测试失败,结果是",mtoken.AST)
		}
	}
}
