package myast

import "testing"

func Test_newlist(t *testing.T) {
	s := NewValueExpr(VALUE_IDENT)
	s.Val = "a"
	b := List{}
	b = append(b, s)
	t.Log(b[0])
}
