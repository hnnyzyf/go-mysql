package preprocess

import "ast"
import "meta"

type Context struct {

	//负责Context的切换工作,初始化的时候,需要设置上一层的Context
	//对当前stmt结束遍历后，需要切换回上一个Context
	//终结条件Last==nil

	Last *Context

	//出现错误的Context是无法进行优化的
	HasError bool

	//FromClause的处理结果
	Size  int
	Table map[string]*ast.Relation
	Meta  map[*ast.Relation]meta.Meta

	//TargetCaluse的处理结果，对应Projection的情形
	Field      []*ast.Tuple
	FieldIndex map[string]int

	HasAgg bool
	Func   []int
	Agg    []int

	//WhereClause的处理结果，对应Filter的情形
	Expr      []ast.ExprNode
	ExprIndex map[ast.ExprNode]int

	//存储and或者or形式的每个部分的表达式
	And0Or    int
	And0OrExpr []ast.ExprNode

	//判断where条件是否是1 true等形式
	IsConstant bool
}

func NewContext(ctx *Context) *Context {
	return &Context{
		Last:     ctx,
		HasError: false,

		//From域
		Size:  0,
		Table: make(map[string]*ast.Relation),
		Meta:  make(map[*ast.Relation]meta.Meta),

		//Target域
		Field:      []*ast.Tuple{},
		FieldIndex: make(map[string]int),

		HasAgg: false,
		Func:   []int{},
		Agg:    []int{},

		//Where域
		Expr:      []ast.ExprNode{},
		ExprIndex: make(map[ast.ExprNode]int),

		And0Or:     0,
		And0OrExpr: []ast.ExprNode{},

		IsConstant:false,
		
	}
}



