package binary

import (
	"io"

	"github.com/juju/errors"
)

var errNoEnoughData = errors.Errorf("Cache:fail to copy data from buffer")

//the default cache size is 4KB
const defaultBufferSize = 4 * 1024

//Cache is a io.Reader with buffer to decrease the frequency of syscall
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
		b:  make([]byte, defaultBufferSize),
		fd: r,
	}
}

func NewCacheWithSize(r io.Reader, size int) *Cache {
	return &Cache{
		b:  make([]byte, size),
		fd: r,
	}
}

//Read copy data from buffer,if no avaliable data,try to copy from io.Reader
//	case1:avali == 0 and size > len(c.b) read from io.reader direct and fill the buffer
//  case2:avali == 0 and size <=len(c.b) fill the buffer and copy directly
//  case3:avali > 0 and size > avali copy remaining and read others from io.Reader
//  case4:avali >0 and size <- avali copy remaining
func (c *Cache) Read(size int) ([]byte, error) {
	buffer := make([]byte, size)
	//case 1
	if c.avail == 0 && len(buffer) > len(c.b) {
		//read
		if err := c.read(buffer); err != nil {
			return nil, errors.Trace(err)
		}
		//fill buffer
		if err := c.fill(); err != nil {
			return nil, errors.Trace(err)
		} else {
			return buffer, nil
		}
	} else if c.avail == 0 && len(buffer) <= len(c.b) {
		//fill buffer
		if err := c.fill(); err != nil {
			return nil, errors.Trace(err)
		}
		//copy data
		if n := copy(buffer, c.b); n != len(buffer) {
			return nil, errNoEnoughData
		} else {
			//fmt.Println("size", len(buffer), n, c.off, c.avail)
			c.off += n
			c.avail -= n
			return buffer, nil
		}
	} else if c.avail > 0 && len(buffer) > c.avail {
		//copy data
		if n := copy(buffer, c.b[c.off:c.off+c.avail]); n != c.avail {
			return nil, errNoEnoughData
		} else {
			//read from io.reader
			if err := c.read(buffer[c.avail:]); err != nil {
				return nil, errors.Trace(err)
			} else {
				c.off += c.avail
				c.avail -= c.avail
				return buffer, nil
			}
		}
	} else if c.avail > 0 && len(buffer) <= c.avail {
		//copy data
		if n := copy(buffer, c.b[c.off:c.off+len(buffer)]); n != len(buffer) {
			return nil, errNoEnoughData
		} else {
			c.off += n
			c.avail -= n
			return buffer, nil
		}
	} else {
		panic("binary:impossible cache state")
	}

}

//Next return a byte slice according to the size
//you need to use it immeddiately because Next does not guarantee the buffer is safe
//	case1:avali == 0 and size > len(c.b) read from io.reader direct
//  case2:avali == 0 and size <=len(c.b) fill the buffer and return
//  case3:avali > 0 and size > avali read remaining and read others from io.Reader
//  case4:avali >0 and size <- avali read remaining
func (c *Cache) Next(size int) ([]byte, error) {
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
		return c.Next(size)
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
	if r, ok := c.fd.(io.Closer); ok {
		r.Close()
	}
}
