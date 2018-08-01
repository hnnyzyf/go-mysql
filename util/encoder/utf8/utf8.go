package utf8

import (
	"unicode/utf8"
)

//utf8 only support 3 bytes utf code
type utf8Encoder struct {
	buffer []byte
	offset int
}

func NewEncoder(s []byte) *utf8Encoder {
	return &utf8Encoder{
		buffer: s,
		offset: 0,
	}
}

//ReadRune implement Encoder interface
func (u *utf8Encoder) ReadRune() ([]byte, error) {
	//return when EOF
	if u.offset >= len(u.buffer) {
		return nil, io.EOF
	}

	r := u.buffer[u.offset]
	n := 0
	switch i := uint8(r); {
	case i <= 127:
		n = 1
	case i <= 223:
		n = 2
	case i <= 239:
		n = 3
	default:
		n = 4
	}

	if u.offset+n > len(u.buffer) {
		n = len(u.buffer) - u.offset
		u.offset = len(u.buffer)
	} else {
		u.offset += n
	}

	return u.buffer[u.offset-n : u.offset], nil
}
