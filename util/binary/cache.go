package binary

import (
	//"fmt"
	"github.com/juju/errors"
	"io"
)

//the default cache size is 4KB
const defaultBufferSize = 4 * 1024

//Cache is a io.Reader with buffer
type Cache struct {
	//buffer represent the buffer
	b []byte

	//fd reprensents a io.Reader
	fd io.Reader

	//off represent the offset in b
	//avail represent the avail buffer
	off, avail int
}

func NewCache(r io.Reader) *Cache {
	return &Cache{
		b:   make([]byte, defaultBufferSize),
		fd:  r,
		off: 0,
	}
}

func NewCacheWithSize(r io.Reader, size int) *Cache {
	return &Cache{
		b:   make([]byte, size),
		fd:  r,
		off: 0,
	}
}

//Read return a byte slice according to the size
//	case1:avali == 0 and size > len(c.b) read from io.reader direct
//  case2:avali == 0 and size <=len(c.b) fill the buffer and return
//  case3:avali > 0 and size > avali read remaining and read others from io.Reader
//  case4:avali >0 and size <- avali read remaining
func (c *Cache) Read(size int) ([]byte, error) {
	//case 1
	if c.avail == 0 && size > len(c.b) {
		//init buffer
		buffer := make([]byte, size)
		//read from io.reader direct
		if err := c.read(buffer); err != nil {
			return nil, errors.Trace(err)
		} else {
			return buffer, nil
		}
		//case 2
	} else if c.avail == 0 && size <= len(c.b) {
		if err := c.fill(); err != nil {
			return nil, err
		}
		//read
		return c.Read(size)
		//case 3
	} else if c.avail > 0 && size > c.avail {
		buffer := make([]byte, size)
		//copy remaining to buffer
		avali := copy(buffer, c.b[c.off:c.off+c.avail])
		if avali != c.avail {
			return nil, errors.Errorf("binary:only copy %d bytes from cache buffer,expect %d bytes", avali, c.avail)
		}
		//read from io.reader
		if err := c.read(buffer[c.avail:]); err != nil {
			return nil, errors.Trace(err)
		} else {
			c.off += avali
			c.avail -= avali
			return buffer, nil
		}
		//case 4
	} else if c.avail > 0 && size <= c.avail {
		buffer := c.b[c.off : c.off+size]
		c.off += size
		c.avail -= size
		return buffer, nil
	} else {
		panic("binary:impossible state in cache reading")
	}

}

//read reads from io.reader
func (c *Cache) read(buffer []byte) error {
	n, err := c.fd.Read(buffer)
	//occurs error
	if err != nil {
		return errors.Trace(err)
	} else if n != len(buffer) {
		return errors.Errorf("binary:no enough bytes to read in cache")
	} else {
		return nil
	}
}

//fill try to fill the buffer
func (c *Cache) fill() error {
	//fill the buffer
	if n, err := c.fd.Read(c.b); err != nil {
		return errors.Trace(err)
	} else {
		c.avail = n
		c.off = 0
		return nil
	}
}

//Close try to close the io.Reader if it is a io.Closer
func (c *Cache) Close() {
	if r, ok := c.(io.Closer); ok {
		r.Close()
	}
}
