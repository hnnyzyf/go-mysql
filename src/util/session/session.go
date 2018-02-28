package session

import "bytes"
import "fmt"

var DRIVER = "mysql"

type Session struct {
	User     string
	Password string
	Host     string
	Port     string
	DB       string
}

func NewSession(User string, Password string, Host string, Port string, DB string) *Session {
	return &Session{
		User:     User,
		Password: Password,
		Host:     Host,
		Port:     Port,
		DB:       DB,
	}
}

//生成连接字符串
func (s *Session) String() string {
	buf := bytes.NewBuffer([]byte{})
	fmt.Fprint(buf, s.User)
	fmt.Fprint(buf, ":")
	fmt.Fprint(buf, s.Password)
	fmt.Fprint(buf, "@tcp(")
	fmt.Fprint(buf, s.Host)
	fmt.Fprint(buf, ":")
	fmt.Fprint(buf, s.Port)
	fmt.Fprint(buf, ")/")
	fmt.Fprint(buf, s.DB)
	return buf.String()
}
