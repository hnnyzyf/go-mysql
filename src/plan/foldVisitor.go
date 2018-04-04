package plan

import "ast"
import server "util/server"

type FoldVisitor struct {
	currentColumn *ast.Column
	currentValue  ast.ExprNode
}

func (pre *FoldVisitor) Notify(n ast.Node) bool {
	return true
}

func (pre *FoldVisitor) Visit(n ast.Node) (ast.Node, bool) {
	switch node := n.(type) {
	case *ast.Column:
		return pre.ReplaceColumnByValue(node), true
	}
	return n, true
}

func (pre *FoldVisitor) ReplaceColumnByValue(node *ast.Column) ast.Node {
	current := pre.currentColumn
	//判断是否是同一个节点
	if current == node {
		return node
	}

	//判断是否是一个column
	if !current.IsSameAs(node, pre) {
		return node
	}

	//均相同的情况,用value替换掉
	return pre.currentValue
}

func (pre *FoldVisitor) Session() server.Server {
	return nil
}
