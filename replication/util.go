package replication

import (
	"github.com/hnnyzyf/go-mysql/util/binary"
)

//IsEventPacket tests whether current packet is event packet
func IsEventPacket(buffer []byte) bool {
	if header, err := binary.ReadInt1(buffer); err != nil {
		return false
	} else {
		return testEventPacket(header)
	}
}

//testEventPacket test whether current packet is err packet
func testEventPacket(header uint8) bool {
	return header == 0x00
}
