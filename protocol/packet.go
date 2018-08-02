package protocol

import (
	"bytes"
	"io"

	"github.com/hnnyzyf/go-mysql/util/binary"
	"github.com/juju/errors"
)

//Packet is composed of three parts
type Packet struct {
	//Length of the payload.
	//The number of bytes in the packet beyond the initial 4 bytes that make up the packet header.
	length []byte
	//Sequence ID
	sid []byte
	//[len=payload_length] payload of the packet
	p *binary.Payload
}

//ReadLength returns the length of payload
func (p *Packet) ReadLength() uint32 {
	return binary.ReadInt3(p.length)
}

//ReadLength returns the id of a packet
func (p *Packet) ReadId() uint8 {
	return binary.ReadInt1(p.sid)
}

//ReadPayload return the payload of a packet
func (p *Packet) ReadPayload() *binary.Payload {
	return p.p
}

//ReadPacket read a packet from a io.Reader
func ReadPacket(r io.Reader) (*Packet, error) {
	//read the first 4 bytes
	header := make([]byte, 4)
	if _, err := io.ReadFull(r, header); err != nil {
		return nil, errors.Trace(err)
	}

	//get the length of payload and the sequence id
	length := binary.ReadInt3(header[0:3])
	if length > PayLoadMaxLen || length < 1 {
		return nil, errors.Errorf("invalid payload length: %d", length)
	}

	//write $length bytes into the buffer
	buffer := bytes.NewBuffer([]byte{})
	if n, err := io.CopyN(buffer, r, int64(length)); err != nil {
		return nil, errors.Trace(err)
	} else if n != int64(length) {
		return nil, errors.Errorf("Error package")
	} else {
		return &Packet{
			length: header[0:3],
			sid:    header[3:],
			p:      binary.NewReader(buffer.Bytes()),
		}, nil
	}
}

//WritePacket write a packet into a io.Writer
func WritePacket(w io.Writer, p *Packet) error {
	length := p.ReadLength()
	if length > PayLoadMaxLen || length < 1 {
		return errors.Errorf("invalid payload length: %d", length)
	}

	if p.p.Len() != int(length) {
		return errors.Errorf("invalid payload: %d", length)
	}

	//write length
	n, err := w.Write(p.length)
	if err != nil {
		return errors.Trace(err)
	} else if n != 3 {
		return errors.Errorf("Fail to write the length of payload,Expect 3 bytes( %d bytes)", n)
	} else {
		//donothing
	}

	n, err = w.Write(p.sid)
	if err != nil {
		return errors.Trace(err)
	} else if n != 1 {
		return errors.Errorf("Fail to write the sequence id!")
	} else {
		//donothing
	}

	n, err = w.Write(p.p.Bytes())
	if err != nil {
		return errors.Trace(err)
	} else if n != int(length) {
		return errors.Errorf("Fail to write the payload!Expect %d bytes( %d bytes)", length, n)
	} else {
		//donothing
	}

	return nil
}
