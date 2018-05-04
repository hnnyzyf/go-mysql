package comp

import "github.com/hnnyzyf/mysql-go/advisor/util/server"

func Compare(query string, s server.Server) bool {
	session := s.CreateSession2DB("")
	if ok, _ := s.QueryTrue0False(session, query, 1); ok {
		return true
	} else {
		return false
	}
}

//true表明原节点需要被替换掉
//false表示新节点要被替换掉
func Min(query string, s server.Server) bool {
	//生成比较的字符串
	return Compare(query, s)
}

func Max(query string, s server.Server) bool {
	//生成比较的字符串
	return Compare(query, s)
}

func Equal(query string, s server.Server) bool {
	//生成比较的字符串
	return Compare(query, s)
}
