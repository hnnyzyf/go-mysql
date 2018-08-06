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
	//the collations of charset
	Collations map[string]*Collation
	//the default collation
	Default *Collation
}

//charsets reprensent the character sets and its method
var charsets = map[string]struct {
	method encoding.Encoding
}{
	"armscii8": {nil},
	"ascii":    {charmap.Windows1252},
	"big5":     {traditionalchinese.Big5},
	"binary":   {encoding.Nop},
	"cp1250":   {charmap.Windows1250},
	"cp1251":   {charmap.Windows1251},
	"cp1256":   {charmap.Windows1256},
	"cp1257":   {charmap.Windows1257},
	"cp850":    {charmap.CodePage850},
	"cp852":    {charmap.CodePage852},
	"cp866":    {charmap.CodePage866},
	"cp932":    {charmap.CodePage932},
	"dec8":     {nil},
	"eucjpms":  {nil},
	"euckr":    {korean.EUCKR},
	"gb18030":  {simplifiedchinese.GB18030},
	"gb2312":   {simplifiedchinese.HZGB2312},
	"gbk":      {simplifiedchinese.GBK},
	"geostd8":  {nil},
	"greek":    {charmap.ISO8859_7},
	"hebrew":   {charmap.ISO8859_8},
	"hp8":      {nil},
	"keybcs2":  {nil},
	"koi8r":    {charmap.KOI8R},
	"koi8u":    {charmap.KOI8U},
	"latin1":   {charmap.Windows1252},
	"latin2":   {charmap.ISO8859_2},
	"latin5":   {charmap.ISO8859_9},
	"latin7":   {charmap.ISO8859_13},
	"macce":    {nil},
	"macroman": {nil},
	"sjis":     {japanese.ShiftJIS},
	"swe7":     {nil},
	"tis620":   {charmap.Windows874},
	"ucs2":     {nil},
	"ujis":     {japanese.EUCJP},
	"utf16":    {unicode.UTF16(unicode.LittleEndian, unicode.IgnoreBOM)},
	"utf16le":  {unicode.UTF16(unicode.LittleEndian, unicode.IgnoreBOM)},
	"utf32":    {utf32.UTF32(utf32.BigEndian, utf32.IgnoreBOM)},
	"utf8":     {encoding.Nop},
	"utf8mb4":  {encoding.Nop},
}

//GetMethod return the encoding and decoding method of a charset
func GetMethod(name string) (encoding.Encoding, error) {
	charset := strings.Lower(strings.Trim(name, "\n\t\f"))
	n, ok := charsets[charset]
	if !ok {
		return nil, errors.Errorf("Charset:Invalid charset name:%d", charset)
	}

	if n == nil.(encoding.Encoding) {
		return nil, errors.Errorf("Charset:Unsupport charset:%d", charset)
	}

	return n, nil
}
