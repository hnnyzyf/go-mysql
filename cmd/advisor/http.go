package main

import "github.com/hnnyzyf/mysql-go/advisor/parserhttp"
import "flag"
import "fmt"

//实现一个http服务的代码
func main() {

	//解析输入参数
	username := flag.String("U", "", "User for login")
	password := flag.String("P", "", "Password to use when connecting to server")

	flag.Parse()

	if *username == "" || *password == "" {
		fmt.Println("请输入正确的Username和Password以启动服务")
		return
	}

	//设置全局的账号密码
	parserhttp.User = *username
	parserhttp.PassWord = *password

	//启动Http服务
	parserhttp.ServerHttp()

}
