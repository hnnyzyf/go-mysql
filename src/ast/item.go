package ast

/******************
*	Relation Item
*******************/

type Relation struct {
	node

	//传统的Relation形式
	DB    string
	Table string

	//关系指向的别名
	Alias string
}

func (e *Relation) Accept(v Visitor) (Node, bool) {
	if v == nil {
		return e, false
	}

	if !v.Notify(e) {
		return v.Visit(e)
	}

	return v.Visit(e)
}

/******************
*	Relation Item
*******************/

type Tuple struct {
	node

	Ref ExprNode
	//关系指向的别名
	Alias string
}

func (e *Tuple) Accept(v Visitor) (Node, bool) {
	if v == nil {
		return e, false
	}

	if !v.Notify(e) {
		return v.Visit(e)
	}

	node, ok := e.Ref.Accept(v)
	if !ok {
		return e, false
	}
	e.Ref = node.(ExprNode)

	return v.Visit(e)
}

/******************
*	Column Item
*******************/

type Column struct {
	node

	IsStar   bool
	Relation *Relation
	Column   string
}

func (e *Column) Accept(v Visitor) (Node, bool) {
	if v == nil {
		return e, false
	}

	if !v.Notify(e) {
		return v.Visit(e)
	}

	node, ok := e.Relation.Accept(v)
	if !ok {
		return e, false
	}
	e.Relation = node.(*Relation)

	return v.Visit(e)
}

const (
	Bit = iota
	Hex
	Digit
)

/******************
*	Number Item
*******************/

type Number struct {
	node
	Symbol bool
	Val    string
}

func (e *Number) Accept(v Visitor) (Node, bool) {
	if v == nil {
		return e, false
	}

	if !v.Notify(e) {
		return v.Visit(e)
	}

	return v.Visit(e)
}

/******************
*	String Item
*******************/

type String struct {
	node
	Str string
}

func (e *String) Accept(v Visitor) (Node, bool) {
	if v == nil {
		return e, false
	}

	if !v.Notify(e) {
		return v.Visit(e)
	}

	return v.Visit(e)
}

/******************
*	True Item
*******************/

type True struct {
	node
	IsTrue bool
}

func (e *True) Accept(v Visitor) (Node, bool) {
	if v == nil {
		return e, false
	}

	if !v.Notify(e) {
		return v.Visit(e)
	}

	return v.Visit(e)
}

/******************
*	Null Item
*******************/

type Null struct {
	node
}

func (e *Null) Accept(v Visitor) (Node, bool) {
	if v == nil {
		return e, false
	}

	if !v.Notify(e) {
		return v.Visit(e)
	}

	return v.Visit(e)
}

/******************
*	SystemVariable Item
*******************/

type SystemVariable struct {
	node

	Schema    string
	Parameter string
}

func (e *SystemVariable) Accept(v Visitor) (Node, bool) {
	if v == nil {
		return e, false
	}

	if !v.Notify(e) {
		return v.Visit(e)
	}

	return v.Visit(e)
}

/******************
*	UserVariable Item
*******************/

type UserVariable struct {
	node

	Variable string
}

func (e *UserVariable) Accept(v Visitor) (Node, bool) {
	if v == nil {
		return e, false
	}

	if !v.Notify(e) {
		return v.Visit(e)
	}

	return v.Visit(e)
}

/******************
*	Marker Item
*******************/

type Marker struct {
	node

	Str string
}

func (e *Marker) Accept(v Visitor) (Node, bool) {
	if v == nil {
		return e, false
	}

	if !v.Notify(e) {
		return v.Visit(e)
	}

	return v.Visit(e)
}

/******************
*	Sort Item
*******************/

const (
	Sort_Asc = iota
	Sort_Desc
	Sort_Empty
)

type Sortby struct {
	node

	Expr    ExprNode
	AscType int
}

func (e *Sortby) Accept(v Visitor) (Node, bool) {
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
