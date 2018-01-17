package myast

import "testing"

func Test_newList(t *testing.T) {
	a := NewValueExpr(3)
	a.Val = "a"
	c := NewList(a)
	t.Log(c)
}
