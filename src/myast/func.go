package myast

//value转换函数
func DoValueExpr(val ExprNode) ExprNode {
	switch val.(type) {
	case *ValueExpr:
		value := val.(*ValueExpr)
		value.Val = "-" + value.Val
		return value
	default:
		return val
	}
}
