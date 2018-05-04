package ast

import "io"
import "fmt"
import "strings"

/***********************************
*
*	实现Func表达式定义
*
* 	FuncCall
* 	AggregationFuncCall
*  	CastFuncCall
*  	CalTimeFuncCall
*  	StringFuncCall
*
************************************/

/******************
*	Func Item
*******************/

type FuncCall struct {
	node

	Name string
	Arg  []ExprNode
}

func (e *FuncCall) Accept(v Visitor) (Node, bool) {

	if !v.Notify(e) {
		return v.Visit(e)
	}

	for index, target := range e.Arg {
		node, ok := target.Accept(v)
		if !ok {
			return e, false
		}
		e.Arg[index] = node.(ExprNode)

	}

	return v.Visit(e)
}

func (e *FuncCall) Format(w io.Writer) {
	fmt.Fprint(w, strings.ToUpper(e.Name))
	fmt.Fprint(w, "(")
	length := len(e.Arg)
	for index, target := range e.Arg {
		target.Format(w)
		if index != length-1 {
			fmt.Fprint(w, ",")
		}
	}
	fmt.Fprint(w, ")")

}

func (e *FuncCall) Type() int {
	return Ast_funccall
}

func (e *FuncCall) IsSameAs(src ExprNode, v Visitor) bool {
	if src.Type() != e.Type() {
		return false
	}

	if e.GetTag() != src.GetTag() {
		return false
	}

	other := src.(*FuncCall)

	if e.Name != other.Name {
		return false
	}

	if len(e.Arg) != len(other.Arg) {
		return false
	}
	for _, expr := range e.Arg {
		hasSame := false
		for _, otherexpr := range other.Arg {
			if expr.IsSameAs(otherexpr, v) {
				hasSame = true
				break
			}
		}
		if !hasSame {
			return false
		}
	}

	return true
}

const (
	Aggregation_Avg = iota + 1
	Aggregation_Count
	Aggregation_Sum
	Aggregation_Max
	Aggregation_Min
)

var Agg = map[int]string{
	Aggregation_Avg:   "AVG",
	Aggregation_Count: "COUNT",
	Aggregation_Sum:   "SUM",
	Aggregation_Max:   "MAX",
	Aggregation_Min:   "MIN",
}

/******************
*	Aggre Item
*******************/

type AggregationFuncCall struct {
	node

	DistinctType int
	Arg          []ExprNode
}

func (e *AggregationFuncCall) Accept(v Visitor) (Node, bool) {

	if !v.Notify(e) {
		return v.Visit(e)
	}

	for index, target := range e.Arg {
		node, ok := target.Accept(v)
		if !ok {
			return e, false
		}
		e.Arg[index] = node.(ExprNode)

	}

	return v.Visit(e)
}

func (e *AggregationFuncCall) Format(w io.Writer) {
	fmt.Fprint(w, Agg[e.GetTag()])
	fmt.Fprint(w, "(")
	fmt.Fprint(w, Distinct[e.DistinctType])
	fmt.Fprint(w, " ")
	length := len(e.Arg)
	for index, target := range e.Arg {
		target.Format(w)
		if index != length-1 {
			fmt.Fprint(w, ",")
		}
	}
	fmt.Fprint(w, ")")

}

func (e *AggregationFuncCall) Type() int {
	return Ast_aggregationfunccall
}

func (e *AggregationFuncCall) IsSameAs(src ExprNode, v Visitor) bool {
	if src.Type() != e.Type() {
		return false
	}

	if e.GetTag() != src.GetTag() {
		return false
	}

	other := src.(*AggregationFuncCall)

	if e.DistinctType != other.DistinctType {
		return false
	}

	if len(e.Arg) != len(other.Arg) {
		return false
	}
	for _, expr := range e.Arg {
		hasSame := false
		for _, otherexpr := range other.Arg {
			if expr.IsSameAs(otherexpr, v) {
				hasSame = true
				break
			}
		}
		if !hasSame {
			return false
		}
	}

	return true
}

/******************
*	Cast Item
*******************/
type CastFuncCall struct {
	node

	Expr     ExprNode
	CastType string
}

func (e *CastFuncCall) Accept(v Visitor) (Node, bool) {

	if !v.Notify(e) {
		return v.Visit(e)
	}

	node, ok := e.Expr.Accept(v)
	if !ok {
		return e, false
	}
	e.Expr = node.(ExprNode)

	return v.Visit(e)
}

func (e *CastFuncCall) Format(w io.Writer) {
	fmt.Fprint(w, "CAST(")
	e.Expr.Format(w)
	fmt.Fprint(w, " AS ")
	fmt.Fprint(w, e.CastType)
	fmt.Fprint(w, ")")

}

func (e *CastFuncCall) Type() int {
	return Ast_castfunccall
}

func (e *CastFuncCall) IsSameAs(src ExprNode, v Visitor) bool {
	if src.Type() != e.Type() {
		return false
	}

	if e.GetTag() != src.GetTag() {
		return false
	}

	other := src.(*CastFuncCall)

	if e.CastType != other.CastType {
		return false
	}

	if !e.Expr.IsSameAs(other.Expr, v) {
		return false
	}

	return true
}

/******************
*	CalTime Item
*******************/

const (
	CalTime_adddate = iota + 1
	CalTime_date_add
	CalTime_timestampadd
	CalTime_subdate
	CalTime_date_sub
	CalTime_timestampsub
)

var FuncTime = map[string]int{
	"ADDDATE":      CalTime_adddate,
	"DATE_ADD":     CalTime_date_add,
	"TIMESTAMPADD": CalTime_timestampadd,
	"SUBDATE":      CalTime_subdate,
	"DATE_SUB":     CalTime_date_sub,
	"TIMESTAMPSUB": CalTime_timestampsub,
}

var ResFuncTime = map[int]string{
	CalTime_adddate:      "ADDDATE",
	CalTime_date_add:     "DATE_ADD",
	CalTime_timestampadd: "TIMESTAMPADD",
	CalTime_subdate:      "SUBDATE",
	CalTime_date_sub:     "DATE_SUB",
	CalTime_timestampsub: "TIMESTAMPSUB",
}

type CalTimeFuncCall struct {
	node

	Interval *IntervalExpr
}

func (e *CalTimeFuncCall) Accept(v Visitor) (Node, bool) {

	if !v.Notify(e) {
		return v.Visit(e)
	}

	return v.Visit(e)
}

func (e *CalTimeFuncCall) Format(w io.Writer) {
	fmt.Fprint(w, ResFuncTime[e.GetTag()])
	fmt.Fprint(w, "(")
	e.Interval.Left.Format(w)
	fmt.Fprint(w, ", INTERVAL ")
	e.Interval.Right.Format(w)
	fmt.Fprint(w, " ")
	fmt.Fprint(w, ReTimeUnit[e.Interval.TimeUnit])
	fmt.Fprint(w, ")")
}

func (e *CalTimeFuncCall) Type() int {
	return Ast_caltimefunccall
}

func (e *CalTimeFuncCall) IsSameAs(src ExprNode, v Visitor) bool {
	if src.Type() != e.Type() {
		return false
	}

	if e.GetTag() != src.GetTag() {
		return false
	}

	other := src.(*CalTimeFuncCall)

	if !e.Interval.IsSameAs(other.Interval, v) {
		return false
	}

	return true
}

type StringFuncCall struct {
	node

	Substr ExprNode
	Str    ExprNode
}

func (e *StringFuncCall) Accept(v Visitor) (Node, bool) {

	if !v.Notify(e) {
		return v.Visit(e)
	}

	return v.Visit(e)
}

func (e *StringFuncCall) Format(w io.Writer) {
	fmt.Fprint(w, "POSITION(")
	e.Substr.Format(w)
	fmt.Fprint(w, " IN ")
	e.Str.Format(w)
	fmt.Fprint(w, ")")
}

func (e *StringFuncCall) Type() int {
	return Ast_stringfunccall
}

func (e *StringFuncCall) IsSameAs(src ExprNode, v Visitor) bool {
	if src.Type() != e.Type() {
		return false
	}

	if e.GetTag() != src.GetTag() {
		return false
	}

	other := src.(*StringFuncCall)

	if !e.Substr.IsSameAs(other.Substr, v) {
		return false
	}

	if !e.Str.IsSameAs(other.Str, v) {
		return false
	}

	return true
}
