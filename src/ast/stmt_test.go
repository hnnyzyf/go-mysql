package ast

import "testing"

func Test_select(t *testing.T) {
	var stmt Node
	stmt = &SelectStmt{}
	stmt.SetTag(1)
	//switch no := stmt.(type){
	//case *SelectStmt:
	//	no.SetTag(1)
	//}

	t.Log(stmt)
}


func Test_select2(t *testing.T) {
	stmt := &SelectStmt{}
	t.Log(stmt)
}
