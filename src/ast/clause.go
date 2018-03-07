package ast

import "io"
import "fmt"

/***********************************
*
*	实现所有的clause定义
*
* 	DistinctClause
* 	TargetClause
*  	IntoClause
*  	FromClause
*  	WhereClause
*  	GroupClause
*  	HavingClause
*  	SortClause
*  	LockClause
*  	LimitClause
*
************************************/

//实现clause 定义

const (
	AST_EMPTY = iota + 1
	AST_ALL
	AST_DISTINCT
	AST_DISTINCTROW
)

var Distinct = map[int]string{
	AST_EMPTY:       "",
	AST_ALL:         "ALL",
	AST_DISTINCT:    "DISTINCT",
	AST_DISTINCTROW: "DISTINCTROW",
}

type DistinctClause struct {
	node
}

func (e *DistinctClause) Accept(v Visitor) (Node, bool) {

	if !v.Notify(e) || e.GetTag() == AST_EMPTY {
		return v.Visit(e)
	}

	return v.Visit(e)
}

func (e *DistinctClause) Format(w io.Writer) {
	switch e.GetTag() {
	case AST_DISTINCT:
		fmt.Fprint(w, "DISTINCT")
	case AST_DISTINCTROW:
		fmt.Fprint(w, "DISTINCTROW")
	}

	fmt.Fprint(w, " ")
}

type TargetClause struct {
	node
	Target_ref []ExprNode
}

func (e *TargetClause) Accept(v Visitor) (Node, bool) {

	if !v.Notify(e) || e.GetTag() == AST_EMPTY {
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

func (e *TargetClause) Format(w io.Writer) {
	length := len(e.Target_ref)
	for index, target := range e.Target_ref {
		target.Format(w)
		if index != length-1 {
			fmt.Fprint(w, ",")
		} else {
			fmt.Fprint(w, " ")
		}
	}
}

type IntoClause struct {
	node
	Relation ExprNode
}

func (e *IntoClause) Accept(v Visitor) (Node, bool) {

	if !v.Notify(e) || e.GetTag() == AST_EMPTY {
		return v.Visit(e)
	}

	node, ok := e.Relation.Accept(v)
	if !ok {
		return e, false
	}
	e.Relation = node.(ExprNode)

	return v.Visit(e)
}

func (e *IntoClause) Format(w io.Writer) {
	if e.GetTag() != AST_EMPTY {
		fmt.Fprint(w, " INTO ")
		e.Relation.Format(w)
	}

}

type FromClause struct {
	node
	Table_ref []ExprNode
}

func (e *FromClause) Accept(v Visitor) (Node, bool) {

	if !v.Notify(e) || e.GetTag() == AST_EMPTY {
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

func (e *FromClause) Format(w io.Writer) {
	if e.GetTag() != AST_EMPTY {
		fmt.Fprint(w, " FROM ")
		length := len(e.Table_ref)
		for index, target := range e.Table_ref {
			target.Format(w)
			if index != length-1 {
				fmt.Fprint(w, ",")
			} else {
				fmt.Fprint(w, " ")
			}
		}
	}
}

type WhereClause struct {
	node
	Where ExprNode
}

func (e *WhereClause) Accept(v Visitor) (Node, bool) {

	if !v.Notify(e) || e.GetTag() == AST_EMPTY {
		return v.Visit(e)
	}

	node, ok := e.Where.Accept(v)
	if !ok {
		return e, false
	}
	e.Where = node.(ExprNode)

	return v.Visit(e)
}

func (e *WhereClause) Format(w io.Writer) {
	if e.GetTag() != AST_EMPTY {
		fmt.Fprint(w, " WHERE ")
		e.Where.Format(w)
	}
}

type GroupClause struct {
	node
	Target_ref []ExprNode
}

func (e *GroupClause) Accept(v Visitor) (Node, bool) {

	if !v.Notify(e) || e.GetTag() == AST_EMPTY {
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

func (e *GroupClause) Format(w io.Writer) {
	if e.GetTag() != AST_EMPTY {
		fmt.Fprint(w, " GROUP BY ")
		length := len(e.Target_ref)
		for index, target := range e.Target_ref {
			target.Format(w)
			if index != length-1 {
				fmt.Fprint(w, ",")
			} else {
				fmt.Fprint(w, " ")
			}
		}
	}
}

type HavingClause struct {
	node
	Having ExprNode
}

func (e *HavingClause) Accept(v Visitor) (Node, bool) {

	if !v.Notify(e) || e.GetTag() == AST_EMPTY {
		return v.Visit(e)
	}

	node, ok := e.Having.Accept(v)
	if !ok {
		return e, false
	}
	e.Having = node.(ExprNode)
	return v.Visit(e)
}

func (e *HavingClause) Format(w io.Writer) {
	if e.GetTag() != AST_EMPTY {
		fmt.Fprint(w, " HAVING ")
		e.Having.Format(w)
	}
}

type SortClause struct {
	node
	Target_ref []ExprNode
}

func (e *SortClause) Accept(v Visitor) (Node, bool) {

	if !v.Notify(e) || e.GetTag() == AST_EMPTY {
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

func (e *SortClause) Format(w io.Writer) {
	if e.GetTag() != AST_EMPTY {
		fmt.Fprint(w, " ORDER BY ")
		length := len(e.Target_ref)
		for index, target := range e.Target_ref {
			target.Format(w)
			if index != length-1 {
				fmt.Fprint(w, ",")
			}
		}
	}
}

const (
	Lock_update = iota + 2
	Lock_share
)

type LockClause struct {
	node
}

func (e *LockClause) Accept(v Visitor) (Node, bool) {

	if !v.Notify(e) || e.GetTag() == AST_EMPTY {
		return v.Visit(e)
	}

	return v.Visit(e)
}

func (e *LockClause) Format(w io.Writer) {
	switch e.GetTag() {
	case Lock_update:
		fmt.Fprint(w, " FOR UPDATE ")
	case Lock_share:
		fmt.Fprint(w, " FOR IN SHARE MODE ")
	}
}

//定义Limitclause
type LimitClause struct {
	node
	Limit  []ExprNode
	Offset ExprNode
}

func (e *LimitClause) Accept(v Visitor) (Node, bool) {

	if !v.Notify(e) || e.GetTag() == AST_EMPTY {
		return v.Visit(e)
	}

	for index, target := range e.Limit {
		node, ok := target.Accept(v)
		if !ok {
			return e, false
		}
		e.Limit[index] = node.(ExprNode)

	}

	if e.Offset != nil {
		node, ok := e.Offset.Accept(v)
		if !ok {
			return e, false
		}
		e.Offset = node.(ExprNode)
	}

	return v.Visit(e)
}

func (e *LimitClause) Format(w io.Writer) {
	if e.GetTag() != AST_EMPTY {
		if e.Limit != nil {
			fmt.Fprint(w, " LIMIT ")
			length := len(e.Limit)
			for index, target := range e.Limit {
				target.Format(w)
				if index != length-1 {
					fmt.Fprint(w, ",")
				} else {
					fmt.Fprint(w, " ")
				}
			}
		}

		if e.Offset != nil {
			fmt.Fprint(w, " OFFSET ")
			e.Offset.Format(w)
		}
	}
}
