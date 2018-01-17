package myast

/*****************************************************************************
 *
 *	布尔表达式
 *
 *****************************************************************************/
//布尔表达式
type BoolExpr struct {
	Ntype NodeType
	Op    string
	Left  ExprNode
	Right ExprNode
}

//初始化函数
func NewBoolExpr(t NodeType) *BoolExpr {
	return &BoolExpr{Ntype: t}
}

//实现ExprNode接口
func (e *BoolExpr) IsExprNode() {}

//实现Node接口
func (e *BoolExpr) Accept(v Visitor) { v.Walk(e) }

/*****************************************************************************
 *
 *	判断表达式
 *
 *****************************************************************************/

type IsExpr struct {
	Ntype NodeType
	Not   bool
	Left  ExprNode
	Right ExprNode
}

//初始化函数
func NewIsExpr(t NodeType) *IsExpr {
	return &IsExpr{Ntype: t}
}

//实现ExprNode接口
func (e *IsExpr) IsExprNode() {}

//实现Node接口
func (e *IsExpr) Accept(v Visitor) { v.Walk(e) }

/*****************************************************************************
 *
 *	between表达式
 *
 *****************************************************************************/

type RangeExpr struct {
	Ntype NodeType
	Expr  ExprNode
	Left  ExprNode
	Right ExprNode
}

//初始化函数
func NewRangeExpr() *RangeExpr {
	return &RangeExpr{Ntype: EXPR_RANGE}
}

//实现ExprNode接口
func (e *RangeExpr) IsExprNode() {}

//实现Node接口
func (e *RangeExpr) Accept(v Visitor) { v.Walk(e) }

/*****************************************************************************
 *
 *	Sublink类型的表达式
 *
 *****************************************************************************/

type InExpr struct {
	Ntype NodeType
	Not   bool
	Left  ExprNode
	Right List
}

//初始化函数
func NewInExpr(t NodeType) *InExpr {
	return &InExpr{Ntype: t}
}

//实现ExprNode接口
func (e *InExpr) IsExprNode() {}

//实现Node接口
func (e *InExpr) Accept(v Visitor) { v.Walk(e) }

/*****************************************************************************
 *
 *	Suquery 类型表达式
 *
 *****************************************************************************/

type SubQueryExpr struct {
	Ntype   NodeType
	Op      string
	Subtype string
	Left    ExprNode
	Right   Node
}

//初始化函数
func NewSubQueryExpr(t NodeType) *SubQueryExpr {
	return &SubQueryExpr{Ntype: t}
}

//实现ExprNode接口
func (e *SubQueryExpr) IsExprNode() {}

//实现Node接口
func (e *SubQueryExpr) Accept(v Visitor) { v.Walk(e) }

/*****************************************************************************
 *
 *	Join 类型表达式
 *
 *****************************************************************************/

type JoinExpr struct {
	Ntype    NodeType
	Jointype string
	Joinqual ExprNode
	Left     ExprNode
	Right    ExprNode
}

//初始化函数
func NewJoinExpr(t NodeType) *JoinExpr {
	return &JoinExpr{Ntype: t}
}

//实现ExprNode接口
func (e *JoinExpr) IsExprNode() {}

//实现Node接口
func (e *JoinExpr) Accept(v Visitor) { v.Walk(e) }

/*****************************************************************************
 *
 *	Join On 类型表达式
 *
 *****************************************************************************/

type OnExpr struct {
	Ntype NodeType
	Expr  ExprNode
}

//初始化函数
func NewOnExpr(t NodeType) *OnExpr {
	return &OnExpr{Ntype: t}
}

//实现ExprNode接口
func (e *OnExpr) IsExprNode() {}

//实现Node接口
func (e *OnExpr) Accept(v Visitor) { v.Walk(e) }

/*****************************************************************************
 *
 *	Join using 类型表达式
 *
 *****************************************************************************/

type UseExpr struct {
	Ntype NodeType
	Expr  List
}

//初始化函数
func NewUseExpr(t NodeType) *UseExpr {
	return &UseExpr{Ntype: t}
}

//实现ExprNode接口
func (e *UseExpr) IsExprNode() {}

//实现Node接口
func (e *UseExpr) Accept(v Visitor) { v.Walk(e) }

/*****************************************************************************
 *
 *	Function 类型表达式
 *
 *****************************************************************************/

type FuncExpr struct {
	Ntype    NodeType
	Funcname string
	Expr     List
	Collate  string
}

//初始化函数
func NewFuncExpr(t NodeType) *FuncExpr {
	return &FuncExpr{Ntype: t}
}

//实现ExprNode接口
func (e *FuncExpr) IsExprNode() {}

//实现Node接口
func (e *FuncExpr) Accept(v Visitor) { v.Walk(e) }

/*****************************************************************************
 *
 *	时间类型 类型表达式
 *
 *****************************************************************************/

type TimeExpr struct {
	Ntype NodeType
	Func  string
}

//初始化函数
func NewTimeExpr(t NodeType) *TimeExpr {
	return &TimeExpr{Ntype: t}
}

//实现ExprNode接口
func (e *TimeExpr) IsExprNode() {}

//实现Node接口
func (e *TimeExpr) Accept(v Visitor) { v.Walk(e) }

/*****************************************************************************
 *
 *	case_expr 类型表达式
 *
 *****************************************************************************/
type CaseExpr struct {
	Ntype NodeType
	Expr  ExprNode
	When  List
	Else  ExprNode
}

//初始化函数
func NewCaseExpr(t NodeType) *CaseExpr {
	return &CaseExpr{Ntype: t}
}

//实现ExprNode接口
func (e *CaseExpr) IsExprNode() {}

//实现Node接口
func (e *CaseExpr) Accept(v Visitor) { v.Walk(e) }
