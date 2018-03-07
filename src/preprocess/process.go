package preprocess

import "ast"
import "log"

/**********************************************************
*	From Item
*
* 	涉及的Relation表的个数 ProcessFromClause
* 	添加Meta元信息 ProcessRelation
* 	设置Relation与Alias的映射关系 ProcessRelation
*
************************************************************/

func (v *PreProcessVisitor) ProcessFromClause(node *ast.FromClause) {
	ctx := v.CurrentContext
	visitor := NewRelationVisitor(v)

	for index, _ := range node.Table_ref {
		node.Table_ref[index].Accept(visitor)
	}

	//设置ctx的Size域
	ctx.Size = visitor.Size
}

//获得所有的Relation信息
func (v *RelationVisitor) ProcessRelation(node *ast.Relation) {
	ctx := v.ParentVisitor.CurrentContext
	v.Size++
	switch node.GetTag() {
	case ast.Relation_Table:
		if node.Alias != "" {
			ctx.Table[node.Alias] = node
		} else {
			//设置别名
			node.Alias = node.Table
			ctx.Table[node.Table] = node
		}

	case ast.Relation_Subquery:
		ctx.Table[node.Alias] = node
	}
}

/**********************************************************
*	Target Item
*
* 	设置每个Tuple的位置 ProcessTargetClause
* 	设置位置和Tuple的映射关系 ProcessTargetClause
* 	添加Alias和位置的映射关系 ProcessTuple
* 	设置column和Relation的映射关系 ProcessColumn
*
* 	设置func的列表
*   设置Agg的列表
*
************************************************************/
func (v *PreProcessVisitor) ProcessTargetClause(node *ast.TargetClause) {
	ctx := v.CurrentContext
	visitor := NewTupleVisitor(v)

	for index, _ := range node.Target_ref {
		target := node.Target_ref[index].(*ast.Tuple)
		target.Location = index

		//设置ctx的Field域
		ctx.Field = append(ctx.Field, target)

		//设置ctx的FieldIndex
		ctx.FieldIndex[target.GetAlias()] = target.Location

		//遍历Tuple
		target.Accept(visitor)

	}
}

func (v *TupleVisitor) ProcessTuple(node *ast.Tuple) {
	ctx := v.ParentVisitor.CurrentContext

	//设置ctx的Agg和Func域
	switch node.Ref.(type) {
	case *ast.AggregationFuncCall:
		ctx.HasAgg = true
		ctx.Agg = append(ctx.Agg, node.Location)
	case *ast.FuncCall:
		ctx.Func = append(ctx.Func, node.Location)
	case *ast.CastFuncCall:
		ctx.Func = append(ctx.Func, node.Location)
	case *ast.CalTimeFuncCall:
		ctx.Func = append(ctx.Func, node.Location)
	case *ast.StringFuncCall:
		ctx.Func = append(ctx.Func, node.Location)
	}

}

//将所有的column添加到对应的ReLation中
func (v *TupleVisitor) ProcessColumn(node *ast.Column) {
	ctx := v.ParentVisitor.CurrentContext

	if node.IsStar {
		// '*'的情况
		if node.Table == "" {
			for key, _ := range ctx.Table {
				ctx.Table[key].HasStar = true
			}
		} else {
			v.RelationAddColumn(ctx, node)
		}
	} else {
		//非'*'的情况
		if node.Table == "" && ctx.Size > 1 {
			ctx.HasError = true
			log.Println("ERROR:Column '", node.Column, "' in field list is ambiguous")
		} else if node.Table == "" && ctx.Size == 1 {
			for key, _ := range ctx.Table {
				v.AddProjection(ctx.Table[key], node)
			}
		} else {
			v.RelationAddColumn(ctx, node)
		}
	}
}

/**********************************************************
*	Where Item
*	获得当前条件的or或者and形式
*   按照深度优先遍历,获得 从左到右的expr顺序列表
*
*
************************************************************/

func (v *PreProcessVisitor) ProcessWhereClause(node *ast.WhereClause) {
	ctx := v.CurrentContext
	visitor := NewConditionVisitor(v)

	//设置And0Or域
	ctx.And0Or=node.Where.(*ast.Expr).GetTag() 

	//设置Expr域和对应的index
	ctx.Expr = visitor.Dfs(node.Where)
	for idx, _ := range ctx.Expr {
		ctx.ExprIndex[ctx.Expr[idx]] = idx
	}

	//设置And0OrExpr域
	ctx.And0OrExpr = visitor.DfsAnd0Or(node.Where)

	//设置Isconstant域
	ctx.IsConstant=visitor.CheckConstantExpr(ctx.And0OrExpr)

}

func (v *ConditionVisitor) ProcessColumn(node *ast.Column) {

}
