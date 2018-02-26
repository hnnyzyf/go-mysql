package ast

/******************
*	Func Item
*******************/

type FuncCall struct {
	node

	Name string
	Arg  []ExprNode
}

func (e *FuncCall) Accept(v Visitor) (Node, bool) {
	if v == nil {
		return e, false
	}

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

const (
	Aggretion_Avg = iota
	Aggretion_Count
	Aggretion_Sum
	Aggretion_Max
	Aggretion_Min
)

/******************
*	Aggre Item
*******************/

type AggretionFuncCall struct {
	node

	DistinctType int
	HasStar      bool
	Arg          []ExprNode
}

func (e *AggretionFuncCall) Accept(v Visitor) (Node, bool) {
	if v == nil {
		return e, false
	}

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

/******************
*	Cast Item
*******************/
type CastFuncCall struct {
	node

	Expr     ExprNode
	CastType string
}

func (e *CastFuncCall) Accept(v Visitor) (Node, bool) {
	if v == nil {
		return e, false
	}

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

/******************
*	CalTime Item
*******************/

const (
	CalTime_adddate = iota
	CalTime_date_add
	CalTime_timestampadd
	CalTime_subdate
	CalTime_date_sub
	CalTime_timestampsub
)

var FuncTime = map[string]int{
	"adddate":      CalTime_adddate,
	"date_add":     CalTime_date_add,
	"timestampadd": CalTime_timestampadd,
	"subdate":      CalTime_subdate,
	"date_sub":     CalTime_date_sub,
	"timestampsub": CalTime_timestampsub,
}

type CalTimeFuncCall struct {
	node

	Interval *IntervalExpr
}

func (e *CalTimeFuncCall) Accept(v Visitor) (Node, bool) {
	if v == nil {
		return e, false
	}

	if !v.Notify(e) {
		return v.Visit(e)
	}


	return v.Visit(e)
}


type StringFuncCall struct {
	node

	Substr ExprNode
	Str ExprNode
}

func (e *StringFuncCall) Accept(v Visitor) (Node, bool) {
	if v == nil {
		return e, false
	}

	if !v.Notify(e) {
		return v.Visit(e)
	}


	return v.Visit(e)
}
