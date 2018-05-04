package parser

import "github.com/hnnyzyf/mysql-go/advisor/ast"
import "errors"

func SqlParse(sql string) (ast.Node, error) {
	mtoken := NewTokener(sql)

	if yyParse(mtoken) != 0 {
		return nil, errors.New(mtoken.LastError)
	}

	return mtoken.Stmt(), nil
}
