package plan

import "ast"
import meta "util/meta"
import "fmt"
import "strings"

//根据每一个relation条件生成
func (pre *IndexVisitor) ExtractIndexFromWhereClause(alias string, expr ast.ExprNode) {
	//遍历where条件
	switch e := expr.(type) {
	case *ast.And0OrExpr:
		//遍历and条件
		if e.GetTag() == ast.Expr_And {
			pre.FindIndexInAndExpr(alias, e)
			pre.GetIndex(alias)
			//遍历Or条件
		} else {
			pre.FindIndexInOrExpr(alias, e)
		}
		//其他单一条件
	default:
		pre.FindIndexInOneExpr(alias, expr)
		pre.GetIndex(alias)
	}

}

//根据每一个relation条件生成
func (pre *IndexVisitor) ExtractIndexFromOrderBy(alias string) {
	ctx := pre.current
	if order, _, ok := ctx.CheckOrderByIndex(alias); ok {
		orderIndex, _ := pre.getOrderbyIndex(nil, alias, order, 0)
		if orderIndex != nil {
			indexQuery := generateIndex(alias, order)
			if _, ok := pre.indexs["Order索引"]; ok {
				pre.indexs["Order索引"] = append(pre.indexs["Order索引"], indexQuery)
			} else {
				pre.indexs["Order索引"] = []string{indexQuery}
			}
		}
	}

}

func (pre *IndexVisitor) FindIndexInAndExpr(alias string, e *ast.And0OrExpr) {
	conditionStack := newCondition()
	pre.current.IndexMap[alias] = conditionStack
	for _, expr := range e.Expr {
		switch node := expr.(type) {
		case *ast.Expr:
			if column, op, ok := pre.checkColumnAndOperator(alias, node); ok {
				switch op {
				//hash index条件之一
				case "=":
					//新索引入栈
					conditionStack.add(column, a_equal, node)
				case "<=>", "<>":
					conditionStack.add(column, a_not_equal, node)
				//betree索引之一
				default:
					conditionStack.add(column, a_range, node)
				}
			}
		//hash索引条件之一
		case *ast.IsNullExpr:
			//a is null的形式
			if !node.HasNot && node.Left.Type() == ast.Ast_column {
				//新索引入栈
				column := node.Left.(*ast.Column)
				if column.GetTable() == alias {
					conditionStack.add(column.GetColumn(), a_equal, node)
				}
			}
		//Btree索引条件之一
		case *ast.BetweenExpr:
			// a between 1 and 2添加入索引当中
			if node.Expr.Type() == ast.Ast_column && isConstant(node.Left) && isConstant(node.Right) {
				pre.CheckImplicitConversion(node)
				if !node.HasNot {
					column := node.Expr.(*ast.Column)
					if column.GetTable() == alias {
						conditionStack.add(column.GetColumn(), a_range, node)
					}
				}
			}
		//Btree索引条件之一
		case *ast.LikeExpr:
			//a like 'A%'
			if !node.HasNot && node.Left.Type() == ast.Ast_column && node.Right.Type() == ast.Ast_string && node.Right.GetTag() != ast.String_math {
				expr := node.Right.(*ast.String)
				if !strings.HasPrefix(expr.Str, "%") {
					column := node.Left.(*ast.Column)
					if column.GetTable() == alias {
						conditionStack.add(column.GetColumn(), a_range, node)
					}
				}
			}
		case *ast.InExpr:
			if isConstantList(node.Right) {
				pre.CheckImplicitConversion(node)
				if !node.HasNot {
					column := node.Left.(*ast.Column)
					if column.GetTable() == alias {
						conditionStack.add(column.GetColumn(), a_range, node)
					}

					if len(node.Right) >= 100 {
						e := fmt.Sprintf("In条件: 包含 %d 个项,请减少In的个数", len(node.Right))
						pre.Error = append(pre.Error, e)
					}
				}
			}
		}
	}
}

func (pre *IndexVisitor) FindIndexInOrExpr(alias string, e *ast.And0OrExpr) {

}

func (pre *IndexVisitor) FindIndexInOneExpr(alias string, expr ast.ExprNode) {
	e := &ast.And0OrExpr{Expr: []ast.ExprNode{expr}}
	e.SetTag(ast.Expr_And)
	pre.FindIndexInAndExpr(alias, e)
}

func (pre *IndexVisitor) FindUsingIndex(node *ast.Using) {
	//undo
}

func (pre *IndexVisitor) FindEuqalConditionIndex(n ast.ExprNode) {
	switch node := n.(type) {
	//多条件
	case *ast.And0OrExpr:
		pre.FindEqualConditionInAnd0OrExpr(node)
	//单一条件
	default:
		pre.FindEqualCondition(n)
	}
}

func (pre *IndexVisitor) FindEqualConditionInAnd0OrExpr(node *ast.And0OrExpr) {
	ctx := pre.current

	indexMap := make(map[string][]string)

	if node.GetTag() == ast.Expr_And {
		//筛选索引
		state := 0
		for _, expr := range node.Expr {
			if col1, col2, ok := isEqualCondition(expr); ok {
				col1Table := col1.Relation
				alias1 := col1Table.GetAlias()
				if col1Table.GetTag() == ast.Relation_Table {
					state = 1
					indexMap[alias1] = append(indexMap[alias1], col1.GetColumn())
				}
				col2Table := col2.Relation
				alias2 := col2Table.GetAlias()
				if col2Table.GetTag() == ast.Relation_Table {
					state = 1
					indexMap[alias2] = append(indexMap[alias2], col2.GetColumn())
				}
			}
		}

		if state == 0 {
			pre.addInfo(fmt.Sprintln("	[Join条件索引校验]Where条件中的不存在连接条件,略过"))
		} else {
			pre.addInfo(fmt.Sprintln("	[Join条件索引校验]当前条件(", generateString(node), ")是AND形式,校验连接条件索引"))
		}

		for alias, columns := range indexMap {
			if len(columns) != 0 {
				currentIndex, currentCtx := ctx.GetIndexMeta(alias)
				indexs := [][]string{}
				for id := 1; id <= len(columns); id++ {
					indexs = append(indexs, combination(columns, id)...)
				}

				if primaryKey, ok := currentCtx.HasPrimary(indexs, alias); ok {
					pre.addInfo(fmt.Sprintln("	[Join条件索引校验]表", alias, "存在主键索引条件:", primaryKey, ",不需要添加索引"))
				} else {
					flag := 0
					for _, index := range indexs {
						if currentIndex.HaveIndex(index) {
							if len(index) == len(columns) {
								pre.addInfo(fmt.Sprintln("	[Join条件索引校验]表", alias, "存在索引条件:", index))
								flag = 1
							} else {
								pre.addInfo(fmt.Sprintln("	[Join条件索引校验]表", alias, "存在索引条件:", index, ",该索引缺少列"))
							}
						}
					}
					if flag == 0 {
						pre.addInfo(fmt.Sprintln("	[Join条件索引校验]表", alias, "不存在索引条件:", columns, ",需要添加索引"))
						table := currentCtx.Table[alias].GetTable()
						indexQuery := generateIndex(table, columns)
						if _, ok := pre.indexs["Join条件索引"]; ok {
							pre.indexs["Join条件索引"] = append(pre.indexs["Join条件索引"], indexQuery)
						} else {
							pre.indexs["Join条件索引"] = []string{indexQuery}
						}
					}
				}

			}
		}

	} else {
		pre.addInfo(fmt.Sprintln("	[Join条件索引校验]当前条件(", generateString(node), ")是OR形式,请调整写法"))
	}
}

func (pre *IndexVisitor) FindEqualCondition(n ast.ExprNode) {
	expr := &ast.And0OrExpr{Expr: []ast.ExprNode{n}}
	expr.SetTag(ast.Expr_And)
	pre.FindEqualConditionInAnd0OrExpr(expr)

}

//如果是a=1 或者 a<2形式的表达式,返回column 和 op
func (pre *IndexVisitor) checkColumnAndOperator(alias string, node *ast.Expr) (string, string, bool) {
	//经过预处理过程
	//所有的常数1<1 均被消除
	//1>a 均被转换为a<1
	//检测左侧节点的类型，不是column则直接返回
	if node.Left.Type() != ast.Ast_column {
		return "", "", false
	}

	//校验是否是当前alias下的column

	col := node.Left.(*ast.Column)

	if col.GetTable() != alias {
		return "", "", false
	}

	//检测右侧节点,是否是constat
	if !isConstant(node.Right) {
		return "", "", false
	}
	pre.CheckImplicitConversion(node)
	//返回column,操作符
	return col.GetColumn(), node.Operator, true
}

//检测饮食转换
func (pre *IndexVisitor) CheckImplicitConversion(n ast.ExprNode) {
	ctx := pre.current
	switch node := n.(type) {
	case *ast.Expr:
		col := node.Left.(*ast.Column)
		val := node.Right
		table := col.GetTable()
		column := col.GetColumn()
		if sqlType, ok := ctx.GetSqlType(table, column); ok {
			if !haveSameType(sqlType, node.Right) {
				e := fmt.Sprintf("条件%s发生隐式转换,列%s的类型为%s,条件的类型为%s", generateString(node), col.Column, meta.ReSqlDict[sqlType], ast.TypeString[val.Type()])
				pre.Error = append(pre.Error, e)
				return
			}
		}
	case *ast.BetweenExpr:
		col := node.Expr.(*ast.Column)
		val1 := node.Left
		val2 := node.Right
		table := col.GetTable()
		column := col.GetColumn()
		if sqlType, ok := ctx.GetSqlType(table, column); ok {
			if !haveSameType(sqlType, val1) {
				e := fmt.Sprintf("条件%s发生隐式转换,列%s的类型为%s,条件的类型为%s", generateString(node), column, meta.ReSqlDict[sqlType], ast.TypeString[val1.Type()])
				pre.Error = append(pre.Error, e)
				return
			}

			if !haveSameType(sqlType, val2) {
				e := fmt.Sprintf("条件%s发生隐式转换,列%s的类型为%s,条件的类型为%s", generateString(node), column, meta.ReSqlDict[sqlType], ast.TypeString[val2.Type()])
				pre.Error = append(pre.Error, e)
				return
			}
		}
	case *ast.InExpr:
		col := node.Left.(*ast.Column)
		table := col.GetTable()
		column := col.GetColumn()
		if sqlType, ok := ctx.GetSqlType(table, column); ok {
			for _, val := range node.Right {
				if !haveSameType(sqlType, val) {
					e := fmt.Sprintf("In条件发生隐式转换,列%s的类型为%s,条件的类型为%s", column, meta.ReSqlDict[sqlType], ast.TypeString[val.Type()])
					pre.Error = append(pre.Error, e)
					return
				}
			}
		}
	}
}
