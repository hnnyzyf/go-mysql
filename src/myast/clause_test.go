package myast 


import "testing"


func Test_NewClasue(t *testing.T){
		Distinctclause:=NewDistinctclause(true)
		if Distinctclause.NType==CLAUSE_DISTINCT{
			t.Log("Distinctclause通过测试")
		}else{
			t.Error("Distinctclause失败了")
		}

		Intoclause:=NewIntoclause(true)
		if Intoclause.NType==CLAUSE_INTO{
			t.Log("Intoclause通过测试")
		}else{
			t.Error("Intoclause失败了")
		}

		Fromclause:=NewFromclause(true)
		if Fromclause.NType==CLAUSE_FROM{
			t.Log("Fromclause通过测试")
		}else{
			t.Error("Fromclause失败了")
		}


		Whereclause:=NewWhereclause(true)
		if Whereclause.NType==CLAUSE_WHERE{
			t.Log("Whereclause通过测试")
		}else{
			t.Error("Whereclause失败了")
		}


		Groupclause:=NewGroupclause(true)
		if Groupclause.NType==CLAUSE_GROUPBY{
			t.Log("Groupclause通过测试")
		}else{
			t.Error("Groupclause失败了")
		}


		Havingclause:=NewHavingclause(true)
		if Havingclause.NType==CLAUSE_HAVING{
			t.Log("Havingclause通过测试")
		}else{
			t.Error("Havingclause失败了")
		}

		Sortclause:=NewSortclause(true)
		if Sortclause.NType==CLAUSE_SORTBY{
			t.Log("Sortclause通过测试")
		}else{
			t.Error("Sortclause失败了")
		}

		Lockingclause:=NewLockingclause(true)
		if Lockingclause.NType==CLAUSE_LOCK{
			t.Log("Lockingclause通过测试")
		}else{
			t.Error("Lockingclause失败了")
		}

		Limitclause:=NewLimitclause(true)
		if Limitclause.NType==CLAUSE_LIMIT{
			t.Log("Limitclause通过测试")
		}else{
			t.Error("Limitclause失败了")
		}
}