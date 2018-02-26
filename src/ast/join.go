package ast

const (
	Cross_join = iota
	Inner_join
	Straight_join
	Left_join
	Right_join
)

type Join struct {
	node

	HasNatural bool
	HasOuter   bool

	//Join形式的结构
	Left     ExprNode
	Right    ExprNode
	Cndition ExprNode
}

func (e *Join) Accept(v Visitor) (Node, bool) {
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

	node, ok = e.Cndition.Accept(v)
	if !ok {
		return e, false
	}
	e.Cndition = node.(ExprNode)

	return v.Visit(e)
}

type Using struct {
	node

	Column []ExprNode
}

func (e *Using) Accept(v Visitor) (Node, bool) {

	if v == nil {
		return e, false
	}

	if !v.Notify(e) {
		return v.Visit(e)
	}

	for index, target := range e.Column {
		node, ok := target.Accept(v)
		if !ok {
			return e, false
		}
		e.Column[index] = node.(ExprNode)

	}

	return v.Visit(e)
}
