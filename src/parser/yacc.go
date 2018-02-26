//line parse.y.go:2
package parser

import __yyfmt__ "fmt"

//line parse.y.go:4
import "ast"

//line parse.y.go:11
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
const NUMBER = 57357
const ALL = 57358
const ANY = 57359
const AS = 57360
const ASC = 57361
const AVG = 57362
const BY = 57363
const CAST = 57364
const CHARACTER = 57365
const COMMENT = 57366
const CONVERT = 57367
const COUNT = 57368
const DESC = 57369
const DISTINCT = 57370
const DISTINCTROW = 57371
const FALSE = 57372
const FOR = 57373
const FROM = 57374
const GROUP = 57375
const HAVING = 57376
const IF = 57377
const INTO = 57378
const LIMIT = 57379
const LOCK = 57380
const MAX = 57381
const MIN = 57382
const MODE = 57383
const NULL = 57384
const OFFSET = 57385
const ORDER = 57386
const POSITION = 57387
const QUARTER = 57388
const SELECT = 57389
const SET = 57390
const SHARE = 57391
const SOME = 57392
const SOUNDS = 57393
const SUM = 57394
const TO = 57395
const TRUE = 57396
const UPDATE = 57397
const WHERE = 57398
const LOW = 57399
const INTERVAL = 57400
const UNION = 57401
const EXCEPT = 57402
const INTERSECT = 57403
const JOIN = 57404
const STRAIGHT_JOIN = 57405
const LEFT = 57406
const RIGHT = 57407
const INNER = 57408
const OUTER = 57409
const CROSS = 57410
const NATURAL = 57411
const USE = 57412
const ON = 57413
const USING = 57414
const OR = 57415
const OO = 57416
const AND = 57417
const AA = 57418
const NOT = 57419
const BETWEEN = 57420
const CASE = 57421
const WHEN = 57422
const THEN = 57423
const ELSE = 57424
const IS = 57425
const LIKE = 57426
const REGEXP = 57427
const IN = 57428
const GE = 57429
const LE = 57430
const NE = 57431
const LG = 57432
const IE = 57433
const LNE = 57434
const SL = 57435
const SR = 57436
const LEG = 57437
const XOR = 57438
const EXISTS = 57439
const DIV = 57440
const MOD = 57441
const OP = 57442
const COLLATE = 57443
const UMINUS = 57444
const END = 57445

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
	"NUMBER",
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
	"'@'",
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
	117, 6,
	-2, 10,
	-1, 9,
	1, 7,
	117, 7,
	-2, 206,
	-1, 28,
	79, 147,
	88, 147,
	89, 147,
	90, 147,
	-2, 36,
	-1, 52,
	116, 122,
	-2, 76,
	-1, 149,
	31, 11,
	37, 11,
	38, 11,
	43, 11,
	44, 11,
	59, 11,
	60, 11,
	61, 11,
	-2, 64,
	-1, 275,
	118, 178,
	-2, 164,
	-1, 276,
	118, 179,
	-2, 165,
}

const yyPrivate = 57344

const yyLast = 1460

var yyAct = [...]int{

	240, 383, 381, 40, 4, 268, 355, 239, 333, 24,
	52, 16, 27, 267, 273, 361, 311, 317, 51, 24,
	108, 264, 242, 189, 109, 350, 116, 179, 187, 93,
	339, 289, 329, 91, 289, 112, 65, 332, 20, 367,
	150, 28, 287, 149, 29, 151, 94, 95, 153, 103,
	104, 106, 107, 103, 104, 106, 107, 152, 132, 133,
	400, 399, 156, 130, 131, 134, 135, 136, 137, 140,
	138, 139, 141, 49, 377, 105, 376, 157, 395, 105,
	145, 146, 147, 148, 65, 194, 405, 129, 276, 370,
	178, 180, 378, 289, 366, 183, 347, 188, 183, 183,
	289, 375, 298, 24, 197, 198, 199, 200, 201, 65,
	114, 192, 115, 113, 142, 191, 65, 194, 98, 110,
	276, 112, 97, 8, 280, 7, 168, 103, 104, 106,
	107, 205, 206, 132, 133, 167, 166, 165, 130, 131,
	134, 135, 136, 137, 140, 138, 139, 141, 343, 158,
	164, 8, 111, 105, 105, 288, 163, 162, 161, 160,
	159, 248, 154, 236, 237, 253, 254, 255, 256, 244,
	346, 235, 217, 218, 219, 220, 221, 223, 225, 226,
	227, 228, 229, 230, 143, 141, 174, 65, 87, 233,
	246, 203, 272, 245, 271, 250, 251, 252, 216, 21,
	243, 275, 134, 135, 136, 137, 140, 138, 139, 141,
	257, 278, 99, 243, 307, 292, 279, 144, 98, 277,
	7, 132, 133, 285, 272, 286, 130, 131, 134, 135,
	136, 137, 140, 138, 139, 141, 103, 104, 106, 107,
	84, 106, 107, 5, 293, 88, 294, 211, 296, 83,
	386, 299, 300, 301, 295, 281, 143, 283, 284, 282,
	180, 188, 105, 202, 291, 105, 313, 362, 103, 104,
	106, 107, 322, 323, 398, 387, 326, 176, 6, 345,
	365, 112, 112, 275, 149, 64, 328, 309, 331, 310,
	338, 324, 364, 340, 105, 335, 325, 100, 337, 330,
	65, 114, 196, 115, 113, 393, 336, 314, 354, 260,
	110, 344, 259, 102, 195, 351, 351, 351, 271, 271,
	271, 12, 360, 102, 170, 275, 275, 275, 358, 359,
	349, 352, 353, 175, 172, 182, 363, 112, 182, 182,
	192, 348, 372, 213, 368, 371, 265, 369, 63, 373,
	374, 325, 214, 215, 212, 173, 258, 379, 382, 90,
	204, 308, 103, 104, 106, 107, 103, 104, 106, 107,
	181, 271, 62, 184, 185, 89, 390, 85, 275, 388,
	89, 389, 391, 392, 86, 394, 90, 396, 105, 100,
	356, 271, 105, 312, 382, 190, 401, 357, 275, 403,
	96, 397, 271, 402, 404, 342, 19, 78, 79, 275,
	65, 48, 60, 47, 46, 43, 44, 45, 247, 65,
	334, 232, 67, 1, 74, 101, 241, 75, 68, 186,
	22, 23, 63, 136, 137, 140, 138, 139, 141, 304,
	177, 69, 70, 17, 50, 15, 3, 76, 103, 104,
	106, 107, 132, 133, 71, 380, 62, 130, 131, 134,
	135, 136, 137, 140, 138, 139, 141, 80, 103, 104,
	106, 107, 266, 32, 105, 103, 104, 106, 107, 56,
	25, 327, 61, 341, 58, 270, 82, 103, 104, 106,
	107, 297, 81, 34, 105, 103, 104, 106, 107, 30,
	14, 105, 269, 41, 33, 290, 35, 36, 249, 31,
	57, 234, 42, 105, 37, 11, 13, 12, 39, 155,
	26, 105, 38, 66, 78, 79, 55, 65, 48, 60,
	47, 46, 43, 44, 45, 210, 209, 53, 77, 67,
	262, 74, 65, 194, 75, 68, 193, 59, 263, 63,
	319, 316, 322, 323, 320, 54, 321, 318, 69, 70,
	85, 50, 274, 119, 76, 9, 89, 86, 2, 208,
	10, 71, 90, 62, 130, 131, 134, 135, 136, 137,
	140, 138, 139, 141, 80, 131, 134, 135, 136, 137,
	140, 138, 139, 141, 169, 103, 104, 106, 107, 61,
	171, 18, 128, 82, 120, 121, 125, 123, 117, 81,
	207, 122, 124, 315, 126, 305, 127, 73, 72, 118,
	41, 105, 231, 35, 36, 64, 306, 103, 104, 106,
	107, 37, 261, 65, 114, 39, 115, 113, 0, 38,
	66, 78, 79, 110, 65, 48, 60, 47, 46, 43,
	44, 45, 0, 105, 0, 303, 67, 0, 74, 302,
	0, 75, 68, 0, 22, 23, 63, 103, 104, 106,
	107, 103, 104, 106, 107, 69, 70, 0, 50, 92,
	0, 76, 103, 104, 106, 107, 0, 0, 71, 0,
	62, 22, 23, 105, 0, 65, 114, 105, 115, 113,
	0, 80, 0, 319, 316, 322, 323, 320, 105, 321,
	318, 0, 384, 385, 25, 0, 61, 0, 0, 0,
	82, 0, 0, 0, 0, 0, 81, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 41, 0, 0,
	35, 36, 64, 0, 0, 0, 0, 0, 37, 0,
	0, 0, 39, 0, 0, 0, 38, 66, 78, 79,
	0, 65, 48, 60, 47, 46, 43, 44, 45, 0,
	0, 0, 0, 67, 0, 74, 0, 0, 75, 68,
	0, 0, 0, 63, 0, 0, 0, 0, 0, 0,
	0, 0, 69, 70, 0, 50, 0, 0, 76, 0,
	8, 0, 0, 0, 0, 71, 0, 62, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 80, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 25, 0, 61, 0, 0, 0, 82, 0, 0,
	0, 0, 0, 81, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 41, 0, 0, 35, 36, 64,
	0, 0, 0, 0, 0, 37, 0, 0, 0, 39,
	0, 0, 0, 38, 66, 78, 79, 0, 65, 48,
	60, 47, 46, 43, 44, 45, 0, 0, 0, 0,
	67, 0, 74, 0, 0, 75, 68, 0, 0, 0,
	63, 0, 0, 0, 0, 0, 0, 0, 0, 69,
	70, 0, 50, 0, 0, 76, 0, 0, 0, 0,
	0, 0, 71, 0, 62, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 80, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 25, 0,
	61, 0, 0, 0, 82, 0, 0, 0, 0, 0,
	81, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 41, 0, 0, 35, 36, 64, 0, 0, 0,
	0, 0, 37, 0, 0, 0, 39, 238, 0, 0,
	38, 66, 78, 79, 0, 65, 48, 60, 47, 46,
	43, 44, 45, 0, 0, 0, 0, 67, 0, 74,
	0, 0, 75, 68, 0, 0, 0, 63, 0, 0,
	0, 0, 0, 0, 0, 0, 69, 70, 0, 50,
	0, 0, 76, 0, 0, 0, 0, 0, 0, 71,
	0, 62, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 80, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 25, 0, 61, 0, 0,
	0, 82, 0, 0, 0, 0, 0, 81, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 41, 0,
	0, 35, 36, 64, 0, 0, 0, 0, 0, 37,
	0, 0, 0, 39, 0, 0, 0, 38, 66, 78,
	79, 0, 65, 48, 60, 47, 46, 43, 44, 45,
	0, 0, 0, 0, 67, 0, 74, 0, 0, 75,
	68, 0, 0, 0, 63, 0, 0, 0, 0, 0,
	0, 0, 0, 69, 70, 0, 50, 0, 0, 76,
	0, 0, 0, 0, 0, 0, 71, 0, 62, 0,
	0, 0, 224, 0, 0, 0, 0, 0, 0, 80,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 61, 0, 0, 0, 82, 0,
	0, 0, 0, 0, 81, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 41, 0, 0, 35, 36,
	64, 0, 0, 0, 0, 0, 37, 0, 0, 0,
	39, 0, 0, 0, 38, 66, 78, 79, 0, 65,
	48, 60, 47, 46, 43, 44, 45, 0, 0, 0,
	0, 67, 0, 74, 0, 0, 75, 68, 0, 0,
	0, 63, 0, 0, 0, 0, 0, 0, 0, 0,
	69, 70, 0, 50, 0, 0, 76, 0, 0, 0,
	0, 0, 0, 71, 0, 62, 0, 0, 0, 222,
	0, 0, 0, 0, 0, 0, 80, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 61, 0, 0, 0, 82, 0, 0, 0, 0,
	0, 81, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 41, 0, 0, 35, 36, 64, 0, 0,
	0, 0, 0, 37, 0, 0, 0, 39, 0, 0,
	0, 38, 66, 78, 79, 0, 65, 48, 60, 47,
	46, 43, 44, 45, 0, 0, 0, 0, 67, 0,
	74, 0, 0, 75, 68, 0, 0, 0, 63, 0,
	0, 0, 0, 0, 0, 0, 0, 69, 70, 0,
	50, 0, 0, 76, 0, 0, 0, 0, 0, 0,
	71, 0, 62, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 80, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 61, 0,
	0, 0, 82, 0, 0, 0, 0, 0, 81, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 41,
	0, 0, 35, 36, 64, 0, 0, 0, 0, 0,
	37, 0, 0, 0, 39, 0, 0, 0, 38, 66,
}
var yyPact = [...]int{

	104, -1000, -1000, -1000, -1000, -1000, 456, 104, 636, -1000,
	529, 663, 663, 663, 379, 5, 1, 353, 987, 251,
	-1000, -1000, -1000, -1000, 292, 987, 520, -1000, 36, 142,
	-1000, -1000, -1000, -1000, -1000, 1338, 1338, 1338, 1338, 753,
	-1000, 9, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -61, -1000, -70, -1000, -1000, -1000, -1000, 46,
	-1000, 987, -1000, -1000, -1000, -1000, 28, 44, 43, 42,
	41, 40, 34, 21, 20, 19, 10, -1000, -1000, -1000,
	-1000, -1000, -1000, 343, 346, 300, 96, 316, 338, 987,
	987, 104, -1000, -1000, 104, 104, 987, -1000, -1000, 363,
	534, 261, 987, 987, 987, 987, 987, 987, -1000, -1000,
	687, -1000, -1000, -1000, -1000, -1000, 54, 318, 1338, 519,
	169, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 264, 110,
	1338, 1338, 1338, 1338, 1221, 1104, 1338, 1338, 1338, 1338,
	1338, 1338, -1000, 417, 1338, 70, 70, 70, 70, 1,
	394, -1000, 179, 411, 870, 119, 608, -1000, 411, 663,
	402, 663, 663, 663, 987, 987, 987, 987, 1338, -1000,
	-1000, -1000, -1000, -1000, 307, -1000, -1000, 250, 608, -1000,
	608, 260, -1000, -1000, -1000, 260, 247, -1000, 521, 290,
	108, -1000, -1000, -1000, -1000, 363, -1000, 165, 165, -1000,
	54, 54, -1000, -1000, -1000, -1000, -1000, 9, -1000, -1000,
	-1000, -1000, 8, 1338, 1338, 1338, 1338, 482, 98, 472,
	472, 327, 987, 327, 987, 74, 74, 74, 74, 74,
	-1000, -1000, -1000, 70, -1000, -1000, -76, -1000, -1000, 38,
	608, 132, -1000, 987, -1000, 987, 987, 987, 374, -15,
	987, 987, 987, 597, 593, 421, 553, 124, 320, 987,
	987, -1000, -1000, -1000, 360, 987, 245, 487, -1000, -1000,
	625, 625, 76, -1000, -86, -1000, -1000, 290, -1000, -1000,
	753, -39, 70, 355, 355, 413, 413, 179, -1000, 987,
	-89, -1000, 987, 401, 288, 31, 194, -1000, -1000, 162,
	53, -21, 283, 272, 411, 411, 411, 987, -1000, -1000,
	-1000, 356, 376, 608, 108, 108, 108, 199, 207, -1000,
	229, 217, -1000, -1000, -1000, -1000, 101, -23, 487, 534,
	360, -28, 1338, -1000, -1000, -1000, -1000, -1000, 608, -1000,
	608, 987, -1000, -1000, -1000, -1000, -1000, -1000, 987, 987,
	-16, -1000, -41, -43, -25, -1000, 987, 987, 487, 640,
	178, 212, -1000, 199, -1000, -1000, -1000, 108, -1000, 356,
	-1000, -1000, 608, 413, 413, -1000, -1000, -1000, -1000, 608,
	243, -1000, 608, -1000, 987, -38, 987, 108, 211, 487,
	-1000, -56, -57, 987, 608, 987, 608, 640, 108, -1000,
	-1000, -1000, -31, -1000, -1000, -1000,
}
var yyPgo = [...]int{

	0, 33, 632, 622, 618, 617, 613, 29, 610, 8,
	608, 602, 73, 17, 15, 601, 249, 23, 16, 6,
	212, 600, 594, 570, 278, 240, 445, 568, 3, 243,
	565, 21, 24, 152, 25, 563, 562, 555, 547, 10,
	18, 538, 20, 537, 14, 526, 41, 520, 519, 512,
	510, 509, 505, 0, 504, 2, 1, 502, 499, 493,
	245, 12, 485, 27, 44, 28, 484, 5, 13, 481,
	199, 479, 473, 22, 7, 472, 455, 188, 443, 440,
	429, 406, 426, 423,
}
var yyR1 = [...]int{

	0, 83, 27, 27, 28, 28, 26, 26, 26, 26,
	24, 24, 29, 29, 29, 29, 29, 53, 53, 53,
	53, 53, 53, 53, 53, 47, 47, 47, 47, 47,
	61, 61, 61, 61, 61, 61, 61, 46, 46, 46,
	46, 46, 46, 46, 46, 46, 46, 46, 46, 46,
	46, 46, 64, 64, 64, 64, 64, 64, 64, 64,
	64, 64, 64, 64, 64, 64, 64, 58, 58, 58,
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
	5, 1, 1, 1, 1, 1, 1, 3, 3, 2,
	2, 2, 2, 3, 1, 2, 1, 1, 1, 1,
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

	-1000, -83, -27, -26, -28, -29, -24, 116, 47, -30,
	-23, 59, 61, 60, 44, -26, -28, -78, -15, -81,
	-7, -70, 28, 29, -53, 78, -47, -61, -46, -64,
	-58, -51, -72, -54, -59, 104, 105, 112, 120, 116,
	-28, 101, -49, 13, 14, 15, 12, 11, 9, -12,
	42, -40, -39, -43, -37, -45, -71, -50, -66, -38,
	10, 80, 54, 30, 106, 8, 121, 20, 26, 39,
	40, 52, -4, -5, 22, 25, 45, -41, 5, 6,
	65, 90, 84, -16, -25, 31, 38, -77, -60, 37,
	43, -1, 16, -7, -1, -1, 21, 117, 117, -20,
	36, -81, 62, 74, 75, 100, 76, 77, -42, -32,
	18, -33, -39, 12, 9, 11, -53, -10, 99, -35,
	84, 85, 91, 87, 92, 86, 94, 96, -11, 51,
	102, 103, 97, 98, 104, 105, 106, 107, 109, 110,
	108, 111, 78, 114, 75, -64, -64, -64, -64, -28,
	-53, -28, 118, 118, 116, -48, -53, -39, 121, 116,
	116, 116, 116, 116, 116, 116, 116, 116, 116, -22,
	-25, -21, -16, 55, 90, -60, -77, -79, -53, -63,
	-53, -24, -29, -28, -24, -24, -80, -65, -53, -17,
	32, -44, -39, 12, 9, -20, -70, -53, -53, -53,
	-53, -53, -33, -12, 42, -61, -61, -8, 50, 17,
	16, 78, 90, 79, 88, 89, 88, -46, -46, -46,
	-46, -46, 58, -46, 58, -46, -46, -46, -46, -46,
	-46, -3, 4, -64, 117, -40, -39, -39, 117, -74,
	-53, -82, -73, 81, -39, -1, -7, 16, -53, 106,
	-1, -1, -1, -53, -53, -53, -53, -46, 49, 62,
	62, -2, 19, 27, -31, 56, -75, -68, -67, -57,
	-62, -28, 116, -44, -36, -39, 12, -17, -28, -28,
	116, -46, -64, -46, -46, -53, -53, 118, 117, 62,
	-52, -73, 83, -53, -53, -74, -53, 117, 117, -53,
	-53, -53, 62, 62, 18, 62, 73, 90, 41, -63,
	-65, -18, 33, -53, 62, -6, 64, -13, 70, 63,
	67, 69, 65, 66, -42, -32, -28, -69, -68, 118,
	-31, -74, 76, -9, 7, -9, -40, -39, -53, 119,
	-53, 82, 117, 117, 117, 117, 117, 117, 58, 58,
	-34, -39, -34, -34, -53, -19, 34, 21, -68, -68,
	-67, -14, 68, -13, 63, 63, 117, 62, -44, -18,
	117, -61, -53, -53, -53, 117, 117, 117, 117, -53,
	-76, -55, -53, -56, 72, 73, 72, 63, -14, -68,
	-19, -9, -9, 62, -53, 116, -53, -68, 63, 117,
	117, -55, -74, -56, -67, 117,
}
var yyDef = [...]int{

	0, -2, 1, 2, -2, -2, 207, 0, 116, -2,
	0, 159, 159, 159, 0, 0, 11, 163, 0, 115,
	156, 117, 160, 161, 125, 0, 23, 29, -2, 51,
	52, 53, 54, 55, 56, 0, 0, 0, 0, 0,
	64, 0, 66, 67, 68, 69, 70, 71, 72, 73,
	74, 81, -2, 78, 79, 84, 85, 86, 87, 0,
	106, 109, 132, 133, 77, 75, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 123, 154, 155,
	150, 151, 152, 220, 216, 0, 0, 223, 224, 0,
	0, 0, 157, 158, 0, 0, 0, 4, 5, 168,
	0, 163, 0, 0, 0, 0, 0, 0, 119, 124,
	0, 127, 128, 129, 130, 131, 22, 0, 0, 0,
	144, 137, 138, 139, 140, 141, 142, 143, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 146, 0, 0, 59, 60, 61, 62, -2,
	0, 65, 0, 0, 0, 0, 108, 120, 0, 159,
	0, 159, 159, 159, 0, 0, 0, 0, 0, 8,
	219, 9, 215, 217, 0, 221, 222, 225, 227, 228,
	229, 14, 10, 11, 15, 16, 208, 209, 214, 198,
	0, 162, 164, 165, 166, 168, 118, 17, 18, 19,
	20, 21, 126, 24, 25, 26, 27, 0, 134, 135,
	136, 145, 0, 0, 0, 0, 0, 37, 38, 39,
	40, 41, 0, 42, 0, 43, 44, 45, 46, 47,
	48, 57, 149, 58, 63, 82, 76, 80, 88, 0,
	104, 114, 110, 0, 121, 0, 0, 0, 0, 77,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 211, 212, 213, 200, 0, 167, 169, 171, 172,
	125, 0, 0, 176, 0, -2, -2, 198, 28, 30,
	0, 0, 34, 35, 33, 0, 0, 0, 89, 0,
	0, 111, 0, 0, 0, 0, 0, 93, 94, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 218, 226,
	210, 205, 0, 197, 0, 0, 0, 194, 0, 188,
	0, 0, 191, 192, 173, 174, 11, 0, 180, 0,
	200, 0, 0, 49, 148, 50, 83, 76, 105, 107,
	113, 0, 90, 91, 92, 95, 96, 97, 0, 0,
	0, 153, 0, 0, 0, 12, 0, 0, 170, 182,
	184, 0, 193, 194, 189, 190, 175, 0, 177, 205,
	31, 32, 112, 0, 0, 100, 101, 102, 103, 204,
	199, 201, 203, 183, 0, 0, 0, 0, 0, 181,
	13, 0, 0, 0, 195, 0, 185, 0, 0, 98,
	99, 202, 0, 186, 187, 196,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 120, 3, 3, 3, 108, 103, 3,
	116, 117, 106, 104, 62, 105, 118, 107, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	86, 85, 87, 3, 121, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 111, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 102, 3, 112,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
	52, 53, 54, 55, 56, 57, 58, 59, 60, 61,
	63, 64, 65, 66, 67, 68, 69, 70, 71, 72,
	73, 74, 75, 76, 77, 78, 79, 80, 81, 82,
	83, 84, 88, 89, 90, 91, 92, 93, 94, 95,
	96, 97, 98, 99, 100, 101, 109, 110, 113, 114,
	115, 119,
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
		//line parse.y.go:228
		{
			yylex.(*Tokener).ParseStmt = yyDollar[1].node
		}
	case 4:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y.go:247
		{
			yyVAL.node = yyDollar[2].node
		}
	case 5:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y.go:251
		{
			yyVAL.node = yyDollar[2].node
		}
	case 6:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y.go:257
		{
			yyVAL.node = yyDollar[1].node
		}
	case 7:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parse.y.go:261
		{
			switch node := yyDollar[1].node.(type) {
			case *ast.SimpleSelectStmt:
				node.Sort = yyDollar[2].node.(*ast.SortClause)
			case *ast.SelectStmt:
				node.Sort = yyDollar[2].node.(*ast.SortClause)
			}

			yyVAL.node = yyDollar[1].node
		}
	case 8:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parse.y.go:272
		{
			switch node := yyDollar[1].node.(type) {
			case *ast.SimpleSelectStmt:
				if yyDollar[2].node != nil {
					node.Sort = yyDollar[2].node.(*ast.SortClause)
				}
				if yyDollar[3].node != nil {
					node.Lock = yyDollar[3].node.(*ast.LockClause)
				}
				if yyDollar[4].node != nil {
					node.Limit = yyDollar[4].node.(*ast.LimitClause)
				}
			case *ast.SelectStmt:
				if yyDollar[2].node != nil {
					node.Sort = yyDollar[2].node.(*ast.SortClause)
				}
				if yyDollar[3].node != nil {
					node.Lock = yyDollar[3].node.(*ast.LockClause)
				}
				if yyDollar[4].node != nil {
					node.Limit = yyDollar[4].node.(*ast.LimitClause)
				}
			}

			yyVAL.node = yyDollar[1].node
		}
	case 9:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parse.y.go:299
		{
			switch node := yyDollar[1].node.(type) {
			case *ast.SimpleSelectStmt:
				if yyDollar[2].node != nil {
					node.Sort = yyDollar[2].node.(*ast.SortClause)
				}
				if yyDollar[4].node != nil {
					node.Lock = yyDollar[4].node.(*ast.LockClause)
				}
				if yyDollar[3].node != nil {
					node.Limit = yyDollar[3].node.(*ast.LimitClause)
				}
			case *ast.SelectStmt:
				if yyDollar[2].node != nil {
					node.Sort = yyDollar[2].node.(*ast.SortClause)
				}
				if yyDollar[4].node != nil {
					node.Lock = yyDollar[4].node.(*ast.LockClause)
				}
				if yyDollar[3].node != nil {
					node.Limit = yyDollar[3].node.(*ast.LimitClause)
				}
			}

			yyVAL.node = yyDollar[1].node
		}
	case 10:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y.go:327
		{
			yyVAL.node = yyDollar[1].node
		}
	case 11:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y.go:328
		{
			yyVAL.node = yyDollar[1].node
		}
	case 12:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parse.y.go:333
		{
			stmt := &ast.SimpleSelectStmt{}
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

			yyVAL.node = stmt
		}
	case 13:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line parse.y.go:355
		{
			stmt := &ast.SimpleSelectStmt{Distinct: yyDollar[2].node.(*ast.DistinctClause), Target: &ast.TargetClause{Target_ref: yyDollar[3].list}}
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

			yyVAL.node = stmt
		}
	case 14:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parse.y.go:376
		{
			stmt := &ast.SelectStmt{DistinctType: yyDollar[3].tag, Left: yyDollar[1].node.(*ast.SelectStmt), Right: yyDollar[4].node.(*ast.SelectStmt)}
			stmt.SetTag(ast.Union)
			yyVAL.node = stmt
		}
	case 15:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parse.y.go:382
		{
			stmt := &ast.SelectStmt{DistinctType: yyDollar[3].tag, Left: yyDollar[1].node.(*ast.SelectStmt), Right: yyDollar[4].node.(*ast.SelectStmt)}
			stmt.SetTag(ast.Except)
			yyVAL.node = stmt
		}
	case 16:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parse.y.go:388
		{
			stmt := &ast.SelectStmt{DistinctType: yyDollar[3].tag, Left: yyDollar[1].node.(*ast.SelectStmt), Right: yyDollar[4].node.(*ast.SelectStmt)}
			stmt.SetTag(ast.Intersect)
			yyVAL.node = stmt
		}
	case 17:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y.go:403
		{
			yyVAL.expr = &ast.Expr{Operator: "OR", Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 18:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y.go:404
		{
			yyVAL.expr = &ast.Expr{Operator: "||", Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 19:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y.go:405
		{
			yyVAL.expr = &ast.Expr{Operator: "XOR", Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 20:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y.go:406
		{
			yyVAL.expr = &ast.Expr{Operator: "AND", Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 21:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y.go:407
		{
			yyVAL.expr = &ast.Expr{Operator: "&&", Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 22:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parse.y.go:408
		{
			yyVAL.expr = &ast.Expr{Operator: "NOT", Left: nil, Right: yyDollar[2].expr}
		}
	case 23:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y.go:409
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 24:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y.go:411
		{
			yyVAL.expr = &ast.IsTrueExpr{HasNot: yyDollar[2].boolean, Left: yyDollar[1].expr}
		}
	case 25:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y.go:417
		{
			yyVAL.expr = &ast.IsNullExpr{HasNot: yyDollar[2].boolean, Left: yyDollar[1].expr}
		}
	case 26:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y.go:421
		{
			yyVAL.expr = &ast.Expr{Operator: "<=>", Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 27:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y.go:425
		{
			yyVAL.expr = &ast.Expr{Operator: yyDollar[2].val, Left: yyDollar[1].expr, Right: yyDollar[3].expr}

		}
	case 28:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parse.y.go:430
		{
			expr := &ast.SubqueryExpr{Operator: yyDollar[2].val, Subtype: yyDollar[3].tag, Left: yyDollar[1].expr, Select: yyDollar[4].node}
			expr.SetTag(ast.Subquery_Operator)
			yyVAL.expr = expr
		}
	case 29:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y.go:436
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 30:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parse.y.go:442
		{
			expr := &ast.SubqueryExpr{HasNot: yyDollar[2].boolean, Left: yyDollar[1].expr, Select: yyDollar[4].node}
			expr.SetTag(ast.Subquery_In)
			yyVAL.expr = expr
		}
	case 31:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parse.y.go:448
		{
			yyVAL.expr = &ast.InExpr{HasNot: yyDollar[2].boolean, Left: yyDollar[1].expr, Right: yyDollar[5].list}
		}
	case 32:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parse.y.go:452
		{
			yyVAL.expr = &ast.BetweenExpr{HasNot: yyDollar[2].boolean, Expr: yyDollar[1].expr, Left: yyDollar[4].expr, Right: yyDollar[6].expr}
		}
	case 33:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parse.y.go:456
		{
			yyVAL.expr = &ast.LikeExpr{HasNot: false, Left: yyDollar[1].expr, Right: yyDollar[4].expr}
		}
	case 34:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parse.y.go:460
		{
			yyVAL.expr = &ast.LikeExpr{HasNot: yyDollar[2].boolean, Left: yyDollar[1].expr, Right: yyDollar[4].expr}
		}
	case 35:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parse.y.go:464
		{
			yyVAL.expr = &ast.RegexpExpr{HasNot: yyDollar[2].boolean, Left: yyDollar[1].expr, Right: yyDollar[4].expr}
		}
	case 36:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y.go:468
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 37:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y.go:474
		{
			yyVAL.expr = &ast.Expr{Operator: "|", Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 38:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y.go:475
		{
			yyVAL.expr = &ast.Expr{Operator: "&", Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 39:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y.go:476
		{
			yyVAL.expr = &ast.Expr{Operator: "<<", Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 40:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y.go:477
		{
			yyVAL.expr = &ast.Expr{Operator: ">>", Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 41:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y.go:478
		{
			yyVAL.expr = &ast.Expr{Operator: "+", Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 42:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y.go:479
		{
			yyVAL.expr = &ast.Expr{Operator: "-", Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 43:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y.go:480
		{
			yyVAL.expr = &ast.Expr{Operator: "*", Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 44:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y.go:481
		{
			yyVAL.expr = &ast.Expr{Operator: "/", Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 45:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y.go:482
		{
			yyVAL.expr = &ast.Expr{Operator: "DIV", Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 46:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y.go:483
		{
			yyVAL.expr = &ast.Expr{Operator: "MOD", Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 47:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y.go:484
		{
			yyVAL.expr = &ast.Expr{Operator: "%", Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 48:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y.go:485
		{
			yyVAL.expr = &ast.Expr{Operator: "^", Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 49:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parse.y.go:487
		{
			yyVAL.expr = &ast.IntervalExpr{Operator: ast.Op_Add, Left: yyDollar[1].expr, Right: yyDollar[4].expr, TimeUnit: yyDollar[5].tag}
		}
	case 50:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parse.y.go:491
		{
			yyVAL.expr = &ast.IntervalExpr{Operator: ast.Op_Minus, Left: yyDollar[1].expr, Right: yyDollar[4].expr, TimeUnit: yyDollar[5].tag}
		}
	case 51:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y.go:495
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 52:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y.go:500
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 53:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y.go:501
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 54:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y.go:502
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 55:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y.go:503
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 56:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y.go:504
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 57:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y.go:506
		{
			yyVAL.expr = &ast.CollateExpr{Expr: yyDollar[1].expr, Collate: yyDollar[3].tag}
		}
	case 58:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y.go:510
		{
			yyVAL.expr = &ast.Expr{Operator: "||", Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 59:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parse.y.go:514
		{
			yyVAL.expr = yyDollar[2].expr
		}
	case 60:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parse.y.go:518
		{
			yyDollar[2].expr.(*ast.Number).Symbol = !yyDollar[2].expr.(*ast.Number).Symbol
			yyVAL.expr = yyDollar[2].expr
		}
	case 61:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parse.y.go:523
		{
			yyVAL.expr = &ast.Expr{Operator: "~", Left: nil, Right: yyDollar[2].expr}
		}
	case 62:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parse.y.go:527
		{
			yyVAL.expr = &ast.Expr{Operator: "!", Left: nil, Right: yyDollar[2].expr}
		}
	case 63:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y.go:531
		{
			yyVAL.expr = yyDollar[2].expr
		}
	case 64:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y.go:536
		{
			expr := &ast.SubqueryExpr{Select: yyDollar[1].node}
			expr.SetTag(ast.Subquery_Tuple)
			yyVAL.expr = expr
		}
	case 65:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parse.y.go:542
		{
			expr := &ast.SubqueryExpr{Select: yyDollar[2].node}
			expr.SetTag(ast.Subquery_Exists)
			yyVAL.expr = expr
		}
	case 66:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y.go:548
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 67:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y.go:554
		{
			Expr := &ast.Number{Symbol: true, Val: yyDollar[1].val}
			Expr.SetTag(ast.Bit)
			yyVAL.expr = Expr
		}
	case 68:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y.go:560
		{
			Expr := &ast.Number{Symbol: true, Val: yyDollar[1].val}
			Expr.SetTag(ast.Hex)
			yyVAL.expr = Expr
		}
	case 69:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y.go:566
		{
			Expr := &ast.Number{Symbol: true, Val: yyDollar[1].val}
			Expr.SetTag(ast.Digit)
			yyVAL.expr = Expr
		}
	case 70:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y.go:571
		{
			yyVAL.expr = &ast.String{Str: yyDollar[1].str}
		}
	case 71:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y.go:572
		{
			yyVAL.expr = &ast.String{Str: yyDollar[1].str}
		}
	case 72:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y.go:573
		{
			yyVAL.expr = &ast.String{Str: yyDollar[1].str}
		}
	case 73:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y.go:574
		{
			yyVAL.expr = &ast.True{IsTrue: yyDollar[1].boolean}
		}
	case 74:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y.go:575
		{
			yyVAL.expr = &ast.Null{}
		}
	case 75:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y.go:579
		{
			yyVAL.val = yyDollar[1].ident
		}
	case 76:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y.go:583
		{
			yyVAL.val = yyDollar[1].val
		}
	case 77:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y.go:587
		{
			yyVAL.val = "*"
		}
	case 78:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y.go:592
		{
			yyVAL.expr = &ast.UserVariable{Variable: yyDollar[1].val}
		}
	case 79:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y.go:593
		{
			yyVAL.expr = &ast.SystemVariable{Schema: yyDollar[1].val}
		}
	case 80:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y.go:595
		{
			yyVAL.expr = &ast.SystemVariable{Schema: yyDollar[1].val, Parameter: yyDollar[3].val}
		}
	case 81:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y.go:601
		{
			if yyDollar[1].val == "*" {
				yyVAL.expr = &ast.Column{IsStar: true}
			} else {
				yyVAL.expr = &ast.Column{Column: yyDollar[1].val}
			}
		}
	case 82:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y.go:609
		{
			if yyDollar[3].val == "*" {
				yyVAL.expr = &ast.Column{Relation: &ast.Relation{Table: yyDollar[1].val}, IsStar: true}
			} else {
				yyVAL.expr = &ast.Column{Relation: &ast.Relation{Table: yyDollar[1].val}, Column: yyDollar[3].val}
			}
		}
	case 83:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parse.y.go:617
		{
			if yyDollar[5].val == "*" {
				yyVAL.expr = &ast.Column{Relation: &ast.Relation{DB: yyDollar[1].val, Table: yyDollar[1].val}, IsStar: true}
			} else {
				yyVAL.expr = &ast.Column{Relation: &ast.Relation{DB: yyDollar[1].val, Table: yyDollar[3].val}, Column: yyDollar[5].val}
			}
		}
	case 84:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y.go:628
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 85:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y.go:632
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 86:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y.go:636
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 87:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y.go:640
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 88:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y.go:644
		{
			yyVAL.expr = &ast.FuncCall{Name: yyDollar[1].val}
		}
	case 89:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parse.y.go:648
		{
			yyVAL.expr = &ast.FuncCall{Name: yyDollar[1].val, Arg: yyDollar[3].list}
		}
	case 90:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parse.y.go:654
		{
			expr := &ast.AggretionFuncCall{DistinctType: yyDollar[3].tag, Arg: []ast.ExprNode{yyDollar[4].expr}}
			expr.SetTag(ast.Aggretion_Avg)
			yyVAL.expr = expr
		}
	case 91:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parse.y.go:660
		{
			expr := &ast.AggretionFuncCall{DistinctType: yyDollar[3].tag, Arg: yyDollar[4].list}
			expr.SetTag(ast.Aggretion_Count)
			yyVAL.expr = expr
		}
	case 92:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parse.y.go:666
		{
			expr := &ast.AggretionFuncCall{DistinctType: ast.AST_ALL, Arg: []ast.ExprNode{yyDollar[4].expr}}
			expr.SetTag(ast.Aggretion_Count)
			yyVAL.expr = expr
		}
	case 93:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parse.y.go:672
		{
			expr := &ast.AggretionFuncCall{DistinctType: ast.AST_EMPTY, Arg: []ast.ExprNode{yyDollar[3].expr}}
			expr.SetTag(ast.Aggretion_Count)
			yyVAL.expr = expr
		}
	case 94:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parse.y.go:678
		{
			expr := &ast.AggretionFuncCall{DistinctType: ast.AST_EMPTY, HasStar: true}
			expr.SetTag(ast.Aggretion_Count)
			yyVAL.expr = expr
		}
	case 95:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parse.y.go:684
		{
			expr := &ast.AggretionFuncCall{DistinctType: yyDollar[3].tag, Arg: []ast.ExprNode{yyDollar[4].expr}}
			expr.SetTag(ast.Aggretion_Max)
			yyVAL.expr = expr
		}
	case 96:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parse.y.go:690
		{
			expr := &ast.AggretionFuncCall{DistinctType: yyDollar[3].tag, Arg: []ast.ExprNode{yyDollar[4].expr}}
			expr.SetTag(ast.Aggretion_Min)
			yyVAL.expr = expr
		}
	case 97:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parse.y.go:696
		{
			expr := &ast.AggretionFuncCall{DistinctType: yyDollar[3].tag, Arg: []ast.ExprNode{yyDollar[4].expr}}
			expr.SetTag(ast.Aggretion_Sum)
			yyVAL.expr = expr
		}
	case 98:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line parse.y.go:704
		{
			Expr := &ast.CalTimeFuncCall{Interval: &ast.IntervalExpr{Operator: ast.Op_Add, Left: yyDollar[3].expr, Right: yyDollar[6].expr, TimeUnit: yyDollar[7].tag}}
			Expr.SetTag(yyDollar[1].tag)
			yyVAL.expr = Expr
		}
	case 99:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line parse.y.go:710
		{
			Expr := &ast.CalTimeFuncCall{Interval: &ast.IntervalExpr{Operator: ast.Op_Minus, Left: yyDollar[3].expr, Right: yyDollar[6].expr, TimeUnit: yyDollar[7].tag}}
			Expr.SetTag(yyDollar[1].tag)
			yyVAL.expr = Expr
		}
	case 100:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parse.y.go:718
		{
			yyVAL.expr = &ast.CastFuncCall{Expr: yyDollar[3].expr, CastType: yyDollar[5].val}
		}
	case 101:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parse.y.go:722
		{
			yyVAL.expr = &ast.CastFuncCall{Expr: yyDollar[3].expr, CastType: yyDollar[5].val}
		}
	case 102:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parse.y.go:726
		{
			yyVAL.expr = &ast.CastFuncCall{Expr: yyDollar[3].expr, CastType: yyDollar[5].val}
		}
	case 103:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parse.y.go:732
		{
			yyVAL.expr = &ast.StringFuncCall{Substr: yyDollar[3].expr, Str: yyDollar[5].expr}
		}
	case 104:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y.go:737
		{
			yyVAL.list = []ast.ExprNode{yyDollar[1].expr}
		}
	case 105:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y.go:738
		{
			yyVAL.list = append(yyDollar[1].list, yyDollar[3].expr)
		}
	case 106:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y.go:741
		{
			yyVAL.expr = &ast.Marker{Str: yyDollar[1].str}
		}
	case 107:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parse.y.go:745
		{
			yyVAL.expr = &ast.CaseExpr{Case: yyDollar[2].expr, When: yyDollar[3].list, Else: yyDollar[4].expr}
		}
	case 108:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y.go:750
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 109:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parse.y.go:751
		{
			yyVAL.expr = nil
		}
	case 110:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y.go:754
		{
			yyVAL.list = []ast.ExprNode{yyDollar[1].expr}
		}
	case 111:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parse.y.go:755
		{
			yyVAL.list = append(yyDollar[1].list, yyDollar[2].expr)
		}
	case 112:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parse.y.go:758
		{
			yyVAL.expr = &ast.Expr{Left: yyDollar[2].expr, Right: yyDollar[4].expr}
		}
	case 113:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parse.y.go:761
		{
			yyVAL.expr = &ast.Expr{Left: yyDollar[2].expr}
		}
	case 114:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parse.y.go:762
		{
			yyVAL.expr = nil
		}
	case 115:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y.go:765
		{
			yyVAL.list = yyDollar[1].list
		}
	case 116:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parse.y.go:766
		{
			yyVAL.list = nil
		}
	case 117:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y.go:769
		{
			yyVAL.list = []ast.ExprNode{yyDollar[1].expr}
		}
	case 118:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y.go:770
		{
			yyVAL.list = append(yyDollar[1].list, yyDollar[3].expr)
		}
	case 119:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parse.y.go:773
		{
			yyVAL.expr = &ast.Tuple{Ref: yyDollar[1].expr, Alias: yyDollar[2].val}
		}
	case 120:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parse.y.go:776
		{
			yyVAL.val = yyDollar[2].val
		}
	case 121:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y.go:779
		{
			yyVAL.val = yyDollar[3].val
		}
	case 122:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y.go:782
		{
			yyVAL.val = yyDollar[1].val
		}
	case 123:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y.go:783
		{
			yyVAL.val = yyDollar[1].val
		}
	case 124:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y.go:786
		{
			yyVAL.val = yyDollar[1].val
		}
	case 125:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parse.y.go:787
		{
			yyVAL.val = ""
		}
	case 126:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parse.y.go:790
		{
			yyVAL.val = yyDollar[2].val
		}
	case 127:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y.go:791
		{
			yyVAL.val = yyDollar[1].val
		}
	case 128:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y.go:794
		{
			yyVAL.val = yyDollar[1].val
		}
	case 129:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y.go:795
		{
			yyVAL.val = yyDollar[1].str
		}
	case 130:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y.go:796
		{
			yyVAL.val = yyDollar[1].str
		}
	case 131:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y.go:797
		{
			yyVAL.val = yyDollar[1].str
		}
	case 132:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y.go:801
		{
			yyVAL.boolean = true
		}
	case 133:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y.go:802
		{
			yyVAL.boolean = false
		}
	case 134:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y.go:807
		{
			yyVAL.tag = ast.Operator_Some
		}
	case 135:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y.go:808
		{
			yyVAL.tag = ast.Operator_All
		}
	case 136:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y.go:809
		{
			yyVAL.tag = ast.Operator_Any
		}
	case 137:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y.go:812
		{
			yyVAL.val = "="
		}
	case 138:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y.go:813
		{
			yyVAL.val = ">="
		}
	case 139:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y.go:814
		{
			yyVAL.val = ">"
		}
	case 140:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y.go:815
		{
			yyVAL.val = "<="
		}
	case 141:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y.go:816
		{
			yyVAL.val = "<"
		}
	case 142:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y.go:817
		{
			yyVAL.val = "<>"
		}
	case 143:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y.go:818
		{
			yyVAL.val = "!="
		}
	case 144:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y.go:821
		{
			yyVAL.boolean = false
		}
	case 145:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parse.y.go:822
		{
			yyVAL.boolean = true
		}
	case 146:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y.go:826
		{
			yyVAL.boolean = true
		}
	case 147:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parse.y.go:827
		{
			yyVAL.boolean = false
		}
	case 148:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y.go:830
		{
			yyVAL.tag = ast.TimeUnit[yyDollar[1].ident]
		}
	case 149:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y.go:834
		{
			yyVAL.tag = ast.Character[yyDollar[1].ident]
		}
	case 150:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y.go:838
		{
			yyVAL.val = "LEFT"
		}
	case 151:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y.go:839
		{
			yyVAL.val = "IN"
		}
	case 152:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y.go:840
		{
			yyVAL.val = "IS"
		}
	case 153:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y.go:843
		{
			yyVAL.val = yyDollar[1].val
		}
	case 154:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y.go:846
		{
			yyVAL.tag = ast.FuncTime[yyDollar[1].ident]
		}
	case 155:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y.go:849
		{
			yyVAL.tag = ast.FuncTime[yyDollar[1].ident]
		}
	case 156:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y.go:860
		{
			clause := &ast.DistinctClause{}
			clause.SetTag(yyDollar[1].tag)
			yyVAL.node = clause
		}
	case 157:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y.go:867
		{
			yyVAL.tag = ast.AST_ALL
		}
	case 158:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y.go:868
		{
			yyVAL.tag = yyDollar[1].tag
		}
	case 159:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parse.y.go:869
		{
			yyVAL.tag = ast.AST_EMPTY
		}
	case 160:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y.go:872
		{
			yyVAL.tag = ast.AST_DISTINCT
		}
	case 161:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y.go:873
		{
			yyVAL.tag = ast.AST_DISTINCTROW
		}
	case 162:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parse.y.go:883
		{
			yyVAL.node = &ast.IntoClause{Relation: &ast.Relation{Table: yyDollar[2].val}}
		}
	case 163:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parse.y.go:884
		{
			yyVAL.node = nil
		}
	case 164:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y.go:887
		{
			yyVAL.val = yyDollar[1].val
		}
	case 165:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y.go:888
		{
			yyVAL.val = yyDollar[1].str
		}
	case 166:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y.go:889
		{
			yyVAL.val = yyDollar[1].str
		}
	case 167:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parse.y.go:899
		{
			yyVAL.node = &ast.FromClause{Table_ref: yyDollar[2].list}
		}
	case 168:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parse.y.go:900
		{
			yyVAL.node = nil
		}
	case 169:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y.go:904
		{
			yyVAL.list = []ast.ExprNode{yyDollar[1].expr}
		}
	case 170:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y.go:905
		{
			yyVAL.list = append(yyDollar[1].list, yyDollar[3].expr)
		}
	case 171:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y.go:908
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 172:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y.go:909
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 173:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parse.y.go:913
		{
			yyDollar[1].expr.(*ast.Relation).Alias = yyDollar[2].val
			yyVAL.expr = yyDollar[1].expr
		}
	case 174:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parse.y.go:918
		{
			tb := &ast.SubqueryExpr{Select: yyDollar[1].node, Alias: yyDollar[2].val}
			tb.SetTag(ast.Subquery_Relation)
			yyVAL.expr = tb
		}
	case 175:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y.go:924
		{
			yyVAL.expr = yyDollar[2].expr
		}
	case 176:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y.go:929
		{
			yyVAL.expr = &ast.Relation{Table: yyDollar[1].val}
		}
	case 177:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y.go:930
		{
			yyVAL.expr = &ast.Relation{DB: yyDollar[1].val, Table: yyDollar[2].val}
		}
	case 178:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y.go:933
		{
			yyVAL.val = yyDollar[1].val
		}
	case 179:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y.go:934
		{
			yyVAL.val = yyDollar[1].str
		}
	case 180:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y.go:937
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 181:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y.go:939
		{
			rel := &ast.Join{Left: yyDollar[1].expr, Right: yyDollar[3].expr}
			rel.SetTag(ast.Cross_join)
			yyVAL.expr = rel
		}
	case 182:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y.go:950
		{
			rel := &ast.Join{Left: yyDollar[1].expr, Right: yyDollar[3].expr}
			rel.SetTag(ast.Cross_join)
			yyVAL.expr = rel
		}
	case 183:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parse.y.go:956
		{
			rel := &ast.Join{Left: yyDollar[1].expr, Right: yyDollar[3].expr, Cndition: yyDollar[4].expr}
			rel.SetTag(ast.Inner_join)
			yyVAL.expr = rel
		}
	case 184:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y.go:962
		{
			rel := &ast.Join{Left: yyDollar[1].expr, Right: yyDollar[3].expr}
			rel.SetTag(ast.Straight_join)
			yyVAL.expr = rel
		}
	case 185:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parse.y.go:968
		{
			rel := &ast.Join{Left: yyDollar[1].expr, Right: yyDollar[3].expr, Cndition: yyDollar[5].expr}
			rel.SetTag(ast.Straight_join)
			yyVAL.expr = rel
		}
	case 186:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parse.y.go:974
		{
			rel := &ast.Join{HasOuter: yyDollar[3].boolean, Left: yyDollar[1].expr, Right: yyDollar[5].expr, Cndition: yyDollar[6].expr}
			if yyDollar[2].boolean {
				rel.SetTag(ast.Left_join)
			} else {
				rel.SetTag(ast.Right_join)
			}
			yyVAL.expr = rel
		}
	case 187:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parse.y.go:984
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
		//line parse.y.go:995
		{
			yyVAL.tag = 1
		}
	case 189:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parse.y.go:996
		{
			yyVAL.tag = 2
		}
	case 190:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parse.y.go:997
		{
			yyVAL.tag = 3
		}
	case 191:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y.go:1001
		{
			yyVAL.boolean = true
		}
	case 192:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y.go:1002
		{
			yyVAL.boolean = false
		}
	case 193:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y.go:1005
		{
			yyVAL.boolean = true
		}
	case 194:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parse.y.go:1006
		{
			yyVAL.boolean = false
		}
	case 195:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parse.y.go:1009
		{
			yyVAL.expr = yyDollar[2].expr
		}
	case 196:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parse.y.go:1010
		{
			yyVAL.expr = &ast.Using{Column: yyDollar[3].list}
		}
	case 197:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parse.y.go:1020
		{
			yyVAL.node = &ast.WhereClause{Where: yyDollar[2].expr}
		}
	case 198:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parse.y.go:1021
		{
			yyVAL.node = nil
		}
	case 199:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y.go:1030
		{
			yyVAL.node = &ast.GroupClause{Target_ref: yyDollar[3].list}
		}
	case 200:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parse.y.go:1031
		{
			yyVAL.node = nil
		}
	case 201:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y.go:1035
		{
			yyVAL.list = []ast.ExprNode{yyDollar[1].expr}
		}
	case 202:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y.go:1039
		{
			yyVAL.list = append(yyDollar[1].list, yyDollar[3].expr)
		}
	case 203:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y.go:1045
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 204:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parse.y.go:1054
		{
			yyVAL.node = &ast.HavingClause{Having: yyDollar[2].expr}
		}
	case 205:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parse.y.go:1055
		{
			yyVAL.node = nil
		}
	case 206:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y.go:1065
		{
			yyVAL.node = yyDollar[1].node
		}
	case 207:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parse.y.go:1066
		{
			yyVAL.node = nil
		}
	case 208:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y.go:1069
		{
			yyVAL.node = &ast.SortClause{Target_ref: yyDollar[3].list}
		}
	case 209:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y.go:1072
		{
			yyVAL.list = []ast.ExprNode{yyDollar[1].expr}
		}
	case 210:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y.go:1073
		{
			yyVAL.list = append(yyDollar[1].list, yyDollar[3].expr)
		}
	case 211:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parse.y.go:1076
		{
			yyVAL.expr = &ast.Sortby{Expr: yyDollar[1].expr, AscType: yyDollar[2].tag}
		}
	case 212:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y.go:1079
		{
			yyVAL.tag = ast.Sort_Asc
		}
	case 213:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y.go:1080
		{
			yyVAL.tag = ast.Sort_Desc
		}
	case 214:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parse.y.go:1081
		{
			yyVAL.tag = ast.Sort_Empty
		}
	case 215:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y.go:1090
		{
			yyVAL.node = yyDollar[1].node
		}
	case 216:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parse.y.go:1091
		{
			yyVAL.node = nil
		}
	case 217:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parse.y.go:1095
		{
			lock := &ast.LockClause{}
			lock.SetTag(ast.Lock_update)
			yyVAL.node = lock
		}
	case 218:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parse.y.go:1101
		{
			lock := &ast.LockClause{}
			lock.SetTag(ast.Lock_share)
			yyVAL.node = lock
		}
	case 219:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y.go:1115
		{
			yyVAL.node = yyDollar[1].node
		}
	case 220:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parse.y.go:1116
		{
			yyVAL.node = nil
		}
	case 221:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parse.y.go:1119
		{
			yyVAL.node = &ast.LimitClause{Limit: yyDollar[1].list, Offset: yyDollar[2].expr}
		}
	case 222:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parse.y.go:1120
		{
			yyVAL.node = &ast.LimitClause{Limit: yyDollar[2].list, Offset: yyDollar[1].expr}
		}
	case 223:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y.go:1121
		{
			yyVAL.node = &ast.LimitClause{Limit: yyDollar[1].list}
		}
	case 224:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y.go:1122
		{
			yyVAL.node = &ast.LimitClause{Offset: yyDollar[1].expr}
		}
	case 225:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parse.y.go:1128
		{
			yyVAL.list = yyDollar[2].list
		}
	case 226:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parse.y.go:1132
		{
			yyVAL.list = append(yyDollar[2].list, yyDollar[4].expr)
		}
	case 227:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y.go:1137
		{
			yyVAL.list = []ast.ExprNode{yyDollar[1].expr}
		}
	case 228:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parse.y.go:1140
		{
			yyVAL.expr = yyDollar[2].expr
		}
	case 229:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y.go:1143
		{
			yyVAL.expr = yyDollar[1].expr
		}
	}
	goto yystack /* stack new state and value */
}
