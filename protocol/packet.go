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
	size uint32
	//Sequence ID
	id uint8
	//[len=payload_length] payload of the packet
	p []byte
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

	//set current sequence
	sequenceId := binary.ReadInt1(header[3:])

	//write $length bytes into the buffer
	buffer := bytes.NewBuffer([]byte{})
	if n, err := io.CopyN(buffer, r, int64(length)); err != nil {
		return nil, errors.Trace(err)
	} else if n != int64(length) {
		return nil, errors.Errorf("Error package")
	} else {
		p := &Packet{
			size: length,
			id:   sequenceId,
			p:    buffer.Bytes(),
		}
		return p, nil
	}
}

//WritePacket write Packets
func WritePacket(w io.Writer,p *Packet)  
