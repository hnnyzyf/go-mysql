package ast

import "testing"

func Test_select(t *testing.T) {
	var stmt Node
	stmt = &SelectStmt{}
	switch no := stmt.(type){
	case *SelectStmt:
		no.SetTag(1)
	}

	t.Log(stmt)
}


func Test_select2(t *testing.T) {
	stmt := &SelectStmt{Tag:1}
	t.Log(stmt)
}
