package protocol

import (
	"crypto/tls"
	"io"
	"net"
	"time"

	"github.com/hnnyzyf/go-mysql/util/binary"
	"github.com/hnnyzyf/go-mysql/util/math"
	"github.com/juju/errors"
)

const (
	OK_Packet      uint8 = 0x00
	ERR_Packet     uint8 = 0xFF
	EOF_Packet     uint8 = 0xFE
	Unknown_Packet uint8 = 0xAA
)

const DefaultBufferSize = 16 * 1024

type Conn struct {
	//the real conn
	conn net.Conn

	//timeout for I/O read and write
	read  time.Duration
	write time.Duration

	//The sequence-id is incremented with each packet and may wrap around.
	//It starts at 0 and is reset to 0 when a new command begins in the Command Phase.
	sid uint8

	//header
	header []byte
}

//init a Conn without Timeout
func NewConn(conn net.Conn) *Conn {
	return &Conn{
		conn:   conn,
		sid:    0,
		header: make([]byte, 4),
	}
}

//Close the connection
func (c *Conn) Close() {
	c.conn.Close()
}

//Reset resets sid to 0
func (c *Conn) Reset() {
	c.sid = 0
}

//TransformToSSL tranform a connection to SSL connection
func (c *Conn) TransformToSSL(cfg *tls.Config) {
	c.conn = tls.Client(c.conn, cfg)
}

//ReadPacket read the packet from a io.Reader
func (c *Conn) ReadPacket() ([]byte, error) {
	//set read timeout
	if c.read != 0 {
		c.conn.SetReadDeadline(time.Now().Add(c.read))
	}

	//read the first 4 bytes
	if _, err := io.ReadFull(c.conn, c.header); err != nil {
		return nil, errors.Trace(err)
	}

	//get the length of payload
	length, err := binary.ReadInt3(c.header[0:3])
	if err != nil {
		return nil, errors.Trace(err)
	}
	if length > PayLoadMaxLen || length < 1 {
		return nil, errors.Errorf("Protocol:invalid payload length: %d", length)
	}

	//check sequenced id
	if sid, err := binary.ReadInt1(c.header[3:]); err != nil {
		return nil, errors.Trace(err)
		//if it is not command packet or handshake packet
	} else if sid != 0 && sid != c.sid {
		return nil, errors.Errorf("Protocol:invalid sequence id %d,expect %d", sid, c.sid)
		//if it is a command packet or handshake packet or just sid roll around from 255 to 0
		//when we read this packet,next sid should be 1
	} else if sid == 0 {
		c.sid = 1
		//sid == c.sid means that current packet is we expect ,so we just increase the sid
	} else {
		c.sid++
	}

	//write $length bytes into the buffer
	buffer := binary.NewBuffer(make([]byte, length))
	if n, err := io.CopyN(buffer, c.conn, int64(length)); err != nil {
		return nil, errors.Trace(err)
	} else if n != int64(length) {
		return nil, errors.Errorf("Protocol:Error package")
	} else {
		return buffer.Bytes(), nil
	}
}

//WritePacket write a packet into a io.Writer,If a packet is bigger than PayLoadMaxLen,
//it will be splited into some small packets
//size represent the Length of the payload.
//we use int to represent the length of payload,it is enough to store any length
//A payload is impossible to more than 18446744073709551615 bytes
//which is nearly 16777216TB
func (c *Conn) WritePacket(size int, payload io.Reader) error {
	//set write timeout
	if c.write != 0 {
		c.conn.SetWriteDeadline(time.Now().Add(c.write))
	}

	for size > 0 {
		//get the payload length and sid
		length := math.MinInt(size, int(PayLoadMaxLen))

		//length and header
		if err := binary.WriteInt3(c.header, uint32(length)); err != nil {
			return errors.Trace(err)
		}
		if err := binary.WriteInt1(c.header[3:], c.sid); err != nil {
			return errors.Trace(err)
		}

		//write header
		if n, err := c.conn.Write(c.header); err != nil {
			return errors.Trace(err)
		} else {
			if n != 4 {
				return errors.Errorf("Protocol:Fail to write 4 bytes header!")
			}
		}

		//write payload
		if n, err := io.CopyN(c.conn, payload, int64(length)); err != nil {
			return errors.Trace(err)
		} else {
			if n != int64(length) {
				return errors.Errorf("Protocol:Fail to write %d bytes payload!", length)
			}
		}

		//Next sid
		c.sid += 1

		//size and payload
		size -= length
	}

	return nil
}

//CopyPacket try to use send file to copy a packet from reader to writer
func CopyPacket(w io.Writer, r io.Reader) error {
	//undo
	return nil
}

func (c *Conn) SetReadTimeOut(read time.Duration) {
	c.read = read
}

func (c *Conn) SetWriteTimeOut(write time.Duration) {
	c.write = write
}
