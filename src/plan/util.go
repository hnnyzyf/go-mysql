package plan

import "ast"
import "bytes"
import "strings"
import "fmt"
import meta "util/meta"

/*******************************

	PreProcessMeta

********************************/

//func combine(src []string, res []string) {
//	src = append(src, res...)
//}

func appendColumn(node *ast.Star, relation string, meta []string) {
	for _, col := range meta {
		column := &ast.Column{Table: relation, Column: col}
		node.Column = append(node.Column, column)
	}
}

func lower(s string) string {
	return strings.ToLower(s)
}

/*******************************

	constantfold

********************************/

//判断是不是a=3形式的表达式
func isEqualExpr(node *ast.Expr) (*ast.Column, ast.ExprNode, bool) {
	if node.Operator != "=" {
		return nil, nil, false
	}
	ltp := node.Left.Type()
	rtp := node.Right.Type()

	if ltp == ast.Ast_column {
		switch rtp {
		case ast.Ast_true, ast.Ast_integer, ast.Ast_float, ast.Ast_string, ast.Ast_null, ast.Ast_marker:
			return node.Left.(*ast.Column), node.Right, true
		default:
			return nil, nil, false
		}
	}

	if rtp == ast.Ast_column {
		switch ltp {
		case ast.Ast_true, ast.Ast_integer, ast.Ast_float, ast.Ast_string, ast.Ast_null, ast.Ast_marker:
			node.Left, node.Right = node.Right, node.Left
			return node.Left.(*ast.Column), node.Right, true
		default:
			return nil, nil, false
		}
	}

	return nil, nil, false
}

func isMathOp(op string) bool {
	if op == "+" || op == "-" || op == "*" || op == "/" {
		return true
	} else {
		return false
	}
}

func isConstant(n ast.ExprNode) bool {
	switch n.(type) {
	case *ast.Integer, *ast.Float, *ast.String, *ast.True, *ast.Null, *ast.Marker:
		return true
	default:
		return false

	}
}

func isConstantList(n []ast.ExprNode) bool {
	for _, expr := range n {
		if !isConstant(expr) {
			return false
		}
	}
	return true
}

//存储符号相对的符号
var ReOperator = map[string]string{
	">":  "<",
	"<":  ">",
	">=": "<=",
	"<=": ">=",
	"<>": "<>",
	"!=": "!=",
}

func IsCompareOperator(op string) bool {
	switch op {
	case ">=", "<=", ">", "<", "<>", "!=":
		return true
	default:
		return false
	}
}

//判断是不是a>3形式的表达式
func isCompareExpr(node *ast.Expr) (*ast.Column, ast.ExprNode, string, bool) {
	if !IsCompareOperator(node.Operator) {
		return nil, nil, "", false
	}

	ltp := node.Left.Type()
	rtp := node.Right.Type()

	if ltp == ast.Ast_column {
		switch rtp {
		case ast.Ast_true, ast.Ast_integer, ast.Ast_float, ast.Ast_string, ast.Ast_null, ast.Ast_marker:
			return node.Left.(*ast.Column), node.Right, node.Operator, true
		default:
			return nil, nil, "", false
		}
	}

	if rtp == ast.Ast_column {
		switch ltp {
		case ast.Ast_true, ast.Ast_integer, ast.Ast_float, ast.Ast_string, ast.Ast_null, ast.Ast_marker:
			node.Left, node.Right = node.Right, node.Left
			node.Operator = ReOperator[node.Operator]
			return node.Left.(*ast.Column), node.Right, node.Operator, true
		default:
			return nil, nil, "", false
		}
	}

	return nil, nil, "", false
}

func isEqualCondition(n ast.ExprNode) (*ast.Column, *ast.Column, bool) {
	if n.Type() != ast.Ast_expr {
		return nil, nil, false
	}

	node := n.(*ast.Expr)

	if node.Operator != "=" {
		return nil, nil, false
	}

	if node.Left.Type() != ast.Ast_column || node.Right.Type() != ast.Ast_column {
		return nil, nil, false
	}

	return node.Left.(*ast.Column), node.Right.(*ast.Column), true
}

//查看是否是simplerelation
func isRelationTable(alias string, ctx *Context) bool {
	relation := ctx.Table[alias]
	if relation.GetTag() == ast.Relation_Table {
		return true
	} else {
		return false
	}
}

//生成格式化字符串
func isSameString(src []string, org []string) bool {
	if len(src) != len(org) {
		return false
	}
	for _, scolumn := range src {
		if !hasColumn(org, scolumn) {
			return false
		}
	}
	return true
}

//生成格式化字符串
func generateString(n ast.ExprNode) string {
	buf := bytes.NewBuffer([]byte{})
	n.Format(buf)
	return buf.String()
}

//生成格式化字符串
func generateQuery(org ast.ExprNode, res ast.ExprNode, op string) string {
	//生成比较的字符串
	buf := bytes.NewBuffer([]byte{})
	org.Format(buf)
	buf.WriteString(op)
	res.Format(buf)

	return buf.String()
}

//select count(1) from (select count(1) from Table Force INDEX(PRIMARY) group by index1,index2) as temp
func generateSelectivityQuery(row int, table string, index []string) string {
	buf := bytes.NewBuffer([]byte{})

	fmt.Fprint(buf, "SELECT COUNT(1) FROM ")

	fmt.Fprint(buf, " ( Select DISTINCT ")
	for idx, column := range index {
		fmt.Fprint(buf, column)
		if idx != len(index)-1 {
			fmt.Fprint(buf, ",")
		} else {
			fmt.Fprint(buf, " ")
		}
	}
	fmt.Fprint(buf, " FROM ")
	fmt.Fprint(buf, table)
	fmt.Fprint(buf, " LIMIT 0, ")
	fmt.Fprint(buf, row)
	fmt.Fprint(buf, ") AS temp2 ")

	return buf.String()
}

//select count(1) from (select count(1) from Table Force INDEX(PRIMARY) group by index1,index2) as temp
func generateRangeQuery(table string, expr []ast.ExprNode) string {
	buf := bytes.NewBuffer([]byte{})

	fmt.Fprint(buf, "SELECT COUNT(1) FROM ")
	fmt.Fprint(buf, table)
	fmt.Fprint(buf, " WHERE  ")
	for idx, e := range expr {
		e.Format(buf)
		if idx != len(expr)-1 {
			fmt.Fprint(buf, " AND ")
		} else {
			fmt.Fprint(buf, " ")
		}
	}

	return buf.String()
}

func generateIndex(table string, idx []string) string {
	buf := bytes.NewBuffer([]byte{})
	fmt.Fprint(buf, "ALTER TABLE ")
	fmt.Fprint(buf, table)
	fmt.Fprint(buf, " ADD INDEX idx_")
	for id, col := range idx {
		fmt.Fprint(buf, col)
		if id < len(idx)-1 {
			fmt.Fprint(buf, "_")
		}
	}

	fmt.Fprint(buf, " ON ")
	fmt.Fprint(buf, table)
	fmt.Fprint(buf, "(")
	for id, col := range idx {
		fmt.Fprint(buf, col)
		if id < len(idx)-1 {
			fmt.Fprint(buf, ",")
		}
	}
	fmt.Fprint(buf, ")")

	return buf.String()
}

//只检测string
func haveSameType(tp meta.SqlType, n ast.ExprNode) bool {
	if n.Type() == ast.Ast_marker {
		return true
	}

	if n.Type() == ast.Ast_string && tp > meta.Type_date_start && tp < meta.Type_string_end {
		return true
	}

	if (n.Type() == ast.Ast_float || n.Type() == ast.Ast_integer) && tp > meta.Type_digit_start && tp < meta.Type_digit_end {
		return true
	}

	return false
}

func withPercent(str string) bool {
	r := []rune(str)
	if len(r) == 0 {
		return false
	} else {
		if r[0] == '%' || r[len(r)-1] == '%' {
			return true
		} else {
			return false
		}
	}
}

func remove(expr []ast.ExprNode, element ast.ExprNode) []ast.ExprNode {
	res := []ast.ExprNode{}
	for i := range expr {
		if expr[i] != element {
			res = append(res, expr[i])
		}
	}

	return res
}

func initialize(length int) []bool {
	res := []bool{}
	for i := 0; i < length; i++ {
		res = append(res, false)
	}
	return res
}

func hasColumn(src []string, column string) bool {
	for _, col := range src {
		if col == column {
			return true
		}
	}

	return false
}
