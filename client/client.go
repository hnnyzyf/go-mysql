package client

import (
	"net"

	"github.com/hnnyzyf/go-mysql/protocol"
	"github.com/hnnyzyf/go-mysql/util/binary"
	"github.com/hnnyzyf/go-mysql/util/charset"
	"github.com/hnnyzyf/go-mysql/util/hack"
	"github.com/hnnyzyf/go-mysql/util/math"
	"github.com/juju/errors"
	"golang.org/x/text/encoding"
)

var errHandShake10 = errors.Errorf("Client:not a valid HandShake10 packets")
var errResponsePacket = errors.Errorf("Client:not a valid response Packet")
var errERRPacket = errors.Errorf("Client:not a valid Err packet")

//Session represent a connection to server from client
type Session struct {
	//the real conn to server,the
	*protocol.Conn

	//connection inforamtion
	username string
	password string
	database string

	//some options
	cfg *config

	//capabilities flag use in client
	capabilities uint32

	//20-bytes random data from serve
	authPluginData []byte
	//authentication method
	authMethod protocol.Method

	//charset
	charset uint8
	//charset method
	decoder *encoding.Decoder
	encoder *encoding.Encoder

	res *response
}

//connect create a tcp connection to server
func connect(host string, user string, passwd string, db string, cfg *config) (*Session, error) {
	//the client connecting to the server with timeout
	conn, err := net.DialTimeout("tcp", host, cfg.timeout)
	if err != nil {
		return nil, errors.Trace(err)
	}

	//init Session
	s := &Session{
		Conn:     protocol.NewConn(conn),
		username: user,
		password: passwd,
		database: db,
		cfg:      cfg,
		res:      &response{},
	}

	//init configuration
	if err := s.init(); err != nil {
		return nil, errors.Trace(err)
	}

	//the server responds with the Initial Handshake Packet
	if err := s.readHandShakeV10(); err != nil {
		return nil, errors.Trace(err)
	}

	if cfg.isSSL {
		//the client use ssl to sends the SSL Request Packet
		if err := s.writeSSLRequeset(); err != nil {
			return nil, errors.Trace(err)
		}
	}

	//the client sends the Handshake Response Packet
	if err := s.writeHandshakeResponse41(); err != nil {
		return nil, errors.Trace(err)
	}

	//read response packet and decide what to do according to the type of packet
	// 		OK-packet:return success
	//		ERR-packet:return false
	//		EOF-packet:return false
	if buffer, err := s.ReadPacket(); err != nil {
		return nil, errors.Trace(err)
	} else {
		switch {
		case s.IsOkPacket(buffer):
			return s, nil
		case s.IsErrPacket(buffer):
			return s, s.ReadErrPacket(buffer)
		default:
			return nil, errors.Errorf("Client:fail to accept correct generic response packet from server!")
		}
	}

	return s, nil
}

//Connect is a wrapper of connect
func Connect(host string, user string, passwd string, db string) (*Session, error) {
	return connect(host, user, passwd, db, defaultConfig)
}

//ConnectWithConfig is a wrapper of connect
func ConnectWithConfig(host string, user string, passwd string, db string, cfg *config) (*Session, error) {
	return connect(host, user, passwd, db, cfg)
}

//init read configuration from config
func (s *Session) init() error {
	cfg := s.cfg

	//always use ok packet to relace eof packet after a Text Resultset
	s.capabilities = cfg.capabilities | protocol.CLIENT_DEPRECATE_EOF

	if len(s.database) == 0 {
		s.capabilities &= (^protocol.CLIENT_CONNECT_WITH_DB)
	} else {
		s.capabilities |= protocol.CLIENT_CONNECT_WITH_DB
	}

	if cfg.isSSL {
		s.capabilities |= protocol.CLIENT_SSL
	}

	if cfg.mutilStatement {
		s.capabilities |= protocol.CLIENT_MULTI_STATEMENTS
	}

	//test ssl
	if cfg.isSSL && cfg.tlsConfig == nil {
		return errors.Errorf("Client:not provide tls configuraions")
	}

	//test connection attr
	if cfg.attrs.size() == 0 {
		s.capabilities &= (^protocol.CLIENT_CONNECT_ATTRS)
	} else {
		s.capabilities |= protocol.CLIENT_CONNECT_ATTRS
	}

	//set charset method
	if method, err := charset.GetMethod(cfg.charset); err != nil {
		return errors.Trace(err)
	} else {
		s.decoder = method.NewDecoder()
		s.encoder = method.NewEncoder()
	}

	//set charset id
	if id, err := charset.GetID(cfg.charset); err != nil {
		return errors.Trace(err)
	} else {
		s.charset = id
	}

	return nil
}

//readHandShakeV10 Read the init handshake packet
func (s *Session) readHandShakeV10() error {
	//Get packet from conn
	buffer, err := s.ReadPacket()
	if err != nil {
		return errors.Trace(err)
	}

	//init payload
	payload := binary.NewBuffer(buffer)

	//read protocol version 1 bytes
	if protocolVersion, err := payload.ReadInt1(); err != nil {
		return errHandShake10
	} else {
		if !testProtocolVersion(protocolVersion) {
			return errors.Errorf("Client:not a valid protocol version: %d ", protocolVersion)
		}
	}
	//read server version string.null
	if _, err := payload.ReadStringWithNull(); err != nil {
		return errHandShake10
	}

	//read connect id 4 bytes
	if _, err := payload.ReadInt4(); err != nil {
		return errHandShake10
	}

	//read auth-plugin-data-part-1 8 bytes
	authPluginData, err := payload.ReadStringWithFixLen(8)
	if err != nil {
		return errHandShake10
	} else {
		s.authPluginData = authPluginData
	}

	//skip filler 1 bytes
	if _, err := payload.ReadInt1(); err != nil {
		return errHandShake10
	}

	//read low capability flags (lower 2 bytes)
	lflag, err := payload.ReadInt2()
	if err != nil {
		return errHandShake10
	}

	//if no more data in the packet
	if payload.IsEOF() {
		return nil
	}

	//read character set 1 bytes
	if _, err := payload.ReadInt1(); err != nil {
		return errHandShake10
	}

	//read status flag
	if _, err := payload.ReadInt2(); err != nil {
		return errHandShake10
	}

	//read upper capability flags (upper 2 bytes)
	uflag, err := payload.ReadInt2()
	if err != nil {
		return errHandShake10
	}

	//set Capabilities (4 bytes)
	capabilities := uint32(uflag)<<16 | uint32(lflag)
	if s.cfg.isSSL && !testSSL(capabilities) {
		return errors.Errorf("Client:server does not support SSL connection!")
	}

	//The client should only announce the capabilities
	//in the Handshake Response Packet that it has in common with the server.
	s.capabilities &= capabilities

	//length of auth-plugin-data 1bytes
	var authPluginDataLen uint8
	if testPluginAuth(capabilities) {
		if length, err := payload.ReadInt1(); err != nil {
			return errHandShake10
		} else {
			authPluginDataLen = length
		}
	} else {
		if err := payload.Skip(1); err != nil {
			return errHandShake10
		} else {
			authPluginDataLen = 0
		}
	}

	//skip reserved 10 bytes
	if err := payload.Skip(10); err != nil {
		return errHandShake10
	}

	//I find it does not work
	//https://dev.mysql.com/doc/internals/en/secure-password-authentication.html shows
	//client-side expects a 20-byte random challenge
	//if we use $len=MAX(13, length of auth-plugin-data - 8 as the length of auth_plugin_data_part2
	//the length is more than 21 instead of 20
	//so I decide to read only 12 bytes data as the challenge
	//maybe I should have a look at Mysql source code!?
	if testSecureConnection(capabilities) {
		length := math.MaxInt(13, int(authPluginDataLen)-8)
		lauthPluginData, err := payload.ReadStringWithFixLen(12)
		if err != nil {
			return errHandShake10
		}
		//skip
		if err := payload.Skip(length - 12); err != nil {
			return errHandShake10
		}
		s.authPluginData = append(s.authPluginData, lauthPluginData...)
	}

	//read pluginName and set authentication method
	var pluginName string
	if testPluginAuth(capabilities) {
		if name, err := payload.ReadStringWithNull(); err != nil {
			return errHandShake10
		} else {
			pluginName = hack.String(name)
		}
	} else {
		//https://dev.mysql.com/doc/internals/en/determining-authentication-method.html
		if !testProtocol41(capabilities) || !testSecureConnection(capabilities) {
			pluginName = "mysql_old_password"
		} else if testProtocol41(capabilities) && testSecureConnection(capabilities) {
			pluginName = "mysql_native_password"
		} else {
			return errors.Errorf("Client:server do not support any authentication method")
		}
	}
	if method, err := protocol.GetMethod(pluginName); err != nil {
		return errors.Trace(err)
	} else {
		s.authMethod = method
	}

	return nil
}

//writeSSLRequeset create ssl communication to server
func (s *Session) writeSSLRequeset() error {
	size := 4 + 4 + 1 + 23

	//create a payload
	payload := binary.NewBuffer(make([]byte, size))

	//write capability
	if err := payload.WriteInt4(s.capabilities); err != nil {
		return errors.Trace(err)
	}

	//write max_packet size
	if err := payload.WriteInt4(protocol.PayLoadMaxLen); err != nil {
		return errors.Trace(err)
	}

	//write chracter set
	if err := payload.WriteInt1(s.charset); err != nil {
		return errors.Trace(err)
	}

	//write packet
	if err := s.WritePacket(size, payload.ToReader()); err != nil {
		return errors.Trace(err)
	}

	//replace the origin connection to ssl connection
	s.TransformToSSL(s.cfg.tlsConfig)

	return nil
}

//https://dev.mysql.com/doc/internals/en/connection-phase-packets.html#packet-Protocol::SSLRequest
//writeHandshakeResponse41 return a HandshakeResponse to server
func (s *Session) writeHandshakeResponse41() error {
	//write a header four bytes
	size := 0
	//add capabilities
	size += 4

	//set max-packet size character set  reserved (all [0])
	size += 4 + 1 + 23

	//username+0x0 and convert database from utf8 to the charset
	var username []byte
	if name, err := s.decoder.Bytes(hack.Slice(s.username)); err != nil {
		return errors.Trace(err)
	} else {
		username = name
	}
	size += len(username) + 1

	//calculate auth-response
	password := s.authMethod.EncodeKey(s.authPluginData, s.password)
	size += len(password)
	if testLenencData(s.capabilities) {
		size += binary.LengthOfInteger(uint64(len(password)))
	} else if testSecureConnection(s.capabilities) {
		size += 1
	} else {
		size += 1
	}

	//set database name
	var database []byte
	if testConnectWithDB(s.capabilities) {
		//convert database from utf8 to the charset
		if db, err := s.decoder.Bytes(hack.Slice(s.database)); err != nil {
			return errors.Trace(err)
		} else {
			database = db
		}
		size += len(database) + 1
	}

	//auth plugin name
	if testPluginAuth(s.capabilities) {
		size += len(s.authMethod.Name()) + 1
	}

	//add attributes
	var attr []byte
	if testConnectAttrs(s.capabilities) {
		if b, err := s.cfg.attrs.marshal(); err != nil {
			return errors.Trace(err)
		} else {
			size += len(b)
			attr = b
		}
	}

	//create a payload
	payload := binary.NewBuffer(make([]byte, size))

	//write capability flags
	if err := payload.WriteInt4(s.capabilities); err != nil {
		return errors.Trace(err)
	}

	//max-packet size
	if err := payload.WriteInt4(protocol.PayLoadMaxLen); err != nil {
		return errors.Trace(err)
	}

	//character set
	if err := payload.WriteInt1(s.charset); err != nil {
		return errors.Trace(err)
	}

	//reserved
	if err := payload.Skip(23); err != nil {
		return errors.Trace(err)
	}

	//username
	if err := payload.WriteStringWithNull(username); err != nil {
		return errors.Trace(err)
	}

	//write auth-response
	if testLenencData(s.capabilities) {
		if _, err := payload.WriteLengthEncodedInteger(uint64(len(password))); err != nil {
			return errors.Trace(err)
		}
		if err := payload.WriteStringWithFixLen(password); err != nil {
			return errors.Trace(err)
		}
	} else if testSecureConnection(s.capabilities) {
		if err := payload.WriteInt1(uint8(len(password))); err != nil {
			return errors.Trace(err)
		}
		if err := payload.WriteStringWithFixLen(password); err != nil {
			return errors.Trace(err)
		}
	} else {
		if err := payload.WriteStringWithNull(password); err != nil {
			return errors.Trace(err)
		}
	}

	//database
	if testConnectWithDB(s.capabilities) {
		if err := payload.WriteStringWithNull(database); err != nil {
			return errors.Trace(err)
		}
	}

	//auth plugin name
	if testPluginAuth(s.capabilities) {
		if err := payload.WriteStringWithNull(hack.Slice(s.authMethod.Name())); err != nil {
			return errors.Trace(err)
		}
	}

	//set attribute
	if testConnectAttrs(s.capabilities) {
		if err := payload.WriteStringWithFixLen(attr); err != nil {
			return errors.Trace(err)
		}
	}

	//write packet
	return s.WritePacket(size, payload.ToReader())
}
