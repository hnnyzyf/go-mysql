package parser

import "testing"
import "ast"
import "log"
import "errors"

func SqlParse(sql string) ast.Node {
	mtoken := NewTokener(sql)
	if yyParse(mtoken) != 0 {
		log.Println(errors.New(mtoken.LastError))
	}
	return mtoken.Stmt()
}

func Test_print(t *testing.T) {
	sql := `SELECT a.userIp as userIp FROM (
		SELECT userIp FROM loginfo_login 
		WHERE source_from = 'CRM' AND subsource_from = 'LOGIN' AND (user_type = 'MemberLogin' OR user_type = 'MemberLoginMask') AND field5 = 'P' AND userip IS NOT NULL AND userip !=0 AND userip != 2130706433 AND ip_range_1 != 0x0A000000 AND user_agent_md5 IS NOT NULL AND request_time >= '2018-01-23 04:55:01' AND request_time < '2018-01-23 16:55:01' 
		GROUP BY userIp 
		HAVING COUNT(userIp) >= 30 AND sum(respond_type = 'T') >= COUNT(Respond_Type) * 0.7 AND COUNT(DISTINCT user_agent_md5) = 1 AND COUNT(DISTINCT uid) >= COUNT(uid) * 0.5 ) AS a  
		INNER JOIN (
			SELECT userIp FROM loginfo_login 
		WHERE source_from = 'CRM' AND subsource_from = 'LOGIN' AND (user_type = 'MemberLogin' OR user_type = 'MemberLoginMask')) as f`

	mtoken := NewTokener(sql)
	i := 0

	for mtoken.EOF == false {
		id, str := mtoken.Scan()
		t.Log(i, id, str)
		i = i + 1
	}
}

func Test_parse(t *testing.T) {
	//sql := `SELECT a.userIp as userIp FROM (
	//	SELECT userIp FROM loginfo_login 
	//	WHERE source_from = 'CRM' AND subsource_from = 'LOGIN' AND (user_type = 'MemberLogin' OR user_type = 'MemberLoginMask') AND field5 = 'P' AND userip IS NOT NULL AND userip !=0 AND userip != 2130706433 AND ip_range_1 != 0x0A000000 AND user_agent_md5 IS NOT NULL AND request_time >= '2018-01-23 04:55:01' AND request_time < '2018-01-23 16:55:01' 
	//	GROUP BY userIp 
	//	HAVING COUNT(userIp) >= 30 AND sum(respond_type = 'T') >= COUNT(Respond_Type) * 0.7 AND COUNT(DISTINCT user_agent_md5) = 1 AND COUNT(DISTINCT uid) >= COUNT(uid) * 0.5 ) AS a  
	//	INNER JOIN (
	//		SELECT userIp FROM loginfo_login 
	//	WHERE source_from = 'CRM' AND subsource_from = 'LOGIN' AND (user_type = 'MemberLogin' OR user_type = 'MemberLoginMask')) as f`
	//sql:="select 1 from test where true and false"
	sql:=`SELECT CASE WHEN EXISTS(SELECT * FROM TMP WHERE TMP.id = SRC.id) THEN 1 ELSE 2 END FROM SRC`
	stmt := SqlParse(sql)
	t.Log(stmt)

}
