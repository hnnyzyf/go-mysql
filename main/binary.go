package main

import "flag"
import "parser"
import server "util/server"
import "bytes"
import p "plan"
import "fmt"

func main() {

	username := flag.String("u", "", "User for login if not current user.")
	password := flag.String("p", "", `Password to use when connecting to server. If password is not given it's asked from the tty.`)
	port := flag.String("P", "3306", `Port number to use for connection or 0 for default to, in order of preference, my.cnf, $MYSQL_TCP_PORT,/etc/services, built-in default (3306).`)
	database := flag.String("D", "Configdb", `Database to use`)
	host := flag.String("h", "", `Connect to host`)
	sql := flag.String("e", "SELECT 1", `Execute command and quit (Command should be quoted with """ ) `)

	flag.Parse()

	fmt.Println("待解析的语句为", *sql)

	if *username == "" || *password == "" || *host == "" || *port == "" || *database == "" {
		fmt.Println("请输入正确的连接信息")
		return
	}

	astree,err:= parser.SqlParse(*sql)
	if err!=nil{
		fmt.Println("语句解析出现错误,错误信息为:",err)
		return 
	}

	mysql := server.NewMysqlServer(*username, *password, *host, *port, *database)

	plan := p.NewPlan(astree, mysql)

	plan.PreProcessMeta()
	plan.PreProcessColumn()
	plan.ConstantFolding()

	buf := bytes.NewBuffer([]byte{})
	astree.Format(buf)
	fmt.Println("格式化后的语句为", buf.String())

	fmt.Println("检查可以添加的索引...")
	plan.ProcessIndex()

	if len(plan.Error) != 0 {
		fmt.Println("出现的错误为:")
	}
	for i, e := range plan.Error {
		fmt.Println("第", i, "个错误是 ", e)
	}
}
