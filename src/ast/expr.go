package ast

import "io"
import "fmt"

/***********************************
*
*	实现表达式定义
*
* 	Expr
* 	UnaryExpr
* 	IsTrueExpr
*  	IsNullExpr
*  	SubqueryExpr
*  	BetweenExpr
*  	LikeExpr
*  	RegexpExpr
*  	IntervalExpr
*  	CollateExpr
*  	CaseExpr
*  	InExpr
*
************************************/

//所有一元二元表达式均可使用
//形如a+b a-b a*b这种形式

type Expr struct {
	node

	Operator string
	Left     ExprNode
	Right    ExprNode

	IsEnclosed bool
}

func (e *Expr) Accept(v Visitor) (Node, bool) {

	if !v.Notify(e) {
		return v.Visit(e)
	}

	if e.Left != nil {
		node, ok := e.Left.Accept(v)
		if !ok {
			return e, false
		}
		e.Left = node.(ExprNode)
	}

	if e.Right != nil {
		node, ok := e.Right.Accept(v)
		if !ok {
			return e, false
		}
		e.Right = node.(ExprNode)
	}
	return v.Visit(e)
}

func (e *Expr) Format(w io.Writer) {
	if e.IsEnclosed {
		fmt.Fprint(w, "(")
	}
	if e.Left != nil {
		e.Left.Format(w)
	}

	if e.Operator != "" {
		fmt.Fprint(w, " ")
		fmt.Fprint(w, e.Operator)
		fmt.Fprint(w, " ")
	}

	if e.Right != nil {
		e.Right.Format(w)
	}
	if e.IsEnclosed {
		fmt.Fprint(w, ")")
	}
}

func (e *Expr) Type() int {
	return Ast_expr
}

func (e *Expr) IsSameAs(src ExprNode, v Visitor) bool {
	if src.Type() != e.Type() {
		return false
	}

	if e.GetTag() != src.GetTag() {
		return false
	}
	other := src.(*Expr)
	if e.Operator != other.Operator {
		return false
	}

	if e.Left != nil && other.Left != nil {
		if !e.Left.IsSameAs(other.Left, v) {
			return false
		}
	}
	if (e.Left != nil && other.Left == nil) || (e.Left == nil && other.Left != nil) {
		return false
	}

	if e.Right != nil && other.Right != nil {
		if !e.Right.IsSameAs(other.Right, v) {
			return false
		}
	}
	if (e.Right != nil && other.Right == nil) || (e.Right == nil && other.Right != nil) {
		return false
	}

	return true
}

const (
	Expr_And = iota + 1
	Expr_Or
	Expr_Single
)

//用于Preprocess过程中的替换
type And0OrExpr struct {
	node

	Operator string
	Expr     []ExprNode

	IsEnclosed bool
}

func (e *And0OrExpr) Accept(v Visitor) (Node, bool) {

	if !v.Notify(e) {
		return v.Visit(e)
	}

	for index, target := range e.Expr {
		node, ok := target.Accept(v)
		if !ok {
			return e, false
		}
		e.Expr[index] = node.(ExprNode)

	}

	return v.Visit(e)
}

func (e *And0OrExpr) Format(w io.Writer) {
	if e.IsEnclosed {
		fmt.Fprint(w, "(")
	}
	length := len(e.Expr)
	for index, target := range e.Expr {
		target.Format(w)

		if index != length-1 {
			fmt.Fprint(w, " ")
			fmt.Fprint(w, e.Operator)
			fmt.Fprint(w, " ")
		} else {
			fmt.Fprint(w, " ")
		}
	}
	if e.IsEnclosed {
		fmt.Fprint(w, ")")
	}
}

func (e *And0OrExpr) Type() int {
	return Ast_and0orexpr
}

func (e *And0OrExpr) IsSameAs(src ExprNode, v Visitor) bool {
	if src.Type() != e.Type() {
		return false
	}

	if e.GetTag() != src.GetTag() {
		return false
	}

	other := src.(*And0OrExpr)

	if e.Operator != other.Operator {
		return false
	}

	if len(e.Expr) != len(other.Expr) {
		return false
	}
	for _, expr := range e.Expr {
		hasSame := false
		for _, otherexpr := range other.Expr {
			if expr.IsSameAs(otherexpr, v) {
				hasSame = true
				break
			}
		}
		if !hasSame {
			return false
		}
	}

	return true
}

func (e *And0OrExpr) Combine(l ExprNode, r ExprNode) {
	switch node := l.(type) {
	case *And0OrExpr:
		if node.GetTag() == e.GetTag() {
			for idx, _ := range node.Expr {
				e.Expr = append(e.Expr, node.Expr[idx])
			}
		} else {
			e.Expr = append(e.Expr, l)
		}
	default:
		e.Expr = append(e.Expr, l)
	}

	switch node := r.(type) {
	case *And0OrExpr:
		if node.GetTag() == e.GetTag() {
			for idx, _ := range node.Expr {
				e.Expr = append(e.Expr, node.Expr[idx])
			}
		} else {
			e.Expr = append(e.Expr, r)
		}
	default:
		e.Expr = append(e.Expr, r)
	}

}

const (
	Unary_Constant = iota + 1
	Unary_Other
)

type UnaryExpr struct {
	node

	Operator string
	Expr     ExprNode
}

func (e *UnaryExpr) Accept(v Visitor) (Node, bool) {
	if !v.Notify(e) {
		return v.Visit(e)
	}

	node, ok := e.Expr.Accept(v)
	if !ok {
		return e, false
	}
	e.Expr = node.(ExprNode)

	return v.Visit(e)

}

func (e *UnaryExpr) Format(w io.Writer) {

	if e.Operator != "" {
		fmt.Fprint(w, " ")
		fmt.Fprint(w, e.Operator)
	}

	e.Expr.Format(w)

}

func (e *UnaryExpr) Type() int {
	return Ast_unaryexpr
}

func (e *UnaryExpr) IsSameAs(src ExprNode, v Visitor) bool {
	if src.Type() != e.Type() {
		return false
	}

	if e.GetTag() != src.GetTag() {
		return false
	}

	other := src.(*UnaryExpr)

	if e.Operator != other.Operator {
		return false
	}

	return e.Expr.IsSameAs(other.Expr, v)
}

type NotExpr struct {
	node

	Expr ExprNode
}

func (e *NotExpr) Accept(v Visitor) (Node, bool) {

	if !v.Notify(e) {
		return v.Visit(e)
	}

	node, ok := e.Expr.Accept(v)
	if !ok {
		return e, false
	}
	e.Expr = node.(ExprNode)

	return v.Visit(e)
}

func (e *NotExpr) Format(w io.Writer) {
	fmt.Fprint(w, " NOT ")
	e.Expr.Format(w)
}

func (e *NotExpr) Type() int {
	return Ast_notexpr
}

func (e *NotExpr) IsSameAs(src ExprNode, v Visitor) bool {
	if src.Type() != e.Type() {
		return false
	}

	if e.GetTag() != src.GetTag() {
		return false
	}

	other := src.(*NotExpr)

	return e.Expr.IsSameAs(other.Expr, v)
}

type IsTrueExpr struct {
	node

	HasNot bool
	IsTrue bool
	Expr   ExprNode
}

func (e *IsTrueExpr) Accept(v Visitor) (Node, bool) {

	if !v.Notify(e) {
		return v.Visit(e)
	}

	node, ok := e.Expr.Accept(v)
	if !ok {
		return e, false
	}
	e.Expr = node.(ExprNode)

	return v.Visit(e)
}

func (e *IsTrueExpr) Format(w io.Writer) {

	e.Expr.Format(w)
	if e.HasNot == true {
		fmt.Fprint(w, " IS NOT ")
	} else {
		fmt.Fprint(w, " IS ")
	}
	if e.IsTrue {
		fmt.Fprint(w, "TRUE ")
	} else {
		fmt.Fprint(w, "FALSE ")
	}
}

func (e *IsTrueExpr) Type() int {
	return Ast_istrueexpr
}

func (e *IsTrueExpr) IsSameAs(src ExprNode, v Visitor) bool {
	if src.Type() != e.Type() {
		return false
	}

	if e.GetTag() != src.GetTag() {
		return false
	}

	other := src.(*IsTrueExpr)

	if e.HasNot != other.HasNot {
		return false
	}

	if e.IsTrue != other.IsTrue {
		return false
	}

	return e.Expr.IsSameAs(other.Expr, v)
}

type IsNullExpr struct {
	node

	HasNot bool
	Left   ExprNode
}

func (e *IsNullExpr) Accept(v Visitor) (Node, bool) {

	if !v.Notify(e) {
		return v.Visit(e)
	}

	node, ok := e.Left.Accept(v)
	if !ok {
		return e, false
	}
	e.Left = node.(ExprNode)

	return v.Visit(e)
}

func (e *IsNullExpr) Format(w io.Writer) {

	e.Left.Format(w)
	if e.HasNot == true {
		fmt.Fprint(w, " IS NOT NULL ")
	} else {
		fmt.Fprint(w, " IS NULL ")
	}
}

func (e *IsNullExpr) Type() int {
	return Ast_isnullexpr
}

func (e *IsNullExpr) IsSameAs(src ExprNode, v Visitor) bool {
	if src.Type() != e.Type() {
		return false
	}

	if e.GetTag() != src.GetTag() {
		return false
	}

	other := src.(*IsNullExpr)

	if e.HasNot != other.HasNot {
		return false
	}

	return e.Left.IsSameAs(other.Left, v)
}

const (
	Operator_Some = iota
	Operator_All
	Operator_Any
)

var ReOp = map[int]string{
	Operator_Some: "SOME",
	Operator_All:  "ALL",
	Operator_Any:  "ANY",
}

const (
	Subquery_Operator = iota
	Subquery_In
	Subquery_Exists
)

type SubqueryExpr struct {
	node

	//比较形式的Subquery
	Operator string
	Subtype  int

	//In形式的Subquery
	HasNot bool

	Left ExprNode

	//select形式
	Select Node
	//realtion形式或者tuple形式的subquery
	Alias string
}

func (e *SubqueryExpr) Accept(v Visitor) (Node, bool) {

	if !v.Notify(e) {
		return v.Visit(e)
	}

	if e.Left != nil {
		node, ok := e.Left.Accept(v)
		if !ok {
			return e, false
		}
		e.Left = node.(ExprNode)
	}

	node, ok := e.Select.Accept(v)

	if !ok {
		return e, false
	}
	e.Select = node

	return v.Visit(e)
}

func (e *SubqueryExpr) Format(w io.Writer) {
	switch e.GetTag() {
	case Subquery_Operator:
		e.Left.Format(w)
		fmt.Fprint(w, e.Operator)
		fmt.Fprint(w, ReOp[e.Subtype])
		e.Select.(ExprNode).Format(w)
	case Subquery_In:
		e.Left.Format(w)
		if e.HasNot {
			fmt.Fprint(w, " NOT IN ")
		} else {
			fmt.Fprint(w, " IN ")
		}
		e.Select.(ExprNode).Format(w)
	case Subquery_Exists:
		if e.HasNot {
			fmt.Fprint(w, " NOT EXISTS ")
		} else {
			fmt.Fprint(w, " EXISTS ")
		}
		e.Select.(ExprNode).Format(w)
	}
}

func (e *SubqueryExpr) Type() int {
	return Ast_subqueryexpr
}

func (e *SubqueryExpr) IsSameAs(src ExprNode, v Visitor) bool {
	if src == nil {
		return false
	}
	if src.Type() != e.Type() {
		return false
	}

	if e.GetTag() != src.GetTag() {
		return false
	}

	other := src.(*SubqueryExpr)

	if e.Operator != other.Operator {
		return false
	}

	if e.Subtype != other.Subtype {
		return false
	}

	if e.HasNot != other.HasNot {
		return false
	}

	if e.Left != nil {
		if !e.Left.IsSameAs(other.Left, v) {
			return false
		}

	}

	if !e.Select.(ExprNode).IsSameAs(other.Select.(ExprNode), v) {
		return false
	}

	return true
}

//形如 between类型的比较
type BetweenExpr struct {
	node

	HasNot bool
	Expr   ExprNode
	Left   ExprNode
	Right  ExprNode
}

func (e *BetweenExpr) Accept(v Visitor) (Node, bool) {

	if !v.Notify(e) {
		return v.Visit(e)
	}

	node, ok := e.Expr.Accept(v)
	if !ok {
		return e, false
	}
	e.Expr = node.(ExprNode)

	node, ok = e.Left.Accept(v)
	if !ok {
		return e, false
	}
	e.Left = node.(ExprNode)

	node, ok = e.Right.Accept(v)
	if !ok {
		return e, false
	}
	e.Right = node.(ExprNode)

	return v.Visit(e)
}

func (e *BetweenExpr) Format(w io.Writer) {
	e.Expr.Format(w)
	if e.HasNot {
		fmt.Fprint(w, " NOT BETWEEN ")
	} else {
		fmt.Fprint(w, " BETWEEN ")
	}
	e.Left.Format(w)
	fmt.Fprint(w, " AND ")
	e.Right.Format(w)
}

func (e *BetweenExpr) Type() int {
	return Ast_betweenexpr
}

func (e *BetweenExpr) IsSameAs(src ExprNode, v Visitor) bool {
	if src.Type() != e.Type() {
		return false
	}

	if e.GetTag() != src.GetTag() {
		return false
	}

	other := src.(*BetweenExpr)

	if e.HasNot != other.HasNot {
		return false
	}
	if !e.Expr.IsSameAs(other.Expr, v) {
		return false
	}

	if !e.Left.IsSameAs(other.Left, v) {
		return false
	}

	if !e.Right.IsSameAs(other.Right, v) {
		return false
	}

	return true
}

type LikeExpr struct {
	node

	HasNot bool
	Left   ExprNode
	Right  ExprNode
}

func (e *LikeExpr) Accept(v Visitor) (Node, bool) {

	if !v.Notify(e) {
		return v.Visit(e)
	}

	node, ok := e.Left.Accept(v)
	if !ok {
		return e, false
	}
	e.Left = node.(ExprNode)

	node, ok = e.Right.Accept(v)
	if !ok {
		return e, false
	}
	e.Right = node.(ExprNode)

	return v.Visit(e)
}

func (e *LikeExpr) Format(w io.Writer) {
	e.Left.Format(w)
	if e.HasNot {
		fmt.Fprint(w, " NOT LIKE ")
	} else {
		fmt.Fprint(w, " LIKE ")
	}
	e.Right.Format(w)
}

func (e *LikeExpr) Type() int {
	return Ast_likeexpr
}

func (e *LikeExpr) IsSameAs(src ExprNode, v Visitor) bool {
	if src.Type() != e.Type() {
		return false
	}

	if e.GetTag() != src.GetTag() {
		return false
	}

	other := src.(*LikeExpr)

	if e.HasNot != other.HasNot {
		return false
	}

	if !e.Left.IsSameAs(other.Left, v) {
		return false
	}

	if !e.Right.IsSameAs(other.Right, v) {
		return false
	}

	return true
}

type RegexpExpr struct {
	node

	HasNot bool
	Left   ExprNode
	Right  ExprNode
}

func (e *RegexpExpr) Accept(v Visitor) (Node, bool) {

	if !v.Notify(e) {
		return v.Visit(e)
	}

	node, ok := e.Left.Accept(v)
	if !ok {
		return e, false
	}
	e.Left = node.(ExprNode)

	node, ok = e.Right.Accept(v)
	if !ok {
		return e, false
	}
	e.Right = node.(ExprNode)

	return v.Visit(e)
}

func (e *RegexpExpr) Format(w io.Writer) {
	e.Left.Format(w)
	if e.HasNot {
		fmt.Fprint(w, " NOT REGEXP ")
	} else {
		fmt.Fprint(w, " REGEXP ")
	}
	e.Right.Format(w)
}

func (e *RegexpExpr) Type() int {
	return Ast_regexpexpr
}

func (e *RegexpExpr) IsSameAs(src ExprNode, v Visitor) bool {
	if src.Type() != e.Type() {
		return false
	}

	if e.GetTag() != src.GetTag() {
		return false
	}

	other := src.(*RegexpExpr)

	if e.HasNot != other.HasNot {
		return false
	}

	if !e.Left.IsSameAs(other.Left, v) {
		return false
	}

	if !e.Right.IsSameAs(other.Right, v) {
		return false
	}

	return true
}

var TimeUnit = map[string]int{
	"YEAR":               Time_year,
	"YEAR_MONTH":         Time_year_month,
	"WEEK":               Time_week,
	"SECOND":             Time_second,
	"SECOND_MICROSECOND": Time_second_microsecond,
	"MONTH":              Time_month,
	"MINUTE":             Time_minute,
	"MINUTE_MICROSECOND": Time_minute_microsecond,
	"MINUTE_SECOND":      Time_minute_second,
	"MICROSECOND":        Time_microsecond,
	"HOUR":               Time_hour,
	"HOUR_MICROSECOND":   Time_hour_microsecond,
	"HOUR_MINUTE":        Time_hour_minute,
	"HOUR_SECOND":        Time_hour_second,
	"DAY":                Time_day,
	"DAY_HOUR":           Time_day_hour,
	"DAY_MICROSECOND":    Time_day_microsecond,
	"DAY_MINUTE":         Time_day_minute,
	"DAY_SECOND":         Time_day_second,
}

var ReTimeUnit = map[int]string{
	Time_year:               "YEAR",
	Time_year_month:         "YEAR_MONTH",
	Time_week:               "WEEK",
	Time_second:             "SECOND",
	Time_second_microsecond: "SECOND_MICROSECOND",
	Time_month:              "MONTH",
	Time_minute:             "MINUTE",
	Time_minute_microsecond: "MINUTE_MICROSECOND",
	Time_minute_second:      "MINUTE_SECOND",
	Time_microsecond:        "MICROSECOND",
	Time_hour:               "HOUR",
	Time_hour_microsecond:   "HOUR_MICROSECOND",
	Time_hour_minute:        "HOUR_MINUTE",
	Time_hour_second:        "HOUR_SECOND",
	Time_day:                "DAY",
	Time_day_hour:           "DAY_HOUR",
	Time_day_microsecond:    "DAY_MICROSECOND",
	Time_day_minute:         "DAY_MINUTE",
	Time_day_second:         "DAY_SECOND",
}

const (
	Time_year = iota
	Time_year_month
	Time_week
	Time_second
	Time_second_microsecond
	Time_month
	Time_minute
	Time_minute_microsecond
	Time_minute_second
	Time_microsecond
	Time_hour
	Time_hour_microsecond
	Time_hour_minute
	Time_hour_second
	Time_day
	Time_day_hour
	Time_day_microsecond
	Time_day_minute
	Time_day_second
)

const (
	Op_Add = iota
	Op_Minus
)

type IntervalExpr struct {
	node

	Operator int
	Left     ExprNode
	Right    ExprNode
	TimeUnit int
}

func (e *IntervalExpr) Accept(v Visitor) (Node, bool) {

	if !v.Notify(e) {
		return v.Visit(e)
	}

	node, ok := e.Left.Accept(v)
	if !ok {
		return e, false
	}
	e.Left = node.(ExprNode)

	node, ok = e.Right.Accept(v)
	if !ok {
		return e, false
	}
	e.Right = node.(ExprNode)

	return v.Visit(e)
}

func (e *IntervalExpr) Format(w io.Writer) {
	e.Left.Format(w)
	if e.Operator == Op_Add {
		fmt.Fprint(w, " + INTERVAL ")
	} else {
		fmt.Fprint(w, " - INTERVAL ")
	}
	e.Right.Format(w)

	fmt.Fprint(w, ReTimeUnit[e.TimeUnit])

}

func (e *IntervalExpr) Type() int {
	return Ast_intervalexpr
}

func (e *IntervalExpr) IsSameAs(src ExprNode, v Visitor) bool {
	if src.Type() != e.Type() {
		return false
	}

	if e.GetTag() != src.GetTag() {
		return false
	}

	other := src.(*IntervalExpr)

	if e.Operator != other.Operator {
		return false
	}

	if e.TimeUnit != other.TimeUnit {
		return false
	}
	if !e.Left.IsSameAs(other.Left, v) {
		return false
	}

	if !e.Right.IsSameAs(other.Right, v) {
		return false
	}

	return true
}

const (
	collate_armscii8_general_ci = iota
	collate_ascii_general_ci
	collate_big5_chinese_ci
	collate_binary
	collate_cp1250_general_ci
	collate_cp1251_general_ci
	collate_cp1256_general_ci
	collate_cp1257_general_ci
	collate_cp850_general_ci
	collate_cp852_general_ci
	collate_cp866_general_ci
	collate_cp932_japanese_ci
	collate_dec8_swedish_ci
	collate_eucjpms_japanese_ci
	collate_euckr_korean_ci
	collate_gb18030_chinese_ci
	collate_gb2312_chinese_ci
	collate_gbk_chinese_ci
	collate_geostd8_general_ci
	collate_greek_general_ci
	collate_hebrew_general_ci
	collate_hp8_english_ci
	collate_keybcs2_general_ci
	collate_koi8r_general_ci
	collate_koi8u_general_ci
	collate_latin1_swedish_ci
	collate_latin2_general_ci
	collate_latin5_turkish_ci
	collate_latin7_general_ci
	collate_macce_general_ci
	collate_macroman_general_ci
	collate_sjis_japanese_ci
	collate_swe7_swedish_ci
	collate_tis620_thai_ci
	collate_ucs2_general_ci
	collate_ujis_japanese_ci
	collate_utf16_general_ci
	collate_utf16le_general_ci
	collate_utf32_general_ci
	collate_utf8_general_ci
	collate_utf8mb4_general_ci
)

var Character = map[string]int{
	"armscii8_general_ci": collate_armscii8_general_ci,
	"ascii_general_ci":    collate_ascii_general_ci,
	"big5_chinese_ci":     collate_big5_chinese_ci,
	"binary":              collate_binary,
	"cp1250_general_ci":   collate_cp1250_general_ci,
	"cp1251_general_ci":   collate_cp1251_general_ci,
	"cp1256_general_ci":   collate_cp1256_general_ci,
	"cp1257_general_ci":   collate_cp1257_general_ci,
	"cp850_general_ci":    collate_cp850_general_ci,
	"cp852_general_ci":    collate_cp852_general_ci,
	"cp866_general_ci":    collate_cp866_general_ci,
	"cp932_japanese_ci":   collate_cp932_japanese_ci,
	"dec8_swedish_ci":     collate_dec8_swedish_ci,
	"eucjpms_japanese_ci": collate_eucjpms_japanese_ci,
	"euckr_korean_ci":     collate_euckr_korean_ci,
	"gb18030_chinese_ci":  collate_gb18030_chinese_ci,
	"gb2312_chinese_ci":   collate_gb2312_chinese_ci,
	"gbk_chinese_ci":      collate_gbk_chinese_ci,
	"geostd8_general_ci":  collate_geostd8_general_ci,
	"greek_general_ci":    collate_greek_general_ci,
	"hebrew_general_ci":   collate_hebrew_general_ci,
	"hp8_english_ci":      collate_hp8_english_ci,
	"keybcs2_general_ci":  collate_keybcs2_general_ci,
	"koi8r_general_ci":    collate_koi8r_general_ci,
	"koi8u_general_ci":    collate_koi8u_general_ci,
	"latin1_swedish_ci":   collate_latin1_swedish_ci,
	"latin2_general_ci":   collate_latin2_general_ci,
	"latin5_turkish_ci":   collate_latin5_turkish_ci,
	"latin7_general_ci":   collate_latin7_general_ci,
	"macce_general_ci":    collate_macce_general_ci,
	"macroman_general_ci": collate_macroman_general_ci,
	"sjis_japanese_ci":    collate_sjis_japanese_ci,
	"swe7_swedish_ci":     collate_swe7_swedish_ci,
	"tis620_thai_ci":      collate_tis620_thai_ci,
	"ucs2_general_ci":     collate_ucs2_general_ci,
	"ujis_japanese_ci":    collate_ujis_japanese_ci,
	"utf16_general_ci":    collate_utf16_general_ci,
	"utf16le_general_ci":  collate_utf16le_general_ci,
	"utf32_general_ci":    collate_utf32_general_ci,
	"utf8_general_ci":     collate_utf8_general_ci,
	"utf8mb4_general_ci":  collate_utf8mb4_general_ci,
}

var ReCharacter = map[int]string{
	collate_armscii8_general_ci: "armscii8_general_ci",
	collate_ascii_general_ci:    "ascii_general_ci",
	collate_big5_chinese_ci:     "big5_chinese_ci",
	collate_binary:              "binary",
	collate_cp1250_general_ci:   "cp1250_general_ci",
	collate_cp1251_general_ci:   "cp1251_general_ci",
	collate_cp1256_general_ci:   "cp1256_general_ci",
	collate_cp1257_general_ci:   "cp1257_general_ci",
	collate_cp850_general_ci:    "cp850_general_ci",
	collate_cp852_general_ci:    "cp852_general_ci",
	collate_cp866_general_ci:    "cp866_general_ci",
	collate_cp932_japanese_ci:   "cp932_japanese_ci",
	collate_dec8_swedish_ci:     "dec8_swedish_ci",
	collate_eucjpms_japanese_ci: "eucjpms_japanese_ci",
	collate_euckr_korean_ci:     "euckr_korean_ci",
	collate_gb18030_chinese_ci:  "gb18030_chinese_ci",
	collate_gb2312_chinese_ci:   "gb2312_chinese_ci",
	collate_gbk_chinese_ci:      "gbk_chinese_ci",
	collate_geostd8_general_ci:  "geostd8_general_ci",
	collate_greek_general_ci:    "greek_general_ci",
	collate_hebrew_general_ci:   "hebrew_general_ci",
	collate_hp8_english_ci:      "hp8_english_ci",
	collate_keybcs2_general_ci:  "keybcs2_general_ci",
	collate_koi8r_general_ci:    "koi8r_general_ci",
	collate_koi8u_general_ci:    "koi8u_general_ci",
	collate_latin1_swedish_ci:   "latin1_swedish_ci",
	collate_latin2_general_ci:   "latin2_general_ci",
	collate_latin5_turkish_ci:   "latin5_turkish_ci",
	collate_latin7_general_ci:   "latin7_general_ci",
	collate_macce_general_ci:    "macce_general_ci",
	collate_macroman_general_ci: "macroman_general_ci",
	collate_sjis_japanese_ci:    "sjis_japanese_ci",
	collate_swe7_swedish_ci:     "swe7_swedish_ci",
	collate_tis620_thai_ci:      "tis620_thai_ci",
	collate_ucs2_general_ci:     "ucs2_general_ci",
	collate_ujis_japanese_ci:    "ujis_japanese_ci",
	collate_utf16_general_ci:    "utf16_general_ci",
	collate_utf16le_general_ci:  "utf16le_general_ci",
	collate_utf32_general_ci:    "utf32_general_ci",
	collate_utf8_general_ci:     "utf8_general_ci",
	collate_utf8mb4_general_ci:  "utf8mb4_general_ci",
}

type CollateExpr struct {
	node

	Expr    ExprNode
	Collate int
}

func (e *CollateExpr) Accept(v Visitor) (Node, bool) {

	if !v.Notify(e) {
		return v.Visit(e)
	}

	node, ok := e.Expr.Accept(v)
	if !ok {
		return e, false
	}
	e.Expr = node.(ExprNode)

	return v.Visit(e)
}

func (e *CollateExpr) Format(w io.Writer) {
	e.Expr.Format(w)
	fmt.Fprint(w, " COLLATE ")

	fmt.Fprint(w, ReCharacter[e.Collate])

}

func (e *CollateExpr) Type() int {
	return Ast_collateexpr
}

func (e *CollateExpr) IsSameAs(src ExprNode, v Visitor) bool {
	if src.Type() != e.Type() {
		return false
	}

	if e.GetTag() != src.GetTag() {
		return false
	}

	other := src.(*CollateExpr)

	if e.Collate != other.Collate {
		return false
	}

	if !e.Expr.IsSameAs(other.Expr, v) {
		return false
	}

	return true
}

type CaseExpr struct {
	node

	Case ExprNode
	When []ExprNode
	Else ExprNode
}

func (e *CaseExpr) Accept(v Visitor) (Node, bool) {

	if !v.Notify(e) {
		return v.Visit(e)
	}

	if e.Case != nil {
		node, ok := e.Case.Accept(v)
		if !ok {
			return e, false
		}
		e.Case = node.(ExprNode)
	}

	for index, target := range e.When {
		node, ok := target.Accept(v)
		if !ok {
			return e, false
		}
		e.When[index] = node.(ExprNode)

	}

	if e.Else != nil {
		node, ok := e.Else.Accept(v)
		if !ok {
			return e, false
		}
		e.Else = node.(ExprNode)
	}
	return v.Visit(e)
}

func (e *CaseExpr) Format(w io.Writer) {
	fmt.Fprint(w, " CASE ")
	if e.Case != nil {
		e.Case.Format(w)
	}

	for _, target := range e.When {
		fmt.Fprint(w, " WHEN ")
		target.(*Expr).Left.Format(w)
		fmt.Fprint(w, " THEN ")
		target.(*Expr).Right.Format(w)
	}

	fmt.Fprint(w, " ELSE ")
	if e.Else != nil {
		e.Else.Format(w)
	}

	fmt.Fprint(w, " END")
}

func (e *CaseExpr) Type() int {
	return Ast_caseexpr
}

func (e *CaseExpr) IsSameAs(src ExprNode, v Visitor) bool {
	if src.Type() != e.Type() {
		return false
	}

	if e.GetTag() != src.GetTag() {
		return false
	}

	other := src.(*CaseExpr)

	if e.Case != nil && other.Case != nil {
		if !e.Case.IsSameAs(other.Case, v) {
			return false
		}
	}
	if (e.Case != nil && other.Case == nil) || (e.Case == nil && other.Case != nil) {
		return false
	}

	if e.Else != nil && other.Else != nil {
		if !e.Else.IsSameAs(other.Else, v) {
			return false
		}
	}
	if (e.Else != nil && other.Else == nil) || (e.Else == nil && other.Else != nil) {
		return false
	}

	if len(e.When) != len(other.When) {
		return false
	}
	for _, expr := range e.When {
		hasSame := false
		for _, otherexpr := range other.When {
			if expr.IsSameAs(otherexpr, v) {
				hasSame = true
				break
			}
		}
		if !hasSame {
			return false
		}
	}

	return true
}

type InExpr struct {
	node

	HasNot bool
	Left   ExprNode
	Right  []ExprNode
}

func (e *InExpr) Accept(v Visitor) (Node, bool) {

	if !v.Notify(e) {
		return v.Visit(e)
	}

	node, ok := e.Left.Accept(v)
	if !ok {
		return e, false
	}
	e.Left = node.(ExprNode)

	for index, target := range e.Right {
		node, ok := target.Accept(v)
		if !ok {
			return e, false
		}
		e.Right[index] = node.(ExprNode)

	}

	return v.Visit(e)
}

func (e *InExpr) Format(w io.Writer) {
	e.Left.Format(w)

	if e.HasNot {
		fmt.Fprint(w, " NOT IN (")
	} else {
		fmt.Fprint(w, " IN (")
	}

	length := len(e.Right)
	for index, target := range e.Right {
		target.Format(w)
		if index != length-1 {
			fmt.Fprint(w, ",")
		}
	}

	fmt.Fprint(w, ")")

}

func (e *InExpr) Type() int {
	return Ast_inexpr
}

func (e *InExpr) IsSameAs(src ExprNode, v Visitor) bool {
	if src.Type() != e.Type() {
		return false
	}

	if e.GetTag() != src.GetTag() {
		return false
	}

	other := src.(*InExpr)

	if e.HasNot != other.HasNot {
		return false
	}

	if !e.Left.IsSameAs(other.Left, v) {
		return false
	}

	if len(e.Right) != len(other.Right) {
		return false
	}
	for _, expr := range e.Right {
		hasSame := false
		for _, otherexpr := range other.Right {
			if expr.IsSameAs(otherexpr, v) {
				hasSame = true
				break
			}
		}
		if !hasSame {
			return false
		}
	}

	return true
}
