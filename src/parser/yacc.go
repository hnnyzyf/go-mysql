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
	-2, 203,
	-1, 28,
	78, 144,
	87, 144,
	88, 144,
	89, 144,
	-2, 36,
	-1, 52,
	115, 119,
	-2, 76,
	-1, 139,
	29, 11,
	35, 11,
	36, 11,
	41, 11,
	42, 11,
	58, 11,
	59, 11,
	60, 11,
	-2, 64,
	-1, 261,
	117, 175,
	-2, 161,
	-1, 262,
	117, 176,
	-2, 162,
}

const yyPrivate = 57344

const yyLast = 1165

var yyAct = [...]int{

	226, 101, 383, 40, 4, 254, 357, 225, 359, 24,
	52, 16, 316, 334, 27, 340, 259, 294, 253, 24,
	300, 51, 99, 250, 98, 173, 106, 165, 233, 83,
	29, 62, 325, 312, 273, 102, 143, 81, 20, 142,
	140, 404, 175, 139, 275, 141, 397, 390, 62, 180,
	84, 85, 262, 275, 346, 275, 389, 275, 388, 147,
	93, 94, 96, 97, 287, 88, 135, 136, 137, 138,
	8, 62, 180, 87, 148, 262, 400, 119, 393, 372,
	164, 166, 351, 278, 8, 169, 95, 174, 169, 169,
	277, 28, 266, 24, 183, 184, 185, 186, 187, 387,
	7, 178, 188, 333, 154, 153, 132, 177, 349, 345,
	329, 102, 274, 152, 151, 150, 144, 62, 104, 133,
	105, 103, 131, 191, 192, 122, 123, 100, 49, 62,
	120, 121, 124, 125, 126, 127, 130, 128, 129, 131,
	315, 7, 95, 149, 124, 125, 126, 127, 130, 128,
	129, 131, 239, 222, 223, 258, 93, 94, 96, 97,
	235, 122, 123, 21, 221, 219, 120, 121, 124, 125,
	126, 127, 130, 128, 129, 131, 96, 97, 258, 89,
	257, 237, 95, 160, 199, 234, 202, 261, 236, 134,
	241, 242, 243, 200, 201, 198, 234, 264, 281, 332,
	95, 197, 265, 126, 127, 130, 128, 129, 131, 271,
	78, 272, 203, 204, 205, 206, 207, 209, 211, 212,
	213, 214, 215, 216, 263, 88, 61, 362, 133, 73,
	341, 268, 77, 305, 306, 282, 189, 283, 375, 285,
	363, 344, 288, 289, 290, 284, 166, 174, 343, 370,
	74, 12, 296, 297, 246, 245, 182, 93, 94, 96,
	97, 280, 309, 93, 94, 96, 97, 102, 102, 261,
	139, 181, 293, 292, 314, 92, 321, 311, 323, 324,
	308, 307, 326, 95, 320, 318, 369, 313, 161, 95,
	5, 267, 90, 269, 270, 319, 399, 368, 392, 322,
	331, 257, 257, 257, 158, 339, 330, 6, 261, 261,
	261, 162, 93, 94, 96, 97, 337, 338, 251, 92,
	102, 159, 342, 178, 156, 244, 80, 79, 354, 347,
	350, 348, 308, 80, 122, 123, 355, 358, 95, 120,
	121, 124, 125, 126, 127, 130, 128, 129, 131, 291,
	257, 79, 367, 75, 90, 328, 295, 261, 364, 335,
	76, 371, 366, 373, 176, 365, 62, 257, 336, 377,
	378, 358, 168, 60, 261, 168, 168, 379, 86, 257,
	380, 382, 374, 381, 19, 190, 261, 384, 218, 167,
	385, 386, 170, 171, 394, 398, 59, 391, 396, 230,
	231, 401, 276, 91, 403, 145, 384, 14, 1, 93,
	94, 96, 97, 102, 62, 48, 57, 47, 46, 43,
	44, 45, 238, 11, 13, 12, 65, 15, 3, 232,
	66, 82, 22, 23, 60, 95, 93, 94, 96, 97,
	310, 22, 23, 67, 68, 172, 50, 402, 93, 94,
	96, 97, 286, 163, 17, 69, 356, 59, 248, 62,
	180, 252, 95, 179, 249, 110, 111, 115, 113, 32,
	70, 227, 112, 114, 95, 116, 256, 117, 395, 220,
	108, 34, 30, 25, 255, 58, 33, 279, 31, 72,
	93, 94, 96, 97, 317, 71, 120, 121, 124, 125,
	126, 127, 130, 128, 129, 131, 41, 42, 146, 35,
	36, 240, 93, 94, 96, 97, 95, 37, 26, 62,
	104, 39, 105, 103, 56, 38, 63, 62, 48, 57,
	47, 46, 43, 44, 45, 196, 195, 193, 95, 65,
	53, 118, 64, 66, 107, 62, 104, 60, 105, 103,
	298, 229, 228, 55, 54, 100, 67, 68, 260, 50,
	93, 94, 96, 97, 109, 217, 194, 247, 69, 9,
	59, 376, 2, 10, 155, 62, 48, 57, 47, 46,
	43, 44, 45, 70, 157, 18, 95, 65, 0, 0,
	0, 66, 0, 22, 23, 60, 62, 104, 58, 105,
	103, 0, 72, 0, 67, 68, 100, 50, 71, 0,
	93, 94, 96, 97, 0, 0, 69, 0, 59, 41,
	0, 0, 35, 36, 61, 0, 93, 94, 96, 97,
	37, 70, 0, 0, 39, 75, 95, 0, 38, 63,
	0, 79, 76, 0, 25, 0, 58, 80, 0, 0,
	72, 0, 95, 0, 0, 0, 71, 121, 124, 125,
	126, 127, 130, 128, 129, 131, 0, 41, 0, 0,
	35, 36, 61, 302, 299, 305, 306, 303, 37, 304,
	301, 0, 39, 0, 353, 0, 38, 63, 62, 48,
	57, 47, 46, 43, 44, 45, 93, 94, 96, 97,
	65, 0, 0, 0, 66, 0, 0, 0, 60, 93,
	94, 96, 97, 352, 0, 0, 0, 67, 68, 0,
	50, 0, 95, 0, 8, 93, 94, 96, 97, 69,
	0, 59, 0, 0, 0, 95, 62, 48, 57, 47,
	46, 43, 44, 45, 70, 0, 0, 0, 65, 0,
	0, 95, 66, 0, 0, 0, 60, 25, 0, 58,
	0, 0, 0, 72, 0, 67, 68, 0, 50, 71,
	93, 94, 96, 97, 0, 0, 0, 69, 327, 59,
	41, 0, 0, 35, 36, 61, 0, 0, 0, 0,
	0, 37, 70, 0, 0, 39, 95, 0, 0, 38,
	63, 0, 0, 0, 0, 25, 0, 58, 0, 0,
	0, 72, 302, 299, 305, 306, 303, 71, 304, 301,
	0, 360, 361, 0, 0, 0, 0, 0, 41, 0,
	0, 35, 36, 61, 0, 0, 0, 0, 0, 37,
	0, 0, 0, 39, 224, 0, 0, 38, 63, 62,
	48, 57, 47, 46, 43, 44, 45, 0, 0, 0,
	0, 65, 0, 0, 0, 66, 0, 0, 0, 60,
	0, 0, 0, 0, 0, 0, 0, 0, 67, 68,
	0, 50, 62, 48, 57, 47, 46, 43, 44, 45,
	69, 0, 59, 0, 65, 0, 0, 0, 66, 0,
	0, 0, 60, 0, 0, 70, 0, 0, 0, 0,
	0, 67, 68, 0, 50, 0, 0, 0, 25, 0,
	58, 0, 0, 69, 72, 59, 0, 0, 0, 0,
	71, 210, 0, 0, 0, 0, 0, 0, 70, 0,
	0, 41, 0, 0, 35, 36, 61, 0, 0, 0,
	0, 0, 37, 58, 0, 0, 39, 72, 0, 0,
	38, 63, 0, 71, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 41, 0, 0, 35, 36, 61,
	0, 0, 0, 0, 0, 37, 0, 0, 0, 39,
	0, 0, 0, 38, 63, 62, 48, 57, 47, 46,
	43, 44, 45, 0, 0, 0, 0, 65, 0, 0,
	0, 66, 0, 0, 0, 60, 0, 0, 0, 0,
	0, 0, 0, 0, 67, 68, 0, 50, 0, 0,
	0, 0, 0, 0, 0, 0, 69, 0, 59, 0,
	0, 0, 0, 0, 208, 0, 0, 0, 0, 0,
	0, 70, 62, 48, 57, 47, 46, 43, 44, 45,
	0, 0, 0, 0, 65, 0, 58, 0, 66, 0,
	72, 0, 60, 0, 0, 0, 71, 0, 0, 0,
	0, 67, 68, 0, 50, 0, 0, 41, 0, 0,
	35, 36, 61, 69, 0, 59, 0, 0, 37, 0,
	0, 0, 39, 0, 0, 0, 38, 63, 70, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 58, 0, 0, 0, 72, 0, 0,
	0, 0, 0, 71, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 41, 0, 0, 35, 36, 61,
	0, 0, 0, 0, 0, 37, 0, 0, 0, 39,
	0, 0, 0, 38, 63,
}
var yyPact = [...]int{

	26, -1000, -1000, -1000, -1000, -1000, 365, 26, 567, -1000,
	606, 415, 415, 415, 357, -43, -51, 320, 841, 214,
	-1000, -1000, -1000, -1000, 537, 841, 382, -1000, 29, 115,
	-1000, -1000, -1000, -1000, -1000, 1044, 1044, 1044, 1044, 680,
	-1000, -15, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -78, -1000, -81, 1, -1000, -1000, 841, -1000,
	-1000, -1000, -1000, 23, -1000, 0, -1, -2, -10, -11,
	-1000, -1000, -1000, 292, 324, 269, 94, 285, 316, 841,
	841, 26, -1000, -1000, 26, 26, 841, -1000, -1000, 334,
	451, 258, 841, 841, 841, 841, 841, 841, -1000, -1000,
	511, -1000, -1000, -1000, -1000, -1000, 43, 345, 1044, 519,
	124, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 106, 99,
	1044, 1044, 1044, 1044, 987, 874, 1044, 1044, 1044, 1044,
	1044, 1044, -1000, 384, 1044, 6, 6, 6, 6, -51,
	363, -1000, 121, 358, 728, 394, 105, 636, -1000, 358,
	415, 406, 415, 415, 415, -1000, -1000, -1000, -1000, -1000,
	279, -1000, -1000, 194, 636, -1000, 636, 191, -1000, -1000,
	-1000, 191, 193, -1000, 439, 265, 63, -1000, -1000, -1000,
	-1000, 334, -1000, 101, 101, -1000, 43, 43, -1000, -1000,
	-1000, -1000, -1000, -15, -1000, -1000, -1000, -1000, -23, 1044,
	1044, 1044, 1044, 555, 41, 395, 395, 98, 841, 98,
	841, 12, 12, 12, 12, 12, -1000, -1000, -1000, 6,
	-1000, -1000, -83, -1000, -1000, -4, 636, -1000, -25, -32,
	-1000, -1000, 116, -1000, 841, -1000, 841, 841, 841, 336,
	-52, 841, 841, 841, 310, 841, 841, -1000, -1000, -1000,
	325, 841, 192, 611, -1000, -1000, 588, 588, 40, -1000,
	-84, -1000, -1000, 265, -1000, -1000, 680, 65, 6, 238,
	238, 487, 487, 121, -1000, 841, 245, 841, 841, -86,
	-1000, 841, 697, 239, -6, 190, -1000, -1000, 184, 83,
	-13, -1000, -1000, -1000, 327, 347, 636, 63, 63, 63,
	163, 169, -1000, 186, 179, -1000, -1000, -1000, -1000, 109,
	-7, 611, 451, 325, -8, 1044, -1000, -1000, -1000, -1000,
	-1000, 636, -33, 652, 623, -1000, 636, 841, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, 841, 841, 611, 750, 156,
	178, -1000, 163, -1000, -1000, -1000, 63, -1000, 327, -1000,
	-1000, 841, 240, 229, 636, 636, 188, -1000, 636, -1000,
	841, -36, 841, 63, 176, 611, -1000, 553, 841, 841,
	841, 636, 841, 636, 750, 63, 358, 487, 487, -1000,
	-17, -1000, -1000, -58, -1000, -60, -69, -1000, -1000, -1000,
	-1000, 243, -37, 841, 417, 358, -70, -1000, 241, -39,
	841, 375, 511, -75, -1000,
}
var yyPgo = [...]int{

	0, 585, 229, 42, 17, 13, 179, 584, 574, 573,
	307, 250, 427, 572, 3, 290, 569, 23, 22, 1,
	37, 567, 2, 565, 564, 558, 554, 553, 552, 551,
	10, 21, 550, 544, 20, 542, 24, 29, 541, 15,
	540, 537, 16, 12, 128, 524, 91, 518, 508, 507,
	488, 487, 0, 486, 6, 8, 484, 482, 481, 210,
	14, 476, 27, 30, 25, 5, 18, 163, 471, 469,
	28, 7, 461, 456, 232, 454, 453, 445, 440, 384,
	429, 408, 405, 402, 397, 395,
}
var yyR1 = [...]int{

	0, 81, 13, 13, 14, 14, 12, 12, 12, 12,
	10, 10, 15, 15, 15, 15, 15, 52, 52, 52,
	52, 52, 52, 52, 52, 47, 47, 47, 47, 47,
	60, 60, 60, 60, 60, 60, 60, 46, 46, 46,
	46, 46, 46, 46, 46, 46, 46, 46, 46, 46,
	46, 46, 63, 63, 63, 63, 63, 63, 63, 63,
	63, 63, 63, 63, 63, 63, 63, 57, 57, 57,
	57, 57, 57, 57, 57, 30, 31, 31, 69, 69,
	69, 50, 50, 50, 53, 53, 82, 83, 84, 85,
	53, 45, 45, 45, 45, 45, 45, 45, 45, 68,
	68, 71, 71, 58, 49, 48, 48, 80, 80, 70,
	51, 51, 75, 75, 79, 79, 67, 40, 26, 27,
	27, 36, 36, 18, 18, 19, 19, 19, 19, 44,
	44, 41, 41, 41, 24, 24, 24, 24, 24, 24,
	24, 33, 33, 38, 38, 43, 23, 35, 35, 35,
	22, 28, 29, 1, 20, 20, 20, 37, 37, 6,
	6, 42, 42, 42, 3, 3, 72, 72, 66, 66,
	65, 65, 65, 61, 61, 25, 25, 78, 78, 56,
	56, 56, 56, 56, 56, 32, 32, 32, 34, 34,
	39, 39, 55, 55, 17, 17, 4, 4, 73, 73,
	54, 5, 5, 9, 9, 16, 77, 77, 64, 21,
	21, 21, 7, 7, 2, 2, 8, 8, 11, 11,
	11, 11, 74, 74, 76, 62, 59,
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
	3, 1, 3, 5, 3, 4, 0, 0, 0, 0,
	24, 5, 5, 5, 4, 4, 5, 5, 5, 8,
	8, 1, 3, 1, 5, 1, 0, 1, 2, 4,
	2, 0, 1, 0, 1, 3, 2, 2, 3, 1,
	1, 1, 0, 2, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 2, 1, 0, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 0, 1, 1, 2,
	0, 1, 1, 1, 2, 0, 1, 3, 1, 1,
	2, 2, 3, 1, 3, 1, 1, 1, 3, 3,
	4, 3, 5, 6, 6, 1, 2, 2, 1, 1,
	1, 0, 2, 4, 2, 0, 3, 0, 1, 3,
	1, 2, 0, 1, 0, 3, 1, 3, 2, 1,
	1, 0, 1, 0, 2, 4, 1, 0, 2, 2,
	1, 1, 2, 4, 1, 1, 2,
}
var yyChk = [...]int{

	-1000, -81, -13, -12, -14, -15, -10, 115, 44, -16,
	-9, 58, 60, 59, 42, -12, -14, -75, -1, -79,
	-37, -67, 26, 27, -52, 77, -47, -60, -46, -63,
	-57, -50, -69, -53, -58, 103, 104, 111, 119, 115,
	-14, 100, -49, 13, 14, 15, 12, 11, 9, -44,
	40, -31, -30, -40, -26, -27, -45, 10, 79, 51,
	28, 105, 8, 120, -35, 20, 24, 37, 38, 49,
	64, 89, 83, -2, -11, 29, 36, -74, -59, 35,
	41, -20, 16, -37, -20, -20, 21, 116, 116, -6,
	34, -79, 61, 73, 74, 99, 75, 76, -36, -18,
	18, -19, -30, 12, 9, 11, -52, -33, 98, -24,
	83, 84, 90, 86, 91, 85, 93, 95, -38, 48,
	101, 102, 96, 97, 103, 104, 105, 106, 108, 109,
	107, 110, 77, 113, 74, -63, -63, -63, -63, -14,
	-52, -14, 117, 117, 115, -82, -48, -52, -30, 120,
	115, 115, 115, 115, 115, -8, -11, -7, -2, 52,
	89, -59, -74, -76, -52, -62, -52, -10, -15, -14,
	-10, -10, -77, -64, -52, -3, 30, -42, -30, 12,
	9, -6, -67, -52, -52, -52, -52, -52, -19, -44,
	40, -60, -60, -41, 47, 17, 16, 77, 89, 78,
	87, 88, 87, -46, -46, -46, -46, -46, 57, -46,
	57, -46, -46, -46, -46, -46, -46, -23, 4, -63,
	116, -31, -30, -30, 116, -71, -52, -68, -28, -29,
	5, 6, -80, -70, 80, -30, -20, -37, 16, -52,
	105, -20, -20, -20, 46, 61, 61, -21, 19, 25,
	-17, 53, -72, -66, -65, -56, -61, -14, 115, -42,
	-25, -30, 12, -3, -14, -14, 115, -46, -63, -46,
	-46, -52, -52, 117, 116, 61, -83, 115, 115, -51,
	-70, 82, -52, -52, -71, -52, 116, 116, -52, -52,
	-52, 39, -62, -64, -4, 31, -52, 61, -32, 63,
	-34, 69, 62, 66, 68, 64, 65, -36, -18, -14,
	-78, -66, 117, -17, -71, 75, -43, 7, -43, -31,
	-30, -52, 54, -52, -52, 118, -52, 81, 116, 116,
	116, 116, 116, 116, -5, 32, 21, -66, -66, -65,
	-39, 67, -34, 62, 62, 116, 61, -42, -4, 116,
	-60, 115, 61, 61, -52, -52, -73, -54, -52, -55,
	71, 72, 71, 62, -39, -66, -5, -52, 57, 57,
	61, -52, 115, -52, -66, 62, 18, -52, -52, -54,
	-71, -55, -65, -22, -30, -43, -43, 116, 116, 116,
	116, -84, 55, 115, -52, 61, -22, 116, -85, 55,
	115, -52, 72, -19, 116,
}
var yyDef = [...]int{

	0, -2, 1, 2, -2, -2, 204, 0, 113, -2,
	0, 156, 156, 156, 0, 0, 11, 160, 0, 112,
	153, 114, 157, 158, 122, 0, 24, 29, -2, 51,
	52, 53, 54, 55, 57, 0, 0, 0, 0, 0,
	64, 0, 66, 67, 68, 69, 70, 71, 72, 73,
	74, 81, -2, 78, 79, 0, 86, 103, 106, 129,
	130, 77, 75, 0, 120, 0, 0, 0, 0, 0,
	147, 148, 149, 217, 213, 0, 0, 220, 221, 0,
	0, 0, 154, 155, 0, 0, 0, 4, 5, 165,
	0, 160, 0, 0, 0, 0, 0, 0, 116, 121,
	0, 124, 125, 126, 127, 128, 22, 0, 0, 0,
	141, 134, 135, 136, 137, 138, 139, 140, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 143, 0, 0, 59, 60, 61, 62, -2,
	0, 65, 0, 0, 0, 0, 0, 105, 117, 0,
	156, 0, 156, 156, 156, 8, 216, 9, 212, 214,
	0, 218, 219, 222, 224, 226, 225, 14, 10, 11,
	15, 16, 205, 206, 211, 195, 0, 159, 161, 162,
	163, 165, 115, 17, 18, 19, 20, 21, 123, 23,
	25, 26, 27, 0, 131, 132, 133, 142, 0, 0,
	0, 0, 0, 37, 38, 39, 40, 41, 0, 42,
	0, 43, 44, 45, 46, 47, 48, 56, 146, 58,
	63, 82, 76, 80, 84, 0, 101, 87, 0, 0,
	151, 152, 111, 107, 0, 118, 0, 0, 0, 0,
	77, 0, 0, 0, 0, 0, 0, 208, 209, 210,
	197, 0, 164, 166, 168, 169, 122, 0, 0, 173,
	0, -2, -2, 195, 28, 30, 0, 0, 34, 35,
	33, 0, 0, 0, 85, 0, 0, 0, 0, 0,
	108, 0, 0, 0, 0, 0, 94, 95, 0, 0,
	0, 215, 223, 207, 202, 0, 194, 0, 0, 0,
	191, 0, 185, 0, 0, 188, 189, 170, 171, 11,
	0, 177, 0, 197, 0, 0, 49, 145, 50, 83,
	76, 102, 0, 0, 0, 104, 110, 0, 91, 92,
	93, 96, 97, 98, 12, 0, 0, 167, 179, 181,
	0, 190, 191, 186, 187, 172, 0, 174, 202, 31,
	32, 0, 0, 0, 109, 201, 196, 198, 200, 180,
	0, 0, 0, 0, 0, 178, 13, 0, 0, 0,
	0, 192, 0, 182, 0, 0, 0, 0, 0, 199,
	0, 183, 184, 0, 150, 0, 0, 193, 88, 99,
	100, 0, 0, 0, 0, 0, 0, 89, 0, 0,
	0, 0, 0, 0, 90,
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
		//line parser.y:208
		{
			yylex.(*Tokener).ParseStmt = yyDollar[1].node
		}
	case 4:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:227
		{
			yyVAL.node = yyDollar[2].node
		}
	case 5:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:231
		{
			yyVAL.node = yyDollar[2].node
		}
	case 6:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:237
		{
			yyVAL.node = yyDollar[1].node
		}
	case 7:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:241
		{

		}
	case 8:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:245
		{
		}
	case 9:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:248
		{
		}
	case 10:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:252
		{
			yyVAL.node = yyDollar[1].node
		}
	case 11:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:253
		{
			yyVAL.node = yyDollar[1].node
		}
	case 12:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.y:258
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
		//line parser.y:280
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
		//line parser.y:301
		{
			stmt := &ast.SelectStmt{All: yyDollar[3].val, Left: yyDollar[1].node.(*ast.SelectStmt), Right: yyDollar[4].node.(*ast.SelectStmt)}
			stmt.SetTag(ast.AST_SELECT_UNION)
			yyVAL.node = stmt
		}
	case 15:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:307
		{
			stmt := &ast.SelectStmt{All: yyDollar[3].val, Left: yyDollar[1].node.(*ast.SelectStmt), Right: yyDollar[4].node.(*ast.SelectStmt)}
			stmt.SetTag(ast.AST_SELECT_INTERSECT)
			yyVAL.node = stmt
		}
	case 16:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:313
		{
			stmt := &ast.SelectStmt{All: yyDollar[3].val, Left: yyDollar[1].node.(*ast.SelectStmt), Right: yyDollar[4].node.(*ast.SelectStmt)}
			stmt.SetTag(ast.AST_SELECT_EXCEPT)
			yyVAL.node = stmt
		}
	case 17:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:329
		{
			expr := &ast.Expr{Operator: "OR", Left: yyDollar[1].expr, Right: yyDollar[3].expr}
			expr.SetTag(ast.AST_EXPR_OR)
			yyVAL.expr = expr
		}
	case 18:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:335
		{
			expr := &ast.Expr{Operator: "||", Left: yyDollar[1].expr, Right: yyDollar[3].expr}
			expr.SetTag(ast.AST_EXPR_DOUBLE_OR)
			yyVAL.expr = expr
		}
	case 19:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:341
		{
			expr := &ast.Expr{Operator: "XOR", Left: yyDollar[1].expr, Right: yyDollar[3].expr}
			expr.SetTag(ast.AST_EXPR_XOR)
			yyVAL.expr = expr
		}
	case 20:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:347
		{
			expr := &ast.Expr{Operator: "AND", Left: yyDollar[1].expr, Right: yyDollar[3].expr}
			expr.SetTag(ast.AST_EXPR_AND)
			yyVAL.expr = expr
		}
	case 21:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:353
		{
			expr := &ast.Expr{Operator: "&&", Left: yyDollar[1].expr, Right: yyDollar[3].expr}
			expr.SetTag(ast.AST_EXPR_DOUBLE_AND)
			yyVAL.expr = expr
		}
	case 22:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:359
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
		//line parser.y:372
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
		//line parser.y:391
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 25:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:397
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
		//line parser.y:409
		{
			expr := &ast.Expr{Operator: "<=>", Left: yyDollar[1].expr, Right: yyDollar[3].expr}
			expr.SetTag(ast.AST_EXPR_LEG)
			yyVAL.expr = expr
		}
	case 27:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:415
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
		//line parser.y:436
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
		//line parser.y:449
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 30:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:455
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
		//line parser.y:466
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
		//line parser.y:477
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
		//line parser.y:489
		{
			expr := &ast.Expr{Operator: "LIKE", Left: yyDollar[1].expr, Right: yyDollar[4].expr}
			expr.SetTag(ast.AST_EXPR_SOUND_LIKE)
			yyVAL.expr = expr
		}
	case 34:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:495
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
		//line parser.y:507
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
		//line parser.y:519
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 37:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:526
		{
			expr := &ast.Expr{Operator: "|", Left: yyDollar[1].expr, Right: yyDollar[3].expr}
			expr.SetTag(ast.AST_EXPR_VETICEL)
			yyVAL.expr = expr
		}
	case 38:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:532
		{
			expr := &ast.Expr{Operator: "&", Left: yyDollar[1].expr, Right: yyDollar[3].expr}
			expr.SetTag(ast.AST_EXPR_AMPERSAND)
			yyVAL.expr = expr
		}
	case 39:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:538
		{
			expr := &ast.Expr{Operator: "<<", Left: yyDollar[1].expr, Right: yyDollar[3].expr}
			expr.SetTag(ast.AST_EXPR_SL)
			yyVAL.expr = expr
		}
	case 40:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:544
		{
			expr := &ast.Expr{Operator: ">>", Left: yyDollar[1].expr, Right: yyDollar[3].expr}
			expr.SetTag(ast.AST_EXPR_SR)
			yyVAL.expr = expr
		}
	case 41:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:550
		{
			expr := &ast.Expr{Operator: "+", Left: yyDollar[1].expr, Right: yyDollar[3].expr}
			expr.SetTag(ast.AST_EXPR_PLUS)
			yyVAL.expr = expr
		}
	case 42:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:556
		{
			expr := &ast.Expr{Operator: "-", Left: yyDollar[1].expr, Right: yyDollar[3].expr}
			expr.SetTag(ast.AST_EXPR_MINUS)
			yyVAL.expr = expr
		}
	case 43:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:562
		{
			expr := &ast.Expr{Operator: "*", Left: yyDollar[1].expr, Right: yyDollar[3].expr}
			expr.SetTag(ast.AST_EXPR_MULT)
			yyVAL.expr = expr
		}
	case 44:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:568
		{
			expr := &ast.Expr{Operator: "/", Left: yyDollar[1].expr, Right: yyDollar[3].expr}
			expr.SetTag(ast.AST_EXPR_DIV)
			yyVAL.expr = expr
		}
	case 45:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:574
		{
			expr := &ast.Expr{Operator: "DIV", Left: yyDollar[1].expr, Right: yyDollar[3].expr}
			expr.SetTag(ast.AST_EXPR_DIV)
			yyVAL.expr = expr
		}
	case 46:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:580
		{
			expr := &ast.Expr{Operator: "MOD", Left: yyDollar[1].expr, Right: yyDollar[3].expr}
			expr.SetTag(ast.AST_EXPR_MOD)
			yyVAL.expr = expr
		}
	case 47:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:586
		{
			expr := &ast.Expr{Operator: "%", Left: yyDollar[1].expr, Right: yyDollar[3].expr}
			expr.SetTag(ast.AST_EXPR_MOD)
			yyVAL.expr = expr
		}
	case 48:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:592
		{
			expr := &ast.Expr{Operator: "^", Left: yyDollar[1].expr, Right: yyDollar[3].expr}
			expr.SetTag(ast.AST_EXPR_XOR)
			yyVAL.expr = expr
		}
	case 49:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:598
		{
			expr := &ast.Expr{Operator: "+", Left: yyDollar[1].expr, Right: yyDollar[4].expr}
			expr.SetTag(ast.AST_TIME_INTERVAL)
			expr.SetAlias(yyDollar[5].val)
			yyVAL.expr = expr
		}
	case 50:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:605
		{
			expr := &ast.Expr{Operator: "-", Left: yyDollar[1].expr, Right: yyDollar[4].expr}
			expr.SetTag(ast.AST_TIME_INTERVAL)
			expr.SetAlias(yyDollar[5].val)
			yyVAL.expr = expr
		}
	case 51:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:612
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 52:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:620
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 53:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:625
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 54:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:630
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 55:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:634
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 56:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:638
		{
			yyDollar[1].expr.SetCollate(yyDollar[3].val)
			yyVAL.expr = yyDollar[1].expr
		}
	case 57:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:643
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 58:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:648
		{
			expr := &ast.Expr{Operator: "||", Left: yyDollar[1].expr, Right: yyDollar[3].expr}
			expr.SetTag(ast.AST_EXPR_DOUBLE_OR)
			yyVAL.expr = expr
		}
	case 59:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:654
		{
			yyVAL.expr = yyDollar[2].expr.(*ast.ValExpr)
		}
	case 60:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:658
		{
			val := yyDollar[2].expr.(*ast.ValExpr)
			val.Symbol = 0 - val.Symbol
			yyVAL.expr = val
		}
	case 61:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:664
		{
			expr := &ast.Expr{Operator: "~", Left: nil, Right: yyDollar[2].expr}
			expr.SetTag(ast.AST_EXPR_TILDE)
			yyVAL.expr = expr
		}
	case 62:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:670
		{
			expr := &ast.Expr{Operator: "!", Left: nil, Right: yyDollar[2].expr}
			expr.SetTag(ast.AST_EXPR_EXCLAMATION_MARK)
			yyVAL.expr = expr
		}
	case 63:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:676
		{
			yyVAL.expr = yyDollar[2].expr
		}
	case 64:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:681
		{
			expr := &ast.SubqueryExpr{Left: nil, Right: yyDollar[1].node.(*ast.SelectStmt)}
			expr.SetTag(ast.AST_SUBQUERY)
			yyVAL.expr = expr
		}
	case 65:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:687
		{
			expr := &ast.SubqueryExpr{Operator: "EXISTS", Left: nil, Right: yyDollar[2].node.(*ast.SelectStmt)}
			expr.SetTag(ast.AST_SUBQUERY_EXISTS)
			yyVAL.expr = expr
		}
	case 66:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:693
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 67:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:699
		{
			expr := &ast.ValExpr{Symbol: 1, Sval: yyDollar[1].val}
			expr.SetTag(ast.AST_VALUE_BIT_NUMBER)
			yyVAL.expr = expr
		}
	case 68:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:705
		{
			expr := &ast.ValExpr{Symbol: 1, Sval: yyDollar[1].val}
			expr.SetTag(ast.AST_VALUE_HEX_NUMBER)
			yyVAL.expr = expr
		}
	case 69:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:711
		{
			expr := &ast.ValExpr{Symbol: 1, Sval: yyDollar[1].val}
			expr.SetTag(ast.AST_VALUE_NUMBER)
			yyVAL.expr = expr
		}
	case 70:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:717
		{
			expr := &ast.ValExpr{Symbol: 0, Sval: yyDollar[1].str}
			expr.SetTag(ast.AST_VALUE_STRING)
			yyVAL.expr = expr
		}
	case 71:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:723
		{
			expr := &ast.ValExpr{Symbol: 0, Sval: yyDollar[1].str}
			expr.SetTag(ast.AST_VALUE_SINGLE_QUOTA_STRING)
			yyVAL.expr = expr
		}
	case 72:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:729
		{
			expr := &ast.ValExpr{Symbol: 0, Sval: yyDollar[1].str}
			expr.SetTag(ast.AST_VALUE_DOUBLE_QUOTA_STRING)
			yyVAL.expr = expr
		}
	case 73:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:735
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
		//line parser.y:747
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
		//line parser.y:761
		{
			yyVAL.val = yyDollar[1].ident
		}
	case 76:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:767
		{
			yyVAL.val = yyDollar[1].val
		}
	case 77:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:771
		{
			yyVAL.val = "*"
		}
	case 78:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:777
		{
			expr := &ast.ColumnExpr{Column: yyDollar[1].val}
			expr.SetTag(ast.AST_EXPR_SINGLE_AT_COLUMN)
			yyVAL.expr = expr
		}
	case 79:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:783
		{
			expr := &ast.ColumnExpr{Column: yyDollar[1].val}
			expr.SetTag(ast.AST_EXPR_DOUBLE_AT_COLUMN)
			yyVAL.expr = expr
		}
	case 80:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:789
		{
			expr := &ast.ColumnExpr{Column: yyDollar[1].val}
			expr.SetTag(ast.AST_EXPR_DOUBLE_AT_COLUMN)
			yyVAL.expr = expr
		}
	case 81:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:800
		{
			expr := &ast.ColumnExpr{Column: yyDollar[1].val}
			expr.SetTag(ast.AST_EXPR_COLUMN)
			yyVAL.expr = expr
		}
	case 82:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:806
		{
			expr := &ast.ColumnExpr{Relation: &ast.RelationExpr{Table: yyDollar[1].val}, Column: yyDollar[2].val}
			expr.SetTag(ast.AST_EXPR_COLUMN)
			yyVAL.expr = expr
		}
	case 83:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:812
		{
			expr := &ast.ColumnExpr{Relation: &ast.RelationExpr{DB: yyDollar[1].val, Table: yyDollar[2].val}, Column: yyDollar[3].val}
			expr.SetTag(ast.AST_EXPR_COLUMN)
			yyVAL.expr = expr
		}
	case 84:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:821
		{
			expr := &ast.FuncExpr{Name: yyDollar[1].val}
			expr.SetTag(ast.AST_EXPR_FUNC)
			yyVAL.expr = expr
		}
	case 85:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:827
		{
			expr := &ast.FuncExpr{Name: yyDollar[1].val, Arg: yyDollar[3].list}
			expr.SetTag(ast.AST_EXPR_FUNC)
			yyVAL.expr = expr
		}
	case 86:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:833
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 87:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:837
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 88:
		yyDollar = yyS[yypt-10 : yypt+1]
		//line parser.y:841
		{
			expr := &ast.FuncExpr{Name: "CAST", Arg: ast.List{yyDollar[3].expr}}
			expr.SetTag(ast.AST_EXPR_CAST)
			expr.SetCollate(yyDollar[5].ident)
			yyVAL.expr = expr
		}
	case 89:
		yyDollar = yyS[yypt-17 : yypt+1]
		//line parser.y:848
		{
			expr := &ast.FuncExpr{Name: "CONVERT", Arg: ast.List{yyDollar[3].expr}}
			expr.SetTag(ast.AST_EXPR_CONVERT_TYPE)
			expr.SetCollate(yyDollar[5].ident)
			yyVAL.expr = expr
		}
	case 90:
		yyDollar = yyS[yypt-24 : yypt+1]
		//line parser.y:855
		{
			expr := &ast.FuncExpr{Name: "CONVERT", Arg: ast.List{yyDollar[3].expr}}
			expr.SetTag(ast.AST_EXPR_CONVERT_ALIAS)
			expr.SetAlias(yyDollar[5].ident)
			yyVAL.expr = expr
		}
	case 91:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:864
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
		//line parser.y:879
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
		//line parser.y:890
		{
			expr := &ast.FuncExpr{Name: "COUNT", Arg: ast.List{yyDollar[4].expr}}
			expr.SetTag(ast.AST_EXPR_COUNT_ALL)
			yyVAL.expr = expr
		}
	case 94:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:896
		{
			expr := &ast.FuncExpr{Name: "COUNT", Arg: ast.List{yyDollar[3].expr}}
			expr.SetTag(ast.AST_EXPR_COUNT)
			yyVAL.expr = expr
		}
	case 95:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:902
		{
			expr := &ast.FuncExpr{Name: "COUNT"}
			expr.SetTag(ast.AST_EXPR_COUNT_STAR)
			yyVAL.expr = expr
		}
	case 96:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:908
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
		//line parser.y:923
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
		//line parser.y:938
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
		//line parser.y:955
		{

			expr := &ast.Expr{Operator: "+", Left: yyDollar[3].expr, Right: yyDollar[6].expr}
			expr.SetTag(ast.AST_TIME_INTERVAL)
			expr.SetAlias(yyDollar[7].val)
			yyVAL.expr = &ast.FuncExpr{Name: yyDollar[1].val, Arg: ast.List{expr}}
		}
	case 100:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line parser.y:963
		{

			expr := &ast.Expr{Operator: "-", Left: yyDollar[3].expr, Right: yyDollar[6].expr}
			expr.SetTag(ast.AST_TIME_INTERVAL)
			expr.SetAlias(yyDollar[7].val)
			yyVAL.expr = &ast.FuncExpr{Name: yyDollar[1].val, Arg: ast.List{expr}}
		}
	case 101:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:974
		{
			yyVAL.list = ast.List{yyDollar[1].expr}
		}
	case 102:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:978
		{
			yyVAL.list = append(yyDollar[1].list, yyDollar[3].expr)
		}
	case 103:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:986
		{
			expr := &ast.ValExpr{Sval: yyDollar[1].str}
			expr.SetTag(ast.AST_VALUE_MARKER)
			yyVAL.expr = expr
		}
	case 104:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:996
		{
			expr := &ast.CaseExpr{Case: yyDollar[2].expr, When: yyDollar[3].list, Else: yyDollar[4].expr}
			expr.SetTag(ast.AST_EXPR_CASE)
			yyVAL.expr = expr
		}
	case 105:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1004
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 106:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:1008
		{
			yyVAL.expr = nil
		}
	case 107:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1015
		{
			yyVAL.list = ast.List{yyDollar[1].expr}
		}
	case 108:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:1019
		{
			yyVAL.list = append(yyDollar[1].list, yyDollar[2].expr)
		}
	case 109:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:1026
		{
			expr := &ast.Expr{Left: yyDollar[2].expr, Right: yyDollar[4].expr}
			expr.SetTag(ast.AST_EXPR_WHEN)
			yyVAL.expr = expr
		}
	case 110:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:1035
		{
			expr := &ast.Expr{Left: yyDollar[2].expr}
			expr.SetTag(ast.AST_EXPR_ELSE)
			yyVAL.expr = expr
		}
	case 111:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:1041
		{
			yyVAL.expr = nil
		}
	case 112:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1049
		{
			yyVAL.list = yyDollar[1].list
		}
	case 113:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:1053
		{
			yyVAL.list = nil
		}
	case 114:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1059
		{
			yyVAL.list = ast.List{yyDollar[1].expr}

		}
	case 115:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:1064
		{
			yyVAL.list = append(yyDollar[1].list, yyDollar[3].expr)
		}
	case 116:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:1071
		{
			yyDollar[1].expr.SetAlias(yyDollar[2].val)
			yyVAL.expr = yyDollar[1].expr
		}
	case 117:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:1078
		{
			yyVAL.val = yyDollar[2].val
		}
	case 118:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:1083
		{
			yyVAL.val = yyDollar[3].val
		}
	case 119:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1089
		{
			yyVAL.val = yyDollar[1].val
		}
	case 120:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1093
		{
			yyVAL.val = yyDollar[1].val
		}
	case 121:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1098
		{
			yyVAL.val = yyDollar[1].val
		}
	case 122:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:1099
		{
			yyVAL.val = ""
		}
	case 123:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:1102
		{
			yyVAL.val = yyDollar[2].val
		}
	case 124:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1103
		{
			yyVAL.val = yyDollar[1].val
		}
	case 125:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1106
		{
			yyVAL.val = yyDollar[1].val
		}
	case 126:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1107
		{
			yyVAL.val = yyDollar[1].str
		}
	case 127:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1108
		{
			yyVAL.val = yyDollar[1].str
		}
	case 128:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1109
		{
			yyVAL.val = yyDollar[1].str
		}
	case 129:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1113
		{
			yyVAL.val = "TRUE"
		}
	case 130:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1114
		{
			yyVAL.val = "FALSE"
		}
	case 131:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1119
		{
			yyVAL.val = "SOME"
		}
	case 132:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1120
		{
			yyVAL.val = "ANY"
		}
	case 133:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1121
		{
			yyVAL.val = "ALL"
		}
	case 134:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1124
		{
			yyVAL.val = "="
		}
	case 135:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1125
		{
			yyVAL.val = ">="
		}
	case 136:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1126
		{
			yyVAL.val = ">"
		}
	case 137:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1127
		{
			yyVAL.val = "<="
		}
	case 138:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1128
		{
			yyVAL.val = "<"
		}
	case 139:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1129
		{
			yyVAL.val = "<>"
		}
	case 140:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1130
		{
			yyVAL.val = "!="
		}
	case 141:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1133
		{
			yyVAL.val = "IS"
		}
	case 142:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:1134
		{
			yyVAL.val = "IS NOT"
		}
	case 143:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1138
		{
			yyVAL.val = "NOT"
		}
	case 144:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:1139
		{
			yyVAL.val = ""
		}
	case 145:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1142
		{
			yyVAL.val = yyDollar[1].ident
		}
	case 146:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1146
		{
			yyVAL.val = yyDollar[1].ident
		}
	case 147:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1150
		{
			yyVAL.val = "LEFT"
		}
	case 148:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1151
		{
			yyVAL.val = "IN"
		}
	case 149:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1152
		{
			yyVAL.val = "IS"
		}
	case 150:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1155
		{
			yyVAL.val = yyDollar[1].val
		}
	case 151:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1158
		{
			yyVAL.val = yyDollar[1].ident
		}
	case 152:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1161
		{
			yyVAL.val = yyDollar[1].ident
		}
	case 153:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1172
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
		//line parser.y:1188
		{
			yyVAL.val = yyDollar[1].ident
		}
	case 155:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1192
		{
			yyVAL.val = yyDollar[1].val
		}
	case 156:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:1196
		{
			yyVAL.val = ""
		}
	case 157:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1201
		{
			yyVAL.val = "DISTINCT"
		}
	case 158:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1205
		{
			yyVAL.val = "DISTINCTROW"
		}
	case 159:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:1217
		{
			yyVAL.node = &ast.IntoClause{Table: yyDollar[2].val}
		}
	case 160:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:1221
		{
			yyVAL.node = nil
		}
	case 161:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1227
		{
			yyVAL.val = yyDollar[1].val
		}
	case 162:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1231
		{
			yyVAL.val = yyDollar[1].str
		}
	case 163:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1235
		{
			yyVAL.val = yyDollar[1].str
		}
	case 164:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:1250
		{
			yyVAL.node = &ast.FromClause{Table_ref: yyDollar[2].list}
		}
	case 165:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:1254
		{
			yyVAL.node = nil
		}
	case 166:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1260
		{
			yyVAL.list = ast.List{yyDollar[1].expr}
		}
	case 167:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:1264
		{
			yyVAL.list = append(yyDollar[1].list, yyDollar[3].expr)
		}
	case 168:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1270
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 169:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1274
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 170:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:1280
		{
			yyDollar[1].expr.SetAlias(yyDollar[2].val)
			yyVAL.expr = yyDollar[1].expr
		}
	case 171:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:1285
		{
			expr := &ast.SubqueryExpr{Left: nil, Right: yyDollar[1].node.(*ast.SelectStmt)}
			expr.SetTag(ast.AST_SUBQUERY)
			expr.SetAlias(yyDollar[2].val)
			yyVAL.expr = expr
		}
	case 172:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:1292
		{
			yyVAL.expr = &ast.RelationListExpr{Relation: yyDollar[2].list}
		}
	case 173:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1298
		{
			tb := &ast.RelationExpr{Table: yyDollar[1].val}
			tb.SetTag(ast.AST_RELATION)
			yyVAL.expr = tb
		}
	case 174:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:1304
		{
			tb := &ast.RelationExpr{DB: yyDollar[1].val, Table: yyDollar[2].val}
			tb.SetTag(ast.AST_RELATION)
			yyVAL.expr = tb
		}
	case 175:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1312
		{
			yyVAL.val = yyDollar[1].val
		}
	case 176:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1316
		{
			yyVAL.val = yyDollar[1].str
		}
	case 177:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1323
		{
			yyVAL.list = ast.List{yyDollar[1].expr}
		}
	case 178:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:1327
		{
			yyVAL.list = append(yyDollar[1].list, yyDollar[3].expr)
		}
	case 179:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:1336
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
		//line parser.y:1351
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
		//line parser.y:1366
		{

			join := &ast.JoinExpr{Jtype: "STRAIGHT_JOIN", Left: yyDollar[1].expr, Right: yyDollar[3].expr}
			join.SetTag(ast.AST_STRAIGHT_JOIN)
			yyVAL.expr = join
		}
	case 182:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:1373
		{
			join := &ast.JoinExpr{Jtype: "STRAIGHT_JOIN", Left: yyDollar[1].expr, Right: yyDollar[3].expr, Jqual: yyDollar[5].expr}
			join.SetTag(ast.AST_STRAIGHT_JOIN)
			yyVAL.expr = join
		}
	case 183:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.y:1379
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
		//line parser.y:1391
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
		//line parser.y:1404
		{
			yyVAL.val = "JOIN"
		}
	case 186:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:1408
		{
			yyVAL.val = "INNER JOIN"
		}
	case 187:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:1412
		{
			yyVAL.val = "CROSS JOIN"
		}
	case 188:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1419
		{
			yyVAL.val = "LEFT"
		}
	case 189:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1423
		{
			yyVAL.val = "RIGHT"
		}
	case 190:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1430
		{
			yyVAL.val = "OUTER"
		}
	case 191:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:1434
		{
			yyVAL.val = ""
		}
	case 192:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:1440
		{
			yyVAL.expr = yyDollar[2].expr
		}
	case 193:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:1444
		{
			yyVAL.expr = &ast.UsingExpr{Column: yyDollar[3].list}
		}
	case 194:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:1458
		{
			yyVAL.node = &ast.WhereClause{Where: yyDollar[2].expr}
		}
	case 195:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:1462
		{
			yyVAL.node = nil
		}
	case 196:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:1474
		{
			yyVAL.node = &ast.GroupClause{Target_ref: yyDollar[3].list}
		}
	case 197:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:1478
		{
			yyVAL.node = nil
		}
	case 198:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1485
		{
			yyVAL.list = ast.List{yyDollar[1].expr}
		}
	case 199:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:1489
		{
			yyVAL.list = append(yyDollar[1].list, yyDollar[3].expr)
		}
	case 200:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1496
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 201:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:1510
		{
			yyVAL.node = &ast.HavingClause{Having: yyDollar[2].expr}
		}
	case 202:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:1514
		{
			yyVAL.node = nil
		}
	case 203:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1529
		{
			yyVAL.node = yyDollar[1].node
		}
	case 204:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:1533
		{
			yyVAL.node = nil
		}
	case 205:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:1539
		{
			yyVAL.node = &ast.SortClause{Target_ref: yyDollar[3].list}
		}
	case 206:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1546
		{
			yyVAL.list = ast.List{yyDollar[1].expr}
		}
	case 207:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:1550
		{
			yyVAL.list = append(yyDollar[1].list, yyDollar[3].expr)
		}
	case 208:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:1556
		{
			switch node := yyDollar[1].expr.(type) {
			case *ast.Expr:
				node.SetAsc(yyDollar[2].val)
			}

			yyVAL.expr = yyDollar[1].expr
		}
	case 209:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1567
		{
			yyVAL.val = "ASC"
		}
	case 210:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1571
		{
			yyVAL.val = "DESC"
		}
	case 211:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:1575
		{
			yyVAL.val = ""
		}
	case 212:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1586
		{
			yyVAL.node = yyDollar[1].node
		}
	case 213:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:1587
		{
			yyVAL.node = nil
		}
	case 214:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:1591
		{
			yyVAL.node = &ast.LockClause{Lock: "FOR UPDATE"}
		}
	case 215:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:1595
		{
			yyVAL.node = &ast.LockClause{Lock: "LOCK IN SHARE MODE"}
		}
	case 216:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1608
		{
			yyVAL.node = yyDollar[1].node
		}
	case 217:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:1612
		{
			yyVAL.node = nil
		}
	case 218:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:1619
		{
			yyVAL.node = &ast.LimitClause{Limit: yyDollar[1].list, Offset: yyDollar[2].expr}
		}
	case 219:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:1623
		{
			yyVAL.node = &ast.LimitClause{Limit: yyDollar[2].list, Offset: yyDollar[1].expr}
		}
	case 220:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1627
		{
			yyVAL.node = &ast.LimitClause{Limit: yyDollar[1].list}
		}
	case 221:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1631
		{
			yyVAL.node = &ast.LimitClause{Offset: yyDollar[1].expr}
		}
	case 222:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:1638
		{
			yyVAL.list = yyDollar[2].list
		}
	case 223:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:1642
		{
			yyVAL.list = append(yyDollar[2].list, yyDollar[4].expr)
		}
	case 224:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1648
		{
			yyVAL.list = ast.List{yyDollar[1].expr}
		}
	case 225:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1655
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 226:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:1662
		{
			yyVAL.expr = yyDollar[2].expr
		}
	}
	goto yystack /* stack new state and value */
}
