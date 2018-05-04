package plan

import "github.com/hnnyzyf/mysql-go/ast"
import "github.com/hnnyzyf/mysql-go/util/server"
import "errors"
import "fmt"

type Plan struct {

	//原始的AST解析树
	ParseTree ast.Node
	//存储每一个SimpleSelectStmt的上下文
	Ctx map[*ast.SimpleSelectStmt]*Context

	HasError bool

	Error []string

	MetaServer *server.Mysql
}

func NewPlan(tree ast.Node, server *server.Mysql) *Plan {
	return &Plan{
		ParseTree:  tree,
		Ctx:        make(map[*ast.SimpleSelectStmt]*Context),
		HasError:   false,
		Error:      []string{},
		MetaServer: server,
	}
}

func (p *Plan) AddError(e []string) {
	p.Error = append(p.Error, e...)
}

var PreprocessError = errors.New("预处理出现错误")

//预处理
//为每一个Relation添加别名
//为每一个非*tuple添加别名
//为每一个非subquery的relation添加meta信息
//扩展*
//生成Subquery的Meta信息
//检查subquery是否存在多个projection
func (p *Plan) PreProcessMeta() error {
	v := NewPreProcessMetaVisitor(p.MetaServer)
	p.ParseTree.Accept(v)
	p.Ctx = v.Ctx

	p.AddError(v.Error)
	p.HasError = v.HasError

	if p.HasError {
		return PreprocessError
	} else {
		return nil
	}
}

var TargetClauseError = errors.New("TargetClause添加归属出现错误")
var FromClauseError = errors.New("FromClause添加归属出现错误")
var WhereClauseError = errors.New("WhereClause添加归属出现错误")

//处理target,from,where,order by,having中的ast.Column
//添加对应的关系信息
func (p *Plan) PreProcessColumn() error {

	fv := NewFromClauseVisitor(p.Ctx)
	for stmt, _ := range p.Ctx {
		stmt.Accept(fv)
	}

	p.AddError(fv.Error)
	p.HasError = fv.HasError

	if p.HasError {
		return FromClauseError
	}

	//为targetclause所有的column添加Table信息
	tv := NewTargetClauseVisitor(p.Ctx)

	for stmt, _ := range p.Ctx {
		stmt.Accept(tv)
	}

	p.AddError(tv.Error)
	p.HasError = tv.HasError

	if p.HasError {
		return TargetClauseError
	}

	wv := NewWhereClauseVisitor(p.Ctx)
	for stmt, _ := range p.Ctx {
		stmt.Accept(wv)
	}

	if p.HasError {
		return WhereClauseError
	}

	return nil
}

var constantFoldError = errors.New("constantFold出现错误")

//处理target from where中的condition
//对常数条件进行折叠
//  对于一个column,保证对应的条件只会出现如下两种情况
//	1.等值条件
//	2.range条件
func (p *Plan) ConstantFolding() error {
	v := NewConstantVisitor(p.MetaServer, p.Ctx)
	p.ParseTree.Accept(v)

	p.AddError(v.Error)
	p.HasError = v.HasError

	if p.HasError {
		return constantFoldError
	}

	return nil
}

//将如下情况加入到候选索引列表中去（最左匹配原则）
// 		1.等值条件 		单一索引
// 		2.等值+等值...  复合索引
// 		3.等值+比较...  复合索引的部分选择

var IndexError = errors.New("查询INDEX的过程中出现错误")

func (p *Plan) ProcessIndex() (error, []string, map[string][]string) {
	v := NewIndexVisitor(p.Ctx, p.MetaServer)
	count := 1
	for stmt, _ := range p.Ctx {
		v.addInfo(fmt.Sprintln("	", count, ". 当前校验的的Select语句是:", generateString(stmt)))
		stmt.Accept(v)
		count = count + 1
		v.addInfo(fmt.Sprintln("	"))
	}
	p.AddError(v.Error)
	p.HasError = v.HasError

	if v.HasError {
		return IndexError, nil, nil
	}

	return nil, v.Info, v.indexs

}
