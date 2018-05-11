//line ./advisor/parser/parser.y:2
package parser

import __yyfmt__ "fmt"

//line ./advisor/parser/parser.y:4
import "github.com/hnnyzyf/mysql-go/advisor/ast"
import "strings"

//line ./advisor/parser/parser.y:11
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
const BuiltinCharacterName = 57350
const IDENT = 57351
const DOUBLE_QUOTA_STRING = 57352
const PARAM_MARKER = 57353
const SINGLE_QUOTA_STRING = 57354
const STRING = 57355
const BIT_NUMBER = 57356
const HEX_NUMBER = 57357
const INTEGER = 57358
const FLOAT = 57359
const ALL = 57360
const ANY = 57361
const AS = 57362
const ASC = 57363
const AVG = 57364
const BY = 57365
const CAST = 57366
const CHARACTER = 57367
const COMMENT = 57368
const CONVERT = 57369
const COUNT = 57370
const DESC = 57371
const DISTINCT = 57372
const DISTINCTROW = 57373
const FALSE = 57374
const FOR = 57375
const FROM = 57376
const GROUP = 57377
const HAVING = 57378
const IF = 57379
const INTO = 57380
const LIMIT = 57381
const LOCK = 57382
const MAX = 57383
const MIN = 57384
const MODE = 57385
const NULL = 57386
const OFFSET = 57387
const ORDER = 57388
const POSITION = 57389
const QUARTER = 57390
const SELECT = 57391
const SET = 57392
const SHARE = 57393
const SOME = 57394
const SOUNDS = 57395
const SUM = 57396
const TO = 57397
const TRUE = 57398
const UPDATE = 57399
const WHERE = 57400
const EXISTS = 57401
const OUTER = 57402
const LOW = 57403
const UNION = 57404
const EXCEPT = 57405
const INTERSECT = 57406
const JOIN = 57407
const STRAIGHT_JOIN = 57408
const NATURAL = 57409
const LEFT = 57410
const RIGHT = 57411
const INNER = 57412
const CROSS = 57413
const USE = 57414
const TABLE_REF_PRIORITY = 57415
const ON = 57416
const USING = 57417
const OR = 57418
const OO = 57419
const XOR = 57420
const AND = 57421
const AA = 57422
const BETWEEN = 57423
const CASE = 57424
const WHEN = 57425
const THEN = 57426
const ELSE = 57427
const LIKE = 57428
const REGEXP = 57429
const IN = 57430
const GE = 57431
const LE = 57432
const NE = 57433
const LG = 57434
const IE = 57435
const LNE = 57436
const LEG = 57437
const IS = 57438
const SL = 57439
const SR = 57440
const DIV = 57441
const MOD = 57442
const NOT = 57443
const OP = 57444
const COLLATE = 57445
const INTERVAL = 57446
const UMINUS = 57447
const END = 57448

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"BuiltinCharacterIdent",
	"BuiltinFucTimeAddIdent",
	"BuiltinFucTimeSubIdent",
	"BuiltinTimeUnitIdent",
	"BuiltinCharacterName",
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
	121, 6,
	-2, 10,
	-1, 8,
	10, 79,
	11, 79,
	12, 79,
	13, 79,
	14, 79,
	15, 79,
	16, 79,
	17, 79,
	32, 79,
	44, 79,
	56, 79,
	124, 79,
	-2, 118,
	-1, 9,
	1, 7,
	121, 7,
	-2, 208,
	-1, 28,
	82, 149,
	90, 149,
	91, 149,
	92, 149,
	-2, 36,
	-1, 43,
	123, 124,
	-2, 76,
	-1, 49,
	84, 111,
	-2, 79,
	-1, 151,
	33, 11,
	39, 11,
	40, 11,
	45, 11,
	46, 11,
	62, 11,
	63, 11,
	64, 11,
	-2, 63,
	-1, 274,
	120, 180,
	-2, 166,
	-1, 275,
	120, 181,
	-2, 167,
}

const yyPrivate = 57344

const yyLast = 1064

var yyAct = [...]int{

	240, 384, 239, 38, 4, 267, 382, 356, 43, 24,
	362, 16, 27, 266, 334, 318, 272, 312, 42, 24,
	351, 263, 93, 28, 94, 178, 101, 186, 242, 8,
	188, 396, 29, 97, 279, 78, 7, 51, 152, 167,
	76, 151, 166, 153, 20, 51, 193, 51, 193, 275,
	157, 275, 165, 79, 80, 88, 89, 90, 91, 92,
	51, 99, 164, 100, 98, 163, 147, 148, 149, 150,
	114, 95, 162, 161, 160, 177, 179, 159, 158, 155,
	182, 290, 187, 182, 182, 8, 340, 401, 24, 196,
	197, 198, 199, 200, 191, 400, 378, 290, 368, 379,
	290, 330, 190, 7, 97, 377, 290, 88, 89, 90,
	91, 92, 376, 299, 83, 82, 204, 205, 115, 116,
	117, 118, 119, 120, 121, 122, 125, 123, 124, 126,
	288, 232, 127, 154, 210, 126, 350, 406, 349, 216,
	217, 218, 219, 220, 222, 224, 225, 226, 227, 228,
	229, 348, 234, 371, 367, 233, 344, 128, 139, 271,
	247, 271, 289, 237, 252, 253, 254, 255, 51, 173,
	215, 96, 83, 236, 138, 143, 137, 136, 132, 133,
	134, 135, 88, 89, 90, 91, 92, 243, 21, 293,
	84, 256, 243, 270, 387, 245, 145, 399, 274, 244,
	388, 249, 250, 251, 366, 106, 110, 108, 140, 5,
	277, 107, 109, 365, 111, 278, 112, 103, 105, 72,
	144, 394, 284, 315, 285, 276, 347, 88, 89, 90,
	91, 92, 90, 91, 92, 73, 280, 69, 282, 283,
	259, 286, 68, 287, 294, 295, 281, 297, 296, 258,
	300, 301, 302, 121, 122, 125, 123, 124, 126, 179,
	187, 202, 385, 386, 212, 314, 50, 201, 91, 92,
	292, 346, 213, 214, 211, 327, 195, 194, 97, 97,
	274, 87, 332, 151, 310, 329, 181, 311, 146, 181,
	181, 339, 325, 175, 341, 326, 305, 338, 331, 6,
	336, 88, 89, 90, 91, 92, 169, 337, 174, 355,
	323, 324, 171, 85, 352, 352, 352, 363, 264, 270,
	270, 270, 172, 361, 274, 274, 274, 353, 354, 359,
	360, 11, 13, 12, 75, 364, 97, 257, 74, 191,
	87, 309, 74, 373, 75, 345, 372, 369, 85, 370,
	374, 375, 326, 88, 89, 90, 91, 92, 380, 383,
	88, 89, 90, 91, 92, 357, 70, 88, 89, 90,
	91, 92, 270, 71, 313, 389, 180, 274, 391, 183,
	184, 189, 390, 358, 81, 51, 395, 231, 397, 392,
	393, 1, 270, 15, 3, 383, 241, 274, 19, 403,
	404, 402, 398, 270, 343, 405, 63, 64, 274, 41,
	51, 298, 88, 89, 90, 91, 92, 86, 185, 246,
	51, 193, 176, 52, 192, 59, 17, 381, 60, 53,
	265, 22, 23, 119, 120, 121, 122, 125, 123, 124,
	126, 130, 54, 55, 45, 145, 328, 47, 61, 63,
	64, 269, 41, 51, 131, 56, 235, 203, 335, 129,
	39, 268, 209, 208, 32, 291, 52, 31, 59, 144,
	65, 60, 53, 116, 117, 118, 119, 120, 121, 122,
	125, 123, 124, 126, 49, 54, 55, 77, 46, 40,
	156, 61, 26, 66, 14, 44, 207, 30, 56, 22,
	23, 67, 141, 39, 62, 48, 33, 34, 248, 142,
	11, 13, 12, 65, 35, 36, 25, 88, 89, 90,
	91, 92, 273, 104, 37, 342, 9, 49, 88, 89,
	90, 91, 92, 2, 10, 168, 66, 63, 64, 170,
	41, 51, 51, 99, 67, 100, 98, 18, 113, 33,
	34, 50, 102, 95, 52, 206, 59, 35, 36, 60,
	53, 316, 22, 23, 58, 57, 230, 37, 260, 306,
	261, 0, 0, 54, 55, 0, 0, 0, 262, 61,
	307, 88, 89, 90, 91, 92, 56, 0, 0, 0,
	0, 39, 88, 89, 90, 91, 92, 0, 0, 63,
	64, 65, 41, 51, 117, 118, 119, 120, 121, 122,
	125, 123, 124, 126, 0, 49, 52, 0, 59, 0,
	0, 60, 53, 0, 66, 0, 88, 89, 90, 91,
	92, 0, 67, 0, 0, 54, 55, 33, 34, 50,
	0, 61, 0, 8, 0, 35, 36, 25, 56, 0,
	0, 0, 304, 39, 0, 37, 63, 64, 0, 41,
	51, 0, 0, 65, 88, 89, 90, 91, 92, 303,
	0, 0, 0, 52, 0, 59, 0, 49, 60, 53,
	0, 88, 89, 90, 91, 92, 66, 0, 0, 0,
	70, 0, 54, 55, 67, 0, 74, 71, 61, 33,
	34, 50, 75, 0, 0, 56, 0, 35, 36, 25,
	39, 0, 308, 0, 0, 51, 99, 37, 100, 98,
	65, 115, 116, 117, 118, 119, 120, 121, 122, 125,
	123, 124, 126, 0, 49, 320, 317, 319, 323, 324,
	321, 322, 0, 66, 385, 386, 0, 0, 0, 0,
	0, 67, 0, 0, 0, 0, 33, 34, 50, 63,
	64, 0, 41, 51, 35, 36, 25, 0, 0, 0,
	0, 0, 238, 0, 37, 0, 52, 0, 59, 0,
	0, 60, 53, 115, 116, 117, 118, 119, 120, 121,
	122, 125, 123, 124, 126, 54, 55, 0, 0, 0,
	0, 61, 63, 64, 0, 41, 51, 0, 56, 0,
	0, 0, 0, 39, 0, 0, 0, 0, 0, 52,
	0, 59, 0, 65, 60, 53, 320, 317, 319, 323,
	324, 321, 322, 0, 0, 0, 0, 49, 54, 55,
	0, 0, 0, 0, 61, 0, 66, 0, 0, 0,
	0, 56, 0, 0, 67, 0, 39, 0, 0, 33,
	34, 50, 0, 0, 0, 0, 65, 35, 36, 25,
	0, 0, 0, 0, 0, 0, 0, 37, 0, 0,
	49, 0, 0, 0, 0, 0, 0, 0, 0, 66,
	63, 64, 0, 41, 51, 51, 99, 67, 100, 98,
	0, 0, 33, 34, 50, 0, 95, 52, 0, 59,
	35, 36, 60, 53, 0, 223, 0, 0, 0, 0,
	37, 0, 0, 0, 0, 0, 54, 55, 0, 0,
	0, 0, 61, 63, 64, 0, 41, 51, 0, 56,
	0, 0, 0, 0, 39, 0, 0, 0, 0, 0,
	52, 0, 59, 0, 65, 60, 53, 0, 0, 0,
	0, 0, 0, 88, 89, 90, 91, 92, 49, 54,
	55, 0, 0, 0, 0, 61, 0, 66, 0, 0,
	0, 0, 56, 0, 0, 67, 0, 39, 0, 0,
	33, 34, 50, 0, 0, 0, 0, 65, 35, 36,
	0, 0, 0, 221, 0, 0, 0, 0, 37, 0,
	0, 49, 0, 0, 0, 0, 0, 0, 0, 0,
	66, 0, 0, 0, 0, 0, 0, 0, 67, 0,
	0, 333, 0, 33, 34, 50, 0, 0, 0, 0,
	0, 35, 36, 0, 0, 0, 0, 0, 0, 0,
	0, 37, 115, 116, 117, 118, 119, 120, 121, 122,
	125, 123, 124, 126,
}
var yyPact = [...]int{

	-20, -1000, -1000, -1000, -1000, -1000, 448, -20, 532, -1000,
	657, 469, 469, 469, 361, -6, -7, 310, 754, 216,
	-1000, -1000, -1000, -1000, 886, 754, 118, -1000, 17, 40,
	164, -1000, -1000, 928, 928, 928, 928, 594, -1000, -87,
	-1000, -1000, -1000, 13, -1000, -1000, -1000, -1000, -44, 754,
	-1000, -1000, -45, -46, -49, -50, -51, -58, -61, -71,
	-81, -84, -1000, -1000, -1000, -1000, -1000, -1000, 299, 333,
	265, 77, 289, 303, 754, 754, -20, -1000, -1000, -20,
	-20, 754, -1000, -1000, 347, 411, 275, 754, 754, 754,
	754, 754, 754, -1000, -1000, 706, -1000, -1000, -1000, -1000,
	-1000, -1000, 413, 928, 444, 19, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 182, 80, 928, 928, 928, 928, 885,
	797, 928, 928, 928, 928, 928, 928, -1000, 383, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, 11, -1000, -1000, -1000, 28, 40, 40, 40,
	40, -7, 335, -1000, 159, 651, 108, 515, 469, 401,
	469, 469, 469, 754, 754, 754, 754, 928, -1000, -1000,
	-1000, -1000, -1000, 286, -1000, -1000, 184, 515, -1000, 515,
	269, -1000, -1000, -1000, 269, 175, -1000, 549, 260, 38,
	-1000, -1000, -1000, -1000, 347, -1000, 153, 153, 188, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -87, -1000, -1000, -1000,
	-1000, -89, 928, 928, 928, 928, 371, 501, 328, 328,
	146, 754, 146, 754, 23, 23, 23, 23, 23, -1000,
	-1000, -1000, 376, -1000, 376, -1000, -1000, 10, -1000, 41,
	515, 103, -1000, 754, 754, 754, 754, 290, -8, 754,
	754, 754, 604, 587, 276, 504, 620, 298, 754, 754,
	-1000, -1000, -1000, 339, 754, 158, 760, -1000, -1000, 533,
	533, 36, -1000, -19, -1000, -1000, 260, -1000, -1000, 594,
	951, 40, 682, 682, 451, 451, -1000, -1000, 159, -1000,
	754, -36, -1000, 754, 440, 283, 35, 224, -1000, -1000,
	150, 105, 30, 20, 18, 376, 376, 376, 754, -1000,
	-1000, -1000, 329, 360, 515, 38, 38, 38, 257, 241,
	-1000, 147, 138, -1000, -1000, -1000, -1000, 51, 33, 760,
	411, 339, 32, 928, -1000, -1000, -1000, -1000, -1000, 515,
	-1000, 515, 754, -1000, -1000, -1000, -1000, -1000, -1000, 754,
	754, -9, -1000, -16, -25, -22, -1000, 754, 754, 760,
	187, 119, 134, -1000, 257, -1000, -1000, -1000, 38, -1000,
	329, -1000, -1000, 515, 451, 451, -1000, -1000, -1000, -1000,
	515, 156, -1000, 515, -1000, 754, -92, 754, 38, 131,
	760, -1000, -26, -34, 754, 515, 754, 515, 669, 38,
	-1000, -1000, -1000, 16, -1000, -1000, -1000,
}
var yyPgo = [...]int{

	0, 40, 568, 566, 565, 564, 561, 35, 555, 14,
	552, 548, 158, 15, 10, 547, 242, 30, 17, 7,
	190, 539, 535, 534, 299, 237, 393, 533, 3, 209,
	526, 21, 24, 171, 20, 523, 522, 509, 505, 8,
	18, 504, 22, 502, 16, 497, 495, 23, 492, 490,
	489, 488, 467, 465, 0, 464, 6, 1, 461, 459,
	454, 235, 12, 451, 25, 32, 27, 447, 5, 13,
	446, 188, 444, 441, 28, 2, 430, 427, 219, 426,
	422, 418, 398, 396, 391,
}
var yyR1 = [...]int{

	0, 84, 27, 27, 28, 28, 26, 26, 26, 26,
	24, 24, 29, 29, 29, 29, 29, 54, 54, 54,
	54, 54, 54, 54, 54, 48, 48, 48, 48, 48,
	62, 62, 62, 62, 62, 62, 62, 47, 47, 47,
	47, 47, 47, 47, 47, 47, 47, 47, 47, 47,
	47, 47, 65, 65, 65, 65, 65, 65, 65, 65,
	65, 65, 65, 65, 65, 65, 59, 59, 59, 59,
	59, 59, 59, 59, 59, 39, 40, 40, 45, 45,
	73, 73, 73, 52, 52, 52, 55, 55, 55, 55,
	55, 55, 46, 46, 46, 46, 46, 46, 46, 46,
	72, 72, 51, 51, 51, 67, 75, 75, 60, 50,
	49, 49, 83, 83, 74, 53, 53, 79, 79, 82,
	82, 71, 43, 37, 38, 38, 42, 42, 32, 32,
	33, 33, 33, 33, 12, 12, 8, 8, 8, 35,
	35, 35, 35, 35, 35, 35, 10, 10, 11, 11,
	9, 3, 41, 41, 41, 34, 4, 5, 15, 1,
	1, 1, 7, 7, 20, 20, 44, 44, 44, 17,
	17, 76, 76, 69, 69, 68, 68, 68, 63, 63,
	36, 36, 70, 70, 58, 58, 58, 58, 58, 58,
	6, 6, 6, 13, 13, 14, 14, 57, 57, 31,
	31, 18, 18, 77, 77, 56, 19, 19, 23, 23,
	30, 81, 81, 66, 2, 2, 2, 21, 21, 16,
	16, 22, 22, 25, 25, 25, 25, 78, 78, 80,
	61, 64,
}
var yyR2 = [...]int{

	0, 1, 1, 1, 3, 3, 1, 2, 4, 4,
	1, 1, 7, 8, 4, 4, 4, 3, 3, 3,
	3, 3, 2, 1, 3, 3, 3, 3, 4, 1,
	4, 6, 6, 4, 4, 4, 1, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 5,
	5, 1, 2, 1, 2, 1, 2, 3, 2, 2,
	2, 2, 3, 1, 2, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 0,
	1, 1, 3, 1, 3, 5, 1, 1, 1, 1,
	3, 4, 5, 5, 5, 4, 4, 5, 5, 5,
	8, 8, 6, 6, 6, 6, 1, 3, 1, 5,
	1, 0, 1, 2, 4, 2, 0, 1, 0, 1,
	3, 2, 2, 3, 1, 1, 1, 0, 2, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 2, 1, 0,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 0, 1, 1, 2, 0, 1, 1, 1, 2,
	0, 1, 3, 1, 1, 2, 2, 3, 1, 3,
	1, 1, 1, 3, 3, 4, 3, 5, 6, 6,
	1, 2, 2, 1, 1, 1, 0, 2, 4, 2,
	0, 3, 0, 1, 3, 1, 2, 0, 1, 0,
	3, 1, 3, 2, 1, 1, 0, 1, 0, 2,
	4, 1, 0, 2, 2, 1, 1, 2, 4, 1,
	2, 1,
}
var yyChk = [...]int{

	-1000, -84, -27, -26, -28, -29, -24, 123, 49, -30,
	-23, 62, 64, 63, 46, -26, -28, -79, -15, -82,
	-7, -71, 30, 31, -54, 115, -48, -62, -47, -65,
	-45, -52, -55, 105, 106, 113, 114, 123, -28, 59,
	-50, 8, -40, -39, -46, -72, -51, -67, -38, 83,
	107, 9, 22, 28, 41, 42, 54, -4, -5, 24,
	27, 47, -41, 5, 6, 69, 92, 100, -16, -25,
	33, 40, -78, -61, 39, 45, -1, 18, -7, -1,
	-1, 23, 121, 121, -20, 38, -82, 65, 77, 78,
	79, 80, 81, -42, -32, 20, -33, -39, 13, 10,
	12, -54, -10, 99, -35, 100, 87, 93, 89, 94,
	88, 96, 98, -11, 53, 101, 102, 103, 104, 105,
	106, 107, 108, 110, 111, 109, 112, 115, 117, -59,
	-73, -60, 14, 15, 16, 17, 13, 12, 10, -12,
	44, -43, -37, 11, 56, 32, 124, -65, -65, -65,
	-65, -28, -54, -28, 120, 123, -49, -54, 123, 123,
	123, 123, 123, 123, 123, 123, 123, 123, -22, -25,
	-21, -16, 57, 92, -61, -78, -80, -54, -64, -54,
	-24, -29, -28, -24, -24, -81, -66, -54, -17, 34,
	-44, -39, 13, 10, -20, -71, -54, -54, -54, -54,
	-54, -33, -12, 44, -62, -62, -8, 52, 19, 18,
	115, 92, 82, 90, 91, 90, -47, -47, -47, -47,
	-47, 118, -47, 118, -47, -47, -47, -47, -47, -47,
	-3, 4, 120, -39, 124, 121, -40, -39, 121, -75,
	-54, -83, -74, 84, -1, -7, 18, -54, 107, -1,
	-1, -1, -54, -54, -54, -54, -47, 51, 65, 65,
	-2, 21, 29, -31, 58, -76, -69, -68, -58, -63,
	-28, 123, -44, -36, -39, 13, -17, -28, -28, 123,
	-47, -65, -47, -47, -54, -54, -39, -39, 120, 121,
	65, -53, -74, 86, -54, -54, -75, -54, 121, 121,
	-54, -54, -54, 65, 65, 20, 65, 76, 92, 43,
	-64, -66, -18, 35, -54, 65, -6, 67, -13, 68,
	66, 71, 72, 69, 70, -42, -32, -28, -70, -69,
	120, -31, -75, 80, -9, 7, -9, -40, -39, -54,
	122, -54, 85, 121, 121, 121, 121, 121, 121, 118,
	118, -34, -39, -34, -34, -54, -19, 36, 23, -69,
	-69, -68, -14, 60, -13, 66, 66, 121, 65, -44,
	-18, 121, -62, -54, -54, -54, 121, 121, 121, 121,
	-54, -77, -56, -54, -57, 75, 76, 75, 66, -14,
	-69, -19, -9, -9, 65, -54, 123, -54, -69, 66,
	121, 121, -56, -75, -57, -68, 121,
}
var yyDef = [...]int{

	0, -2, 1, 2, -2, -2, 209, 0, -2, -2,
	0, 161, 161, 161, 0, 0, 11, 165, 79, 117,
	158, 119, 162, 163, 127, 79, 23, 29, -2, 51,
	0, 53, 55, 79, 79, 79, 79, 79, 63, 0,
	65, 78, 83, -2, 86, 87, 88, 89, 0, -2,
	77, 75, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 125, 156, 157, 152, 153, 154, 222, 218,
	0, 0, 225, 226, 79, 79, 0, 159, 160, 0,
	0, 79, 4, 5, 170, 0, 165, 79, 79, 79,
	79, 79, 79, 121, 126, 0, 129, 130, 131, 132,
	133, 22, 0, 79, 79, 146, 139, 140, 141, 142,
	143, 144, 145, 0, 0, 79, 79, 79, 79, 79,
	79, 79, 79, 79, 79, 79, 79, 148, 0, 52,
	54, 56, 66, 67, 68, 69, 70, 71, 72, 73,
	74, 80, 81, 108, 134, 135, 0, 58, 59, 60,
	61, -2, 0, 64, 0, 79, 0, 110, 161, 79,
	161, 161, 161, 79, 79, 79, 79, 79, 8, 221,
	9, 217, 219, 0, 223, 224, 227, 229, 230, 231,
	14, 10, 11, 15, 16, 210, 211, 216, 200, 0,
	164, 166, 167, 168, 170, 120, 17, 18, 19, 20,
	21, 128, 24, 25, 26, 27, 0, 136, 137, 138,
	147, 0, 79, 79, 79, 79, 37, 38, 39, 40,
	41, 79, 42, 79, 43, 44, 45, 46, 47, 48,
	57, 151, 0, 122, 0, 62, 84, 76, 90, 0,
	106, 116, 112, 79, 79, 79, 79, 0, 77, 79,
	79, 79, 0, 0, 0, 0, 0, 0, 79, 79,
	213, 214, 215, 202, 79, 169, 171, 173, 174, 127,
	0, 0, 178, 0, -2, -2, 200, 28, 30, 79,
	0, 34, 35, 33, 0, 0, 82, 123, 0, 91,
	79, 0, 113, 79, 0, 0, 0, 0, 95, 96,
	0, 0, 0, 0, 0, 0, 0, 0, 79, 220,
	228, 212, 207, 0, 199, 0, 0, 0, 196, 0,
	190, 0, 0, 193, 194, 175, 176, 11, 0, 182,
	0, 202, 0, 79, 49, 150, 50, 85, 76, 107,
	109, 115, 79, 92, 93, 94, 97, 98, 99, 79,
	79, 0, 155, 0, 0, 0, 12, 79, 79, 172,
	184, 186, 0, 195, 196, 191, 192, 177, 0, 179,
	207, 31, 32, 114, 0, 0, 102, 103, 104, 105,
	206, 201, 203, 205, 185, 79, 0, 79, 0, 0,
	183, 13, 0, 0, 79, 197, 79, 187, 0, 0,
	100, 101, 204, 0, 188, 189, 198,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 114, 3, 3, 3, 109, 102, 3,
	123, 121, 107, 105, 65, 106, 120, 108, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	88, 87, 89, 3, 124, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 112, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 101, 3, 113,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
	52, 53, 54, 55, 56, 57, 58, 59, 60, 61,
	62, 63, 64, 66, 67, 68, 69, 70, 71, 72,
	73, 74, 75, 76, 77, 78, 79, 80, 81, 82,
	83, 84, 85, 86, 90, 91, 92, 93, 94, 95,
	96, 97, 98, 99, 100, 103, 104, 110, 111, 115,
	116, 117, 118, 119, 122,
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

	res := "syntax error: unexpected " + yyCurrentToken

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
		//line ./advisor/parser/parser.y:216
		{
			yylex.(*Tokener).ParseStmt = yyDollar[1].node
		}
	case 4:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./advisor/parser/parser.y:235
		{
			yyVAL.node = yyDollar[2].node
		}
	case 5:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./advisor/parser/parser.y:239
		{
			yyVAL.node = yyDollar[2].node
		}
	case 6:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./advisor/parser/parser.y:245
		{
			yyVAL.node = yyDollar[1].node
		}
	case 7:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./advisor/parser/parser.y:249
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
		//line ./advisor/parser/parser.y:260
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
		//line ./advisor/parser/parser.y:275
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
		//line ./advisor/parser/parser.y:293
		{
			yyVAL.node = yyDollar[1].node
		}
	case 11:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./advisor/parser/parser.y:294
		{
			yyVAL.node = yyDollar[1].node
		}
	case 12:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line ./advisor/parser/parser.y:299
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
		//line ./advisor/parser/parser.y:310
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
		//line ./advisor/parser/parser.y:322
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
		//line ./advisor/parser/parser.y:332
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
		//line ./advisor/parser/parser.y:342
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
		//line ./advisor/parser/parser.y:362
		{
			expr := &ast.And0OrExpr{Operator: "OR", Expr: []ast.ExprNode{}}
			expr.SetTag(ast.Expr_Or)
			expr.Combine(yyDollar[1].expr, yyDollar[3].expr)
			yyVAL.expr = expr

		}
	case 18:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./advisor/parser/parser.y:370
		{
			expr := &ast.And0OrExpr{Operator: "OR", Expr: []ast.ExprNode{}}
			expr.SetTag(ast.Expr_Or)
			expr.Combine(yyDollar[1].expr, yyDollar[3].expr)
			yyVAL.expr = expr
		}
	case 19:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./advisor/parser/parser.y:376
		{
			yyVAL.expr = &ast.Expr{Operator: "XOR", Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 20:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./advisor/parser/parser.y:378
		{
			expr := &ast.And0OrExpr{Operator: "AND", Expr: []ast.ExprNode{}}
			expr.SetTag(ast.Expr_And)
			expr.Combine(yyDollar[1].expr, yyDollar[3].expr)
			yyVAL.expr = expr
		}
	case 21:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./advisor/parser/parser.y:385
		{
			expr := &ast.And0OrExpr{Operator: "AND", Expr: []ast.ExprNode{}}
			expr.SetTag(ast.Expr_And)
			expr.Combine(yyDollar[1].expr, yyDollar[3].expr)
			yyVAL.expr = expr
		}
	case 22:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./advisor/parser/parser.y:391
		{
			yyVAL.expr = &ast.NotExpr{Expr: yyDollar[2].expr}
		}
	case 23:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./advisor/parser/parser.y:392
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 24:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./advisor/parser/parser.y:394
		{
			yyVAL.expr = &ast.IsTrueExpr{HasNot: yyDollar[2].boolean, IsTrue: yyDollar[3].boolean, Expr: yyDollar[1].expr}
		}
	case 25:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./advisor/parser/parser.y:400
		{
			yyVAL.expr = &ast.IsNullExpr{HasNot: yyDollar[2].boolean, Left: yyDollar[1].expr}
		}
	case 26:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./advisor/parser/parser.y:404
		{
			yyVAL.expr = &ast.Expr{Operator: "=", Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 27:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./advisor/parser/parser.y:408
		{
			yyVAL.expr = &ast.Expr{Operator: yyDollar[2].val, Left: yyDollar[1].expr, Right: yyDollar[3].expr}

		}
	case 28:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./advisor/parser/parser.y:413
		{
			expr := &ast.SubqueryExpr{Operator: yyDollar[2].val, Subtype: yyDollar[3].tag, Left: yyDollar[1].expr, Select: yyDollar[4].node}
			expr.SetTag(ast.Subquery_Operator)
			yyVAL.expr = expr
		}
	case 29:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./advisor/parser/parser.y:419
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 30:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./advisor/parser/parser.y:425
		{
			expr := &ast.SubqueryExpr{HasNot: yyDollar[2].boolean, Left: yyDollar[1].expr, Select: yyDollar[4].node}
			expr.SetTag(ast.Subquery_In)
			yyVAL.expr = expr
		}
	case 31:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./advisor/parser/parser.y:431
		{
			yyVAL.expr = &ast.InExpr{HasNot: yyDollar[2].boolean, Left: yyDollar[1].expr, Right: yyDollar[5].list}
		}
	case 32:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./advisor/parser/parser.y:435
		{
			yyVAL.expr = &ast.BetweenExpr{HasNot: yyDollar[2].boolean, Expr: yyDollar[1].expr, Left: yyDollar[4].expr, Right: yyDollar[6].expr}
		}
	case 33:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./advisor/parser/parser.y:439
		{
			yyVAL.expr = &ast.LikeExpr{HasNot: false, Left: yyDollar[1].expr, Right: yyDollar[4].expr}
		}
	case 34:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./advisor/parser/parser.y:443
		{
			yyVAL.expr = &ast.LikeExpr{HasNot: yyDollar[2].boolean, Left: yyDollar[1].expr, Right: yyDollar[4].expr}
		}
	case 35:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./advisor/parser/parser.y:447
		{
			yyVAL.expr = &ast.RegexpExpr{HasNot: yyDollar[2].boolean, Left: yyDollar[1].expr, Right: yyDollar[4].expr}
		}
	case 36:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./advisor/parser/parser.y:451
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 37:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./advisor/parser/parser.y:457
		{
			yyVAL.expr = &ast.Expr{Operator: "|", Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 38:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./advisor/parser/parser.y:458
		{
			yyVAL.expr = &ast.Expr{Operator: "&", Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 39:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./advisor/parser/parser.y:459
		{
			yyVAL.expr = &ast.Expr{Operator: "<<", Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 40:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./advisor/parser/parser.y:460
		{
			yyVAL.expr = &ast.Expr{Operator: ">>", Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 41:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./advisor/parser/parser.y:461
		{
			yyVAL.expr = &ast.Expr{Operator: "+", Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 42:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./advisor/parser/parser.y:462
		{
			yyVAL.expr = &ast.Expr{Operator: "-", Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 43:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./advisor/parser/parser.y:463
		{
			yyVAL.expr = &ast.Expr{Operator: "*", Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 44:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./advisor/parser/parser.y:464
		{
			yyVAL.expr = &ast.Expr{Operator: "/", Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 45:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./advisor/parser/parser.y:465
		{
			yyVAL.expr = &ast.Expr{Operator: "DIV", Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 46:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./advisor/parser/parser.y:466
		{
			yyVAL.expr = &ast.Expr{Operator: "MOD", Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 47:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./advisor/parser/parser.y:467
		{
			yyVAL.expr = &ast.Expr{Operator: "%", Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 48:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./advisor/parser/parser.y:468
		{
			yyVAL.expr = &ast.Expr{Operator: "^", Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 49:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./advisor/parser/parser.y:470
		{
			yyVAL.expr = &ast.IntervalExpr{Operator: ast.Op_Add, Left: yyDollar[1].expr, Right: yyDollar[4].expr, TimeUnit: yyDollar[5].tag}
		}
	case 50:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./advisor/parser/parser.y:474
		{
			yyVAL.expr = &ast.IntervalExpr{Operator: ast.Op_Minus, Left: yyDollar[1].expr, Right: yyDollar[4].expr, TimeUnit: yyDollar[5].tag}
		}
	case 51:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./advisor/parser/parser.y:478
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 52:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./advisor/parser/parser.y:483
		{
			yyVAL.expr = yyDollar[2].expr
		}
	case 53:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./advisor/parser/parser.y:484
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 54:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./advisor/parser/parser.y:485
		{
			yyVAL.expr = yyDollar[2].expr
		}
	case 55:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./advisor/parser/parser.y:486
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 56:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./advisor/parser/parser.y:487
		{
			yyVAL.expr = yyDollar[2].expr
		}
	case 57:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./advisor/parser/parser.y:489
		{
			yyVAL.expr = &ast.CollateExpr{Expr: yyDollar[1].expr, Collate: yyDollar[3].tag}
		}
	case 58:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./advisor/parser/parser.y:493
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "+", Expr: yyDollar[2].expr}
		}
	case 59:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./advisor/parser/parser.y:497
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "-", Expr: yyDollar[2].expr}
		}
	case 60:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./advisor/parser/parser.y:501
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "~", Expr: yyDollar[2].expr}
		}
	case 61:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./advisor/parser/parser.y:505
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "!", Expr: yyDollar[2].expr}
		}
	case 62:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./advisor/parser/parser.y:509
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
		//line ./advisor/parser/parser.y:520
		{
			yyVAL.expr = yyDollar[1].node.(ast.ExprNode)
		}
	case 64:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./advisor/parser/parser.y:524
		{
			expr := &ast.SubqueryExpr{Select: yyDollar[2].node}
			expr.SetTag(ast.Subquery_Exists)
			yyVAL.expr = expr
		}
	case 65:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./advisor/parser/parser.y:530
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 66:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./advisor/parser/parser.y:536
		{
			Expr := &ast.Integer{Operator: []string{}, Val: yyDollar[1].val}
			Expr.SetTag(ast.Bit)
			yyVAL.expr = Expr
		}
	case 67:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./advisor/parser/parser.y:542
		{
			Expr := &ast.Integer{Operator: []string{}, Val: yyDollar[1].val}
			Expr.SetTag(ast.Hex)
			yyVAL.expr = Expr
		}
	case 68:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./advisor/parser/parser.y:548
		{
			yyVAL.expr = &ast.Integer{Operator: []string{}, Val: yyDollar[1].val}
		}
	case 69:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./advisor/parser/parser.y:552
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
		//line ./advisor/parser/parser.y:560
		{
			yyVAL.expr = &ast.Marker{Str: yyDollar[1].str}
		}
	case 71:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./advisor/parser/parser.y:562
		{
			str := &ast.String{Operator: []string{}, Str: yyDollar[1].str}
			str.SetTag(ast.String_single)
			yyVAL.expr = str
		}
	case 72:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./advisor/parser/parser.y:568
		{
			str := &ast.String{Operator: []string{}, Str: yyDollar[1].str}
			str.SetTag(ast.String_double)
			yyVAL.expr = str
		}
	case 73:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./advisor/parser/parser.y:573
		{
			yyVAL.expr = &ast.True{IsTrue: yyDollar[1].boolean}
		}
	case 74:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./advisor/parser/parser.y:574
		{
			yyVAL.expr = &ast.Null{}
		}
	case 75:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./advisor/parser/parser.y:578
		{
			yyVAL.val = yyDollar[1].ident
		}
	case 76:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./advisor/parser/parser.y:582
		{
			yyVAL.val = yyDollar[1].val
		}
	case 77:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./advisor/parser/parser.y:586
		{
			yyVAL.val = "*"
		}
	case 78:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./advisor/parser/parser.y:592
		{
			yyVAL.val = ""
		}
	case 79:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./advisor/parser/parser.y:596
		{
			yyVAL.val = ""
		}
	case 80:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./advisor/parser/parser.y:601
		{
			yyVAL.expr = &ast.UserVariable{Variable: yyDollar[1].val}
		}
	case 81:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./advisor/parser/parser.y:602
		{
			yyVAL.expr = &ast.SystemVariable{Schema: yyDollar[1].val}
		}
	case 82:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./advisor/parser/parser.y:604
		{
			yyVAL.expr = &ast.SystemVariable{Schema: yyDollar[1].val, Parameter: yyDollar[3].val}
		}
	case 83:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./advisor/parser/parser.y:610
		{
			if yyDollar[1].val == "*" {
				yyVAL.expr = &ast.Star{Column: []*ast.Column{}}
			} else {
				yyVAL.expr = &ast.Column{Column: yyDollar[1].val}
			}
		}
	case 84:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./advisor/parser/parser.y:618
		{
			if yyDollar[3].val == "*" {
				yyVAL.expr = &ast.Star{Table: yyDollar[1].val, Column: []*ast.Column{}}
			} else {
				yyVAL.expr = &ast.Column{Table: yyDollar[1].val, Column: yyDollar[3].val}
			}
		}
	case 85:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./advisor/parser/parser.y:626
		{
			if yyDollar[5].val == "*" {
				yyVAL.expr = &ast.Star{DB: yyDollar[1].val, Table: yyDollar[3].val, Column: []*ast.Column{}}
			} else {
				yyVAL.expr = &ast.Column{DB: yyDollar[1].val, Table: yyDollar[3].val, Column: yyDollar[5].val}
			}
		}
	case 86:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./advisor/parser/parser.y:637
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 87:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./advisor/parser/parser.y:641
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 88:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./advisor/parser/parser.y:645
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 89:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./advisor/parser/parser.y:649
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 90:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./advisor/parser/parser.y:653
		{
			yyVAL.expr = &ast.FuncCall{Name: yyDollar[1].val}
		}
	case 91:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./advisor/parser/parser.y:657
		{
			yyVAL.expr = &ast.FuncCall{Name: yyDollar[1].val, Arg: yyDollar[3].list}
		}
	case 92:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./advisor/parser/parser.y:663
		{
			expr := &ast.AggregationFuncCall{DistinctType: yyDollar[3].tag, Arg: []ast.ExprNode{yyDollar[4].expr}}
			expr.SetTag(ast.Aggregation_Avg)
			yyVAL.expr = expr
		}
	case 93:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./advisor/parser/parser.y:669
		{
			expr := &ast.AggregationFuncCall{DistinctType: yyDollar[3].tag, Arg: yyDollar[4].list}
			expr.SetTag(ast.Aggregation_Count)
			yyVAL.expr = expr
		}
	case 94:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./advisor/parser/parser.y:675
		{
			expr := &ast.AggregationFuncCall{DistinctType: ast.AST_ALL, Arg: []ast.ExprNode{yyDollar[4].expr}}
			expr.SetTag(ast.Aggregation_Count)
			yyVAL.expr = expr
		}
	case 95:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./advisor/parser/parser.y:681
		{
			expr := &ast.AggregationFuncCall{DistinctType: ast.AST_EMPTY, Arg: []ast.ExprNode{yyDollar[3].expr}}
			expr.SetTag(ast.Aggregation_Count)
			yyVAL.expr = expr
		}
	case 96:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./advisor/parser/parser.y:687
		{
			expr := &ast.AggregationFuncCall{DistinctType: ast.AST_EMPTY, Arg: []ast.ExprNode{&ast.Star{Column: []*ast.Column{}}}}
			expr.SetTag(ast.Aggregation_Count)
			yyVAL.expr = expr
		}
	case 97:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./advisor/parser/parser.y:693
		{
			expr := &ast.AggregationFuncCall{DistinctType: yyDollar[3].tag, Arg: []ast.ExprNode{yyDollar[4].expr}}
			expr.SetTag(ast.Aggregation_Max)
			yyVAL.expr = expr
		}
	case 98:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./advisor/parser/parser.y:699
		{
			expr := &ast.AggregationFuncCall{DistinctType: yyDollar[3].tag, Arg: []ast.ExprNode{yyDollar[4].expr}}
			expr.SetTag(ast.Aggregation_Min)
			yyVAL.expr = expr
		}
	case 99:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./advisor/parser/parser.y:705
		{
			expr := &ast.AggregationFuncCall{DistinctType: yyDollar[3].tag, Arg: []ast.ExprNode{yyDollar[4].expr}}
			expr.SetTag(ast.Aggregation_Sum)
			yyVAL.expr = expr
		}
	case 100:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line ./advisor/parser/parser.y:713
		{
			Expr := &ast.CalTimeFuncCall{Interval: &ast.IntervalExpr{Operator: ast.Op_Add, Left: yyDollar[3].expr, Right: yyDollar[6].expr, TimeUnit: yyDollar[7].tag}}
			Expr.SetTag(yyDollar[1].tag)
			yyVAL.expr = Expr
		}
	case 101:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line ./advisor/parser/parser.y:719
		{
			Expr := &ast.CalTimeFuncCall{Interval: &ast.IntervalExpr{Operator: ast.Op_Minus, Left: yyDollar[3].expr, Right: yyDollar[6].expr, TimeUnit: yyDollar[7].tag}}
			Expr.SetTag(yyDollar[1].tag)
			yyVAL.expr = Expr
		}
	case 102:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./advisor/parser/parser.y:727
		{
			yyVAL.expr = &ast.CastFuncCall{Expr: yyDollar[3].expr, CastType: yyDollar[5].val}
		}
	case 103:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./advisor/parser/parser.y:731
		{
			yyVAL.expr = &ast.CastFuncCall{Expr: yyDollar[3].expr, CastType: yyDollar[5].val}
		}
	case 104:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./advisor/parser/parser.y:735
		{
			yyVAL.expr = &ast.CastFuncCall{Expr: yyDollar[3].expr, CastType: yyDollar[5].val}
		}
	case 105:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./advisor/parser/parser.y:741
		{
			yyVAL.expr = &ast.StringFuncCall{Substr: yyDollar[3].expr, Str: yyDollar[5].expr}
		}
	case 106:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./advisor/parser/parser.y:746
		{
			yyVAL.list = []ast.ExprNode{yyDollar[1].expr}
		}
	case 107:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./advisor/parser/parser.y:747
		{
			yyVAL.list = append(yyDollar[1].list, yyDollar[3].expr)
		}
	case 108:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./advisor/parser/parser.y:750
		{
			yyVAL.expr = &ast.Marker{Str: yyDollar[1].str}
		}
	case 109:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./advisor/parser/parser.y:754
		{
			yyVAL.expr = &ast.CaseExpr{Case: yyDollar[2].expr, When: yyDollar[3].list, Else: yyDollar[4].expr}
		}
	case 110:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./advisor/parser/parser.y:759
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 111:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./advisor/parser/parser.y:760
		{
			yyVAL.expr = nil
		}
	case 112:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./advisor/parser/parser.y:763
		{
			yyVAL.list = []ast.ExprNode{yyDollar[1].expr}
		}
	case 113:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./advisor/parser/parser.y:764
		{
			yyVAL.list = append(yyDollar[1].list, yyDollar[2].expr)
		}
	case 114:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./advisor/parser/parser.y:767
		{
			yyVAL.expr = &ast.Expr{Left: yyDollar[2].expr, Right: yyDollar[4].expr}
		}
	case 115:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./advisor/parser/parser.y:770
		{
			yyVAL.expr = &ast.Expr{Left: yyDollar[2].expr}
		}
	case 116:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./advisor/parser/parser.y:771
		{
			yyVAL.expr = nil
		}
	case 117:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./advisor/parser/parser.y:774
		{
			yyVAL.list = yyDollar[1].list
		}
	case 118:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./advisor/parser/parser.y:775
		{
			yyVAL.list = nil
		}
	case 119:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./advisor/parser/parser.y:778
		{
			yyVAL.list = []ast.ExprNode{yyDollar[1].expr}
		}
	case 120:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./advisor/parser/parser.y:779
		{
			yyVAL.list = append(yyDollar[1].list, yyDollar[3].expr)
		}
	case 121:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./advisor/parser/parser.y:783
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
	case 122:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./advisor/parser/parser.y:805
		{
			yyVAL.val = yyDollar[2].val
		}
	case 123:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./advisor/parser/parser.y:808
		{
			yyVAL.val = yyDollar[3].val
		}
	case 124:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./advisor/parser/parser.y:811
		{
			yyVAL.val = yyDollar[1].val
		}
	case 125:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./advisor/parser/parser.y:812
		{
			yyVAL.val = yyDollar[1].val
		}
	case 126:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./advisor/parser/parser.y:815
		{
			yyVAL.val = yyDollar[1].val
		}
	case 127:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./advisor/parser/parser.y:816
		{
			yyVAL.val = ""
		}
	case 128:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./advisor/parser/parser.y:819
		{
			yyVAL.val = yyDollar[2].val
		}
	case 129:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./advisor/parser/parser.y:820
		{
			yyVAL.val = yyDollar[1].val
		}
	case 130:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./advisor/parser/parser.y:823
		{
			yyVAL.val = yyDollar[1].val
		}
	case 131:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./advisor/parser/parser.y:824
		{
			yyVAL.val = yyDollar[1].str
		}
	case 132:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./advisor/parser/parser.y:825
		{
			yyVAL.val = yyDollar[1].str
		}
	case 133:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./advisor/parser/parser.y:826
		{
			yyVAL.val = yyDollar[1].str
		}
	case 134:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./advisor/parser/parser.y:830
		{
			yyVAL.boolean = true
		}
	case 135:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./advisor/parser/parser.y:831
		{
			yyVAL.boolean = false
		}
	case 136:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./advisor/parser/parser.y:836
		{
			yyVAL.tag = ast.Operator_Some
		}
	case 137:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./advisor/parser/parser.y:837
		{
			yyVAL.tag = ast.Operator_All
		}
	case 138:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./advisor/parser/parser.y:838
		{
			yyVAL.tag = ast.Operator_Any
		}
	case 139:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./advisor/parser/parser.y:841
		{
			yyVAL.val = "="
		}
	case 140:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./advisor/parser/parser.y:842
		{
			yyVAL.val = ">="
		}
	case 141:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./advisor/parser/parser.y:843
		{
			yyVAL.val = ">"
		}
	case 142:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./advisor/parser/parser.y:844
		{
			yyVAL.val = "<="
		}
	case 143:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./advisor/parser/parser.y:845
		{
			yyVAL.val = "<"
		}
	case 144:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./advisor/parser/parser.y:846
		{
			yyVAL.val = "<>"
		}
	case 145:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./advisor/parser/parser.y:847
		{
			yyVAL.val = "!="
		}
	case 146:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./advisor/parser/parser.y:850
		{
			yyVAL.boolean = false
		}
	case 147:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./advisor/parser/parser.y:851
		{
			yyVAL.boolean = true
		}
	case 148:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./advisor/parser/parser.y:855
		{
			yyVAL.boolean = true
		}
	case 149:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./advisor/parser/parser.y:856
		{
			yyVAL.boolean = false
		}
	case 150:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./advisor/parser/parser.y:859
		{
			yyVAL.tag = ast.TimeUnit[yyDollar[1].ident]
		}
	case 151:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./advisor/parser/parser.y:863
		{
			yyVAL.tag = ast.Character[yyDollar[1].ident]
		}
	case 152:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./advisor/parser/parser.y:867
		{
			yyVAL.val = "LEFT"
		}
	case 153:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./advisor/parser/parser.y:868
		{
			yyVAL.val = "IN"
		}
	case 154:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./advisor/parser/parser.y:869
		{
			yyVAL.val = "IS"
		}
	case 155:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./advisor/parser/parser.y:872
		{
			yyVAL.val = yyDollar[1].val
		}
	case 156:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./advisor/parser/parser.y:875
		{
			yyVAL.tag = ast.FuncTime[yyDollar[1].ident]
		}
	case 157:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./advisor/parser/parser.y:878
		{
			yyVAL.tag = ast.FuncTime[yyDollar[1].ident]
		}
	case 158:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./advisor/parser/parser.y:889
		{
			clause := &ast.DistinctClause{}
			clause.SetTag(yyDollar[1].tag)
			yyVAL.node = clause
		}
	case 159:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./advisor/parser/parser.y:896
		{
			yyVAL.tag = ast.AST_ALL
		}
	case 160:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./advisor/parser/parser.y:897
		{
			yyVAL.tag = yyDollar[1].tag
		}
	case 161:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./advisor/parser/parser.y:898
		{
			yyVAL.tag = ast.AST_EMPTY
		}
	case 162:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./advisor/parser/parser.y:901
		{
			yyVAL.tag = ast.AST_DISTINCT
		}
	case 163:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./advisor/parser/parser.y:902
		{
			yyVAL.tag = ast.AST_DISTINCTROW
		}
	case 164:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./advisor/parser/parser.y:912
		{
			yyVAL.node = &ast.IntoClause{Relation: &ast.Relation{Table: yyDollar[2].val}}
		}
	case 165:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./advisor/parser/parser.y:914
		{
			clause := &ast.IntoClause{}
			clause.SetTag(ast.AST_EMPTY)
			yyVAL.node = clause
		}
	case 166:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./advisor/parser/parser.y:921
		{
			yyVAL.val = yyDollar[1].val
		}
	case 167:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./advisor/parser/parser.y:922
		{
			yyVAL.val = yyDollar[1].str
		}
	case 168:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./advisor/parser/parser.y:923
		{
			yyVAL.val = yyDollar[1].str
		}
	case 169:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./advisor/parser/parser.y:933
		{
			yyVAL.node = &ast.FromClause{Table_ref: yyDollar[2].list}
		}
	case 170:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./advisor/parser/parser.y:935
		{
			clause := &ast.FromClause{}
			clause.SetTag(ast.AST_EMPTY)
			yyVAL.node = clause
		}
	case 171:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./advisor/parser/parser.y:943
		{
			yyVAL.list = []ast.ExprNode{yyDollar[1].expr}
		}
	case 172:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./advisor/parser/parser.y:944
		{
			yyVAL.list = append(yyDollar[1].list, yyDollar[3].expr)
		}
	case 173:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./advisor/parser/parser.y:947
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 174:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./advisor/parser/parser.y:948
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 175:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./advisor/parser/parser.y:952
		{
			yyDollar[1].expr.(*ast.Relation).Alias = yyDollar[2].val
			yyVAL.expr = yyDollar[1].expr
		}
	case 176:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./advisor/parser/parser.y:957
		{
			tb := &ast.Relation{Select: yyDollar[1].node, Alias: yyDollar[2].val}
			tb.SetTag(ast.Relation_Subquery)
			yyVAL.expr = tb
		}
	case 177:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./advisor/parser/parser.y:963
		{
			switch node := yyDollar[2].expr.(type) {
			case *ast.Join:
				node.IsEnclosed = true
			}
			yyVAL.expr = yyDollar[2].expr
		}
	case 178:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./advisor/parser/parser.y:973
		{
			tb := &ast.Relation{Table: yyDollar[1].val}
			tb.SetTag(ast.Relation_Table)
			yyVAL.expr = tb
		}
	case 179:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./advisor/parser/parser.y:979
		{
			tb := &ast.Relation{DB: yyDollar[1].val, Table: yyDollar[3].val}
			tb.SetTag(ast.Relation_Table)
			yyVAL.expr = tb
		}
	case 180:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./advisor/parser/parser.y:986
		{
			yyVAL.val = yyDollar[1].val
		}
	case 181:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./advisor/parser/parser.y:987
		{
			yyVAL.val = yyDollar[1].str
		}
	case 182:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./advisor/parser/parser.y:990
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 183:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./advisor/parser/parser.y:992
		{
			rel := &ast.Join{Left: yyDollar[1].expr, Right: yyDollar[3].expr}
			rel.SetTag(ast.Cross_join)
			yyVAL.expr = rel
		}
	case 184:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./advisor/parser/parser.y:1003
		{
			rel := &ast.Join{Left: yyDollar[1].expr, Right: yyDollar[3].expr}
			rel.SetTag(ast.Cross_join)
			yyVAL.expr = rel
		}
	case 185:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./advisor/parser/parser.y:1009
		{
			rel := &ast.Join{Left: yyDollar[1].expr, Right: yyDollar[3].expr, Condition: yyDollar[4].expr}
			rel.SetTag(ast.Inner_join)
			yyVAL.expr = rel
		}
	case 186:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./advisor/parser/parser.y:1015
		{
			rel := &ast.Join{Left: yyDollar[1].expr, Right: yyDollar[3].expr}
			rel.SetTag(ast.Straight_join)
			yyVAL.expr = rel
		}
	case 187:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./advisor/parser/parser.y:1021
		{
			rel := &ast.Join{Left: yyDollar[1].expr, Right: yyDollar[3].expr, Condition: yyDollar[5].expr}
			rel.SetTag(ast.Straight_join)
			yyVAL.expr = rel
		}
	case 188:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./advisor/parser/parser.y:1027
		{
			rel := &ast.Join{HasOuter: yyDollar[3].boolean, Left: yyDollar[1].expr, Right: yyDollar[5].expr, Condition: yyDollar[6].expr}
			if yyDollar[2].boolean {
				rel.SetTag(ast.Left_join)
			} else {
				rel.SetTag(ast.Right_join)
			}
			yyVAL.expr = rel
		}
	case 189:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./advisor/parser/parser.y:1037
		{
			rel := &ast.Join{HasNatural: true, HasOuter: yyDollar[4].boolean, Left: yyDollar[1].expr, Right: yyDollar[6].expr}
			if yyDollar[3].boolean {
				rel.SetTag(ast.Left_join)
			} else {
				rel.SetTag(ast.Right_join)
			}
			yyVAL.expr = rel
		}
	case 190:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./advisor/parser/parser.y:1048
		{
			yyVAL.tag = 1
		}
	case 191:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./advisor/parser/parser.y:1049
		{
			yyVAL.tag = 2
		}
	case 192:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./advisor/parser/parser.y:1050
		{
			yyVAL.tag = 3
		}
	case 193:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./advisor/parser/parser.y:1054
		{
			yyVAL.boolean = true
		}
	case 194:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./advisor/parser/parser.y:1055
		{
			yyVAL.boolean = false
		}
	case 195:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./advisor/parser/parser.y:1058
		{
			yyVAL.boolean = true
		}
	case 196:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./advisor/parser/parser.y:1059
		{
			yyVAL.boolean = false
		}
	case 197:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./advisor/parser/parser.y:1062
		{
			yyVAL.expr = yyDollar[2].expr
		}
	case 198:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./advisor/parser/parser.y:1063
		{
			yyVAL.expr = &ast.Using{Column: yyDollar[3].list}
		}
	case 199:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./advisor/parser/parser.y:1074
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
	case 200:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./advisor/parser/parser.y:1087
		{
			clause := &ast.WhereClause{}
			clause.SetTag(ast.AST_EMPTY)
			yyVAL.node = clause
		}
	case 201:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./advisor/parser/parser.y:1100
		{
			yyVAL.node = &ast.GroupClause{Target_ref: yyDollar[3].list}
		}
	case 202:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./advisor/parser/parser.y:1102
		{
			clause := &ast.GroupClause{}
			clause.SetTag(ast.AST_EMPTY)
			yyVAL.node = clause
		}
	case 203:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./advisor/parser/parser.y:1110
		{
			yyVAL.list = []ast.ExprNode{yyDollar[1].expr}
		}
	case 204:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./advisor/parser/parser.y:1114
		{
			yyVAL.list = append(yyDollar[1].list, yyDollar[3].expr)
		}
	case 205:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./advisor/parser/parser.y:1120
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 206:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./advisor/parser/parser.y:1129
		{
			yyVAL.node = &ast.HavingClause{Having: yyDollar[2].expr}
		}
	case 207:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./advisor/parser/parser.y:1131
		{
			clause := &ast.HavingClause{}
			clause.SetTag(ast.AST_EMPTY)
			yyVAL.node = clause
		}
	case 208:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./advisor/parser/parser.y:1145
		{
			yyVAL.node = yyDollar[1].node
		}
	case 209:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./advisor/parser/parser.y:1147
		{
			clause := &ast.SortClause{}
			clause.SetTag(ast.AST_EMPTY)
			yyVAL.node = clause
		}
	case 210:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./advisor/parser/parser.y:1154
		{
			yyVAL.node = &ast.SortClause{Target_ref: yyDollar[3].list}
		}
	case 211:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./advisor/parser/parser.y:1157
		{
			yyVAL.list = []ast.ExprNode{yyDollar[1].expr}
		}
	case 212:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./advisor/parser/parser.y:1158
		{
			yyVAL.list = append(yyDollar[1].list, yyDollar[3].expr)
		}
	case 213:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./advisor/parser/parser.y:1161
		{
			yyVAL.expr = &ast.Sortby{Expr: yyDollar[1].expr, AscType: yyDollar[2].tag}
		}
	case 214:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./advisor/parser/parser.y:1164
		{
			yyVAL.tag = ast.Sort_Asc
		}
	case 215:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./advisor/parser/parser.y:1165
		{
			yyVAL.tag = ast.Sort_Desc
		}
	case 216:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./advisor/parser/parser.y:1166
		{
			yyVAL.tag = ast.Sort_Empty
		}
	case 217:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./advisor/parser/parser.y:1175
		{
			yyVAL.node = yyDollar[1].node
		}
	case 218:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./advisor/parser/parser.y:1177
		{
			clause := &ast.LockClause{}
			clause.SetTag(ast.AST_EMPTY)
			yyVAL.node = clause
		}
	case 219:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./advisor/parser/parser.y:1185
		{
			lock := &ast.LockClause{}
			lock.SetTag(ast.Lock_update)
			yyVAL.node = lock
		}
	case 220:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./advisor/parser/parser.y:1191
		{
			lock := &ast.LockClause{}
			lock.SetTag(ast.Lock_share)
			yyVAL.node = lock
		}
	case 221:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./advisor/parser/parser.y:1205
		{
			yyVAL.node = yyDollar[1].node
		}
	case 222:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./advisor/parser/parser.y:1207
		{
			clause := &ast.LimitClause{}
			clause.SetTag(ast.AST_EMPTY)
			yyVAL.node = clause
		}
	case 223:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./advisor/parser/parser.y:1214
		{
			yyVAL.node = &ast.LimitClause{Limit: yyDollar[1].list, Offset: yyDollar[2].expr}
		}
	case 224:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./advisor/parser/parser.y:1215
		{
			yyVAL.node = &ast.LimitClause{Limit: yyDollar[2].list, Offset: yyDollar[1].expr}
		}
	case 225:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./advisor/parser/parser.y:1216
		{
			yyVAL.node = &ast.LimitClause{Limit: yyDollar[1].list}
		}
	case 226:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./advisor/parser/parser.y:1217
		{
			yyVAL.node = &ast.LimitClause{Offset: yyDollar[1].expr}
		}
	case 227:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./advisor/parser/parser.y:1223
		{
			yyVAL.list = yyDollar[2].list
		}
	case 228:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./advisor/parser/parser.y:1227
		{
			yyVAL.list = append(yyDollar[2].list, yyDollar[4].expr)
		}
	case 229:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./advisor/parser/parser.y:1232
		{
			yyVAL.list = []ast.ExprNode{yyDollar[1].expr}
		}
	case 230:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./advisor/parser/parser.y:1235
		{
			yyVAL.expr = yyDollar[2].expr
		}
	case 231:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./advisor/parser/parser.y:1238
		{
			yyVAL.expr = yyDollar[1].expr
		}
	}
	goto yystack /* stack new state and value */
}
