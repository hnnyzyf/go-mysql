package ast

import "io"
import "fmt"

const (
	Cross_join = iota + 1
	Inner_join
	Straight_join
	Left_join
	Right_join
)

var Rejoin = map[int]string{
	Cross_join:    "INNER",
	Inner_join:    "INNER",
	Straight_join: "STRAIGHT_JOIN",
	Left_join:     "LEFT",
	Right_join:    "RIGHT",
}

type Join struct {
	node

	IsEnclosed bool
	HasNatural bool
	HasOuter   bool

	//Join形式的结构
	Left      ExprNode
	Right     ExprNode
	Condition ExprNode
}

func (e *Join) Accept(v Visitor) (Node, bool) {

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

	if e.Condition != nil {
		node, ok = e.Condition.Accept(v)
		if !ok {
			return e, false
		}
		e.Condition = node.(ExprNode)
	}

	return v.Visit(e)
}

func (e *Join) Format(w io.Writer) {
	if e.IsEnclosed {
		fmt.Fprint(w, "(")
	}
	e.Left.Format(w)
	if e.HasNatural {
		fmt.Fprint(w, " NATURAL ")
	} else {
		fmt.Fprint(w, " ")
	}
	fmt.Fprint(w, Rejoin[e.GetTag()])
	if e.HasOuter {
		fmt.Fprint(w, " OUTER JOIN ")
	} else {
		fmt.Fprint(w, " JOIN ")
	}
	e.Right.Format(w)

	if e.Condition != nil {
		switch e.Condition.(type) {
		case *Using:
			e.Condition.Format(w)
		default:
			fmt.Fprint(w, " ON ")
			e.Condition.Format(w)
		}
	}
	if e.IsEnclosed {
		fmt.Fprint(w, ")")
	}
}

func (e *Join) Type() int {
	return Ast_join
}

func (e *Join) IsSameAs(src ExprNode, v Visitor) bool {
	if src.Type() != e.Type() {
		return false
	}

	if e.GetTag() != src.GetTag() {
		return false
	}

	other := src.(*Join)

	if e.HasNatural != other.HasNatural {
		return false
	}

	if e.HasOuter != other.HasOuter {
		return false
	}

	if !e.Left.IsSameAs(other.Left, v) {
		return false
	}

	if !e.Right.IsSameAs(other.Right, v) {
		return false
	}

	if e.Condition != nil && other.Condition != nil {
		if !e.Condition.IsSameAs(other.Condition, v) {
			return false
		}
	}

	if (e.Condition == nil && other.Condition != nil) || (e.Condition != nil && other.Condition == nil) {
		return false

	}

	return true
}

type Using struct {
	node

	Column []ExprNode
}

func (e *Using) Accept(v Visitor) (Node, bool) {

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

func (e *Using) Format(w io.Writer) {
	fmt.Fprint(w, " USING (")

	length := len(e.Column)
	for index, target := range e.Column {
		target.Format(w)
		if index < length-1 {
			fmt.Fprint(w, ",")
		} else {
			fmt.Fprint(w, ")")
		}
	}
}

func (e *Using) Type() int {
	return Ast_using
}

func (e *Using) IsSameAs(src ExprNode, v Visitor) bool {
	if src.Type() != e.Type() {
		return false
	}

	if e.GetTag() != src.GetTag() {
		return false
	}

	other := src.(*Using)

	if len(e.Column) != len(other.Column) {
		return false
	}

	for _, expr := range e.Column {
		hasSame := false
		for _, otherexpr := range other.Column {
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
