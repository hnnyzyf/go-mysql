package protocol

import (
	"io"
	"net"

	"github.com/hnnyzyf/go-mysql/util/binary"
	"github.com/hnnyzyf/go-mysql/util/math"
	"github.com/juju/errors"
)

const DefaultBufferSize = 16 * 1024

type Conn struct {
	//the real conn
	conn net.Conn

	//the sequence id
	sid uint8

	//header
	header []byte
	//the buffer we used to store payload,it may be resued
	//the size if 16KB
	buffer []byte
}

func NewConn(conn net.Conn) *Conn {
	return &Conn{
		conn:   conn,
		sid:    0,
		header: make([]byte, 4),
		buffer: make([]byte, DefaultBufferSize),
	}
}

//ReadPacket read the packet from a io.Reader
func (c *Conn) ReadPacket() ([]byte, error) {
	//read the first 4 bytes
	header := c.header
	if _, err := io.ReadFull(c.conn, header); err != nil {
		return nil, errors.Trace(err)
	}

	//get the length of payload
	length, err := binary.ReadInt3(header[0:3])
	if err != nil {
		return nil, errors.Trace(err)
	}
	if length > PayLoadMaxLen || length < 1 {
		return nil, errors.Errorf("Protocol:invalid payload length: %d", length)
	}

	//check sequenced id
	sid, err := binary.ReadInt1(header[3:])
	if err != nil {
		return nil, errors.Trace(err)
	}
	if sid != c.sid {
		return nil, errors.Errorf("Protocol:invalid sequence id %d,expect %d", sid, c.sid)
	} else {
		//turn next sid
		c.sid++
	}

	//write $length bytes into the buffer
	var buffer *binary.Buffer
	if length > DefaultBufferSize {
		buffer = binary.NewBuffer(make([]byte, length))
	} else {
		buffer = binary.NewBuffer(c.buffer[:length])
	}

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
	//alloc memroy for header
	header := c.header

	for size > 0 {
		//get the payload length and sid
		length := math.MinInt(size, int(PayLoadMaxLen))

		//length and header
		if err := binary.WriteInt3(header, uint32(length)); err != nil {
			return errors.Trace(err)
		}
		if err := binary.WriteInt1(header[3:], c.sid); err != nil {
			return errors.Trace(err)
		}

		//write header
		if n, err := c.conn.Write(header); err != nil {
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
