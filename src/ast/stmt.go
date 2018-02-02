package ast

//实现StmtNode接口定义
type SelectStmt struct {
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

	//左右子树，只在SelectStmt类型为非Simple时使用
	All   string
	Left  *SelectStmt
	Right *SelectStmt

	//select语句的其他信息
	Sort  *SortClause
	Lock  *LockClause
	Limit *LimitClause
}

//实现stmt接口
func (s *SelectStmt) IsStmt() {}

//实现Accept接口
func (s *SelectStmt) Accept(v Visitor) (Node, bool) {
	if v == nil {
		return s, false
	}

	//允许访问子节点的话,返回
	if !v.Notify(s) {
		return v.Visit(s)
	}

	if s.Tag == AST_SELECT_SIMPLE {
		node, ok := s.Distinct.Accept(v)
		if !ok {
			return s, false
		}
		s.Distinct = node.(*DistinctClause)

		node, ok = s.Target.Accept(v)
		if !ok {
			return s, false
		}
		s.Target = node.(*TargetClause)

		node, ok = s.Into.Accept(v)
		if !ok {
			return s, false
		}
		s.Into = node.(*IntoClause)

		node, ok = s.From.Accept(v)
		if !ok {
			return s, false
		}
		s.From = node.(*FromClause)

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
	} else {
		node, ok := s.Left.Accept(v)
		if !ok {
			return s, false
		}
		s.Left = node.(*SelectStmt)

		node, ok = s.Right.Accept(v)
		if !ok {
			return s, false
		}
		s.Right = node.(*SelectStmt)
	}

	node, ok := s.Sort.Accept(v)
	if !ok {
		return s, false
	}
	s.Sort = node.(*SortClause)

	node, ok = s.Lock.Accept(v)
	if !ok {
		return s, false
	}
	s.Lock = node.(*LockClause)

	node, ok = s.Limit.Accept(v)
	if !ok {
		return s, false
	}
	s.Limit = node.(*LimitClause)

	//正常处理完后，返回处理完的结果
	return v.Visit(s)

}
