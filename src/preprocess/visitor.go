package preprocess

import "ast"
import "log"
import "meta"
import "server"

/*************************************************
*
* 	PreProcessVisitor
*	定义从ParseTree的预处理，理清每一个Node的关系
*
*  	对clause节点调用对应的Visitor进行处理
*
**************************************************/

type PreProcessVisitor struct {
	ParseTree ast.Node

	//存储每一个SelectStmt的中间状态信息
	State map[*ast.SimpleSelectStmt]*Context

	//Sender用来将任务发送到mysql连接groutine，以获得元信息
	RegisteredServer server.Server

	//当前使用的Context
	CurrentContext *Context

	//Error 信息
	Error []string

	//warning 信息
	Warning []string
}

//每一个PreprocessVistor需要传入两个参数
func NewPreProcessVisitor(p ast.Node, s server.Server) *PreProcessVisitor {
	return &PreProcessVisitor{
		ParseTree:      p,
		State:          make(map[*ast.SimpleSelectStmt]*Context),
		RegisteredServer:        s,
		CurrentContext: nil,
	}
}

func (p *PreProcessVisitor) Notify(n ast.Node) bool {
	switch node := n.(type) {
	case *ast.SimpleSelectStmt:
		//设置CurrentContext域为新的Context
		p.CurrentContext = NewContext(p.CurrentContext)

		p.State[node] = p.CurrentContext

		return true

	case *ast.FromClause, *ast.TargetClause, *ast.WhereClause:
		return false

	default:
		return true
	}
}

func (p *PreProcessVisitor) Visit(n ast.Node) (ast.Node, bool) {
	switch node := n.(type) {
	case *ast.SimpleSelectStmt:
		p.CurrentContext = p.CurrentContext.Last

	case *ast.FromClause:
		p.ProcessFromClause(node)

	case *ast.TargetClause:
		p.ProcessTargetClause(node)

	case *ast.WhereClause:
		if node.GetTag() != ast.AST_EMPTY {
			p.ProcessWhereClause(node)
		}

	}

	return n, true
}



/*************************************************
*
* 	RelationVisitor
*
*	功能:负责遍历FromClause
*  	范围:对于所有的节点,均处理
*
**************************************************/

type RelationVisitor struct {
	Size          int
	ParentVisitor *PreProcessVisitor
}

func NewRelationVisitor(p *PreProcessVisitor) *RelationVisitor {
	return &RelationVisitor{
		Size:          0,
		ParentVisitor: p,
	}
}

func (p *RelationVisitor) Notify(n ast.Node) bool {
	switch n.(type) {
	case *ast.Join, *ast.Relation:
		return true
	default:
		return false
	}
}

func (p *RelationVisitor) Visit(n ast.Node) (ast.Node, bool) {
	switch node := n.(type) {
	case *ast.Relation:
		//设置表的别名
		p.ProcessRelation(node)
		//获取表的元信息
		p.GetMetaInfo(node)
	case *ast.Join:
		return n, true
	default:
		node.Accept(p.ParentVisitor)
	}

	return n, true
}

func (p *RelationVisitor) GetMetaInfo(node *ast.Relation) {
	var Meta meta.Meta
	//获得simpleTable的元信息
	if node.GetTag() == ast.Relation_Table {
		session:=p.ParentVisitor.RegisteredServer.CreateSession2DB(node.DB)
		Meta = p.ParentVisitor.RegisteredServer.Query(session,server.Task_relation,node.Table)
	}

	//获得子查询的元信息
	if node.GetTag() == ast.Relation_Subquery {
		Meta = p.GetTempTableMetaInfo(node.Select)
	}

	p.AddMeta(node, Meta)
}

func (p *RelationVisitor) GetTempTableMetaInfo(node ast.Node) meta.Meta{
	return nil
}

func (p *RelationVisitor) AddMeta(node *ast.Relation, table meta.Meta) {
	p.ParentVisitor.CurrentContext.Meta[node] = table
}



/*************************************************
*
* 	TupleVisitor
*	功能:负责遍历TargetClause
*  	范围:对于所有的节点,均处理
*
**************************************************/

type TupleVisitor struct {
	ParentVisitor *PreProcessVisitor
}

func NewTupleVisitor(p *PreProcessVisitor) *TupleVisitor {
	return &TupleVisitor{
		ParentVisitor: p,
	}
}

func (p *TupleVisitor) Notify(n ast.Node) bool {
	switch n.(type) {
	case *ast.Tuple, *ast.Column:
		return true
	default:
		return false
	}
}

func (p *TupleVisitor) Visit(n ast.Node) (ast.Node, bool) {
	switch node := n.(type) {
	case *ast.Tuple:
		p.ProcessTuple(node)
	case *ast.Column:
		p.ProcessColumn(node)

	default:
		node.Accept(p.ParentVisitor)
	}
	return n, true
}

func (p *TupleVisitor) AddProjection(t *ast.Relation, node *ast.Column) {
	t.Projection = append(t.Projection, node)
}

func (p *TupleVisitor) RelationAddColumn(ctx *Context, node *ast.Column) {
	table := ctx.Table[node.Table]
	if table == nil {
		ctx.HasError = true
		log.Println("ERROR:Table '", table.Alias, "' doesn't exist")
	} else {
		if node.Column == "*" {
			table.HasStar = true
		} else {
			p.AddProjection(table, node)
		}
	}
}

/*************************************************
*
* 	ConditionVisitor
*	功能:负责遍历WhereClause
*  	范围:对于所有的节点,均处理
*
**************************************************/

type ConditionVisitor struct {
	ParentVisitor *PreProcessVisitor
}

func NewConditionVisitor(p *PreProcessVisitor) *ConditionVisitor {
	return &ConditionVisitor{
		ParentVisitor: p,
	}
}

func (p *ConditionVisitor) Notify(n ast.Node) bool {
	return true
}

func (p *ConditionVisitor) Visit(n ast.Node) (ast.Node, bool) {
	switch node := n.(type) {
	case *ast.Column:
		p.ProcessColumn(node)
	default:
		node.Accept(p.ParentVisitor)
	}
	return n, true
}

//按照深度优先遍历ExprNode,获得location和Expr的映射关系
func (p *ConditionVisitor) Dfs(n ast.ExprNode) []ast.ExprNode {
	switch node := n.(type) {
	case *ast.Expr:
		switch node.GetTag() {
		case ast.Expr_And, ast.Expr_Or:
			l := p.Dfs(node.Left)
			r := p.Dfs(node.Right)
			for i, _ := range r {
				l = append(l, r[i])
			}
			return l
		default:
			return []ast.ExprNode{node}
		}
	default:
		return []ast.ExprNode{node}
	}
}

//按照深度优先遍历ExprNode,获得location和Expr的映射关系
func (p *ConditionVisitor) DfsAnd0Or(n ast.ExprNode) []ast.ExprNode {
	switch node := n.(type) {
	case *ast.Expr:
		switch node.GetTag() {
		case p.ParentVisitor.CurrentContext.And0Or:
			l := p.DfsAnd0Or(node.Left)
			r := p.DfsAnd0Or(node.Right)
			for i, _ := range r {
				l = append(l, r[i])
			}
			return l
		default:
			return []ast.ExprNode{node}
		}
	default:
		return []ast.ExprNode{node}
	}
}

//检查是否存在常数形式的表达式
func (p *ConditionVisitor) CheckConstantExpr(list []ast.ExprNode) bool{

	return false
}

