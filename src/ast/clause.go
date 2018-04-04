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
		fmt.Fprint(w, "DISTINCT ")
	case AST_DISTINCTROW:
		fmt.Fprint(w, "DISTINCTROW ")
	}

	fmt.Fprint(w, " ")
}

func (e *DistinctClause) Type() int {
	return Ast_distinctclause
}

func (e *DistinctClause) IsSameAs(src ExprNode, v Visitor) bool {
	if src.Type() != e.Type() {
		return false
	}

	if e.GetTag() != src.GetTag() {
		return false
	}

	return true
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

func (e *TargetClause) Type() int {
	return Ast_targetclause
}

func (e *TargetClause) IsSameAs(src ExprNode, v Visitor) bool {
	if src.Type() != e.Type() {
		return false
	}

	if e.GetTag() != src.GetTag() {
		return false
	}

	other := src.(*TargetClause)
	if len(e.Target_ref) != len(other.Target_ref) {
		return false
	}

	for _, expr := range e.Target_ref {
		hasSame := false
		for _, otherexpr := range other.Target_ref {
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

func (e *IntoClause) Type() int {
	return Ast_intoclause
}

func (e *IntoClause) IsSameAs(src ExprNode, v Visitor) bool {
	if src.Type() != e.Type() {
		return false
	}

	if e.GetTag() != src.GetTag() {
		return false
	}

	other := src.(*IntoClause)

	if e.Relation != nil && other.Relation != nil {
		if !e.Relation.IsSameAs(other.Relation, v) {
			return false
		}
	}

	if (e.Relation != nil && other.Relation == nil) || (e.Relation == nil && other.Relation != nil) {
		return false
	}

	return true
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

func (e *FromClause) Type() int {
	return Ast_fromclause
}

func (e *FromClause) IsSameAs(src ExprNode, v Visitor) bool {
	if src.Type() != e.Type() {
		return false
	}

	if e.GetTag() != src.GetTag() {
		return false
	}

	other := src.(*FromClause)

	if len(e.Table_ref) != len(other.Table_ref) {
		return false
	}

	for _, expr := range e.Table_ref {
		hasSame := false
		for _, otherexpr := range other.Table_ref {
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

func (e *WhereClause) Type() int {
	return Ast_whereclause
}

func (e *WhereClause) IsSameAs(src ExprNode, v Visitor) bool {
	if src.Type() != e.Type() {
		return false
	}

	if e.GetTag() != src.GetTag() {
		return false
	}

	other := src.(*WhereClause)

	if e.Where != nil && other.Where != nil {
		if !e.Where.IsSameAs(other.Where, v) {
			return false
		}
	}

	if (e.Where != nil && other.Where == nil) || (e.Where == nil && other.Where != nil) {
		return false
	}

	return true

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

func (e *GroupClause) Type() int {
	return Ast_groupclause
}

func (e *GroupClause) IsSameAs(src ExprNode, v Visitor) bool {
	if src.Type() != e.Type() {
		return false
	}

	if e.GetTag() != src.GetTag() {
		return false
	}

	other := src.(*GroupClause)
	if len(e.Target_ref) != len(other.Target_ref) {
		return false
	}

	for _, expr := range e.Target_ref {
		hasSame := false
		for _, otherexpr := range other.Target_ref {
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

func (e *HavingClause) Type() int {
	return Ast_havingclause
}

func (e *HavingClause) IsSameAs(src ExprNode, v Visitor) bool {
	if src.Type() != e.Type() {
		return false
	}

	if e.GetTag() != src.GetTag() {
		return false
	}

	other := src.(*HavingClause)

	if e.Having != nil && other.Having != nil {
		if !e.Having.IsSameAs(other.Having, v) {
			return false
		}
	}

	if (e.Having != nil && other.Having == nil) || (e.Having == nil && other.Having != nil) {
		return false
	}
	return true
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

func (e *SortClause) Type() int {
	return Ast_sortclause
}

func (e *SortClause) IsSameAs(src ExprNode, v Visitor) bool {
	if src.Type() != e.Type() {
		return false
	}

	if e.GetTag() != src.GetTag() {
		return false
	}

	other := src.(*SortClause)
	if len(e.Target_ref) != len(other.Target_ref) {
		return false
	}

	for _, expr := range e.Target_ref {
		hasSame := false
		for _, otherexpr := range other.Target_ref {
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

func (e *LockClause) Type() int {
	return Ast_lockclause
}

func (e *LockClause) IsSameAs(src ExprNode, v Visitor) bool {
	if src.Type() != e.Type() {
		return false
	}

	if e.GetTag() != src.GetTag() {
		return false
	}

	return true

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

func (e *LimitClause) Type() int {
	return Ast_limitclause
}

func (e *LimitClause) IsSameAs(src ExprNode, v Visitor) bool {
	if src.Type() != e.Type() {
		return false
	}

	if e.GetTag() != src.GetTag() {
		return false
	}

	other := src.(*LimitClause)
	if len(e.Limit) != len(other.Limit) {
		return false
	}

	for _, expr := range e.Limit {
		hasSame := false
		for _, otherexpr := range other.Limit {
			if expr.IsSameAs(otherexpr, v) {
				hasSame = true
				break
			}
		}
		if !hasSame {
			return false
		}
	}

	if e.Offset != nil && other.Offset != nil {
		if !e.Offset.IsSameAs(other.Offset, v) {
			return false
		}
	}

	if (e.Offset != nil && other.Offset == nil) || (e.Offset == nil && other.Offset != nil) {
		return false
	}

	return true

}
