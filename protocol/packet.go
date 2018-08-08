package protocol

import (
	"io"

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

//Packet is composed of three parts
type Packet struct {
	//Length of the payload.
	//we use uint64 to represent the length of payload,it is enough to store any length
	//A payload is impossible to more than 18446744073709551615 bytes
	//which is nearly 16777216TB
	length int
	//Sequence ID
	sid uint8
	//[len=payload_length] payload of the packet
	payload io.Reader
}

//test whether the payload should be reset offset
type testResetOffset interface {
	ResetOffset()
}

func NewPacket(l int, i uint8, p io.Reader) *Packet {
	//reset offset if the payload need
	if buffer, ok := p.(testResetOffset); ok {
		buffer.ResetOffset()
	}

	return &Packet{
		length:  l,
		sid:     i,
		payload: p,
	}

}

//ReadLength returns the length of payload
func (p *Packet) ReadLength() int {
	return p.length
}

//ReadLength returns the id of a packet
func (p *Packet) ReadId() uint8 {
	return p.sid
}

//ReadPayload return the payload of a packet
func (p *Packet) ReadPayload() (*binary.Buffer, error) {
	if b, ok := p.payload.(*binary.Buffer); !ok {
		return nil, errors.Errorf("Protocol:Not a corret Packet!")
	} else {
		return b, nil
	}
}

//ReadPacket read the id packet from a io.Reader and write it into io.Writer
func ReadPacket(r io.Reader, id uint8) (*Packet, error) {
	//read the first 4 bytes
	header := make([]byte, 4)
	if _, err := io.ReadFull(r, header); err != nil {
		return nil, errors.Trace(err)
	}

	//get the length of payload
	length, err := binary.ReadInt3(header[0:3])
	if err != nil {
		return nil, errors.Trace(err)
	}
	if length > PayLoadMaxLen || length < 1 {
		return nil, errors.Errorf("invalid payload length: %d", length)
	}

	//check sequenced id
	sid, err := binary.ReadInt1(header[3:])
	if err != nil {
		return nil, errors.Trace(err)
	}
	if sid != id {
		return nil, errors.Errorf("invalid sequence id %d,expect %d", sid, id)
	}

	//write $length bytes into the buffer
	payload := binary.NewBuffer(make([]byte, length))

	if n, err := io.CopyN(payload, r, int64(length)); err != nil {
		return nil, errors.Trace(err)
	} else if n != int64(length) {
		return nil, errors.Errorf("Error package")
	} else {
		return NewPacket(int(length), sid, payload), nil
	}
}

//WritePacket write a packet into a io.Writer
//If a packet is bigger than PayLoadMaxLen,
//it will be splited into some small packets
func WritePacket(w io.Writer, packet *Packet) (uint8, error) {
	//get size and payload
	size := packet.length
	sid := packet.sid
	payload := packet.payload

	//alloc memroy for header
	header := make([]byte, 4)

	//while length<=0
	for size > 0 {
		//get the payload length and sid
		length := math.MinInt(size, int(PayLoadMaxLen))
		sid += 1

		//length and header
		if err := binary.WriteInt3(header, uint32(length)); err != nil {
			return sid, errors.Trace(err)
		}
		if err := binary.WriteInt1(header[3:], sid); err != nil {
			return sid, errors.Trace(err)
		}

		//write header
		if n, err := w.Write(header); err != nil {
			return sid, errors.Trace(err)
		} else {
			if n != 4 {
				return sid, errors.Errorf("Fail to write 4 bytes header!")
			}
		}

		//write payload
		if n, err := io.CopyN(w, payload, int64(length)); err != nil {
			return sid, errors.Trace(err)
		} else {
			if n != int64(length) {
				return sid, errors.Errorf("Fail to write %d bytes payload!", length)
			}
		}

		//size and payload
		size -= length
	}

	return sid, nil
}

//CopyPacket try to use send file to copy a packet from reader to writer
func CopyPacket(w io.Writer, r io.Reader) error {
	//undo
	return nil
}
