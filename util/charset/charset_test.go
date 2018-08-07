package charset

import (
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/encoding/korean"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/encoding/traditionalchinese"

	. "gopkg.in/check.v1"
)

type MyCharsetSuite struct{}

var _ = Suite(&MyCharsetSuite{})

func (s *MyCharsetSuite) TestCharset(c *C) {
	data := []string{
		"UTF8  ",
		"big5",
		"cp1250",
		"euckr",
		"  GBK  ",
		"sjis",
	}

	res := []encoding.Encoding{
		encoding.Nop,
		traditionalchinese.Big5,
		charmap.Windows1250,
		korean.EUCKR,
		simplifiedchinese.GBK,
		japanese.ShiftJIS,
	}

	for i := range data {
		charset, err := GetMethod(data[i])
		c.Assert(err, Equals, nil)
		c.Assert(charset, DeepEquals, res[i])
	}

}
