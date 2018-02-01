package ast

//所有一元二元表达式均可意识用
//形如a+b a-b a*b这种形式
type Expr struct {
	node
	Not bool
	Operator string
	Left     ExprNode
	Right    ExprNode
}

func (e *Expr) IsExpr() {}

func (e *Expr) Accept(v Visitor) (Node, bool) {
	if v == nil {
		return e, false
	}

	if !v.Notify(e) {
		return v.Visit(e)
	}

	node, ok := e.Left.Accept(v)
	if !ok {
		return e, false
	}
	e.Left = node.(ExprNode)

	node, ok = e.Right.Accept(v)
	if !ok {
		return e, false
	}
	e.Right = node.(ExprNode)

	return v.Visit(e)
}

//形如 between类型的比较
type BetweenExpr struct {
	node
	Not   bool
	Expr  ExprNode
	Left  ExprNode
	Right ExprNode
}

func (e *BetweenExpr) IsExpr() {}

func (e *BetweenExpr) Accept(v Visitor) (Node, bool) {
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

	node, ok = e.Left.Accept(v)
	if !ok {
		return e, false
	}
	e.Left = node.(ExprNode)

	node, ok = e.Right.Accept(v)
	if !ok {
		return e, false
	}
	e.Right = node.(ExprNode)

	return v.Visit(e)
}

type SubqueryExpr struct{
	node
	Not bool
	Operator string
	Left ExprNode
	Right *SelectStmt 
}

func (e *SubqueryExpr) IsExpr() {}

func (e *SubqueryExpr) Accept(v Visitor) (Node, bool) {
	if v == nil {
		return e, false
	}

	if !v.Notify(e) {
		return v.Visit(e)
	}
	node, ok := e.Left.Accept(v)
	if !ok {
		return e, false
	}
	e.Left = node.(ExprNode)

	node, ok = e.Right.Accept(v)
	if !ok {
		return e, false
	}
	e.Right = node.(*SelectStmt)

	return v.Visit(e)
}


type CaseExpr struct{
	node
	Case  ExprNode
	When  List
	Else  ExprNode
}


func (e *CaseExpr) IsExpr() {}

func (e *CaseExpr) Accept(v Visitor) (Node, bool) {
	if v == nil {
		return e, false
	}

	if !v.Notify(e) {
		return v.Visit(e)
	}

	node, ok := e.Case.Accept(v)
	if !ok {
		return e, false
	}
	e.Case = node.(ExprNode)


	for index, target := range e.When {
		node, ok := target.Accept(v)
		if !ok {
			return e, false
		}
		e.When[index] = node

	}

	node, ok = e.Else.Accept(v)
	if !ok {
		return e, false
	}
	e.Else = node.(ExprNode)
	return v.Visit(e)
}



type IntervalExpr struct{
	node
	Expr  ExprNode
}


func (e *IntervalExpr) IsExpr() {}

func (e *IntervalExpr) Accept(v Visitor) (Node, bool) {
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

type InExpr struct{
	node
	Not bool
	Left ExprNode
	Right List
}


func (e *InExpr) IsExpr() {}

func (e *InExpr) Accept(v Visitor) (Node, bool) {
	if v == nil {
		return e, false
	}

	if !v.Notify(e) {
		return v.Visit(e)
	}

	node, ok := e.Left.Accept(v)
	if !ok {
		return e, false
	}
	e.Left = node.(ExprNode)

	
	for index, target := range e.Right {
		node, ok := target.Accept(v)
		if !ok {
			return e, false
		}
		e.Right[index] = node

	}

	return v.Visit(e)
}



type FuncExpr struct{
	node
	Name string
	Arg List
}


func (e *FuncExpr) IsExpr() {}

func (e *FuncExpr) Accept(v Visitor) (Node, bool) {
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
		e.Arg[index] = node

	}
	
	return v.Visit(e)
}



