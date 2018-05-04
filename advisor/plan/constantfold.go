package plan

import "github.com/hnnyzyf/mysql-go/advisor/ast"
import "github.com/hnnyzyf/mysql-go/advisor/util/server"

type ConstantVisitor struct {
	//And0OrExpr的Stack
	stack

	MetaServer server.Server

	ConstantFoldIndexMap map[*ast.And0OrExpr][]bool

	hasConstantExprCombine bool

	HasError bool
	Error    []string

	Ctx map[*ast.SimpleSelectStmt]*Context

	current *Context
}

func NewConstantVisitor(s server.Server, ctx map[*ast.SimpleSelectStmt]*Context) *ConstantVisitor {
	return &ConstantVisitor{
		MetaServer:           s,
		ConstantFoldIndexMap: make(map[*ast.And0OrExpr][]bool),
		Ctx:                  ctx,
		current:              nil,
		Error:                []string{},
	}
}

func (pre *ConstantVisitor) Notify(n ast.Node) bool {
	switch node := n.(type) {
	case *ast.SimpleSelectStmt:
		//set current context
		pre.current = pre.Ctx[node]
		return true
	case *ast.And0OrExpr:
		pre.push(node)

		if _, ok := pre.ConstantFoldIndexMap[node]; !ok {
			pre.ConstantFoldIndexMap[node] = initialize(len(node.Expr))
		}

		if node.GetTag() == ast.Expr_And {
			pre.CheckAnd0OrExpr(node)
		}

		return true
	default:
		return true
	}
}

func (pre *ConstantVisitor) Visit(n ast.Node) (ast.Node, bool) {
	switch node := n.(type) {
	case *ast.SimpleSelectStmt:
		pre.current = pre.current.Pctx
		return n, true
	case *ast.And0OrExpr:
		defer pre.pop()
		return pre.RemoveAnd0OrExpr(node), true
	//对当前And0OrExpr下的节点进行的处理
	case *ast.Expr:
		return pre.RemoveExpr(node), true
	case *ast.InExpr:
		return pre.RemoveInExpr(node), true
	case *ast.UnaryExpr:
		return pre.RemoveUnaryExpr(node), true
	case *ast.BetweenExpr:
		return pre.RemoveBetweenExpr(node), true
	case *ast.IsNullExpr:
		return pre.RemoveIsNullExpr(node), true
	case *ast.LikeExpr:
		return pre.RemoveLikeExpr(node), true
	default:
		return n, true

	}
}

func (pre *ConstantVisitor) Session() server.Server {
	return pre.MetaServer
}

//检查And0Orexpr
func (pre *ConstantVisitor) CheckAnd0OrExpr(node *ast.And0OrExpr) {

	//记录已经进行过constantFolding的index,初始化为false
	HasConstantFoldIndexMap := pre.ConstantFoldIndexMap[node]

	for i := 0; i < len(node.Expr); i++ {
		HasConstantFoldIndexMap = append(HasConstantFoldIndexMap, false)
	}

	for currentIndex := 0; currentIndex < len(node.Expr); currentIndex++ {
		if !HasConstantFoldIndexMap[currentIndex] {
			switch n := node.Expr[currentIndex].(type) {
			//处理Expr
			case *ast.Expr:
				if pre.CheckExpr(n) {
					HasConstantFoldIndexMap[currentIndex] = true
					currentIndex = -1
				}
			case *ast.IsNullExpr:
				if pre.CheckIsNullExpr(n) {
					HasConstantFoldIndexMap[currentIndex] = true
					currentIndex = -1
				}
			case *ast.InExpr:
				if pre.CheckInExpr(n, currentIndex, node.Expr) {
					HasConstantFoldIndexMap[currentIndex] = true
					currentIndex = -1
				}

			case *ast.BetweenExpr:
				if pre.CheckBetweenExpr(n, currentIndex, node.Expr) {
					HasConstantFoldIndexMap[currentIndex] = true
					currentIndex = -1
				}

			case *ast.LikeExpr:
				if pre.CheckLikeExpr(n, currentIndex, node.Expr) {
					HasConstantFoldIndexMap[currentIndex] = true
					currentIndex = -1
				}
			//where a 等价于 where a!=0
			case *ast.Column:
				node.Expr[currentIndex] = &ast.Expr{Operator: "<>", Left: n, Right: &ast.Integer{Operator: []string{}, Val: "0"}}
			}
		}
	}

}

//and中出现false
//or中出现true
func (pre *ConstantVisitor) RemoveAnd0OrExpr(node *ast.And0OrExpr) ast.Node {
	pre.ReplaceValAndColumn(node)

	if !pre.ReplaceLogicFalseExpr(node) {
		//再执行一次
		n, _ := node.Accept(pre)
		return n
	}

	if pre.hasConstantExprCombine {
		pre.hasConstantExprCombine = false
		//再执行一次
		n, _ := node.Accept(pre)
		return n
	}

	tp := node.GetTag()
	switch tp {
	case ast.Expr_And:
		for index, expr := range node.Expr {
			if expr.Type() == ast.Ast_true {
				e := expr.(*ast.True)
				if e.IsTrue {
					node.Expr[index] = nil
				} else {
					return &ast.True{IsTrue: false}
				}
			} else {
				for i := index + 1; i < len(node.Expr); i++ {
					if expr.IsSameAs(node.Expr[i], pre) {
						node.Expr[i] = &ast.True{IsTrue: true}
					}
				}
			}
		}
	case ast.Expr_Or:
		for index, expr := range node.Expr {
			if expr.Type() == ast.Ast_true {
				e := expr.(*ast.True)
				if e.IsTrue {
					return &ast.True{IsTrue: true}
				} else {
					node.Expr[index] = nil
				}
			} else {
				for i := index + 1; i < len(node.Expr); i++ {
					if expr.IsSameAs(node.Expr[i], pre) {
						node.Expr[i] = &ast.True{IsTrue: false}
					}
				}
			}
		}
	}

	node.Expr = remove(node.Expr, nil)

	switch len(node.Expr) {
	case 0:
		if tp == ast.Expr_And {
			return &ast.True{IsTrue: true}
		} else {
			return &ast.True{IsTrue: false}
		}
	case 1:
		return node.Expr[0]
	default:
		return node
	}

}

func (pre *ConstantVisitor) ConstantFolding(column *ast.Column, value ast.ExprNode, node *ast.And0OrExpr) {
	v := &FoldVisitor{currentColumn: column, currentValue: value}
	node.Accept(v)
}
