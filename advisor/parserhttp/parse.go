package parserhttp

import p "github.com/hnnyzyf/mysql-go/advisor/plan"
import "github.com/hnnyzyf/mysql-go/advisor/util/server"
import "github.com/hnnyzyf/mysql-go/advisor/parser"
import "bytes"

//经过sql解析获得的结果
type parseResult struct {
	Sql   string              `json:"sql"`
	Index map[string][]string `json:"index"`
	Info  []string            `json:info`
	Error []string            `json:error`
}

func newparseResult() *parseResult {
	return &parseResult{
		Info:  []string{},
		Error: []string{},
	}
}

func parseStmt(ask *request) (*parseResult, error) {
	//创建数据库连接
	mysql := server.NewMysqlServer(User, PassWord, ask.Host, ask.Port, ask.DB)
	sql := ask.Sql
	res := newparseResult()

	//解析数据库语句
	ast, err := parser.SqlParse(sql)
	if err != nil {
		return res, err
	}

	//语句预处理
	plan := p.NewPlan(ast, mysql)
	err = plan.PreProcessMeta()
	if err != nil {
		res.Error = plan.Error
		return res, err
	}

	//处理Column
	err = plan.PreProcessColumn()
	if err != nil {
		res.Error = plan.Error
		return res, err
	}

	//处理constantfold
	err = plan.ConstantFolding()

	if err != nil {
		return res, err
	}

	buf := bytes.NewBuffer([]byte{})
	ast.Format(buf)

	res.Sql = buf.String()

	//输出索引数据
	ok, info, index := plan.ProcessIndex()

	if ok != nil {
		return res, err
	}

	res.Info = info
	res.Index = index
	res.Error = plan.Error

	return res, nil
}
