package plan

import "ast"
import comp "util/compare"

//将1=1 1="a" 2="23a" 1<2替换掉,忽略掉function
func (pre *ConstantVisitor) RemoveExpr(node *ast.Expr) ast.Node {
	//非常数直接返回
	if isConstant(node.Left) && isConstant(node.Right) {
		if isMathOp(node.Operator) {
			res := &ast.String{Str: generateString(node)}
			res.SetTag(ast.String_math)
			pre.hasConstantExprCombine = true
			return res
		} else {
			return pre.checkTure0false(node, 1)
		}
	} else {
		return node
	}

}

func (pre *ConstantVisitor) RemoveInExpr(node *ast.InExpr) ast.Node {
	if isConstant(node.Left) && isConstantList(node.Right) {
		return pre.checkTure0false(node, 1)
	} else {
		return node
	}
}

func (pre *ConstantVisitor) RemoveUnaryExpr(node *ast.UnaryExpr) ast.Node {
	//查看Unary的对象
	switch n := node.Expr.(type) {
	case *ast.Integer:
		n.Operator = append(n.Operator, node.Operator)
		node.SetTag(ast.Unary_Constant)
		return n
	case *ast.Float:
		n.Operator = append(n.Operator, node.Operator)
		node.SetTag(ast.Unary_Constant)
		return n
	case *ast.String:
		n.Operator = append(n.Operator, node.Operator)
		node.SetTag(ast.Unary_Constant)
		return n
	case *ast.UnaryExpr:
		node.SetTag(n.GetTag())
		return node
	default:
		node.SetTag(ast.Unary_Other)
		return node
	}

}

func (pre *ConstantVisitor) RemoveBetweenExpr(node *ast.BetweenExpr) ast.Node {

	//判断所有条件是否是常数条件
	//1 between 1 and 2
	//1 not between 2 and 2
	if isConstant(node.Expr) && isConstant(node.Left) && isConstant(node.Right) {
		return pre.checkTure0false(node, 1)
	}

	if !isConstant(node.Expr) && isConstant(node.Left) && isConstant(node.Right) {
		//不存在n.left=n.right的情况，已经被折叠了
		// a between 2 and 1
		// a not between 2 and 1
		if comp.Max(generateQuery(node.Left, node.Right, ">"), pre.Session()) {
			return &ast.True{IsTrue: node.HasNot}
		}

	}

	//a between a and a的形式
	//a not between a and a 的形式
	if node.Expr.IsSameAs(node.Left, pre) && node.Expr.IsSameAs(node.Right, pre) {
		return &ast.True{IsTrue: !node.HasNot}
	}

	// a between a and ? 的形式
	if node.Expr.IsSameAs(node.Left, pre) && !node.Expr.IsSameAs(node.Right, pre) {
		if node.HasNot {
			return &ast.Expr{Operator: ">", Left: node.Expr, Right: node.Right}
		} else {
			return &ast.Expr{Operator: "<=", Left: node.Expr, Right: node.Right}
		}
	}

	// a between ? and a 的形式
	if !node.Expr.IsSameAs(node.Left, pre) && node.Expr.IsSameAs(node.Right, pre) {
		if node.HasNot {
			return &ast.Expr{Operator: "<", Left: node.Expr, Right: node.Left}
		} else {
			return &ast.Expr{Operator: ">=", Left: node.Expr, Right: node.Left}
		}
	}

	return node
}

func (pre *ConstantVisitor) RemoveIsNullExpr(node *ast.IsNullExpr) ast.Node {
	//是Null的情况
	if node.Left.Type() == ast.Ast_null {
		if node.HasNot {
			return &ast.True{IsTrue: false}
		} else {
			return &ast.True{IsTrue: true}
		}
	}

	//非null的情况
	if isConstant(node.Left) {
		return pre.checkTure0false(node, 1)
	} else {
		return node
	}

}

func (pre *ConstantVisitor) RemoveLikeExpr(node *ast.LikeExpr) ast.Node {
	//判断所有条件是否是常数条件
	//1 like 1 的情况
	if isConstant(node.Left) && isConstant(node.Right) {
		return pre.checkTure0false(node, 1)
	} else {
		return node
	}

	//a like '1%'

}

func (pre *ConstantVisitor) checkTure0false(n ast.ExprNode, tag int) ast.Node {
	session := pre.MetaServer.CreateSession2DB("")
	if ok, _ := pre.MetaServer.QueryTrue0False(session, generateString(n), tag); ok {
		return &ast.True{IsTrue: true}
	} else {
		return &ast.True{IsTrue: false}
	}
}

//将所有的值表达式转换为true 或者 false
func (pre *ConstantVisitor) ReplaceValAndColumn(node *ast.And0OrExpr) {
	session := pre.MetaServer.CreateSession2DB("")
	for index, expr := range node.Expr {
		switch expr.(type) {
		case *ast.Integer, *ast.Float, *ast.String:
			if ok, _ := pre.MetaServer.QueryTrue0False(session, generateString(expr), 0); ok {
				node.Expr[index] = &ast.True{IsTrue: true}
			} else {
				node.Expr[index] = &ast.True{IsTrue: false}
			}
		}
	}
}

func (pre *ConstantVisitor) ReplaceLogicFalseExpr(node *ast.And0OrExpr) bool {
	if node.GetTag() == ast.Expr_And {
		return pre.ReplaceLogicFalseInAndExpr(node)
	}
	return true
}

//移除a>1 and a<-1 形式的语句
func (pre *ConstantVisitor) ReplaceLogicFalseInAndExpr(node *ast.And0OrExpr) bool {
	state := true
	RangerMap := make(map[ast.Column]*ranger)
	for index, expr := range node.Expr {
		if expr.Type() == ast.Ast_expr {
			//是compare节点的情况下
			if col, val, op, ok := isCompareExpr(expr.(*ast.Expr)); ok {
				//true代表idx的节点要被替换掉
				if _, ok := RangerMap[*col]; !ok {
					RangerMap[*col] = newRanger()
				}
				//log.Println(col, val, op)
				if idx, ok := RangerMap[*col].add(index, val, op, pre.MetaServer); ok {
					node.Expr[idx] = &ast.True{IsTrue: true}
				}
			}
		}
		//已经留下a between 1 and 2的形式
		if expr.Type() == ast.Ast_betweenexpr {
			//undo
		}
	}

	//移除a>1 and a<-1的情况
	for _, ranger := range RangerMap {
		if hasNewEqualExpr := ranger.compare(node.Expr, pre.MetaServer); hasNewEqualExpr {
			//进行一次constantfolding
			expr := node.Expr[ranger.high_index].(*ast.Expr)
			col := expr.Left.(*ast.Column)
			val := expr.Right

			//进行一次constantfolding
			currentAnd0OrExpr := pre.top().(*ast.And0OrExpr)
			pre.ConstantFolding(col, val, currentAnd0OrExpr)
			state = false
		}
	}

	return state

}
