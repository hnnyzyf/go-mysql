package client

import (
	"time"
	
	"github.com/hnnyzyf/go-mysql/protocol"
	"github.com/hnnyzyf/go-mysql/util/character"
)

var DefaultConfig = &Config{
	deadline:0,
	readDeadline:0,
	writeDeadline:0,
	keepAlive:0,
	pluginName:protocol.DefaultPluginName,
	autoCommit:true,
	character:

}


type Config struct {
	//the timeout of the I/O operation 
	deadline time.Duration

	//the timeout of the read operation
	readDeadline time.Duration

	//the timeout of the wirte operaion 
	writeDeadline time.Duration

	//the period of a keepalive packet
	keepAlive time.Duration

	//the name of auth plugin 
	pluginName string

	//autocommit
	autoCommit bool

	//charset 
	charset *character.Charset

	//Default SQL_MODE to use
	sqlMode uint8

	//capabilities flag use in client
	capabilities uint32

	//wether use ssl to connct
	isSSL bool

	//compress:not support!
}

func NewConfig() *Config{
	return &Config{}
}