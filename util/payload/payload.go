package payload

type size_t = int

//payload represent a buffer in a packet
type payload struct {
	b   []byte
	off size_t
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
	return uint8(buffer[0])
}

//ReadInt2 reads 2 bytes
func (p *payload) ReadInt2() uint16 {
	buffer := p.b[p.off:]
	p.off += 2
	b := uint16(buffer[0])
	var i uint8 = 1
	for i < 2 {
		b = b | uint16(buffer[i])<<(i*8)
		i++
	}
	return b
}

//ReadInt3 reads 3 bytes
func (p *payload) ReadInt3() uint32 {
	buffer := p.b[p.off:]
	p.off += 3
	b := uint32(buffer[0])
	var i uint8 = 1
	for i < 3 {
		b = b | uint32(buffer[i])<<(i*8)
		i++
	}
	return b
}

//ReadInt4 reads 4 bytes
func (p *payload) ReadInt4() uint32 {
	buffer := p.b[p.off:]
	p.off += 4
	b := uint32(buffer[0])
	var i uint8 = 1
	for i < 4 {
		b = b | uint32(buffer[i])<<(i*8)
		i++
	}
	return b
}

//ReadInt6 reads 6 bytes
func (p *payload) ReadInt6() uint64 {
	buffer := p.b[p.off:]
	p.off += 6
	b := uint64(buffer[0])
	var i uint8 = 1
	for i < 6 {
		b = b | uint64(buffer[i])<<(i*8)
		i++
	}
	return b
}

//ReadInt8 reads 8 bytes
func (p *payload) ReadInt8() uint64 {
	buffer := p.b[p.off:]
	p.off += 8
	b := uint64(buffer[0])
	var i uint8 = 1
	for i < 8 {
		b = b | uint64(buffer[i])<<(i*8)
		i++
	}
	return b
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

func (p *payload) WriteInt1(val uint8) {
	p.b[p.off] = val
	p.off += 1
}

func (p *payload) WriteInt2(val uint16) {
	p.b[p.off] = uint8(val)
	p.b[p.off+1] = uint8(val >> 8)
	p.off += 2
}

func (p *payload) WriteInt3(val uint32) {
	p.b[p.off] = uint8(val)
	p.b[p.off+1] = uint8(val >> 8)
	p.b[p.off+2] = uint8(val >> 16)
	p.off += 3
}

func (p *payload) WriteInt4(val uint32) {
	p.b[p.off] = uint8(val)
	p.b[p.off+1] = uint8(val >> 8)
	p.b[p.off+2] = uint8(val >> 16)
	p.b[p.off+3] = uint8(val >> 24)
	p.off += 4
}

func (p *payload) WriteInt6(val uint64) {
	p.b[p.off] = uint8(val)
	p.b[p.off+1] = uint8(val >> 8)
	p.b[p.off+2] = uint8(val >> 16)
	p.b[p.off+3] = uint8(val >> 24)
	p.b[p.off+4] = uint8(val >> 32)
	p.b[p.off+5] = uint8(val >> 40)
	p.off += 6
}

func (p *payload) WriteInt8(val uint64) {
	p.b[p.off] = uint8(val)
	p.b[p.off+1] = uint8(val >> 8)
	p.b[p.off+2] = uint8(val >> 16)
	p.b[p.off+3] = uint8(val >> 24)
	p.b[p.off+4] = uint8(val >> 32)
	p.b[p.off+5] = uint8(val >> 40)
	p.b[p.off+6] = uint8(val >> 48)
	p.b[p.off+7] = uint8(val >> 56)
	p.off += 8
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
