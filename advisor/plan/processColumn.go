package plan

import "github.com/hnnyzyf/mysql-go/advisor/ast"
import "fmt"
import "github.com/hnnyzyf/mysql-go/advisor/util/server"

/*******************************

	1.将*替换成StarExpr
	2.生成location列表

********************************/

type TargetClauseVisitor struct {
	//记录所有的错误信息
	HasError bool
	Error    []string

	Ctx map[*ast.SimpleSelectStmt]*Context

	current *Context
}

func NewTargetClauseVisitor(ctx map[*ast.SimpleSelectStmt]*Context) *TargetClauseVisitor {
	return &TargetClauseVisitor{
		HasError: false,
		Error:    []string{},
		Ctx:      ctx,
		current:  nil,
	}
}

func (pre *TargetClauseVisitor) Notify(n ast.Node) bool {
	switch node := n.(type) {
	case *ast.SimpleSelectStmt:
		//set current context
		pre.current = pre.Ctx[node]
		return true
	//如下Clause不进行处理
	case *ast.FromClause, *ast.WhereClause:
		return false
	default:
		return true
	}
}

func (pre *TargetClauseVisitor) Visit(n ast.Node) (ast.Node, bool) {
	switch node := n.(type) {
	case *ast.SimpleSelectStmt:
		pre.current = pre.current.Pctx
	case *ast.Column:
		pre.SetColumn2Relation(node)
	case *ast.SortClause:
		pre.SetSortbyMap(node)

	}

	return n, true
}

func (pre *TargetClauseVisitor) Session() server.Server {
	return nil
}

func (pre *TargetClauseVisitor) SetColumn2Relation(node *ast.Column) {
	ctx := pre.current
	table := node.GetTable()
	column := node.GetColumn()
	switch node.GetTable() {
	case "":
		res := pre.GetColumn2Relation(column)
		switch len(res) {
		case 0:
			pre.HasError = true
			e := fmt.Sprintf("SetColumn2Relation发生错误,Column:%s不属于任何关系,请检查", column)
			pre.Error = append(pre.Error, e)
		case 1:
			node.Table = res[0]
		default:
			//检查是否在using中
			fmt.Print(column)
			fmt.Print(ctx.Using)
			if !ctx.HasUsing(column) {
				pre.HasError = true
				e := fmt.Sprintf("SetColumn2Relation发生错误,Column:%s属于多个关系,请检查", column)
				pre.Error = append(pre.Error, e)
			}
		}
	default:
		if !ctx.HasColumn(table, column) {
			pre.HasError = true
			e := fmt.Sprintf("SetColumn2Relation发生错误,Column:%s不属于关系:%s,请检查", table, column)
			pre.Error = append(pre.Error, e)
		}
	}
}

func (pre *TargetClauseVisitor) SetSortbyMap(node *ast.SortClause) {
	ctx := pre.current
	ctx.Sortby = []*ast.Sortby{}
	//检测可以使用的Group Item
	for _, expr := range node.Target_ref {
		ctx.Sortby = append(ctx.Sortby, expr.(*ast.Sortby))
	}
}

func (pre *TargetClauseVisitor) GetColumn2Relation(column string) []string {
	ctx := pre.current
	res := []string{}
	relations := ctx.GetAllRelationAlias()
	for _, table := range relations {
		if ctx.HasColumn(table, column) {
			res = append(res, table)
		}
	}

	return res
}

type FromClauseVisitor struct {
	//记录所有的错误信息
	HasError bool
	Error    []string

	Ctx map[*ast.SimpleSelectStmt]*Context

	current *Context
}

func NewFromClauseVisitor(ctx map[*ast.SimpleSelectStmt]*Context) *FromClauseVisitor {
	return &FromClauseVisitor{
		HasError: false,
		Error:    []string{},
		Ctx:      ctx,
		current:  nil,
	}
}

func (pre *FromClauseVisitor) Notify(n ast.Node) bool {
	switch node := n.(type) {
	case *ast.SimpleSelectStmt:
		//set current context
		pre.current = pre.Ctx[node]
		return true
	case *ast.Join:
		pre.SetUsingColumn(node)
		pre.current.push(node)
		return true
	case *ast.Relation:
		return false
	case *ast.Using:
		return false
	default:
		return true
	}
}

func (pre *FromClauseVisitor) Visit(n ast.Node) (ast.Node, bool) {
	switch node := n.(type) {
	case *ast.SimpleSelectStmt:
		pre.current = pre.current.Pctx
		//只处理FromClause
	case *ast.FromClause:
		return n, false
	case *ast.Join:
		pre.current.pop()
	case *ast.Column:
		pre.SetColumn2Relation(node)

	}
	return n, true
}

func (pre *FromClauseVisitor) Session() server.Server {
	return nil
}

//记录using的column
func (pre *FromClauseVisitor) SetUsingColumn(node *ast.Join) {
	ctx := pre.current
	if node.Condition != nil {
		if node.Condition.Type() == ast.Ast_using {
			using := node.Condition.(*ast.Using)
			for _, item := range using.Column {
				col := item.(*ast.Column)
				ctx.Using = append(ctx.Using, col.GetColumn())
			}
		}
	}
}

func (pre *FromClauseVisitor) SetColumn2Relation(node *ast.Column) {
	switch node.GetTable() {
	case "":
		//判断是否属于当前join的子节点
		pre.CheckCurrentJoin(node)
	default:
		if pre.CheckColumn2Relation(node.GetTable(), node.GetColumn()) {
			node.Relation = pre.current.Table[node.GetTable()]
		}
	}
}

func (pre *FromClauseVisitor) CheckCurrentJoin(node *ast.Column) {
	ctx := pre.current
	join := ctx.top().(*ast.Join)
	column := node.GetColumn()
	res := pre.DfscurrentJoin(join, column)
	switch len(res) {
	case 0:
		pre.HasError = true
		e := fmt.Sprintf("CheckCurrentJoin发生错误,%s不属于任何relation,请检查", column)
		pre.Error = append(pre.Error, e)
		return
	case 1:
		node.Table = res[0]
		node.Relation = ctx.GetRelation(node.GetTable())
	default:
		pre.HasError = true
		e := fmt.Sprintf("CheckCurrentJoin发生错误,%s属于多个relation,列表为:%s,请检查", column, res)
		pre.Error = append(pre.Error, e)
	}

}

//Dfs当前join下的所有relation
func (pre *FromClauseVisitor) DfscurrentJoin(node *ast.Join, column string) []string {
	ctx := pre.current
	relations := []string{}
	switch n := node.Left.(type) {
	case *ast.Join:
		relations = append(relations, pre.DfscurrentJoin(n, column)...)
	case *ast.Relation:
		alias := n.GetAlias()
		if ctx.HasColumn(alias, column) {
			relations = append(relations, alias)

		}
	}

	switch n := node.Right.(type) {
	case *ast.Join:
		relations = append(relations, pre.DfscurrentJoin(n, column)...)
	case *ast.Relation:
		alias := n.GetAlias()
		if ctx.HasColumn(alias, column) {
			relations = append(relations, alias)

		}
	}

	return relations
}

func (pre *FromClauseVisitor) CheckColumn2Relation(table string, column string) bool {
	ctx := pre.current
	if !ctx.HasColumn(table, column) {
		pre.HasError = true
		e := fmt.Sprintf("CheckColumn2Relation发生错误,%s不属于relation,请检查", column)
		pre.Error = append(pre.Error, e)
		return false
	} else {

		return true
	}
}

type WhereClauseVisitor struct {
	//记录所有的错误信息
	HasError bool
	Error    []string

	Ctx map[*ast.SimpleSelectStmt]*Context

	current *Context
}

func NewWhereClauseVisitor(ctx map[*ast.SimpleSelectStmt]*Context) *WhereClauseVisitor {
	return &WhereClauseVisitor{
		HasError: false,
		Error:    []string{},
		Ctx:      ctx,
		current:  nil,
	}
}

func (pre *WhereClauseVisitor) Notify(n ast.Node) bool {
	switch node := n.(type) {
	case *ast.SimpleSelectStmt:
		//set current context
		pre.current = pre.Ctx[node]
		return true
	//不处理FromClause
	case *ast.TargetClause, *ast.FromClause:
		return false
	default:
		return true
	}
}

func (pre *WhereClauseVisitor) Visit(n ast.Node) (ast.Node, bool) {
	switch node := n.(type) {
	case *ast.SimpleSelectStmt:
		pre.current = pre.current.Pctx
	case *ast.WhereClause:
		return node, false
	case *ast.Column:
		pre.SetColumn2Relation(node)
	}
	return n, true
}

func (pre *WhereClauseVisitor) Session() server.Server {
	return nil
}

func (pre *WhereClauseVisitor) SetColumn2Relation(node *ast.Column) {
	table := node.GetTable()
	column := node.GetColumn()
	switch table {
	case "":
		res, context := pre.GetColumn2Relation(column)
		switch len(res) {
		case 0:
			pre.HasError = true
			e := fmt.Sprintf("SetColumn2Relation发生错误,Column:%s不属于任何关系,请检查", column)
			pre.Error = append(pre.Error, e)
		case 1:
			node.Table = res[0]
			if context {
				//当前上下文
				ctx := pre.current
				node.Relation = ctx.GetRelation(table)
			} else {
				ctx := pre.current.Pctx
				node.Relation = ctx.GetRelation(table)
			}
		default:
			pre.HasError = true
			e := fmt.Sprintf("SetColumn2Relation发生错误,Column:%s属于多个关系,请检查", column)
			pre.Error = append(pre.Error, e)
		}
	default:
		ctx := pre.current
		if ctx.HasRelation(table) {
			if !ctx.HasColumn(table, column) {
				pre.HasError = true
				e := fmt.Sprintf("SetColumn2Relation发生错误,Column:%s不属于关系:%s,请检查", column, table)
				pre.Error = append(pre.Error, e)
			}
			node.Relation = ctx.Table[table]
		} else {
			if pctx := ctx.Pctx; pctx != nil {
				if pctx.HasRelation(table) {
					if !pctx.HasColumn(table, column) {
						pre.HasError = true
						e := fmt.Sprintf("SetColumn2Relation发生错误,Column:%s不属于关系%s,请检查", column, table)
						pre.Error = append(pre.Error, e)
					}
					node.Relation = pctx.Table[table]
				} else {
					pre.HasError = true
					e := fmt.Sprintf("SetColumn2Relation发生错误,关系:%s不存在于任何当前上下文和父上下文,请检查%s的归属", table, column)
					pre.Error = append(pre.Error, e)
				}
			}
		}
	}
}

//第一个返回值是所属表的个数
//第二个返回值表明属于currentctx还是parentcontext
func (pre *WhereClauseVisitor) GetColumn2Relation(column string) ([]string, bool) {
	//在current上下文检查
	if res, ok := pre.current.CheckBelongings(column); ok {
		return res, true
	} else {
		//在parent的上下文检查
		if ctx := pre.current.Pctx; ctx != nil {
			res, _ = pre.current.Pctx.CheckBelongings(column)
		}
		return res, false
	}
}
