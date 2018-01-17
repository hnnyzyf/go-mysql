package myast

import "testing"


func Test_DoValueexpr(t *testing.T){
	a:=NewValueExpr(VALUE_STRING)
	a.Val="abcdefg"
	t.Log(a.Val)
}