package protocol

type Capability uint32

//read https//dev.mysql.com/doc/internals/en/Capability-flags.html#packet-ProtocolCapabilityFlags
const (
	CLIENT_LONG_PASSWORD                  uint32 = 1 << iota
	CLIENT_FOUND_ROWS                     uint32 = 1 << 1
	CLIENT_LONG_FLAG                      uint32 = 1 << 2
	CLIENT_CONNECT_WITH_DB                uint32 = 1 << 3
	CLIENT_NO_SCHEMA                      uint32 = 1 << 4
	CLIENT_COMPRESS                       uint32 = 1 << 5
	CLIENT_ODBC                           uint32 = 1 << 6
	CLIENT_LOCAL_FILES                    uint32 = 1 << 7
	CLIENT_IGNORE_SPACE                   uint32 = 1 << 8
	CLIENT_PROTOCOL_41                    uint32 = 1 << 9
	CLIENT_INTERACTIVE                    uint32 = 1 << 10
	CLIENT_SSL                            uint32 = 1 << 11
	CLIENT_IGNORE_SIGPIPE                 uint32 = 1 << 12
	CLIENT_TRANSACTIONS                   uint32 = 1 << 13
	CLIENT_RESERVED                       uint32 = 1 << 14
	CLIENT_SECURE_CONNECTION              uint32 = 1 << 15
	CLIENT_MULTI_STATEMENTS               uint32 = 1 << 16
	CLIENT_MULTI_RESULTS                  uint32 = 1 << 17
	CLIENT_PS_MULTI_RESULTS               uint32 = 1 << 18
	CLIENT_PLUGIN_AUTH                    uint32 = 1 << 19
	CLIENT_CONNECT_ATTRS                  uint32 = 1 << 20
	CLIENT_PLUGIN_AUTH_LENENC_CLIENT_DATA uint32 = 1 << 21
	CLIENT_CAN_HANDLE_EXPIRED_PASSWORDS   uint32 = 1 << 22
	CLIENT_SESSION_TRACK                  uint32 = 1 << 23
	CLIENT_DEPRECATE_EOF                  uint32 = 1 << 24

	CLIENT_SSL_VERIFY_SERVER_CERT uint32 = 1 << 30
	CLIENT_REMEMBER_OPTIONS       uint32 = 1 << 31
)

//read https//dev.mysql.com/doc/internals/en/status-flags.html#packet-ProtocolStatusFlags
const (
	SERVER_STATUS_IN_TRANS uint16 = 1 << iota
	SERVER_STATUS_AUTOCOMMIT
	SERVER_MORE_RESULTS_EXISTS
	SERVER_STATUS_NO_GOOD_INDEX_USED
	SERVER_STATUS_NO_INDEX_USED
	SERVER_STATUS_CURSOR_EXISTS
	SERVER_STATUS_LAST_ROW_SENT
	SERVER_STATUS_DB_DROPPED
	SERVER_STATUS_NO_BACKSLASH_ESCAPES
	SERVER_STATUS_METADATA_CHANGED
	SERVER_QUERY_WAS_SLOW
	SERVER_PS_OUT_PARAMS
	SERVER_STATUS_IN_TRANS_READONLY
	SERVER_SESSION_STATE_CHANGE
)

//read https//dev.mysql.com/doc/internals/en/command-phase.html
type Com uint8

const (
	COM_SLEEP               Com = 0x00
	COM_QUIT                Com = 0x01
	COM_INIT_DB             Com = 0x02
	COM_QUERY               Com = 0x03
	COM_FIELD_LIST          Com = 0x04
	COM_CREATE_DB           Com = 0x05
	COM_DROP_DB             Com = 0x06
	COM_REFRESH             Com = 0x07
	COM_SHUTDOWN            Com = 0x08
	COM_STATISTICS          Com = 0x09
	COM_PROCESS_INFO        Com = 0x0A
	COM_CONNECT             Com = 0x0B
	COM_PROCESS_KILL        Com = 0x0C
	COM_DEBUG               Com = 0x0D
	COM_PING                Com = 0x0E
	COM_TIME                Com = 0x0F
	COM_DELAYED_INSERT      Com = 0x10
	COM_CHANGE_USER         Com = 0x11
	COM_BINLOG_DUMP         Com = 0x12
	COM_TABLE_DUMP          Com = 0x13
	COM_CONNECT_OUT         Com = 0x14
	COM_REGISTER_SLAVE      Com = 0x15
	COM_STMT_PREPARE        Com = 0x16
	COM_STMT_EXECUTE        Com = 0x17
	COM_STMT_SEND_LONG_DATA Com = 0x18
	COM_STMT_CLOSE          Com = 0x19
	COM_STMT_RESET          Com = 0x1A
	COM_SET_OPTION          Com = 0x1B
	COM_STMT_FETCH          Com = 0x1C
	COM_DAEMON              Com = 0x1D
	COM_BINLOG_DUMP_GTID    Com = 0x1E
	COM_RESET_CONNECTION    Com = 0x1F
)

type Collation uint8

const (
	Big5_chinese_ci          uint8 = 0x01
	Latin2_czech_cs          uint8 = 0x02
	Dec8_swedish_ci          uint8 = 0x03
	Cp850_general_ci         uint8 = 0x04
	Latin1_german1_ci        uint8 = 0x05
	Hp8_english_ci           uint8 = 0x06
	Koi8r_general_ci         uint8 = 0x07
	Latin1_swedish_ci        uint8 = 0x08
	Latin2_general_ci        uint8 = 0x09
	Swe7_swedish_ci          uint8 = 0x0A
	Ascii_general_ci         uint8 = 0x0B
	Ujis_japanese_ci         uint8 = 0x0C
	Sjis_japanese_ci         uint8 = 0x0D
	Cp1251_bulgarian_ci      uint8 = 0x0E
	Latin1_danish_ci         uint8 = 0x0F
	Hebrew_general_ci        uint8 = 0x10
	Tis620_thai_ci           uint8 = 0x11
	Euckr_korean_ci          uint8 = 0x12
	Latin7_estonian_cs       uint8 = 0x13
	Latin2_hungarian_ci      uint8 = 0x14
	Koi8u_general_ci         uint8 = 0x15
	Cp1251_ukrainian_ci      uint8 = 0x16
	Gb2312_chinese_ci        uint8 = 0x17
	Greek_general_ci         uint8 = 0x18
	Cp1250_general_ci        uint8 = 0x19
	Latin2_croatian_ci       uint8 = 0x1A
	Gbk_chinese_ci           uint8 = 0x1B
	Cp1257_lithuanian_ci     uint8 = 0x1C
	Latin5_turkish_ci        uint8 = 0x1D
	Latin1_german2_ci        uint8 = 0x1E
	Armscii8_general_ci      uint8 = 0x1F
	Utf8_general_ci          uint8 = 0x20
	Cp1250_czech_cs          uint8 = 0x21
	Ucs2_general_ci          uint8 = 0x22
	Cp866_general_ci         uint8 = 0x23
	Keybcs2_general_ci       uint8 = 0x24
	Macce_general_ci         uint8 = 0x25
	Macroman_general_ci      uint8 = 0x26
	Cp852_general_ci         uint8 = 0x27
	Latin7_general_ci        uint8 = 0x28
	Latin7_general_cs        uint8 = 0x29
	Macce_bin                uint8 = 0x3A
	Cp1250_croatian_ci       uint8 = 0x3B
	Utf8mb4_general_ci       uint8 = 0x3C
	Utf8mb4_bin              uint8 = 0x3D
	Latin1_bin               uint8 = 0x3E
	Latin1_general_ci        uint8 = 0x3F
	Latin1_general_cs        uint8 = 0x40
	Cp1251_bin               uint8 = 0x41
	Cp1251_general_ci        uint8 = 0x42
	Cp1251_general_cs        uint8 = 0x43
	Macroman_bin             uint8 = 0x44
	Utf16_general_ci         uint8 = 0x45
	Utf16_bin                uint8 = 0x46
	Utf16le_general_ci       uint8 = 0x47
	Cp1256_general_ci        uint8 = 0x48
	Cp1257_bin               uint8 = 0x49
	Cp1257_general_ci        uint8 = 0x4A
	Utf32_general_ci         uint8 = 0x4B
	Utf32_bin                uint8 = 0x4C
	Utf16le_bin              uint8 = 0x4D
	Binary                   uint8 = 0x4E
	Armscii8_bin             uint8 = 0x4F
	Ascii_bin                uint8 = 0x50
	Cp1250_bin               uint8 = 0x51
	Cp1256_bin               uint8 = 0x52
	Cp866_bin                uint8 = 0x53
	Dec8_bin                 uint8 = 0x54
	Greek_bin                uint8 = 0x55
	Hebrew_bin               uint8 = 0x56
	Hp8_bin                  uint8 = 0x57
	keybcs2_bin              uint8 = 0x58
	Koi8r_bin                uint8 = 0x59
	Koi8u_bin                uint8 = 0x5A
	Latin2_bin               uint8 = 0x5B
	Latin5_bin               uint8 = 0x5C
	Latin7_bin               uint8 = 0x5D
	Cp850_bin                uint8 = 0x5E
	Cp852_bin                uint8 = 0x5F
	Swe7_bin                 uint8 = 0x60
	Utf8_bin                 uint8 = 0x61
	Big5_bin                 uint8 = 0x62
	Euckr_bin                uint8 = 0x63
	Gb2312_bin               uint8 = 0x64
	Gbk_bin                  uint8 = 0x65
	Sjis_bin                 uint8 = 0x66
	Tis620_bin               uint8 = 0x67
	Ucs2_bin                 uint8 = 0x68
	Ujis_bin                 uint8 = 0x69
	Geostd8_general_ci       uint8 = 0x6A
	Geostd8_bin              uint8 = 0x6B
	Latin1_spanish_ci        uint8 = 0x6C
	Cp932_japanese_ci        uint8 = 0x6D
	Cp932_bin                uint8 = 0x6E
	Eucjpms_japanese_ci      uint8 = 0x6F
	Eucjpms_bin              uint8 = 0x70
	Cp1250_polish_ci         uint8 = 0x71
	Utf16_unicode_ci         uint8 = 0x72
	Utf16_icelandic_ci       uint8 = 0x73
	Utf16_latvian_ci         uint8 = 0x74
	Utf16_romanian_ci        uint8 = 0x75
	Utf16_slovenian_ci       uint8 = 0x76
	Utf16_polish_ci          uint8 = 0x77
	Utf16_estonian_ci        uint8 = 0x78
	Utf16_spanish_ci         uint8 = 0x79
	Utf16_swedish_ci         uint8 = 0x7A
	Utf16_turkish_ci         uint8 = 0x7B
	Utf16_czech_ci           uint8 = 0x7C
	Utf16_danish_ci          uint8 = 0x7D
	Utf16_lithuanian_ci      uint8 = 0x7E
	Utf16_slovak_ci          uint8 = 0x7F
	Utf16_spanish2_ci        uint8 = 0x80
	Utf16_roman_ci           uint8 = 0x81
	Utf16_persian_ci         uint8 = 0x82
	Utf16_esperanto_ci       uint8 = 0x83
	Utf16_hungarian_ci       uint8 = 0x84
	Utf16_sinhala_ci         uint8 = 0x85
	Utf16_german2_ci         uint8 = 0x86
	Utf16_croatian_ci        uint8 = 0x87
	Utf16_unicode_520_ci     uint8 = 0x88
	Utf16_vietnamese_ci      uint8 = 0x89
	Ucs2_unicode_ci          uint8 = 0x8A
	Ucs2_icelandic_ci        uint8 = 0x8B
	Ucs2_latvian_ci          uint8 = 0x8C
	Ucs2_romanian_ci         uint8 = 0x8D
	Ucs2_slovenian_ci        uint8 = 0x8E
	Ucs2_polish_ci           uint8 = 0x8F
	Ucs2_estonian_ci         uint8 = 0x90
	Ucs2_spanish_ci          uint8 = 0x91
	Ucs2_swedish_ci          uint8 = 0x92
	Ucs2_turkish_ci          uint8 = 0x93
	Ucs2_czech_ci            uint8 = 0x94
	Ucs2_danish_ci           uint8 = 0x95
	Ucs2_lithuanian_ci       uint8 = 0x96
	Ucs2_slovak_ci           uint8 = 0x97
	Ucs2_spanish2_ci         uint8 = 0x98
	Ucs2_roman_ci            uint8 = 0x99
	Ucs2_persian_ci          uint8 = 0x9A
	Ucs2_esperanto_ci        uint8 = 0x9B
	Ucs2_hungarian_ci        uint8 = 0x9C
	Ucs2_sinhala_ci          uint8 = 0x9D
	Ucs2_german2_ci          uint8 = 0x9E
	Ucs2_croatian_ci         uint8 = 0x9F
	Ucs2_unicode_520_ci      uint8 = 0xA0
	Ucs2_vietnamese_ci       uint8 = 0xA1
	Ucs2_general_mysql500_ci uint8 = 0xA2
	Utf32_unicode_ci         uint8 = 0xA3
	Utf32_icelandic_ci       uint8 = 0xA4
	Utf32_latvian_ci         uint8 = 0xA5
	Utf32_romanian_ci        uint8 = 0xA6
	Utf32_slovenian_ci       uint8 = 0xA7
	Utf32_polish_ci          uint8 = 0xA8
	Utf32_estonian_ci        uint8 = 0xA9
	Utf32_spanish_ci         uint8 = 0xAA
	Utf32_swedish_ci         uint8 = 0xAB
	Utf32_turkish_ci         uint8 = 0xAC
	Utf32_czech_ci           uint8 = 0xAD
	Utf32_danish_ci          uint8 = 0xAE
	Utf32_lithuanian_ci      uint8 = 0xAF
	Utf32_slovak_ci          uint8 = 0xB0
	Utf32_spanish2_ci        uint8 = 0xB1
	Utf32_roman_ci           uint8 = 0xB2
	Utf32_persian_ci         uint8 = 0xB3
	Utf32_esperanto_ci       uint8 = 0xB4
	Utf32_hungarian_ci       uint8 = 0xB5
	Utf32_sinhala_ci         uint8 = 0xB6
	Utf32_german2_ci         uint8 = 0xB7
	Utf32_croatian_ci        uint8 = 0xB8
	Utf32_unicode_520_ci     uint8 = 0xB9
	Utf32_vietnamese_ci      uint8 = 0xBA
	Utf8_unicode_ci          uint8 = 0xBB
	Utf8_icelandic_ci        uint8 = 0xBC
	Utf8_latvian_ci          uint8 = 0xBD
	Utf8_romanian_ci         uint8 = 0xBE
	Utf8_slovenian_ci        uint8 = 0xBF
	Utf8_polish_ci           uint8 = 0xC0
	Utf8_estonian_ci         uint8 = 0xC1
	Utf8_spanish_ci          uint8 = 0xC2
	Utf8_swedish_ci          uint8 = 0xC3
	Utf8_turkish_ci          uint8 = 0xC4
	Utf8_czech_ci            uint8 = 0xC5
	Utf8_danish_ci           uint8 = 0xC6
	Utf8_lithuanian_ci       uint8 = 0xC7
	Utf8_slovak_ci           uint8 = 0xC8
	Utf8_spanish2_ci         uint8 = 0xC9
	Utf8_roman_ci            uint8 = 0xCA
	Utf8_persian_ci          uint8 = 0xCB
	Utf8_esperanto_ci        uint8 = 0xCC
	Utf8_hungarian_ci        uint8 = 0xCD
	Utf8_sinhala_ci          uint8 = 0xCE
	Utf8_german2_ci          uint8 = 0xCF
	Utf8_croatian_ci         uint8 = 0xD0
	Utf8_unicode_520_ci      uint8 = 0xD1
	Utf8_vietnamese_ci       uint8 = 0xD2
	Utf8_general_mysql500_ci uint8 = 0xD3
	Utf8mb4_unicode_ci       uint8 = 0xD4
	Utf8mb4_icelandic_ci     uint8 = 0xD5
	Utf8mb4_latvian_ci       uint8 = 0xD6
	Utf8mb4_romanian_ci      uint8 = 0xD7
	Utf8mb4_slovenian_ci     uint8 = 0xD8
	Utf8mb4_polish_ci        uint8 = 0xD9
	Utf8mb4_estonian_ci      uint8 = 0xDA
	Utf8mb4_spanish_ci       uint8 = 0xDB
	Utf8mb4_swedish_ci       uint8 = 0xDC
	Utf8mb4_turkish_ci       uint8 = 0xDD
	Utf8mb4_czech_ci         uint8 = 0xDE
	Utf8mb4_danish_ci        uint8 = 0xDF
	Utf8mb4_lithuanian_ci    uint8 = 0xE0
	Utf8mb4_slovak_ci        uint8 = 0xE1
	Utf8mb4_spanish2_ci      uint8 = 0xE2
	Utf8mb4_roman_ci         uint8 = 0xE3
	Utf8mb4_persian_ci       uint8 = 0xE4
	Utf8mb4_esperanto_ci     uint8 = 0xE5
	Utf8mb4_hungarian_ci     uint8 = 0xE6
	Utf8mb4_sinhala_ci       uint8 = 0xE7
	Utf8mb4_german2_ci       uint8 = 0xE8
	Utf8mb4_croatian_ci      uint8 = 0xE9
	Utf8mb4_unicode_520_ci   uint8 = 0xEA
	Utf8mb4_vietnamese_ci    uint8 = 0xE1
)

const (
	mysql_native_password uint8 = 0x01
	ssl                   uint8 = 0x02
)
