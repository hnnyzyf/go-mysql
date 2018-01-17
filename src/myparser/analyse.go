package myparser

import AST "myast"
import "log"
import "errors"

func SqlParse(sql string) AST.Node {
	//sql := "select b,c,d,e from x,y"
	mtoken := NewTokener(sql)
	if yyParse(mtoken) != 0 {
		log.Println(errors.New(mtoken.LastError))
	}
	return mtoken.ASTree.(*AST.SimpleSelectStmt)
}
