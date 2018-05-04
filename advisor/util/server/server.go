package server

import "github.com/hnnyzyf/mysql-go/advisor/util/meta"
import "bytes"
import "fmt"

const (
	_ = iota
	Task_relation
	Task_index
)

type Server interface {
	//创建一个新的连接
	CreateSession2DB(string) string
	//Server负责查询
	QueryTable(string, string) (*meta.Meta, error)
	QueryTrue0False(string, string, int) (bool, error)
	QuerySelectivity(string, string) (int, error)
	QueryIndex(string, string) (*meta.Index, error)
	QueryRow(string, string) (int, error)
	QueryCardinality(string, string) (int, error)
}

//Server的实现
type Mysql struct {
	User     string
	Password string
	Host     string
	Port     string
	DB       string
}

func NewMysqlServer(u string, p string, h string, o string, d string) *Mysql {
	return &Mysql{
		User:     u,
		Password: p,
		Host:     h,
		Port:     o,
		DB:       d,
	}
}

//生成连接字符串
func (m *Mysql) CreateConnectionString(DB string) string {
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

func (m *Mysql) CreateSession2DB(DB string) string {
	switch DB {
	case m.DB:
		return m.CreateConnectionString(m.DB)
	case "":
		return m.CreateConnectionString(m.DB)
	default:
		return m.CreateConnectionString(DB)
	}
}
