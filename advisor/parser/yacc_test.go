package parser

import "testing"
import "github.com/hnnyzyf/mysql-go/advisor/ast"
import "log"
import "errors"
import "bytes"

func SqlParse(sql string) ast.Node {
	mtoken := NewTokener(sql)
	if yyParse(mtoken) != 0 {
		log.Println(errors.New(mtoken.LastError))
	}
	return mtoken.Stmt()
}

func Test_print(t *testing.T) {
	sqls := []string{`select product_baseinfo.* , product_sellcount.showprice, product_sellcount.SellCount,
	 product_sellcount.SellPoint, product_sellcount.ShowPriceStrage, product_sellcount.ShowPriceStrageValue, 
	 product_presell.IsPresell, IFNULL(product_presell.PresellEndTime,null) as presellEndTime, 
	 product_exclusiveinfo.ExclusiveType, IFNULL(product_exclusiveinfo.Status,0) as exclusiveStatus,
	  IFNULL(product_presell.Status,0) as presellStatus, '' as brandName, '' as cateName, 0 as SaleCount 
	  from product_baseinfo inner join mktypproductdb.product_sellcount on product_baseinfo.productID = 
	  product_sellcount.productID left join product_exclusiveinfo on product_baseinfo.productID = product_exclusiveinfo.ProductID and product_exclusiveinfo.status = 1 left join product_presell on product_presell.productID = product_baseinfo.ProductID and product_presell.status = 1 inner join (SELECT distinct(product_inventory.ProductID) FROM mktypproductdb.product_inventory inner join mktypproductdb.product_sku on product_inventory.SKUID = product_sku.SKUID and product_inventory.ProductID = product_sku.ProductID where product_sku.SKUStatus = 1 and CurrentCount > 0 ) as tbinventory on
	   tbinventory.ProductID = product_baseinfo.productID where product_baseinfo.ProductStatus = 4 and product_baseinfo.ProductType =
	    0 and product_baseinfo.ProductionLineID = 1 and product_sellcount.SellPoint > 0 and product_sellcount.SellPoint < 1044 and 
	    product_baseinfo.IsVirtual = 0 order by SellPoint, ProductID desc limit 12 offset 0`,
		"select CASE WHEN HOST LIKE ? THEN LEFT ( HOST , POSITION ( a in b ) - ? ) ELSE LEFT ( HOST , POSITION ( a in b) - ? ) END ",
		"SELECT something FROM tbl_name WHERE DATE_SUB(CURDATE(),INTERVAL 30 DAY) <= date_col",
		"select a.b,a.* from b.c left join a.c using (a,bc,d)",
		`select 1 as a from a left join (a,b,c,d,e) on a=b where a='abcde' `,
		"select case when 1 then 1 else 1 end",
		"select 1 from (a left join (c,d,e) on a=b ) right join (b right join (a,b,c,e) on c =d) on a=d",
		"select 1 from a where e>t || (a>e or c<d)",
	}
	for index, sql := range sqls {
		t.Log(index, sql)
		mtoken := NewTokener(sql)
		i := 0
		for mtoken.EOF == false {
			id, str := mtoken.Scan()
			i = i + 1
			t.Log(i, id, str)
		}
	}

}

func Test_parse(t *testing.T) {
	sqls := []string{`select product_baseinfo.* , product_sellcount.showprice, product_sellcount.SellCount,
	 product_sellcount.SellPoint, product_sellcount.ShowPriceStrage, product_sellcount.ShowPriceStrageValue, 
	 product_presell.IsPresell, IFNULL(product_presell.PresellEndTime,null) as presellEndTime, 
	 product_exclusiveinfo.ExclusiveType, IFNULL(product_exclusiveinfo.Status,0) as exclusiveStatus,
	  IFNULL(product_presell.Status,0) as presellStatus, '' as brandName, '' as cateName, 0 as SaleCount 
	  from product_baseinfo inner join mktypproductdb.product_sellcount on product_baseinfo.productID = 
	  product_sellcount.productID left join product_exclusiveinfo on product_baseinfo.productID = product_exclusiveinfo.ProductID and product_exclusiveinfo.status = 1 left join product_presell on product_presell.productID = product_baseinfo.ProductID and product_presell.status = 1 inner join (SELECT distinct(product_inventory.ProductID) FROM mktypproductdb.product_inventory inner join mktypproductdb.product_sku on product_inventory.SKUID = product_sku.SKUID and product_inventory.ProductID = product_sku.ProductID where product_sku.SKUStatus = 1 and CurrentCount > 0 ) as tbinventory on
	   tbinventory.ProductID = product_baseinfo.productID where product_baseinfo.ProductStatus = 4 and product_baseinfo.ProductType =
	    0 and product_baseinfo.ProductionLineID = 1 and product_sellcount.SellPoint > 0 and product_sellcount.SellPoint < 1044 and 
	    product_baseinfo.IsVirtual = 0 order by SellPoint, ProductID desc limit 12 offset 0`,
		"select CASE WHEN HOST LIKE ? THEN LEFT ( HOST , POSITION ( a in b ) - ? ) ELSE LEFT ( HOST , POSITION ( a in b) - ? ) END ",
		"SELECT something FROM tbl_name WHERE DATE_SUB(CURDATE(),INTERVAL 30 DAY) <= date_col",
		"select a.b,a.* from b.c left join a.c using (a,bc,d)",
		`select 1 as a from a left join (a,b,c,d,e) on a=b where a='abcde' `,
		"select case when 1 then 1 else 1 end",
		"select 1 from (a left join (c,d,e) on a=b ) right join (b right join (a,b,c,e) on c =d) on a=d",
	}
	for index, sql := range sqls {
		s := SqlParse(sql)
		t.Log(index, s)
	}
}

func Test_format(t *testing.T) {
	sqls := []string{`select product_baseinfo.* , product_sellcount.showprice, product_sellcount.SellCount,
	 product_sellcount.SellPoint, product_sellcount.ShowPriceStrage, product_sellcount.ShowPriceStrageValue, 
	 product_presell.IsPresell, IFNULL(product_presell.PresellEndTime,null) as presellEndTime, 
	 product_exclusiveinfo.ExclusiveType, IFNULL(product_exclusiveinfo.Status,0) as exclusiveStatus,
	  IFNULL(product_presell.Status,0) as presellStatus, '' as brandName, '' as cateName, 0 as SaleCount 
	  from product_baseinfo inner join mktypproductdb.product_sellcount on product_baseinfo.productID = 
	  product_sellcount.productID left join product_exclusiveinfo on product_baseinfo.productID = product_exclusiveinfo.ProductID and product_exclusiveinfo.status = 1 left join product_presell on product_presell.productID = product_baseinfo.ProductID and product_presell.status = 1 inner join (SELECT distinct(product_inventory.ProductID) FROM mktypproductdb.product_inventory inner join mktypproductdb.product_sku on product_inventory.SKUID = product_sku.SKUID and product_inventory.ProductID = product_sku.ProductID where product_sku.SKUStatus = 1 and CurrentCount > 0 ) as tbinventory on
	   tbinventory.ProductID = product_baseinfo.productID where product_baseinfo.ProductStatus = 4 and product_baseinfo.ProductType =
	    0 and product_baseinfo.ProductionLineID = 1 and product_sellcount.SellPoint > 0 and product_sellcount.SellPoint < 1044 and 
	    product_baseinfo.IsVirtual = 0 order by SellPoint, ProductID desc limit 12 offset 0`,
		"select CASE WHEN HOST LIKE ? THEN LEFT ( HOST , POSITION ( a in b ) - ? ) ELSE LEFT ( HOST , POSITION ( a in b) - ? ) END ",
		"SELECT something FROM tbl_name WHERE DATE_SUB(CURDATE(),INTERVAL 30 DAY) <= date_col",
		"select a.b,a.* from b.c left join a.c using (a,bc,d)",
		`select 1 as a from a left join (a,b,c,d,e) on a=b where a='abcde' `,
		"select case when 1 then 1 else 1 end",
		"select 1 from (a left join (c,d,e) on a=b ) right join (b right join (a,b,c,e) on c =d) on a=d",
		"select 1 from a where (a>b and c<d and e>t or (a>b or c<d)) or (a<d and c>1) and a in (b)",
	}
	for index, sql := range sqls {
		s := SqlParse(sql)
		w := bytes.NewBuffer([]byte{})
		s.Format(w)
		t.Log(index)
		t.Log(w.String())
	}

}

type TestVistor struct {
	Name int
}

func (t *TestVistor) Notify(n ast.Node) bool {
	if n == nil {
		return false
	}

	switch n.(type) {
	case *ast.SimpleSelectStmt:
		return true
	default:
		return true
	}
	return true
}

func (t *TestVistor) Visit(n ast.Node) (ast.Node, bool) {
	if n == nil {
		return n, true
	}

	switch node := n.(type) {
	case *ast.SimpleSelectStmt:
		return node, true
	default:
		return node, true
	}
	return n, true
}

func Test_simple(t *testing.T) {
	sqls := []string{`select product_baseinfo.* , product_sellcount.showprice, product_sellcount.SellCount,
	 product_sellcount.SellPoint, product_sellcount.ShowPriceStrage, product_sellcount.ShowPriceStrageValue, 
	 product_presell.IsPresell, IFNULL(product_presell.PresellEndTime,null) as presellEndTime, 
	 product_exclusiveinfo.ExclusiveType, IFNULL(product_exclusiveinfo.Status,0) as exclusiveStatus,
	  IFNULL(product_presell.Status,0) as presellStatus, '' as brandName, '' as cateName, 0 as SaleCount 
	  from product_baseinfo inner join mktypproductdb.product_sellcount on product_baseinfo.productID = 
	  product_sellcount.productID left join product_exclusiveinfo on product_baseinfo.productID = product_exclusiveinfo.ProductID and product_exclusiveinfo.status = 1 left join product_presell on product_presell.productID = product_baseinfo.ProductID and product_presell.status = 1 inner join (SELECT distinct(product_inventory.ProductID) FROM mktypproductdb.product_inventory inner join mktypproductdb.product_sku on product_inventory.SKUID = product_sku.SKUID and product_inventory.ProductID = product_sku.ProductID where product_sku.SKUStatus = 1 and CurrentCount > 0 ) as tbinventory on
	   tbinventory.ProductID = product_baseinfo.productID where product_baseinfo.ProductStatus = 4 and product_baseinfo.ProductType =
	    0 and product_baseinfo.ProductionLineID = 1 and product_sellcount.SellPoint > 0 and product_sellcount.SellPoint < 1044 and 
	    product_baseinfo.IsVirtual = 0 order by SellPoint, ProductID desc limit 12 offset 0`,
		"select CASE WHEN HOST LIKE ? THEN LEFT ( HOST , POSITION ( a in b ) - ? ) ELSE LEFT ( HOST , POSITION ( a in b) - ? ) END ",
		"SELECT something FROM tbl_name WHERE DATE_SUB(CURDATE(),INTERVAL 30 DAY) <= date_col",
		"select a.b,a.* from b.c join a.c using (a,bc,d)",
		`select 1 as a from a join (a,b,c,d,e) where a='abcde' `,
		"select case when 1 then 1 else 1 end",
	}
	for index, sql := range sqls {
		s := SqlParse(sql)
		v := &TestVistor{Name: 1}
		t.Log(index)
		t.Log(s.Accept(v))
	}
}

func Test_expr(t *testing.T) {
	//sql := "select 1 from test where ((a>b && a<b and a>b )|| (a>b or a<b)) or (a<b and a>b) and a in (1)"
	sql := "select 1 from test where 1 or (a>b && (!a<b or ~a>b || (!\"a\">b or +a<null))) or (a<b or a>b) or a in (1)"
	s := SqlParse(sql)
	w := bytes.NewBuffer([]byte{})
	s.Format(w)
	t.Log(w.String())

	c := s.(*ast.SimpleSelectStmt).Where.Where.(*ast.And0OrExpr)

	for _, x := range c.Expr {
		t.Log(x)
	}

	t.Log(c)

}
