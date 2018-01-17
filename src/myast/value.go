package myast

/*****************************************************************************
 *
 *	Target for SELECT
 *
 *****************************************************************************/

type TargetExpr struct {
	Ntype NodeType
	Tuple ExprNode
	Alias string
}

//初始化函数
func NewTargetExpr() *TargetExpr {
	return &TargetExpr{Ntype: VALUE_TARGET}
}

//实现ExprNode接口
func (e *TargetExpr) IsExprNode() {}

//实现Node接口
func (e *TargetExpr) Accept(v Visitor) { v.Walk(e) }

/*****************************************************************************
 *
 *	Column for SELECT
 *
 *****************************************************************************/

type ColumnExpr struct {
	Ntype  NodeType
	Table  string
	Column string
}

//初始化函数
func NewColumnExpr() *ColumnExpr {
	return &ColumnExpr{Ntype: VALUE_COLUMN}
}

//实现ExprNode接口
func (e *ColumnExpr) IsExprNode() {}

//实现Node接口
func (e *ColumnExpr) Accept(v Visitor) { v.Walk(e) }

/*****************************************************************************
 *
 *	tablename for SELECT
 *
 *****************************************************************************/

type SimpleTableExpr struct {
	Ntype NodeType
	Table ExprNode
	Alias string
}

//初始化函数
func NewSimpleTableExpr() *SimpleTableExpr {
	return &SimpleTableExpr{Ntype: VALUE_TABLE_SIMPLE}
}

//实现ExprNode接口
func (e *SimpleTableExpr) IsExprNode() {}

//实现Node接口
func (e *SimpleTableExpr) Accept(v Visitor) { v.Walk(e) }

type TableExpr struct {
	Ntype NodeType
	DB    string
	Table string
}

//初始化函数
func NewTableExpr() *TableExpr {
	return &TableExpr{Ntype: VALUE_TABLE_SIMPLE}
}

//实现ExprNode接口
func (e *TableExpr) IsExprNode() {}

//实现Node接口
func (e *TableExpr) Accept(v Visitor) { v.Walk(e) }

type SubTableExpr struct {
	Ntype NodeType
	Stmt  SelectNode
	Alias string
}

//初始化函数
func NewSubTableExpr() *SubTableExpr {
	return &SubTableExpr{Ntype: VALUE_TABLE_SUBTABLE}
}

//实现ExprNode接口
func (e *SubTableExpr) IsExprNode() {}

//实现Node接口
func (e *SubTableExpr) Accept(v Visitor) { v.Walk(e) }

/*****************************************************************************
 *
 *	value for SELECT
 *
 *****************************************************************************/

//实现
type ValueExpr struct {
	Ntype NodeType
	Val   string
}

//初始化函数
func NewValueExpr(t NodeType) *ValueExpr {
	return &ValueExpr{Ntype: t}
}

//实现Expr接口
func (e *ValueExpr) IsExprNode() {}

//实现Node接口
func (e *ValueExpr) Accept(v Visitor) { v.Walk(e) }

/*****************************************************************************
 *
 *	Sortby item for SELECT
 *
 *****************************************************************************/

//实现
type AscExpr struct {
	Ntype NodeType
	Expr  ExprNode
	Not   bool   //判断是否有Asc
	Asc   string //判断是Asc还是Desc
}

//初始化函数
func NewAscExpr() *AscExpr {
	return &AscExpr{Ntype: VALUE_SORTBY_ITEM}
}

//实现Expr接口
func (e *AscExpr) IsExprNode() {}

//实现Node接口
func (e *AscExpr) Accept(v Visitor) { v.Walk(e) }

/*****************************************************************************
 *
 *	Sortby item for SELECT
 *
 *****************************************************************************/

//实现
type StartExpr struct {
	Ntype NodeType
}

//初始化函数
func NewStartExpr() *StartExpr {
	return &StartExpr{Ntype: VALUE_STAR}
}

//实现Expr接口
func (e *StartExpr) IsExprNode() {}

//实现Node接口
func (e *StartExpr) Accept(v Visitor) { v.Walk(e) }

/*****************************************************************************
 *
 *	func arg item for SELECT
 *
 *****************************************************************************/

//实现
type FuncArgExpr struct {
	Ntype     NodeType
	Distinct  string
	Alias     string
	Character string
	TypeAlias ExprNode
	Expr      ExprNode
}

//初始化函数
func NewFuncArgExpr() *FuncArgExpr {
	return &FuncArgExpr{Ntype: EXPR_FUNC_ARG}
}

//实现Expr接口
func (e *FuncArgExpr) IsExprNode() {}

//实现Node接口
func (e *FuncArgExpr) Accept(v Visitor) { v.Walk(e) }



/*****************************************************************************
 *
 *	when arg item for SELECT
 *
 *****************************************************************************/

//实现
type WhenExpr struct {
	Ntype     NodeType
	Left ExprNode
	Right ExprNode
}

//初始化函数
func NewWhenExpr() *WhenExpr {
	return &WhenExpr{Ntype: EXPR_CASE_WHEN}
}

//实现Expr接口
func (e *WhenExpr) IsExprNode() {}

//实现Node接口
func (e *WhenExpr) Accept(v Visitor) { v.Walk(e) }


//实现
type ElseExpr struct {
	Ntype     NodeType
	Expr ExprNode
}

//初始化函数
func NewElseExpr() *ElseExpr {
	return &ElseExpr{Ntype: EXPR_CASE_ELSE}
}

//实现Expr接口
func (e *ElseExpr) IsExprNode() {}

//实现Node接口
func (e *ElseExpr) Accept(v Visitor) { v.Walk(e) }