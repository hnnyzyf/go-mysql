//line parse.y:2
package myparser

import __yyfmt__ "fmt"

//line parse.y:4
import AST "myast"

func SetParseTree(yylex interface{}, stmt AST.SelectNode) {
	yylex.(*Tokener).ASTree = stmt
}

//line parse.y:15
type yySymType struct {
	yys     int
	keyword string
	value   string
	node    AST.Node
	snode   AST.SelectNode
	expr    AST.ExprNode
	list    AST.List
}

const IDENT = 57346
const STRING = 57347
const NUMBER = 57348
const AS = 57349
const ASC = 57350
const ALL = 57351
const ANY = 57352
const BY = 57353
const COMMENT = 57354
const CHARACTER = 57355
const COLLATE = 57356
const DISTINCT = 57357
const DISTINCTROW = 57358
const DESC = 57359
const DAY = 57360
const DAY_HOUR = 57361
const DAY_MICROSECOND = 57362
const DAY_MINUTE = 57363
const DAY_SECOND = 57364
const EXISTS = 57365
const FROM = 57366
const FOR = 57367
const FALSE = 57368
const GROUP = 57369
const HAVING = 57370
const HOUR = 57371
const HOUR_MICROSECOND = 57372
const HOUR_MINUTE = 57373
const HOUR_SECOND = 57374
const INTO = 57375
const IF = 57376
const INTERVAL = 57377
const LIMIT = 57378
const LOCK = 57379
const MINUTE = 57380
const MINUTE_SECOND = 57381
const MODE = 57382
const MONTH = 57383
const MICROSECOND = 57384
const MINUTE_MICROSECOND = 57385
const NULL = 57386
const ORDER = 57387
const OFFSET = 57388
const QUARTER = 57389
const SELECT = 57390
const SOME = 57391
const SHARE = 57392
const SECOND = 57393
const SECOND_MICROSECOND = 57394
const SET = 57395
const TO = 57396
const TRUE = 57397
const UPDATE = 57398
const USING = 57399
const WHERE = 57400
const WEEK = 57401
const YEAR = 57402
const YEAR_MONTH = 57403
const UNION = 57404
const EXCEPT = 57405
const INTERSECT = 57406
const JOIN = 57407
const STRAIGHT_JOIN = 57408
const LEFT = 57409
const RIGHT = 57410
const INNER = 57411
const OUTER = 57412
const CROSS = 57413
const NATURAL = 57414
const USE = 57415
const ON = 57416
const OR = 57417
const AND = 57418
const NOT = 57419
const BETWEEN = 57420
const CASE = 57421
const WHEN = 57422
const THEN = 57423
const ELSE = 57424
const IS = 57425
const LIKE = 57426
const IN = 57427
const GE = 57428
const LE = 57429
const NE = 57430
const IE = 57431
const LNE = 57432
const SL = 57433
const SR = 57434
const LEG = 57435
const AA = 57436
const OO = 57437
const OP = 57438
const UMINUS = 57439
const END = 57440

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"IDENT",
	"STRING",
	"NUMBER",
	"AS",
	"ASC",
	"ALL",
	"ANY",
	"BY",
	"COMMENT",
	"CHARACTER",
	"COLLATE",
	"DISTINCT",
	"DISTINCTROW",
	"DESC",
	"DAY",
	"DAY_HOUR",
	"DAY_MICROSECOND",
	"DAY_MINUTE",
	"DAY_SECOND",
	"EXISTS",
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
	"INTERVAL",
	"LIMIT",
	"LOCK",
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
	"SHARE",
	"SECOND",
	"SECOND_MICROSECOND",
	"SET",
	"TO",
	"TRUE",
	"UPDATE",
	"USING",
	"WHERE",
	"WEEK",
	"YEAR",
	"YEAR_MONTH",
	"'('",
	"'~'",
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
	"OR",
	"AND",
	"NOT",
	"BETWEEN",
	"CASE",
	"WHEN",
	"THEN",
	"ELSE",
	"'='",
	"'<'",
	"'>'",
	"IS",
	"LIKE",
	"IN",
	"GE",
	"LE",
	"NE",
	"IE",
	"LNE",
	"SL",
	"SR",
	"LEG",
	"AA",
	"OO",
	"OP",
	"'|'",
	"'&'",
	"'+'",
	"'-'",
	"'*'",
	"'/'",
	"'%'",
	"'^'",
	"UMINUS",
	"'.'",
	"END",
	"')'",
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
	114, 6,
	-2, 10,
	-1, 9,
	1, 7,
	114, 7,
	-2, 67,
	-1, 40,
	62, 199,
	-2, 46,
	-1, 269,
	13, 37,
	-2, 202,
}

const yyPrivate = 57344

const yyLast = 957

var yyAct = [...]int{

	23, 37, 380, 41, 364, 360, 354, 335, 277, 62,
	105, 223, 32, 282, 183, 220, 226, 104, 292, 127,
	171, 212, 119, 345, 57, 94, 56, 96, 97, 98,
	129, 99, 70, 69, 63, 64, 65, 66, 67, 68,
	109, 95, 68, 331, 290, 185, 186, 106, 274, 118,
	120, 213, 383, 326, 268, 372, 128, 63, 64, 65,
	66, 67, 68, 132, 135, 136, 137, 138, 139, 140,
	141, 142, 143, 144, 145, 146, 147, 148, 149, 150,
	151, 152, 153, 154, 155, 156, 157, 158, 69, 63,
	64, 65, 66, 67, 68, 178, 239, 240, 241, 382,
	325, 267, 367, 225, 4, 95, 42, 340, 188, 320,
	319, 16, 257, 339, 213, 187, 272, 280, 244, 242,
	243, 258, 256, 216, 254, 253, 251, 252, 250, 247,
	248, 249, 245, 246, 229, 100, 215, 237, 238, 239,
	240, 241, 61, 12, 181, 65, 66, 67, 68, 25,
	375, 57, 95, 266, 123, 59, 173, 123, 123, 103,
	231, 160, 159, 161, 230, 237, 238, 239, 240, 241,
	166, 7, 221, 366, 264, 114, 343, 324, 285, 286,
	287, 232, 233, 95, 283, 269, 285, 286, 287, 61,
	281, 284, 14, 365, 172, 321, 283, 165, 285, 286,
	287, 22, 281, 284, 8, 58, 323, 371, 189, 334,
	228, 11, 13, 12, 273, 162, 120, 128, 7, 349,
	177, 176, 279, 40, 42, 29, 164, 8, 291, 291,
	49, 322, 229, 271, 20, 21, 276, 298, 275, 167,
	295, 228, 31, 289, 47, 294, 297, 296, 348, 6,
	235, 163, 350, 357, 38, 59, 48, 40, 42, 29,
	175, 45, 5, 134, 99, 172, 133, 327, 234, 43,
	48, 94, 48, 332, 46, 336, 31, 16, 265, 330,
	49, 30, 49, 328, 229, 44, 278, 229, 38, 45,
	130, 115, 338, 19, 344, 341, 51, 291, 342, 28,
	121, 39, 52, 124, 125, 102, 346, 187, 185, 186,
	355, 337, 60, 122, 113, 30, 122, 122, 167, 167,
	167, 55, 116, 107, 26, 27, 24, 351, 377, 111,
	95, 178, 293, 28, 356, 39, 50, 358, 361, 259,
	260, 261, 353, 229, 229, 95, 180, 229, 290, 53,
	54, 362, 363, 214, 369, 368, 95, 42, 26, 27,
	370, 40, 42, 15, 3, 1, 374, 16, 373, 117,
	16, 229, 110, 361, 112, 217, 126, 381, 378, 376,
	9, 10, 359, 379, 224, 381, 384, 167, 167, 167,
	167, 167, 167, 167, 167, 167, 167, 167, 167, 167,
	167, 167, 167, 167, 167, 167, 167, 222, 299, 300,
	301, 302, 303, 304, 305, 306, 307, 308, 309, 310,
	311, 312, 313, 314, 315, 316, 317, 318, 85, 84,
	87, 89, 131, 18, 174, 91, 73, 71, 72, 88,
	86, 90, 83, 82, 80, 81, 79, 76, 77, 78,
	74, 75, 227, 70, 69, 63, 64, 65, 66, 67,
	68, 33, 95, 42, 352, 329, 211, 270, 108, 167,
	35, 190, 34, 101, 184, 182, 36, 262, 288, 92,
	17, 2, 0, 85, 84, 87, 89, 0, 0, 0,
	347, 73, 71, 72, 88, 86, 90, 83, 82, 80,
	81, 79, 76, 77, 78, 74, 75, 0, 70, 69,
	63, 64, 65, 66, 67, 68, 0, 95, 42, 179,
	93, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 85, 84, 87, 89,
	0, 0, 0, 0, 73, 71, 72, 88, 86, 90,
	83, 82, 80, 81, 79, 76, 77, 78, 74, 75,
	218, 70, 69, 63, 64, 65, 66, 67, 68, 219,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 85, 84, 87, 89, 0, 0, 0, 0, 73,
	71, 72, 88, 86, 90, 83, 82, 80, 81, 79,
	76, 77, 78, 74, 75, 0, 70, 69, 63, 64,
	65, 66, 67, 68, 0, 0, 0, 0, 0, 0,
	85, 84, 87, 89, 0, 0, 0, 0, 73, 71,
	72, 88, 86, 90, 83, 82, 80, 81, 79, 76,
	77, 78, 74, 75, 0, 70, 69, 63, 64, 65,
	66, 67, 68, 85, 84, 87, 89, 0, 0, 333,
	0, 73, 71, 72, 88, 86, 90, 83, 82, 80,
	81, 79, 76, 77, 78, 74, 75, 0, 70, 69,
	63, 64, 65, 66, 67, 68, 85, 84, 87, 89,
	0, 0, 0, 0, 73, 71, 72, 88, 86, 90,
	83, 82, 80, 81, 79, 76, 77, 78, 74, 75,
	0, 70, 69, 63, 64, 65, 66, 67, 68, 84,
	87, 89, 0, 0, 0, 0, 73, 71, 72, 88,
	86, 90, 83, 82, 80, 81, 79, 76, 77, 78,
	74, 75, 0, 70, 69, 63, 64, 65, 66, 67,
	68, 87, 89, 0, 0, 0, 0, 73, 71, 72,
	88, 86, 90, 83, 82, 80, 81, 79, 76, 77,
	78, 74, 75, 0, 70, 69, 63, 64, 65, 66,
	67, 68, 73, 71, 72, 88, 86, 90, 83, 82,
	80, 81, 79, 76, 77, 78, 74, 75, 0, 70,
	69, 63, 64, 65, 66, 67, 68, 255, 236, 257,
	0, 0, 40, 42, 29, 244, 242, 243, 258, 256,
	0, 254, 253, 251, 252, 250, 247, 248, 249, 245,
	246, 31, 0, 0, 237, 238, 239, 240, 241, 0,
	40, 42, 29, 38, 40, 42, 29, 0, 0, 0,
	0, 0, 0, 191, 192, 193, 194, 195, 0, 31,
	0, 0, 0, 31, 196, 197, 198, 199, 0, 0,
	30, 38, 0, 201, 203, 38, 204, 200, 202, 0,
	0, 0, 205, 0, 8, 0, 206, 207, 28, 0,
	39, 0, 0, 0, 208, 209, 210, 0, 263, 0,
	0, 0, 30, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 26, 27, 24, 28, 0, 39, 0,
	170, 0, 39, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 26, 27, 0, 0, 168, 169,
}
var yyPact = [...]int{

	156, -1000, -1000, -1000, -1000, -1000, 147, 156, 219, -1000,
	236, 287, 287, 287, 310, -88, -90, 222, 818, 75,
	-1000, -1000, -1000, 513, -1000, -1000, 253, 253, 253, -1000,
	253, 109, 291, -1000, -1000, -1000, 97, -65, 317, 253,
	-1000, -1000, -1000, 234, 264, 119, 184, 220, 253, 253,
	156, -1000, -1000, 156, 156, 253, -1000, -1000, 266, 326,
	122, 818, -1000, 253, 253, 253, 253, 253, 253, 253,
	253, 253, 253, 253, 253, 253, 253, 253, 253, 253,
	253, 253, 253, 253, 253, 253, 253, 72, 171, 850,
	94, 211, -1000, 352, -1000, -1000, -1000, -1000, 681, 405,
	-1000, -1000, 342, 30, -65, -1000, 101, 845, -32, 618,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 69, 618, -1000,
	618, 77, -1000, -1000, -1000, 77, 56, -1000, 552, 114,
	148, -1000, -1000, 266, -1000, 38, 38, -68, -68, -68,
	-1000, -48, -16, -71, -71, -71, -71, -71, -71, -71,
	-71, -71, -71, -71, -71, -71, 681, 650, -71, 253,
	126, 94, -1000, 224, -1000, -1000, 739, -1000, 850, 850,
	850, -1000, -1000, 846, 91, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -13, -1000, 253, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 31, -1000, 253, -43, 253, 253, -1000, -1000, -1000,
	259, 253, 50, 128, 341, 341, -1000, -1000, 179, -65,
	114, -71, -1000, -1000, -1000, -1000, 253, 850, 850, 850,
	850, 850, 850, 850, 850, 850, 850, 850, 850, 850,
	850, 850, 850, 850, 850, 850, 850, 20, 151, -1000,
	-1000, 32, -14, 846, 618, -1000, 846, -1000, 293, 458,
	-70, -1000, 253, 585, 159, -1000, -1000, 247, 300, 618,
	148, 45, 39, 148, 108, -1000, -1000, -1000, -1000, -1000,
	326, -1000, -1000, 37, -91, 128, -65, 259, 706, -11,
	-11, -1000, -1000, -1000, 60, 60, 60, 60, 60, 60,
	60, 60, 60, 60, 60, 60, 60, 32, 60, 850,
	193, -1000, 208, -1000, -1000, -1000, 253, 350, -1000, 357,
	297, -1000, 618, 253, 213, -1000, 253, 253, 128, 148,
	148, 116, 34, 148, -1000, -1000, 247, 60, -1000, -1000,
	-1000, 618, -1000, 297, -1000, 154, 618, -1000, 618, -12,
	-1000, 618, -1000, 116, -1000, 253, 88, 148, -1000, -1000,
	-1000, 324, 253, -1000, 618, 352, -1000, -1000, -1000, -15,
	-1000, -1000, -1000, 352, -1000,
}
var yyPgo = [...]int{

	0, 481, 363, 103, 262, 249, 480, 293, 201, 9,
	479, 478, 18, 1, 3, 0, 170, 149, 20, 477,
	12, 14, 476, 475, 474, 6, 473, 472, 471, 470,
	468, 467, 21, 466, 461, 452, 17, 10, 435, 434,
	433, 336, 205, 432, 30, 407, 11, 384, 16, 4,
	2, 13, 383, 15, 8, 382, 5, 7, 381, 380,
	376, 19, 375, 374, 269, 372, 285, 274, 244, 369,
	22, 365, 353,
}
var yyR1 = [...]int{

	0, 71, 1, 1, 3, 3, 2, 2, 2, 2,
	5, 5, 4, 4, 4, 4, 4, 40, 40, 41,
	41, 41, 6, 6, 7, 7, 8, 8, 35, 35,
	34, 34, 36, 36, 37, 37, 9, 9, 10, 10,
	13, 13, 11, 11, 12, 12, 14, 42, 42, 43,
	44, 44, 45, 45, 46, 46, 46, 47, 53, 53,
	54, 54, 55, 55, 56, 57, 57, 58, 58, 59,
	60, 60, 61, 62, 62, 62, 48, 48, 48, 48,
	48, 48, 51, 51, 51, 49, 49, 52, 52, 50,
	63, 63, 72, 64, 65, 65, 66, 66, 66, 66,
	67, 67, 69, 70, 68, 15, 15, 15, 15, 15,
	15, 15, 15, 15, 15, 15, 15, 15, 15, 15,
	15, 15, 15, 15, 15, 15, 15, 15, 15, 15,
	15, 15, 15, 15, 15, 15, 15, 15, 15, 15,
	15, 15, 15, 15, 15, 16, 16, 16, 16, 16,
	16, 16, 16, 16, 16, 16, 16, 16, 16, 16,
	16, 16, 16, 16, 16, 16, 16, 16, 16, 16,
	16, 16, 16, 16, 16, 16, 17, 17, 17, 17,
	17, 17, 17, 18, 18, 19, 19, 38, 38, 38,
	38, 38, 38, 38, 39, 39, 39, 20, 20, 22,
	23, 23, 21, 21, 21, 24, 24, 24, 25, 26,
	26, 27, 28, 28, 28, 28, 28, 28, 28, 28,
	28, 28, 28, 28, 28, 28, 28, 28, 28, 28,
	28, 28, 29, 33, 33, 32, 31, 31, 30, 30,
}
var yyR2 = [...]int{

	0, 1, 1, 1, 3, 3, 1, 2, 4, 4,
	1, 1, 7, 8, 4, 4, 4, 1, 1, 1,
	1, 0, 1, 0, 1, 3, 2, 1, 1, 2,
	1, 2, 1, 2, 2, 2, 1, 0, 2, 1,
	1, 1, 1, 0, 2, 1, 1, 2, 0, 1,
	2, 0, 1, 3, 2, 2, 1, 1, 2, 0,
	3, 0, 1, 3, 1, 2, 0, 1, 0, 3,
	1, 3, 2, 1, 1, 0, 3, 4, 5, 4,
	5, 4, 1, 1, 1, 2, 4, 1, 3, 1,
	1, 0, 0, 6, 1, 0, 2, 2, 1, 1,
	2, 4, 1, 1, 2, 1, 2, 2, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 2, 3, 4, 3, 4, 3, 4, 3, 4,
	5, 3, 4, 4, 6, 1, 2, 2, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 2, 3, 4,
	3, 4, 3, 4, 3, 4, 1, 3, 2, 2,
	1, 1, 1, 1, 3, 1, 3, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 3, 4, 1,
	1, 3, 2, 5, 4, 1, 1, 0, 3, 2,
	0, 3, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 5, 1, 2, 4, 2, 0, 1, 0,
}
var yyChk = [...]int{

	-1000, -71, -1, -2, -3, -4, -5, 62, 48, -59,
	-58, 64, 66, 65, 45, -2, -3, -6, -40, -7,
	15, 16, -8, -15, 107, -17, 105, 106, 80, 6,
	62, 23, -20, -34, -27, -29, -22, -13, 35, 82,
	4, -14, 5, -64, -66, 25, -67, -68, 36, 46,
	-41, 9, 15, -41, -41, 11, 114, 114, -42, 33,
	-7, 67, -9, 105, 106, 107, 108, 109, 110, 104,
	103, 87, 88, 86, 100, 101, 97, 98, 99, 96,
	94, 95, 93, 92, 79, 78, 90, 80, 89, 81,
	91, -38, -10, 7, -13, 4, -15, -15, -15, -15,
	-3, -26, 14, 62, -36, -37, 112, 6, -30, -15,
	-65, -66, -63, -64, 56, -68, -67, -69, -15, -70,
	-15, -5, -4, -3, -5, -5, -60, -61, -15, -44,
	24, -43, -14, -42, -8, -15, -15, -15, -15, -15,
	-15, -15, -15, -15, -15, -15, -15, -15, -15, -15,
	-15, -15, -15, -15, -15, -15, -15, -15, -15, 90,
	89, 91, 44, 80, 55, 26, -16, -17, 105, 106,
	80, -18, -3, 62, -39, 49, 10, 9, -13, 114,
	4, 114, -23, -21, -24, 15, 16, -37, -13, 107,
	-28, 18, 19, 20, 21, 22, 29, 30, 31, 32,
	42, 38, 43, 39, 41, 47, 51, 52, 59, 60,
	61, -33, -32, 83, -72, 67, 67, -62, 8, 17,
	-53, 58, -45, -46, -47, -3, -48, -35, 62, -14,
	-44, -15, 55, -18, 44, 26, 79, 105, 106, 107,
	108, 109, 87, 88, 86, 100, 101, 97, 98, 99,
	96, 94, 95, 93, 92, 78, 90, 80, 89, -16,
	-16, -16, -19, 62, -15, -3, 62, 114, 67, -15,
	-31, -32, 85, -15, 91, -70, -61, -54, 27, -15,
	67, 74, -51, 68, 75, 70, 71, 72, -11, -12,
	7, -14, -12, -3, -48, -46, -36, -53, -15, -16,
	-16, -16, -16, -16, -16, -16, -16, -16, -16, -16,
	-16, -16, -16, -16, -16, -16, -16, -16, -16, 90,
	89, 44, 80, 55, 26, 114, 67, -15, -21, 7,
	-9, 113, -15, 84, 50, -57, 28, 11, -46, 68,
	68, -46, -51, 68, -14, 114, -54, -16, 55, 26,
	44, -15, 114, -20, -25, 13, -15, 40, -15, -55,
	-56, -15, -46, -46, -49, 77, 57, 68, -46, -57,
	-25, 53, 67, -49, -15, 62, -46, 4, -56, -52,
	-50, -13, 114, 67, -50,
}
var yyDef = [...]int{

	0, -2, 1, 2, -2, -2, 68, 0, 23, -2,
	0, 21, 21, 21, 0, 0, 11, 48, 0, 22,
	17, 18, 24, 37, 27, 105, 0, 0, 0, 176,
	0, 0, 210, 180, 181, 182, 0, 30, 0, 239,
	-2, 40, 41, 95, 91, 0, 98, 99, 0, 0,
	0, 19, 20, 0, 0, 0, 4, 5, 51, 0,
	48, 0, 26, 0, 0, 0, 0, 0, 0, 0,
	0, 187, 188, 189, 0, 0, 0, 0, 0, 0,
	193, 190, 192, 191, 0, 0, 0, 0, 0, 0,
	0, 0, 36, 0, 39, 46, 106, 107, 131, 0,
	178, 179, 0, 207, 31, 32, 0, 0, 0, 238,
	8, 94, 9, 90, 92, 96, 97, 100, 102, 104,
	103, 14, 10, 11, 15, 16, 69, 70, 75, 59,
	0, 47, 49, 51, 25, 108, 109, 110, 111, 112,
	113, 114, 115, 116, 117, 118, 119, 120, 121, 122,
	123, 124, 125, 126, 127, 128, 129, 130, 132, 0,
	0, 0, 134, 0, 136, 138, 0, 145, 0, 0,
	0, 141, 183, 0, 0, 194, 195, 196, 38, 177,
	209, 197, 0, 200, 0, 205, 206, 33, 34, 35,
	211, 212, 213, 214, 215, 216, 217, 218, 219, 220,
	221, 222, 223, 224, 225, 226, 227, 228, 229, 230,
	231, 237, 233, 0, 0, 0, 0, 72, 73, 74,
	61, 0, 50, 52, 43, 0, 56, 57, 0, 28,
	59, 133, 137, 142, 135, 139, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 146,
	147, 167, 0, 0, 185, 143, 0, 198, 207, -2,
	0, 234, 0, 0, 0, 101, 71, 66, 0, 58,
	0, 0, 0, 0, 0, 82, 83, 84, 54, 42,
	0, 45, 55, 11, 56, 0, 29, 61, 140, 148,
	149, 150, 151, 152, 153, 154, 155, 156, 157, 158,
	159, 160, 161, 162, 163, 164, 165, 166, 168, 0,
	0, 170, 0, 172, 174, 184, 0, 0, 201, 0,
	0, 232, 236, 0, 0, 12, 0, 0, 53, 0,
	0, 0, 0, 0, 44, 76, 66, 169, 173, 175,
	171, 186, 144, 0, 204, 0, 235, 93, 65, 60,
	62, 64, 77, 0, 79, 0, 0, 0, 81, 13,
	203, 0, 0, 78, 85, 0, 80, 208, 63, 0,
	87, 89, 86, 0, 88,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 109, 104, 3,
	62, 114, 107, 105, 67, 106, 112, 108, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	87, 86, 88, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 110, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 103, 3, 63,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
	52, 53, 54, 55, 56, 57, 58, 59, 60, 61,
	64, 65, 66, 68, 69, 70, 71, 72, 73, 74,
	75, 76, 77, 78, 79, 80, 81, 82, 83, 84,
	85, 89, 90, 91, 92, 93, 94, 95, 96, 97,
	98, 99, 100, 101, 102, 111, 113,
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
		//line parse.y:167
		{
			SetParseTree(yylex, yyDollar[1].snode)
		}
	case 4:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:179
		{
			yyVAL.snode = yyDollar[2].snode
		}
	case 5:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:180
		{
			yyVAL.snode = yyDollar[2].snode
		}
	case 6:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:184
		{
			yyVAL.snode = yyDollar[1].snode
		}
	case 7:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parse.y:186
		{
			yyDollar[1].snode.(*AST.SimpleSelectStmt).Sortby = yyDollar[2].node.(*AST.Sortclause)
			yyVAL.snode = yyDollar[1].snode
		}
	case 8:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parse.y:191
		{
			if yyDollar[2].node != nil {
				yyDollar[1].snode.(*AST.SimpleSelectStmt).Sortby = yyDollar[2].node.(*AST.Sortclause)
			}
			yyDollar[1].snode.(*AST.SimpleSelectStmt).Lock = yyDollar[3].node.(*AST.Lockingclause)
			if yyDollar[4].node != nil {
				yyDollar[1].snode.(*AST.SimpleSelectStmt).Limit = yyDollar[4].node.(*AST.Limitclause)
			}
			yyVAL.snode = yyDollar[1].snode
		}
	case 9:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parse.y:202
		{
			if yyDollar[2].node != nil {
				yyDollar[1].snode.(*AST.SimpleSelectStmt).Sortby = yyDollar[2].node.(*AST.Sortclause)
			}
			if yyDollar[4].node != nil {
				yyDollar[1].snode.(*AST.SimpleSelectStmt).Lock = yyDollar[4].node.(*AST.Lockingclause)
			}
			yyDollar[1].snode.(*AST.SimpleSelectStmt).Limit = yyDollar[3].node.(*AST.Limitclause)
			yyVAL.snode = yyDollar[1].snode
		}
	case 10:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:214
		{
			yyVAL.snode = yyDollar[1].snode
		}
	case 11:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:215
		{
			yyVAL.snode = yyDollar[1].snode
		}
	case 12:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parse.y:220
		{
			//创建一个新的Select结构体
			Select := AST.NewSimpleSelectStmt()
			Select.NType = AST.AST_SELECT
			Select.Op = AST.SETOP_SIMPLE
			Select.Target = yyDollar[2].list
			Select.Into = yyDollar[3].node
			Select.From = yyDollar[4].node
			Select.Where = yyDollar[5].node
			Select.Groupby = yyDollar[6].node
			Select.Having = yyDollar[7].node
			yyVAL.snode = Select
		}
	case 13:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line parse.y:234
		{
			//创建一个新的Select结构体
			Select := AST.NewSimpleSelectStmt()
			Select.NType = AST.AST_SELECT
			Select.Op = AST.SETOP_SIMPLE
			Select.Distinct = yyDollar[2].node
			Select.Target = yyDollar[3].list
			Select.Into = yyDollar[4].node
			Select.From = yyDollar[5].node
			Select.Where = yyDollar[6].node
			Select.Groupby = yyDollar[7].node
			Select.Having = yyDollar[8].node
			yyVAL.snode = Select
		}
	case 14:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parse.y:249
		{
			Stmt := AST.NewSelectStmt()
			Stmt.NType = AST.AST_SELECT
			Stmt.Having = yyDollar[2].keyword
			Stmt.Op = AST.SETOP_UNION
			Stmt.Left = yyDollar[1].snode
			Stmt.Right = yyDollar[4].snode
			yyVAL.snode = Stmt
		}
	case 15:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parse.y:259
		{
			Stmt := AST.NewSelectStmt()
			Stmt.NType = AST.AST_SELECT
			Stmt.Having = yyDollar[2].keyword
			Stmt.Op = AST.SETOP_INTERSECT
			Stmt.Left = yyDollar[1].snode
			Stmt.Right = yyDollar[4].snode
			yyVAL.snode = Stmt
		}
	case 16:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parse.y:269
		{
			Stmt := AST.NewSelectStmt()
			Stmt.NType = AST.AST_SELECT
			Stmt.Having = yyDollar[2].keyword
			Stmt.Op = AST.SETOP_EXCEPT
			Stmt.Left = yyDollar[1].snode
			Stmt.Right = yyDollar[4].snode
			yyVAL.snode = Stmt
		}
	case 17:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:287
		{
			distinct := AST.NewDistinctclause()
			yyVAL.node = distinct
		}
	case 18:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:292
		{
			distinct := AST.NewDistinctclause()
			yyVAL.node = distinct
		}
	case 19:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:301
		{
			yyVAL.keyword = "ALL"
		}
	case 20:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:305
		{
			yyVAL.keyword = "DISTINCT"
		}
	case 21:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parse.y:309
		{
			yyVAL.keyword = ""
		}
	case 22:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:323
		{
			yyVAL.list = yyDollar[1].list
		}
	case 23:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parse.y:327
		{
			yyVAL.list = nil
		}
	case 24:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:333
		{
			yyVAL.list = AST.NewList(yyDollar[1].expr)

		}
	case 25:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:338
		{
			yyVAL.list = append(yyDollar[1].list, yyDollar[3].expr)
		}
	case 26:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parse.y:345
		{
			target := AST.NewTargetExpr()
			target.Tuple = yyDollar[1].expr
			target.Alias = yyDollar[2].value
			yyVAL.expr = target
		}
	case 27:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:352
		{
			target := AST.NewTargetExpr()
			star := AST.NewStartExpr()
			target.Tuple = star
			yyVAL.expr = target
		}
	case 28:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:369
		{
			tb := AST.NewTableExpr()
			tb.Table = yyDollar[1].value
			yyVAL.expr = tb
		}
	case 29:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parse.y:375
		{
			tb := AST.NewTableExpr()
			tb.DB = yyDollar[1].value
			tb.Table = yyDollar[2].value
			yyVAL.expr = tb
		}
	case 30:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:385
		{
			col := AST.NewColumnExpr()
			col.Column = yyDollar[1].value

			yyVAL.expr = col
		}
	case 31:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parse.y:392
		{
			col := AST.NewColumnExpr()
			col.Table = yyDollar[1].value
			col.Column = yyDollar[2].value

			yyVAL.expr = col
		}
	case 32:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:402
		{
			//返回字符串
			yyVAL.value = yyDollar[1].value
		}
	case 33:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parse.y:407
		{
			//返回字符串
			yyVAL.value = yyDollar[1].value + yyDollar[2].value
		}
	case 34:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parse.y:415
		{
			//返回字符串
			yyVAL.value = yyDollar[2].value
		}
	case 35:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parse.y:420
		{
			//返回字符串
			yyVAL.value = "*"
		}
	case 36:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:433
		{
			yyVAL.value = yyDollar[1].value
		}
	case 37:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parse.y:437
		{
			yyVAL.value = ""
		}
	case 38:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parse.y:443
		{
			//可以为INDET和STRING
			yyVAL.value = yyDollar[2].value
		}
	case 39:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:448
		{
			yyVAL.value = yyDollar[1].value
		}
	case 40:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:454
		{
			yyVAL.value = yyDollar[1].value
		}
	case 41:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:458
		{
			yyVAL.value = yyDollar[1].value
		}
	case 42:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:465
		{
			yyVAL.value = yyDollar[1].value
		}
	case 43:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parse.y:469
		{
			yyVAL.value = ""
		}
	case 44:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parse.y:475
		{
			//可以为INDET和STRING
			yyVAL.value = yyDollar[2].value
		}
	case 45:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:480
		{
			yyVAL.value = yyDollar[1].value
		}
	case 46:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:486
		{
			yyVAL.value = yyDollar[1].value
		}
	case 47:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parse.y:498
		{
			Into := AST.NewIntoclause()
			Into.Tbname = yyDollar[2].value
			yyVAL.node = Into
		}
	case 48:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parse.y:504
		{
			yyVAL.node = nil
		}
	case 49:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:510
		{
			yyVAL.value = yyDollar[1].value
		}
	case 50:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parse.y:522
		{
			From := AST.NewFromclause()
			From.Expr = yyDollar[2].list
			yyVAL.node = From
		}
	case 51:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parse.y:528
		{
			yyVAL.node = nil
		}
	case 52:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:534
		{
			switch yyDollar[1].expr.(type) {
			case *AST.SimpleTableExpr:
				yyVAL.list = AST.NewList(yyDollar[1].expr)
			case *AST.SubTableExpr:
				yyVAL.list = AST.NewList(yyDollar[1].expr)
			case *AST.JoinExpr:
				yyVAL.list = AST.NewList(yyDollar[1].expr)
			}
		}
	case 53:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:545
		{
			yyVAL.list = append(yyDollar[1].list, yyDollar[3].expr)
		}
	case 54:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parse.y:551
		{
			SimpleTable := AST.NewSimpleTableExpr()
			SimpleTable.Table = yyDollar[1].expr.(*AST.TableExpr)
			SimpleTable.Alias = yyDollar[2].value
			yyVAL.expr = SimpleTable
		}
	case 55:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parse.y:558
		{
			SubTable := AST.NewSubTableExpr()
			SubTable.Stmt = yyDollar[1].snode
			SubTable.Alias = yyDollar[2].value
			yyVAL.expr = SubTable
		}
	case 56:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:565
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 57:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:571
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 58:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parse.y:584
		{
			Where := AST.NewWhereclause()
			Where.Expr = yyDollar[2].expr
			yyVAL.node = Where
		}
	case 59:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parse.y:590
		{
			yyVAL.node = nil
		}
	case 60:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:603
		{
			Group := AST.NewGroupclause()
			Group.Expr = yyDollar[3].list
			yyVAL.node = Group
		}
	case 61:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parse.y:609
		{
			yyVAL.node = nil
		}
	case 62:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:616
		{
			yyVAL.list = AST.NewList(yyDollar[1].expr)
		}
	case 63:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:620
		{
			yyVAL.list = append(yyDollar[1].list, yyDollar[3].expr)
		}
	case 64:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:627
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 65:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parse.y:640
		{
			Having := AST.NewHavingclause()
			Having.Expr = yyDollar[2].expr
			yyVAL.node = Having
		}
	case 66:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parse.y:646
		{
			yyVAL.node = nil
		}
	case 67:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:661
		{
			yyVAL.node = yyDollar[1].node
		}
	case 68:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parse.y:665
		{
			yyVAL.node = nil
		}
	case 69:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:671
		{
			Sort := AST.NewSortclause()
			Sort.Expr = yyDollar[3].list
			yyVAL.node = Sort
		}
	case 70:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:680
		{
			yyVAL.list = AST.NewList(yyDollar[1].expr)
		}
	case 71:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:684
		{
			yyVAL.list = append(yyDollar[1].list, yyDollar[3].expr)
		}
	case 72:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parse.y:690
		{
			SortBy := AST.NewAscExpr()
			SortBy.Expr = yyDollar[1].expr
			SortBy.Asc = yyDollar[2].keyword
			if SortBy.Asc == "" {
				SortBy.Not = false
			} else {
				SortBy.Not = true
			}
			yyVAL.expr = SortBy
		}
	case 73:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:705
		{
			yyVAL.keyword = "ASC"
		}
	case 74:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:709
		{
			yyVAL.keyword = "DESC"
		}
	case 75:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parse.y:713
		{
			yyVAL.keyword = ""
		}
	case 76:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:725
		{
			yyVAL.expr = yyDollar[2].expr
		}
	case 77:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parse.y:729
		{
			Join := AST.NewJoinExpr(AST.EXPR_JOIN_CROSS)
			Join.Jointype = "CROSS"
			Join.Left = yyDollar[1].expr
			Join.Right = yyDollar[4].expr

			yyVAL.expr = Join
		}
	case 78:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parse.y:738
		{
			Join := AST.NewJoinExpr(AST.EXPR_JOIN)
			switch yyDollar[2].keyword {
			case "LEFT":
				Join.Ntype = AST.EXPR_JOIN_LEFT
				Join.Jointype = "LEFT"
			case "RIGHT":
				Join.Ntype = AST.EXPR_JOIN_RIGHT
				Join.Jointype = "RIGHT"
			case "INNER":
				Join.Ntype = AST.EXPR_JOIN_INNER
				Join.Jointype = "INNER"
			}
			Join.Joinqual = yyDollar[5].expr
			Join.Left = yyDollar[1].expr
			Join.Right = yyDollar[4].expr

			yyVAL.expr = Join
		}
	case 79:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parse.y:758
		{
			Join := AST.NewJoinExpr(AST.EXPR_JOIN)
			Join.Jointype = ""
			Join.Joinqual = yyDollar[4].expr
			Join.Left = yyDollar[1].expr
			Join.Right = yyDollar[3].expr

			yyVAL.expr = Join
		}
	case 80:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parse.y:768
		{
			Join := AST.NewJoinExpr(AST.EXPR_JOIN_NATURAL)
			switch yyDollar[2].keyword {
			case "LEFT":
				Join.Ntype = AST.EXPR_JOIN_NATURAL_LEFT
				Join.Jointype = "NATURAL LEFT"
			case "RIGHT":
				Join.Ntype = AST.EXPR_JOIN_NATURAL_RIGHT
				Join.Jointype = "NATURAL RIGHT"
			case "INNER":
				Join.Ntype = AST.EXPR_JOIN_NATURAL_INNER
				Join.Jointype = "NATURAL INNER"
			}
			Join.Left = yyDollar[1].expr
			Join.Right = yyDollar[5].expr

			yyVAL.expr = Join
		}
	case 81:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parse.y:787
		{
			Join := AST.NewJoinExpr(AST.EXPR_JOIN_NATURAL)
			Join.Jointype = "NATURAL"
			Join.Left = yyDollar[1].expr
			Join.Right = yyDollar[4].expr

			yyVAL.expr = Join
		}
	case 82:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:798
		{
			yyVAL.keyword = "LEFT"
		}
	case 83:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:802
		{
			yyVAL.keyword = "RIGHT"
		}
	case 84:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:806
		{
			yyVAL.keyword = "INNER"
		}
	case 85:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parse.y:812
		{
			On := AST.NewOnExpr(AST.EXPR_JOIN_ON)
			On.Expr = yyDollar[2].expr

			yyVAL.expr = On
		}
	case 86:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parse.y:819
		{
			Use := AST.NewUseExpr(AST.EXPR_JOIN_USING)
			Use.Expr = yyDollar[3].list

			yyVAL.expr = Use
		}
	case 87:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:829
		{
			yyVAL.list = AST.NewList(yyDollar[1].expr)
		}
	case 88:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:833
		{
			yyVAL.list = append(yyDollar[1].list, yyDollar[3].expr)
		}
	case 89:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:839
		{
			col := AST.NewColumnExpr()
			col.Column = yyDollar[1].value
			yyVAL.expr = col
		}
	case 90:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:857
		{
			yyVAL.node = yyDollar[1].node
		}
	case 91:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parse.y:861
		{
			yyVAL.node = nil
		}
	case 92:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parse.y:867
		{
			Lock := AST.NewLockingclause()
			Lock.Expr = "FOR UPDATE"
			yyVAL.node = Lock
		}
	case 93:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parse.y:873
		{
			Lock := AST.NewLockingclause()
			Lock.Expr = "IN SHARE MODE"
			yyVAL.node = Lock
		}
	case 94:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:888
		{
			yyVAL.node = yyDollar[1].node
		}
	case 95:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parse.y:892
		{
			yyVAL.node = nil
		}
	case 96:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parse.y:899
		{
			Limit := AST.NewLimitclause()
			Limit.Limit = yyDollar[1].list
			Limit.Offset = yyDollar[2].expr
			yyVAL.node = Limit
		}
	case 97:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parse.y:906
		{
			Limit := AST.NewLimitclause()
			Limit.Limit = yyDollar[2].list
			Limit.Offset = yyDollar[1].expr
			yyVAL.node = Limit
		}
	case 98:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:913
		{
			Limit := AST.NewLimitclause()
			Limit.Limit = yyDollar[1].list
			yyVAL.node = Limit
		}
	case 99:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:919
		{
			Limit := AST.NewLimitclause()
			Limit.Offset = yyDollar[1].expr
			yyVAL.node = Limit
		}
	case 100:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parse.y:928
		{
			yyVAL.list = yyDollar[2].list
		}
	case 101:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parse.y:932
		{

			yyVAL.list = append(yyDollar[2].list, yyDollar[4].expr)
		}
	case 102:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:939
		{
			yyVAL.list = AST.NewList(yyDollar[1].expr)
		}
	case 103:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:945
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 104:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parse.y:952
		{
			yyVAL.expr = yyDollar[2].expr
		}
	case 105:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:964
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 106:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parse.y:968
		{
			yyVAL.expr = yyDollar[2].expr
		}
	case 107:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parse.y:972
		{
			yyVAL.expr = AST.DoValueExpr(yyDollar[2].expr)
		}
	case 108:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:976
		{
			bool := AST.NewBoolExpr(AST.EXPR_BOOL_PLUS)
			bool.Op = "+"
			bool.Left = yyDollar[1].expr
			bool.Right = yyDollar[3].expr
			yyVAL.expr = bool

		}
	case 109:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:985
		{
			bool := AST.NewBoolExpr(AST.EXPR_BOOL_MINUS)
			bool.Op = "-"
			bool.Left = yyDollar[1].expr
			bool.Right = yyDollar[3].expr
			yyVAL.expr = bool

		}
	case 110:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:994
		{
			bool := AST.NewBoolExpr(AST.EXPR_BOOL_MULTIPLY)
			bool.Op = "*"
			bool.Left = yyDollar[1].expr
			bool.Right = yyDollar[3].expr
			yyVAL.expr = bool

		}
	case 111:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:1003
		{
			bool := AST.NewBoolExpr(AST.EXPR_BOOL_DIVIDE)
			bool.Op = "/"
			bool.Left = yyDollar[1].expr
			bool.Right = yyDollar[3].expr
			yyVAL.expr = bool

		}
	case 112:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:1012
		{
			bool := AST.NewBoolExpr(AST.EXPR_BOOL_REMAINDER)
			bool.Op = "%"
			bool.Left = yyDollar[1].expr
			bool.Right = yyDollar[3].expr
			yyVAL.expr = bool

		}
	case 113:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:1021
		{
			bool := AST.NewBoolExpr(AST.EXPR_BOOL_XOR)
			bool.Op = "^"
			bool.Left = yyDollar[1].expr
			bool.Right = yyDollar[3].expr
			yyVAL.expr = bool

		}
	case 114:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:1030
		{
			bool := AST.NewBoolExpr(AST.EXPR_BOOL_AND)
			bool.Op = "&"
			bool.Left = yyDollar[1].expr
			bool.Right = yyDollar[3].expr
			yyVAL.expr = bool
		}
	case 115:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:1038
		{
			bool := AST.NewBoolExpr(AST.EXPR_BOOL_OR)
			bool.Op = "|"
			bool.Left = yyDollar[1].expr
			bool.Right = yyDollar[3].expr
			yyVAL.expr = bool

		}
	case 116:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:1047
		{
			bool := AST.NewBoolExpr(AST.EXPR_BOOL_L)
			bool.Op = "<"
			bool.Left = yyDollar[1].expr
			bool.Right = yyDollar[3].expr
			yyVAL.expr = bool

		}
	case 117:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:1056
		{
			bool := AST.NewBoolExpr(AST.EXPR_BOOL_G)
			bool.Op = ">"
			bool.Left = yyDollar[1].expr
			bool.Right = yyDollar[3].expr
			yyVAL.expr = bool

		}
	case 118:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:1065
		{
			bool := AST.NewBoolExpr(AST.EXPR_BOOL_E)
			bool.Op = "="
			bool.Left = yyDollar[1].expr
			bool.Right = yyDollar[3].expr
			yyVAL.expr = bool

		}
	case 119:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:1074
		{
			bool := AST.NewBoolExpr(AST.EXPR_BOOL_AA)
			bool.Op = "&&"
			bool.Left = yyDollar[1].expr
			bool.Right = yyDollar[3].expr
			yyVAL.expr = bool

		}
	case 120:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:1083
		{
			bool := AST.NewBoolExpr(AST.EXPR_BOOL_OO)
			bool.Op = "||"
			bool.Left = yyDollar[1].expr
			bool.Right = yyDollar[3].expr
			yyVAL.expr = bool

		}
	case 121:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:1092
		{
			bool := AST.NewBoolExpr(AST.EXPR_BOOL_SL)
			bool.Op = "<<"
			bool.Left = yyDollar[1].expr
			bool.Right = yyDollar[3].expr
			yyVAL.expr = bool

		}
	case 122:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:1101
		{
			bool := AST.NewBoolExpr(AST.EXPR_BOOL_SR)
			bool.Op = ">>"
			bool.Left = yyDollar[1].expr
			bool.Right = yyDollar[3].expr
			yyVAL.expr = bool

		}
	case 123:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:1110
		{
			bool := AST.NewBoolExpr(AST.EXPR_BOOL_LEG)
			bool.Op = "<=>"
			bool.Left = yyDollar[1].expr
			bool.Right = yyDollar[3].expr
			yyVAL.expr = bool
		}
	case 124:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:1118
		{
			bool := AST.NewBoolExpr(AST.EXPR_BOOL_NE)
			bool.Op = "!="
			bool.Left = yyDollar[1].expr
			bool.Right = yyDollar[3].expr
			yyVAL.expr = bool

		}
	case 125:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:1127
		{
			bool := AST.NewBoolExpr(AST.EXPR_BOOL_NE)
			bool.Op = "<>"
			bool.Left = yyDollar[1].expr
			bool.Right = yyDollar[3].expr
			yyVAL.expr = bool

		}
	case 126:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:1136
		{
			bool := AST.NewBoolExpr(AST.EXPR_BOOL_IE)
			bool.Op = "=="
			bool.Left = yyDollar[1].expr
			bool.Right = yyDollar[3].expr
			yyVAL.expr = bool

		}
	case 127:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:1145
		{
			bool := AST.NewBoolExpr(AST.EXPR_BOOL_LE)
			bool.Op = "<="
			bool.Left = yyDollar[1].expr
			bool.Right = yyDollar[3].expr
			yyVAL.expr = bool

		}
	case 128:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:1154
		{
			bool := AST.NewBoolExpr(AST.EXPR_BOOL_GE)
			bool.Op = ">="
			bool.Left = yyDollar[1].expr
			bool.Right = yyDollar[3].expr
			yyVAL.expr = bool

		}
	case 129:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:1163
		{
			bool := AST.NewBoolExpr(AST.EXPR_BOOL_AND)
			bool.Op = "AND"
			bool.Left = yyDollar[1].expr
			bool.Right = yyDollar[3].expr
			yyVAL.expr = bool

		}
	case 130:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:1172
		{
			bool := AST.NewBoolExpr(AST.EXPR_BOOL_OR)
			bool.Op = "OR"
			bool.Left = yyDollar[1].expr
			bool.Right = yyDollar[3].expr
			yyVAL.expr = bool

		}
	case 131:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parse.y:1181
		{
			bool := AST.NewBoolExpr(AST.EXPR_BOOL_NOT)
			bool.Op = "NOT"
			bool.Left = nil
			bool.Right = yyDollar[2].expr
			yyVAL.expr = bool
		}
	case 132:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:1190
		{
			bool := AST.NewBoolExpr(AST.EXPR_BOOL_LIKE)
			bool.Op = "LIKE"
			bool.Left = yyDollar[1].expr
			bool.Right = yyDollar[3].expr
			yyVAL.expr = bool
		}
	case 133:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parse.y:1198
		{
			bool := AST.NewBoolExpr(AST.EXPR_BOOL_NOT_LIKE)
			bool.Op = "NOT LIKE"
			bool.Left = yyDollar[1].expr
			bool.Right = yyDollar[4].expr
			yyVAL.expr = bool
		}
	case 134:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:1206
		{
			Is := AST.NewIsExpr(AST.EXPR_NULL)
			Is.Not = false
			Is.Left = yyDollar[1].expr
			yyVAL.expr = Is
		}
	case 135:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parse.y:1213
		{
			Is := AST.NewIsExpr(AST.EXPR_NULL)
			Is.Not = true
			Is.Left = yyDollar[1].expr
			yyVAL.expr = Is
		}
	case 136:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:1220
		{
			Is := AST.NewIsExpr(AST.EXPR_TRUE)
			Is.Not = false
			Is.Left = yyDollar[1].expr
			yyVAL.expr = Is
		}
	case 137:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parse.y:1227
		{
			Is := AST.NewIsExpr(AST.EXPR_TRUE)
			Is.Not = true
			Is.Left = yyDollar[1].expr
			yyVAL.expr = Is
		}
	case 138:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:1234
		{
			Is := AST.NewIsExpr(AST.EXPR_FALSE)
			Is.Not = false
			Is.Left = yyDollar[1].expr
			yyVAL.expr = Is
		}
	case 139:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parse.y:1241
		{
			Is := AST.NewIsExpr(AST.EXPR_FALSE)
			Is.Not = true
			Is.Left = yyDollar[1].expr
			yyVAL.expr = Is
		}
	case 140:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parse.y:1248
		{
			Between := AST.NewRangeExpr()
			Between.Expr = yyDollar[1].expr
			Between.Left = yyDollar[3].expr
			Between.Right = yyDollar[5].expr
			yyVAL.expr = Between
		}
	case 141:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:1256
		{
			In := AST.NewInExpr(AST.EXPR_IN_SIMPLE)
			In.Not = false
			In.Left = yyDollar[1].expr
			In.Right = yyDollar[3].list
			//判断InExpr的类型
			//for expr := range In.Right{
			//	switch expr.(type){
			//		case *AST.SubTableExpr:
			//			In.Ntype=AST.EXPR_IN_SUBLINK
			//	}
			//}
			yyVAL.expr = In
		}
	case 142:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parse.y:1271
		{
			In := AST.NewInExpr(AST.EXPR_IN_SIMPLE)
			In.Not = true
			In.Left = yyDollar[1].expr
			In.Right = yyDollar[4].list
			//判断InExpr的类型
			//for expr := range In.Right{
			//	switch expr.(type){
			//		case *AST.SubTableExpr:
			//			In.Ntype=AST.EXPR_IN_SUBLINK
			//	}
			//}
			yyVAL.expr = In
		}
	case 143:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parse.y:1286
		{
			Sub := AST.NewSubQueryExpr(AST.EXPR_SUBQUERY)
			Sub.Op = yyDollar[2].value
			Sub.Subtype = yyDollar[3].value
			Sub.Left = yyDollar[1].expr
			Sub.Right = yyDollar[4].snode

			switch Sub.Subtype {
			case "SOME":
				Sub.Ntype = AST.EXPR_SUBQUERY_SOME
			case "ANY":
				Sub.Ntype = AST.EXPR_SUBQUERY_ANY
			case "ALL":
				Sub.Ntype = AST.EXPR_SUBQUERY_ALL
			}

			yyVAL.expr = Sub
		}
	case 144:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parse.y:1305
		{
			Sub := AST.NewSubQueryExpr(AST.EXPR_SUBQUERY)
			Sub.Op = yyDollar[2].value
			Sub.Subtype = yyDollar[3].value
			Sub.Left = yyDollar[1].expr
			Sub.Right = yyDollar[5].expr

			switch Sub.Subtype {
			case "SOME":
				Sub.Ntype = AST.EXPR_SUBQUERY_SOME
			case "ANY":
				Sub.Ntype = AST.EXPR_SUBQUERY_ANY
			case "ALL":
				Sub.Ntype = AST.EXPR_SUBQUERY_ALL
			}

			yyVAL.expr = Sub
		}
	case 145:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:1326
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 146:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parse.y:1330
		{
			yyVAL.expr = yyDollar[2].expr
		}
	case 147:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parse.y:1334
		{
			yyVAL.expr = AST.DoValueExpr(yyDollar[2].expr)
		}
	case 148:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:1338
		{
			bool := AST.NewBoolExpr(AST.EXPR_BOOL_PLUS)
			bool.Op = "+"
			bool.Left = yyDollar[1].expr
			bool.Right = yyDollar[3].expr
			yyVAL.expr = bool

		}
	case 149:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:1347
		{
			bool := AST.NewBoolExpr(AST.EXPR_BOOL_MINUS)
			bool.Op = "-"
			bool.Left = yyDollar[1].expr
			bool.Right = yyDollar[3].expr
			yyVAL.expr = bool

		}
	case 150:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:1356
		{
			bool := AST.NewBoolExpr(AST.EXPR_BOOL_MULTIPLY)
			bool.Op = "*"
			bool.Left = yyDollar[1].expr
			bool.Right = yyDollar[3].expr
			yyVAL.expr = bool

		}
	case 151:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:1365
		{
			bool := AST.NewBoolExpr(AST.EXPR_BOOL_DIVIDE)
			bool.Op = "/"
			bool.Left = yyDollar[1].expr
			bool.Right = yyDollar[3].expr
			yyVAL.expr = bool

		}
	case 152:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:1374
		{
			bool := AST.NewBoolExpr(AST.EXPR_BOOL_REMAINDER)
			bool.Op = "%"
			bool.Left = yyDollar[1].expr
			bool.Right = yyDollar[3].expr
			yyVAL.expr = bool

		}
	case 153:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:1383
		{
			bool := AST.NewBoolExpr(AST.EXPR_BOOL_L)
			bool.Op = "<"
			bool.Left = yyDollar[1].expr
			bool.Right = yyDollar[3].expr
			yyVAL.expr = bool

		}
	case 154:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:1392
		{
			bool := AST.NewBoolExpr(AST.EXPR_BOOL_G)
			bool.Op = ">"
			bool.Left = yyDollar[1].expr
			bool.Right = yyDollar[3].expr
			yyVAL.expr = bool

		}
	case 155:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:1401
		{
			bool := AST.NewBoolExpr(AST.EXPR_BOOL_E)
			bool.Op = "="
			bool.Left = yyDollar[1].expr
			bool.Right = yyDollar[3].expr
			yyVAL.expr = bool

		}
	case 156:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:1410
		{
			bool := AST.NewBoolExpr(AST.EXPR_BOOL_AA)
			bool.Op = "&&"
			bool.Left = yyDollar[1].expr
			bool.Right = yyDollar[3].expr
			yyVAL.expr = bool

		}
	case 157:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:1419
		{
			bool := AST.NewBoolExpr(AST.EXPR_BOOL_OO)
			bool.Op = "||"
			bool.Left = yyDollar[1].expr
			bool.Right = yyDollar[3].expr
			yyVAL.expr = bool

		}
	case 158:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:1428
		{
			bool := AST.NewBoolExpr(AST.EXPR_BOOL_SL)
			bool.Op = "<<"
			bool.Left = yyDollar[1].expr
			bool.Right = yyDollar[3].expr
			yyVAL.expr = bool

		}
	case 159:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:1437
		{
			bool := AST.NewBoolExpr(AST.EXPR_BOOL_SR)
			bool.Op = ">>"
			bool.Left = yyDollar[1].expr
			bool.Right = yyDollar[3].expr
			yyVAL.expr = bool

		}
	case 160:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:1446
		{
			bool := AST.NewBoolExpr(AST.EXPR_BOOL_LEG)
			bool.Op = "<=>"
			bool.Left = yyDollar[1].expr
			bool.Right = yyDollar[3].expr
			yyVAL.expr = bool
		}
	case 161:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:1454
		{
			bool := AST.NewBoolExpr(AST.EXPR_BOOL_NE)
			bool.Op = "!="
			bool.Left = yyDollar[1].expr
			bool.Right = yyDollar[3].expr
			yyVAL.expr = bool

		}
	case 162:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:1463
		{
			bool := AST.NewBoolExpr(AST.EXPR_BOOL_NE)
			bool.Op = "<>"
			bool.Left = yyDollar[1].expr
			bool.Right = yyDollar[3].expr
			yyVAL.expr = bool

		}
	case 163:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:1472
		{
			bool := AST.NewBoolExpr(AST.EXPR_BOOL_IE)
			bool.Op = "=="
			bool.Left = yyDollar[1].expr
			bool.Right = yyDollar[3].expr
			yyVAL.expr = bool

		}
	case 164:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:1481
		{
			bool := AST.NewBoolExpr(AST.EXPR_BOOL_LE)
			bool.Op = "<="
			bool.Left = yyDollar[1].expr
			bool.Right = yyDollar[3].expr
			yyVAL.expr = bool

		}
	case 165:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:1491
		{
			bool := AST.NewBoolExpr(AST.EXPR_BOOL_GE)
			bool.Op = ">="
			bool.Left = yyDollar[1].expr
			bool.Right = yyDollar[3].expr
			yyVAL.expr = bool

		}
	case 166:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:1500
		{
			bool := AST.NewBoolExpr(AST.EXPR_BOOL_OR)
			bool.Op = "OR"
			bool.Left = yyDollar[1].expr
			bool.Right = yyDollar[3].expr
			yyVAL.expr = bool

		}
	case 167:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parse.y:1509
		{
			bool := AST.NewBoolExpr(AST.EXPR_BOOL_NOT)
			bool.Op = "NOT"
			bool.Left = nil
			bool.Right = yyDollar[2].expr
			yyVAL.expr = bool
		}
	case 168:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:1517
		{
			bool := AST.NewBoolExpr(AST.EXPR_BOOL_LIKE)
			bool.Op = "LIKE"
			bool.Left = yyDollar[1].expr
			bool.Right = yyDollar[3].expr
			yyVAL.expr = bool
		}
	case 169:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parse.y:1525
		{
			bool := AST.NewBoolExpr(AST.EXPR_BOOL_NOT_LIKE)
			bool.Op = "NOT LIKE"
			bool.Left = yyDollar[1].expr
			bool.Right = yyDollar[4].expr
			yyVAL.expr = bool
		}
	case 170:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:1533
		{
			Is := AST.NewIsExpr(AST.EXPR_NULL)
			Is.Not = false
			Is.Left = yyDollar[1].expr
			yyVAL.expr = Is
		}
	case 171:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parse.y:1540
		{
			Is := AST.NewIsExpr(AST.EXPR_NULL)
			Is.Not = true
			Is.Left = yyDollar[1].expr
			yyVAL.expr = Is
		}
	case 172:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:1547
		{
			Is := AST.NewIsExpr(AST.EXPR_TRUE)
			Is.Not = false
			Is.Left = yyDollar[1].expr
			yyVAL.expr = Is
		}
	case 173:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parse.y:1554
		{
			Is := AST.NewIsExpr(AST.EXPR_TRUE)
			Is.Not = true
			Is.Left = yyDollar[1].expr
			yyVAL.expr = Is
		}
	case 174:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:1561
		{
			Is := AST.NewIsExpr(AST.EXPR_FALSE)
			Is.Not = false
			Is.Left = yyDollar[1].expr
			yyVAL.expr = Is
		}
	case 175:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parse.y:1568
		{
			Is := AST.NewIsExpr(AST.EXPR_FALSE)
			Is.Not = true
			Is.Left = yyDollar[1].expr
			yyVAL.expr = Is
		}
	case 176:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:1577
		{
			s := AST.NewValueExpr(AST.VALUE_NUMBER)
			s.Val = yyDollar[1].value
			yyVAL.expr = s
		}
	case 177:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:1583
		{
			yyVAL.expr = yyDollar[2].expr
		}
	case 178:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parse.y:1587
		{
			Sub := AST.NewSubQueryExpr(AST.EXPR_SUBQUERY_EXISTS)
			Sub.Subtype = "EXISTS"
			Sub.Left = nil
			Sub.Right = yyDollar[2].snode
			yyVAL.expr = Sub
		}
	case 179:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parse.y:1595
		{
			yyVAL.expr.(*AST.FuncExpr).Collate = yyDollar[2].value

		}
	case 180:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:1600
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 181:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:1604
		{
			s := AST.NewValueExpr(AST.VALUE_STRING)
			s.Val = yyDollar[1].value
			yyVAL.expr = s
		}
	case 182:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:1610
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 183:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:1627
		{
			yyVAL.list = AST.NewList(yyDollar[1].snode)
		}
	case 184:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:1631
		{
			yyVAL.list = yyDollar[2].list
		}
	case 185:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:1638
		{
			yyVAL.list = AST.NewList(yyDollar[1].expr)
		}
	case 186:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:1642
		{
			yyVAL.list = append(yyDollar[1].list, yyDollar[3].expr)
		}
	case 187:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:1649
		{
			yyVAL.value = "<"
		}
	case 188:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:1653
		{
			yyVAL.value = ">"
		}
	case 189:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:1657
		{
			yyVAL.value = "="
		}
	case 190:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:1661
		{
			yyVAL.value = "=="
		}
	case 191:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:1665
		{
			yyVAL.value = ">="
		}
	case 192:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:1669
		{
			yyVAL.value = "<="
		}
	case 193:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:1673
		{
			yyVAL.value = "<>"
		}
	case 194:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:1680
		{
			yyVAL.value = "SOME"
		}
	case 195:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:1684
		{
			yyVAL.value = "ANY"
		}
	case 196:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:1688
		{
			yyVAL.value = "ALL"
		}
	case 197:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:1700
		{
			Func := AST.NewFuncExpr(AST.EXPR_FUNC)
			Func.Funcname = yyDollar[1].keyword
			Func.Expr = nil

			yyVAL.expr = Func
		}
	case 198:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parse.y:1708
		{
			Func := AST.NewFuncExpr(AST.EXPR_FUNC)
			Func.Funcname = yyDollar[1].keyword
			Func.Expr = yyDollar[3].list

			yyVAL.expr = Func
		}
	case 199:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:1718
		{
			yyVAL.keyword = yyDollar[1].value
		}
	case 200:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:1725
		{
			yyVAL.list = AST.NewList(yyDollar[1].expr)
		}
	case 201:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:1729
		{
			yyVAL.list = append(yyDollar[1].list, yyDollar[3].expr)
		}
	case 202:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parse.y:1736
		{
			arg := AST.NewFuncArgExpr()
			arg.Distinct = yyDollar[1].keyword
			arg.Expr = yyDollar[2].expr
			yyVAL.expr = arg
		}
	case 203:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parse.y:1743
		{
			arg := AST.NewFuncArgExpr()
			arg.Distinct = yyDollar[1].keyword
			arg.Expr = yyDollar[2].expr
			arg.TypeAlias = yyDollar[4].expr
			arg.Character = yyDollar[5].value
			yyVAL.expr = arg
		}
	case 204:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parse.y:1752
		{
			arg := AST.NewFuncArgExpr()
			arg.Distinct = yyDollar[1].keyword
			arg.Expr = yyDollar[2].expr
			arg.Alias = yyDollar[3].value
			arg.Character = yyDollar[4].value
			yyVAL.expr = arg
		}
	case 205:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:1762
		{
			yyVAL.keyword = " DISTINCT "
		}
	case 206:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:1766
		{
			yyVAL.keyword = " DISTINCTROW "
		}
	case 207:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parse.y:1770
		{
			yyVAL.keyword = ""
		}
	case 208:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:1776
		{
			yyVAL.value = yyDollar[3].value
		}
	case 209:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parse.y:1783
		{
			yyVAL.value = yyDollar[2].value
		}
	case 210:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parse.y:1787
		{
			yyVAL.value = ""
		}
	case 211:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:1801
		{
			yyVAL.value = " INTERVAL " + yyDollar[2].value + yyDollar[3].value
		}
	case 212:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:1805
		{
			yyVAL.value = "DAY"
		}
	case 213:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:1806
		{
			yyVAL.value = "DAY_HOUR"
		}
	case 214:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:1807
		{
			yyVAL.value = "DAY_MICROSECOND"
		}
	case 215:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:1808
		{
			yyVAL.value = "DAY_MINUTE"
		}
	case 216:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:1809
		{
			yyVAL.value = "DAY_SECOND"
		}
	case 217:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:1810
		{
			yyVAL.value = "HOUR"
		}
	case 218:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:1811
		{
			yyVAL.value = "HOUR_MICROSECOND"
		}
	case 219:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:1812
		{
			yyVAL.value = "HOUR_MINUTE"
		}
	case 220:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:1813
		{
			yyVAL.value = "HOUR_SECOND"
		}
	case 221:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:1814
		{
			yyVAL.value = "MICROSECOND"
		}
	case 222:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:1815
		{
			yyVAL.value = "MINUTE"
		}
	case 223:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:1816
		{
			yyVAL.value = "MINUTE_MICROSECOND"
		}
	case 224:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:1817
		{
			yyVAL.value = "MINUTE_SECOND"
		}
	case 225:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:1818
		{
			yyVAL.value = "MONTH"
		}
	case 226:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:1819
		{
			yyVAL.value = "QUARTER"
		}
	case 227:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:1820
		{
			yyVAL.value = "SECOND"
		}
	case 228:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:1821
		{
			yyVAL.value = "SECOND_MICROSECOND"
		}
	case 229:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:1822
		{
			yyVAL.value = "WEEK"
		}
	case 230:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:1823
		{
			yyVAL.value = "YEAR"
		}
	case 231:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:1824
		{
			yyVAL.value = "YEAR_MONTH"
		}
	case 232:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parse.y:1835
		{
			ca := AST.NewCaseExpr(AST.EXPR_CASE)
			ca.Expr = yyDollar[2].expr
			ca.When = yyDollar[3].list
			ca.Else = yyDollar[4].expr

			yyVAL.expr = ca
		}
	case 233:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:1847
		{
			yyVAL.list = AST.NewList(yyDollar[1].expr)
		}
	case 234:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parse.y:1851
		{
			yyVAL.list = append(yyDollar[1].list, yyDollar[2].expr)
		}
	case 235:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parse.y:1858
		{
			when := AST.NewWhenExpr()
			when.Left = yyDollar[2].expr
			when.Right = yyDollar[4].expr

			yyVAL.expr = when
		}
	case 236:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parse.y:1869
		{
			el := AST.NewElseExpr()
			el.Expr = yyDollar[2].expr
			yyVAL.expr = el
		}
	case 237:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parse.y:1875
		{
			yyVAL.expr = nil
		}
	case 238:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:1881
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 239:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parse.y:1885
		{
			yyVAL.expr = nil
		}
	}
	goto yystack /* stack new state and value */
}
