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
	//sql:=`SELECT CASE WHEN EXISTS(SELECT * FROM TMP WHERE TMP.id = SRC.id) THEN 1 ELSE 2 END FROM SRC`
	//sql:="select a.*"
	//sql:=`select product_baseinfo.* , product_sellcount.showprice, product_sellcount.SellCount,
	// product_sellcount.SellPoint, product_sellcount.ShowPriceStrage, product_sellcount.ShowPriceStrageValue, 
	// product_presell.IsPresell, IFNULL(product_presell.PresellEndTime,null) as presellEndTime, 
	// product_exclusiveinfo.ExclusiveType, IFNULL(product_exclusiveinfo.Status,0) as exclusiveStatus,
	//  IFNULL(product_presell.Status,0) as presellStatus, '' as brandName, '' as cateName, 0 as SaleCount 
	//  from product_baseinfo inner join mktypproductdb.product_sellcount on product_baseinfo.productID = 
	//  product_sellcount.productID left join product_exclusiveinfo on product_baseinfo.productID = product_exclusiveinfo.ProductID and product_exclusiveinfo.status = 1 left join product_presell on product_presell.productID = product_baseinfo.ProductID and product_presell.status = 1 inner join (SELECT distinct(product_inventory.ProductID) FROM mktypproductdb.product_inventory inner join mktypproductdb.product_sku on product_inventory.SKUID = product_sku.SKUID and product_inventory.ProductID = product_sku.ProductID where product_sku.SKUStatus = 1 and CurrentCount > 0 ) as tbinventory on
	//   tbinventory.ProductID = product_baseinfo.productID where product_baseinfo.ProductStatus = 4 and product_baseinfo.ProductType =
	//    0 and product_baseinfo.ProductionLineID = 1 and product_sellcount.SellPoint > 0 and product_sellcount.SellPoint < 1044 and 
	//    product_baseinfo.IsVirtual = 0 order by SellPoint, ProductID desc limit 12 offset 0`
	//sql:="select CASE WHEN HOST LIKE ? THEN LEFT ( HOST , POSITION ( a in b ) - ? ) ELSE LEFT ( HOST , POSITION ( ) - ? ) END "
	//sql:="SELECT something FROM tbl_name WHERE DATE_SUB(CURDATE(),INTERVAL 30 DAY) <= date_col"
	sql:="select a.b,a.* from b.c join a.c using (a,bc,d)"
	stmt := SqlParse(sql)
	t.Log(stmt)

}
