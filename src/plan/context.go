package plan

import "ast"
import meta "util/meta"

type Context struct {

	//指向上一层的Context
	Pctx *Context
	Cctx []*Context

	//表的映射关系
	Size int

	//alias全部存储小写
	Table map[string]*ast.Relation

	//每个表的元数据信息,alias全部存储小写
	Meta map[string]*meta.Meta
	//每个表的索引信息
	IndexMeta map[string]*meta.Index
	//表有多少行
	Row map[string]int

	//Select的投影信息
	Projection *meta.Meta

	//targetclause中的Tuple信息
	Tuple    []*ast.Tuple
	TupleMap map[string][]*ast.Tuple

	//join条件中的using字符
	Using []string

	//Join栈信息
	stack

	//index中间信息
	IndexMap map[string]*condition

	//order by信息
	Sortby []*ast.Sortby
}

func NewContext() *Context {
	return &Context{
		Pctx: nil,
		Cctx: []*Context{},

		Size:      0,
		Table:     make(map[string]*ast.Relation),
		Meta:      make(map[string]*meta.Meta),
		IndexMeta: make(map[string]*meta.Index),
		Row:       make(map[string]int),

		Projection: meta.NewMeta(),
		Tuple:      []*ast.Tuple{},
		TupleMap:   make(map[string][]*ast.Tuple),

		Using: []string{},

		IndexMap: make(map[string]*condition),
		Sortby:   nil,
	}
}

func (ctx *Context) Parent(parent *Context) {
	ctx.Pctx = parent
	if parent != nil {
		parent.Cctx = append(parent.Cctx, ctx)
	}
}

func (ctx *Context) GetAllRelationAlias() []string {
	res := []string{}
	for key, _ := range ctx.Table {
		res = append(res, key)
	}
	return res
}

func (ctx *Context) GetMetaFromSubquery() *meta.Meta {
	return ctx.Projection
}

func (ctx *Context) GetTable(alias string) string {
	relation, ok := ctx.Table[alias]
	if !ok {
		return ""
	}
	return relation.Table
}

func (ctx *Context) GetRelation(alias string) *ast.Relation {
	return ctx.Table[alias]
}

func (ctx *Context) GetSqlType(alias string, column string) (meta.SqlType, bool) {

	relation, ok := ctx.Table[alias]
	if !ok {
		return meta.Type_unknown, false
	}

	if relation.GetTag() != ast.Relation_Table {
		return meta.Type_unknown, false
	}

	m, ok := ctx.Meta[alias]
	if !ok {
		return meta.Type_unknown, false
	}

	return m.GetSqlType(column), true
}

//判断是不是当前上下稳重的column
func (ctx *Context) HasColumn(table string, column string) bool {
	if meta, ok := ctx.Meta[table]; !ok {
		return false
	} else {
		for _, col := range meta.GetField() {
			if column == lower(col) {
				return true
			}
		}
		return false
	}
}

func (ctx *Context) HasRelation(table string) bool {
	if _, ok := ctx.Meta[lower(table)]; ok {
		return true
	} else {
		return false
	}
}

func (ctx *Context) CheckBelongings(column string) ([]string, bool) {
	res := []string{}

	for table, _ := range ctx.Meta {
		if ctx.HasColumn(table, column) {
			res = append(res, table)
		}
	}

	if len(res) > 0 {
		return res, true
	} else {
		return res, false
	}
}

func (ctx *Context) CheckOrderByIndex(alias string) ([]string, int, bool) {
	if len(ctx.Table) != 1 {
		return nil, 0, false
	}
	//不存在order by
	if ctx.Sortby == nil {
		return nil, 0, false
	}

	sortbyMap := make(map[string][]*ast.Sortby)

	for _, sortby := range ctx.Sortby {
		if sortby.Expr.Type() != ast.Ast_column {
			return nil, 0, false
		} else {
			col := sortby.Expr.(*ast.Column)
			table := col.GetTable()
			if _, ok := sortbyMap[table]; ok {
				sortbyMap[table] = append(sortbyMap[table], sortby)
			} else {
				sortbyMap[table] = []*ast.Sortby{sortby}
			}
		}
	}

	//存在多个条件
	if len(sortbyMap) > 1 {
		return nil, 0, false
	}

	//不属于当前条件
	if _, ok := sortbyMap[alias]; !ok {
		return nil, 0, false
	}

	//查看投影是否只包含主键和order By条件
	asctype := 0
	res := []string{}
	for index, sortby := range sortbyMap[alias] {
		col := sortby.Expr.(*ast.Column).GetColumn()
		if index == 0 {
			asctype = sortby.AscType
			res = append(res, col)
			continue
		} else {
			if sortby.AscType != asctype {
				return nil, 0, false
			}
			res = append(res, col)
		}

	}

	return res, asctype, true
}

func (ctx *Context) GetPrimaryKey(alias string) []string {
	if indexs, ok := ctx.IndexMeta[alias]; ok {
		if indexs.HasPrimary {
			return indexs.Idx["PRIMARY"]
		}
	}

	return nil
}

func (ctx *Context) GetRow(alias string, sample int) (int, int) {
	row := ctx.Row[alias]
	if row > sample {
		row = sample
	}

	return row, ctx.Row[alias]
}

func (ctx *Context) HasPrimary(equalIndexs [][]string, alias string) ([]string, bool) {
	primaryKey := ctx.GetPrimaryKey(alias)
	if primaryKey == nil {
		return nil, false
	}

	for _, index := range equalIndexs {
		if isSameString(index, primaryKey) {
			return index, true
		}
	}

	return nil, false
}

func (ctx *Context) HasUsing(column string) bool {
	for _, col := range ctx.Using {
		if col == column {
			return true
		}
	}
	return false
}

func (ctx *Context) GetIndexMeta(alias string) (*meta.Index, bool, *Context) {
	meta, ok := ctx.IndexMeta[alias]
	if ok {
		return meta, true, ctx
	}

	if ctx.Pctx == nil {
		return nil, false, ctx
	}

	meta, ok = ctx.Pctx.IndexMeta[alias]

	if ok {
		return meta, true, ctx.Pctx
	}

	return nil, false, ctx.Pctx
}
