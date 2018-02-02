package ast

//所有一元二元表达式均可意识用
//形如a+b a-b a*b这种形式
type Expr struct {
	node
	Not      bool
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

type SubqueryExpr struct {
	node
	Not      bool
	Operator string
	Left     ExprNode
	Right    *SelectStmt
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

type CaseExpr struct {
	node
	Case ExprNode
	When List
	Else ExprNode
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

type IntervalExpr struct {
	node
	Expr ExprNode
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

type InExpr struct {
	node
	Not   bool
	Left  ExprNode
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

type FuncExpr struct {
	node
	Name string
	Arg  List
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

type ValExpr struct {
	node
	Symbol int //-1代表正，1代表负
	Sval   string
}

func (e *ValExpr) IsExpr() {}

func (e *ValExpr) Accept(v Visitor) (Node, bool) {
	if v == nil {
		return e, false
	}

	if !v.Notify(e) {
		return v.Visit(e)
	}

	return v.Visit(e)
}

type ColumnExpr struct {
	node

	DB     string
	Table  string
	Column string
}

func (e *ColumnExpr) IsExpr() {}

func (e *ColumnExpr) Accept(v Visitor) (Node, bool) {
	if v == nil {
		return e, false
	}

	if !v.Notify(e) {
		return v.Visit(e)
	}

	return v.Visit(e)
}

type RelationExpr struct {
	node

	//传统的表结构
	DB    string
	Table string
}

func (e *RelationExpr) IsExpr() {}

func (e *RelationExpr) Accept(v Visitor) (Node, bool) {
	if v == nil {
		return e, false
	}

	if !v.Notify(e) {
		return v.Visit(e)
	}

	return v.Visit(e)
}

type RelationListExpr struct {
	node

	Relation List
}

func (e *RelationListExpr) IsExpr() {}

func (e *RelationListExpr) Accept(v Visitor) (Node, bool) {
	if v == nil {
		return e, false
	}

	if !v.Notify(e) {
		return v.Visit(e)
	}

	for index, target := range e.Relation {
		node, ok := target.Accept(v)
		if !ok {
			return e, false
		}
		e.Relation[index] = node

	}

	return v.Visit(e)
}

type JoinExpr struct {
	node

	Jtype string
	Outer string
	//Join形式的结构
	Left  ExprNode
	Right ExprNode
	Jqual ExprNode
}

func (e *JoinExpr) IsExpr() {}

func (e *JoinExpr) Accept(v Visitor) (Node, bool) {
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

	node, ok = e.Jqual.Accept(v)
	if !ok {
		return e, false
	}
	e.Jqual = node.(ExprNode)

	return v.Visit(e)
}
