package character

import (
	"strings"

	"github.com/juju/errors"
	"golang.org/x/text/encoding"
)

//charSet is the default 
var CharSet = map[string]string{
	"big5":     "big5_chinese_ci",
	"dec8":     "dec8_swedish_ci",
	"cp850":    "cp850_general_ci",
	"hp8":      "hp8_english_ci",
	"koi8r":    "koi8r_general_ci",
	"latin1":   "latin1_swedish_ci",
	"latin2":   "latin2_general_ci",
	"swe7":     "swe7_swedish_ci",
	"ascii":    "ascii_general_ci",
	"ujis":     "ujis_japanese_ci",
	"sjis":     "sjis_japanese_ci",
	"hebrew":   "hebrew_general_ci",
	"tis620":   "tis620_thai_ci",
	"euckr":    "euckr_korean_ci",
	"koi8u":    "koi8u_general_ci",
	"gb2312":   "gb2312_chinese_ci",
	"greek":    "greek_general_ci",
	"cp1250":   "cp1250_general_ci",
	"gbk":      "gbk_chinese_ci",
	"latin5":   "latin5_turkish_ci",
	"armscii8": "armscii8_general_ci",
	"utf8":     "utf8_general_ci",
	"ucs2":     "ucs2_general_ci",
	"cp866":    "cp866_general_ci",
	"keybcs2":  "keybcs2_general_ci",
	"macce":    "macce_general_ci",
	"macroman": "macroman_general_ci",
	"cp852":    "cp852_general_ci",
	"latin7":   "latin7_general_ci",
	"utf8mb4":  "utf8mb4_general_ci",
	"cp1251":   "cp1251_general_ci",
	"utf16":    "utf16_general_ci",
	"utf16le":  "utf16le_general_ci",
	"cp1256":   "cp1256_general_ci",
	"cp1257":   "cp1257_general_ci",
	"utf32":    "utf32_general_ci",
	"binary":   "binary",
	"geostd8":  "geostd8_general_ci",
	"cp932":    "cp932_japanese_ci",
	"eucjpms":  "eucjpms_japanese_ci",
}

//SetDefaultCollation set the collation of a charset
func SetDefaultCollation(charset string,collation string) error {
	charset = strings.ToLower(strings.Trim(name, "\t\n\r\f "))
	collation = strings.ToLower(strings.Trim(name, "\t\n\r\f "))

	//check charset
	if _,ok:=CharSet[charset];!ok{
		return  errors.Errorf("character:Invalid character set:%s", charset)
	}

	//check collation
	if attr,ok:=Collations[collation];!ok {
		return errors.Errorf("character:Invalid collation :%s", collation)
	}else{
		if attr.charset!=charset{
			return errors.Errorf("character:Collation %s does not belong to CharSet:%s",collation,charset)
		}
	}

	CharSet[]

}


//GetEncoding return the Encoding instance
func GetEncoding(name string) (encoding.Encoding, error) {
	//check character name
	name = strings.ToLower(strings.Trim(name, "\t\n\r\f "))
	c, ok := charSet[name]
	if !ok {
		return nil, errors.Errorf("character:Invalid character set:%s", name)
	}

	//check collation name
	method, ok := collation[c]
	if !ok {
		return nil, errors.Errorf("character:Invalid collation :%s", c)
	}

	//check method
	if method == nil.(encoding.Encoding) {
		return nil, errors.Errorf("character:Unimplemented collation :%s", c)
	}

	return method, nil
}
