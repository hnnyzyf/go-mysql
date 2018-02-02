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

const IDENT = 57346
const BuiltinCharacterIdent = 57347
const BuiltinFucTimeAddIdent = 57348
const BuiltinFucTimeSubIdent = 57349
const BuiltinTimeUnitIdent = 57350
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
const CHARACTER = 57364
const COMMENT = 57365
const COUNT = 57366
const DESC = 57367
const DISTINCT = 57368
const DISTINCTROW = 57369
const FALSE = 57370
const FOR = 57371
const FROM = 57372
const GROUP = 57373
const HAVING = 57374
const IF = 57375
const INTO = 57376
const LIMIT = 57377
const LOCK = 57378
const MAX = 57379
const MIN = 57380
const MODE = 57381
const NULL = 57382
const OFFSET = 57383
const ORDER = 57384
const QUARTER = 57385
const SELECT = 57386
const SET = 57387
const SHARE = 57388
const SOME = 57389
const SOUNDS = 57390
const SUM = 57391
const TO = 57392
const TRUE = 57393
const UPDATE = 57394
const WHERE = 57395
const CAST = 57396
const CONVERT = 57397
const LOW = 57398
const INTERVAL = 57399
const UNION = 57400
const EXCEPT = 57401
const INTERSECT = 57402
const JOIN = 57403
const STRAIGHT_JOIN = 57404
const LEFT = 57405
const RIGHT = 57406
const INNER = 57407
const OUTER = 57408
const CROSS = 57409
const NATURAL = 57410
const USE = 57411
const ON = 57412
const USING = 57413
const OR = 57414
const OO = 57415
const AND = 57416
const AA = 57417
const NOT = 57418
const BETWEEN = 57419
const CASE = 57420
const WHEN = 57421
const THEN = 57422
const ELSE = 57423
const IS = 57424
const LIKE = 57425
const REGEXP = 57426
const IN = 57427
const GE = 57428
const LE = 57429
const NE = 57430
const LG = 57431
const IE = 57432
const LNE = 57433
const SL = 57434
const SR = 57435
const LEG = 57436
const XOR = 57437
const EXISTS = 57438
const DIV = 57439
const MOD = 57440
const OP = 57441
const COLLATE = 57442
const UMINUS = 57443
const END = 57444

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"IDENT",
	"BuiltinCharacterIdent",
	"BuiltinFucTimeAddIdent",
	"BuiltinFucTimeSubIdent",
	"BuiltinTimeUnitIdent",
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
	"CHARACTER",
	"COMMENT",
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
	"CAST",
	"CONVERT",
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
	116, 6,
	-2, 10,
	-1, 9,
	1, 7,
	116, 7,
	-2, 205,
	-1, 28,
	78, 144,
	87, 144,
	88, 144,
	89, 144,
	-2, 36,
	-1, 52,
	115, 119,
	-2, 76,
	-1, 146,
	29, 11,
	35, 11,
	36, 11,
	41, 11,
	42, 11,
	58, 11,
	59, 11,
	60, 11,
	-2, 64,
	-1, 270,
	117, 175,
	-2, 161,
	-1, 271,
	117, 176,
	-2, 162,
}

const yyPrivate = 57344

const yyLast = 1460

var yyAct = [...]int{

	24, 265, 375, 348, 40, 4, 263, 373, 305, 27,
	327, 52, 16, 354, 262, 268, 106, 311, 333, 108,
	51, 105, 259, 240, 235, 183, 113, 337, 185, 175,
	29, 399, 90, 323, 65, 65, 109, 284, 88, 190,
	147, 20, 271, 282, 146, 65, 148, 28, 150, 149,
	190, 91, 92, 271, 392, 391, 65, 100, 101, 103,
	104, 111, 155, 112, 110, 367, 142, 143, 144, 145,
	107, 366, 100, 101, 103, 104, 360, 284, 156, 365,
	296, 126, 95, 102, 8, 8, 398, 174, 176, 284,
	94, 140, 363, 179, 184, 387, 179, 179, 102, 49,
	345, 193, 194, 195, 196, 197, 275, 7, 164, 188,
	139, 163, 162, 187, 138, 344, 161, 160, 159, 109,
	100, 101, 103, 104, 158, 201, 202, 198, 153, 129,
	130, 359, 341, 152, 127, 128, 131, 132, 133, 134,
	137, 135, 136, 138, 283, 267, 102, 100, 101, 103,
	104, 157, 236, 237, 238, 7, 267, 151, 21, 102,
	246, 232, 233, 343, 251, 252, 65, 96, 95, 242,
	231, 328, 229, 102, 170, 213, 214, 215, 216, 217,
	219, 221, 222, 223, 224, 225, 226, 212, 241, 207,
	342, 266, 244, 141, 378, 355, 326, 243, 270, 248,
	249, 250, 133, 134, 137, 135, 136, 138, 273, 100,
	101, 103, 104, 274, 199, 84, 390, 129, 130, 280,
	272, 281, 127, 128, 131, 132, 133, 134, 137, 135,
	136, 138, 140, 85, 385, 102, 100, 101, 103, 104,
	379, 277, 291, 80, 292, 236, 294, 103, 104, 297,
	298, 299, 340, 316, 317, 176, 184, 276, 192, 278,
	279, 307, 102, 289, 209, 81, 191, 64, 241, 293,
	290, 102, 320, 210, 211, 208, 236, 109, 109, 270,
	146, 304, 322, 319, 303, 332, 358, 318, 357, 308,
	255, 338, 329, 254, 331, 324, 12, 334, 334, 109,
	325, 172, 99, 330, 97, 335, 347, 336, 100, 101,
	103, 104, 346, 266, 266, 266, 339, 353, 171, 260,
	270, 270, 270, 351, 352, 168, 100, 101, 103, 104,
	356, 99, 109, 362, 102, 188, 364, 319, 169, 361,
	368, 253, 100, 101, 103, 104, 166, 369, 370, 87,
	371, 374, 102, 127, 128, 131, 132, 133, 134, 137,
	135, 136, 138, 302, 82, 266, 382, 86, 102, 295,
	380, 83, 270, 87, 86, 381, 82, 386, 97, 388,
	383, 384, 86, 83, 266, 230, 374, 14, 87, 395,
	349, 270, 396, 393, 389, 266, 306, 397, 186, 270,
	350, 400, 270, 11, 13, 12, 93, 65, 65, 78,
	79, 270, 48, 60, 47, 46, 43, 44, 45, 245,
	228, 15, 3, 68, 1, 19, 239, 69, 5, 22,
	23, 63, 131, 132, 133, 134, 137, 135, 136, 138,
	70, 71, 321, 50, 98, 89, 182, 301, 173, 394,
	17, 372, 72, 261, 62, 22, 23, 58, 59, 100,
	101, 103, 104, 129, 130, 32, 57, 75, 127, 128,
	131, 132, 133, 134, 137, 135, 136, 138, 34, 63,
	25, 30, 61, 264, 6, 102, 77, 100, 101, 103,
	104, 200, 76, 128, 131, 132, 133, 134, 137, 135,
	136, 138, 62, 41, 65, 33, 35, 36, 247, 111,
	288, 112, 110, 102, 37, 31, 42, 178, 39, 154,
	178, 178, 38, 66, 65, 26, 78, 79, 56, 48,
	60, 47, 46, 43, 44, 45, 206, 205, 65, 53,
	68, 67, 65, 111, 69, 112, 110, 111, 63, 112,
	110, 286, 107, 74, 73, 55, 107, 70, 71, 54,
	50, 285, 287, 100, 101, 103, 104, 204, 203, 72,
	125, 62, 114, 177, 58, 59, 180, 181, 117, 118,
	122, 120, 309, 269, 75, 119, 121, 116, 123, 102,
	124, 227, 300, 115, 313, 310, 316, 317, 314, 61,
	315, 312, 256, 77, 100, 101, 103, 104, 9, 76,
	2, 100, 101, 103, 104, 257, 100, 101, 103, 104,
	41, 258, 10, 35, 36, 64, 165, 167, 65, 18,
	102, 37, 0, 190, 0, 39, 271, 102, 0, 38,
	66, 65, 102, 78, 79, 0, 48, 60, 47, 46,
	43, 44, 45, 0, 65, 0, 0, 68, 0, 190,
	0, 69, 189, 22, 23, 63, 0, 0, 0, 100,
	101, 103, 104, 0, 70, 71, 0, 50, 0, 0,
	0, 0, 0, 0, 0, 0, 72, 0, 62, 0,
	0, 58, 59, 0, 0, 102, 313, 310, 316, 317,
	314, 75, 315, 312, 0, 376, 377, 0, 0, 0,
	0, 0, 0, 0, 25, 0, 61, 0, 0, 0,
	77, 0, 0, 0, 0, 0, 76, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 41, 0, 0,
	35, 36, 64, 0, 0, 0, 0, 0, 37, 0,
	0, 0, 39, 0, 0, 0, 38, 66, 65, 0,
	78, 79, 0, 48, 60, 47, 46, 43, 44, 45,
	0, 0, 0, 0, 68, 0, 0, 0, 69, 0,
	0, 0, 63, 0, 0, 0, 0, 0, 0, 0,
	0, 70, 71, 0, 50, 0, 0, 0, 8, 0,
	0, 0, 0, 72, 0, 62, 0, 0, 58, 59,
	0, 0, 0, 0, 0, 0, 0, 0, 75, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 25, 0, 61, 0, 0, 0, 77, 0, 0,
	0, 0, 0, 76, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 41, 0, 0, 35, 36, 64,
	0, 0, 0, 0, 0, 37, 0, 0, 0, 39,
	0, 0, 0, 38, 66, 65, 0, 78, 79, 0,
	48, 60, 47, 46, 43, 44, 45, 0, 0, 0,
	0, 68, 0, 0, 0, 69, 0, 0, 0, 63,
	0, 0, 0, 0, 0, 0, 0, 0, 70, 71,
	0, 50, 0, 0, 0, 0, 0, 0, 0, 0,
	72, 0, 62, 0, 0, 58, 59, 0, 0, 0,
	0, 0, 0, 0, 0, 75, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 25, 0,
	61, 0, 0, 0, 77, 0, 0, 0, 0, 0,
	76, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 41, 0, 0, 35, 36, 64, 0, 0, 0,
	0, 0, 37, 0, 0, 0, 39, 234, 0, 0,
	38, 66, 65, 0, 78, 79, 0, 48, 60, 47,
	46, 43, 44, 45, 0, 0, 0, 0, 68, 0,
	0, 0, 69, 0, 0, 0, 63, 0, 0, 0,
	0, 0, 0, 0, 0, 70, 71, 0, 50, 0,
	0, 0, 0, 0, 0, 0, 0, 72, 0, 62,
	0, 0, 58, 59, 0, 0, 0, 0, 0, 0,
	0, 0, 75, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 25, 0, 61, 0, 0,
	0, 77, 0, 0, 0, 0, 0, 76, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 41, 0,
	0, 35, 36, 64, 0, 0, 0, 0, 0, 37,
	0, 0, 0, 39, 0, 0, 0, 38, 66, 65,
	0, 78, 79, 0, 48, 60, 47, 46, 43, 44,
	45, 0, 0, 0, 0, 68, 0, 0, 0, 69,
	0, 0, 0, 63, 0, 0, 0, 0, 0, 0,
	0, 0, 70, 71, 0, 50, 0, 0, 0, 0,
	0, 0, 0, 0, 72, 0, 62, 0, 0, 58,
	59, 0, 220, 0, 0, 0, 0, 0, 0, 75,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 61, 0, 0, 0, 77, 0,
	0, 0, 0, 0, 76, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 41, 0, 0, 35, 36,
	64, 0, 0, 0, 0, 0, 37, 0, 0, 0,
	39, 0, 0, 0, 38, 66, 65, 0, 78, 79,
	0, 48, 60, 47, 46, 43, 44, 45, 0, 0,
	0, 0, 68, 0, 0, 0, 69, 0, 0, 0,
	63, 0, 0, 0, 0, 0, 0, 0, 0, 70,
	71, 0, 50, 0, 0, 0, 0, 0, 0, 0,
	0, 72, 0, 62, 0, 0, 58, 59, 0, 218,
	0, 0, 0, 0, 0, 0, 75, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 61, 0, 0, 0, 77, 0, 0, 0, 0,
	0, 76, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 41, 0, 0, 35, 36, 64, 0, 0,
	0, 0, 0, 37, 0, 0, 0, 39, 0, 0,
	0, 38, 66, 65, 0, 78, 79, 0, 48, 60,
	47, 46, 43, 44, 45, 0, 0, 0, 0, 68,
	0, 0, 0, 69, 0, 0, 0, 63, 0, 0,
	0, 0, 0, 0, 0, 0, 70, 71, 0, 50,
	0, 0, 0, 0, 0, 0, 0, 0, 72, 0,
	62, 0, 0, 58, 59, 0, 0, 0, 0, 0,
	0, 0, 0, 75, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 61, 0,
	0, 0, 77, 0, 0, 0, 0, 0, 76, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 41,
	0, 0, 35, 36, 64, 0, 0, 0, 0, 0,
	37, 0, 0, 0, 39, 0, 0, 0, 38, 66,
}
var yyPact = [...]int{

	40, -1000, -1000, -1000, -1000, -1000, 345, 40, 637, -1000,
	347, 429, 429, 429, 385, -26, -34, 344, 988, 241,
	-1000, -1000, -1000, -1000, 538, 988, 495, -1000, 33, 119,
	-1000, -1000, -1000, -1000, -1000, 1339, 1339, 1339, 1339, 754,
	-1000, -8, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -68, -1000, -69, 42, -1000, -1000, 18, 13,
	-1000, 988, -1000, -1000, -1000, -1000, 31, -1000, 9, 3,
	2, 1, -3, -4, -7, -1000, -1000, -1000, -1000, -1000,
	332, 335, 286, 85, 308, 339, 988, 988, 40, -1000,
	-1000, 40, 40, 988, -1000, -1000, 368, 650, 270, 988,
	988, 988, 988, 988, 988, -1000, -1000, 500, -1000, -1000,
	-1000, -1000, -1000, 60, 451, 1339, 520, 112, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, 186, 100, 1339, 1339, 1339,
	1339, 1222, 1105, 1339, 1339, 1339, 1339, 1339, 1339, -1000,
	415, 1339, -22, -22, -22, -22, -34, 269, -1000, 162,
	404, 871, 988, 988, 108, 414, -1000, 404, 429, 403,
	429, 429, 429, 988, 988, -1000, -1000, -1000, -1000, -1000,
	295, -1000, -1000, 232, 414, -1000, 414, 236, -1000, -1000,
	-1000, 236, 229, -1000, 596, 266, 30, -1000, -1000, -1000,
	-1000, 368, -1000, 172, 172, -1000, 60, 60, -1000, -1000,
	-1000, -1000, -1000, -8, -1000, -1000, -1000, -1000, -9, 1339,
	1339, 1339, 1339, 391, 329, 252, 252, 97, 988, 97,
	988, 4, 4, 4, 4, 4, -1000, -1000, -1000, -22,
	-1000, -1000, -74, -1000, -1000, 28, 414, 543, 490, 188,
	-1000, 988, -1000, 988, 988, 988, 253, -36, 988, 988,
	988, 531, 386, 324, 988, 988, -1000, -1000, -1000, 365,
	988, 228, 532, -1000, -1000, 534, 534, 41, -1000, -84,
	-1000, -1000, 266, -1000, -1000, 754, 121, -22, 367, 367,
	163, 163, 162, -1000, 988, 404, 404, 500, -91, -1000,
	988, 235, 136, 16, 74, -1000, -1000, 47, -1, -16,
	255, 249, -1000, -1000, -1000, 358, 379, 414, 30, 30,
	30, 128, 189, -1000, 226, 224, -1000, -1000, -1000, -1000,
	52, 15, 532, 650, 365, -24, 1339, -1000, -1000, -1000,
	-1000, -1000, 414, -37, -1000, -45, -51, -1000, 414, 988,
	-1000, -1000, -1000, -1000, -1000, -1000, 988, 988, -1000, 988,
	988, 532, 634, 123, 178, -1000, 128, -1000, -1000, -1000,
	30, -1000, 358, -1000, -1000, -1000, -1000, -1000, 414, 163,
	163, 414, 173, -1000, 414, -1000, 988, -20, 988, 30,
	154, 532, -1000, -61, -62, 988, 414, 624, 414, 634,
	30, -1000, -1000, -1000, -30, -1000, -1000, -1000, -1000, 624,
	-1000,
}
var yyPgo = [...]int{

	0, 629, 243, 28, 8, 3, 167, 627, 626, 622,
	484, 265, 421, 610, 4, 428, 608, 22, 16, 19,
	38, 602, 591, 587, 583, 11, 582, 572, 17, 21,
	32, 570, 13, 568, 15, 10, 99, 18, 559, 555,
	554, 553, 20, 541, 539, 528, 47, 525, 519, 516,
	515, 510, 0, 505, 7, 2, 483, 481, 478, 233,
	9, 1, 29, 30, 25, 6, 14, 158, 466, 465,
	23, 24, 453, 451, 450, 449, 448, 446, 442, 425,
	426, 215, 424,
}
var yyR1 = [...]int{

	0, 82, 13, 13, 14, 14, 12, 12, 12, 12,
	10, 10, 15, 15, 15, 15, 15, 52, 52, 52,
	52, 52, 52, 52, 52, 47, 47, 47, 47, 47,
	60, 60, 60, 60, 60, 60, 60, 46, 46, 46,
	46, 46, 46, 46, 46, 46, 46, 46, 46, 46,
	46, 46, 63, 63, 63, 63, 63, 63, 63, 63,
	63, 63, 63, 63, 63, 63, 63, 57, 57, 57,
	57, 57, 57, 57, 57, 25, 42, 42, 69, 69,
	69, 50, 50, 50, 53, 53, 53, 53, 53, 53,
	53, 45, 45, 45, 45, 45, 45, 45, 45, 68,
	68, 71, 71, 58, 49, 48, 48, 80, 80, 70,
	51, 51, 74, 74, 79, 79, 67, 44, 38, 39,
	39, 29, 29, 18, 18, 19, 19, 19, 19, 36,
	36, 33, 33, 33, 23, 23, 23, 23, 23, 23,
	23, 27, 27, 31, 31, 35, 22, 43, 43, 43,
	37, 40, 41, 1, 20, 20, 20, 30, 30, 6,
	6, 34, 34, 34, 3, 3, 72, 72, 66, 66,
	65, 65, 65, 61, 61, 24, 24, 78, 78, 56,
	56, 56, 56, 56, 56, 26, 26, 26, 28, 28,
	32, 32, 55, 55, 75, 75, 17, 17, 4, 4,
	73, 73, 54, 5, 5, 9, 9, 16, 77, 77,
	64, 21, 21, 21, 7, 7, 2, 2, 8, 8,
	11, 11, 11, 11, 81, 81, 76, 62, 59,
}
var yyR2 = [...]int{

	0, 1, 1, 1, 3, 3, 1, 2, 4, 4,
	1, 1, 7, 8, 4, 4, 4, 3, 3, 3,
	3, 3, 2, 3, 1, 3, 3, 3, 4, 1,
	4, 6, 6, 4, 4, 4, 1, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 5,
	5, 1, 1, 1, 1, 1, 3, 1, 3, 2,
	2, 2, 2, 3, 1, 2, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	3, 1, 3, 5, 3, 4, 1, 1, 6, 6,
	6, 5, 5, 5, 4, 4, 5, 5, 5, 8,
	8, 1, 3, 1, 5, 1, 0, 1, 2, 4,
	2, 0, 1, 0, 1, 3, 2, 2, 3, 1,
	1, 1, 0, 2, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 2, 1, 0, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 0, 1, 1, 2,
	0, 1, 1, 1, 2, 0, 1, 3, 1, 1,
	2, 2, 3, 1, 3, 1, 1, 1, 3, 3,
	4, 3, 5, 6, 6, 1, 2, 2, 1, 1,
	1, 0, 2, 4, 1, 3, 2, 0, 3, 0,
	1, 3, 1, 2, 0, 1, 0, 3, 1, 3,
	2, 1, 1, 0, 1, 0, 2, 4, 1, 0,
	2, 2, 1, 1, 2, 4, 1, 1, 2,
}
var yyChk = [...]int{

	-1000, -82, -13, -12, -14, -15, -10, 115, 44, -16,
	-9, 58, 60, 59, 42, -12, -14, -74, -1, -79,
	-30, -67, 26, 27, -52, 77, -47, -60, -46, -63,
	-57, -50, -69, -53, -58, 103, 104, 111, 119, 115,
	-14, 100, -49, 13, 14, 15, 12, 11, 9, -36,
	40, -42, -25, -44, -38, -39, -45, -68, 54, 55,
	10, 79, 51, 28, 105, 4, 120, -43, 20, 24,
	37, 38, 49, -40, -41, 64, 89, 83, 6, 7,
	-2, -11, 29, 36, -81, -59, 35, 41, -20, 16,
	-30, -20, -20, 21, 116, 116, -6, 34, -79, 61,
	73, 74, 99, 75, 76, -29, -18, 18, -19, -25,
	12, 9, 11, -52, -27, 98, -23, 83, 84, 90,
	86, 91, 85, 93, 95, -31, 48, 101, 102, 96,
	97, 103, 104, 105, 106, 108, 109, 107, 110, 77,
	113, 74, -63, -63, -63, -63, -14, -52, -14, 117,
	117, 115, 115, 115, -48, -52, -25, 120, 115, 115,
	115, 115, 115, 115, 115, -8, -11, -7, -2, 52,
	89, -59, -81, -76, -52, -62, -52, -10, -15, -14,
	-10, -10, -77, -64, -52, -3, 30, -34, -25, 12,
	9, -6, -67, -52, -52, -52, -52, -52, -19, -36,
	40, -60, -60, -33, 47, 17, 16, 77, 89, 78,
	87, 88, 87, -46, -46, -46, -46, -46, 57, -46,
	57, -46, -46, -46, -46, -46, -46, -22, 5, -63,
	116, -42, -25, -25, 116, -71, -52, -52, -52, -80,
	-70, 80, -25, -20, -30, 16, -52, 105, -20, -20,
	-20, -52, -52, 46, 61, 61, -21, 19, 25, -17,
	53, -72, -66, -65, -56, -61, -14, 115, -34, -24,
	-25, 12, -3, -14, -14, 115, -46, -63, -46, -46,
	-52, -52, 117, 116, 61, 18, 61, 72, -51, -70,
	82, -52, -52, -71, -52, 116, 116, -52, -52, -52,
	61, 61, 39, -62, -64, -4, 31, -52, 61, -26,
	63, -28, 69, 62, 66, 68, 64, 65, -29, -18,
	-14, -78, -66, 117, -17, -71, 75, -35, 8, -35,
	-42, -25, -52, -37, -25, -37, -19, 118, -52, 81,
	116, 116, 116, 116, 116, 116, 57, 57, -5, 32,
	21, -66, -66, -65, -32, 67, -28, 62, 62, 116,
	61, -34, -4, 116, -60, 116, 116, 116, -52, -52,
	-52, -52, -73, -54, -52, -55, 71, 72, 71, 62,
	-32, -66, -5, -35, -35, 61, -52, 115, -52, -66,
	62, 116, 116, -54, -75, -61, -55, -65, 116, 61,
	-61,
}
var yyDef = [...]int{

	0, -2, 1, 2, -2, -2, 206, 0, 113, -2,
	0, 156, 156, 156, 0, 0, 11, 160, 0, 112,
	153, 114, 157, 158, 122, 0, 24, 29, -2, 51,
	52, 53, 54, 55, 57, 0, 0, 0, 0, 0,
	64, 0, 66, 67, 68, 69, 70, 71, 72, 73,
	74, 81, -2, 78, 79, 0, 86, 87, 0, 0,
	103, 106, 129, 130, 77, 75, 0, 120, 0, 0,
	0, 0, 0, 0, 0, 147, 148, 149, 151, 152,
	219, 215, 0, 0, 222, 223, 0, 0, 0, 154,
	155, 0, 0, 0, 4, 5, 165, 0, 160, 0,
	0, 0, 0, 0, 0, 116, 121, 0, 124, 125,
	126, 127, 128, 22, 0, 0, 0, 141, 134, 135,
	136, 137, 138, 139, 140, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 143,
	0, 0, 59, 60, 61, 62, -2, 0, 65, 0,
	0, 0, 0, 0, 0, 105, 117, 0, 156, 0,
	156, 156, 156, 0, 0, 8, 218, 9, 214, 216,
	0, 220, 221, 224, 226, 228, 227, 14, 10, 11,
	15, 16, 207, 208, 213, 197, 0, 159, 161, 162,
	163, 165, 115, 17, 18, 19, 20, 21, 123, 23,
	25, 26, 27, 0, 131, 132, 133, 142, 0, 0,
	0, 0, 0, 37, 38, 39, 40, 41, 0, 42,
	0, 43, 44, 45, 46, 47, 48, 56, 146, 58,
	63, 82, 76, 80, 84, 0, 101, 0, 0, 111,
	107, 0, 118, 0, 0, 0, 0, 77, 0, 0,
	0, 0, 0, 0, 0, 0, 210, 211, 212, 199,
	0, 164, 166, 168, 169, 122, 0, 0, 173, 0,
	-2, -2, 197, 28, 30, 0, 0, 34, 35, 33,
	0, 0, 0, 85, 0, 0, 0, 0, 0, 108,
	0, 0, 0, 0, 0, 94, 95, 0, 0, 0,
	0, 0, 217, 225, 209, 204, 0, 196, 0, 0,
	0, 191, 0, 185, 0, 0, 188, 189, 170, 171,
	11, 0, 177, 0, 199, 0, 0, 49, 145, 50,
	83, 76, 102, 0, 150, 0, 0, 104, 110, 0,
	91, 92, 93, 96, 97, 98, 0, 0, 12, 0,
	0, 167, 179, 181, 0, 190, 191, 186, 187, 172,
	0, 174, 204, 31, 32, 88, 89, 90, 109, 0,
	0, 203, 198, 200, 202, 180, 0, 0, 0, 0,
	0, 178, 13, 0, 0, 0, 192, 0, 182, 0,
	0, 99, 100, 201, 0, 194, 183, 184, 193, 0,
	195,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 119, 3, 3, 3, 107, 102, 3,
	115, 116, 105, 103, 61, 104, 117, 106, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	85, 84, 86, 3, 120, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 110, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 101, 3, 111,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
	52, 53, 54, 55, 56, 57, 58, 59, 60, 62,
	63, 64, 65, 66, 67, 68, 69, 70, 71, 72,
	73, 74, 75, 76, 77, 78, 79, 80, 81, 82,
	83, 87, 88, 89, 90, 91, 92, 93, 94, 95,
	96, 97, 98, 99, 100, 108, 109, 112, 113, 114,
	118,
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
		//line parser.y:213
		{
			yylex.(*Tokener).ParseStmt = yyDollar[1].node
		}
	case 4:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:232
		{
			yyVAL.node = yyDollar[2].node
		}
	case 5:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:236
		{
			yyVAL.node = yyDollar[2].node
		}
	case 6:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:242
		{
			yyVAL.node = yyDollar[1].node
		}
	case 7:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:246
		{

		}
	case 8:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:250
		{
		}
	case 9:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:253
		{
		}
	case 10:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:257
		{
			yyVAL.node = yyDollar[1].node
		}
	case 11:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:258
		{
			yyVAL.node = yyDollar[1].node
		}
	case 12:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.y:263
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
		//line parser.y:285
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
		//line parser.y:306
		{
			stmt := &ast.SelectStmt{All: yyDollar[3].val, Left: yyDollar[1].node.(*ast.SelectStmt), Right: yyDollar[4].node.(*ast.SelectStmt)}
			stmt.SetTag(ast.AST_SELECT_UNION)
			yyVAL.node = stmt
		}
	case 15:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:312
		{
			stmt := &ast.SelectStmt{All: yyDollar[3].val, Left: yyDollar[1].node.(*ast.SelectStmt), Right: yyDollar[4].node.(*ast.SelectStmt)}
			stmt.SetTag(ast.AST_SELECT_INTERSECT)
			yyVAL.node = stmt
		}
	case 16:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:318
		{
			stmt := &ast.SelectStmt{All: yyDollar[3].val, Left: yyDollar[1].node.(*ast.SelectStmt), Right: yyDollar[4].node.(*ast.SelectStmt)}
			stmt.SetTag(ast.AST_SELECT_EXCEPT)
			yyVAL.node = stmt
		}
	case 17:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:334
		{
			expr := &ast.Expr{Operator: "OR", Left: yyDollar[1].expr, Right: yyDollar[3].expr}
			expr.SetTag(ast.AST_EXPR_OR)
			yyVAL.expr = expr
		}
	case 18:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:340
		{
			expr := &ast.Expr{Operator: "||", Left: yyDollar[1].expr, Right: yyDollar[3].expr}
			expr.SetTag(ast.AST_EXPR_DOUBLE_OR)
			yyVAL.expr = expr
		}
	case 19:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:346
		{
			expr := &ast.Expr{Operator: "XOR", Left: yyDollar[1].expr, Right: yyDollar[3].expr}
			expr.SetTag(ast.AST_EXPR_XOR)
			yyVAL.expr = expr
		}
	case 20:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:352
		{
			expr := &ast.Expr{Operator: "AND", Left: yyDollar[1].expr, Right: yyDollar[3].expr}
			expr.SetTag(ast.AST_EXPR_AND)
			yyVAL.expr = expr
		}
	case 21:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:358
		{
			expr := &ast.Expr{Operator: "&&", Left: yyDollar[1].expr, Right: yyDollar[3].expr}
			expr.SetTag(ast.AST_EXPR_DOUBLE_AND)
			yyVAL.expr = expr
		}
	case 22:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:364
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
		//line parser.y:377
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
		//line parser.y:396
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 25:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:402
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
		//line parser.y:414
		{
			expr := &ast.Expr{Operator: "<=>", Left: yyDollar[1].expr, Right: yyDollar[3].expr}
			expr.SetTag(ast.AST_EXPR_LEG)
			yyVAL.expr = expr
		}
	case 27:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:420
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
		//line parser.y:441
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
		//line parser.y:454
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 30:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:460
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
		//line parser.y:471
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
		//line parser.y:482
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
		//line parser.y:494
		{
			expr := &ast.Expr{Operator: "LIKE", Left: yyDollar[1].expr, Right: yyDollar[4].expr}
			expr.SetTag(ast.AST_EXPR_SOUND_LIKE)
			yyVAL.expr = expr
		}
	case 34:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:500
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
		//line parser.y:512
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
		//line parser.y:524
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 37:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:531
		{
			expr := &ast.Expr{Operator: "|", Left: yyDollar[1].expr, Right: yyDollar[3].expr}
			expr.SetTag(ast.AST_EXPR_VETICEL)
			yyVAL.expr = expr
		}
	case 38:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:537
		{
			expr := &ast.Expr{Operator: "&", Left: yyDollar[1].expr, Right: yyDollar[3].expr}
			expr.SetTag(ast.AST_EXPR_AMPERSAND)
			yyVAL.expr = expr
		}
	case 39:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:543
		{
			expr := &ast.Expr{Operator: "<<", Left: yyDollar[1].expr, Right: yyDollar[3].expr}
			expr.SetTag(ast.AST_EXPR_SL)
			yyVAL.expr = expr
		}
	case 40:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:549
		{
			expr := &ast.Expr{Operator: ">>", Left: yyDollar[1].expr, Right: yyDollar[3].expr}
			expr.SetTag(ast.AST_EXPR_SR)
			yyVAL.expr = expr
		}
	case 41:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:555
		{
			expr := &ast.Expr{Operator: "+", Left: yyDollar[1].expr, Right: yyDollar[3].expr}
			expr.SetTag(ast.AST_EXPR_PLUS)
			yyVAL.expr = expr
		}
	case 42:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:561
		{
			expr := &ast.Expr{Operator: "-", Left: yyDollar[1].expr, Right: yyDollar[3].expr}
			expr.SetTag(ast.AST_EXPR_MINUS)
			yyVAL.expr = expr
		}
	case 43:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:567
		{
			expr := &ast.Expr{Operator: "*", Left: yyDollar[1].expr, Right: yyDollar[3].expr}
			expr.SetTag(ast.AST_EXPR_MULT)
			yyVAL.expr = expr
		}
	case 44:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:573
		{
			expr := &ast.Expr{Operator: "/", Left: yyDollar[1].expr, Right: yyDollar[3].expr}
			expr.SetTag(ast.AST_EXPR_DIV)
			yyVAL.expr = expr
		}
	case 45:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:579
		{
			expr := &ast.Expr{Operator: "DIV", Left: yyDollar[1].expr, Right: yyDollar[3].expr}
			expr.SetTag(ast.AST_EXPR_DIV)
			yyVAL.expr = expr
		}
	case 46:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:585
		{
			expr := &ast.Expr{Operator: "MOD", Left: yyDollar[1].expr, Right: yyDollar[3].expr}
			expr.SetTag(ast.AST_EXPR_MOD)
			yyVAL.expr = expr
		}
	case 47:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:591
		{
			expr := &ast.Expr{Operator: "%", Left: yyDollar[1].expr, Right: yyDollar[3].expr}
			expr.SetTag(ast.AST_EXPR_MOD)
			yyVAL.expr = expr
		}
	case 48:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:597
		{
			expr := &ast.Expr{Operator: "^", Left: yyDollar[1].expr, Right: yyDollar[3].expr}
			expr.SetTag(ast.AST_EXPR_XOR)
			yyVAL.expr = expr
		}
	case 49:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:603
		{
			expr := &ast.Expr{Operator: "+", Left: yyDollar[1].expr, Right: yyDollar[4].expr}
			expr.SetTag(ast.AST_TIME_INTERVAL)
			expr.SetAlias(yyDollar[5].val)
			yyVAL.expr = expr
		}
	case 50:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:610
		{
			expr := &ast.Expr{Operator: "-", Left: yyDollar[1].expr, Right: yyDollar[4].expr}
			expr.SetTag(ast.AST_TIME_INTERVAL)
			expr.SetAlias(yyDollar[5].val)
			yyVAL.expr = expr
		}
	case 51:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:617
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 52:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:625
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 53:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:630
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 54:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:635
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 55:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:639
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 56:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:643
		{
			yyDollar[1].expr.SetCollate(yyDollar[3].val)
			yyVAL.expr = yyDollar[1].expr
		}
	case 57:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:648
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 58:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:653
		{
			expr := &ast.Expr{Operator: "||", Left: yyDollar[1].expr, Right: yyDollar[3].expr}
			expr.SetTag(ast.AST_EXPR_DOUBLE_OR)
			yyVAL.expr = expr
		}
	case 59:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:659
		{
			yyVAL.expr = yyDollar[2].expr.(*ast.ValExpr)
		}
	case 60:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:663
		{
			val := yyDollar[2].expr.(*ast.ValExpr)
			val.Symbol = 0 - val.Symbol
			yyVAL.expr = val
		}
	case 61:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:669
		{
			expr := &ast.Expr{Operator: "~", Left: nil, Right: yyDollar[2].expr}
			expr.SetTag(ast.AST_EXPR_TILDE)
			yyVAL.expr = expr
		}
	case 62:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:675
		{
			expr := &ast.Expr{Operator: "!", Left: nil, Right: yyDollar[2].expr}
			expr.SetTag(ast.AST_EXPR_EXCLAMATION_MARK)
			yyVAL.expr = expr
		}
	case 63:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:681
		{
			yyVAL.expr = yyDollar[2].expr
		}
	case 64:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:686
		{
			expr := &ast.SubqueryExpr{Left: nil, Right: yyDollar[1].node.(*ast.SelectStmt)}
			expr.SetTag(ast.AST_SUBQUERY)
			yyVAL.expr = expr
		}
	case 65:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:692
		{
			expr := &ast.SubqueryExpr{Operator: "EXISTS", Left: nil, Right: yyDollar[2].node.(*ast.SelectStmt)}
			expr.SetTag(ast.AST_SUBQUERY_EXISTS)
			yyVAL.expr = expr
		}
	case 66:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:698
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 67:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:704
		{
			expr := &ast.ValExpr{Symbol: 1, Sval: yyDollar[1].val}
			expr.SetTag(ast.AST_VALUE_BIT_NUMBER)
			yyVAL.expr = expr
		}
	case 68:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:710
		{
			expr := &ast.ValExpr{Symbol: 1, Sval: yyDollar[1].val}
			expr.SetTag(ast.AST_VALUE_HEX_NUMBER)
			yyVAL.expr = expr
		}
	case 69:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:716
		{
			expr := &ast.ValExpr{Symbol: 1, Sval: yyDollar[1].val}
			expr.SetTag(ast.AST_VALUE_NUMBER)
			yyVAL.expr = expr
		}
	case 70:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:722
		{
			expr := &ast.ValExpr{Symbol: 0, Sval: yyDollar[1].str}
			expr.SetTag(ast.AST_VALUE_STRING)
			yyVAL.expr = expr
		}
	case 71:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:728
		{
			expr := &ast.ValExpr{Symbol: 0, Sval: yyDollar[1].str}
			expr.SetTag(ast.AST_VALUE_SINGLE_QUOTA_STRING)
			yyVAL.expr = expr
		}
	case 72:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:734
		{
			expr := &ast.ValExpr{Symbol: 0, Sval: yyDollar[1].str}
			expr.SetTag(ast.AST_VALUE_DOUBLE_QUOTA_STRING)
			yyVAL.expr = expr
		}
	case 73:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:740
		{
			expr := &ast.ValExpr{Symbol: 0, Sval: yyDollar[1].val}
			switch yyDollar[1].val {
			case "TRUE":
				expr.SetTag(ast.AST_VALUE_TRUE)
			case "FALSE":
				expr.SetTag(ast.AST_VALUE_FALSE)
			}

			yyVAL.expr = expr
		}
	case 74:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:752
		{
			expr := &ast.ValExpr{Symbol: 0, Sval: "NULL"}
			switch yyDollar[1].ident {
			case "TRUE":
				expr.SetTag(ast.AST_VALUE_NULL)
			case "FALSE":
				expr.SetTag(ast.AST_VALUE_NULL)
			}

			yyVAL.expr = expr
		}
	case 75:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:766
		{
			yyVAL.val = yyDollar[1].ident
		}
	case 76:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:772
		{
			yyVAL.val = yyDollar[1].val
		}
	case 77:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:776
		{
			yyVAL.val = "*"
		}
	case 78:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:782
		{
			expr := &ast.ColumnExpr{Column: yyDollar[1].val}
			expr.SetTag(ast.AST_EXPR_SINGLE_AT_COLUMN)
			yyVAL.expr = expr
		}
	case 79:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:788
		{
			expr := &ast.ColumnExpr{Column: yyDollar[1].val}
			expr.SetTag(ast.AST_EXPR_DOUBLE_AT_COLUMN)
			yyVAL.expr = expr
		}
	case 80:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:794
		{
			expr := &ast.ColumnExpr{Column: yyDollar[1].val}
			expr.SetTag(ast.AST_EXPR_DOUBLE_AT_COLUMN)
			yyVAL.expr = expr
		}
	case 81:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:805
		{
			expr := &ast.ColumnExpr{Column: yyDollar[1].val}
			expr.SetTag(ast.AST_EXPR_COLUMN)
			yyVAL.expr = expr
		}
	case 82:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:811
		{
			expr := &ast.ColumnExpr{Table: yyDollar[1].val, Column: yyDollar[2].val}
			expr.SetTag(ast.AST_EXPR_COLUMN)
			yyVAL.expr = expr
		}
	case 83:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:817
		{
			expr := &ast.ColumnExpr{DB: yyDollar[1].val, Table: yyDollar[2].val, Column: yyDollar[3].val}
			expr.SetTag(ast.AST_EXPR_COLUMN)
			yyVAL.expr = expr
		}
	case 84:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:826
		{
			expr := &ast.FuncExpr{Name: yyDollar[1].val}
			expr.SetTag(ast.AST_EXPR_FUNC)
			yyVAL.expr = expr
		}
	case 85:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:832
		{
			expr := &ast.FuncExpr{Name: yyDollar[1].val, Arg: yyDollar[3].list}
			expr.SetTag(ast.AST_EXPR_FUNC)
			yyVAL.expr = expr
		}
	case 86:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:838
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 87:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:842
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 88:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.y:846
		{
			expr := &ast.FuncExpr{Name: "CAST", Arg: ast.List{yyDollar[3].expr}}
			expr.SetTag(ast.AST_EXPR_CAST)
			expr.SetCollate(yyDollar[5].val)
			yyVAL.expr = expr
		}
	case 89:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.y:853
		{
			expr := &ast.FuncExpr{Name: "CONVERT", Arg: ast.List{yyDollar[3].expr}}
			expr.SetTag(ast.AST_EXPR_CONVERT_TYPE)
			expr.SetCollate(yyDollar[5].val)
			yyVAL.expr = expr
		}
	case 90:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.y:860
		{
			expr := &ast.FuncExpr{Name: "CONVERT", Arg: ast.List{yyDollar[3].expr}}
			expr.SetTag(ast.AST_EXPR_CONVERT_ALIAS)
			expr.SetAlias(yyDollar[5].val)
			yyVAL.expr = expr
		}
	case 91:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:869
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
	case 92:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:884
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
	case 93:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:895
		{
			expr := &ast.FuncExpr{Name: "COUNT", Arg: ast.List{yyDollar[4].expr}}
			expr.SetTag(ast.AST_EXPR_COUNT_ALL)
			yyVAL.expr = expr
		}
	case 94:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:901
		{
			expr := &ast.FuncExpr{Name: "COUNT", Arg: ast.List{yyDollar[3].expr}}
			expr.SetTag(ast.AST_EXPR_COUNT)
			yyVAL.expr = expr
		}
	case 95:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:907
		{
			expr := &ast.FuncExpr{Name: "COUNT"}
			expr.SetTag(ast.AST_EXPR_COUNT_STAR)
			yyVAL.expr = expr
		}
	case 96:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:913
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
	case 97:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:928
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
	case 98:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:943
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
	case 99:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line parser.y:960
		{

			expr := &ast.Expr{Operator: "+", Left: yyDollar[3].expr, Right: yyDollar[6].expr}
			expr.SetTag(ast.AST_TIME_INTERVAL)
			expr.SetAlias(yyDollar[7].val)
			yyVAL.expr = &ast.FuncExpr{Name: yyDollar[1].val, Arg: ast.List{expr}}
		}
	case 100:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line parser.y:968
		{

			expr := &ast.Expr{Operator: "-", Left: yyDollar[3].expr, Right: yyDollar[6].expr}
			expr.SetTag(ast.AST_TIME_INTERVAL)
			expr.SetAlias(yyDollar[7].val)
			yyVAL.expr = &ast.FuncExpr{Name: yyDollar[1].val, Arg: ast.List{expr}}
		}
	case 101:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:979
		{
			yyVAL.list = ast.List{yyDollar[1].expr}
		}
	case 102:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:983
		{
			yyVAL.list = append(yyDollar[1].list, yyDollar[3].expr)
		}
	case 103:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:991
		{
			expr := &ast.ValExpr{Sval: yyDollar[1].str}
			expr.SetTag(ast.AST_VALUE_MARKER)
			yyVAL.expr = expr
		}
	case 104:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:1001
		{
			expr := &ast.CaseExpr{Case: yyDollar[2].expr, When: yyDollar[3].list, Else: yyDollar[4].expr}
			expr.SetTag(ast.AST_EXPR_CASE)
			yyVAL.expr = expr
		}
	case 105:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1009
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 106:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:1013
		{
			yyVAL.expr = nil
		}
	case 107:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1020
		{
			yyVAL.list = ast.List{yyDollar[1].expr}
		}
	case 108:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:1024
		{
			yyVAL.list = append(yyDollar[1].list, yyDollar[2].expr)
		}
	case 109:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:1031
		{
			expr := &ast.Expr{Left: yyDollar[2].expr, Right: yyDollar[4].expr}
			expr.SetTag(ast.AST_EXPR_WHEN)
			yyVAL.expr = expr
		}
	case 110:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:1040
		{
			expr := &ast.Expr{Left: yyDollar[2].expr}
			expr.SetTag(ast.AST_EXPR_ELSE)
			yyVAL.expr = expr
		}
	case 111:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:1046
		{
			yyVAL.expr = nil
		}
	case 112:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1054
		{
			yyVAL.list = yyDollar[1].list
		}
	case 113:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:1058
		{
			yyVAL.list = nil
		}
	case 114:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1064
		{
			yyVAL.list = ast.List{yyDollar[1].expr}

		}
	case 115:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:1069
		{
			yyVAL.list = append(yyDollar[1].list, yyDollar[3].expr)
		}
	case 116:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:1076
		{
			yyDollar[1].expr.SetAlias(yyDollar[2].val)
			yyVAL.expr = yyDollar[1].expr
		}
	case 117:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:1083
		{
			yyVAL.val = yyDollar[2].val
		}
	case 118:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:1088
		{
			yyVAL.val = yyDollar[3].val
		}
	case 119:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1094
		{
			yyVAL.val = yyDollar[1].val
		}
	case 120:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1098
		{
			yyVAL.val = yyDollar[1].val
		}
	case 121:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1103
		{
			yyVAL.val = yyDollar[1].val
		}
	case 122:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:1104
		{
			yyVAL.val = ""
		}
	case 123:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:1107
		{
			yyVAL.val = yyDollar[2].val
		}
	case 124:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1108
		{
			yyVAL.val = yyDollar[1].val
		}
	case 125:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1111
		{
			yyVAL.val = yyDollar[1].val
		}
	case 126:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1112
		{
			yyVAL.val = yyDollar[1].str
		}
	case 127:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1113
		{
			yyVAL.val = yyDollar[1].str
		}
	case 128:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1114
		{
			yyVAL.val = yyDollar[1].str
		}
	case 129:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1118
		{
			yyVAL.val = "TRUE"
		}
	case 130:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1119
		{
			yyVAL.val = "FALSE"
		}
	case 131:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1124
		{
			yyVAL.val = "SOME"
		}
	case 132:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1125
		{
			yyVAL.val = "ANY"
		}
	case 133:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1126
		{
			yyVAL.val = "ALL"
		}
	case 134:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1129
		{
			yyVAL.val = "="
		}
	case 135:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1130
		{
			yyVAL.val = ">="
		}
	case 136:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1131
		{
			yyVAL.val = ">"
		}
	case 137:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1132
		{
			yyVAL.val = "<="
		}
	case 138:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1133
		{
			yyVAL.val = "<"
		}
	case 139:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1134
		{
			yyVAL.val = "<>"
		}
	case 140:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1135
		{
			yyVAL.val = "!="
		}
	case 141:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1138
		{
			yyVAL.val = "IS"
		}
	case 142:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:1139
		{
			yyVAL.val = "IS NOT"
		}
	case 143:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1143
		{
			yyVAL.val = "NOT"
		}
	case 144:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:1144
		{
			yyVAL.val = ""
		}
	case 145:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1147
		{
			yyVAL.val = yyDollar[1].ident
		}
	case 146:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1151
		{
			yyVAL.val = yyDollar[1].ident
		}
	case 147:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1155
		{
			yyVAL.val = "LEFT"
		}
	case 148:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1156
		{
			yyVAL.val = "IN"
		}
	case 149:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1157
		{
			yyVAL.val = "IS"
		}
	case 150:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1160
		{
			yyVAL.val = yyDollar[1].val
		}
	case 151:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1163
		{
			yyVAL.val = yyDollar[1].ident
		}
	case 152:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1166
		{
			yyVAL.val = yyDollar[1].ident
		}
	case 153:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1177
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
	case 154:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1193
		{
			yyVAL.val = yyDollar[1].ident
		}
	case 155:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1197
		{
			yyVAL.val = yyDollar[1].val
		}
	case 156:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:1201
		{
			yyVAL.val = ""
		}
	case 157:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1206
		{
			yyVAL.val = "DISTINCT"
		}
	case 158:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1210
		{
			yyVAL.val = "DISTINCTROW"
		}
	case 159:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:1222
		{
			yyVAL.node = &ast.IntoClause{Table: yyDollar[2].val}
		}
	case 160:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:1226
		{
			yyVAL.node = nil
		}
	case 161:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1232
		{
			yyVAL.val = yyDollar[1].val
		}
	case 162:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1236
		{
			yyVAL.val = yyDollar[1].str
		}
	case 163:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1240
		{
			yyVAL.val = yyDollar[1].str
		}
	case 164:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:1255
		{
			yyVAL.node = &ast.FromClause{Table_ref: yyDollar[2].list}
		}
	case 165:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:1259
		{
			yyVAL.node = nil
		}
	case 166:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1265
		{
			yyVAL.list = ast.List{yyDollar[1].expr}
		}
	case 167:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:1269
		{
			yyVAL.list = append(yyDollar[1].list, yyDollar[3].expr)
		}
	case 168:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1275
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 169:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1279
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 170:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:1285
		{
			yyDollar[1].expr.SetAlias(yyDollar[2].val)
			yyVAL.expr = yyDollar[1].expr
		}
	case 171:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:1290
		{
			expr := &ast.SubqueryExpr{Left: nil, Right: yyDollar[1].node.(*ast.SelectStmt)}
			expr.SetTag(ast.AST_SUBQUERY)
			expr.SetAlias(yyDollar[2].val)
			yyVAL.expr = expr
		}
	case 172:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:1297
		{
			yyVAL.expr = &ast.RelationListExpr{Relation: yyDollar[2].list}
		}
	case 173:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1303
		{
			tb := &ast.RelationExpr{Table: yyDollar[1].val}
			tb.SetTag(ast.AST_RELATION)
			yyVAL.expr = tb
		}
	case 174:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:1309
		{
			tb := &ast.RelationExpr{DB: yyDollar[1].val, Table: yyDollar[2].val}
			tb.SetTag(ast.AST_RELATION)
			yyVAL.expr = tb
		}
	case 175:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1317
		{
			yyVAL.val = yyDollar[1].val
		}
	case 176:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1321
		{
			yyVAL.val = yyDollar[1].str
		}
	case 177:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1328
		{
			yyVAL.list = ast.List{yyDollar[1].expr}
		}
	case 178:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:1332
		{
			yyVAL.list = append(yyDollar[1].list, yyDollar[3].expr)
		}
	case 179:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:1341
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
	case 180:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:1356
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
	case 181:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:1371
		{

			join := &ast.JoinExpr{Jtype: "STRAIGHT_JOIN", Left: yyDollar[1].expr, Right: yyDollar[3].expr}
			join.SetTag(ast.AST_STRAIGHT_JOIN)
			yyVAL.expr = join
		}
	case 182:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:1378
		{
			join := &ast.JoinExpr{Jtype: "STRAIGHT_JOIN", Left: yyDollar[1].expr, Right: yyDollar[3].expr, Jqual: yyDollar[5].expr}
			join.SetTag(ast.AST_STRAIGHT_JOIN)
			yyVAL.expr = join
		}
	case 183:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.y:1384
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
	case 184:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.y:1395
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
	case 185:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1408
		{
			yyVAL.val = "JOIN"
		}
	case 186:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:1412
		{
			yyVAL.val = "INNER JOIN"
		}
	case 187:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:1416
		{
			yyVAL.val = "CROSS JOIN"
		}
	case 188:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1423
		{
			yyVAL.val = "LEFT"
		}
	case 189:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1427
		{
			yyVAL.val = "RIGHT"
		}
	case 190:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1434
		{
			yyVAL.val = "OUTER"
		}
	case 191:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:1438
		{
			yyVAL.val = ""
		}
	case 192:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:1444
		{
			yyVAL.expr = yyDollar[2].expr
		}
	case 193:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:1448
		{
			yyVAL.expr = &ast.RelationListExpr{Relation: yyDollar[3].list}
		}
	case 194:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1455
		{
			yyVAL.list = ast.List{yyDollar[1].expr}
		}
	case 195:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:1459
		{
			yyVAL.list = append(yyDollar[1].list, yyDollar[3].expr)
		}
	case 196:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:1472
		{
			yyVAL.node = &ast.WhereClause{Where: yyDollar[2].expr}
		}
	case 197:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:1476
		{
			yyVAL.node = nil
		}
	case 198:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:1488
		{
			yyVAL.node = &ast.GroupClause{Target_ref: yyDollar[3].list}
		}
	case 199:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:1492
		{
			yyVAL.node = nil
		}
	case 200:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1499
		{
			yyVAL.list = ast.List{yyDollar[1].expr}
		}
	case 201:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:1503
		{
			yyVAL.list = append(yyDollar[1].list, yyDollar[3].expr)
		}
	case 202:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1510
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 203:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:1524
		{
			yyVAL.node = &ast.HavingClause{Having: yyDollar[2].expr}
		}
	case 204:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:1528
		{
			yyVAL.node = nil
		}
	case 205:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1543
		{
			yyVAL.node = yyDollar[1].node
		}
	case 206:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:1547
		{
			yyVAL.node = nil
		}
	case 207:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:1553
		{
			yyVAL.node = &ast.SortClause{Target_ref: yyDollar[3].list}
		}
	case 208:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1560
		{
			yyVAL.list = ast.List{yyDollar[1].expr}
		}
	case 209:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:1564
		{
			yyVAL.list = append(yyDollar[1].list, yyDollar[3].expr)
		}
	case 210:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:1570
		{
			switch node := yyDollar[1].expr.(type) {
			case *ast.Expr:
				node.SetAsc(yyDollar[2].val)
			}

			yyVAL.expr = yyDollar[1].expr
		}
	case 211:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1581
		{
			yyVAL.val = "ASC"
		}
	case 212:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1585
		{
			yyVAL.val = "DESC"
		}
	case 213:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:1589
		{
			yyVAL.val = ""
		}
	case 214:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1600
		{
			yyVAL.node = yyDollar[1].node
		}
	case 215:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:1601
		{
			yyVAL.node = nil
		}
	case 216:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:1605
		{
			yyVAL.node = &ast.LockClause{Lock: "FOR UPDATE"}
		}
	case 217:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:1609
		{
			yyVAL.node = &ast.LockClause{Lock: "LOCK IN SHARE MODE"}
		}
	case 218:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1622
		{
			yyVAL.node = yyDollar[1].node
		}
	case 219:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:1626
		{
			yyVAL.node = nil
		}
	case 220:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:1633
		{
			yyVAL.node = &ast.LimitClause{Limit: yyDollar[1].list, Offset: yyDollar[2].expr}
		}
	case 221:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:1637
		{
			yyVAL.node = &ast.LimitClause{Limit: yyDollar[2].list, Offset: yyDollar[1].expr}
		}
	case 222:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1641
		{
			yyVAL.node = &ast.LimitClause{Limit: yyDollar[1].list}
		}
	case 223:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1645
		{
			yyVAL.node = &ast.LimitClause{Offset: yyDollar[1].expr}
		}
	case 224:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:1652
		{
			yyVAL.list = yyDollar[2].list
		}
	case 225:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:1656
		{
			yyVAL.list = append(yyDollar[2].list, yyDollar[4].expr)
		}
	case 226:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1662
		{
			yyVAL.list = ast.List{yyDollar[1].expr}
		}
	case 227:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1669
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 228:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:1676
		{
			yyVAL.expr = yyDollar[2].expr
		}
	}
	goto yystack /* stack new state and value */
}
