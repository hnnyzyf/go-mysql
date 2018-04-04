package ast

import "io"
import "fmt"

//实现StmtNode接口定义
type SimpleSelectStmt struct {
	//包含node和文本定义
	node

	//clasuse定义
	Distinct *DistinctClause
	Target   *TargetClause
	Into     *IntoClause
	From     *FromClause
	Where    *WhereClause
	Group    *GroupClause
	Having   *HavingClause

	//select语句的其他信息
	Sort  *SortClause
	Lock  *LockClause
	Limit *LimitClause
}

//实现Accept接口
func (s *SimpleSelectStmt) Accept(v Visitor) (Node, bool) {

	//允许访问子节点的话,返回
	if !v.Notify(s) {
		return v.Visit(s)
	}

	//优先处理From的信息
	node, ok := s.From.Accept(v)
	if !ok {
		return s, false
	}
	s.From = node.(*FromClause)

	node, ok = s.Target.Accept(v)
	if !ok {
		return s, false
	}
	s.Target = node.(*TargetClause)

	node, ok = s.Where.Accept(v)
	if !ok {
		return s, false
	}
	s.Where = node.(*WhereClause)

	node, ok = s.Group.Accept(v)
	if !ok {
		return s, false
	}
	s.Group = node.(*GroupClause)

	if s.Sort != nil {
		node, ok = s.Sort.Accept(v)
		if !ok {
			return s, false
		}
		s.Sort = node.(*SortClause)
	}

	if s.Lock != nil {
		node, ok = s.Lock.Accept(v)
		if !ok {
			return s, false
		}
		s.Lock = node.(*LockClause)
	}

	if s.Limit != nil {
		node, ok = s.Limit.Accept(v)
		if !ok {
			return s, false
		}
		s.Limit = node.(*LimitClause)
	}

	if s.Distinct != nil {
		node, ok = s.Distinct.Accept(v)
		if !ok {
			return s, false
		}
		s.Distinct = node.(*DistinctClause)
	}

	node, ok = s.Into.Accept(v)
	if !ok {
		return s, false
	}
	s.Into = node.(*IntoClause)
	//正常处理完后，返回处理完的结果
	return v.Visit(s)

}

func (s *SimpleSelectStmt) Format(w io.Writer) {
	fmt.Fprint(w, "( SELECT ")
	if s.Distinct != nil {
		s.Distinct.Format(w)
	}
	s.Target.Format(w)
	s.Into.Format(w)
	s.From.Format(w)
	s.Where.Format(w)
	s.Group.Format(w)
	s.Having.Format(w)

	if s.Sort != nil {
		s.Sort.Format(w)
	}
	if s.Lock != nil {
		s.Lock.Format(w)
	}
	if s.Limit != nil {
		s.Limit.Format(w)
	}
	fmt.Fprint(w, ")")
}

func (e *SimpleSelectStmt) Type() int {
	return Ast_simpleselectstmt
}

func (e *SimpleSelectStmt) IsSameAs(src ExprNode, v Visitor) bool {
	if src.Type() != e.Type() {
		return false
	}

	if e.GetTag() != src.GetTag() {
		return false
	}

	other := src.(*SimpleSelectStmt)

	if e.Distinct != nil && other.Distinct != nil {
		if !e.Distinct.IsSameAs(other.Distinct, v) {
			return false
		}
	}

	if (e.Distinct != nil && other.Distinct == nil) || (e.Distinct == nil && other.Distinct != nil) {
		return false
	}

	if !e.Target.IsSameAs(other.Target, v) {
		return false
	}
	if !e.Into.IsSameAs(other.Into, v) {
		return false
	}
	if !e.From.IsSameAs(other.From, v) {
		return false
	}
	if !e.Where.IsSameAs(other.Where, v) {
		return false
	}
	if !e.Group.IsSameAs(other.Group, v) {
		return false
	}
	if !e.Having.IsSameAs(other.Having, v) {
		return false
	}

	if e.Sort != nil && other.Sort != nil {
		if !e.Sort.IsSameAs(other.Sort, v) {
			return false
		}
	}

	if (e.Sort != nil && other.Sort == nil) || (e.Sort == nil && other.Sort != nil) {
		return false
	}

	if e.Lock != nil && other.Lock != nil {
		if !e.Lock.IsSameAs(other.Lock, v) {
			return false
		}
	}
	if (e.Lock != nil && other.Lock == nil) || (e.Lock == nil && other.Lock != nil) {
		return false
	}

	if e.Limit != nil && other.Limit != nil {
		if !e.Limit.IsSameAs(other.Limit, v) {
			return false
		}
	}

	if (e.Limit != nil && other.Limit == nil) || (e.Limit == nil && other.Limit != nil) {
		return false
	}

	return true
}

const (
	Union     = iota // 0
	Except           // 1
	Intersect        // 2
)

//实现SelectStmt接口定义
type UnionStmt struct {
	//包含node和文本定义
	node

	//左右子树，只在SelectStmt类型为非Simple时使用
	DistinctType int

	Left  Node
	Right Node

	//select语句的其他信息
	Sort  *SortClause
	Lock  *LockClause
	Limit *LimitClause
}

func (s *UnionStmt) Accept(v Visitor) (Node, bool) {

	//允许访问子节点的话,返回
	if !v.Notify(s) {
		return v.Visit(s)
	}

	node, ok := s.Left.Accept(v)
	if !ok {
		return s, false
	}
	s.Left = node

	node, ok = s.Right.Accept(v)
	if !ok {
		return s, false
	}
	s.Right = node

	if s.Sort != nil {
		node, ok = s.Sort.Accept(v)
		if !ok {
			return s, false
		}
		s.Sort = node.(*SortClause)
	}

	if s.Lock != nil {
		node, ok = s.Lock.Accept(v)
		if !ok {
			return s, false
		}
		s.Lock = node.(*LockClause)
	}

	if s.Limit != nil {
		node, ok = s.Limit.Accept(v)
		if !ok {
			return s, false
		}
		s.Limit = node.(*LimitClause)
	}

	//正常处理完后，返回处理完的结果
	return v.Visit(s)
}

func (s *UnionStmt) Format(w io.Writer) {
	fmt.Fprint(w, "(")
	s.Left.(ExprNode).Format(w)
	fmt.Fprint(w, " UNION ")
	fmt.Fprint(w, Distinct[s.DistinctType])
	fmt.Fprint(w, " ")

	if s.Sort != nil {
		s.Sort.Format(w)
	}
	if s.Lock != nil {
		s.Lock.Format(w)
	}
	if s.Limit != nil {
		s.Limit.Format(w)
	}

	s.Right.(ExprNode).Format(w)
	fmt.Fprint(w, ")")
}

func (e *UnionStmt) Type() int {
	return Ast_unionstmt
}

func (e *UnionStmt) IsSameAs(src ExprNode, v Visitor) bool {
	if src.Type() != e.Type() {
		return false
	}

	if e.GetTag() != src.GetTag() {
		return false
	}

	other := src.(*UnionStmt)

	if e.DistinctType != other.DistinctType {
		return false
	}

	if !e.Left.(ExprNode).IsSameAs(other.Left.(ExprNode), v) {
		return false
	}
	if !e.Right.(ExprNode).IsSameAs(other.Right.(ExprNode), v) {
		return false
	}

	if e.Sort != nil && other.Sort != nil {
		if !e.Sort.IsSameAs(other.Sort, v) {
			return false
		}
	}

	if (e.Sort != nil && other.Sort == nil) || (e.Sort == nil && other.Sort != nil) {
		return false
	}

	if e.Lock != nil && other.Lock != nil {
		if !e.Lock.IsSameAs(other.Lock, v) {
			return false
		}
	}
	if (e.Lock != nil && other.Lock == nil) || (e.Lock == nil && other.Lock != nil) {
		return false
	}

	if e.Limit != nil && other.Limit != nil {
		if !e.Limit.IsSameAs(other.Limit, v) {
			return false
		}
	}

	if (e.Limit != nil && other.Limit == nil) || (e.Limit == nil && other.Limit != nil) {
		return false
	}

	return true
}
