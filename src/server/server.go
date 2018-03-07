package server

import "fmt"
import "bytes"
import "meta"
import "database/sql"
import _ "github.com/go-sql-driver/mysql"

type Server interface{
	//创建一个新的连接
	CreateSession2DB(string) string
	//Server负责查询
	Query(string,int,string) meta.Meta
}

//Server的实现
type Mysql struct {
	User     string
	Password string
	Host     string
	Port     string
	DB       string
}

const (
	_ = iota
	Task_relation
	Task_index
)

func NewMysqlServer(u string,p string,h string,o string,d string) *Mysql{
	return &Mysql{
		User:u,
		Password:p,
		Host:h,
		Port:o,
		DB:d,
	}
}


//生成连接字符串
func (m *Mysql)CreateConnectionString(DB string) string {
	buf := bytes.NewBuffer([]byte{})
	fmt.Fprint(buf, m.User)
	fmt.Fprint(buf, ":")
	fmt.Fprint(buf, m.Password)
	fmt.Fprint(buf, "@tcp(")
	fmt.Fprint(buf, m.Host)
	fmt.Fprint(buf, ":")
	fmt.Fprint(buf, m.Port)
	fmt.Fprint(buf, ")/")
	fmt.Fprint(buf, DB)
	return buf.String()
}


func (m *Mysql) CreateSession2DB (DB string) string{
	switch DB {
	case m.DB:
		return m.CreateConnectionString(m.DB)
	case "":
		return m.CreateConnectionString(m.DB)
	default:
		return m.CreateConnectionString(DB)
	}
}


func (m *Mysql) Query (session string,tag int,Object string) meta.Meta{
	var info meta.Meta
	switch tag {
	case Task_relation:
		info, _ = m.QueryTable(session,Object)
	case Task_index:
		info, _ = m.QueryIndex(session,Object)
	default:
		info, _ = m.QueryTable(session,Object)
	}
	return info 
}

func (m *Mysql) QueryTable(session string,ob string) (*meta.ColumnTable, error) {

	var Field string
	var Tp string
	var Null string
	var Key sql.NullString
	var Default sql.NullString
	var Extra sql.NullString

	//连接数据库
	db, err := sql.Open("mysql", session)
	if err != nil {
		return nil, err
	}

	//查询表结构
	buf := bytes.NewBuffer([]byte{})
	fmt.Fprint(buf, "DESC ")
	fmt.Fprint(buf, ob)
	rows, err := db.Query(buf.String())

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	//插入元数据
	Meta:=meta.NewColumnTable()

	for rows.Next() {
		err := rows.Scan(&Field, &Tp, &Null, &Key, &Default, &Extra)
		if err != nil {
			return nil, err
		}

		Meta.Insert(Field,Tp,Null)
	}

	//返回错误
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	defer db.Close()

	return Meta,nil
}


func (m *Mysql) QueryIndex(session string,ob string) (meta.Meta, error) {
	return nil,nil
}

func (m *Mysql) QueryNormal(session string,ob string) (meta.Meta, error) {
	return nil,nil
}