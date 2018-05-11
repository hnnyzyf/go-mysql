package parser

import (
	"bytes"
	"github.com/hnnyzyf/mysql-go/advisor/ast"
	"testing"
)

func ssql(n ast.Node) string {
	b := bytes.NewBuffer([]byte{})
	n.Format(b)
	return b.String()
}

func Test_parse(t *testing.T) {
	sqls := []string{
		"select _binary '1'",
	}

	for i, sql := range sqls {
		if ast, err := SqlParse(sql); err != nil {
			t.Log("The", i, "test:", sql)
			t.Error("Error:", err, "\n Sql:", sql)
		} else {
			t.Log("The", i, "test:", ssql(ast))
			t.Log("Pass")
		}

	}
}
