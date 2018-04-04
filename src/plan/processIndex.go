package plan

import "ast"
import server "util/server"
import "fmt"
import "sort"
import "math"

/*******************************

	按照Relation的顺序,遍历Where条件
	获得候选索引的列表

********************************/

// A data structure to hold a key/value pair.
type Pair struct {
	Key   int
	Value float32
	Len   int
}

// A slice of Pairs that implements sort.Interface to sort by Value.
type PairList []Pair

func (p PairList) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

func (p PairList) Len() int { return len(p) }

func (p PairList) Less(i, j int) bool {
	if p[i].Value > p[j].Value {
		return true
	}

	if p[i].Value == p[j].Value && p[i].Len < p[j].Len {
		return true
	}

	return false
}

// A function to turn a map into a PairList, then sort and return it.
func sortMapByValue(m map[int]float32, index [][]string) PairList {
	p := make(PairList, len(m))
	i := 0
	for k, v := range m {
		p[i] = Pair{k, v, len(index[k])}
		i++
	}
	sort.Sort(p)
	return p
}

const SAMPLE = 500000

type IndexVisitor struct {
	//上下文信息
	Ctx map[*ast.SimpleSelectStmt]*Context

	//当前正在遍历的select
	current *Context

	//数据库服务信息
	MetaServer server.Server

	Info []string

	HasError bool
	Error    []string

	//可以建立的索引
	indexs map[string][]string
}

func NewIndexVisitor(ctx map[*ast.SimpleSelectStmt]*Context, s server.Server) *IndexVisitor {
	return &IndexVisitor{
		Ctx:        ctx,
		current:    nil,
		MetaServer: s,
		Info:       []string{},
		HasError:   false,
		Error:      []string{},
		indexs:     make(map[string][]string),
	}
}

func (pre *IndexVisitor) Notify(n ast.Node) bool {
	switch node := n.(type) {
	case *ast.SimpleSelectStmt:
		pre.current = pre.Ctx[node]
		return true
	case *ast.TargetClause:
		return false
	case *ast.FromClause:
		return true
	//不遍历SubqueryExpr
	case *ast.SubqueryExpr:
		return false
	//不对relation进行遍历
	case *ast.Relation:
		return false
	default:
		return true
	}

	return true
}

func (pre *IndexVisitor) Visit(n ast.Node) (ast.Node, bool) {
	switch node := n.(type) {
	//遍历完whereclause后，则停止遍历
	case *ast.SimpleSelectStmt:
		pre.current = pre.current.Pctx
		return node, true
	//除了whereClause和FromClause，其他clause均不遍历
	case *ast.Join:
		if node.Condition != nil {
			pre.CollectConditionIndex(node)
		}
		return node, true
	case *ast.WhereClause:
		pre.CollectCandidateIndex(node)
		return node, true
	default:
		return node, true
	}
}

func (pre *IndexVisitor) Session() server.Server {
	return pre.MetaServer
}

func (pre *IndexVisitor) addInfo(info string) {
	pre.Info = append(pre.Info, info)
}

//根据每一个relation条件生成
func (pre *IndexVisitor) CollectConditionIndex(node *ast.Join) {
	switch condition := node.Condition.(type) {
	case *ast.Using:
		pre.FindUsingIndex(condition)
	default:
		pre.FindEuqalConditionIndex(node.Condition)
	}

}

//根据每一个relation条件生成
func (pre *IndexVisitor) CollectCandidateIndex(node *ast.WhereClause) {
	ctx := pre.current
	if node.GetTag() == ast.AST_EMPTY {
		pre.addInfo(fmt.Sprintln("	[Where条件索引校验]没有Where条件,检查Order或者Group条件"))
		for alias, _ := range ctx.Table {
			//只处理relation_table的情况
			if isRelationTable(alias, ctx) {
				//根据relation,生成索引列表
				pre.ExtractIndexFromOrderBy(alias)
			}
		}
		return
	}

	//检查连接条件
	pre.addInfo(fmt.Sprintln("	[Join条件索引校验]检查Where中的连接条件"))
	pre.FindEuqalConditionIndex(node.Where)

	for alias, _ := range ctx.Table {
		//只处理relation_table的情况
		if isRelationTable(alias, ctx) {
			//根据relation,生成索引列表
			pre.ExtractIndexFromWhereClause(alias, node.Where)
		}
	}
}

func (pre *IndexVisitor) GetIndex(alias string) {
	ctx := pre.current
	currentIndex := ctx.IndexMeta[alias]
	conditionStack := ctx.IndexMap[alias]
	_, row := ctx.GetRow(alias, SAMPLE)

	pre.addInfo(fmt.Sprintln("**********[关系", alias, "(总行数:", row, ")]索引校验开始**********"))

	//getEqualIndex
	equalIndexs, select1 := pre.getEqualIndex(conditionStack, alias)
	rangeIndexs, select2 := pre.getRangeIndex(conditionStack, alias)
	leftIndexs, select3 := pre.getLeftMostPrefixIndex(conditionStack, alias)

	//主键是第一优先级
	if equalIndexs != nil {
		//判断主键
		if index, ok := ctx.HasPrimary(equalIndexs, alias); ok {
			pre.addInfo(fmt.Sprintln("	1.[主键条件校验] 索引", index, "是主键,无需添加索引"))
			return
		}
	}

	if rangeIndexs != nil {
		result := sortMapByValue(select2, rangeIndexs)
		for rank, cardinality := range result {
			percent := cardinality.Value
			scanRow := int((100 - cardinality.Value) * float32(row) / 100)
			pre.addInfo(fmt.Sprintln("	2.[Range条件校验]排名第", rank+1, "的索引为", rangeIndexs[cardinality.Key], ",过滤掉", percent, "%的数据,扫描的行数预估为", scanRow))
			if percent >= 95 {
				index := rangeIndexs[cardinality.Key]
				indexQuery := generateIndex(ctx.Table[alias].GetTable(), index)
				lable := "Range索引"
				if currentIndex.HaveSameIndex(index) {
					lable = "Range索引(已存在)"
				}
				if _, ok := pre.indexs["Range索引"]; ok {
					pre.indexs[lable] = append(pre.indexs[lable], indexQuery)
				} else {
					pre.indexs[lable] = []string{indexQuery}
				}
			}
		}

	}

	//最左匹配索引是第二优先级
	if leftIndexs != nil {
		if equalIndexs == nil {
			equalIndexs = [][]string{}
			select1 = make(map[int]float32)
		}
		//优先选择等值条件和最左匹配条件
		candidateIndex := append(equalIndexs, leftIndexs...)
		for idx, cardinality := range select3 {
			select1[len(equalIndexs)+idx] = cardinality
		}

		result := sortMapByValue(select1, candidateIndex)

		for rank, cardinality := range result {
			percent := (1 - 1/cardinality.Value) * 100
			scanRow := float32(row) / cardinality.Value
			pre.addInfo(fmt.Sprintln("	2.[最左匹配索引校验]选择度排名第", rank+1, "的索引为", candidateIndex[cardinality.Key], ",其选择度为", cardinality.Value, "/", SAMPLE, ",可过滤掉:", percent, "%数据,扫描的行数预估为", scanRow))
			//输出前三的索引建立语句
			if rank <= 2 {
				index := candidateIndex[cardinality.Key]
				indexQuery := generateIndex(ctx.Table[alias].GetTable(), candidateIndex[cardinality.Key])
				lable := "联合索引"
				if currentIndex.HaveSameIndex(index) {
					lable = "联合索引(已存在)"
				}
				if _, ok := pre.indexs["联合索引"]; ok {
					pre.indexs[lable] = append(pre.indexs[lable], indexQuery)
				} else {
					pre.indexs[lable] = []string{indexQuery}
				}
			}
		}
	}

}

//获得潜在的可能的索引列表
func (pre *IndexVisitor) getEqualIndex(conditionStack *condition, alias string) ([][]string, map[int]float32) {
	ctx := pre.current
	sample, _ := ctx.GetRow(alias, SAMPLE)
	//获得所有的等值条件
	equals := conditionStack.keys

	if len(equals) == 0 {
		pre.addInfo(fmt.Sprintln("	[等值索引校验]没有等值条件,略过"))
		return nil, nil
	}

	//获得候选索引的组合列表
	candidateIndex := conditionStack.candidateIndex(equals, 1, 1)

	//计算每一个索引的选择度
	//即每一个index条件的需要扫描的行数
	selectivity := make(map[int]float32)
	for idx, index := range candidateIndex {
		selectivity[idx] = float32(pre.GetCardinality(sample, alias, index))
		pre.addInfo(fmt.Sprintln("	[等值索引校验]索引:", index, " 的选择度为", selectivity[idx]))

	}

	return candidateIndex, selectivity
}

//获得潜在的可能的索引列表
func (pre *IndexVisitor) getRangeIndex(conditionStack *condition, alias string) ([][]string, map[int]float32) {
	ctx := pre.current
	_, row := ctx.GetRow(alias, SAMPLE)

	//获得所有的range条件
	ranges := conditionStack.ranges

	if len(ranges) == 0 {
		pre.addInfo(fmt.Sprintln("	[Range索引校验]没有Range条件,略过"))
		return nil, nil
	}

	//获得候选索引的组合列表
	candidateIndex := conditionStack.candidateIndex(ranges, 1, 1)

	//计算每一个索引的筛选了多少行
	selectivity := make(map[int]float32)

	for idx, index := range candidateIndex {
		expr := conditionStack.getCondition(index)
		filter:=float64(pre.GetFilter(alias, index, expr))


		selectivity[idx] = (float32)(1 - math.Min(filter,float64(row))/float64(row)) * 100
		pre.addInfo(fmt.Sprintln("	[Range索引校验]索引:", index, " 过滤掉", selectivity[idx], "%的数据"))
	}

	return candidateIndex, selectivity
}

//获得潜在的可能的索引列表
func (pre *IndexVisitor) getLeftMostPrefixIndex(conditionStack *condition, alias string) ([][]string, map[int]float32) {
	ctx := pre.current
	sample, _ := ctx.GetRow(alias, SAMPLE)

	//获得所有条件
	equals := append(conditionStack.keys, conditionStack.nkeys...)
	ranges := conditionStack.ranges

	if len(equals) == 0 {
		return nil, nil
	}

	//获得hash条件的组合
	hashIndex := conditionStack.candidateIndex(equals, 2, len(equals))

	//获得btree条件的组合
	rangesIndex := conditionStack.candidateIndex(ranges, 1, 1)
	btreeIndex := combineBtreeIndex(hashIndex, rangesIndex)

	candidateIndex := append(hashIndex, btreeIndex...)

	//计算每一个索引的选择度
	selectivity := make(map[int]float32)
	for idx, index := range candidateIndex {
		selectivity[idx] = float32(pre.GetCardinality(sample, alias, index))
		pre.addInfo(fmt.Sprintln("	[联合索引校验]索引:", index, " 的选择度为", selectivity[idx]))
	}

	return candidateIndex, selectivity
}

//获得潜在的可能的索引列表,只处理单表的情况
func (pre *IndexVisitor) getOrderbyIndex(conditionStack *condition, alias string, orderColumns []string, asc int) ([][]string, map[int]float32) {
	//ctx := pre.current
	//_, row := ctx.GetRow(alias, SAMPLE)

	//检测投影中是否包含非order by 或者 primary的column
	if !pre.checkProjectionHasPrimaryAndOrderBy(alias, orderColumns) {
		pre.addInfo(fmt.Sprintln("	[Order条件检查],投影含有额外的Column,不能使用order中的条件"))
		return nil, nil
	}

	if conditionStack == nil {
		return [][]string{orderColumns}, nil
	}

	equals := conditionStack.keys
	nequals := conditionStack.nkeys
	ranges := conditionStack.ranges

	candidateIndex := []string{}

	//select * from t1 order by key_part1,key_part2
	if len(equals) == 0 && len(nequals) == 0 && len(ranges) == 0 {
		candidateIndex = append(candidateIndex, orderColumns...)
	} else if len(equals) != 0 && len(nequals) == 0 && len(ranges) == 0 {
		//select * from t1 where key_part1=2 order by key_part1,key_part2
		if hasColumn(equals, orderColumns[0]) {
		}
	}

	return nil, nil

}

type columnVisitor struct {
	alias  string
	column []string
	//是否
	IsTrue bool
}

func (pre *columnVisitor) Notify(n ast.Node) bool {
	switch n.(type) {
	case *ast.SubqueryExpr:
		return false
	default:
		return true
	}
}

func (pre *columnVisitor) Visit(n ast.Node) (ast.Node, bool) {
	switch node := n.(type) {
	case *ast.Column:
		pre.checkcolumn(node)
	default:
		return n, true

	}
	return n, true
}

func (pre *columnVisitor) Session() server.Server {
	return nil
}

func (pre *columnVisitor) checkcolumn(node *ast.Column) {
	if node.GetTable() != pre.alias {
		pre.IsTrue = false
		return
	}
	column := node.GetColumn()

	for _, col := range pre.column {
		if col == column {
			return
		}
	}
	pre.IsTrue = false
}

func (pre *IndexVisitor) checkProjectionHasPrimaryAndOrderBy(alias string, sortby []string) bool {
	ctx := pre.current
	primaryKey := ctx.GetPrimaryKey(alias)
	v := &columnVisitor{alias: alias, column: append(primaryKey, sortby...), IsTrue: true}
	for _, tuple := range ctx.Tuple {
		tuple.Accept(v)
		if !v.IsTrue {
			return false
		}
	}
	return true
}

//计算指定索引列表的选择度
func (pre *IndexVisitor) GetCardinality(row int, table string, idx []string) int {
	ctx := pre.current
	query := generateSelectivityQuery(row, generateString(ctx.Table[table]), idx)
	pre.addInfo(fmt.Sprintln("	使用", query, "计算索引 ", idx, "的选择度"))
	session := pre.MetaServer.CreateSession2DB("")
	cardinality, err := pre.MetaServer.QueryCardinality(session, query)
	if err != nil {
		pre.HasError = true
		pre.Error = append(pre.Error, err.Error())
		return 1
	}
	return cardinality
}

//计算指定索引列表的选择度
func (pre *IndexVisitor) GetFilter(table string, idx []string, expr []ast.ExprNode) int {
	ctx := pre.current
	query := generateRangeQuery(generateString(ctx.Table[table]), expr)
	pre.addInfo(fmt.Sprintln("	使用", query, "计算,", idx, "的选择度"))
	session := pre.MetaServer.CreateSession2DB("")
	cardinality, err := pre.MetaServer.QueryCardinality(session, query)
	if err != nil {
		pre.HasError = true
		pre.Error = append(pre.Error, err.Error())
		return 1
	}
	return cardinality
}
