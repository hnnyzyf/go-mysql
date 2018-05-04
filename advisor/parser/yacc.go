//line parser.y:2
package parser

import __yyfmt__ "fmt"

//line parser.y:4
import "github.com/hnnyzyf/mysql-go/advisor/ast"
import "strings"

//line parser.y:11
type yySymType struct {
	yys     int
	ident   string
	str     string
	val     string
	tag     int
	boolean bool

	node ast.Node
	list []ast.ExprNode
	expr ast.ExprNode
}

const BuiltinCharacterIdent = 57346
const BuiltinFucTimeAddIdent = 57347
const BuiltinFucTimeSubIdent = 57348
const BuiltinTimeUnitIdent = 57349
const IDENT = 57350
const DOUBLE_QUOTA_STRING = 57351
const PARAM_MARKER = 57352
const SINGLE_QUOTA_STRING = 57353
const STRING = 57354
const BIT_NUMBER = 57355
const HEX_NUMBER = 57356
const INTEGER = 57357
const FLOAT = 57358
const ALL = 57359
const ANY = 57360
const AS = 57361
const ASC = 57362
const AVG = 57363
const BY = 57364
const CAST = 57365
const CHARACTER = 57366
const COMMENT = 57367
const CONVERT = 57368
const COUNT = 57369
const DESC = 57370
const DISTINCT = 57371
const DISTINCTROW = 57372
const FALSE = 57373
const FOR = 57374
const FROM = 57375
const GROUP = 57376
const HAVING = 57377
const IF = 57378
const INTO = 57379
const LIMIT = 57380
const LOCK = 57381
const MAX = 57382
const MIN = 57383
const MODE = 57384
const NULL = 57385
const OFFSET = 57386
const ORDER = 57387
const POSITION = 57388
const QUARTER = 57389
const SELECT = 57390
const SET = 57391
const SHARE = 57392
const SOME = 57393
const SOUNDS = 57394
const SUM = 57395
const TO = 57396
const TRUE = 57397
const UPDATE = 57398
const WHERE = 57399
const EXISTS = 57400
const OUTER = 57401
const LOW = 57402
const UNION = 57403
const EXCEPT = 57404
const INTERSECT = 57405
const JOIN = 57406
const STRAIGHT_JOIN = 57407
const NATURAL = 57408
const LEFT = 57409
const RIGHT = 57410
const INNER = 57411
const CROSS = 57412
const USE = 57413
const TABLE_REF_PRIORITY = 57414
const ON = 57415
const USING = 57416
const OR = 57417
const OO = 57418
const XOR = 57419
const AND = 57420
const AA = 57421
const BETWEEN = 57422
const CASE = 57423
const WHEN = 57424
const THEN = 57425
const ELSE = 57426
const LIKE = 57427
const REGEXP = 57428
const IN = 57429
const GE = 57430
const LE = 57431
const NE = 57432
const LG = 57433
const IE = 57434
const LNE = 57435
const LEG = 57436
const IS = 57437
const SL = 57438
const SR = 57439
const DIV = 57440
const MOD = 57441
const NOT = 57442
const OP = 57443
const COLLATE = 57444
const INTERVAL = 57445
const UMINUS = 57446
const END = 57447

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"BuiltinCharacterIdent",
	"BuiltinFucTimeAddIdent",
	"BuiltinFucTimeSubIdent",
	"BuiltinTimeUnitIdent",
	"IDENT",
	"DOUBLE_QUOTA_STRING",
	"PARAM_MARKER",
	"SINGLE_QUOTA_STRING",
	"STRING",
	"BIT_NUMBER",
	"HEX_NUMBER",
	"INTEGER",
	"FLOAT",
	"ALL",
	"ANY",
	"AS",
	"ASC",
	"AVG",
	"BY",
	"CAST",
	"CHARACTER",
	"COMMENT",
	"CONVERT",
	"COUNT",
	"DESC",
	"DISTINCT",
	"DISTINCTROW",
	"FALSE",
	"FOR",
	"FROM",
	"GROUP",
	"HAVING",
	"IF",
	"INTO",
	"LIMIT",
	"LOCK",
	"MAX",
	"MIN",
	"MODE",
	"NULL",
	"OFFSET",
	"ORDER",
	"POSITION",
	"QUARTER",
	"SELECT",
	"SET",
	"SHARE",
	"SOME",
	"SOUNDS",
	"SUM",
	"TO",
	"TRUE",
	"UPDATE",
	"WHERE",
	"EXISTS",
	"OUTER",
	"LOW",
	"UNION",
	"EXCEPT",
	"INTERSECT",
	"','",
	"JOIN",
	"STRAIGHT_JOIN",
	"NATURAL",
	"LEFT",
	"RIGHT",
	"INNER",
	"CROSS",
	"USE",
	"TABLE_REF_PRIORITY",
	"ON",
	"USING",
	"OR",
	"OO",
	"XOR",
	"AND",
	"AA",
	"BETWEEN",
	"CASE",
	"WHEN",
	"THEN",
	"ELSE",
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
	"LEG",
	"IS",
	"'|'",
	"'&'",
	"SL",
	"SR",
	"'+'",
	"'-'",
	"'*'",
	"'/'",
	"'%'",
	"DIV",
	"MOD",
	"'^'",
	"'~'",
	"'!'",
	"NOT",
	"OP",
	"COLLATE",
	"INTERVAL",
	"UMINUS",
	"'.'",
	"')'",
	"END",
	"'('",
	"'@'",
}
var yyStatenames = [...]string{}

var yyCurrentToken = ""

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
	120, 6,
	-2, 10,
	-1, 9,
	1, 7,
	120, 7,
	-2, 206,
	-1, 28,
	81, 147,
	89, 147,
	90, 147,
	91, 147,
	-2, 36,
	-1, 53,
	122, 122,
	-2, 76,
	-1, 149,
	32, 11,
	38, 11,
	39, 11,
	44, 11,
	45, 11,
	61, 11,
	62, 11,
	63, 11,
	-2, 63,
	-1, 274,
	119, 178,
	-2, 164,
	-1, 275,
	119, 179,
	-2, 165,
}

const yyPrivate = 57344

const yyLast = 1477

var yyAct = [...]int{

	239, 382, 380, 40, 4, 267, 354, 238, 332, 24,
	53, 16, 27, 266, 272, 360, 310, 316, 52, 24,
	109, 263, 241, 29, 110, 349, 117, 179, 187, 94,
	189, 130, 394, 92, 279, 113, 7, 66, 20, 8,
	150, 28, 168, 149, 167, 151, 95, 96, 66, 194,
	166, 165, 275, 104, 105, 106, 107, 108, 164, 145,
	146, 147, 148, 156, 163, 162, 104, 105, 106, 107,
	108, 66, 194, 161, 160, 275, 159, 154, 157, 131,
	132, 133, 134, 135, 136, 137, 138, 141, 139, 140,
	142, 178, 180, 143, 338, 399, 183, 377, 188, 183,
	183, 398, 376, 288, 24, 197, 198, 199, 200, 201,
	346, 8, 192, 7, 288, 375, 191, 66, 115, 374,
	116, 114, 113, 297, 99, 98, 328, 286, 111, 153,
	152, 366, 205, 206, 348, 50, 288, 347, 144, 211,
	21, 142, 104, 105, 106, 107, 108, 104, 105, 106,
	107, 108, 158, 288, 104, 105, 106, 107, 108, 404,
	174, 247, 271, 235, 236, 252, 253, 254, 255, 243,
	369, 234, 112, 217, 218, 219, 220, 221, 223, 225,
	226, 227, 228, 229, 230, 271, 345, 365, 100, 88,
	245, 344, 342, 244, 270, 249, 250, 251, 343, 216,
	242, 274, 104, 105, 106, 107, 108, 89, 66, 287,
	256, 277, 242, 385, 291, 5, 278, 104, 105, 106,
	107, 108, 397, 284, 84, 285, 276, 392, 85, 99,
	104, 105, 106, 107, 108, 386, 6, 364, 281, 106,
	107, 108, 363, 292, 196, 293, 341, 295, 107, 108,
	298, 299, 300, 294, 203, 280, 313, 282, 283, 180,
	188, 296, 259, 290, 361, 312, 135, 136, 137, 138,
	141, 139, 140, 142, 233, 325, 383, 384, 258, 176,
	113, 113, 274, 149, 202, 327, 308, 330, 309, 337,
	323, 195, 339, 101, 334, 324, 175, 336, 329, 137,
	138, 141, 139, 140, 142, 335, 65, 353, 182, 103,
	172, 182, 182, 170, 350, 350, 350, 270, 270, 270,
	103, 359, 321, 322, 274, 274, 274, 357, 358, 181,
	351, 352, 184, 185, 264, 362, 113, 331, 173, 192,
	257, 371, 91, 367, 370, 213, 368, 307, 372, 373,
	324, 90, 101, 214, 215, 212, 378, 381, 131, 132,
	133, 134, 135, 136, 137, 138, 141, 139, 140, 142,
	270, 14, 11, 13, 12, 389, 86, 274, 387, 90,
	388, 390, 391, 87, 393, 91, 395, 11, 13, 12,
	270, 355, 311, 381, 190, 400, 356, 274, 402, 97,
	396, 270, 401, 403, 19, 66, 79, 80, 274, 66,
	49, 61, 48, 47, 43, 44, 45, 46, 246, 232,
	1, 64, 68, 102, 75, 66, 194, 76, 69, 193,
	22, 23, 64, 204, 240, 104, 105, 106, 107, 108,
	186, 70, 71, 340, 51, 63, 177, 77, 333, 15,
	3, 122, 126, 124, 72, 93, 63, 123, 125, 41,
	127, 306, 128, 119, 121, 17, 379, 22, 23, 81,
	131, 132, 133, 134, 135, 136, 137, 138, 141, 139,
	140, 142, 265, 62, 66, 115, 32, 116, 114, 57,
	326, 86, 82, 59, 269, 111, 302, 90, 87, 34,
	83, 30, 268, 91, 33, 35, 36, 248, 104, 105,
	106, 107, 108, 37, 38, 25, 289, 104, 105, 106,
	107, 108, 31, 39, 67, 79, 80, 58, 66, 49,
	61, 48, 47, 43, 44, 45, 46, 210, 209, 42,
	155, 68, 26, 75, 56, 54, 76, 69, 78, 261,
	60, 64, 104, 105, 106, 107, 108, 262, 55, 273,
	70, 71, 120, 51, 9, 2, 77, 10, 304, 169,
	303, 208, 171, 72, 18, 63, 129, 118, 41, 305,
	104, 105, 106, 107, 108, 207, 314, 74, 81, 131,
	132, 133, 134, 135, 136, 137, 138, 141, 139, 140,
	142, 301, 62, 73, 231, 104, 105, 106, 107, 108,
	260, 82, 0, 104, 105, 106, 107, 108, 0, 83,
	0, 0, 0, 0, 35, 36, 65, 104, 105, 106,
	107, 108, 37, 38, 104, 105, 106, 107, 108, 0,
	0, 0, 39, 67, 79, 80, 0, 66, 49, 61,
	48, 47, 43, 44, 45, 46, 0, 0, 0, 0,
	68, 0, 75, 0, 0, 76, 69, 0, 22, 23,
	64, 318, 315, 317, 321, 322, 319, 320, 0, 70,
	71, 0, 51, 0, 0, 77, 0, 66, 115, 0,
	116, 114, 72, 0, 63, 0, 0, 41, 111, 66,
	115, 0, 116, 114, 0, 0, 0, 81, 132, 133,
	134, 135, 136, 137, 138, 141, 139, 140, 142, 0,
	0, 62, 318, 315, 317, 321, 322, 319, 320, 0,
	82, 383, 384, 0, 0, 0, 0, 0, 83, 0,
	0, 0, 0, 35, 36, 65, 0, 0, 0, 0,
	0, 37, 38, 25, 0, 0, 0, 0, 0, 0,
	0, 39, 67, 79, 80, 0, 66, 49, 61, 48,
	47, 43, 44, 45, 46, 0, 0, 0, 0, 68,
	0, 75, 0, 0, 76, 69, 0, 0, 0, 64,
	0, 0, 0, 0, 0, 0, 0, 0, 70, 71,
	0, 51, 0, 0, 77, 0, 8, 0, 0, 0,
	0, 72, 0, 63, 0, 0, 41, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 81, 133, 134, 135,
	136, 137, 138, 141, 139, 140, 142, 0, 0, 0,
	62, 0, 0, 0, 0, 0, 0, 0, 0, 82,
	0, 0, 0, 0, 0, 0, 0, 83, 0, 0,
	0, 0, 35, 36, 65, 0, 0, 0, 0, 0,
	37, 38, 25, 0, 0, 0, 0, 0, 0, 0,
	39, 67, 79, 80, 0, 66, 49, 61, 48, 47,
	43, 44, 45, 46, 0, 0, 0, 0, 68, 0,
	75, 0, 0, 76, 69, 0, 0, 0, 64, 0,
	0, 0, 0, 0, 0, 0, 0, 70, 71, 0,
	51, 0, 0, 77, 0, 0, 0, 0, 0, 0,
	72, 0, 63, 0, 0, 41, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 81, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 62,
	0, 0, 0, 0, 0, 0, 0, 0, 82, 0,
	0, 0, 0, 0, 0, 0, 83, 0, 0, 0,
	0, 35, 36, 65, 0, 0, 0, 0, 0, 37,
	38, 25, 0, 0, 0, 0, 0, 237, 0, 39,
	67, 79, 80, 0, 66, 49, 61, 48, 47, 43,
	44, 45, 46, 0, 0, 0, 0, 68, 0, 75,
	0, 0, 76, 69, 0, 0, 0, 64, 0, 0,
	0, 0, 0, 0, 0, 0, 70, 71, 0, 51,
	0, 0, 77, 0, 0, 0, 0, 0, 0, 72,
	0, 63, 0, 0, 41, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 81, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 62, 0,
	0, 0, 0, 0, 0, 0, 0, 82, 0, 0,
	0, 0, 0, 0, 0, 83, 0, 0, 0, 0,
	35, 36, 65, 0, 0, 0, 0, 0, 37, 38,
	25, 0, 0, 0, 0, 0, 0, 0, 39, 67,
	79, 80, 0, 66, 49, 61, 48, 47, 43, 44,
	45, 46, 0, 0, 0, 0, 68, 0, 75, 0,
	0, 76, 69, 0, 0, 0, 64, 0, 0, 0,
	0, 0, 0, 0, 0, 70, 71, 0, 51, 0,
	0, 77, 0, 0, 0, 0, 0, 0, 72, 0,
	63, 0, 0, 41, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 81, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 62, 0, 0,
	0, 0, 0, 0, 0, 0, 82, 0, 0, 0,
	0, 0, 0, 0, 83, 0, 0, 0, 0, 35,
	36, 65, 0, 0, 0, 0, 0, 37, 38, 0,
	0, 0, 224, 0, 0, 0, 0, 39, 67, 79,
	80, 0, 66, 49, 61, 48, 47, 43, 44, 45,
	46, 0, 0, 0, 0, 68, 0, 75, 0, 0,
	76, 69, 0, 0, 0, 64, 0, 0, 0, 0,
	0, 0, 0, 0, 70, 71, 0, 51, 0, 0,
	77, 0, 0, 0, 0, 0, 0, 72, 0, 63,
	0, 0, 41, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 81, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 62, 0, 0, 0,
	0, 0, 0, 0, 0, 82, 0, 0, 0, 0,
	0, 0, 0, 83, 0, 0, 0, 0, 35, 36,
	65, 0, 0, 0, 0, 0, 37, 38, 0, 0,
	0, 222, 0, 0, 0, 0, 39, 67, 79, 80,
	0, 66, 49, 61, 48, 47, 43, 44, 45, 46,
	0, 0, 0, 0, 68, 0, 75, 0, 0, 76,
	69, 0, 0, 0, 64, 0, 0, 0, 0, 0,
	0, 0, 0, 70, 71, 0, 51, 0, 0, 77,
	0, 0, 0, 0, 0, 0, 72, 0, 63, 0,
	0, 41, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 81, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 62, 0, 0, 0, 0,
	0, 0, 0, 0, 82, 0, 0, 0, 0, 0,
	0, 0, 83, 0, 0, 0, 0, 35, 36, 65,
	0, 0, 0, 0, 0, 37, 38, 0, 0, 0,
	0, 0, 0, 0, 0, 39, 67,
}
var yyPact = [...]int{

	-9, -1000, -1000, -1000, -1000, -1000, 326, -9, 639, -1000,
	459, 438, 438, 438, 377, 5, 4, 315, 996, 245,
	-1000, -1000, -1000, -1000, 476, 996, 365, -1000, -21, 22,
	-1000, -1000, -1000, -1000, -1000, 1353, 1353, 1353, 1353, 758,
	-1000, -86, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 11, -1000, 10, -1000, -1000, -1000, -1000,
	-45, -1000, 996, -1000, -1000, -1000, -1000, 29, -46, -48,
	-49, -57, -58, -64, -71, -72, -78, -80, -1000, -1000,
	-1000, -1000, -1000, -1000, 341, 344, 282, 69, 298, 313,
	996, 996, -9, -1000, -1000, -9, -9, 996, -1000, -1000,
	361, 417, 256, 996, 996, 996, 996, 996, 996, -1000,
	-1000, 691, -1000, -1000, -1000, -1000, -1000, -1000, 390, 1353,
	520, 25, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 264,
	110, 1353, 1353, 1353, 1353, 1234, 1115, 1353, 1353, 1353,
	1353, 1353, 1353, -1000, 415, 22, 22, 22, 22, 4,
	154, -1000, 200, 397, 877, 117, 558, -1000, 397, 438,
	401, 438, 438, 438, 996, 996, 996, 996, 1353, -1000,
	-1000, -1000, -1000, -1000, 290, -1000, -1000, 214, 558, -1000,
	558, 311, -1000, -1000, -1000, 311, 198, -1000, 529, 277,
	40, -1000, -1000, -1000, -1000, 361, -1000, 161, 161, 169,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -86, -1000, -1000,
	-1000, -1000, -88, 1353, 1353, 1353, 1353, 607, 725, 162,
	162, 193, 996, 193, 996, 30, 30, 30, 30, 30,
	-1000, -1000, -1000, -1000, -1000, 8, -1000, -1000, 89, 558,
	129, -1000, 996, -1000, 996, 996, 996, 141, 3, 996,
	996, 996, 537, 432, 551, 504, 370, 305, 996, 996,
	-1000, -1000, -1000, 358, 996, 192, 606, -1000, -1000, 679,
	679, 63, -1000, 7, -1000, -1000, 277, -1000, -1000, 758,
	258, 22, 489, 489, 441, 441, 200, -1000, 996, -27,
	-1000, 996, 359, 126, 72, 78, -1000, -1000, 71, 66,
	-10, 20, 17, 397, 397, 397, 996, -1000, -1000, -1000,
	356, 374, 558, 40, 40, 40, 205, 254, -1000, 177,
	172, -1000, -1000, -1000, -1000, 109, 67, 606, 417, 358,
	50, 1353, -1000, -1000, -1000, -1000, -1000, 558, -1000, 558,
	996, -1000, -1000, -1000, -1000, -1000, -1000, 996, 996, -1,
	-1000, -5, -18, -23, -1000, 996, 996, 606, 202, 139,
	170, -1000, 205, -1000, -1000, -1000, 40, -1000, 356, -1000,
	-1000, 558, 441, 441, -1000, -1000, -1000, -1000, 558, 163,
	-1000, 558, -1000, 996, -90, 996, 40, 157, 606, -1000,
	-19, -25, 996, 558, 996, 558, 657, 40, -1000, -1000,
	-1000, 39, -1000, -1000, -1000,
}
var yyPgo = [...]int{

	0, 33, 610, 604, 603, 587, 586, 29, 585, 8,
	577, 576, 135, 17, 15, 574, 224, 30, 16, 6,
	188, 572, 569, 567, 236, 228, 449, 565, 3, 215,
	564, 21, 24, 172, 25, 562, 559, 558, 550, 10,
	18, 548, 20, 545, 14, 544, 41, 542, 540, 539,
	527, 522, 516, 0, 504, 2, 1, 502, 501, 499,
	207, 12, 494, 27, 23, 28, 493, 5, 13, 490,
	140, 489, 486, 22, 7, 482, 466, 189, 465, 446,
	440, 404, 434, 420,
}
var yyR1 = [...]int{

	0, 83, 27, 27, 28, 28, 26, 26, 26, 26,
	24, 24, 29, 29, 29, 29, 29, 53, 53, 53,
	53, 53, 53, 53, 53, 47, 47, 47, 47, 47,
	61, 61, 61, 61, 61, 61, 61, 46, 46, 46,
	46, 46, 46, 46, 46, 46, 46, 46, 46, 46,
	46, 46, 64, 64, 64, 64, 64, 64, 64, 64,
	64, 64, 64, 64, 64, 64, 58, 58, 58, 58,
	58, 58, 58, 58, 58, 39, 40, 40, 72, 72,
	72, 51, 51, 51, 54, 54, 54, 54, 54, 54,
	45, 45, 45, 45, 45, 45, 45, 45, 71, 71,
	50, 50, 50, 66, 74, 74, 59, 49, 48, 48,
	82, 82, 73, 52, 52, 78, 78, 81, 81, 70,
	43, 37, 38, 38, 42, 42, 32, 32, 33, 33,
	33, 33, 12, 12, 8, 8, 8, 35, 35, 35,
	35, 35, 35, 35, 10, 10, 11, 11, 9, 3,
	41, 41, 41, 34, 4, 5, 15, 1, 1, 1,
	7, 7, 20, 20, 44, 44, 44, 17, 17, 75,
	75, 68, 68, 67, 67, 67, 62, 62, 36, 36,
	69, 69, 57, 57, 57, 57, 57, 57, 6, 6,
	6, 13, 13, 14, 14, 56, 56, 31, 31, 18,
	18, 76, 76, 55, 19, 19, 23, 23, 30, 80,
	80, 65, 2, 2, 2, 21, 21, 16, 16, 22,
	22, 25, 25, 25, 25, 77, 77, 79, 60, 63,
}
var yyR2 = [...]int{

	0, 1, 1, 1, 3, 3, 1, 2, 4, 4,
	1, 1, 7, 8, 4, 4, 4, 3, 3, 3,
	3, 3, 2, 1, 3, 3, 3, 3, 4, 1,
	4, 6, 6, 4, 4, 4, 1, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 5,
	5, 1, 1, 1, 1, 1, 1, 3, 2, 2,
	2, 2, 3, 1, 2, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	3, 1, 3, 5, 1, 1, 1, 1, 3, 4,
	5, 5, 5, 4, 4, 5, 5, 5, 8, 8,
	6, 6, 6, 6, 1, 3, 1, 5, 1, 0,
	1, 2, 4, 2, 0, 1, 0, 1, 3, 2,
	2, 3, 1, 1, 1, 0, 2, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 2, 1, 0, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 0,
	1, 1, 2, 0, 1, 1, 1, 2, 0, 1,
	3, 1, 1, 2, 2, 3, 1, 3, 1, 1,
	1, 3, 3, 4, 3, 5, 6, 6, 1, 2,
	2, 1, 1, 1, 0, 2, 4, 2, 0, 3,
	0, 1, 3, 1, 2, 0, 1, 0, 3, 1,
	3, 2, 1, 1, 0, 1, 0, 2, 4, 1,
	0, 2, 2, 1, 1, 2, 4, 1, 2, 1,
}
var yyChk = [...]int{

	-1000, -83, -27, -26, -28, -29, -24, 122, 48, -30,
	-23, 61, 63, 62, 45, -26, -28, -78, -15, -81,
	-7, -70, 29, 30, -53, 114, -47, -61, -46, -64,
	-58, -51, -72, -54, -59, 104, 105, 112, 113, 122,
	-28, 58, -49, 13, 14, 15, 16, 12, 11, 9,
	-12, 43, -40, -39, -43, -37, -45, -71, -50, -66,
	-38, 10, 82, 55, 31, 106, 8, 123, 21, 27,
	40, 41, 53, -4, -5, 23, 26, 46, -41, 5,
	6, 68, 91, 99, -16, -25, 32, 39, -77, -60,
	38, 44, -1, 17, -7, -1, -1, 22, 120, 120,
	-20, 37, -81, 64, 76, 77, 78, 79, 80, -42,
	-32, 19, -33, -39, 12, 9, 11, -53, -10, 98,
	-35, 99, 86, 92, 88, 93, 87, 95, 97, -11,
	52, 100, 101, 102, 103, 104, 105, 106, 107, 109,
	110, 108, 111, 114, 116, -64, -64, -64, -64, -28,
	-53, -28, 119, 119, 122, -48, -53, -39, 123, 122,
	122, 122, 122, 122, 122, 122, 122, 122, 122, -22,
	-25, -21, -16, 56, 91, -60, -77, -79, -53, -63,
	-53, -24, -29, -28, -24, -24, -80, -65, -53, -17,
	33, -44, -39, 12, 9, -20, -70, -53, -53, -53,
	-53, -53, -33, -12, 43, -61, -61, -8, 51, 18,
	17, 114, 91, 81, 89, 90, 89, -46, -46, -46,
	-46, -46, 117, -46, 117, -46, -46, -46, -46, -46,
	-46, -3, 4, 120, -40, -39, -39, 120, -74, -53,
	-82, -73, 83, -39, -1, -7, 17, -53, 106, -1,
	-1, -1, -53, -53, -53, -53, -46, 50, 64, 64,
	-2, 20, 28, -31, 57, -75, -68, -67, -57, -62,
	-28, 122, -44, -36, -39, 12, -17, -28, -28, 122,
	-46, -64, -46, -46, -53, -53, 119, 120, 64, -52,
	-73, 85, -53, -53, -74, -53, 120, 120, -53, -53,
	-53, 64, 64, 19, 64, 75, 91, 42, -63, -65,
	-18, 34, -53, 64, -6, 66, -13, 67, 65, 70,
	71, 68, 69, -42, -32, -28, -69, -68, 119, -31,
	-74, 79, -9, 7, -9, -40, -39, -53, 121, -53,
	84, 120, 120, 120, 120, 120, 120, 117, 117, -34,
	-39, -34, -34, -53, -19, 35, 22, -68, -68, -67,
	-14, 59, -13, 65, 65, 120, 64, -44, -18, 120,
	-61, -53, -53, -53, 120, 120, 120, 120, -53, -76,
	-55, -53, -56, 74, 75, 74, 65, -14, -68, -19,
	-9, -9, 64, -53, 122, -53, -68, 65, 120, 120,
	-55, -74, -56, -67, 120,
}
var yyDef = [...]int{

	0, -2, 1, 2, -2, -2, 207, 0, 116, -2,
	0, 159, 159, 159, 0, 0, 11, 163, 0, 115,
	156, 117, 160, 161, 125, 0, 23, 29, -2, 51,
	52, 53, 54, 55, 56, 0, 0, 0, 0, 0,
	63, 0, 65, 66, 67, 68, 69, 70, 71, 72,
	73, 74, 81, -2, 78, 79, 84, 85, 86, 87,
	0, 106, 109, 132, 133, 77, 75, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 123, 154,
	155, 150, 151, 152, 220, 216, 0, 0, 223, 224,
	0, 0, 0, 157, 158, 0, 0, 0, 4, 5,
	168, 0, 163, 0, 0, 0, 0, 0, 0, 119,
	124, 0, 127, 128, 129, 130, 131, 22, 0, 0,
	0, 144, 137, 138, 139, 140, 141, 142, 143, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 146, 0, 58, 59, 60, 61, -2,
	0, 64, 0, 0, 0, 0, 108, 120, 0, 159,
	0, 159, 159, 159, 0, 0, 0, 0, 0, 8,
	219, 9, 215, 217, 0, 221, 222, 225, 227, 228,
	229, 14, 10, 11, 15, 16, 208, 209, 214, 198,
	0, 162, 164, 165, 166, 168, 118, 17, 18, 19,
	20, 21, 126, 24, 25, 26, 27, 0, 134, 135,
	136, 145, 0, 0, 0, 0, 0, 37, 38, 39,
	40, 41, 0, 42, 0, 43, 44, 45, 46, 47,
	48, 57, 149, 62, 82, 76, 80, 88, 0, 104,
	114, 110, 0, 121, 0, 0, 0, 0, 77, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	211, 212, 213, 200, 0, 167, 169, 171, 172, 125,
	0, 0, 176, 0, -2, -2, 198, 28, 30, 0,
	0, 34, 35, 33, 0, 0, 0, 89, 0, 0,
	111, 0, 0, 0, 0, 0, 93, 94, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 218, 226, 210,
	205, 0, 197, 0, 0, 0, 194, 0, 188, 0,
	0, 191, 192, 173, 174, 11, 0, 180, 0, 200,
	0, 0, 49, 148, 50, 83, 76, 105, 107, 113,
	0, 90, 91, 92, 95, 96, 97, 0, 0, 0,
	153, 0, 0, 0, 12, 0, 0, 170, 182, 184,
	0, 193, 194, 189, 190, 175, 0, 177, 205, 31,
	32, 112, 0, 0, 100, 101, 102, 103, 204, 199,
	201, 203, 183, 0, 0, 0, 0, 0, 181, 13,
	0, 0, 0, 195, 0, 185, 0, 0, 98, 99,
	202, 0, 186, 187, 196,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 113, 3, 3, 3, 108, 101, 3,
	122, 120, 106, 104, 64, 105, 119, 107, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	87, 86, 88, 3, 123, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 111, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 100, 3, 112,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
	52, 53, 54, 55, 56, 57, 58, 59, 60, 61,
	62, 63, 65, 66, 67, 68, 69, 70, 71, 72,
	73, 74, 75, 76, 77, 78, 79, 80, 81, 82,
	83, 84, 85, 89, 90, 91, 92, 93, 94, 95,
	96, 97, 98, 99, 102, 103, 109, 110, 114, 115,
	116, 117, 118, 121,
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
	yyErrorVerbose = true
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

func yyGetToken(character int, token int, lval *yySymType) string {
	switch character {
	case IDENT, BuiltinCharacterIdent, BuiltinFucTimeAddIdent, BuiltinFucTimeSubIdent, BuiltinTimeUnitIdent:
		return lval.ident
	case PARAM_MARKER, DOUBLE_QUOTA_STRING, SINGLE_QUOTA_STRING:
		return lval.str
	case BIT_NUMBER, HEX_NUMBER, INTEGER, FLOAT:
		return lval.val
	}

	return yyTokname(token)
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

	res := "语法分析出现错误:unexpected " + yyCurrentToken //yyTokname(lookAhead)

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
	yyCurrentToken = yyGetToken(char, token, lval)
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
		//line parser.y:214
		{
			yylex.(*Tokener).ParseStmt = yyDollar[1].node
		}
	case 4:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:233
		{
			yyVAL.node = yyDollar[2].node
		}
	case 5:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:237
		{
			yyVAL.node = yyDollar[2].node
		}
	case 6:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:243
		{
			yyVAL.node = yyDollar[1].node
		}
	case 7:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:247
		{
			switch node := yyDollar[1].node.(type) {
			case *ast.SimpleSelectStmt:
				node.Sort = yyDollar[2].node.(*ast.SortClause)
			case *ast.UnionStmt:
				node.Sort = yyDollar[2].node.(*ast.SortClause)
			}

			yyVAL.node = yyDollar[1].node
		}
	case 8:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:258
		{
			switch node := yyDollar[1].node.(type) {
			case *ast.SimpleSelectStmt:
				node.Sort = yyDollar[2].node.(*ast.SortClause)
				node.Lock = yyDollar[3].node.(*ast.LockClause)
				node.Limit = yyDollar[4].node.(*ast.LimitClause)
			case *ast.UnionStmt:
				node.Sort = yyDollar[2].node.(*ast.SortClause)
				node.Lock = yyDollar[3].node.(*ast.LockClause)
				node.Limit = yyDollar[4].node.(*ast.LimitClause)
			}

			yyVAL.node = yyDollar[1].node
		}
	case 9:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:273
		{
			switch node := yyDollar[1].node.(type) {
			case *ast.SimpleSelectStmt:
				node.Sort = yyDollar[2].node.(*ast.SortClause)
				node.Lock = yyDollar[4].node.(*ast.LockClause)
				node.Limit = yyDollar[3].node.(*ast.LimitClause)
			case *ast.UnionStmt:
				node.Sort = yyDollar[2].node.(*ast.SortClause)
				node.Lock = yyDollar[4].node.(*ast.LockClause)
				node.Limit = yyDollar[3].node.(*ast.LimitClause)
			}

			yyVAL.node = yyDollar[1].node

			yyVAL.node = yyDollar[1].node
		}
	case 10:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:291
		{
			yyVAL.node = yyDollar[1].node
		}
	case 11:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:292
		{
			yyVAL.node = yyDollar[1].node
		}
	case 12:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.y:297
		{
			yyVAL.node = &ast.SimpleSelectStmt{
				Target: &ast.TargetClause{Target_ref: yyDollar[2].list},
				Into:   yyDollar[3].node.(*ast.IntoClause),
				From:   yyDollar[4].node.(*ast.FromClause),
				Where:  yyDollar[5].node.(*ast.WhereClause),
				Group:  yyDollar[6].node.(*ast.GroupClause),
				Having: yyDollar[7].node.(*ast.HavingClause),
			}
		}
	case 13:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line parser.y:308
		{
			yyVAL.node = &ast.SimpleSelectStmt{
				Distinct: yyDollar[2].node.(*ast.DistinctClause),
				Target:   &ast.TargetClause{Target_ref: yyDollar[3].list},
				Into:     yyDollar[4].node.(*ast.IntoClause),
				From:     yyDollar[5].node.(*ast.FromClause),
				Where:    yyDollar[6].node.(*ast.WhereClause),
				Group:    yyDollar[7].node.(*ast.GroupClause),
				Having:   yyDollar[8].node.(*ast.HavingClause),
			}
		}
	case 14:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:320
		{
			stmt := &ast.UnionStmt{
				DistinctType: yyDollar[3].tag,
				Left:         yyDollar[1].node,
				Right:        yyDollar[4].node,
			}
			stmt.SetTag(ast.Union)
			yyVAL.node = stmt
		}
	case 15:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:330
		{
			stmt := &ast.UnionStmt{
				DistinctType: yyDollar[3].tag,
				Left:         yyDollar[1].node,
				Right:        yyDollar[4].node,
			}
			stmt.SetTag(ast.Except)
			yyVAL.node = stmt
		}
	case 16:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:340
		{
			stmt := &ast.UnionStmt{
				DistinctType: yyDollar[3].tag,
				Left:         yyDollar[1].node,
				Right:        yyDollar[4].node,
			}
			stmt.SetTag(ast.Intersect)
			yyVAL.node = stmt
		}
	case 17:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:360
		{
			expr := &ast.And0OrExpr{Operator: "OR", Expr: []ast.ExprNode{}}
			expr.SetTag(ast.Expr_Or)
			expr.Combine(yyDollar[1].expr, yyDollar[3].expr)
			yyVAL.expr = expr

		}
	case 18:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:368
		{
			expr := &ast.And0OrExpr{Operator: "OR", Expr: []ast.ExprNode{}}
			expr.SetTag(ast.Expr_Or)
			expr.Combine(yyDollar[1].expr, yyDollar[3].expr)
			yyVAL.expr = expr
		}
	case 19:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:374
		{
			yyVAL.expr = &ast.Expr{Operator: "XOR", Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 20:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:376
		{
			expr := &ast.And0OrExpr{Operator: "AND", Expr: []ast.ExprNode{}}
			expr.SetTag(ast.Expr_And)
			expr.Combine(yyDollar[1].expr, yyDollar[3].expr)
			yyVAL.expr = expr
		}
	case 21:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:383
		{
			expr := &ast.And0OrExpr{Operator: "AND", Expr: []ast.ExprNode{}}
			expr.SetTag(ast.Expr_And)
			expr.Combine(yyDollar[1].expr, yyDollar[3].expr)
			yyVAL.expr = expr
		}
	case 22:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:389
		{
			yyVAL.expr = &ast.NotExpr{Expr: yyDollar[2].expr}
		}
	case 23:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:390
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 24:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:392
		{
			yyVAL.expr = &ast.IsTrueExpr{HasNot: yyDollar[2].boolean, IsTrue: yyDollar[3].boolean, Expr: yyDollar[1].expr}
		}
	case 25:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:398
		{
			yyVAL.expr = &ast.IsNullExpr{HasNot: yyDollar[2].boolean, Left: yyDollar[1].expr}
		}
	case 26:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:402
		{
			yyVAL.expr = &ast.Expr{Operator: "=", Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 27:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:406
		{
			yyVAL.expr = &ast.Expr{Operator: yyDollar[2].val, Left: yyDollar[1].expr, Right: yyDollar[3].expr}

		}
	case 28:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:411
		{
			expr := &ast.SubqueryExpr{Operator: yyDollar[2].val, Subtype: yyDollar[3].tag, Left: yyDollar[1].expr, Select: yyDollar[4].node}
			expr.SetTag(ast.Subquery_Operator)
			yyVAL.expr = expr
		}
	case 29:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:417
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 30:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:423
		{
			expr := &ast.SubqueryExpr{HasNot: yyDollar[2].boolean, Left: yyDollar[1].expr, Select: yyDollar[4].node}
			expr.SetTag(ast.Subquery_In)
			yyVAL.expr = expr
		}
	case 31:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.y:429
		{
			yyVAL.expr = &ast.InExpr{HasNot: yyDollar[2].boolean, Left: yyDollar[1].expr, Right: yyDollar[5].list}
		}
	case 32:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.y:433
		{
			yyVAL.expr = &ast.BetweenExpr{HasNot: yyDollar[2].boolean, Expr: yyDollar[1].expr, Left: yyDollar[4].expr, Right: yyDollar[6].expr}
		}
	case 33:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:437
		{
			yyVAL.expr = &ast.LikeExpr{HasNot: false, Left: yyDollar[1].expr, Right: yyDollar[4].expr}
		}
	case 34:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:441
		{
			yyVAL.expr = &ast.LikeExpr{HasNot: yyDollar[2].boolean, Left: yyDollar[1].expr, Right: yyDollar[4].expr}
		}
	case 35:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:445
		{
			yyVAL.expr = &ast.RegexpExpr{HasNot: yyDollar[2].boolean, Left: yyDollar[1].expr, Right: yyDollar[4].expr}
		}
	case 36:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:449
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 37:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:455
		{
			yyVAL.expr = &ast.Expr{Operator: "|", Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 38:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:456
		{
			yyVAL.expr = &ast.Expr{Operator: "&", Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 39:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:457
		{
			yyVAL.expr = &ast.Expr{Operator: "<<", Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 40:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:458
		{
			yyVAL.expr = &ast.Expr{Operator: ">>", Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 41:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:459
		{
			yyVAL.expr = &ast.Expr{Operator: "+", Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 42:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:460
		{
			yyVAL.expr = &ast.Expr{Operator: "-", Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 43:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:461
		{
			yyVAL.expr = &ast.Expr{Operator: "*", Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 44:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:462
		{
			yyVAL.expr = &ast.Expr{Operator: "/", Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 45:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:463
		{
			yyVAL.expr = &ast.Expr{Operator: "DIV", Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 46:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:464
		{
			yyVAL.expr = &ast.Expr{Operator: "MOD", Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 47:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:465
		{
			yyVAL.expr = &ast.Expr{Operator: "%", Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 48:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:466
		{
			yyVAL.expr = &ast.Expr{Operator: "^", Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 49:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:468
		{
			yyVAL.expr = &ast.IntervalExpr{Operator: ast.Op_Add, Left: yyDollar[1].expr, Right: yyDollar[4].expr, TimeUnit: yyDollar[5].tag}
		}
	case 50:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:472
		{
			yyVAL.expr = &ast.IntervalExpr{Operator: ast.Op_Minus, Left: yyDollar[1].expr, Right: yyDollar[4].expr, TimeUnit: yyDollar[5].tag}
		}
	case 51:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:476
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 52:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:481
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 53:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:482
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 54:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:483
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 55:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:484
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 56:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:485
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 57:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:487
		{
			yyVAL.expr = &ast.CollateExpr{Expr: yyDollar[1].expr, Collate: yyDollar[3].tag}
		}
	case 58:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:491
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "+", Expr: yyDollar[2].expr}
		}
	case 59:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:495
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "-", Expr: yyDollar[2].expr}
		}
	case 60:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:499
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "~", Expr: yyDollar[2].expr}
		}
	case 61:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:503
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "!", Expr: yyDollar[2].expr}
		}
	case 62:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:507
		{
			switch node := yyDollar[2].expr.(type) {
			case *ast.Expr:
				node.IsEnclosed = true
			case *ast.And0OrExpr:
				node.IsEnclosed = true
			}
			yyVAL.expr = yyDollar[2].expr
		}
	case 63:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:518
		{
			yyVAL.expr = yyDollar[1].node.(ast.ExprNode)
		}
	case 64:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:522
		{
			expr := &ast.SubqueryExpr{Select: yyDollar[2].node}
			expr.SetTag(ast.Subquery_Exists)
			yyVAL.expr = expr
		}
	case 65:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:528
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 66:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:534
		{
			Expr := &ast.Integer{Operator: []string{}, Val: yyDollar[1].val}
			Expr.SetTag(ast.Bit)
			yyVAL.expr = Expr
		}
	case 67:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:540
		{
			Expr := &ast.Integer{Operator: []string{}, Val: yyDollar[1].val}
			Expr.SetTag(ast.Hex)
			yyVAL.expr = Expr
		}
	case 68:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:546
		{
			yyVAL.expr = &ast.Integer{Operator: []string{}, Val: yyDollar[1].val}
		}
	case 69:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:550
		{
			val := strings.Split(yyDollar[1].val, ".")
			yyVAL.expr = &ast.Float{
				Operator: []string{},
				Integer:  val[0],
				Decimal:  val[1],
			}
		}
	case 70:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:558
		{
			yyVAL.expr = &ast.Marker{Str: yyDollar[1].str}
		}
	case 71:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:560
		{
			str := &ast.String{Operator: []string{}, Str: yyDollar[1].str}
			str.SetTag(ast.String_single)
			yyVAL.expr = str
		}
	case 72:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:566
		{
			str := &ast.String{Operator: []string{}, Str: yyDollar[1].str}
			str.SetTag(ast.String_double)
			yyVAL.expr = str
		}
	case 73:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:571
		{
			yyVAL.expr = &ast.True{IsTrue: yyDollar[1].boolean}
		}
	case 74:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:572
		{
			yyVAL.expr = &ast.Null{}
		}
	case 75:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:576
		{
			yyVAL.val = yyDollar[1].ident
		}
	case 76:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:580
		{
			yyVAL.val = yyDollar[1].val
		}
	case 77:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:584
		{
			yyVAL.val = "*"
		}
	case 78:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:589
		{
			yyVAL.expr = &ast.UserVariable{Variable: yyDollar[1].val}
		}
	case 79:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:590
		{
			yyVAL.expr = &ast.SystemVariable{Schema: yyDollar[1].val}
		}
	case 80:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:592
		{
			yyVAL.expr = &ast.SystemVariable{Schema: yyDollar[1].val, Parameter: yyDollar[3].val}
		}
	case 81:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:598
		{
			if yyDollar[1].val == "*" {
				yyVAL.expr = &ast.Star{Column: []*ast.Column{}}
			} else {
				yyVAL.expr = &ast.Column{Column: yyDollar[1].val}
			}
		}
	case 82:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:606
		{
			if yyDollar[3].val == "*" {
				yyVAL.expr = &ast.Star{Table: yyDollar[1].val, Column: []*ast.Column{}}
			} else {
				yyVAL.expr = &ast.Column{Table: yyDollar[1].val, Column: yyDollar[3].val}
			}
		}
	case 83:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:614
		{
			if yyDollar[5].val == "*" {
				yyVAL.expr = &ast.Star{DB: yyDollar[1].val, Table: yyDollar[3].val, Column: []*ast.Column{}}
			} else {
				yyVAL.expr = &ast.Column{DB: yyDollar[1].val, Table: yyDollar[3].val, Column: yyDollar[5].val}
			}
		}
	case 84:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:625
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 85:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:629
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 86:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:633
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 87:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:637
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 88:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:641
		{
			yyVAL.expr = &ast.FuncCall{Name: yyDollar[1].val}
		}
	case 89:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:645
		{
			yyVAL.expr = &ast.FuncCall{Name: yyDollar[1].val, Arg: yyDollar[3].list}
		}
	case 90:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:651
		{
			expr := &ast.AggregationFuncCall{DistinctType: yyDollar[3].tag, Arg: []ast.ExprNode{yyDollar[4].expr}}
			expr.SetTag(ast.Aggregation_Avg)
			yyVAL.expr = expr
		}
	case 91:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:657
		{
			expr := &ast.AggregationFuncCall{DistinctType: yyDollar[3].tag, Arg: yyDollar[4].list}
			expr.SetTag(ast.Aggregation_Count)
			yyVAL.expr = expr
		}
	case 92:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:663
		{
			expr := &ast.AggregationFuncCall{DistinctType: ast.AST_ALL, Arg: []ast.ExprNode{yyDollar[4].expr}}
			expr.SetTag(ast.Aggregation_Count)
			yyVAL.expr = expr
		}
	case 93:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:669
		{
			expr := &ast.AggregationFuncCall{DistinctType: ast.AST_EMPTY, Arg: []ast.ExprNode{yyDollar[3].expr}}
			expr.SetTag(ast.Aggregation_Count)
			yyVAL.expr = expr
		}
	case 94:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:675
		{
			expr := &ast.AggregationFuncCall{DistinctType: ast.AST_EMPTY, Arg: []ast.ExprNode{&ast.Star{Column: []*ast.Column{}}}}
			expr.SetTag(ast.Aggregation_Count)
			yyVAL.expr = expr
		}
	case 95:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:681
		{
			expr := &ast.AggregationFuncCall{DistinctType: yyDollar[3].tag, Arg: []ast.ExprNode{yyDollar[4].expr}}
			expr.SetTag(ast.Aggregation_Max)
			yyVAL.expr = expr
		}
	case 96:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:687
		{
			expr := &ast.AggregationFuncCall{DistinctType: yyDollar[3].tag, Arg: []ast.ExprNode{yyDollar[4].expr}}
			expr.SetTag(ast.Aggregation_Min)
			yyVAL.expr = expr
		}
	case 97:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:693
		{
			expr := &ast.AggregationFuncCall{DistinctType: yyDollar[3].tag, Arg: []ast.ExprNode{yyDollar[4].expr}}
			expr.SetTag(ast.Aggregation_Sum)
			yyVAL.expr = expr
		}
	case 98:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line parser.y:701
		{
			Expr := &ast.CalTimeFuncCall{Interval: &ast.IntervalExpr{Operator: ast.Op_Add, Left: yyDollar[3].expr, Right: yyDollar[6].expr, TimeUnit: yyDollar[7].tag}}
			Expr.SetTag(yyDollar[1].tag)
			yyVAL.expr = Expr
		}
	case 99:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line parser.y:707
		{
			Expr := &ast.CalTimeFuncCall{Interval: &ast.IntervalExpr{Operator: ast.Op_Minus, Left: yyDollar[3].expr, Right: yyDollar[6].expr, TimeUnit: yyDollar[7].tag}}
			Expr.SetTag(yyDollar[1].tag)
			yyVAL.expr = Expr
		}
	case 100:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.y:715
		{
			yyVAL.expr = &ast.CastFuncCall{Expr: yyDollar[3].expr, CastType: yyDollar[5].val}
		}
	case 101:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.y:719
		{
			yyVAL.expr = &ast.CastFuncCall{Expr: yyDollar[3].expr, CastType: yyDollar[5].val}
		}
	case 102:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.y:723
		{
			yyVAL.expr = &ast.CastFuncCall{Expr: yyDollar[3].expr, CastType: yyDollar[5].val}
		}
	case 103:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.y:729
		{
			yyVAL.expr = &ast.StringFuncCall{Substr: yyDollar[3].expr, Str: yyDollar[5].expr}
		}
	case 104:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:734
		{
			yyVAL.list = []ast.ExprNode{yyDollar[1].expr}
		}
	case 105:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:735
		{
			yyVAL.list = append(yyDollar[1].list, yyDollar[3].expr)
		}
	case 106:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:738
		{
			yyVAL.expr = &ast.Marker{Str: yyDollar[1].str}
		}
	case 107:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:742
		{
			yyVAL.expr = &ast.CaseExpr{Case: yyDollar[2].expr, When: yyDollar[3].list, Else: yyDollar[4].expr}
		}
	case 108:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:747
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 109:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:748
		{
			yyVAL.expr = nil
		}
	case 110:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:751
		{
			yyVAL.list = []ast.ExprNode{yyDollar[1].expr}
		}
	case 111:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:752
		{
			yyVAL.list = append(yyDollar[1].list, yyDollar[2].expr)
		}
	case 112:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:755
		{
			yyVAL.expr = &ast.Expr{Left: yyDollar[2].expr, Right: yyDollar[4].expr}
		}
	case 113:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:758
		{
			yyVAL.expr = &ast.Expr{Left: yyDollar[2].expr}
		}
	case 114:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:759
		{
			yyVAL.expr = nil
		}
	case 115:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:762
		{
			yyVAL.list = yyDollar[1].list
		}
	case 116:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:763
		{
			yyVAL.list = nil
		}
	case 117:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:766
		{
			yyVAL.list = []ast.ExprNode{yyDollar[1].expr}
		}
	case 118:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:767
		{
			yyVAL.list = append(yyDollar[1].list, yyDollar[3].expr)
		}
	case 119:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:771
		{
			tb := &ast.Tuple{Ref: yyDollar[1].expr, Alias: yyDollar[2].val}
			switch yyDollar[1].expr.(type) {
			case *ast.AggregationFuncCall:
				tb.SetTag(ast.Tuple_agg)
			case *ast.FuncCall, *ast.CastFuncCall, *ast.CalTimeFuncCall, *ast.StringFuncCall:
				tb.SetTag(ast.Tuple_funcall)
			case *ast.Column:
				tb.SetTag(ast.Tuple_column)
			case *ast.Star:
				tb.SetTag(ast.Tuple_star)
			case *ast.SimpleSelectStmt:
				tb.SetTag(ast.Tuple_SimpleSubquery)
			case *ast.UnionStmt:
				tb.SetTag(ast.Tuple_UnionSubquery)
			default:
				tb.SetTag(ast.Tuple_expr)
			}
			yyVAL.expr = tb
		}
	case 120:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:793
		{
			yyVAL.val = yyDollar[2].val
		}
	case 121:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:796
		{
			yyVAL.val = yyDollar[3].val
		}
	case 122:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:799
		{
			yyVAL.val = yyDollar[1].val
		}
	case 123:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:800
		{
			yyVAL.val = yyDollar[1].val
		}
	case 124:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:803
		{
			yyVAL.val = yyDollar[1].val
		}
	case 125:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:804
		{
			yyVAL.val = ""
		}
	case 126:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:807
		{
			yyVAL.val = yyDollar[2].val
		}
	case 127:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:808
		{
			yyVAL.val = yyDollar[1].val
		}
	case 128:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:811
		{
			yyVAL.val = yyDollar[1].val
		}
	case 129:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:812
		{
			yyVAL.val = yyDollar[1].str
		}
	case 130:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:813
		{
			yyVAL.val = yyDollar[1].str
		}
	case 131:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:814
		{
			yyVAL.val = yyDollar[1].str
		}
	case 132:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:818
		{
			yyVAL.boolean = true
		}
	case 133:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:819
		{
			yyVAL.boolean = false
		}
	case 134:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:824
		{
			yyVAL.tag = ast.Operator_Some
		}
	case 135:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:825
		{
			yyVAL.tag = ast.Operator_All
		}
	case 136:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:826
		{
			yyVAL.tag = ast.Operator_Any
		}
	case 137:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:829
		{
			yyVAL.val = "="
		}
	case 138:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:830
		{
			yyVAL.val = ">="
		}
	case 139:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:831
		{
			yyVAL.val = ">"
		}
	case 140:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:832
		{
			yyVAL.val = "<="
		}
	case 141:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:833
		{
			yyVAL.val = "<"
		}
	case 142:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:834
		{
			yyVAL.val = "<>"
		}
	case 143:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:835
		{
			yyVAL.val = "!="
		}
	case 144:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:838
		{
			yyVAL.boolean = false
		}
	case 145:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:839
		{
			yyVAL.boolean = true
		}
	case 146:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:843
		{
			yyVAL.boolean = true
		}
	case 147:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:844
		{
			yyVAL.boolean = false
		}
	case 148:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:847
		{
			yyVAL.tag = ast.TimeUnit[yyDollar[1].ident]
		}
	case 149:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:851
		{
			yyVAL.tag = ast.Character[yyDollar[1].ident]
		}
	case 150:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:855
		{
			yyVAL.val = "LEFT"
		}
	case 151:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:856
		{
			yyVAL.val = "IN"
		}
	case 152:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:857
		{
			yyVAL.val = "IS"
		}
	case 153:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:860
		{
			yyVAL.val = yyDollar[1].val
		}
	case 154:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:863
		{
			yyVAL.tag = ast.FuncTime[yyDollar[1].ident]
		}
	case 155:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:866
		{
			yyVAL.tag = ast.FuncTime[yyDollar[1].ident]
		}
	case 156:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:877
		{
			clause := &ast.DistinctClause{}
			clause.SetTag(yyDollar[1].tag)
			yyVAL.node = clause
		}
	case 157:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:884
		{
			yyVAL.tag = ast.AST_ALL
		}
	case 158:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:885
		{
			yyVAL.tag = yyDollar[1].tag
		}
	case 159:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:886
		{
			yyVAL.tag = ast.AST_EMPTY
		}
	case 160:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:889
		{
			yyVAL.tag = ast.AST_DISTINCT
		}
	case 161:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:890
		{
			yyVAL.tag = ast.AST_DISTINCTROW
		}
	case 162:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:900
		{
			yyVAL.node = &ast.IntoClause{Relation: &ast.Relation{Table: yyDollar[2].val}}
		}
	case 163:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:902
		{
			clause := &ast.IntoClause{}
			clause.SetTag(ast.AST_EMPTY)
			yyVAL.node = clause
		}
	case 164:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:909
		{
			yyVAL.val = yyDollar[1].val
		}
	case 165:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:910
		{
			yyVAL.val = yyDollar[1].str
		}
	case 166:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:911
		{
			yyVAL.val = yyDollar[1].str
		}
	case 167:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:921
		{
			yyVAL.node = &ast.FromClause{Table_ref: yyDollar[2].list}
		}
	case 168:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:923
		{
			clause := &ast.FromClause{}
			clause.SetTag(ast.AST_EMPTY)
			yyVAL.node = clause
		}
	case 169:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:931
		{
			yyVAL.list = []ast.ExprNode{yyDollar[1].expr}
		}
	case 170:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:932
		{
			yyVAL.list = append(yyDollar[1].list, yyDollar[3].expr)
		}
	case 171:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:935
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 172:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:936
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 173:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:940
		{
			yyDollar[1].expr.(*ast.Relation).Alias = yyDollar[2].val
			yyVAL.expr = yyDollar[1].expr
		}
	case 174:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:945
		{
			tb := &ast.Relation{Select: yyDollar[1].node, Alias: yyDollar[2].val}
			tb.SetTag(ast.Relation_Subquery)
			yyVAL.expr = tb
		}
	case 175:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:951
		{
			switch node := yyDollar[2].expr.(type) {
			case *ast.Join:
				node.IsEnclosed = true
			}
			yyVAL.expr = yyDollar[2].expr
		}
	case 176:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:961
		{
			tb := &ast.Relation{Table: yyDollar[1].val}
			tb.SetTag(ast.Relation_Table)
			yyVAL.expr = tb
		}
	case 177:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:967
		{
			tb := &ast.Relation{DB: yyDollar[1].val, Table: yyDollar[3].val}
			tb.SetTag(ast.Relation_Table)
			yyVAL.expr = tb
		}
	case 178:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:974
		{
			yyVAL.val = yyDollar[1].val
		}
	case 179:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:975
		{
			yyVAL.val = yyDollar[1].str
		}
	case 180:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:978
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 181:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:980
		{
			rel := &ast.Join{Left: yyDollar[1].expr, Right: yyDollar[3].expr}
			rel.SetTag(ast.Cross_join)
			yyVAL.expr = rel
		}
	case 182:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:991
		{
			rel := &ast.Join{Left: yyDollar[1].expr, Right: yyDollar[3].expr}
			rel.SetTag(ast.Cross_join)
			yyVAL.expr = rel
		}
	case 183:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:997
		{
			rel := &ast.Join{Left: yyDollar[1].expr, Right: yyDollar[3].expr, Condition: yyDollar[4].expr}
			rel.SetTag(ast.Inner_join)
			yyVAL.expr = rel
		}
	case 184:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:1003
		{
			rel := &ast.Join{Left: yyDollar[1].expr, Right: yyDollar[3].expr}
			rel.SetTag(ast.Straight_join)
			yyVAL.expr = rel
		}
	case 185:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:1009
		{
			rel := &ast.Join{Left: yyDollar[1].expr, Right: yyDollar[3].expr, Condition: yyDollar[5].expr}
			rel.SetTag(ast.Straight_join)
			yyVAL.expr = rel
		}
	case 186:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.y:1015
		{
			rel := &ast.Join{HasOuter: yyDollar[3].boolean, Left: yyDollar[1].expr, Right: yyDollar[5].expr, Condition: yyDollar[6].expr}
			if yyDollar[2].boolean {
				rel.SetTag(ast.Left_join)
			} else {
				rel.SetTag(ast.Right_join)
			}
			yyVAL.expr = rel
		}
	case 187:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.y:1025
		{
			rel := &ast.Join{HasNatural: true, HasOuter: yyDollar[4].boolean, Left: yyDollar[1].expr, Right: yyDollar[6].expr}
			if yyDollar[3].boolean {
				rel.SetTag(ast.Left_join)
			} else {
				rel.SetTag(ast.Right_join)
			}
			yyVAL.expr = rel
		}
	case 188:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1036
		{
			yyVAL.tag = 1
		}
	case 189:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:1037
		{
			yyVAL.tag = 2
		}
	case 190:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:1038
		{
			yyVAL.tag = 3
		}
	case 191:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1042
		{
			yyVAL.boolean = true
		}
	case 192:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1043
		{
			yyVAL.boolean = false
		}
	case 193:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1046
		{
			yyVAL.boolean = true
		}
	case 194:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:1047
		{
			yyVAL.boolean = false
		}
	case 195:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:1050
		{
			yyVAL.expr = yyDollar[2].expr
		}
	case 196:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:1051
		{
			yyVAL.expr = &ast.Using{Column: yyDollar[3].list}
		}
	case 197:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:1062
		{
			clause := &ast.WhereClause{}
			switch yyDollar[2].expr.(type) {
			case *ast.And0OrExpr:
				clause.Where = yyDollar[2].expr
			default:
				clause.Where = &ast.And0OrExpr{Operator: "And", Expr: []ast.ExprNode{yyDollar[2].expr}}
				clause.Where.SetTag(ast.Expr_And)
			}
			yyVAL.node = clause

		}
	case 198:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:1075
		{
			clause := &ast.WhereClause{}
			clause.SetTag(ast.AST_EMPTY)
			yyVAL.node = clause
		}
	case 199:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:1088
		{
			yyVAL.node = &ast.GroupClause{Target_ref: yyDollar[3].list}
		}
	case 200:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:1090
		{
			clause := &ast.GroupClause{}
			clause.SetTag(ast.AST_EMPTY)
			yyVAL.node = clause
		}
	case 201:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1098
		{
			yyVAL.list = []ast.ExprNode{yyDollar[1].expr}
		}
	case 202:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:1102
		{
			yyVAL.list = append(yyDollar[1].list, yyDollar[3].expr)
		}
	case 203:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1108
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 204:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:1117
		{
			yyVAL.node = &ast.HavingClause{Having: yyDollar[2].expr}
		}
	case 205:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:1119
		{
			clause := &ast.HavingClause{}
			clause.SetTag(ast.AST_EMPTY)
			yyVAL.node = clause
		}
	case 206:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1133
		{
			yyVAL.node = yyDollar[1].node
		}
	case 207:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:1135
		{
			clause := &ast.SortClause{}
			clause.SetTag(ast.AST_EMPTY)
			yyVAL.node = clause
		}
	case 208:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:1142
		{
			yyVAL.node = &ast.SortClause{Target_ref: yyDollar[3].list}
		}
	case 209:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1145
		{
			yyVAL.list = []ast.ExprNode{yyDollar[1].expr}
		}
	case 210:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:1146
		{
			yyVAL.list = append(yyDollar[1].list, yyDollar[3].expr)
		}
	case 211:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:1149
		{
			yyVAL.expr = &ast.Sortby{Expr: yyDollar[1].expr, AscType: yyDollar[2].tag}
		}
	case 212:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1152
		{
			yyVAL.tag = ast.Sort_Asc
		}
	case 213:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1153
		{
			yyVAL.tag = ast.Sort_Desc
		}
	case 214:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:1154
		{
			yyVAL.tag = ast.Sort_Empty
		}
	case 215:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1163
		{
			yyVAL.node = yyDollar[1].node
		}
	case 216:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:1165
		{
			clause := &ast.LockClause{}
			clause.SetTag(ast.AST_EMPTY)
			yyVAL.node = clause
		}
	case 217:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:1173
		{
			lock := &ast.LockClause{}
			lock.SetTag(ast.Lock_update)
			yyVAL.node = lock
		}
	case 218:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:1179
		{
			lock := &ast.LockClause{}
			lock.SetTag(ast.Lock_share)
			yyVAL.node = lock
		}
	case 219:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1193
		{
			yyVAL.node = yyDollar[1].node
		}
	case 220:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:1195
		{
			clause := &ast.LimitClause{}
			clause.SetTag(ast.AST_EMPTY)
			yyVAL.node = clause
		}
	case 221:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:1202
		{
			yyVAL.node = &ast.LimitClause{Limit: yyDollar[1].list, Offset: yyDollar[2].expr}
		}
	case 222:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:1203
		{
			yyVAL.node = &ast.LimitClause{Limit: yyDollar[2].list, Offset: yyDollar[1].expr}
		}
	case 223:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1204
		{
			yyVAL.node = &ast.LimitClause{Limit: yyDollar[1].list}
		}
	case 224:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1205
		{
			yyVAL.node = &ast.LimitClause{Offset: yyDollar[1].expr}
		}
	case 225:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:1211
		{
			yyVAL.list = yyDollar[2].list
		}
	case 226:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:1215
		{
			yyVAL.list = append(yyDollar[2].list, yyDollar[4].expr)
		}
	case 227:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1220
		{
			yyVAL.list = []ast.ExprNode{yyDollar[1].expr}
		}
	case 228:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:1223
		{
			yyVAL.expr = yyDollar[2].expr
		}
	case 229:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1226
		{
			yyVAL.expr = yyDollar[1].expr
		}
	}
	goto yystack /* stack new state and value */
}
