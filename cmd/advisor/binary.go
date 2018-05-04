package main

import "fmt"
import "flag"
import "bytes"
import "github.com/hnnyzyf/mysql-go/advisor/util/server"
import p "github.com/hnnyzyf/mysql-go/advisor/plan"
import "github.com/hnnyzyf/mysql-go/advisor/parser"

func printSlice(information []string) {
	for _, info := range information {
		fmt.Println(info)
	}
}

func printMap(information map[string][]string) {
	for name, info := range information {
		fmt.Println("可以添加的", name, "类型索引如下所示:")
		printSlice(info)
	}
}

func main() {

	username := flag.String("u", "", "User for login")
	password := flag.String("p", "", "Password to use when connecting to server")
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

	mysql := server.NewMysqlServer(*username, *password, *host, *port, *database)

	//解析数据库语句
	ast, err := parser.SqlParse(*sql)

	if err != nil {
		fmt.Println(err)
		return
	}

	//语句预处理
	plan := p.NewPlan(ast, mysql)

	err = plan.PreProcessMeta()
	if err != nil {
		fmt.Println(err)
		printSlice(plan.Error)
		return
	}

	//处理Column
	err = plan.PreProcessColumn()
	if err != nil {
		fmt.Println(err)
		printSlice(plan.Error)
		return
	}

	//处理constantfold
	err = plan.ConstantFolding()

	if err != nil {
		fmt.Println(err)
		return
	}

	buf := bytes.NewBuffer([]byte{})
	ast.Format(buf)

	fmt.Println("格式化后的语句如下所示:")
	fmt.Println(buf.String())

	//输出索引数据
	ok, info, index := plan.ProcessIndex()

	if ok != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("执行如下过程校验索引......")
	printSlice(info)

	fmt.Println("索引建议如下:")
	printMap(index)

	fmt.Println("警告信息如下:")
	printSlice(plan.Error)

}
