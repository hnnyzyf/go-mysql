package payload

//ReadInt1 reads 1 bytes
func ReadInt1(buffer []byte) uint8 {
	return uint8(buffer[0])
}

//ReadInt2 reads 2 bytes
func ReadInt2(buffer []byte) uint16 {
	b := uint16(buffer[0])
	var i uint8 = 1
	for i < 2 {
		b = b | uint16(buffer[i])<<(i*8)
		i++
	}
	return b
}

//ReadInt3 reads 3 bytes
func ReadInt3(buffer []byte) uint32 {
	b := uint32(buffer[0])
	var i uint8 = 1
	for i < 3 {
		b = b | uint32(buffer[i])<<(i*8)
		i++
	}
	return b
}

//ReadInt4 reads 4 bytes
func ReadInt4(buffer []byte) uint32 {
	b := uint32(buffer[0])
	var i uint8 = 1
	for i < 4 {
		b = b | uint32(buffer[i])<<(i*8)
		i++
	}
	return b
}

//ReadInt6 reads 6 bytes
func ReadInt6(buffer []byte) uint64 {
	b := uint64(buffer[0])
	var i uint8 = 1
	for i < 6 {
		b = b | uint64(buffer[i])<<(i*8)
		i++
	}
	return b
}

//ReadInt8 reads 8 bytes
func ReadInt8(buffer []byte) uint64 {
	b := uint64(buffer[0])
	var i uint8 = 1
	for i < 8 {
		b = b | uint64(buffer[i])<<(i*8)
		i++
	}
	return b
}

//WriteInt1 writes 1 byte into buffer
func WriteInt1(b []byte, val uint8) {
	b[0] = val
}

//WriteInt2 writes 2 bytes into buffer
func WriteInt2(b []byte, val uint16) {
	b[0] = uint8(val)
	b[1] = uint8(val >> 8)
}

//WriteInt3 writes 3 bytes into buffer
func WriteInt3(b []byte, val uint32) {
	b[0] = uint8(val)
	b[1] = uint8(val >> 8)
	b[2] = uint8(val >> 16)
}

//WriteInt4 writes 4 bytes into buffer
func WriteInt4(b []byte, val uint32) {
	b[0] = uint8(val)
	b[1] = uint8(val >> 8)
	b[2] = uint8(val >> 16)
	b[3] = uint8(val >> 24)

}

//WriteInt6 writes 6 bytes into buffer
func WriteInt6(b []byte, val uint64) {
	b[0] = uint8(val)
	b[1] = uint8(val >> 8)
	b[2] = uint8(val >> 16)
	b[3] = uint8(val >> 24)
	b[4] = uint8(val >> 32)
	b[5] = uint8(val >> 40)

}

//WriteInt8 writes 8 bytes into buffer
func WriteInt8(b []byte, val uint64) {
	b[0] = uint8(val)
	b[1] = uint8(val >> 8)
	b[2] = uint8(val >> 16)
	b[3] = uint8(val >> 24)
	b[4] = uint8(val >> 32)
	b[5] = uint8(val >> 40)
	b[6] = uint8(val >> 48)
	b[7] = uint8(val >> 56)
}

//payload represent a buffer in a packet
type payload struct {
	b   []byte
	off int
}

func New(b []byte) *payload {
	return &payload{
		b:   b,
		off: 0,
	}
}

//IsEOf return
func (p *payload) IsEOF() bool {
	return p.off >= len(p.b)
}

func (p *payload) Bytes() []byte {
	return p.b
}

//ReadInt1 reads 1 bytes
func (p *payload) ReadInt1() uint8 {
	buffer := p.b[p.off:]
	p.off += 1
	return ReadInt1(buffer)
}

//ReadInt2 reads 2 bytes
func (p *payload) ReadInt2() uint16 {
	buffer := p.b[p.off:]
	p.off += 2
	return ReadInt2(buffer)
}

//ReadInt3 reads 3 bytes
func (p *payload) ReadInt3() uint32 {
	buffer := p.b[p.off:]
	p.off += 3
	return ReadInt3(buffer)
}

//ReadInt4 reads 4 bytes
func (p *payload) ReadInt4() uint32 {
	buffer := p.b[p.off:]
	p.off += 4
	return ReadInt4(buffer)
}

//ReadInt6 reads 6 bytes
func (p *payload) ReadInt6() uint64 {
	buffer := p.b[p.off:]
	p.off += 6
	return ReadInt6(buffer)
}

//ReadInt8 reads 8 bytes
func (p *payload) ReadInt8() uint64 {
	buffer := p.b[p.off:]
	p.off += 8
	return ReadInt8(buffer)
}

//ReadZero reads length bytes
func (p *payload) ReadZero(l int) {
	p.off += l
}

//ReadLengthEncodedInteger read undetermined length integer
func (p *payload) ReadLengthEncodedInteger() (uint64, int) {
	switch i := p.b[p.off]; {
	case i < 0xfb:
		return uint64(p.ReadInt1()), 1
	case i == 0xfc:
		return uint64(p.ReadInt2()), 2
	case i == 0xfd:
		return uint64(p.ReadInt3()), 3
	case i == 0xfe:
		return uint64(p.ReadInt8()), 8
	default:
		panic("Invaid LengthEncodedInteger flag!")
	}

	return 0, -1
}

//ReadStringWithNull reads a string terminated with null(0x00)
func (p *payload) ReadStringWithNull() []byte {
	for off := p.off; off < len(p.b); off++ {
		if p.b[off] == 0x00 {
			//return not include 0x00
			b := p.b[p.off:off]
			p.off = off + 1
			return b
		}
	}
	//return all bytes
	return p.b[p.off:]
}

//ReadStringWithEOF reads a string util EOF
func (p *payload) ReadStringWithEOF() []byte {
	return p.b[p.off:]
}

//ReadStringWithFixLen reads a string  with fixed length
func (p *payload) ReadStringWithFixLen(l int) []byte {
	b := p.b[p.off : p.off+l]
	p.off += l
	return b
}

//WriteInt1 write 1 byte into buffer
func (p *payload) WriteInt1(val uint8) {
	buffer := p.b[p.off:]
	p.off += 1
	WriteInt1(buffer, val)
}

func (p *payload) WriteInt2(val uint16) {
	buffer := p.b[p.off:]
	p.off += 2
	WriteInt2(buffer, val)
}

func (p *payload) WriteInt3(val uint32) {
	buffer := p.b[p.off:]
	p.off += 3
	WriteInt3(buffer, val)
}

func (p *payload) WriteInt4(val uint32) {
	buffer := p.b[p.off:]
	p.off += 4
	WriteInt4(buffer, val)
}

func (p *payload) WriteInt6(val uint64) {
	buffer := p.b[p.off:]
	p.off += 6
	WriteInt6(buffer, val)
}

func (p *payload) WriteInt8(val uint64) {
	buffer := p.b[p.off:]
	p.off += 8
	WriteInt8(buffer, val)
}

func (p *payload) WriteZero(l int) {
	p.off += l
}

func (p *payload) WriteStringWithNull(str []byte) {
	_ = copy(p.b[p.off:], str)
	p.off += len(str)
	p.WriteZero(1)
}

func (p *payload) WriteStringWithFixLen(str []byte) {
	_ = copy(p.b[p.off:], str)
	p.off += len(str)
}
