package parser

import "ast"
import "log"
import "errors"

func SqlParse(sql string) ast.Node {
	mtoken := NewTokener(sql)
	//i := 0
//
	//for mtoken.EOF == false {
	//	id, str := mtoken.Scan()
	//	log.Println(i, id, str)
	//	i = i + 1
	//}
	if yyParse(mtoken) != 0 {
		log.Println(errors.New(mtoken.LastError))
	}
	return mtoken.Stmt()
}
