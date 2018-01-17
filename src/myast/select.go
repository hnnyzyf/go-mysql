package myast

type SetOption int

const (
	_ SetOption = iota
	SETOP_SIMPLE
	SETOP_UNION
	SETOP_INTERSECT
	SETOP_EXCEPT
)

//定义select接口
type SelectNode interface {
	StmtNode
	IsSelectStmt()
}

/*****************************************************************************
 *
 *	Union Except Intereset for SELECT
 *
 *****************************************************************************/

type SelectStmt struct {
	NType  StmtType
	Having string
	Op     SetOption
	Left   SelectNode
	Right  SelectNode
	Sortby Node
	Lock   Node
	Limit  Node
}

//初始化函数
func NewSelectStmt() *SelectStmt {
	return &SelectStmt{}
}

//实现Selectnode接口
func (s *SelectStmt) IsSelectStmt() {}

//实现stmtnode接口
func (s *SelectStmt) IsStmt() {}

//实现node接口
func (s *SelectStmt) Accept(v Visitor) {
	v.Walk(s)
}

/*****************************************************************************
 *
 *	simple select for SELECT
 *
 *****************************************************************************/

//实现select接口
//simpleselect的定义
type SimpleSelectStmt struct {
	//Stmt的类型
	NType    StmtType
	Op       SetOption
	Distinct Node
	Target   List
	Into     Node
	From     Node
	Where    Node
	Groupby  Node
	Having   Node
	Sortby   Node
	Lock     Node
	Limit    Node
}

//初始化函数
func NewSimpleSelectStmt() *SimpleSelectStmt {
	return &SimpleSelectStmt{}
}

//实现Selectnode接口
func (s *SimpleSelectStmt) IsSelectStmt() {}

//实现stmtnode接口
func (s *SimpleSelectStmt) IsStmt() {}

//实现node接口
func (s *SimpleSelectStmt) Accept(v Visitor) {
	v.Walk(s)

}
