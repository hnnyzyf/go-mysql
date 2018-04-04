package ast

import "io"
import server "util/server"

//定义一个Node，定义vistor访问方式，该Node可以接受vistor，vistor会遍历这个节点
type Node interface {
	//每一个Node都需要能接受一个Visitor进行访问
	//返回的Node是由Vistor来定义的
	//bool表示是否要继续向下访问，用于控制子节点的遍历
	Accept(Visitor) (Node, bool)

	//实现格式化
	Format(io.Writer)

	//判断节点是否是相似的节点
	IsSameAs(ExprNode, Visitor) bool

	//获得当前节点的类型
	Type() int
}

//实现walk为vistor，该vistor对当前的Node进行处理，给出我们想要的操作
type Visitor interface {
	//判断Node是否需要访问
	//使用switch判断节点是否需要继续遍历自己的子节点
	Notify(Node) bool
	//处理节点
	//返回处理后的新的节点以及是否继续往下运行
	Visit(Node) (Node, bool)

	//返回一个数据库服务,原信息需要从数据库服务中获得
	Session() server.Server
}

//定义一个表达式类型的接口
type ExprNode interface {
	//实现Node接口
	Node

	//由struct node实现
	Setnode
	Getnode
}

type Setnode interface {
	SetTag(int)
	SetParent(ExprNode)
}

type Getnode interface {
	GetTag() int
	GetParent() ExprNode
}

const (
	Ast_distinctclause = iota + 1
	Ast_targetclause
	Ast_intoclause
	Ast_fromclause
	Ast_whereclause
	Ast_groupclause
	Ast_havingclause
	Ast_sortclause
	Ast_lockclause
	Ast_limitclause

	Ast_expr
	Ast_and0orexpr
	Ast_unaryexpr
	Ast_notexpr
	Ast_istrueexpr
	Ast_isnullexpr
	Ast_subqueryexpr
	Ast_betweenexpr
	Ast_likeexpr
	Ast_regexpexpr
	Ast_intervalexpr
	Ast_collateexpr
	Ast_caseexpr
	Ast_inexpr

	Ast_funccall
	Ast_aggregationfunccall
	Ast_castfunccall
	Ast_caltimefunccall
	Ast_stringfunccall

	Ast_relation
	Ast_tuple
	Ast_column
	Ast_star
	Ast_integer
	Ast_float
	Ast_string
	Ast_true
	Ast_null
	Ast_systemvariable
	Ast_uservariable
	Ast_marker
	Ast_sortby

	Ast_join
	Ast_using

	Ast_simpleselectstmt
	Ast_unionstmt

	Ast_Unknown
)

var TypeString = map[int]string{
	Ast_integer: "Integer",
	Ast_float:   "Float",
	Ast_string:  "String",
}
