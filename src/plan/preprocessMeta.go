package plan

import "ast"
import "fmt"
import server "util/server"
import meta "util/meta"

type PreProcessMetaVisitor struct {

	//记录所有的错误信息
	HasError bool
	Error    []string

	Ctx map[*ast.SimpleSelectStmt]*Context

	current *Context

	MetaServer *server.Mysql
}

func NewPreProcessMetaVisitor(server *server.Mysql) *PreProcessMetaVisitor {
	return &PreProcessMetaVisitor{
		HasError:   false,
		Error:      []string{},
		Ctx:        make(map[*ast.SimpleSelectStmt]*Context),
		current:    nil,
		MetaServer: server,
	}
}

func (pre *PreProcessMetaVisitor) Notify(n ast.Node) bool {
	switch node := n.(type) {
	case *ast.SimpleSelectStmt:
		//Enter SelectStmt and Create a new Context
		ctx := NewContext()
		ctx.Parent(pre.current)
		pre.Ctx[node] = ctx
		pre.current = ctx
	case *ast.Join:
		pre.ChangeRightJoin2LeftJoin(node)
	}
	return true
}

func (pre *PreProcessMetaVisitor) Visit(n ast.Node) (ast.Node, bool) {
	switch node := n.(type) {
	case *ast.SimpleSelectStmt:
		pre.current = pre.current.Pctx
	case *ast.Relation:
		pre.current.Size++
		pre.SetRelationAlias(node)
		pre.GetRelationMeta(node)
	case *ast.Tuple:
		switch node.GetTag() {
		case ast.Tuple_star:
			pre.current.Tuple = append(pre.current.Tuple, node)
			pre.AddStar2Projection(node.Ref.(*ast.Star))
		default:
			pre.SetTupleAlias(node)
		}
	}
	return n, true
}

func (pre *PreProcessMetaVisitor) Session() server.Server {
	return pre.MetaServer
}

//设置别名
func (pre *PreProcessMetaVisitor) ChangeRightJoin2LeftJoin(node *ast.Join) {
	if node.GetTag() == ast.Right_join {
		node.Left, node.Right = node.Right, node.Left
		node.SetTag(ast.Left_join)
		node.IsEnclosed = true
	}
}

//设置别名(均为小写)
func (pre *PreProcessMetaVisitor) SetRelationAlias(node *ast.Relation) {
	ctx := pre.current
	//移除别名的``
	alias := node.GetAlias()
	//判断是否有冲突的别名
	if _, ok := ctx.Table[alias]; ok {
		pre.HasError = true
		e := fmt.Sprintf("SetRelationAlias发生错误,已经存在别名为%s的relation,请检查", alias)
		pre.Error = append(pre.Error, e)
	} else {
		//alias移除了``并且为小写
		ctx.Table[alias] = node
	}
}

//这只projection,不区分大小写
func (pre *PreProcessMetaVisitor) AddStar2Projection(node *ast.Star) {
	ctx := pre.current
	//移除table的``
	table := node.GetTable()
	if table == "" {
		//将当前上下文中的所有Meta添加到Tuple中去
		for relation, meta := range ctx.Meta {
			ctx.Projection.AddMeta(meta)
			appendColumn(node, relation, meta.GetField())
		}
	} else {
		//检查是否属于当前上下文
		relation := ctx.Table[table]
		if relation != nil {
			//移除alias的``
			alias := relation.GetAlias()
			ctx.Projection.AddMeta(ctx.Meta[alias])
			appendColumn(node, alias, ctx.Meta[alias].GetField())
		} else {
			pre.HasError = true
			e := fmt.Sprintf("AddStar2Tuple发生错误,*所属的Relation不存在,请检查")
			pre.Error = append(pre.Error, e)
		}
	}
}

func (pre *PreProcessMetaVisitor) SetTupleAlias(node *ast.Tuple) {
	ctx := pre.current
	alias := node.GetAlias()
	//添加对应的投影信息,为子查询提供元数据信息
	ctx.Projection.Add(alias, "", "Yes")
	ctx.Tuple = append(ctx.Tuple, node)

	if _, ok := ctx.TupleMap[alias]; ok {
		ctx.TupleMap[alias] = append(ctx.TupleMap[alias], node)
	} else {
		ctx.TupleMap[alias] = []*ast.Tuple{node}
	}

}

//获得表的元数据信息
func (pre *PreProcessMetaVisitor) GetRelationMeta(node *ast.Relation) {
	ctx := pre.current
	alias := node.GetAlias()
	switch node.GetTag() {
	case ast.Relation_Table:
		db := node.GetDB()
		table := node.GetTable()
		ctx.Meta[alias] = pre.GetTableMeta(db, table)
		ctx.IndexMeta[alias] = pre.GetTableIndexMeta(db, table)
		ctx.Row[alias] = pre.GetTableRowMeta(db, table)
	case ast.Relation_Subquery:
		ctx.Meta[alias] = pre.GetSubqueryMeta(node.Select)
	}
}

//获得Table的元数据信息
func (pre *PreProcessMetaVisitor) GetTableMeta(DB string, table string) *meta.Meta {
	session := pre.MetaServer.CreateSession2DB(DB)
	meta, err := pre.MetaServer.QueryTable(session, table)
	if err != nil {
		pre.HasError = true
		pre.Error = append(pre.Error, err.Error())
	}
	return meta
}

//获得Table的元数据信息
func (pre *PreProcessMetaVisitor) GetTableIndexMeta(DB string, table string) *meta.Index {
	session := pre.MetaServer.CreateSession2DB(DB)
	meta, err := pre.MetaServer.QueryIndex(session, table)
	if err != nil {
		pre.HasError = true
		pre.Error = append(pre.Error, err.Error())
	}
	return meta
}

//获得Table的元数据信息
func (pre *PreProcessMetaVisitor) GetTableRowMeta(DB string, table string) int {
	session := pre.MetaServer.CreateSession2DB(DB)
	row, err := pre.MetaServer.QueryRow(session, table)
	if err != nil {
		pre.HasError = true
		pre.Error = append(pre.Error, err.Error())
	}
	return row
}

//获得Subquery元数据信息
func (pre *PreProcessMetaVisitor) GetSubqueryMeta(n ast.Node) *meta.Meta {
	//拿到Tuple别名对应信息
	switch stmt := n.(type) {
	case *ast.SimpleSelectStmt:
		return pre.Ctx[stmt].GetMetaFromSubquery()
	case *ast.UnionStmt:
		//mysql 选择左边的
		return pre.GetSubqueryMeta(stmt.Left)
	default:
		return nil
	}
}
