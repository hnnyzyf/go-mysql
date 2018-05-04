package plan

import "github.com/hnnyzyf/mysql-go/advisor/ast"
import "github.com/hnnyzyf/mysql-go/advisor/util/comp"

//对于a=1 形式,进行constantfold
func (pre *ConstantVisitor) CheckExpr(node *ast.Expr) bool {
	if col, val, ok := isEqualExpr(node); ok {
		//进行常数转换
		currentAnd0OrExpr := pre.top().(*ast.And0OrExpr)
		pre.ConstantFolding(col, val, currentAnd0OrExpr)

		return true
	} else {
		return false
	}
}

//对于a is null 形式,进行constantfold
func (pre *ConstantVisitor) CheckIsNullExpr(node *ast.IsNullExpr) bool {
	if !node.HasNot {
		currentAnd0OrExpr := pre.top().(*ast.And0OrExpr)
		if node.Left.Type() == ast.Ast_column {
			col := node.Left.(*ast.Column)
			val := &ast.Null{}
			pre.ConstantFolding(col, val, currentAnd0OrExpr)
			return true
		}
	}
	return false
}

//对于a is null 形式,进行constantfold
func (pre *ConstantVisitor) CheckInExpr(node *ast.InExpr, index int, expr []ast.ExprNode) bool {
	if e, col, val, ok := pre.ExtractInfoFromInExpr(node); ok {
		currentAnd0OrExpr := pre.top().(*ast.And0OrExpr)
		pre.ConstantFolding(col, val, currentAnd0OrExpr)
		expr[index] = e.(ast.ExprNode)
		return true
	}
	return false
}

func (pre *ConstantVisitor) ExtractInfoFromInExpr(node *ast.InExpr) (ast.Node, *ast.Column, ast.ExprNode, bool) {
	//判断是否是 a in (1)的形式
	ltp := node.Left.Type()
	if ltp != ast.Ast_column {
		return node, nil, nil, false
	}

	if len(node.Right) != 1 {
		return node, nil, nil, false
	}

	expr := &ast.Expr{Operator: "=", Left: node.Left, Right: node.Right[0]}
	return expr, node.Left.(*ast.Column), node.Right[0], true
}

//对于a is null 形式,进行constantfold
func (pre *ConstantVisitor) CheckBetweenExpr(n *ast.BetweenExpr, index int, expr []ast.ExprNode) bool {
	//a between -1 and 2
	if !isConstant(n.Expr) && isConstant(n.Left) && isConstant(n.Right) {
		if comp.Equal(generateQuery(n.Left, n.Right, "="), pre.Session()) {
			if n.HasNot {
				expr[index] = &ast.Expr{Operator: "<>", Left: n.Expr, Right: n.Left}
			} else {
				expr[index] = &ast.Expr{Operator: "=", Left: n.Expr, Right: n.Left}
				if n.Expr.Type() == ast.Ast_column {
					//进行常数折叠
					col := n.Expr.(*ast.Column)
					val := n.Left
					currentAnd0OrExpr := pre.top().(*ast.And0OrExpr)
					pre.ConstantFolding(col, val, currentAnd0OrExpr)
					return true
				}
			}
		}
	}
	return false
}

//对于a like  形式,进行constantfold
//a (not) like 1 true
//a (not) like '1' true
////a (not) like '1%' false
func (pre *ConstantVisitor) CheckLikeExpr(node *ast.LikeExpr, index int, expr []ast.ExprNode) bool {
	if node.Right.Type() == ast.Ast_column && isConstant(node.Left) {
		node.Left, node.Right = node.Right, node.Left
	}

	if node.Left.Type() == ast.Ast_column && isConstant(node.Right) {
		if node.Right.Type() == ast.Ast_string {
			if str := node.Right.(*ast.String).Str; node.Right.GetTag() != ast.String_math {
				if withPercent(str) {
					return false
				}
			}
		}
		switch node.HasNot {
		case true:
			expr[index] = &ast.Expr{Operator: "<>", Left: node.Left, Right: node.Right}
			return false
		default:
			expr[index] = &ast.Expr{Operator: "=", Left: node.Left, Right: node.Right}
			//进行常数折叠
			col := node.Left.(*ast.Column)
			val := node.Right
			currentAnd0OrExpr := pre.top().(*ast.And0OrExpr)
			pre.ConstantFolding(col, val, currentAnd0OrExpr)
			return true
		}
	}

	return false
}
