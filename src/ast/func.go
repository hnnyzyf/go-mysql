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

	length := len(e.Arg)
	for index, target := range e.Arg {
		target.Format(w)
		if index != length-1 {
			fmt.Fprint(w, ",")
		}
	}
	fmt.Fprint(w, ")")

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
