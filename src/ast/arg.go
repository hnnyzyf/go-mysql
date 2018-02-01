package ast 


type ValExpr struct{
	node
	Symbol int 	//-1代表正，1代表负
	Sval string
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


type ColumnExpr struct{
	node

	DB string
	Table string
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


type RelationExpr struct{
	node

	//传统的表结构
	DB string
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

type RelationListExpr struct{
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


type JoinExpr struct{	
	node

	Jtype string
	Outer string
	//Join形式的结构
	Left ExprNode
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


