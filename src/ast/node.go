package ast

//希望所有的节点都有类型和文本定义，获得类型和文本直接从Text中和Tag获取
//所有的节点都会嵌套node
type node struct {
	Tag     AstOption //节点的类型
	Collate string
	Text    string //节点的Text形式
	Alias   string
	Asc     string
}

func (n node) SetTag(t AstOption) {
	n.Tag = t
}

func (n node) SetCollate(c string) {
	n.Collate = c
}

func (n node) SetText(text string) {
	n.Text = text
}

func (n node) SetAlias(a string) {
	n.Alias = a
}

func (n node) SetAsc(a string) {
	n.Asc = a
}

func (n node) GetTag() AstOption{
	return n.Tag
}


//定义stmt
//所有的Stmt均要实现如下的方法
type StmtNode interface {
	IsStmt()
	Node
	Set
}

//定义expr
//所有的表达式均实现如下的方法

type ExprNode interface {
	IsExpr()
	Node
	Set
}

