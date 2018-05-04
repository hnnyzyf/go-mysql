package plan

import "github.com/hnnyzyf/mysql-go/advisor/ast"

const (
	a_equal = iota + 1
	a_not_equal
	a_range
	a_like
)

//定义index列表
type condition struct {
	//等值条件
	keys []string

	//不等值条件
	nkeys []string

	//范围条件
	ranges []string

	value map[string][]ast.ExprNode
}

func newCondition() *condition {
	return &condition{
		keys:   []string{},
		nkeys:  []string{},
		ranges: []string{},
		value:  make(map[string][]ast.ExprNode),
	}
}

func (i *condition) add(column string, op int, val ast.ExprNode) {

	//添加条件列
	switch op {
	case a_equal:
		i.keys = append(i.keys, column)
	case a_not_equal:
		i.nkeys = append(i.nkeys, column)
	default:
		if _, ok := i.value[column]; !ok {
			i.ranges = append(i.ranges, column)
		}
	}

	//添加具体条件
	if _, ok := i.value[column]; ok {
		i.value[column] = append(i.value[column], val)
	} else {
		i.value[column] = []ast.ExprNode{val}
	}

}

func (i *condition) candidateIndex(columns []string, start int, end int) [][]string {
	res := [][]string{}
	if end > len(columns) {
		end = len(columns)
	}
	for id := start; id <= end; id++ {
		res = append(res, combination(columns, id)...)
	}
	return res
}

func (i *condition) keySize() int {
	return len(i.keys)
}

func (i *condition) rangeSize() int {
	return len(i.ranges)
}

func (i *condition) getCondition(index []string) []ast.ExprNode {
	res := []ast.ExprNode{}
	for _, col := range index {
		res = append(res, i.value[col]...)
	}
	return res
}

//求解组合
func combination(src []string, bucket int) [][]string {
	res := [][]string{}

	if bucket > len(src) {
		return res
	}

	if bucket == 0 {
		bucket = len(src)
	}

	if bucket == 1 {
		for _, column := range src {
			res = append(res, []string{column})
		}
		return res
	}

	for index, column := range src {
		subset := src[index+1:]
		c := combination(subset, bucket-1)
		for idx, _ := range c {
			c[idx] = append(c[idx], column)
		}
		res = append(res, c...)
	}

	return res
}

//求解组合
func combineBtreeIndex(org [][]string, src [][]string) [][]string {
	res := [][]string{}
	for _, oindex := range org {
		for _, sindex := range src {
			res = append(res, (append(oindex, sindex...)))
		}
	}
	return res
}
