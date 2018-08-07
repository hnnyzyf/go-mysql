package charset

import (
	"strings"

	"github.com/juju/errors"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/encoding/korean"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/encoding/traditionalchinese"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/encoding/unicode/utf32"
)

//CharSet represent a collection of collations
type CharSet struct {
	//the name of charset
	Name string

	ID uint8
	//the collations of charset
	Collations map[string]*Collation
	//the default collation
	Default *Collation
	//the method of encoding and decoding
	Method encoding.Encoding
}

//charsets reprensent the character sets and its method
var charsets = map[string]*CharSet{
	"armscii8": {"armscii8", 0x00, nil, nil, nil},
	"ascii":    {"ascii", 0x00, nil, nil, charmap.Windows1252},
	"big5":     {"big5", 0x00, nil, nil, traditionalchinese.Big5},
	"binary":   {"binary", 0x00, nil, nil, encoding.Nop},
	"cp1250":   {"cp1250", 0x00, nil, nil, charmap.Windows1250},
	"cp1251":   {"cp1251", 0x00, nil, nil, charmap.Windows1251},
	"cp1256":   {"cp1256", 0x00, nil, nil, charmap.Windows1256},
	"cp1257":   {"cp1257", 0x00, nil, nil, charmap.Windows1257},
	"cp850":    {"cp850", 0x00, nil, nil, charmap.CodePage850},
	"cp852":    {"cp852", 0x00, nil, nil, charmap.CodePage852},
	"cp866":    {"cp866", 0x00, nil, nil, charmap.CodePage866},
	"cp932":    {"cp932", 0x00, nil, nil, nil},
	"dec8":     {"dec8", 0x00, nil, nil, nil},
	"eucjpms":  {"eucjpms", 0x00, nil, nil, nil},
	"euckr":    {"euckr", 0x00, nil, nil, korean.EUCKR},
	"gb18030":  {"gb18030", 0x00, nil, nil, simplifiedchinese.GB18030},
	"gb2312":   {"gb2312", 0x00, nil, nil, simplifiedchinese.HZGB2312},
	"gbk":      {"gbk", 0x00, nil, nil, simplifiedchinese.GBK},
	"geostd8":  {"geostd8", 0x00, nil, nil, nil},
	"greek":    {"greek", 0x00, nil, nil, charmap.ISO8859_7},
	"hebrew":   {"hebrew", 0x00, nil, nil, charmap.ISO8859_8},
	"hp8":      {"hp8", 0x00, nil, nil, nil},
	"keybcs2":  {"keybcs2", 0x00, nil, nil, nil},
	"koi8r":    {"koi8r", 0x00, nil, nil, charmap.KOI8R},
	"koi8u":    {"koi8u", 0x00, nil, nil, charmap.KOI8U},
	"latin1":   {"latin1", 0x00, nil, nil, charmap.Windows1252},
	"latin2":   {"latin2", 0x00, nil, nil, charmap.ISO8859_2},
	"latin5":   {"latin5", 0x00, nil, nil, charmap.ISO8859_9},
	"latin7":   {"latin7", 0x00, nil, nil, charmap.ISO8859_13},
	"macce":    {"macce", 0x00, nil, nil, charmap.Macintosh},
	"macroman": {"macroman", 0x00, nil, nil, charmap.Macintosh},
	"sjis":     {"sjis", 0x00, nil, nil, japanese.ShiftJIS},
	"swe7":     {"swe7", 0x00, nil, nil, nil},
	"tis620":   {"tis620", 0x00, nil, nil, charmap.Windows874},
	"ucs2":     {"ucs2", 0x00, nil, nil, nil},
	"ujis":     {"ujis", 0x00, nil, nil, japanese.EUCJP},
	"utf16be":  {"utf16be", 0x00, nil, nil, unicode.UTF16(unicode.BigEndian, unicode.IgnoreBOM)},
	"utf16le":  {"utf16le", 0x00, nil, nil, unicode.UTF16(unicode.LittleEndian, unicode.IgnoreBOM)},
	"utf32be":  {"utf32be", 0x00, nil, nil, utf32.UTF32(utf32.BigEndian, utf32.IgnoreBOM)},
	"utf32le":  {"utf32le", 0x00, nil, nil, utf32.UTF32(utf32.LittleEndian, utf32.IgnoreBOM)},
	"utf8":     {"utf8", 0x21, nil, nil, encoding.Nop},
	"utf8mb4":  {"utf8mb4", 0x00, nil, nil, encoding.Nop},
}

//GetMethod return the encoding and decoding method of a charset
func GetMethod(name string) (encoding.Encoding, error) {
	charset := strings.ToLower(strings.Trim(name, "\t\n\r\f "))
	n, ok := charsets[charset]
	if !ok {
		return nil, errors.Errorf("Charset:Invalid charset name:%d", charset)
	}

	method := n.Method
	if method == encoding.Encoding(nil) {
		return nil, errors.Errorf("Charset:Unsupport charset:%d", charset)
	}
	return method, nil
}

//GetMethod return the encoding and decoding method of a charset
func GetID(name string) (uint8, error) {
	charset := strings.ToLower(strings.Trim(name, "\t\n\r\f "))
	n, ok := charsets[charset]
	if !ok {
		return nil, errors.Errorf("Charset:Invalid charset name:%d", charset)
	}

	return n.ID, nil
}
