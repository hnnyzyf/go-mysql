package ast

import "io"
import "fmt"
import "bytes"

/***********************************
*
*	实现Item表达式定义
*
* 	Relation
* 	Tuple
*  	Column
*  	Number
*  	String
*  	True
*  	Null
*  	SystemVariable
*  	UserVariable
*  	Marker
*  	Sortby
*
************************************/

/******************
*	Relation Item
*******************/

const (
	Relation_Table = iota
	Relation_Subquery
)

type Relation struct {
	node

	//传统的Relation形式
	DB    string
	Table string
	//subquery形式的Relation
	Select Node
	//关系指向的别名
	Alias string

	//preprocess过程结果
	//投影中出现*
	HasStar bool
	//投影中涉及的Column信息
	Projection []*Column
}

func (e *Relation) Accept(v Visitor) (Node, bool) {

	if !v.Notify(e) {
		return v.Visit(e)
	}

	if e.GetTag() == Relation_Subquery {
		node, ok := e.Select.Accept(v)
		if !ok {
			return e, false
		}
		e.Select = node
	}

	return v.Visit(e)
}

func (e *Relation) Format(w io.Writer) {
	switch e.GetTag() {
	case Relation_Table:
		if e.DB != "" {
			fmt.Fprint(w, e.DB)
			fmt.Fprint(w, ".")
		}
		fmt.Fprint(w, e.Table)
	case Relation_Subquery:
		e.Select.(ExprNode).Format(w)
	}

	if e.Alias != "" {
		fmt.Fprint(w, " AS ")
	}
	fmt.Fprint(w, e.Alias)
}

func (e *Relation) GetAlias() string {
	if e.Alias != "" {
		return e.Alias
	} else {
		buf := bytes.NewBuffer([]byte{})
		e.Format(buf)
		e.Alias = buf.String()
		return e.Alias
	}
}

/******************
*	Tuple Item
*******************/

const (
	Tuple_expr = iota
	Tuple_column
	Tuple_subquery
	Tuple_funcall
	Tuple_agg
)

type Tuple struct {
	node

	Ref ExprNode
	//关系指向的别名

	Alias string

	//Tuple在Field中的位置
	Location int
}

func (e *Tuple) Accept(v Visitor) (Node, bool) {

	if !v.Notify(e) {
		return v.Visit(e)
	}

	node, ok := e.Ref.Accept(v)
	if !ok {
		return e, false
	}
	e.Ref = node.(ExprNode)

	return v.Visit(e)
}

func (e *Tuple) Format(w io.Writer) {
	e.Ref.Format(w)
	if e.Alias != "" {
		fmt.Fprint(w, " AS ")
	}
	fmt.Fprint(w, e.Alias)
}

func (e *Tuple) GetAlias() string {
	if e.Alias != "" {
		return e.Alias
	} else {
		buf := bytes.NewBuffer([]byte{})
		e.Format(buf)
		e.Alias = buf.String()
		return e.Alias
	}
}

/******************
*	Column Item
*******************/

type Column struct {
	node

	IsStar bool
	DB     string
	Table  string
	Column string
}

func (e *Column) Accept(v Visitor) (Node, bool) {

	if !v.Notify(e) {
		return v.Visit(e)
	}

	return v.Visit(e)
}

func (e *Column) Format(w io.Writer) {
	if e.DB != "" {
		fmt.Fprint(w, e.DB)
		fmt.Fprint(w, ".")
	}
	if e.Table != "" {
		fmt.Fprint(w, e.Table)
		fmt.Fprint(w, ".")
	}
	fmt.Fprint(w, e.Column)
}

const (
	Bit = iota
	Hex
	Digit
)

/******************
*	Number Item
*******************/

type Number struct {
	node
	Symbol bool
	Val    string
}

func (e *Number) Accept(v Visitor) (Node, bool) {

	if !v.Notify(e) {
		return v.Visit(e)
	}

	return v.Visit(e)
}

func (e *Number) Format(w io.Writer) {
	if !e.Symbol {
		fmt.Fprint(w, "-")
	}
	fmt.Fprint(w, e.Val)
}

/******************
*	String Item
*******************/

type String struct {
	node
	Str string
}

func (e *String) Accept(v Visitor) (Node, bool) {

	if !v.Notify(e) {
		return v.Visit(e)
	}

	return v.Visit(e)
}

func (e *String) Format(w io.Writer) {
	fmt.Fprint(w, e.Str)
}

/******************
*	True Item
*******************/

type True struct {
	node
	IsTrue bool
}

func (e *True) Accept(v Visitor) (Node, bool) {

	if !v.Notify(e) {
		return v.Visit(e)
	}

	return v.Visit(e)
}

func (e *True) Format(w io.Writer) {
	if e.IsTrue {
		fmt.Fprint(w, " TRUE ")
	} else {
		fmt.Fprint(w, " FALSE ")
	}
}

/******************
*	Null Item
*******************/

type Null struct {
	node
}

func (e *Null) Accept(v Visitor) (Node, bool) {

	if !v.Notify(e) {
		return v.Visit(e)
	}

	return v.Visit(e)
}

func (e *Null) Format(w io.Writer) {
	fmt.Fprint(w, " NULL ")
}

/******************
*	SystemVariable Item
*******************/

type SystemVariable struct {
	node
	Schema    string
	Parameter string
}

func (e *SystemVariable) Accept(v Visitor) (Node, bool) {

	if !v.Notify(e) {
		return v.Visit(e)
	}

	return v.Visit(e)
}

func (e *SystemVariable) Format(w io.Writer) {
	fmt.Fprint(w, e.Schema)
	if e.Parameter != "" {
		fmt.Fprint(w, ",")
		fmt.Fprint(w, e.Parameter)
	}
}

/******************
*	UserVariable Item
*******************/

type UserVariable struct {
	node

	Variable string
}

func (e *UserVariable) Accept(v Visitor) (Node, bool) {

	if !v.Notify(e) {
		return v.Visit(e)
	}

	return v.Visit(e)
}

func (e *UserVariable) Format(w io.Writer) {
	fmt.Fprint(w, e.Variable)
}

/******************
*	Marker Item
*******************/

type Marker struct {
	node

	Str string
}

func (e *Marker) Accept(v Visitor) (Node, bool) {

	if !v.Notify(e) {
		return v.Visit(e)
	}

	return v.Visit(e)
}

func (e *Marker) Format(w io.Writer) {
	fmt.Fprint(w, e.Str)
}

/******************
*	Sort Item
*******************/

const (
	Sort_Asc = iota
	Sort_Desc
	Sort_Empty
)

var ReAsc = map[int]string{
	Sort_Asc:   "ASC",
	Sort_Desc:  "DESC",
	Sort_Empty: "",
}

type Sortby struct {
	node

	Expr    ExprNode
	AscType int
}

func (e *Sortby) Accept(v Visitor) (Node, bool) {

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

func (e *Sortby) Format(w io.Writer) {
	e.Expr.Format(w)
	fmt.Fprint(w, " ")
	fmt.Fprint(w, ReAsc[e.AscType])
}
