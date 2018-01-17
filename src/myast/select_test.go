package myast

import "testing"

func Test_NewSelectStmt(t *testing.T) {
	stmt := NewSelectStmt(SETOP_SELECT, true)
	if stmt.NodeType == SETOP_SELECT && stmt.HavingAll == true {
		t.Log("第一个测试通过", stmt)
	} else {
		t.Error("第一个测试失败", stmt)
	}
}
