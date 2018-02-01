//line parser.y:2
package parser

import __yyfmt__ "fmt"

//line parser.y:4
import "ast"

//line parser.y:11
type yySymType struct {
	yys   int
	ident string
	str   string
	val   string
	node  ast.Node
	expr  ast.ExprNode
	list  ast.List
}

const DOUBLE_AT_IDENT = 57346
const IDENT = 57347
const SINGLE_AT_IDENT = 57348
const PARAM_MARKER = 57349
const DOUBLE_QUOTA_STRING = 57350
const SINGLE_QUOTA_STRING = 57351
const STRING = 57352
const BIT_NUMBER = 57353
const HEX_NUMBER = 57354
const NUMBER = 57355
const AS = 57356
const ASC = 57357
const ALL = 57358
const ANY = 57359
const AVG = 57360
const BY = 57361
const COMMENT = 57362
const COUNT = 57363
const CHARACTER = 57364
const DISTINCT = 57365
const DISTINCTROW = 57366
const DESC = 57367
const DAY = 57368
const DAY_HOUR = 57369
const DAY_MICROSECOND = 57370
const DAY_MINUTE = 57371
const DAY_SECOND = 57372
const FROM = 57373
const FOR = 57374
const FALSE = 57375
const GROUP = 57376
const HAVING = 57377
const HOUR = 57378
const HOUR_MICROSECOND = 57379
const HOUR_MINUTE = 57380
const HOUR_SECOND = 57381
const INTO = 57382
const IF = 57383
const LIMIT = 57384
const LOCK = 57385
const MAX = 57386
const MIN = 57387
const MINUTE = 57388
const MINUTE_SECOND = 57389
const MODE = 57390
const MONTH = 57391
const MICROSECOND = 57392
const MINUTE_MICROSECOND = 57393
const NULL = 57394
const ORDER = 57395
const OFFSET = 57396
const QUARTER = 57397
const SELECT = 57398
const SOME = 57399
const SUM = 57400
const SHARE = 57401
const SOUNDS = 57402
const SECOND = 57403
const SECOND_MICROSECOND = 57404
const SET = 57405
const TO = 57406
const TRUE = 57407
const UPDATE = 57408
const WHERE = 57409
const WEEK = 57410
const YEAR = 57411
const YEAR_MONTH = 57412
const armscii8_general_ci = 57413
const ascii_general_ci = 57414
const big5_chinese_ci = 57415
const binary = 57416
const cp1250_general_ci = 57417
const cp1251_general_ci = 57418
const cp1256_general_ci = 57419
const cp1257_general_ci = 57420
const cp850_general_ci = 57421
const cp852_general_ci = 57422
const cp866_general_ci = 57423
const cp932_japanese_ci = 57424
const dec8_swedish_ci = 57425
const eucjpms_japanese_ci = 57426
const euckr_korean_ci = 57427
const gb18030_chinese_ci = 57428
const gb2312_chinese_ci = 57429
const gbk_chinese_ci = 57430
const geostd8_general_ci = 57431
const greek_general_ci = 57432
const hebrew_general_ci = 57433
const hp8_english_ci = 57434
const keybcs2_general_ci = 57435
const koi8r_general_ci = 57436
const koi8u_general_ci = 57437
const latin1_swedish_ci = 57438
const latin2_general_ci = 57439
const latin5_turkish_ci = 57440
const latin7_general_ci = 57441
const macce_general_ci = 57442
const macroman_general_ci = 57443
const sjis_japanese_ci = 57444
const swe7_swedish_ci = 57445
const tis620_thai_ci = 57446
const ucs2_general_ci = 57447
const ujis_japanese_ci = 57448
const utf16_general_ci = 57449
const utf16le_general_ci = 57450
const utf32_general_ci = 57451
const utf8_general_ci = 57452
const utf8mb4_general_ci = 57453
const LOW = 57454
const INTERVAL = 57455
const UNION = 57456
const EXCEPT = 57457
const INTERSECT = 57458
const JOIN = 57459
const STRAIGHT_JOIN = 57460
const LEFT = 57461
const RIGHT = 57462
const INNER = 57463
const OUTER = 57464
const CROSS = 57465
const NATURAL = 57466
const USE = 57467
const ON = 57468
const USING = 57469
const OR = 57470
const OO = 57471
const AND = 57472
const AA = 57473
const NOT = 57474
const BETWEEN = 57475
const CASE = 57476
const WHEN = 57477
const THEN = 57478
const ELSE = 57479
const IS = 57480
const LIKE = 57481
const REGEXP = 57482
const IN = 57483
const GE = 57484
const LE = 57485
const NE = 57486
const LG = 57487
const IE = 57488
const LNE = 57489
const SL = 57490
const SR = 57491
const LEG = 57492
const XOR = 57493
const EXISTS = 57494
const DIV = 57495
const MOD = 57496
const OP = 57497
const COLLATE = 57498
const UMINUS = 57499
const END = 57500

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"DOUBLE_AT_IDENT",
	"IDENT",
	"SINGLE_AT_IDENT",
	"PARAM_MARKER",
	"DOUBLE_QUOTA_STRING",
	"SINGLE_QUOTA_STRING",
	"STRING",
	"BIT_NUMBER",
	"HEX_NUMBER",
	"NUMBER",
	"AS",
	"ASC",
	"ALL",
	"ANY",
	"AVG",
	"BY",
	"COMMENT",
	"COUNT",
	"CHARACTER",
	"DISTINCT",
	"DISTINCTROW",
	"DESC",
	"DAY",
	"DAY_HOUR",
	"DAY_MICROSECOND",
	"DAY_MINUTE",
	"DAY_SECOND",
	"FROM",
	"FOR",
	"FALSE",
	"GROUP",
	"HAVING",
	"HOUR",
	"HOUR_MICROSECOND",
	"HOUR_MINUTE",
	"HOUR_SECOND",
	"INTO",
	"IF",
	"LIMIT",
	"LOCK",
	"MAX",
	"MIN",
	"MINUTE",
	"MINUTE_SECOND",
	"MODE",
	"MONTH",
	"MICROSECOND",
	"MINUTE_MICROSECOND",
	"NULL",
	"ORDER",
	"OFFSET",
	"QUARTER",
	"SELECT",
	"SOME",
	"SUM",
	"SHARE",
	"SOUNDS",
	"SECOND",
	"SECOND_MICROSECOND",
	"SET",
	"TO",
	"TRUE",
	"UPDATE",
	"WHERE",
	"WEEK",
	"YEAR",
	"YEAR_MONTH",
	"armscii8_general_ci",
	"ascii_general_ci",
	"big5_chinese_ci",
	"binary",
	"cp1250_general_ci",
	"cp1251_general_ci",
	"cp1256_general_ci",
	"cp1257_general_ci",
	"cp850_general_ci",
	"cp852_general_ci",
	"cp866_general_ci",
	"cp932_japanese_ci",
	"dec8_swedish_ci",
	"eucjpms_japanese_ci",
	"euckr_korean_ci",
	"gb18030_chinese_ci",
	"gb2312_chinese_ci",
	"gbk_chinese_ci",
	"geostd8_general_ci",
	"greek_general_ci",
	"hebrew_general_ci",
	"hp8_english_ci",
	"keybcs2_general_ci",
	"koi8r_general_ci",
	"koi8u_general_ci",
	"latin1_swedish_ci",
	"latin2_general_ci",
	"latin5_turkish_ci",
	"latin7_general_ci",
	"macce_general_ci",
	"macroman_general_ci",
	"sjis_japanese_ci",
	"swe7_swedish_ci",
	"tis620_thai_ci",
	"ucs2_general_ci",
	"ujis_japanese_ci",
	"utf16_general_ci",
	"utf16le_general_ci",
	"utf32_general_ci",
	"utf8_general_ci",
	"utf8mb4_general_ci",
	"LOW",
	"INTERVAL",
	"UNION",
	"EXCEPT",
	"INTERSECT",
	"','",
	"JOIN",
	"STRAIGHT_JOIN",
	"LEFT",
	"RIGHT",
	"INNER",
	"OUTER",
	"CROSS",
	"NATURAL",
	"USE",
	"ON",
	"USING",
	"OR",
	"OO",
	"AND",
	"AA",
	"NOT",
	"BETWEEN",
	"CASE",
	"WHEN",
	"THEN",
	"ELSE",
	"IS",
	"'='",
	"'<'",
	"'>'",
	"LIKE",
	"REGEXP",
	"IN",
	"GE",
	"LE",
	"NE",
	"LG",
	"IE",
	"LNE",
	"SL",
	"SR",
	"LEG",
	"XOR",
	"EXISTS",
	"'|'",
	"'&'",
	"'+'",
	"'-'",
	"'*'",
	"'/'",
	"'%'",
	"DIV",
	"MOD",
	"'^'",
	"'~'",
	"OP",
	"COLLATE",
	"UMINUS",
	"'('",
	"')'",
	"'.'",
	"END",
	"'!'",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 4,
	1, 3,
	-2, 11,
	-1, 5,
	1, 6,
	172, 6,
	-2, 10,
	-1, 9,
	1, 7,
	172, 7,
	-2, 243,
	-1, 29,
	134, 130,
	143, 130,
	144, 130,
	145, 130,
	-2, 36,
	-1, 128,
	32, 11,
	42, 11,
	43, 11,
	53, 11,
	54, 11,
	114, 11,
	115, 11,
	116, 11,
	-2, 64,
	-1, 241,
	173, 213,
	-2, 199,
	-1, 242,
	173, 214,
	-2, 200,
}

const yyPrivate = 57344

const yyLast = 1026

var yyAct = [...]int{

	24, 236, 348, 350, 41, 4, 234, 327, 272, 333,
	28, 50, 16, 278, 239, 233, 294, 210, 88, 230,
	151, 87, 214, 30, 108, 159, 56, 95, 318, 93,
	94, 92, 290, 373, 56, 89, 91, 166, 256, 242,
	29, 129, 132, 339, 131, 128, 254, 130, 56, 256,
	256, 166, 133, 242, 265, 77, 135, 8, 76, 361,
	124, 125, 126, 127, 82, 83, 85, 86, 246, 150,
	152, 7, 140, 139, 138, 155, 160, 137, 155, 155,
	136, 123, 122, 169, 170, 171, 172, 173, 372, 120,
	84, 164, 84, 342, 163, 146, 190, 121, 338, 8,
	161, 91, 215, 185, 322, 255, 353, 326, 179, 180,
	115, 116, 119, 117, 118, 120, 111, 112, 85, 86,
	122, 109, 110, 113, 114, 115, 116, 119, 117, 118,
	120, 334, 215, 211, 259, 283, 284, 364, 219, 359,
	14, 354, 84, 209, 337, 212, 336, 207, 79, 293,
	191, 192, 193, 194, 195, 197, 199, 200, 201, 202,
	203, 204, 82, 83, 85, 86, 275, 237, 448, 226,
	111, 112, 7, 225, 241, 109, 110, 113, 114, 115,
	116, 119, 117, 118, 120, 187, 244, 81, 84, 72,
	12, 245, 446, 77, 188, 189, 186, 251, 20, 252,
	238, 11, 13, 12, 444, 325, 442, 82, 83, 85,
	86, 440, 248, 438, 238, 320, 260, 261, 211, 263,
	436, 434, 266, 267, 268, 81, 152, 160, 247, 432,
	249, 250, 274, 84, 430, 262, 258, 428, 426, 82,
	83, 85, 86, 287, 424, 422, 270, 211, 91, 91,
	241, 128, 271, 420, 289, 418, 286, 317, 285, 416,
	319, 414, 412, 291, 292, 84, 316, 410, 243, 314,
	82, 83, 85, 86, 408, 82, 83, 85, 86, 406,
	237, 237, 237, 404, 332, 402, 400, 241, 241, 241,
	398, 330, 331, 335, 396, 394, 84, 392, 390, 91,
	341, 84, 164, 388, 343, 340, 286, 82, 83, 85,
	86, 90, 386, 324, 82, 83, 85, 86, 323, 384,
	21, 345, 82, 83, 85, 86, 78, 217, 382, 346,
	349, 380, 378, 84, 376, 280, 277, 283, 284, 281,
	84, 282, 279, 371, 237, 355, 358, 315, 84, 357,
	321, 241, 360, 206, 362, 356, 231, 264, 67, 237,
	349, 145, 366, 368, 224, 208, 241, 369, 5, 237,
	363, 370, 62, 241, 63, 375, 241, 113, 114, 115,
	116, 119, 117, 118, 120, 241, 52, 56, 51, 54,
	49, 48, 47, 44, 45, 46, 66, 269, 218, 68,
	57, 174, 168, 58, 69, 22, 23, 167, 68, 178,
	79, 69, 328, 273, 162, 52, 56, 51, 54, 49,
	48, 47, 44, 45, 46, 147, 59, 60, 176, 57,
	56, 329, 58, 64, 22, 23, 144, 142, 75, 154,
	61, 177, 154, 154, 65, 52, 56, 51, 54, 49,
	48, 47, 44, 45, 46, 59, 60, 184, 183, 57,
	15, 3, 58, 56, 148, 447, 93, 94, 92, 61,
	70, 445, 89, 52, 56, 51, 54, 49, 48, 47,
	44, 45, 46, 73, 74, 59, 60, 57, 443, 56,
	58, 441, 166, 19, 242, 111, 112, 439, 182, 61,
	109, 110, 113, 114, 115, 116, 119, 117, 118, 120,
	6, 71, 80, 59, 60, 26, 228, 55, 22, 23,
	437, 435, 433, 431, 429, 8, 229, 61, 109, 110,
	113, 114, 115, 116, 119, 117, 118, 120, 42, 427,
	425, 36, 37, 220, 26, 56, 55, 423, 166, 38,
	165, 56, 421, 40, 93, 94, 92, 39, 110, 113,
	114, 115, 116, 119, 117, 118, 120, 42, 419, 64,
	36, 37, 25, 417, 415, 413, 55, 411, 38, 68,
	65, 153, 40, 409, 156, 157, 39, 82, 83, 85,
	86, 69, 407, 405, 403, 401, 399, 42, 397, 395,
	36, 37, 26, 393, 55, 391, 389, 216, 38, 221,
	222, 223, 40, 84, 387, 385, 39, 280, 277, 283,
	284, 281, 383, 282, 279, 42, 351, 352, 36, 37,
	82, 83, 85, 86, 56, 381, 38, 93, 94, 92,
	40, 379, 377, 89, 39, 52, 56, 51, 54, 49,
	48, 47, 44, 45, 46, 374, 84, 365, 344, 57,
	253, 1, 58, 52, 56, 51, 54, 49, 48, 47,
	44, 45, 46, 149, 213, 288, 158, 57, 367, 17,
	58, 99, 100, 104, 102, 59, 60, 347, 101, 103,
	232, 105, 33, 106, 35, 31, 97, 235, 34, 61,
	257, 32, 43, 59, 60, 52, 56, 51, 54, 49,
	48, 47, 44, 45, 46, 134, 27, 61, 53, 57,
	175, 181, 58, 52, 56, 51, 54, 49, 48, 47,
	44, 45, 46, 107, 96, 276, 240, 57, 98, 205,
	58, 227, 9, 2, 10, 59, 60, 141, 143, 18,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 61,
	0, 0, 0, 59, 60, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 26, 0, 55, 61, 0, 0,
	0, 52, 56, 51, 54, 49, 48, 47, 44, 45,
	46, 0, 26, 0, 55, 57, 0, 42, 58, 0,
	36, 37, 25, 0, 0, 0, 0, 0, 38, 0,
	0, 0, 40, 0, 198, 42, 39, 0, 36, 37,
	0, 59, 60, 0, 0, 0, 38, 0, 0, 0,
	40, 0, 196, 0, 39, 61, 55, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 55, 0, 0, 42, 0, 0,
	36, 37, 0, 0, 0, 0, 0, 0, 38, 0,
	0, 0, 40, 0, 0, 42, 39, 0, 36, 37,
	0, 0, 0, 0, 0, 0, 38, 0, 0, 0,
	40, 0, 0, 0, 39, 0, 295, 296, 297, 298,
	299, 0, 0, 0, 0, 0, 300, 301, 302, 303,
	0, 0, 55, 0, 0, 0, 304, 305, 0, 306,
	307, 308, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 309, 310, 42, 0, 0, 36, 37, 311, 312,
	313, 0, 0, 0, 38, 0, 0, 0, 40, 0,
	0, 0, 39, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 82,
	83, 85, 86, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 84,
}
var yyPact = [...]int{

	1, -1000, -1000, -1000, -1000, -1000, 87, 1, 411, -1000,
	537, 495, 495, 495, 419, -114, -117, 370, 641, 70,
	-1000, -1000, -1000, -1000, 458, -1000, 659, 542, -1000, -36,
	-49, -1000, -1000, -1000, -1000, -1000, 777, 777, 777, 777,
	469, -1000, -100, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-129, -1000, -121, -1000, -1000, 659, -1000, -91, -94, -97,
	-98, -99, 357, 401, 295, -50, 350, 366, 659, 659,
	1, -1000, -1000, 1, 1, 659, -1000, -1000, 383, 540,
	108, 641, 659, 659, 659, 659, 659, -1000, -1000, 546,
	-1000, -1000, -1000, -1000, -1000, -63, 376, 777, 441, -30,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 51, -47, 777,
	777, 777, 777, 719, 701, 777, 777, 777, 777, 777,
	777, -1000, 282, 777, -87, -87, -87, -87, -117, 193,
	-1000, 425, 659, 425, -34, 110, 495, 382, 495, 495,
	495, -1000, -1000, -1000, -1000, -1000, 305, -1000, -1000, 56,
	110, -1000, 110, 74, -1000, -1000, -1000, 74, 52, -1000,
	501, 289, 29, -1000, -1000, -1000, -1000, 383, -1000, -13,
	-13, -1000, -63, -63, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -100, -1000, -1000, -1000, -1000, -103, 777, 777, 777,
	777, 400, 218, 371, 371, -51, 659, -51, 659, -77,
	-77, -77, -77, -77, -1000, -1000, -1000, -87, -1000, -127,
	-67, 110, -1000, -4, -1000, 659, 659, 659, 659, 185,
	-118, 659, 659, 659, 349, 659, 659, -1000, -1000, -1000,
	379, 659, 49, 217, -1000, -1000, 629, 629, 43, -1000,
	-141, -1000, -1000, 289, -1000, -1000, 469, 18, -87, 343,
	343, 870, 870, 275, 425, -1000, 659, -146, -1000, 659,
	78, 178, -68, 146, -1000, -1000, 141, 33, -65, -1000,
	-1000, -1000, 377, 412, 110, 29, 29, 29, 8, 15,
	-1000, 28, 26, -1000, -1000, -1000, -1000, 21, -74, 217,
	540, 379, -79, 777, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 110, -1000, 110,
	659, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 659, 659,
	217, 499, -21, 23, -1000, 8, -1000, -1000, -1000, 29,
	-1000, 377, -1000, -1000, 273, 110, 110, 22, -1000, 110,
	-1000, 659, -112, 659, 29, 19, 217, -1000, -1000, 659,
	110, 484, 110, 499, 29, 269, -1000, -84, -1000, -1000,
	-1000, -1000, -1000, 484, 259, -1000, -1000, 256, -1000, 254,
	-1000, 250, -1000, 240, -1000, 232, -1000, 222, -1000, 216,
	-1000, 214, -1000, 211, -1000, 209, -1000, 204, -1000, 199,
	-1000, 197, -1000, 194, -1000, 189, -1000, 183, -1000, 175,
	-1000, 169, -1000, 167, -1000, 164, -1000, 159, -1000, 156,
	-1000, 147, -1000, 145, -1000, 138, -1000, 136, -1000, 132,
	-1000, 126, -1000, 117, -1000, 115, -1000, 107, -1000, 104,
	-1000, 98, -1000, 95, -1000, 82, -1000, 57, -1000,
}
var yyPgo = [...]int{

	0, 749, 372, 100, 8, 7, 326, 748, 747, 744,
	510, 374, 460, 743, 4, 368, 742, 19, 18, 311,
	470, 741, 739, 738, 736, 11, 735, 734, 13, 21,
	189, 733, 9, 721, 14, 16, 720, 718, 40, 716,
	715, 702, 701, 700, 0, 698, 2, 3, 697, 695,
	694, 358, 10, 1, 23, 25, 6, 15, 320, 692,
	22, 20, 17, 690, 687, 679, 678, 676, 675, 493,
	674, 673, 396, 661, 660, 658, 657, 655, 642, 641,
	635, 622, 615, 614, 606, 605, 603, 599, 598, 596,
	595, 594, 593, 592, 583, 577, 575, 574, 573, 568,
	552, 547, 540, 539, 524, 523, 522, 521, 520, 497,
	491, 488, 471, 465,
}
var yyR1 = [...]int{

	0, 73, 13, 13, 14, 14, 12, 12, 12, 12,
	10, 10, 15, 15, 15, 15, 15, 44, 44, 44,
	44, 44, 44, 44, 44, 39, 39, 39, 39, 39,
	52, 52, 52, 52, 52, 52, 52, 38, 38, 38,
	38, 38, 38, 38, 38, 38, 38, 38, 38, 38,
	38, 38, 54, 54, 54, 54, 54, 54, 54, 54,
	54, 54, 54, 54, 54, 54, 54, 49, 49, 49,
	49, 49, 49, 25, 59, 59, 59, 42, 42, 42,
	45, 45, 37, 37, 37, 37, 37, 37, 37, 37,
	62, 62, 50, 41, 40, 40, 70, 70, 60, 43,
	43, 65, 65, 69, 69, 58, 58, 29, 29, 18,
	18, 19, 19, 19, 19, 36, 36, 33, 33, 33,
	23, 23, 23, 23, 23, 23, 23, 27, 27, 31,
	31, 35, 35, 35, 35, 35, 35, 35, 35, 35,
	35, 35, 35, 35, 35, 35, 35, 35, 35, 35,
	74, 75, 76, 77, 78, 79, 80, 81, 82, 83,
	84, 85, 86, 87, 88, 89, 90, 91, 92, 93,
	94, 95, 96, 97, 98, 99, 100, 101, 102, 103,
	104, 105, 106, 107, 108, 109, 110, 111, 112, 113,
	22, 1, 20, 20, 20, 30, 30, 6, 6, 34,
	34, 34, 3, 3, 63, 63, 57, 57, 56, 56,
	56, 53, 53, 24, 24, 68, 68, 48, 48, 48,
	48, 48, 48, 26, 26, 26, 28, 28, 32, 32,
	47, 47, 66, 66, 17, 17, 4, 4, 64, 64,
	46, 5, 5, 9, 9, 16, 67, 67, 55, 21,
	21, 21, 7, 7, 2, 2, 8, 8, 11, 11,
	11, 11, 72, 72, 71, 61, 51,
}
var yyR2 = [...]int{

	0, 1, 1, 1, 3, 3, 1, 2, 4, 4,
	1, 1, 7, 8, 4, 4, 4, 3, 3, 3,
	3, 3, 2, 3, 1, 3, 3, 3, 4, 1,
	4, 6, 6, 4, 4, 4, 1, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 5,
	5, 1, 1, 1, 1, 1, 3, 1, 3, 2,
	2, 2, 2, 3, 1, 2, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 3, 1, 3, 5,
	4, 1, 5, 5, 5, 4, 4, 5, 5, 5,
	1, 3, 1, 5, 1, 0, 1, 2, 4, 2,
	0, 1, 0, 1, 3, 2, 1, 1, 0, 2,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 2, 1,
	0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	81, 1, 1, 1, 0, 1, 1, 2, 0, 1,
	1, 1, 2, 0, 1, 3, 1, 1, 2, 2,
	3, 1, 3, 1, 1, 1, 3, 3, 4, 3,
	5, 6, 6, 1, 2, 2, 1, 1, 1, 0,
	2, 4, 1, 3, 2, 0, 3, 0, 1, 3,
	1, 2, 0, 1, 0, 3, 1, 3, 2, 1,
	1, 0, 1, 0, 2, 4, 1, 0, 2, 2,
	1, 1, 2, 4, 1, 1, 2,
}
var yyChk = [...]int{

	-1000, -73, -13, -12, -14, -15, -10, 171, 56, -16,
	-9, 114, 116, 115, 53, -12, -14, -65, -1, -69,
	-30, -58, 23, 24, -44, 161, 133, -39, -52, -38,
	-54, -49, -42, -59, -45, -50, 159, 160, 167, 175,
	171, -14, 156, -41, 11, 12, 13, 10, 9, 8,
	-25, 6, 4, -37, 7, 135, 5, 18, 21, 44,
	45, 58, -2, -11, 32, 43, -72, -51, 42, 54,
	-20, 16, -30, -20, -20, 19, 172, 172, -6, 40,
	-69, 117, 129, 130, 155, 131, 132, -29, -18, 14,
	-19, -25, 10, 8, 9, -44, -27, 154, -23, 139,
	140, 146, 142, 147, 141, 149, 151, -31, 60, 157,
	158, 152, 153, 159, 160, 161, 162, 164, 165, 163,
	166, 133, 169, 130, -54, -54, -54, -54, -14, -44,
	-14, 173, 171, 173, -40, -44, 171, 171, 171, 171,
	171, -8, -11, -7, -2, 66, 145, -51, -72, -71,
	-44, -61, -44, -10, -15, -14, -10, -10, -67, -55,
	-44, -3, 31, -34, -25, 10, 8, -6, -58, -44,
	-44, -44, -44, -44, -19, -36, 52, 65, 33, -52,
	-52, -33, 57, 17, 16, 133, 145, 134, 143, 144,
	143, -38, -38, -38, -38, -38, 113, -38, 113, -38,
	-38, -38, -38, -38, -38, -22, 71, -54, 172, -25,
	-62, -44, -25, -70, -60, 136, -20, -30, 16, -44,
	161, -20, -20, -20, 59, 117, 117, -21, 15, 25,
	-17, 67, -63, -57, -56, -48, -53, -14, 171, -34,
	-24, -25, 10, -3, -14, -14, 171, -38, -54, -38,
	-38, -44, -44, -74, 173, 172, 117, -43, -60, 138,
	-44, -44, -62, -44, 172, 172, -44, -44, -44, 48,
	-61, -55, -4, 34, -44, 117, -26, 119, -28, 125,
	118, 122, 124, 120, 121, -29, -18, -14, -68, -57,
	173, -17, -62, 131, -35, 26, 27, 28, 29, 30,
	36, 37, 38, 39, 46, 47, 49, 50, 51, 61,
	62, 68, 69, 70, -35, 72, -25, -44, 174, -44,
	137, 172, 172, 172, 172, 172, 172, -5, 35, 19,
	-57, -57, -56, -32, 123, -28, 118, 118, 172, 117,
	-34, -4, 172, -52, -75, -44, -44, -64, -46, -44,
	-47, 127, 128, 127, 118, -32, -57, -5, 73, 117,
	-44, 171, -44, -57, 118, -76, -46, -66, -53, -47,
	-56, 74, 172, 117, -77, -53, 75, -78, 76, -79,
	77, -80, 78, -81, 79, -82, 80, -83, 81, -84,
	82, -85, 83, -86, 84, -87, 85, -88, 86, -89,
	87, -90, 88, -91, 89, -92, 90, -93, 91, -94,
	92, -95, 93, -96, 94, -97, 95, -98, 96, -99,
	97, -100, 98, -101, 99, -102, 100, -103, 101, -104,
	102, -105, 103, -106, 104, -107, 105, -108, 106, -109,
	107, -110, 108, -111, 109, -112, 110, -113, 111,
}
var yyDef = [...]int{

	0, -2, 1, 2, -2, -2, 244, 0, 102, -2,
	0, 194, 194, 194, 0, 0, 11, 198, 0, 101,
	191, 103, 195, 196, 108, 106, 0, 24, 29, -2,
	51, 52, 53, 54, 55, 57, 0, 0, 0, 0,
	0, 64, 0, 66, 67, 68, 69, 70, 71, 72,
	77, 74, 75, 81, 92, 95, 73, 0, 0, 0,
	0, 0, 257, 253, 0, 0, 260, 261, 0, 0,
	0, 192, 193, 0, 0, 0, 4, 5, 203, 0,
	198, 0, 0, 0, 0, 0, 0, 105, 107, 0,
	110, 111, 112, 113, 114, 22, 0, 0, 0, 127,
	120, 121, 122, 123, 124, 125, 126, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 129, 0, 0, 59, 60, 61, 62, -2, 0,
	65, 0, 0, 0, 0, 94, 194, 0, 194, 194,
	194, 8, 256, 9, 252, 254, 0, 258, 259, 262,
	264, 266, 265, 14, 10, 11, 15, 16, 245, 246,
	251, 235, 0, 197, 199, 200, 201, 203, 104, 17,
	18, 19, 20, 21, 109, 23, 25, 115, 116, 26,
	27, 0, 117, 118, 119, 128, 0, 0, 0, 0,
	0, 37, 38, 39, 40, 41, 0, 42, 0, 43,
	44, 45, 46, 47, 48, 56, 150, 58, 63, 78,
	0, 90, 76, 100, 96, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 248, 249, 250,
	237, 0, 202, 204, 206, 207, 108, 0, 0, 211,
	0, -2, -2, 235, 28, 30, 0, 0, 34, 35,
	33, 0, 0, 0, 0, 80, 0, 0, 97, 0,
	0, 0, 0, 0, 85, 86, 0, 0, 0, 255,
	263, 247, 242, 0, 234, 0, 0, 0, 229, 0,
	223, 0, 0, 226, 227, 208, 209, 11, 0, 215,
	0, 237, 0, 0, 49, 131, 132, 133, 134, 135,
	136, 137, 138, 139, 140, 141, 142, 143, 144, 145,
	146, 147, 148, 149, 50, 151, 79, 91, 93, 99,
	0, 82, 83, 84, 87, 88, 89, 12, 0, 0,
	205, 217, 219, 0, 228, 229, 224, 225, 210, 0,
	212, 242, 31, 32, 0, 98, 241, 236, 238, 240,
	218, 0, 0, 0, 0, 0, 216, 13, 152, 0,
	230, 0, 220, 0, 0, 0, 239, 0, 232, 221,
	222, 153, 231, 0, 0, 233, 154, 0, 155, 0,
	156, 0, 157, 0, 158, 0, 159, 0, 160, 0,
	161, 0, 162, 0, 163, 0, 164, 0, 165, 0,
	166, 0, 167, 0, 168, 0, 169, 0, 170, 0,
	171, 0, 172, 0, 173, 0, 174, 0, 175, 0,
	176, 0, 177, 0, 178, 0, 179, 0, 180, 0,
	181, 0, 182, 0, 183, 0, 184, 0, 185, 0,
	186, 0, 187, 0, 188, 0, 189, 0, 190,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 175, 3, 3, 3, 163, 158, 3,
	171, 172, 161, 159, 117, 160, 173, 162, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	141, 140, 142, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 166, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 157, 3, 167,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
	52, 53, 54, 55, 56, 57, 58, 59, 60, 61,
	62, 63, 64, 65, 66, 67, 68, 69, 70, 71,
	72, 73, 74, 75, 76, 77, 78, 79, 80, 81,
	82, 83, 84, 85, 86, 87, 88, 89, 90, 91,
	92, 93, 94, 95, 96, 97, 98, 99, 100, 101,
	102, 103, 104, 105, 106, 107, 108, 109, 110, 111,
	112, 113, 114, 115, 116, 118, 119, 120, 121, 122,
	123, 124, 125, 126, 127, 128, 129, 130, 131, 132,
	133, 134, 135, 136, 137, 138, 139, 143, 144, 145,
	146, 147, 148, 149, 150, 151, 152, 153, 154, 155,
	156, 164, 165, 168, 169, 170, 174,
}
var yyTok3 = [...]int{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := yyPact[state]
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && yyChk[yyAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || yyExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := yyExca[i]
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		token = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = yyTok3[i+0]
		if token == char {
			token = yyTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yyrcvr.char = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:263
		{
			yylex.(*Tokener).ParseStmt = yyDollar[1].node
		}
	case 4:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:282
		{
			yyVAL.node = yyDollar[2].node
		}
	case 5:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:286
		{
			yyVAL.node = yyDollar[2].node
		}
	case 6:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:292
		{
			yyVAL.node = yyDollar[1].node
		}
	case 7:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:296
		{

		}
	case 8:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:300
		{
		}
	case 9:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:303
		{
		}
	case 10:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:307
		{
			yyVAL.node = yyDollar[1].node
		}
	case 11:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:308
		{
			yyVAL.node = yyDollar[1].node
		}
	case 12:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.y:313
		{
			stmt := &ast.SelectStmt{}
			stmt.Target = &ast.TargetClause{Target_ref: yyDollar[2].list}
			if yyDollar[3].node != nil {
				stmt.Into = yyDollar[3].node.(*ast.IntoClause)
			}
			if yyDollar[4].node != nil {
				stmt.From = yyDollar[4].node.(*ast.FromClause)
			}
			if yyDollar[5].node != nil {
				stmt.Where = yyDollar[5].node.(*ast.WhereClause)
			}
			if yyDollar[6].node != nil {
				stmt.Group = yyDollar[6].node.(*ast.GroupClause)
			}
			if yyDollar[7].node != nil {
				stmt.Having = yyDollar[7].node.(*ast.HavingClause)
			}
			stmt.SetTag(ast.AST_SELECT_SIMPLE)
			yyVAL.node = stmt
		}
	case 13:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line parser.y:335
		{
			stmt := &ast.SelectStmt{Distinct: yyDollar[2].node.(*ast.DistinctClause), Target: &ast.TargetClause{Target_ref: yyDollar[3].list}}
			if yyDollar[4].node != nil {
				stmt.Into = yyDollar[4].node.(*ast.IntoClause)
			}
			if yyDollar[5].node != nil {
				stmt.From = yyDollar[5].node.(*ast.FromClause)
			}
			if yyDollar[6].node != nil {
				stmt.Where = yyDollar[6].node.(*ast.WhereClause)
			}
			if yyDollar[7].node != nil {
				stmt.Group = yyDollar[7].node.(*ast.GroupClause)
			}
			if yyDollar[8].node != nil {
				stmt.Having = yyDollar[8].node.(*ast.HavingClause)
			}
			stmt.SetTag(ast.AST_SELECT_SIMPLE)
			yyVAL.node = stmt
		}
	case 14:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:356
		{
			stmt := &ast.SelectStmt{All: yyDollar[3].val, Left: yyDollar[1].node.(*ast.SelectStmt), Right: yyDollar[4].node.(*ast.SelectStmt)}
			stmt.SetTag(ast.AST_SELECT_UNION)
			yyVAL.node = stmt
		}
	case 15:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:362
		{
			stmt := &ast.SelectStmt{All: yyDollar[3].val, Left: yyDollar[1].node.(*ast.SelectStmt), Right: yyDollar[4].node.(*ast.SelectStmt)}
			stmt.SetTag(ast.AST_SELECT_INTERSECT)
			yyVAL.node = stmt
		}
	case 16:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:368
		{
			stmt := &ast.SelectStmt{All: yyDollar[3].val, Left: yyDollar[1].node.(*ast.SelectStmt), Right: yyDollar[4].node.(*ast.SelectStmt)}
			stmt.SetTag(ast.AST_SELECT_EXCEPT)
			yyVAL.node = stmt
		}
	case 17:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:384
		{
			expr := &ast.Expr{Operator: "OR", Left: yyDollar[1].expr, Right: yyDollar[3].expr}
			expr.SetTag(ast.AST_EXPR_OR)
			yyVAL.expr = expr
		}
	case 18:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:390
		{
			expr := &ast.Expr{Operator: "||", Left: yyDollar[1].expr, Right: yyDollar[3].expr}
			expr.SetTag(ast.AST_EXPR_DOUBLE_OR)
			yyVAL.expr = expr
		}
	case 19:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:396
		{
			expr := &ast.Expr{Operator: "XOR", Left: yyDollar[1].expr, Right: yyDollar[3].expr}
			expr.SetTag(ast.AST_EXPR_XOR)
			yyVAL.expr = expr
		}
	case 20:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:402
		{
			expr := &ast.Expr{Operator: "AND", Left: yyDollar[1].expr, Right: yyDollar[3].expr}
			expr.SetTag(ast.AST_EXPR_AND)
			yyVAL.expr = expr
		}
	case 21:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:408
		{
			expr := &ast.Expr{Operator: "&&", Left: yyDollar[1].expr, Right: yyDollar[3].expr}
			expr.SetTag(ast.AST_EXPR_DOUBLE_AND)
			yyVAL.expr = expr
		}
	case 22:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:414
		{
			switch node := yyDollar[2].expr.(type) {
			case *ast.SubqueryExpr:
				node.Not = true
				yyVAL.expr = node
			default:
				expr := &ast.Expr{Operator: "NOT", Left: nil, Right: yyDollar[2].expr}
				expr.SetTag(ast.AST_EXPR_NOT)
				yyVAL.expr = expr
			}

		}
	case 23:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:427
		{
			expr := &ast.Expr{Left: yyDollar[1].expr}
			if yyDollar[2].val == "IS" {
				expr.Not = true
			} else {
				expr.Not = false
			}
			switch yyDollar[3].val {
			case "TRUE":
				expr.SetTag(ast.AST_EXPR_TRUE)
			case "FALSE":
				expr.SetTag(ast.AST_EXPR_FALSE)
			case "UNKNOWN":
				expr.SetTag(ast.AST_EXPR_UNKNOWN)
			}

			yyVAL.expr = expr
		}
	case 24:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:446
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 25:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:452
		{
			expr := &ast.Expr{Left: yyDollar[1].expr}
			if yyDollar[2].val == "IS" {
				expr.Not = false
			} else {
				expr.Not = true
			}
			expr.SetTag(ast.AST_EXPR_NULL)
			yyVAL.expr = expr

		}
	case 26:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:464
		{
			expr := &ast.Expr{Operator: "<=>", Left: yyDollar[1].expr, Right: yyDollar[3].expr}
			expr.SetTag(ast.AST_EXPR_LEG)
			yyVAL.expr = expr
		}
	case 27:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:470
		{
			expr := &ast.Expr{Operator: yyDollar[2].val, Left: yyDollar[1].expr, Right: yyDollar[3].expr}
			switch yyDollar[2].val {
			case "=":
				expr.SetTag(ast.AST_EXPR_E)
			case ">=":
				expr.SetTag(ast.AST_EXPR_GE)
			case ">":
				expr.SetTag(ast.AST_EXPR_G)
			case "<=":
				expr.SetTag(ast.AST_EXPR_LE)
			case "<":
				expr.SetTag(ast.AST_EXPR_L)
			case "<>":
				expr.SetTag(ast.AST_EXPR_LG)
			case "!=":
				expr.SetTag(ast.AST_EXPR_NE)
			}
			yyVAL.expr = expr
		}
	case 28:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:491
		{
			expr := &ast.SubqueryExpr{Operator: yyDollar[2].val, Left: yyDollar[1].expr, Right: yyDollar[4].node.(*ast.SelectStmt)}
			switch yyDollar[3].val {
			case "SOME":
				expr.SetTag(ast.AST_SUBQUERY_SOME)
			case "ALL":
				expr.SetTag(ast.AST_SUBQUERY_ALL)
			case "ANY":
				expr.SetTag(ast.AST_SUBQUERY_ANY)
			}
			yyVAL.expr = expr
		}
	case 29:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:504
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 30:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:510
		{
			expr := &ast.SubqueryExpr{Operator: "IN", Left: yyDollar[1].expr, Right: yyDollar[4].node.(*ast.SelectStmt)}
			if yyDollar[2].val == "" {
				expr.Not = false
			} else {
				expr.Not = true
			}
			expr.SetTag(ast.AST_SUBQUERY_IN)
			yyVAL.expr = expr
		}
	case 31:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.y:521
		{
			expr := &ast.InExpr{Left: yyDollar[1].expr, Right: yyDollar[5].list}
			if yyDollar[2].val == "" {
				expr.Not = false
			} else {
				expr.Not = true
			}
			expr.SetTag(ast.AST_EXPR_IN)
			yyVAL.expr = expr
		}
	case 32:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.y:532
		{
			expr := &ast.BetweenExpr{Expr: yyDollar[1].expr, Left: yyDollar[4].expr, Right: yyDollar[6].expr}
			switch yyDollar[2].val {
			case "":
				expr.Not = false
			case "NOT":
				expr.Not = true
			}
			expr.SetTag(ast.AST_EXPR_BETWEEN)
			yyVAL.expr = expr
		}
	case 33:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:544
		{
			expr := &ast.Expr{Operator: "LIKE", Left: yyDollar[1].expr, Right: yyDollar[4].expr}
			expr.SetTag(ast.AST_EXPR_SOUND_LIKE)
			yyVAL.expr = expr
		}
	case 34:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:550
		{
			expr := &ast.Expr{Operator: "LIKE", Left: yyDollar[1].expr, Right: yyDollar[4].expr}
			switch yyDollar[2].val {
			case "":
				expr.Not = false
			case "NOT":
				expr.Not = true
			}
			expr.SetTag(ast.AST_EXPR_LIKE)
			yyVAL.expr = expr
		}
	case 35:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:562
		{
			expr := &ast.Expr{Operator: "REGEXP", Left: yyDollar[1].expr, Right: yyDollar[4].expr}
			switch yyDollar[2].val {
			case "":
				expr.Not = false
			case "NOT":
				expr.Not = true
			}
			expr.SetTag(ast.AST_EXPR_REGEXP)
			yyVAL.expr = expr
		}
	case 36:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:574
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 37:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:581
		{
			expr := &ast.Expr{Operator: "|", Left: yyDollar[1].expr, Right: yyDollar[3].expr}
			expr.SetTag(ast.AST_EXPR_VETICEL)
			yyVAL.expr = expr
		}
	case 38:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:587
		{
			expr := &ast.Expr{Operator: "&", Left: yyDollar[1].expr, Right: yyDollar[3].expr}
			expr.SetTag(ast.AST_EXPR_AMPERSAND)
			yyVAL.expr = expr
		}
	case 39:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:593
		{
			expr := &ast.Expr{Operator: "<<", Left: yyDollar[1].expr, Right: yyDollar[3].expr}
			expr.SetTag(ast.AST_EXPR_SL)
			yyVAL.expr = expr
		}
	case 40:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:599
		{
			expr := &ast.Expr{Operator: ">>", Left: yyDollar[1].expr, Right: yyDollar[3].expr}
			expr.SetTag(ast.AST_EXPR_SR)
			yyVAL.expr = expr
		}
	case 41:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:605
		{
			expr := &ast.Expr{Operator: "+", Left: yyDollar[1].expr, Right: yyDollar[3].expr}
			expr.SetTag(ast.AST_EXPR_PLUS)
			yyVAL.expr = expr
		}
	case 42:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:611
		{
			expr := &ast.Expr{Operator: "-", Left: yyDollar[1].expr, Right: yyDollar[3].expr}
			expr.SetTag(ast.AST_EXPR_MINUS)
			yyVAL.expr = expr
		}
	case 43:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:617
		{
			expr := &ast.Expr{Operator: "*", Left: yyDollar[1].expr, Right: yyDollar[3].expr}
			expr.SetTag(ast.AST_EXPR_MULT)
			yyVAL.expr = expr
		}
	case 44:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:623
		{
			expr := &ast.Expr{Operator: "/", Left: yyDollar[1].expr, Right: yyDollar[3].expr}
			expr.SetTag(ast.AST_EXPR_DIV)
			yyVAL.expr = expr
		}
	case 45:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:629
		{
			expr := &ast.Expr{Operator: "DIV", Left: yyDollar[1].expr, Right: yyDollar[3].expr}
			expr.SetTag(ast.AST_EXPR_DIV)
			yyVAL.expr = expr
		}
	case 46:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:635
		{
			expr := &ast.Expr{Operator: "MOD", Left: yyDollar[1].expr, Right: yyDollar[3].expr}
			expr.SetTag(ast.AST_EXPR_MOD)
			yyVAL.expr = expr
		}
	case 47:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:641
		{
			expr := &ast.Expr{Operator: "%", Left: yyDollar[1].expr, Right: yyDollar[3].expr}
			expr.SetTag(ast.AST_EXPR_MOD)
			yyVAL.expr = expr
		}
	case 48:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:647
		{
			expr := &ast.Expr{Operator: "^", Left: yyDollar[1].expr, Right: yyDollar[3].expr}
			expr.SetTag(ast.AST_EXPR_XOR)
			yyVAL.expr = expr
		}
	case 49:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:653
		{
			expr := &ast.Expr{Operator: "+", Left: yyDollar[1].expr, Right: yyDollar[4].expr}
			expr.SetTag(ast.AST_TIME_INTERVAL)
			yyVAL.expr = expr
		}
	case 50:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:659
		{
			expr := &ast.Expr{Operator: "-", Left: yyDollar[1].expr, Right: yyDollar[4].expr}
			expr.SetTag(ast.AST_TIME_INTERVAL)
			yyVAL.expr = expr
		}
	case 51:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:665
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 52:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:673
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 53:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:678
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 54:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:683
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 55:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:687
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 56:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:691
		{
			switch node := yyDollar[1].expr.(type) {
			case *ast.ValExpr:
				node.SetCollate(yyDollar[3].val)
			case *ast.ColumnExpr:
				node.SetCollate(yyDollar[3].val)
			case *ast.Expr:
				node.SetCollate(yyDollar[3].val)
			}
			yyVAL.expr = yyDollar[1].expr
		}
	case 57:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:703
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 58:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:708
		{
			expr := &ast.Expr{Operator: "||", Left: yyDollar[1].expr, Right: yyDollar[3].expr}
			expr.SetTag(ast.AST_EXPR_DOUBLE_OR)
			yyVAL.expr = expr
		}
	case 59:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:714
		{
			yyVAL.expr = yyDollar[2].expr.(*ast.ValExpr)
		}
	case 60:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:718
		{
			val := yyDollar[2].expr.(*ast.ValExpr)
			val.Symbol = 0 - val.Symbol
			yyVAL.expr = val
		}
	case 61:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:724
		{
			expr := &ast.Expr{Operator: "~", Left: nil, Right: yyDollar[2].expr}
			expr.SetTag(ast.AST_EXPR_TILDE)
			yyVAL.expr = expr
		}
	case 62:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:730
		{
			expr := &ast.Expr{Operator: "!", Left: nil, Right: yyDollar[2].expr}
			expr.SetTag(ast.AST_EXPR_EXCLAMATION_MARK)
			yyVAL.expr = expr
		}
	case 63:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:736
		{
			yyVAL.expr = yyDollar[2].expr
		}
	case 64:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:741
		{
			expr := &ast.SubqueryExpr{Left: nil, Right: yyDollar[1].node.(*ast.SelectStmt)}
			expr.SetTag(ast.AST_SUBQUERY)
			yyVAL.expr = expr
		}
	case 65:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:747
		{
			expr := &ast.SubqueryExpr{Operator: "EXISTS", Left: nil, Right: yyDollar[2].node.(*ast.SelectStmt)}
			expr.SetTag(ast.AST_SUBQUERY_EXISTS)
			yyVAL.expr = expr
		}
	case 66:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:753
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 67:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:759
		{
			expr := &ast.ValExpr{Symbol: 1, Sval: yyDollar[1].val}
			expr.SetTag(ast.AST_VALUE_BIT_NUMBER)
			yyVAL.expr = expr
		}
	case 68:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:765
		{
			expr := &ast.ValExpr{Symbol: 1, Sval: yyDollar[1].val}
			expr.SetTag(ast.AST_VALUE_HEX_NUMBER)
			yyVAL.expr = expr
		}
	case 69:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:771
		{
			expr := &ast.ValExpr{Symbol: 1, Sval: yyDollar[1].val}
			expr.SetTag(ast.AST_VALUE_NUMBER)
			yyVAL.expr = expr
		}
	case 70:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:777
		{
			expr := &ast.ValExpr{Symbol: 0, Sval: yyDollar[1].str}
			expr.SetTag(ast.AST_VALUE_STRING)
			yyVAL.expr = expr
		}
	case 71:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:783
		{
			expr := &ast.ValExpr{Symbol: 0, Sval: yyDollar[1].str}
			expr.SetTag(ast.AST_VALUE_SINGLE_QUOTA_STRING)
			yyVAL.expr = expr
		}
	case 72:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:789
		{
			expr := &ast.ValExpr{Symbol: 0, Sval: yyDollar[1].str}
			expr.SetTag(ast.AST_VALUE_DOUBLE_QUOTA_STRING)
			yyVAL.expr = expr
		}
	case 73:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:797
		{
			yyVAL.val = yyDollar[1].ident
		}
	case 74:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:804
		{
			expr := &ast.ColumnExpr{Column: yyDollar[1].ident}
			expr.SetTag(ast.AST_EXPR_COLUMN)
			yyVAL.expr = expr
		}
	case 75:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:810
		{
			expr := &ast.ColumnExpr{Column: yyDollar[1].ident}
			expr.SetTag(ast.AST_EXPR_COLUMN)
			yyVAL.expr = expr
		}
	case 76:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:816
		{
			expr := &ast.ColumnExpr{Column: yyDollar[1].ident}
			expr.SetTag(ast.AST_EXPR_COLUMN)
			yyVAL.expr = expr
		}
	case 77:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:825
		{
			expr := &ast.ColumnExpr{Column: yyDollar[1].val}
			expr.SetTag(ast.AST_EXPR_COLUMN)
			yyVAL.expr = expr
		}
	case 78:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:831
		{
			expr := &ast.ColumnExpr{Table: yyDollar[1].val, Column: yyDollar[2].val}
			expr.SetTag(ast.AST_EXPR_COLUMN)
			yyVAL.expr = expr
		}
	case 79:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:837
		{
			expr := &ast.ColumnExpr{DB: yyDollar[1].val, Table: yyDollar[2].val, Column: yyDollar[3].val}
			expr.SetTag(ast.AST_EXPR_COLUMN)
			yyVAL.expr = expr
		}
	case 80:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:846
		{
			expr := &ast.FuncExpr{Name: yyDollar[1].val, Arg: yyDollar[3].list}
			expr.SetTag(ast.AST_EXPR_FUNC)
			yyVAL.expr = expr
		}
	case 81:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:852
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 82:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:858
		{
			expr := &ast.FuncExpr{Name: "AVG", Arg: ast.List{yyDollar[4].expr}}
			switch yyDollar[3].val {
			case "":
				expr.SetTag(ast.AST_EXPR_AVG)
			case "DISTINCT":
				expr.SetTag(ast.AST_EXPR_AVG_DISTINCT)
			case "DISTINCTROW":
				expr.SetTag(ast.AST_EXPR_AVG_DISTINCTROW)
			case "ALL":
				expr.SetTag(ast.AST_EXPR_AVG_ALL)
			}
			yyVAL.expr = expr
		}
	case 83:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:873
		{
			expr := &ast.FuncExpr{Name: "COUNT", Arg: yyDollar[4].list}
			switch yyDollar[3].val {
			case "DISTINCT":
				expr.SetTag(ast.AST_EXPR_COUNT_DISTINCT)
			case "DISTINCTROW":
				expr.SetTag(ast.AST_EXPR_COUNT_DISTINCTROW)
			}
			yyVAL.expr = expr
		}
	case 84:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:884
		{
			expr := &ast.FuncExpr{Name: "COUNT", Arg: ast.List{yyDollar[4].expr}}
			expr.SetTag(ast.AST_EXPR_COUNT_ALL)
			yyVAL.expr = expr
		}
	case 85:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:890
		{
			expr := &ast.FuncExpr{Name: "COUNT", Arg: ast.List{yyDollar[3].expr}}
			expr.SetTag(ast.AST_EXPR_COUNT)
			yyVAL.expr = expr
		}
	case 86:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:896
		{
			expr := &ast.FuncExpr{Name: "COUNT"}
			expr.SetTag(ast.AST_EXPR_COUNT_STAR)
			yyVAL.expr = expr
		}
	case 87:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:902
		{
			expr := &ast.FuncExpr{Name: "MAX", Arg: ast.List{yyDollar[4].expr}}
			switch yyDollar[3].val {
			case "":
				expr.SetTag(ast.AST_EXPR_MAX)
			case "DISTINCT":
				expr.SetTag(ast.AST_EXPR_MAX_DISTINCT)
			case "DISTINCTROW":
				expr.SetTag(ast.AST_EXPR_MAX_DISTINCTROW)
			case "ALL":
				expr.SetTag(ast.AST_EXPR_MAX_ALL)
			}
			yyVAL.expr = expr
		}
	case 88:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:917
		{
			expr := &ast.FuncExpr{Name: "MIN", Arg: ast.List{yyDollar[4].expr}}
			switch yyDollar[3].val {
			case "":
				expr.SetTag(ast.AST_EXPR_MIN)
			case "DISTINCT":
				expr.SetTag(ast.AST_EXPR_MIN_DISTINCT)
			case "DISTINCTROW":
				expr.SetTag(ast.AST_EXPR_MIN_DISTINCTROW)
			case "ALL":
				expr.SetTag(ast.AST_EXPR_MIN_ALL)
			}
			yyVAL.expr = expr
		}
	case 89:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:932
		{
			expr := &ast.FuncExpr{Name: "SUM", Arg: ast.List{yyDollar[4].expr}}
			switch yyDollar[3].val {
			case "":
				expr.SetTag(ast.AST_EXPR_MAX)
			case "DISTINCT":
				expr.SetTag(ast.AST_EXPR_MAX_DISTINCT)
			case "DISTINCTROW":
				expr.SetTag(ast.AST_EXPR_MAX_DISTINCTROW)
			case "ALL":
				expr.SetTag(ast.AST_EXPR_MAX_ALL)
			}
			yyVAL.expr = expr
		}
	case 90:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:950
		{
			yyVAL.list = ast.List{yyDollar[1].expr}
		}
	case 91:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:954
		{
			yyVAL.list = append(yyDollar[1].list, yyDollar[3].expr)
		}
	case 92:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:962
		{
			expr := &ast.ValExpr{Sval: yyDollar[1].str}
			expr.SetTag(ast.AST_VALUE_MARKER)
			yyVAL.expr = expr
		}
	case 93:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:972
		{
			expr := &ast.CaseExpr{Case: yyDollar[2].expr, When: yyDollar[3].list, Else: yyDollar[4].expr}
			expr.SetTag(ast.AST_EXPR_CASE)
			yyVAL.expr = expr
		}
	case 94:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:980
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 95:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:984
		{
			yyVAL.expr = nil
		}
	case 96:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:991
		{
			yyVAL.list = ast.List{yyDollar[1].expr}
		}
	case 97:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:995
		{
			yyVAL.list = append(yyDollar[1].list, yyDollar[2].expr)
		}
	case 98:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:1002
		{
			expr := &ast.Expr{Left: yyDollar[2].expr, Right: yyDollar[4].expr}
			expr.SetTag(ast.AST_EXPR_WHEN)
			yyVAL.expr = expr
		}
	case 99:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:1011
		{
			expr := &ast.Expr{Left: yyDollar[2].expr}
			expr.SetTag(ast.AST_EXPR_ELSE)
			yyVAL.expr = expr
		}
	case 100:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:1017
		{
			yyVAL.expr = nil
		}
	case 101:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1025
		{
			yyVAL.list = yyDollar[1].list
		}
	case 102:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:1029
		{
			yyVAL.list = nil
		}
	case 103:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1035
		{
			yyVAL.list = ast.List{yyDollar[1].expr}

		}
	case 104:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:1040
		{
			yyVAL.list = append(yyDollar[1].list, yyDollar[3].expr)
		}
	case 105:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:1047
		{
			switch node := yyDollar[1].expr.(type) {
			case *ast.Expr:
				node.SetAlias(yyDollar[2].val)
			case *ast.SubqueryExpr:
				node.SetAlias(yyDollar[2].val)
			case *ast.CaseExpr:
				node.SetAlias(yyDollar[2].val)
			case *ast.FuncExpr:
				node.SetAlias(yyDollar[2].val)
			case *ast.ValExpr:
				node.SetAlias(yyDollar[2].val)
			case *ast.ColumnExpr:
				node.SetAlias(yyDollar[2].val)
			}
			yyVAL.expr = yyDollar[1].expr
		}
	case 106:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1065
		{
			expr := &ast.ColumnExpr{Column: "*"}
			expr.SetTag(ast.AST_EXPR_COLUMN_STAR)
			yyVAL.expr = expr
		}
	case 107:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1073
		{
			yyVAL.val = yyDollar[1].val
		}
	case 108:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:1074
		{
			yyVAL.val = ""
		}
	case 109:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:1077
		{
			yyVAL.val = yyDollar[2].val
		}
	case 110:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1078
		{
			yyVAL.val = yyDollar[1].val
		}
	case 111:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1081
		{
			yyVAL.val = yyDollar[1].val
		}
	case 112:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1082
		{
			yyVAL.val = yyDollar[1].str
		}
	case 113:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1083
		{
			yyVAL.val = yyDollar[1].str
		}
	case 114:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1084
		{
			yyVAL.val = yyDollar[1].str
		}
	case 115:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1088
		{
			yyVAL.val = "TRUE"
		}
	case 116:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1089
		{
			yyVAL.val = "FALSE"
		}
	case 117:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1093
		{
			yyVAL.val = "SOME"
		}
	case 118:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1094
		{
			yyVAL.val = "ANY"
		}
	case 119:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1095
		{
			yyVAL.val = "ALL"
		}
	case 120:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1098
		{
			yyVAL.val = "="
		}
	case 121:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1099
		{
			yyVAL.val = ">="
		}
	case 122:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1100
		{
			yyVAL.val = ">"
		}
	case 123:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1101
		{
			yyVAL.val = "<="
		}
	case 124:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1102
		{
			yyVAL.val = "<"
		}
	case 125:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1103
		{
			yyVAL.val = "<>"
		}
	case 126:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1104
		{
			yyVAL.val = "!="
		}
	case 127:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1107
		{
			yyVAL.val = "IS"
		}
	case 128:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:1108
		{
			yyVAL.val = "IS NOT"
		}
	case 129:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1112
		{
			yyVAL.val = "NOT"
		}
	case 130:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:1113
		{
			yyVAL.val = ""
		}
	case 131:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1116
		{
			yyVAL.val = yyDollar[1].ident
		}
	case 132:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1117
		{
			yyVAL.val = yyDollar[1].ident
		}
	case 133:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1118
		{
			yyVAL.val = yyDollar[1].ident
		}
	case 134:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1119
		{
			yyVAL.val = yyDollar[1].ident
		}
	case 135:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1120
		{
			yyVAL.val = yyDollar[1].ident
		}
	case 136:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1121
		{
			yyVAL.val = yyDollar[1].ident
		}
	case 137:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1122
		{
			yyVAL.val = yyDollar[1].ident
		}
	case 138:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1123
		{
			yyVAL.val = yyDollar[1].ident
		}
	case 139:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1124
		{
			yyVAL.val = yyDollar[1].ident
		}
	case 140:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1125
		{
			yyVAL.val = yyDollar[1].ident
		}
	case 141:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1126
		{
			yyVAL.val = yyDollar[1].ident
		}
	case 142:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1127
		{
			yyVAL.val = yyDollar[1].ident
		}
	case 143:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1128
		{
			yyVAL.val = yyDollar[1].ident
		}
	case 144:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1129
		{
			yyVAL.val = yyDollar[1].ident
		}
	case 145:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1130
		{
			yyVAL.val = yyDollar[1].ident
		}
	case 146:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1131
		{
			yyVAL.val = yyDollar[1].ident
		}
	case 147:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1132
		{
			yyVAL.val = yyDollar[1].ident
		}
	case 148:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1133
		{
			yyVAL.val = yyDollar[1].ident
		}
	case 149:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1134
		{
			yyVAL.val = yyDollar[1].ident
		}
	case 150:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1138
		{
			yyVAL.val = yyDollar[1].ident
		}
	case 151:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:1139
		{
			yyVAL.val = yyDollar[1].ident
		}
	case 152:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:1140
		{
			yyVAL.val = yyDollar[1].ident
		}
	case 153:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.y:1141
		{
			yyVAL.val = yyDollar[1].ident
		}
	case 154:
		yyDollar = yyS[yypt-9 : yypt+1]
		//line parser.y:1142
		{
			yyVAL.val = yyDollar[1].ident
		}
	case 155:
		yyDollar = yyS[yypt-11 : yypt+1]
		//line parser.y:1143
		{
			yyVAL.val = yyDollar[1].ident
		}
	case 156:
		yyDollar = yyS[yypt-13 : yypt+1]
		//line parser.y:1144
		{
			yyVAL.val = yyDollar[1].ident
		}
	case 157:
		yyDollar = yyS[yypt-15 : yypt+1]
		//line parser.y:1145
		{
			yyVAL.val = yyDollar[1].ident
		}
	case 158:
		yyDollar = yyS[yypt-17 : yypt+1]
		//line parser.y:1146
		{
			yyVAL.val = yyDollar[1].ident
		}
	case 159:
		yyDollar = yyS[yypt-19 : yypt+1]
		//line parser.y:1147
		{
			yyVAL.val = yyDollar[1].ident
		}
	case 160:
		yyDollar = yyS[yypt-21 : yypt+1]
		//line parser.y:1148
		{
			yyVAL.val = yyDollar[1].ident
		}
	case 161:
		yyDollar = yyS[yypt-23 : yypt+1]
		//line parser.y:1149
		{
			yyVAL.val = yyDollar[1].ident
		}
	case 162:
		yyDollar = yyS[yypt-25 : yypt+1]
		//line parser.y:1150
		{
			yyVAL.val = yyDollar[1].ident
		}
	case 163:
		yyDollar = yyS[yypt-27 : yypt+1]
		//line parser.y:1151
		{
			yyVAL.val = yyDollar[1].ident
		}
	case 164:
		yyDollar = yyS[yypt-29 : yypt+1]
		//line parser.y:1152
		{
			yyVAL.val = yyDollar[1].ident
		}
	case 165:
		yyDollar = yyS[yypt-31 : yypt+1]
		//line parser.y:1153
		{
			yyVAL.val = yyDollar[1].ident
		}
	case 166:
		yyDollar = yyS[yypt-33 : yypt+1]
		//line parser.y:1154
		{
			yyVAL.val = yyDollar[1].ident
		}
	case 167:
		yyDollar = yyS[yypt-35 : yypt+1]
		//line parser.y:1155
		{
			yyVAL.val = yyDollar[1].ident
		}
	case 168:
		yyDollar = yyS[yypt-37 : yypt+1]
		//line parser.y:1156
		{
			yyVAL.val = yyDollar[1].ident
		}
	case 169:
		yyDollar = yyS[yypt-39 : yypt+1]
		//line parser.y:1157
		{
			yyVAL.val = yyDollar[1].ident
		}
	case 170:
		yyDollar = yyS[yypt-41 : yypt+1]
		//line parser.y:1158
		{
			yyVAL.val = yyDollar[1].ident
		}
	case 171:
		yyDollar = yyS[yypt-43 : yypt+1]
		//line parser.y:1159
		{
			yyVAL.val = yyDollar[1].ident
		}
	case 172:
		yyDollar = yyS[yypt-45 : yypt+1]
		//line parser.y:1160
		{
			yyVAL.val = yyDollar[1].ident
		}
	case 173:
		yyDollar = yyS[yypt-47 : yypt+1]
		//line parser.y:1161
		{
			yyVAL.val = yyDollar[1].ident
		}
	case 174:
		yyDollar = yyS[yypt-49 : yypt+1]
		//line parser.y:1162
		{
			yyVAL.val = yyDollar[1].ident
		}
	case 175:
		yyDollar = yyS[yypt-51 : yypt+1]
		//line parser.y:1163
		{
			yyVAL.val = yyDollar[1].ident
		}
	case 176:
		yyDollar = yyS[yypt-53 : yypt+1]
		//line parser.y:1164
		{
			yyVAL.val = yyDollar[1].ident
		}
	case 177:
		yyDollar = yyS[yypt-55 : yypt+1]
		//line parser.y:1165
		{
			yyVAL.val = yyDollar[1].ident
		}
	case 178:
		yyDollar = yyS[yypt-57 : yypt+1]
		//line parser.y:1166
		{
			yyVAL.val = yyDollar[1].ident
		}
	case 179:
		yyDollar = yyS[yypt-59 : yypt+1]
		//line parser.y:1167
		{
			yyVAL.val = yyDollar[1].ident
		}
	case 180:
		yyDollar = yyS[yypt-61 : yypt+1]
		//line parser.y:1168
		{
			yyVAL.val = yyDollar[1].ident
		}
	case 181:
		yyDollar = yyS[yypt-63 : yypt+1]
		//line parser.y:1169
		{
			yyVAL.val = yyDollar[1].ident
		}
	case 182:
		yyDollar = yyS[yypt-65 : yypt+1]
		//line parser.y:1170
		{
			yyVAL.val = yyDollar[1].ident
		}
	case 183:
		yyDollar = yyS[yypt-67 : yypt+1]
		//line parser.y:1171
		{
			yyVAL.val = yyDollar[1].ident
		}
	case 184:
		yyDollar = yyS[yypt-69 : yypt+1]
		//line parser.y:1172
		{
			yyVAL.val = yyDollar[1].ident
		}
	case 185:
		yyDollar = yyS[yypt-71 : yypt+1]
		//line parser.y:1173
		{
			yyVAL.val = yyDollar[1].ident
		}
	case 186:
		yyDollar = yyS[yypt-73 : yypt+1]
		//line parser.y:1174
		{
			yyVAL.val = yyDollar[1].ident
		}
	case 187:
		yyDollar = yyS[yypt-75 : yypt+1]
		//line parser.y:1175
		{
			yyVAL.val = yyDollar[1].ident
		}
	case 188:
		yyDollar = yyS[yypt-77 : yypt+1]
		//line parser.y:1176
		{
			yyVAL.val = yyDollar[1].ident
		}
	case 189:
		yyDollar = yyS[yypt-79 : yypt+1]
		//line parser.y:1177
		{
			yyVAL.val = yyDollar[1].ident
		}
	case 190:
		yyDollar = yyS[yypt-81 : yypt+1]
		//line parser.y:1178
		{
			yyVAL.val = yyDollar[1].ident
		}
	case 191:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1190
		{
			clause := &ast.DistinctClause{}
			switch yyDollar[1].val {
			case "DISTINCT":
				clause.SetTag(ast.AST_CLAUSE_DISTINCT)
			case "DISTINCTROW":
				clause.SetTag(ast.AST_CLAUSE_DISTINCTROW)
			}
			yyVAL.node = clause
		}
	case 192:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1206
		{
			yyVAL.val = yyDollar[1].ident
		}
	case 193:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1210
		{
			yyVAL.val = yyDollar[1].val
		}
	case 194:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:1214
		{
			yyVAL.val = ""
		}
	case 195:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1219
		{
			yyVAL.val = "DISTINCT"
		}
	case 196:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1223
		{
			yyVAL.val = "DISTINCTROW"
		}
	case 197:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:1235
		{
			yyVAL.node = &ast.IntoClause{Table: yyDollar[2].val}
		}
	case 198:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:1239
		{
			yyVAL.node = nil
		}
	case 199:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1245
		{
			yyVAL.val = yyDollar[1].val
		}
	case 200:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1249
		{
			yyVAL.val = yyDollar[1].str
		}
	case 201:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1253
		{
			yyVAL.val = yyDollar[1].str
		}
	case 202:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:1268
		{
			yyVAL.node = &ast.FromClause{Table_ref: yyDollar[2].list}
		}
	case 203:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:1272
		{
			yyVAL.node = nil
		}
	case 204:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1278
		{
			yyVAL.list = ast.List{yyDollar[1].expr}
		}
	case 205:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:1282
		{
			yyVAL.list = append(yyDollar[1].list, yyDollar[3].expr)
		}
	case 206:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1288
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 207:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1292
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 208:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:1298
		{
			switch node := yyDollar[1].expr.(type) {
			case *ast.RelationExpr:
				node.SetAlias(yyDollar[2].val)
			case *ast.JoinExpr:
				node.SetAlias(yyDollar[2].val)
			case *ast.RelationListExpr:
				node.SetAlias(yyDollar[2].val)
			}
			yyVAL.expr = yyDollar[1].expr
		}
	case 209:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:1310
		{
			expr := &ast.SubqueryExpr{Left: nil, Right: yyDollar[1].node.(*ast.SelectStmt)}
			expr.SetTag(ast.AST_SUBQUERY)
			expr.SetAlias(yyDollar[2].val)
			yyVAL.expr = expr
		}
	case 210:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:1317
		{
			yyVAL.expr = &ast.RelationListExpr{Relation: yyDollar[2].list}
		}
	case 211:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1323
		{
			tb := &ast.RelationExpr{Table: yyDollar[1].val}
			tb.SetTag(ast.AST_RELATION)
			yyVAL.expr = tb
		}
	case 212:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:1329
		{
			tb := &ast.RelationExpr{DB: yyDollar[1].val, Table: yyDollar[2].val}
			tb.SetTag(ast.AST_RELATION)
			yyVAL.expr = tb
		}
	case 213:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1337
		{
			yyVAL.val = yyDollar[1].val
		}
	case 214:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1341
		{
			yyVAL.val = yyDollar[1].str
		}
	case 215:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1348
		{
			yyVAL.list = ast.List{yyDollar[1].expr}
		}
	case 216:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:1352
		{
			yyVAL.list = append(yyDollar[1].list, yyDollar[3].expr)
		}
	case 217:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:1361
		{
			join := &ast.JoinExpr{Left: yyDollar[1].expr, Right: yyDollar[3].expr}
			switch yyDollar[2].val {
			case "JOIN":
				join.SetTag(ast.AST_JOIN)
			case "INNER JOIN":
				join.SetTag(ast.AST_INNER_JOIN)
				join.Jtype = "INNER"
			case "CROSS JOIN":
				join.SetTag(ast.AST_INNER_JOIN)
				join.Jtype = "CROSS"
			}
			yyVAL.expr = join
		}
	case 218:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:1376
		{
			join := &ast.JoinExpr{Left: yyDollar[1].expr, Right: yyDollar[3].expr, Jqual: yyDollar[4].expr}
			switch yyDollar[2].val {
			case "JOIN":
				join.SetTag(ast.AST_JOIN)
			case "INNER JOIN":
				join.SetTag(ast.AST_INNER_JOIN)
				join.Jtype = "INNER"
			case "CROSS JOIN":
				join.SetTag(ast.AST_INNER_JOIN)
				join.Jtype = "CROSS"
			}
			yyVAL.expr = join
		}
	case 219:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:1391
		{

			join := &ast.JoinExpr{Jtype: "STRAIGHT_JOIN", Left: yyDollar[1].expr, Right: yyDollar[3].expr}
			join.SetTag(ast.AST_STRAIGHT_JOIN)
			yyVAL.expr = join
		}
	case 220:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:1398
		{
			join := &ast.JoinExpr{Jtype: "STRAIGHT_JOIN", Left: yyDollar[1].expr, Right: yyDollar[3].expr, Jqual: yyDollar[5].expr}
			join.SetTag(ast.AST_STRAIGHT_JOIN)
			yyVAL.expr = join
		}
	case 221:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.y:1404
		{
			join := &ast.JoinExpr{Jtype: yyDollar[2].val, Outer: yyDollar[3].val, Left: yyDollar[1].expr, Right: yyDollar[5].expr, Jqual: yyDollar[6].expr}
			switch yyDollar[2].val {
			case "LEFT":
				join.SetTag(ast.AST_LEFT_JOIN)
			case "RIGHT":
				join.SetTag(ast.AST_RIGHT_JOIN)
			}
			yyVAL.expr = join
		}
	case 222:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.y:1415
		{
			join := &ast.JoinExpr{Jtype: yyDollar[3].val, Outer: yyDollar[4].val, Left: yyDollar[1].expr, Right: yyDollar[6].expr}
			switch yyDollar[3].val {
			case "LEFT":
				join.SetTag(ast.AST_NATURAL_LEFT_JOIN)
			case "RIGHT":
				join.SetTag(ast.AST_NATURAL_RIGHT_JOIN)
			}
			yyVAL.expr = join
		}
	case 223:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1428
		{
			yyVAL.val = "JOIN"
		}
	case 224:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:1432
		{
			yyVAL.val = "INNER JOIN"
		}
	case 225:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:1436
		{
			yyVAL.val = "CROSS JOIN"
		}
	case 226:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1443
		{
			yyVAL.val = "LEFT"
		}
	case 227:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1447
		{
			yyVAL.val = "RIGHT"
		}
	case 228:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1454
		{
			yyVAL.val = "OUTER"
		}
	case 229:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:1458
		{
			yyVAL.val = ""
		}
	case 230:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:1464
		{
			yyVAL.expr = yyDollar[2].expr
		}
	case 231:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:1468
		{
			yyVAL.expr = &ast.RelationListExpr{Relation: yyDollar[3].list}
		}
	case 232:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1475
		{
			yyVAL.list = ast.List{yyDollar[1].expr}
		}
	case 233:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:1479
		{
			yyVAL.list = append(yyDollar[1].list, yyDollar[3].expr)
		}
	case 234:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:1492
		{
			yyVAL.node = &ast.WhereClause{Where: yyDollar[2].expr}
		}
	case 235:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:1496
		{
			yyVAL.node = nil
		}
	case 236:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:1508
		{
			yyVAL.node = &ast.GroupClause{Target_ref: yyDollar[3].list}
		}
	case 237:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:1512
		{
			yyVAL.node = nil
		}
	case 238:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1519
		{
			yyVAL.list = ast.List{yyDollar[1].expr}
		}
	case 239:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:1523
		{
			yyVAL.list = append(yyDollar[1].list, yyDollar[3].expr)
		}
	case 240:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1530
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 241:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:1544
		{
			yyVAL.node = &ast.HavingClause{Having: yyDollar[2].expr}
		}
	case 242:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:1548
		{
			yyVAL.node = nil
		}
	case 243:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1563
		{
			yyVAL.node = yyDollar[1].node
		}
	case 244:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:1567
		{
			yyVAL.node = nil
		}
	case 245:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:1573
		{
			yyVAL.node = &ast.SortClause{Target_ref: yyDollar[3].list}
		}
	case 246:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1580
		{
			yyVAL.list = ast.List{yyDollar[1].expr}
		}
	case 247:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:1584
		{
			yyVAL.list = append(yyDollar[1].list, yyDollar[3].expr)
		}
	case 248:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:1590
		{
			switch node := yyDollar[1].expr.(type) {
			case *ast.Expr:
				node.SetAsc(yyDollar[2].val)
			}

			yyVAL.expr = yyDollar[1].expr
		}
	case 249:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1601
		{
			yyVAL.val = "ASC"
		}
	case 250:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1605
		{
			yyVAL.val = "DESC"
		}
	case 251:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:1609
		{
			yyVAL.val = ""
		}
	case 252:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1620
		{
			yyVAL.node = yyDollar[1].node
		}
	case 253:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:1621
		{
			yyVAL.node = nil
		}
	case 254:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:1625
		{
			yyVAL.node = &ast.LockClause{Lock: "FOR UPDATE"}
		}
	case 255:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:1629
		{
			yyVAL.node = &ast.LockClause{Lock: "IN SHARE MODE"}
		}
	case 256:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1642
		{
			yyVAL.node = yyDollar[1].node
		}
	case 257:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:1646
		{
			yyVAL.node = nil
		}
	case 258:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:1653
		{
			yyVAL.node = &ast.LimitClause{Limit: yyDollar[1].list, Offset: yyDollar[2].expr}
		}
	case 259:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:1657
		{
			yyVAL.node = &ast.LimitClause{Limit: yyDollar[2].list, Offset: yyDollar[1].expr}
		}
	case 260:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1661
		{
			yyVAL.node = &ast.LimitClause{Limit: yyDollar[1].list}
		}
	case 261:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1665
		{
			yyVAL.node = &ast.LimitClause{Offset: yyDollar[1].expr}
		}
	case 262:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:1672
		{
			yyVAL.list = yyDollar[2].list
		}
	case 263:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:1676
		{
			yyVAL.list = append(yyDollar[2].list, yyDollar[4].expr)
		}
	case 264:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1682
		{
			yyVAL.list = ast.List{yyDollar[1].expr}
		}
	case 265:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1689
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 266:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:1696
		{
			yyVAL.expr = yyDollar[2].expr
		}
	}
	goto yystack /* stack new state and value */
}
