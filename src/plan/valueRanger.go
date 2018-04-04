package plan

import comp "util/compare"
import "ast"
import server "util/server"

//一个属性的取值范围
type ranger struct {
	has_low_equal_limit bool
	low_limit           ast.ExprNode
	low_index           int

	high_limit           ast.ExprNode
	high_index           int
	has_high_equal_limit bool

	NotEqualMap map[int]ast.ExprNode
}

func newRanger() *ranger {
	return &ranger{NotEqualMap: make(map[int]ast.ExprNode)}
}

//false表明无节点替换
//true表明有节点替换
func (r *ranger) add(i int, val2 ast.ExprNode, op string, s server.Server) (int, bool) {
	var ok bool
	index := r.high_index
	switch op {
	case "<":
		if r.high_limit == nil {
			r.high_index = i
			r.high_limit = val2
			return -1, false
		}
		if ok = comp.Min(generateQuery(r.high_limit, val2, "<="), s); ok {
			//老节点被替换掉
			r.high_limit = val2
			r.high_index = i
			r.has_high_equal_limit = false
			return index, true
		} else {
			//替换当前节点
			return i, true
		}
	case ">":
		if r.low_limit == nil {
			r.low_index = i
			r.low_limit = val2
			return -1, false
		}
		if ok = comp.Max(generateQuery(r.low_limit, val2, ">="), s); ok {
			//老节点被替换掉
			r.low_limit = val2
			r.low_index = i
			r.has_low_equal_limit = false
			return index, true
		} else {
			//潘慧当前节点
			return i, true
		}
	case "!=", "<>":
		r.NotEqualMap[i] = val2
	case "<=":
		if r.high_limit == nil {
			r.high_limit = val2
			r.high_index = i
			r.has_high_equal_limit = true
			return -1, false
		}
		if ok = comp.Min(generateQuery(r.high_limit, val2, "<"), s); ok {
			//老节点被替换掉
			r.high_limit = val2
			r.high_index = i
			r.has_high_equal_limit = true
			return index, true
		} else {
			//替换当前节点
			return i, true
		}
	case ">=":
		if r.low_limit == nil {
			r.low_limit = val2
			r.low_index = i
			r.has_low_equal_limit = true
			return -1, false
		}
		if ok = comp.Max(generateQuery(r.low_limit, val2, ">"), s); ok {
			//老节点被替换掉
			r.low_limit = val2
			r.low_index = i
			r.has_low_equal_limit = true
			return index, true
		} else {
			//潘慧当前节点
			return i, true
		}
	}

	return -1, false
}

//true表示设置了=条件
//false表示没有等号条件
func (r *ranger) compare(expr []ast.ExprNode, s server.Server) bool {
	//情况一
	if r.low_limit != nil && r.high_limit == nil {
		return r.checkOnlyLowLimit(expr, s)
	}

	//情况二
	if r.low_limit == nil && r.high_limit != nil {
		return r.checkOnlyHighLimit(expr, s)
	}

	//情况三
	if r.low_limit != nil && r.high_limit != nil {
		return r.checkBothLowAndHigh(expr, s)
	}

	return false

}

func (r *ranger) checkOnlyLowLimit(expr []ast.ExprNode, s server.Server) bool {
	for index, e := range r.NotEqualMap {
		//比最小值还小
		if comp.Min(generateQuery(e, r.low_limit, "<"), s) {
			expr[index] = &ast.True{IsTrue: true}
			continue
		}

		//和最小值相等
		if comp.Equal(generateQuery(e, r.low_limit, "="), s) {
			if r.has_low_equal_limit {
				r.has_low_equal_limit = false
				expr[r.low_index].(*ast.Expr).Operator = ">"
				expr[index] = &ast.True{IsTrue: true}
			} else {
				expr[index] = &ast.True{IsTrue: true}
			}
			continue
		}
	}
	return false
}

func (r *ranger) checkOnlyHighLimit(expr []ast.ExprNode, s server.Server) bool {
	for index, e := range r.NotEqualMap {
		//和最大值相等
		if comp.Equal(generateQuery(e, r.high_limit, "="), s) {
			if r.has_high_equal_limit {
				r.has_high_equal_limit = false
				expr[r.high_index].(*ast.Expr).Operator = "<"
				expr[index] = &ast.True{IsTrue: true}
			} else {
				expr[index] = &ast.True{IsTrue: true}
			}
			continue
		}
		//比最大值还大
		if comp.Max(generateQuery(e, r.high_limit, ">"), s) {
			expr[index] = &ast.True{IsTrue: true}
			continue
		}
	}

	return false
}

func (r *ranger) checkBothLowAndHigh(expr []ast.ExprNode, s server.Server) bool {
	state := 0
	//如果上限比下限还小
	if comp.Min(generateQuery(r.high_limit, r.low_limit, "<"), s) {
		state = 1
		//如果上下限相同
	} else if comp.Equal(generateQuery(r.high_limit, r.low_limit, "="), s) {
		state = 2
		//下限小于上限
	} else {
		state = 3
	}

	switch state {
	default:
		return false
	case 1:
		expr[r.high_index] = &ast.True{IsTrue: false}
		return false
	case 2:
		//a>=1 and a<=1
		if r.has_high_equal_limit && r.has_low_equal_limit {
			//将上下限设置为同一个=表达式
			e := expr[r.high_index].(*ast.Expr)
			e.Operator = "="
			expr[r.high_index] = e
			expr[r.low_index] = &ast.True{IsTrue: true}

			for index, val := range r.NotEqualMap {
				if comp.Equal(generateQuery(r.high_limit, val, "="), s) {
					expr[index] = &ast.True{IsTrue: true}
				} else {
					expr[index] = &ast.True{IsTrue: false}
				}
			}
			return true
		} else {
			//1<a<1  1<=a<1  1<a<=1
			expr[r.high_index] = &ast.True{IsTrue: false}
			return false
		}
	case 3:
		for index, e := range r.NotEqualMap {
			//小于下限
			if comp.Min(generateQuery(e, r.low_limit, "<"), s) {
				expr[index] = &ast.True{IsTrue: true}
				continue
			}
			//等于下限
			if comp.Equal(generateQuery(e, r.low_limit, "="), s) {
				if r.has_low_equal_limit {
					r.has_low_equal_limit = false
					expr[r.low_index].(*ast.Expr).Operator = "<"
					expr[index] = &ast.True{IsTrue: true}
				} else {
					expr[index] = &ast.True{IsTrue: true}
				}
				continue
			}
			//等于上限
			if comp.Equal(generateQuery(e, r.high_limit, "="), s) {
				if r.has_high_equal_limit {
					r.has_high_equal_limit = false
					expr[r.high_index].(*ast.Expr).Operator = ">"
					expr[index] = &ast.True{IsTrue: true}
				} else {
					expr[index] = &ast.True{IsTrue: true}
				}
				continue
			}
			//大于上限
			if comp.Max(generateQuery(e, r.high_limit, ">"), s) {
				expr[index] = &ast.True{IsTrue: true}
				continue
			}
		}
		return false
	}

}
