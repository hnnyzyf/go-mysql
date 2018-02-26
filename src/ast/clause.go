package ast

//实现clause 定义

const (
	AST_EMPTY = iota
	AST_ALL
	AST_DISTINCT
	AST_DISTINCTROW
)

type DistinctClause struct {
	node
}

func (e *DistinctClause) Accept(v Visitor) (Node, bool) {
	if v == nil {
		return e, false
	}

	if !v.Notify(e) {
		return v.Visit(e)
	}

	return v.Visit(e)
}

type TargetClause struct {
	node
	Target_ref []ExprNode
}

func (e *TargetClause) Accept(v Visitor) (Node, bool) {
	if v == nil {
		return e, false
	}

	if !v.Notify(e) {
		return v.Visit(e)
	}

	for index, target := range e.Target_ref {
		node, ok := target.Accept(v)
		if !ok {
			return e, false
		}
		e.Target_ref[index] = node.(ExprNode)

	}

	return v.Visit(e)
}

type IntoClause struct {
	node
	Relation ExprNode
}

func (e *IntoClause) Accept(v Visitor) (Node, bool) {
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
	e.Relation = node.(ExprNode)

	return v.Visit(e)
}

type FromClause struct {
	node
	Table_ref []ExprNode
}

func (e *FromClause) Accept(v Visitor) (Node, bool) {
	if v == nil {
		return e, false
	}

	if !v.Notify(e) {
		return v.Visit(e)
	}

	for index, table := range e.Table_ref {
		node, ok := table.Accept(v)
		if !ok {
			return e, false
		}
		e.Table_ref[index] = node.(ExprNode)

	}

	return v.Visit(e)
}

type WhereClause struct {
	node
	Where ExprNode
}

func (e *WhereClause) Accept(v Visitor) (Node, bool) {
	if v == nil {
		return e, false
	}

	if !v.Notify(e) {
		return v.Visit(e)
	}

	node, ok := e.Where.Accept(v)
	if !ok {
		return e, false
	}
	e.Where = node.(ExprNode)
	return v.Visit(e)
}

type GroupClause struct {
	node
	Target_ref []ExprNode
}

func (e *GroupClause) Accept(v Visitor) (Node, bool) {
	if v == nil {
		return e, false
	}

	if !v.Notify(e) {
		return v.Visit(e)
	}

	for index, target := range e.Target_ref {
		node, ok := target.Accept(v)
		if !ok {
			return e, false
		}
		e.Target_ref[index] = node.(ExprNode)

	}

	return v.Visit(e)
}

type HavingClause struct {
	node
	Having ExprNode
}

func (e *HavingClause) Accept(v Visitor) (Node, bool) {
	if v == nil {
		return e, false
	}

	if !v.Notify(e) {
		return v.Visit(e)
	}

	node, ok := e.Having.Accept(v)
	if !ok {
		return e, false
	}
	e.Having = node.(ExprNode)
	return v.Visit(e)
}

type SortClause struct {
	node
	Target_ref []ExprNode
}

func (e *SortClause) Accept(v Visitor) (Node, bool) {
	if v == nil {
		return e, false
	}

	if !v.Notify(e) {
		return v.Visit(e)
	}

	for index, target := range e.Target_ref {
		node, ok := target.Accept(v)
		if !ok {
			return e, false
		}
		e.Target_ref[index] = node.(ExprNode)

	}

	return v.Visit(e)
}

const (
	Lock_update = iota
	Lock_share
)

type LockClause struct {
	node
}

func (e *LockClause) Accept(v Visitor) (Node, bool) {
	if v == nil {
		return e, false
	}

	if !v.Notify(e) {
		return v.Visit(e)
	}

	return v.Visit(e)
}

//定义Limitclause
type LimitClause struct {
	node
	Limit  []ExprNode
	Offset ExprNode
}

func (e *LimitClause) Accept(v Visitor) (Node, bool) {
	if v == nil {
		return e, false
	}

	if !v.Notify(e) {
		return v.Visit(e)
	}

	for index, target := range e.Limit {
		node, ok := target.Accept(v)
		if !ok {
			return e, false
		}
		e.Limit[index] = node.(ExprNode)

	}

	node, ok := e.Offset.Accept(v)
	if !ok {
		return e, false
	}
	e.Offset = node.(ExprNode)

	return v.Visit(e)
}
