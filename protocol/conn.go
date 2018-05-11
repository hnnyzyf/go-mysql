package protocol

import (
	"github.com/juju/errors"
	"net"
)

//定义到Mysql的一个连接
type Session struct {
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

func NewSession(c *net.TCPConn, user string, passwd string, db string) *Session {
	return &Session{
		conn:     c,
		Username: user,
		Passwd:   passwd,
		Database: db,
	}
}

//Connect implements the content in https://dev.mysql.com/doc/internals/en/successful-authentication.html
func Connect(address string, username string, passwd string, db string) (*Session, error) {
	//the client connecting to the server
	raddr, err := net.ResolveTCPAddr("tcp", address)
	if err != nil {
		return nil, errors.Trace(err)
	}

	c, err := net.DialTCP("tcp", nil, raddr)
	if err != nil {
		return nil, errors.Trace(err)
	}

	conn := NewSession(c, username, passwd, db)

	//the server responds with the Initial Handshake Packet
	if err = conn.ReadHandshakeV10(); err != nil {
		return nil, errors.Trace(err)
	}

	//the client sends the Handshake Response Packet
	if err = conn.WriteHandshakeResponse41(); err != nil {
		return nil, errors.Trace(err)
	}

	//read ok packet
	if err := conn.ReadOkPacket(); err != nil {
		return nil, errors.Trace(err)
	}

	return conn, nil
}

//Close closes a Tcp connection
func (c *Session) Close() {
	c.conn.Close()
}

//ReadPacket reads a packet and return the payload
func (c *Session) ReadPacket() ([]byte, error) {
	//read the first 4 bytes
	header := make([]byte, 4)
	if _, err := io.ReadFull(c.conn, header); err != nil {
		return nil, errors.Trace(err)
	}

	//get the length of payload and the sequence id
	length := ReadInt3(header[0:3])

	if length > PayLoadMaxLen || length < 1 {
		return nil, errors.Errorf("invalid payload length: %d", length)
	}

	//set current sequence
	sequenceId := ReadInt1(header[3:])

	if sequenceId != c.SequenceId {
		return nil, errors.Errorf("invalid sequence id %d,expect %d", sequenceId, c.SequenceId)
	}

	c.SequenceId++

	//write $length bytes into the buffer
	buffer := bytes.NewBuffer([]byte{})
	if n, err := io.CopyN(buffer, c.conn, int64(length)); err != nil {
		return nil, errors.Trace(err)
	} else if n != int64(length) {
		return nil, errors.Errorf("Error package")
	} else {
		return buffer.Bytes(), nil
	}
}

//WritePacket write the payload into a packet 
func (c *Session) WritePacket(payload []byte) error {
	//the payload has 4 bytes header
	length := uint64(len(payload) - 4)

	//create packet,modify payload length
	//The sequence-id is incremented with each packet
	for length >= uint64(PayLoadMaxLen) {
		payload[0], payload[1], payload[2], payload[3] = 0xff, 0xff, 0xff, c.SequenceId
		if _, err := c.conn.Write(payload[:3+PayLoadMaxLen]); err != nil {
			return errors.Trace(err)
		} else {
			payload = payload[PayLoadMaxLen:]
			c.SequenceId++
		}
	}

	//modify the last payload
	payload[0] = uint8((len(payload) - 4))
	payload[1] = uint8((len(payload) - 4) >> 8)
	payload[2] = uint8((len(payload) - 4) >> 16)
	payload[3] = c.SequenceId
	c.SequenceId++
	if _, err := c.conn.Write(payload); err != nil {
		return errors.Trace(err)
	}

	return nil
}



type PayLoad struct {
	buffer []byte
	offset uint64
}

func NewPayLoad(buf []byte) *PayLoad {
	return &PayLoad{
		buffer: buf,
		offset: 0,
	}
}

func (p *PayLoad) IsEOF() bool {
	return uint64(len(p.buffer)) == p.offset+1
}

func (p *PayLoad) Bytes() []byte {
	return p.buffer
}

//ReadInt1 reads 1 bytes from payload
func (p *PayLoad) ReadInt1() uint8 {
	buffer := p.buffer[p.offset:]
	p.offset += 1
	return ReadInt1(buffer)
}

func (p *PayLoad) WriteInt1(val uint8) {
	p.buffer[p.offset] = byte(val)
	p.offset += 1
}

//ReadInt2 reads 2 bytes 
func (p *PayLoad) ReadInt2() uint16 {
	buffer := p.buffer[p.offset:]
	p.offset += 2
	return ReadInt2(buffer)
}

func (p *PayLoad) WriteInt2(val uint16) {
	p.buffer[p.offset] = uint8(val)
	p.buffer[p.offset+1] = uint8(val >> 8)
	p.offset += 2
}

//ReadInt3 reads 3 bytes 
func (p *PayLoad) ReadInt3() uint32 {
	buffer := p.buffer[p.offset:]
	p.offset += 3
	return ReadInt3(buffer)
}

func (p *PayLoad) WriteInt3(val uint32) {
	p.buffer[p.offset] = uint8(val)
	p.buffer[p.offset+1] = uint8(val >> 8)
	p.buffer[p.offset+2] = uint8(val >> 16)
	p.offset += 3
}

//ReadInt4 reads 4 bytes 
func (p *PayLoad) ReadInt4() uint32 {

	buffer := p.buffer[p.offset:]
	p.offset += 4
	return ReadInt4(buffer)
}

func (p *PayLoad) WriteInt4(val uint32) {
	p.buffer[p.offset] = uint8(val)
	p.buffer[p.offset+1] = uint8(val >> 8)
	p.buffer[p.offset+2] = uint8(val >> 16)
	p.buffer[p.offset+3] = uint8(val >> 24)
	p.offset += 4
}

//ReadInt6 reads 6 bytes 
func (p *PayLoad) ReadInt6() uint64 {
	buffer := p.buffer[p.offset:]
	p.offset += 6
	return ReadInt6(buffer)
}

func (p *PayLoad) WriteInt6(val uint32) {
	p.buffer[p.offset] = uint8(val)
	p.buffer[p.offset+1] = uint8(val >> 8)
	p.buffer[p.offset+2] = uint8(val >> 16)
	p.buffer[p.offset+3] = uint8(val >> 24)
	p.buffer[p.offset+4] = uint8(val >> 32)
	p.buffer[p.offset+5] = uint8(val >> 40)
	p.offset += 6
}

//ReadInt8 reads 8 bytes 
func (p *PayLoad) ReadInt8() uint64 {
	buffer := p.buffer[p.offset:]
	p.offset += 8
	return ReadInt8(buffer)
}

func (p *PayLoad) WriteInt8(val uint32) {
	p.buffer[p.offset] = uint8(val)
	p.buffer[p.offset+1] = uint8(val >> 8)
	p.buffer[p.offset+2] = uint8(val >> 16)
	p.buffer[p.offset+3] = uint8(val >> 24)
	p.buffer[p.offset+4] = uint8(val >> 32)
	p.buffer[p.offset+5] = uint8(val >> 40)
	p.buffer[p.offset+6] = uint8(val >> 48)
	p.buffer[p.offset+7] = uint8(val >> 54)
	p.offset += 8
}

func (p *PayLoad) ReadZero(length uint64) {
	p.offset += length
}

func (p *PayLoad) WriteZero(length uint64) {
	p.offset += length
}

//ReadStringWithNull reads a string terminated with null(0x00)
func (p *PayLoad) ReadStringWithNull() []byte {
	for off := p.offset; off < uint64(len(p.buffer)); off++ {
		if p.buffer[off] == 0x00 {
			offset := p.offset
			p.offset = off + 1
			//return not include 0x00
			return p.buffer[offset:off]
		}
	}
	return p.buffer[p.offset:]
}

func (p *PayLoad) WriteStringWithNull(str []byte) {
	_ = copy(p.buffer[p.offset:], str)
	p.offset += uint64(len(str))
	p.WriteZero(1)
}

//ReadStringWithFixLen reads a string  with fixed length
func (p *PayLoad) ReadStringWithFixLen(length uint64) []byte {
	p.offset = p.offset + length
	return p.buffer[p.offset-length : p.offset]
}

func (p *PayLoad) WriteStringWithFixLen(str []byte) {
	_ = copy(p.buffer[p.offset:], str)
	p.offset += uint64(len(str))
}

//ReadStringWithEOF reads a string util EOF
func (p *PayLoad) ReadStringWithEOF() []byte {
	return p.buffer[p.offset:]
}

func ReadInt1(buffer []byte) uint8 {
	_ = buffer[0]
	return uint8(buffer[0])
}

func ReadInt2(buffer []byte) uint16 {
	_ = buffer[1]
	return uint16(buffer[0] | buffer[1]<<8)
}

func ReadInt3(buffer []byte) uint32 {
	_ = buffer[2]
	return uint32(buffer[0] | buffer[1]<<8 | buffer[2]<<16)
}

func ReadInt4(buffer []byte) uint32 {
	_ = buffer[4]
	return uint32(buffer[0] | buffer[1]<<8 | buffer[2]<<16 | buffer[3]<<24)
}

func ReadInt6(buffer []byte) uint64 {
	_ = buffer[5]
	return uint64(buffer[0] | buffer[1]<<8 | buffer[2]<<16 | buffer[3]<<24 | buffer[4]<<32 | buffer[5]<<40)
}

func ReadInt8(buffer []byte) uint64 {
	_ = buffer[7]
	return uint64(buffer[0] | buffer[1]<<8 | buffer[2]<<16 | buffer[3]<<24 | buffer[4]<<32 | buffer[5]<<40 | buffer[6]<<48 | buffer[7]<<56)
}
