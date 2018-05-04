package ast

import "io"
import "fmt"
import "bytes"
import "strings"
import "github.com/hnnyzyf/mysql-go/util/comp"
import "github.com/hnnyzyf/mysql-go/util/conv"

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
********************************I****/

/******************
*	Relation Item
*******************/

const (
	Relation_Table = iota + 1
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

func (e *Relation) Type() int {
	return Ast_relation
}

func (e *Relation) IsSameAs(src ExprNode, v Visitor) bool {
	if src.Type() != e.Type() {
		return false
	}

	if e.GetTag() != src.GetTag() {
		return false
	}

	other := src.(*Relation)
	switch e.GetTag() {
	case Relation_Table:
		if e.GetDB() != other.GetDB() {
			return false
		}

		if e.GetTable() != other.GetTable() {
			return false
		}
	case Relation_Subquery:
		if !e.Select.(ExprNode).IsSameAs(other.Select.(ExprNode), v) {
			return false
		}
	}

	return true
}

func (e *Relation) GetDB() string {
	return conv.Removebacktick(e.DB)
}

func (e *Relation) GetTable() string {
	return conv.Removebacktick(e.Table)
}

func (e *Relation) GetAlias() string {
	if e.Alias != "" {
		return strings.ToLower(conv.Removebacktick(e.Alias))
	} else {
		switch e.GetTag() {
		case Relation_Table:
			return strings.ToLower(e.GetTable())
		default:
			buf := bytes.NewBuffer([]byte{})
			e.Format(buf)
			return buf.String()
		}
	}
}

/******************
*	Tuple Item
*******************/

const (
	Tuple_expr = iota + 1
	Tuple_column
	Tuple_funcall
	Tuple_agg
	Tuple_star
	Tuple_SimpleSubquery
	Tuple_UnionSubquery
)

type Tuple struct {
	node

	Ref ExprNode
	//关系指向的别名

	Alias string
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
		fmt.Fprint(w, e.Alias)
	}
}

func (e *Tuple) Type() int {
	return Ast_tuple
}

func (e *Tuple) IsSameAs(src ExprNode, v Visitor) bool {
	if src.Type() != e.Type() {
		return false
	}

	if e.GetTag() != src.GetTag() {
		return false
	}

	other := src.(*Tuple)

	if !e.Ref.IsSameAs(other.Ref, v) {
		return false
	}

	return true
}

func (e *Tuple) GetAlias() string {
	if e.Alias != "" {
		return strings.ToLower(e.Alias)
	} else {
		switch e.GetTag() {
		case Tuple_column:
			return strings.ToLower(e.Ref.(*Column).Column)
		default:
			buf := bytes.NewBuffer([]byte{})
			e.Format(buf)
			return buf.String()
		}
	}
}

/******************
*	Column Item
*******************/

type Column struct {
	node

	DB     string
	Table  string
	Column string

	//column所属的Relation
	Relation *Relation
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

func (e *Column) Type() int {
	return Ast_column
}

func (e *Column) GetDB() string {
	return conv.Removebacktick(e.DB)
}

func (e *Column) GetTable() string {
	return conv.Removebacktick(e.Table)
}

func (e *Column) GetColumn() string {
	return conv.Removebacktick(e.Column)
}

func (e *Column) IsSameAs(src ExprNode, v Visitor) bool {
	if src.Type() != e.Type() {
		return false
	}

	if e.GetTag() != src.GetTag() {
		return false
	}

	other := src.(*Column)
	if e.GetDB() != other.GetDB() {
		return false
	}

	if e.GetTable() != other.GetTable() {
		return false
	}

	if e.GetColumn() != other.GetColumn() {
		return false
	}

	if e.Relation != other.Relation {
		return false
	}

	return true
}

type Star struct {
	node

	DB    string
	Table string

	Column []*Column
}

func (e *Star) Accept(v Visitor) (Node, bool) {

	if !v.Notify(e) {
		return v.Visit(e)
	}

	if e.Column != nil {
		for index, table := range e.Column {
			node, ok := table.Accept(v)
			if !ok {
				return e, false
			}
			e.Column[index] = node.(*Column)

		}
	}

	return v.Visit(e)
}

func (e *Star) Format(w io.Writer) {
	if e.DB != "" {
		fmt.Fprint(w, e.DB)
		fmt.Fprint(w, ".")
	}
	if e.Table != "" {
		fmt.Fprint(w, e.Table)
		fmt.Fprint(w, ".")
	}
	fmt.Fprint(w, "*")
}

func (e *Star) Type() int {
	return Ast_star
}

func (e *Star) GetDB() string {
	return conv.Removebacktick(e.DB)
}

func (e *Star) GetTable() string {
	return conv.Removebacktick(e.Table)
}

func (e *Star) IsSameAs(src ExprNode, v Visitor) bool {
	if src.Type() != e.Type() {
		return false
	}

	if e.GetTag() != src.GetTag() {
		return false
	}

	other := src.(*Star)
	if e.GetDB() != other.GetDB() {
		return false
	}

	if e.GetTable() != other.GetTable() {
		return false
	}

	if len(e.Column) != len(other.Column) {
		return false
	}

	for _, expr := range e.Column {
		hasSame := false
		for _, otherexpr := range other.Column {
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

const (
	Bit = iota + 1
	Hex
	Int
)

/******************
*	Number Item
*******************/

type Integer struct {
	node
	Operator []string
	Val      string
}

func (e *Integer) Accept(v Visitor) (Node, bool) {

	if !v.Notify(e) {
		return v.Visit(e)
	}

	return v.Visit(e)
}

func (e *Integer) Format(w io.Writer) {
	for index := len(e.Operator) - 1; index >= 0; index-- {
		fmt.Fprint(w, e.Operator[index])
	}
	fmt.Fprint(w, e.Val)
}

func (e *Integer) Type() int {
	return Ast_integer
}

func (e *Integer) IsSameAs(src ExprNode, v Visitor) bool {
	if src.Type() != e.Type() {
		return false
	}

	if e.GetTag() != src.GetTag() {
		return false
	}

	other := src.(*Integer)

	buf := bytes.NewBuffer([]byte{})
	e.Format(buf)
	buf.WriteString(" = ")
	other.Format(buf)

	if !comp.Compare(buf.String(), v.Session()) {
		return false
	}

	return true
}

type Float struct {
	node
	Operator []string
	Integer  string
	Decimal  string
}

func (e *Float) Accept(v Visitor) (Node, bool) {

	if !v.Notify(e) {
		return v.Visit(e)
	}

	return v.Visit(e)
}

func (e *Float) Format(w io.Writer) {
	for index := len(e.Operator) - 1; index >= 0; index-- {
		fmt.Fprint(w, e.Operator[index])
	}

	fmt.Fprint(w, e.Integer)
	fmt.Fprint(w, ".")
	fmt.Fprint(w, e.Decimal)
}

func (e *Float) Type() int {
	return Ast_float
}

func (e *Float) IsSameAs(src ExprNode, v Visitor) bool {
	if src.Type() != e.Type() {
		return false
	}

	if e.GetTag() != src.GetTag() {
		return false
	}

	other := src.(*Float)

	buf := bytes.NewBuffer([]byte{})
	e.Format(buf)
	buf.WriteString(" = ")
	other.Format(buf)

	if !comp.Compare(buf.String(), v.Session()) {
		return false
	}

	return true
}

/******************
*	String Item
*******************/
const (
	String_single = iota + 1
	String_double
	String_math
)

type String struct {
	node
	Operator []string
	Str      string
}

func (e *String) Accept(v Visitor) (Node, bool) {

	if !v.Notify(e) {
		return v.Visit(e)
	}

	return v.Visit(e)
}

func (e *String) Format(w io.Writer) {
	for index := len(e.Operator) - 1; index >= 0; index-- {
		fmt.Fprint(w, e.Operator[index])
	}

	switch e.GetTag() {
	case String_single:
		fmt.Fprint(w, "'")
	case String_double:
		fmt.Fprint(w, "\"")
	default:
		//fmt.Fprint(w, "(")

	}

	fmt.Fprint(w, e.Str)
	switch e.GetTag() {
	case String_single:
		fmt.Fprint(w, "'")
	case String_double:
		fmt.Fprint(w, "\"")
	default:
		//fmt.Fprint(w, ")")

	}
}

func (e *String) Type() int {
	return Ast_string
}

func (e *String) IsSameAs(src ExprNode, v Visitor) bool {
	if src.Type() != e.Type() {
		return false
	}

	if e.GetTag() != src.GetTag() {
		return false
	}

	other := src.(*String)
	buf := bytes.NewBuffer([]byte{})
	e.Format(buf)
	buf.WriteString(" = ")
	other.Format(buf)

	if !comp.Compare(buf.String(), v.Session()) {
		return false
	}

	return true
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

func (e *True) Type() int {
	return Ast_true
}

func (e *True) IsSameAs(src ExprNode, v Visitor) bool {
	if src.Type() != e.Type() {
		return false
	}

	if e.GetTag() != src.GetTag() {
		return false
	}

	other := src.(*True)

	if e.IsTrue != other.IsTrue {
		return false
	}

	return true
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

func (e *Null) Type() int {
	return Ast_null
}

func (e *Null) IsSameAs(src ExprNode, v Visitor) bool {
	if src.Type() != e.Type() {
		return false
	}

	if e.GetTag() != src.GetTag() {
		return false
	}

	return true
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

func (e *SystemVariable) Type() int {
	return Ast_systemvariable
}

func (e *SystemVariable) IsSameAs(src ExprNode, v Visitor) bool {
	if src.Type() != e.Type() {
		return false
	}

	if e.GetTag() != src.GetTag() {
		return false
	}

	other := src.(*SystemVariable)
	if e.Schema != other.Schema {
		return false
	}

	if e.Parameter != other.Parameter {
		return false
	}

	return true
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

func (e *UserVariable) Type() int {
	return Ast_uservariable
}

func (e *UserVariable) IsSameAs(src ExprNode, v Visitor) bool {
	if src.Type() != e.Type() {
		return false
	}

	if e.GetTag() != src.GetTag() {
		return false
	}

	other := src.(*UserVariable)

	if e.Variable != other.Variable {
		return false
	}

	return true
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

func (e *Marker) Type() int {
	return Ast_marker
}

func (e *Marker) IsSameAs(src ExprNode, v Visitor) bool {
	if src.Type() != e.Type() {
		return false
	}

	if e.GetTag() != src.GetTag() {
		return false
	}

	return true
}

/******************
*	Sort Item
*******************/

const (
	Sort_Asc = iota + 1
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

func (e *Sortby) Type() int {
	return Ast_sortby
}

func (e *Sortby) IsSameAs(src ExprNode, v Visitor) bool {
	if src.Type() != e.Type() {
		return false
	}

	if e.GetTag() != src.GetTag() {
		return false
	}

	other := src.(*Sortby)

	if e.AscType != other.AscType {
		return false
	}

	if !e.Expr.IsSameAs(other.Expr, v) {
		return false
	}
	return true
}
