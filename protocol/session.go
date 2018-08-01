package protocol

//Session represents a connection to server
type Session struct {
	//the real connect
	conn *net.TCPConn

	Username string
	Passwd   string
	Database string

	SequenceId   uint8
	ConnectionId uint32
	Character    uint8

	Capabilities uint32

	PluginName     []byte
	AuthPluginData []byte
}