package myast

//定义DIstinctclause
type Distinctclause struct {
	NType NodeType
}

func NewDistinctclause() *Distinctclause {
	return &Distinctclause{NType: CLAUSE_DISTINCT}
}

func (c *Distinctclause) Accept(v Visitor) {
	v.Walk(c)
}

//定义DIstinctclause
type Intoclause struct {
	NType  NodeType
	Tbname string
}

func NewIntoclause() *Intoclause {
	return &Intoclause{NType: CLAUSE_INTO}
}

func (c *Intoclause) Accept(v Visitor) {
	v.Walk(c)
}

//定义Fromclause
type Fromclause struct {
	NType NodeType
	Expr  List
}

func NewFromclause() *Fromclause {
	return &Fromclause{NType: CLAUSE_FROM}
}

func (c *Fromclause) Accept(v Visitor) {
	v.Walk(c)
}

//定义Whereclause
type Whereclause struct {
	NType NodeType
	Expr  ExprNode
}

func NewWhereclause() *Whereclause {
	return &Whereclause{NType: CLAUSE_WHERE}
}

func (c *Whereclause) Accept(v Visitor) {
	v.Walk(c)
}

//定义Whereclause
type Groupclause struct {
	NType NodeType
	Expr  List
}

func NewGroupclause() *Groupclause {
	return &Groupclause{NType: CLAUSE_GROUPBY}
}

func (c *Groupclause) Accept(v Visitor) {
	v.Walk(c)
}

//定义Havingclause
type Havingclause struct {
	NType NodeType
	Expr  ExprNode
}

func NewHavingclause() *Havingclause {
	return &Havingclause{NType: CLAUSE_HAVING}
}

func (c *Havingclause) Accept(v Visitor) {
	v.Walk(c)
}

//定义Sortclause
type Sortclause struct {
	NType NodeType
	Expr  List
}

func NewSortclause() *Sortclause {
	return &Sortclause{NType: CLAUSE_SORTBY}
}

func (c *Sortclause) Accept(v Visitor) {
	v.Walk(c)
}

//定义Lockingclause
type Lockingclause struct {
	NType NodeType
	Expr  string
}

func NewLockingclause() *Lockingclause {
	return &Lockingclause{NType: CLAUSE_LOCK}
}

func (c *Lockingclause) Accept(v Visitor) {
	v.Walk(c)
}

//定义Limitclause
type Limitclause struct {
	NType  NodeType
	Limit  List
	Offset ExprNode
}

func NewLimitclause() *Limitclause {
	return &Limitclause{NType: CLAUSE_LIMIT}
}

func (c *Limitclause) Accept(v Visitor) {
	v.Walk(c)
}
