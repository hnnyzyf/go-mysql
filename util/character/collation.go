package character

import (
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/encoding/korean"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/encoding/traditionalchinese"
	"golang.org/x/text/encoding/unicode"
)

const (
	big5_chinese_ci          uint8 = 0x01
	latin2_czech_cs          uint8 = 0x02
	dec8_swedish_ci          uint8 = 0x03
	cp850_general_ci         uint8 = 0x04
	latin1_german1_ci        uint8 = 0x05
	hp8_english_ci           uint8 = 0x06
	koi8r_general_ci         uint8 = 0x07
	latin1_swedish_ci        uint8 = 0x08
	latin2_general_ci        uint8 = 0x09
	swe7_swedish_ci          uint8 = 0x0A
	ascii_general_ci         uint8 = 0x0B
	ujis_japanese_ci         uint8 = 0x0C
	sjis_japanese_ci         uint8 = 0x0D
	cp1251_bulgarian_ci      uint8 = 0x0E
	latin1_danish_ci         uint8 = 0x0F
	hebrew_general_ci        uint8 = 0x10
	unknown_general_ci       uint8 = 0x11
	tis620_thai_ci           uint8 = 0x12
	euckr_korean_ci          uint8 = 0x13
	latin7_estonian_cs       uint8 = 0x14
	latin2_hungarian_ci      uint8 = 0x15
	koi8u_general_ci         uint8 = 0x16
	cp1251_ukrainian_ci      uint8 = 0x17
	gb2312_chinese_ci        uint8 = 0x18
	greek_general_ci         uint8 = 0x19
	cp1250_general_ci        uint8 = 0x1A
	latin2_croatian_ci       uint8 = 0x1B
	gbk_chinese_ci           uint8 = 0x1C
	cp1257_lithuanian_ci     uint8 = 0x1D
	latin5_turkish_ci        uint8 = 0x1E
	latin1_german2_ci        uint8 = 0x1F
	armscii8_general_ci      uint8 = 0x20
	utf8_general_ci          uint8 = 0x21
	cp1250_czech_cs          uint8 = 0x22
	ucs2_general_ci          uint8 = 0x23
	cp866_general_ci         uint8 = 0x24
	keybcs2_general_ci       uint8 = 0x25
	macce_general_ci         uint8 = 0x26
	macroman_general_ci      uint8 = 0x27
	cp852_general_ci         uint8 = 0x28
	latin7_general_ci        uint8 = 0x29
	latin7_general_cs        uint8 = 0x2A
	macce_bin                uint8 = 0x2B
	cp1250_croatian_ci       uint8 = 0x2C
	utf8mb4_general_ci       uint8 = 0x2D
	utf8mb4_bin              uint8 = 0x2E
	latin1_bin               uint8 = 0x2F
	latin1_general_ci        uint8 = 0x30
	latin1_general_cs        uint8 = 0x31
	cp1251_bin               uint8 = 0x32
	cp1251_general_ci        uint8 = 0x33
	cp1251_general_cs        uint8 = 0x34
	macroman_bin             uint8 = 0x35
	utf16_general_ci         uint8 = 0x36
	utf16_bin                uint8 = 0x37
	utf16le_general_ci       uint8 = 0x38
	cp1256_general_ci        uint8 = 0x39
	cp1257_bin               uint8 = 0x3A
	cp1257_general_ci        uint8 = 0x3B
	utf32_general_ci         uint8 = 0x3C
	utf32_bin                uint8 = 0x3D
	utf16le_bin              uint8 = 0x3E
	binary                   uint8 = 0x3F
	armscii8_bin             uint8 = 0x40
	ascii_bin                uint8 = 0x41
	cp1250_bin               uint8 = 0x42
	cp1256_bin               uint8 = 0x43
	cp866_bin                uint8 = 0x44
	dec8_bin                 uint8 = 0x45
	greek_bin                uint8 = 0x46
	hebrew_bin               uint8 = 0x47
	hp8_bin                  uint8 = 0x48
	keybcs2_bin              uint8 = 0x49
	koi8r_bin                uint8 = 0x4A
	koi8u_bin                uint8 = 0x4B
	latin2_bin               uint8 = 0x4C
	latin5_bin               uint8 = 0x4D
	latin7_bin               uint8 = 0x4E
	cp850_bin                uint8 = 0x4F
	cp852_bin                uint8 = 0x50
	swe7_bin                 uint8 = 0x51
	utf8_bin                 uint8 = 0x52
	big5_bin                 uint8 = 0x53
	euckr_bin                uint8 = 0x54
	gb2312_bin               uint8 = 0x55
	gbk_bin                  uint8 = 0x56
	sjis_bin                 uint8 = 0x57
	tis620_bin               uint8 = 0x58
	ucs2_bin                 uint8 = 0x59
	ujis_bin                 uint8 = 0x5A
	geostd8_general_ci       uint8 = 0x5B
	geostd8_bin              uint8 = 0x5C
	latin1_spanish_ci        uint8 = 0x5D
	cp932_japanese_ci        uint8 = 0x5E
	cp932_bin                uint8 = 0x5F
	eucjpms_japanese_ci      uint8 = 0x60
	eucjpms_bin              uint8 = 0x61
	cp1250_polish_ci         uint8 = 0x62
	utf16_unicode_ci         uint8 = 0x63
	utf16_icelandic_ci       uint8 = 0x64
	utf16_latvian_ci         uint8 = 0x65
	utf16_romanian_ci        uint8 = 0x66
	utf16_slovenian_ci       uint8 = 0x67
	utf16_polish_ci          uint8 = 0x68
	utf16_estonian_ci        uint8 = 0x69
	utf16_spanish_ci         uint8 = 0x6A
	utf16_swedish_ci         uint8 = 0x6B
	utf16_turkish_ci         uint8 = 0x6C
	utf16_czech_ci           uint8 = 0x6D
	utf16_danish_ci          uint8 = 0x6E
	utf16_lithuanian_ci      uint8 = 0x6F
	utf16_slovak_ci          uint8 = 0x70
	utf16_spanish2_ci        uint8 = 0x71
	utf16_roman_ci           uint8 = 0x72
	utf16_persian_ci         uint8 = 0x73
	utf16_esperanto_ci       uint8 = 0x74
	utf16_hungarian_ci       uint8 = 0x75
	utf16_sinhala_ci         uint8 = 0x76
	utf16_german2_ci         uint8 = 0x77
	utf16_croatian_ci        uint8 = 0x78
	utf16_unicode_520_ci     uint8 = 0x79
	utf16_vietnamese_ci      uint8 = 0x7A
	ucs2_unicode_ci          uint8 = 0x7B
	ucs2_icelandic_ci        uint8 = 0x7C
	ucs2_latvian_ci          uint8 = 0x7D
	ucs2_romanian_ci         uint8 = 0x7E
	ucs2_slovenian_ci        uint8 = 0x7F
	ucs2_polish_ci           uint8 = 0x80
	ucs2_estonian_ci         uint8 = 0x81
	ucs2_spanish_ci          uint8 = 0x82
	ucs2_swedish_ci          uint8 = 0x83
	ucs2_turkish_ci          uint8 = 0x84
	ucs2_czech_ci            uint8 = 0x85
	ucs2_danish_ci           uint8 = 0x86
	ucs2_lithuanian_ci       uint8 = 0x87
	ucs2_slovak_ci           uint8 = 0x88
	ucs2_spanish2_ci         uint8 = 0x89
	ucs2_roman_ci            uint8 = 0x8A
	ucs2_persian_ci          uint8 = 0x8B
	ucs2_esperanto_ci        uint8 = 0x8C
	ucs2_hungarian_ci        uint8 = 0x8D
	ucs2_sinhala_ci          uint8 = 0x8E
	ucs2_german2_ci          uint8 = 0x8F
	ucs2_croatian_ci         uint8 = 0x90
	ucs2_unicode_520_ci      uint8 = 0x91
	ucs2_vietnamese_ci       uint8 = 0x92
	ucs2_general_mysql500_ci uint8 = 0x93
	utf32_unicode_ci         uint8 = 0x94
	utf32_icelandic_ci       uint8 = 0x95
	utf32_latvian_ci         uint8 = 0x96
	utf32_romanian_ci        uint8 = 0x97
	utf32_slovenian_ci       uint8 = 0x98
	utf32_polish_ci          uint8 = 0x99
	utf32_estonian_ci        uint8 = 0x9A
	utf32_spanish_ci         uint8 = 0x9B
	utf32_swedish_ci         uint8 = 0x9C
	utf32_turkish_ci         uint8 = 0x9D
	utf32_czech_ci           uint8 = 0x9E
	utf32_danish_ci          uint8 = 0x9F
	utf32_lithuanian_ci      uint8 = 0xA0
	utf32_slovak_ci          uint8 = 0xA1
	utf32_spanish2_ci        uint8 = 0xA2
	utf32_roman_ci           uint8 = 0xA3
	utf32_persian_ci         uint8 = 0xA4
	utf32_esperanto_ci       uint8 = 0xA5
	utf32_hungarian_ci       uint8 = 0xA6
	utf32_sinhala_ci         uint8 = 0xA7
	utf32_german2_ci         uint8 = 0xA8
	utf32_croatian_ci        uint8 = 0xA9
	utf32_unicode_520_ci     uint8 = 0xAA
	utf32_vietnamese_ci      uint8 = 0xAB
	utf8_unicode_ci          uint8 = 0xAC
	utf8_icelandic_ci        uint8 = 0xAD
	utf8_latvian_ci          uint8 = 0xAE
	utf8_romanian_ci         uint8 = 0xAF
	utf8_slovenian_ci        uint8 = 0xB0
	utf8_polish_ci           uint8 = 0xB1
	utf8_estonian_ci         uint8 = 0xB2
	utf8_spanish_ci          uint8 = 0xB3
	utf8_swedish_ci          uint8 = 0xB4
	utf8_turkish_ci          uint8 = 0xB5
	utf8_czech_ci            uint8 = 0xB6
	utf8_danish_ci           uint8 = 0xB7
	utf8_lithuanian_ci       uint8 = 0xB8
	utf8_slovak_ci           uint8 = 0xB9
	utf8_spanish2_ci         uint8 = 0xBA
	utf8_roman_ci            uint8 = 0xBB
	utf8_persian_ci          uint8 = 0xBC
	utf8_esperanto_ci        uint8 = 0xBD
	utf8_hungarian_ci        uint8 = 0xBE
	utf8_sinhala_ci          uint8 = 0xBF
	utf8_german2_ci          uint8 = 0xC0
	utf8_croatian_ci         uint8 = 0xC1
	utf8_unicode_520_ci      uint8 = 0xC2
	utf8_vietnamese_ci       uint8 = 0xC3
	utf8_general_mysql500_ci uint8 = 0xC4
	utf8mb4_unicode_ci       uint8 = 0xC5
	utf8mb4_icelandic_ci     uint8 = 0xC6
	utf8mb4_latvian_ci       uint8 = 0xC7
	utf8mb4_romanian_ci      uint8 = 0xC8
	utf8mb4_slovenian_ci     uint8 = 0xC9
	utf8mb4_polish_ci        uint8 = 0xCA
	utf8mb4_estonian_ci      uint8 = 0xCB
	utf8mb4_spanish_ci       uint8 = 0xCC
	utf8mb4_swedish_ci       uint8 = 0xCD
	utf8mb4_turkish_ci       uint8 = 0xCE
	utf8mb4_czech_ci         uint8 = 0xCF
	utf8mb4_danish_ci        uint8 = 0xD0
	utf8mb4_lithuanian_ci    uint8 = 0xD1
	utf8mb4_slovak_ci        uint8 = 0xD2
	utf8mb4_spanish2_ci      uint8 = 0xD3
	utf8mb4_roman_ci         uint8 = 0xD4
	utf8mb4_persian_ci       uint8 = 0xD5
	utf8mb4_esperanto_ci     uint8 = 0xD6
	utf8mb4_hungarian_ci     uint8 = 0xD7
	utf8mb4_sinhala_ci       uint8 = 0xD8
	utf8mb4_german2_ci       uint8 = 0xD9
	utf8mb4_croatian_ci      uint8 = 0xDA
	utf8mb4_unicode_520_ci   uint8 = 0xD1
	utf8mb4_vietnamese_ci    uint8 = 0xD2
)

//collations store the id and charset of a collation
var Collations = map[string]struct {
	id      uint8
	charset string
}{
	"big5_chinese_ci":          {big5_chinese_ci, nil},
	"latin2_czech_cs":          {latin2_czech_cs, charmap.ISO8859_2},
	"dec8_swedish_ci":          {dec8_swedish_ci, nil},
	"cp850_general_ci":         {cp850_general_ci, nil},
	"latin1_german1_ci":        {latin1_german1_ci, charmap.ISO8859_1},
	"hp8_english_ci":           {hp8_english_ci, nil},
	"koi8r_general_ci":         {koi8r_general_ci, charmap.KOI8R},
	"latin1_swedish_ci":        {latin1_swedish_ci, charmap.ISO8859_1},
	"latin2_general_ci":        {latin2_general_ci, charmap.ISO8859_2},
	"swe7_swedish_ci":          {swe7_swedish_ci, nil},
	"ascii_general_ci":         {ascii_general_ci, nil},
	"ujis_japanese_ci":         {ujis_japanese_ci, japanese.EUCJP},
	"sjis_japanese_ci":         {sjis_japanese_ci, japanese.ShiftJIS},
	"cp1251_bulgarian_ci":      {cp1251_bulgarian_ci, nil},
	"latin1_danish_ci":         {latin1_danish_ci, charmap.ISO8859_1},
	"hebrew_general_ci":        {hebrew_general_ci, charmap.ISO8859_8},
	"unknown_general_ci":       {unknown_general_ci, nil},
	"tis620_thai_ci":           {tis620_thai_ci, charmap.Windows874},
	"euckr_korean_ci":          {euckr_korean_ci, nil},
	"latin7_estonian_cs":       {latin7_estonian_cs, charmap.ISO8859_13},
	"latin2_hungarian_ci":      {latin2_hungarian_ci, charmap.ISO8859_2},
	"koi8u_general_ci":         {koi8u_general_ci, charmap.KOI8U},
	"cp1251_ukrainian_ci":      {cp1251_ukrainian_ci, nil},
	"gb2312_chinese_ci":        {gb2312_chinese_ci, nil},
	"greek_general_ci":         {greek_general_ci, charmap.ISO8859_7},
	"cp1250_general_ci":        {cp1250_general_ci, nil},
	"latin2_croatian_ci":       {latin2_croatian_ci, charmap.ISO8859_2},
	"gbk_chinese_ci":           {gbk_chinese_ci, nil},
	"cp1257_lithuanian_ci":     {cp1257_lithuanian_ci, nil},
	"latin5_turkish_ci":        {latin5_turkish_ci, charmap.ISO8859_9},
	"latin1_german2_ci":        {latin1_german2_ci, charmap.ISO8859_1},
	"armscii8_general_ci":      {armscii8_general_ci, nil},
	"utf8_general_ci":          {utf8_general_ci, nil},
	"cp1250_czech_cs":          {cp1250_czech_cs, nil},
	"ucs2_general_ci":          {ucs2_general_ci, nil},
	"cp866_general_ci":         {cp866_general_ci, nil},
	"keybcs2_general_ci":       {keybcs2_general_ci, nil},
	"macce_general_ci":         {macce_general_ci, nil},
	"macroman_general_ci":      {macroman_general_ci, nil},
	"cp852_general_ci":         {cp852_general_ci, nil},
	"latin7_general_ci":        {latin7_general_ci, charmap.ISO8859_13},
	"latin7_general_cs":        {latin7_general_cs, charmap.ISO8859_13},
	"macce_bin":                {macce_bin, nil},
	"cp1250_croatian_ci":       {cp1250_croatian_ci, nil},
	"utf8mb4_general_ci":       {utf8mb4_general_ci, nil},
	"utf8mb4_bin":              {utf8mb4_bin, nil},
	"latin1_bin":               {latin1_bin, charmap.ISO8859_1},
	"latin1_general_ci":        {latin1_general_ci, charmap.ISO8859_1},
	"latin1_general_cs":        {latin1_general_cs, charmap.ISO8859_1},
	"cp1251_bin":               {cp1251_bin, nil},
	"cp1251_general_ci":        {cp1251_general_ci, nil},
	"cp1251_general_cs":        {cp1251_general_cs, nil},
	"macroman_bin":             {macroman_bin, nil},
	"utf16_general_ci":         {utf16_general_ci, nil},
	"utf16_bin":                {utf16_bin, nil},
	"utf16le_general_ci":       {utf16le_general_ci, nil},
	"cp1256_general_ci":        {cp1256_general_ci, nil},
	"cp1257_bin":               {cp1257_bin, nil},
	"cp1257_general_ci":        {cp1257_general_ci, nil},
	"utf32_general_ci":         {utf32_general_ci, nil},
	"utf32_bin":                {utf32_bin, nil},
	"utf16le_bin":              {utf16le_bin, nil},
	"binary":                   {binary, nil},
	"armscii8_bin":             {armscii8_bin, nil},
	"ascii_bin":                {ascii_bin, nil},
	"cp1250_bin":               {cp1250_bin, nil},
	"cp1256_bin":               {cp1256_bin, nil},
	"cp866_bin":                {cp866_bin, nil},
	"dec8_bin":                 {dec8_bin, nil},
	"greek_bin":                {greek_bin, charmap.ISO8859_7},
	"hebrew_bin":               {hebrew_bin, charmap.ISO8859_8},
	"hp8_bin":                  {hp8_bin, nil},
	"keybcs2_bin":              {keybcs2_bin, nil},
	"koi8r_bin":                {koi8r_bin, nil},
	"koi8u_bin":                {koi8u_bin, nil},
	"latin2_bin":               {latin2_bin, charmap.ISO8859_2},
	"latin5_bin":               {latin5_bin, charmap.ISO8859_9},
	"latin7_bin":               {latin7_bin, charmap.ISO8859_13},
	"cp850_bin":                {cp850_bin, nil},
	"cp852_bin":                {cp852_bin, nil},
	"swe7_bin":                 {swe7_bin, nil},
	"utf8_bin":                 {utf8_bin, nil},
	"big5_bin":                 {big5_bin, nil},
	"euckr_bin":                {euckr_bin, nil},
	"gb2312_bin":               {gb2312_bin, nil},
	"gbk_bin":                  {gbk_bin, nil},
	"sjis_bin":                 {sjis_bin, nil},
	"tis620_bin":               {tis620_bin, charmap.Windows874},
	"ucs2_bin":                 {ucs2_bin, nil},
	"ujis_bin":                 {ujis_bin, nil},
	"geostd8_general_ci":       {geostd8_general_ci, nil},
	"geostd8_bin":              {geostd8_bin, nil},
	"latin1_spanish_ci":        {latin1_spanish_ci, charmap.ISO8859_1},
	"cp932_japanese_ci":        {cp932_japanese_ci, nil},
	"cp932_bin":                {cp932_bin, nil},
	"eucjpms_japanese_ci":      {eucjpms_japanese_ci, japanese.ISO2022JP},
	"eucjpms_bin":              {eucjpms_bin, japanese.ISO2022JP},
	"cp1250_polish_ci":         {cp1250_polish_ci, nil},
	"utf16_unicode_ci":         {utf16_unicode_ci, nil},
	"utf16_icelandic_ci":       {utf16_icelandic_ci, nil},
	"utf16_latvian_ci":         {utf16_latvian_ci, nil},
	"utf16_romanian_ci":        {utf16_romanian_ci, nil},
	"utf16_slovenian_ci":       {utf16_slovenian_ci, nil},
	"utf16_polish_ci":          {utf16_polish_ci, nil},
	"utf16_estonian_ci":        {utf16_estonian_ci, nil},
	"utf16_spanish_ci":         {utf16_spanish_ci, nil},
	"utf16_swedish_ci":         {utf16_swedish_ci, nil},
	"utf16_turkish_ci":         {utf16_turkish_ci, nil},
	"utf16_czech_ci":           {utf16_czech_ci, nil},
	"utf16_danish_ci":          {utf16_danish_ci, nil},
	"utf16_lithuanian_ci":      {utf16_lithuanian_ci, nil},
	"utf16_slovak_ci":          {utf16_slovak_ci, nil},
	"utf16_spanish2_ci":        {utf16_spanish2_ci, nil},
	"utf16_roman_ci":           {utf16_roman_ci, nil},
	"utf16_persian_ci":         {utf16_persian_ci, nil},
	"utf16_esperanto_ci":       {utf16_esperanto_ci, nil},
	"utf16_hungarian_ci":       {utf16_hungarian_ci, nil},
	"utf16_sinhala_ci":         {utf16_sinhala_ci, nil},
	"utf16_german2_ci":         {utf16_german2_ci, nil},
	"utf16_croatian_ci":        {utf16_croatian_ci, nil},
	"utf16_unicode_520_ci":     {utf16_unicode_520_ci, nil},
	"utf16_vietnamese_ci":      {utf16_vietnamese_ci, nil},
	"ucs2_unicode_ci":          {ucs2_unicode_ci, nil},
	"ucs2_icelandic_ci":        {ucs2_icelandic_ci, nil},
	"ucs2_latvian_ci":          {ucs2_latvian_ci, nil},
	"ucs2_romanian_ci":         {ucs2_romanian_ci, nil},
	"ucs2_slovenian_ci":        {ucs2_slovenian_ci, nil},
	"ucs2_polish_ci":           {ucs2_polish_ci, nil},
	"ucs2_estonian_ci":         {ucs2_estonian_ci, nil},
	"ucs2_spanish_ci":          {ucs2_spanish_ci, nil},
	"ucs2_swedish_ci":          {ucs2_swedish_ci, nil},
	"ucs2_turkish_ci":          {ucs2_turkish_ci, nil},
	"ucs2_czech_ci":            {ucs2_czech_ci, nil},
	"ucs2_danish_ci":           {ucs2_danish_ci, nil},
	"ucs2_lithuanian_ci":       {ucs2_lithuanian_ci, nil},
	"ucs2_slovak_ci":           {ucs2_slovak_ci, nil},
	"ucs2_spanish2_ci":         {ucs2_spanish2_ci, nil},
	"ucs2_roman_ci":            {ucs2_roman_ci, nil},
	"ucs2_persian_ci":          {ucs2_persian_ci, nil},
	"ucs2_esperanto_ci":        {ucs2_esperanto_ci, nil},
	"ucs2_hungarian_ci":        {ucs2_hungarian_ci, nil},
	"ucs2_sinhala_ci":          {ucs2_sinhala_ci, nil},
	"ucs2_german2_ci":          {ucs2_german2_ci, nil},
	"ucs2_croatian_ci":         {ucs2_croatian_ci, nil},
	"ucs2_unicode_520_ci":      {ucs2_unicode_520_ci, nil},
	"ucs2_vietnamese_ci":       {ucs2_vietnamese_ci, nil},
	"ucs2_general_mysql500_ci": {ucs2_general_mysql500_ci, nil},
	"utf32_unicode_ci":         {utf32_unicode_ci, nil},
	"utf32_icelandic_ci":       {utf32_icelandic_ci, nil},
	"utf32_latvian_ci":         {utf32_latvian_ci, nil},
	"utf32_romanian_ci":        {utf32_romanian_ci, nil},
	"utf32_slovenian_ci":       {utf32_slovenian_ci, nil},
	"utf32_polish_ci":          {utf32_polish_ci, nil},
	"utf32_estonian_ci":        {utf32_estonian_ci, nil},
	"utf32_spanish_ci":         {utf32_spanish_ci, nil},
	"utf32_swedish_ci":         {utf32_swedish_ci, nil},
	"utf32_turkish_ci":         {utf32_turkish_ci, nil},
	"utf32_czech_ci":           {utf32_czech_ci, nil},
	"utf32_danish_ci":          {utf32_danish_ci, nil},
	"utf32_lithuanian_ci":      {utf32_lithuanian_ci, nil},
	"utf32_slovak_ci":          {utf32_slovak_ci, nil},
	"utf32_spanish2_ci":        {utf32_spanish2_ci, nil},
	"utf32_roman_ci":           {utf32_roman_ci, nil},
	"utf32_persian_ci":         {utf32_persian_ci, nil},
	"utf32_esperanto_ci":       {utf32_esperanto_ci, nil},
	"utf32_hungarian_ci":       {utf32_hungarian_ci, nil},
	"utf32_sinhala_ci":         {utf32_sinhala_ci, nil},
	"utf32_german2_ci":         {utf32_german2_ci, nil},
	"utf32_croatian_ci":        {utf32_croatian_ci, nil},
	"utf32_unicode_520_ci":     {utf32_unicode_520_ci, nil},
	"utf32_vietnamese_ci":      {utf32_vietnamese_ci, nil},
	"utf8_unicode_ci":          {utf8_unicode_ci, nil},
	"utf8_icelandic_ci":        {utf8_icelandic_ci, nil},
	"utf8_latvian_ci":          {utf8_latvian_ci, nil},
	"utf8_romanian_ci":         {utf8_romanian_ci, nil},
	"utf8_slovenian_ci":        {utf8_slovenian_ci, nil},
	"utf8_polish_ci":           {utf8_polish_ci, nil},
	"utf8_estonian_ci":         {utf8_estonian_ci, nil},
	"utf8_spanish_ci":          {utf8_spanish_ci, nil},
	"utf8_swedish_ci":          {utf8_swedish_ci, nil},
	"utf8_turkish_ci":          {utf8_turkish_ci, nil},
	"utf8_czech_ci":            {utf8_czech_ci, nil},
	"utf8_danish_ci":           {utf8_danish_ci, nil},
	"utf8_lithuanian_ci":       {utf8_lithuanian_ci, nil},
	"utf8_slovak_ci":           {utf8_slovak_ci, nil},
	"utf8_spanish2_ci":         {utf8_spanish2_ci, nil},
	"utf8_roman_ci":            {utf8_roman_ci, nil},
	"utf8_persian_ci":          {utf8_persian_ci, nil},
	"utf8_esperanto_ci":        {utf8_esperanto_ci, nil},
	"utf8_hungarian_ci":        {utf8_hungarian_ci, nil},
	"utf8_sinhala_ci":          {utf8_sinhala_ci, nil},
	"utf8_german2_ci":          {utf8_german2_ci, nil},
	"utf8_croatian_ci":         {utf8_croatian_ci, nil},
	"utf8_unicode_520_ci":      {utf8_unicode_520_ci, nil},
	"utf8_vietnamese_ci":       {utf8_vietnamese_ci, nil},
	"utf8_general_mysql500_ci": {utf8_general_mysql500_ci, nil},
	"utf8mb4_unicode_ci":       {utf8mb4_unicode_ci, nil},
	"utf8mb4_icelandic_ci":     {utf8mb4_icelandic_ci, nil},
	"utf8mb4_latvian_ci":       {utf8mb4_latvian_ci, nil},
	"utf8mb4_romanian_ci":      {utf8mb4_romanian_ci, nil},
	"utf8mb4_slovenian_ci":     {utf8mb4_slovenian_ci, nil},
	"utf8mb4_polish_ci":        {utf8mb4_polish_ci, nil},
	"utf8mb4_estonian_ci":      {utf8mb4_estonian_ci, nil},
	"utf8mb4_spanish_ci":       {utf8mb4_spanish_ci, nil},
	"utf8mb4_swedish_ci":       {utf8mb4_swedish_ci, nil},
	"utf8mb4_turkish_ci":       {utf8mb4_turkish_ci, nil},
	"utf8mb4_czech_ci":         {utf8mb4_czech_ci, nil},
	"utf8mb4_danish_ci":        {utf8mb4_danish_ci, nil},
	"utf8mb4_lithuanian_ci":    {utf8mb4_lithuanian_ci, nil},
	"utf8mb4_slovak_ci":        {utf8mb4_slovak_ci, nil},
	"utf8mb4_spanish2_ci":      {utf8mb4_spanish2_ci, nil},
	"utf8mb4_roman_ci":         {utf8mb4_roman_ci, nil},
	"utf8mb4_persian_ci":       {utf8mb4_persian_ci, nil},
	"utf8mb4_esperanto_ci":     {utf8mb4_esperanto_ci, nil},
	"utf8mb4_hungarian_ci":     {utf8mb4_hungarian_ci, nil},
	"utf8mb4_sinhala_ci":       {utf8mb4_sinhala_ci, nil},
	"utf8mb4_german2_ci":       {utf8mb4_german2_ci, nil},
	"utf8mb4_croatian_ci":      {utf8mb4_croatian_ci, nil},
	"utf8mb4_unicode_520_ci":   {utf8mb4_unicode_520_ci, nil},
	"utf8mb4_vietnamese_ci":    {utf8mb4_vietnamese_ci, nil},
}
