package myast

import "testing"

func Test_NewExpr(t *testing.T) {
	a := NewBoolExpr()
	t.Log("BoolExpr通过测试")

	b := NewIsExpr()
	t.Log("IsExpr通过测试")

	c := NewRangeExpr()
	t.Log("RangeExpr通过测试")

	d := NewInExpr()
	t.Log("InExpr通过测试")

	e := NewSubQueryExpr()
	t.Log("SubQueryExpr通过测试")


	t.Log(a,b,c,d,e)

}
