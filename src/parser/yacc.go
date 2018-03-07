//line parser.y:2
package parser

import __yyfmt__ "fmt"

//line parser.y:4
import "ast"

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
const EXISTS = 57399
const OUTER = 57400
const LOW = 57401
const UNION = 57402
const EXCEPT = 57403
const INTERSECT = 57404
const JOIN = 57405
const STRAIGHT_JOIN = 57406
const NATURAL = 57407
const LEFT = 57408
const RIGHT = 57409
const INNER = 57410
const CROSS = 57411
const USE = 57412
const TABLE_REF_PRIORITY = 57413
const ON = 57414
const USING = 57415
const OR = 57416
const OO = 57417
const XOR = 57418
const AND = 57419
const AA = 57420
const BETWEEN = 57421
const CASE = 57422
const WHEN = 57423
const THEN = 57424
const ELSE = 57425
const LIKE = 57426
const REGEXP = 57427
const IN = 57428
const GE = 57429
const LE = 57430
const NE = 57431
const LG = 57432
const IE = 57433
const LNE = 57434
const LEG = 57435
const IS = 57436
const SL = 57437
const SR = 57438
const DIV = 57439
const MOD = 57440
const NOT = 57441
const OP = 57442
const COLLATE = 57443
const INTERVAL = 57444
const UMINUS = 57445
const END = 57446

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
	119, 6,
	-2, 10,
	-1, 9,
	1, 7,
	119, 7,
	-2, 205,
	-1, 28,
	80, 146,
	88, 146,
	89, 146,
	90, 146,
	-2, 36,
	-1, 52,
	121, 121,
	-2, 75,
	-1, 148,
	31, 11,
	37, 11,
	38, 11,
	43, 11,
	44, 11,
	60, 11,
	61, 11,
	62, 11,
	-2, 63,
	-1, 273,
	118, 177,
	-2, 163,
	-1, 274,
	118, 178,
	-2, 164,
}

const yyPrivate = 57344

const yyLast = 1467

var yyAct = [...]int{

	238, 381, 379, 40, 4, 266, 353, 237, 331, 24,
	52, 16, 27, 265, 271, 359, 309, 315, 51, 24,
	108, 262, 240, 188, 109, 348, 116, 178, 186, 29,
	93, 65, 393, 91, 278, 112, 8, 7, 167, 20,
	149, 28, 166, 148, 165, 150, 94, 95, 65, 193,
	164, 163, 274, 162, 161, 160, 65, 193, 159, 158,
	274, 153, 155, 287, 287, 144, 145, 146, 147, 103,
	104, 105, 106, 107, 365, 337, 398, 156, 65, 114,
	397, 115, 113, 375, 374, 129, 373, 8, 110, 296,
	177, 179, 287, 287, 98, 182, 97, 187, 182, 182,
	327, 285, 152, 24, 196, 197, 198, 199, 200, 151,
	7, 191, 347, 376, 346, 190, 210, 143, 141, 403,
	368, 112, 49, 136, 137, 140, 138, 139, 141, 65,
	364, 204, 205, 130, 131, 132, 133, 134, 135, 136,
	137, 140, 138, 139, 141, 157, 99, 142, 341, 286,
	111, 173, 21, 103, 104, 105, 106, 107, 87, 215,
	246, 270, 234, 235, 251, 252, 253, 254, 242, 270,
	233, 241, 216, 217, 218, 219, 220, 222, 224, 225,
	226, 227, 228, 229, 103, 104, 105, 106, 107, 98,
	244, 83, 243, 269, 248, 249, 250, 345, 384, 84,
	273, 103, 104, 105, 106, 107, 5, 88, 396, 255,
	276, 212, 241, 385, 290, 277, 6, 363, 275, 213,
	214, 211, 283, 391, 284, 362, 64, 312, 344, 103,
	104, 105, 106, 107, 103, 104, 105, 106, 107, 258,
	202, 257, 291, 280, 292, 343, 294, 175, 194, 297,
	298, 299, 293, 360, 279, 195, 281, 282, 179, 187,
	102, 201, 289, 100, 311, 103, 104, 105, 106, 107,
	105, 106, 107, 342, 324, 263, 171, 172, 340, 112,
	112, 273, 148, 169, 326, 307, 329, 308, 336, 322,
	102, 338, 256, 333, 323, 174, 335, 328, 181, 106,
	107, 181, 181, 90, 334, 89, 352, 302, 180, 295,
	306, 183, 184, 349, 349, 349, 269, 269, 269, 100,
	358, 382, 383, 273, 273, 273, 356, 357, 354, 350,
	351, 320, 321, 89, 361, 112, 310, 85, 191, 90,
	370, 189, 366, 369, 86, 367, 355, 371, 372, 323,
	103, 104, 105, 106, 107, 377, 380, 96, 339, 103,
	104, 105, 106, 107, 103, 104, 105, 106, 107, 269,
	14, 11, 13, 12, 388, 65, 273, 386, 92, 387,
	389, 390, 231, 392, 1, 394, 11, 13, 12, 269,
	22, 23, 380, 239, 399, 185, 273, 401, 19, 395,
	269, 400, 402, 232, 176, 78, 79, 273, 65, 48,
	60, 47, 46, 43, 44, 45, 245, 101, 15, 3,
	67, 17, 74, 65, 193, 75, 68, 192, 22, 23,
	63, 134, 135, 136, 137, 140, 138, 139, 141, 69,
	70, 63, 50, 378, 264, 76, 332, 330, 32, 121,
	125, 123, 71, 203, 62, 122, 124, 41, 126, 56,
	127, 118, 120, 325, 58, 62, 268, 80, 130, 131,
	132, 133, 134, 135, 136, 137, 140, 138, 139, 141,
	34, 61, 65, 114, 30, 115, 113, 267, 33, 85,
	81, 288, 110, 31, 301, 89, 86, 57, 82, 42,
	154, 90, 26, 35, 36, 247, 103, 104, 105, 106,
	107, 37, 38, 25, 103, 104, 105, 106, 107, 55,
	53, 39, 66, 78, 79, 77, 65, 48, 60, 47,
	46, 43, 44, 45, 209, 208, 59, 54, 67, 272,
	74, 119, 9, 75, 68, 2, 260, 10, 63, 103,
	104, 105, 106, 107, 261, 168, 170, 69, 70, 18,
	50, 65, 114, 76, 115, 113, 128, 303, 207, 117,
	71, 110, 62, 206, 313, 41, 73, 305, 304, 103,
	104, 105, 106, 107, 72, 80, 130, 131, 132, 133,
	134, 135, 136, 137, 140, 138, 139, 141, 300, 61,
	230, 259, 103, 104, 105, 106, 107, 0, 81, 0,
	103, 104, 105, 106, 107, 0, 82, 0, 0, 0,
	0, 35, 36, 64, 103, 104, 105, 106, 107, 37,
	38, 317, 314, 316, 320, 321, 318, 319, 0, 39,
	66, 78, 79, 0, 65, 48, 60, 47, 46, 43,
	44, 45, 0, 0, 0, 0, 67, 0, 74, 0,
	0, 75, 68, 0, 22, 23, 63, 65, 114, 0,
	115, 113, 0, 0, 0, 69, 70, 0, 50, 0,
	0, 76, 0, 0, 0, 0, 0, 0, 71, 0,
	62, 0, 0, 41, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 80, 130, 131, 132, 133, 134, 135,
	136, 137, 140, 138, 139, 141, 0, 61, 317, 314,
	316, 320, 321, 318, 319, 0, 81, 382, 383, 0,
	0, 0, 0, 0, 82, 0, 0, 0, 0, 35,
	36, 64, 0, 0, 0, 0, 0, 37, 38, 25,
	0, 0, 0, 0, 0, 0, 0, 39, 66, 78,
	79, 0, 65, 48, 60, 47, 46, 43, 44, 45,
	0, 0, 0, 0, 67, 0, 74, 0, 0, 75,
	68, 0, 0, 0, 63, 0, 0, 0, 0, 0,
	0, 0, 0, 69, 70, 0, 50, 0, 0, 76,
	0, 8, 0, 0, 0, 0, 71, 0, 62, 0,
	0, 41, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 80, 131, 132, 133, 134, 135, 136, 137, 140,
	138, 139, 141, 0, 0, 61, 0, 0, 0, 0,
	0, 0, 0, 0, 81, 0, 0, 0, 0, 0,
	0, 0, 82, 0, 0, 0, 0, 35, 36, 64,
	0, 0, 0, 0, 0, 37, 38, 25, 0, 0,
	0, 0, 0, 0, 0, 39, 66, 78, 79, 0,
	65, 48, 60, 47, 46, 43, 44, 45, 0, 0,
	0, 0, 67, 0, 74, 0, 0, 75, 68, 0,
	0, 0, 63, 0, 0, 0, 0, 0, 0, 0,
	0, 69, 70, 0, 50, 0, 0, 76, 0, 0,
	0, 0, 0, 0, 71, 0, 62, 0, 0, 41,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 80,
	132, 133, 134, 135, 136, 137, 140, 138, 139, 141,
	0, 0, 0, 61, 0, 0, 0, 0, 0, 0,
	0, 0, 81, 0, 0, 0, 0, 0, 0, 0,
	82, 0, 0, 0, 0, 35, 36, 64, 0, 0,
	0, 0, 0, 37, 38, 25, 0, 0, 0, 0,
	0, 236, 0, 39, 66, 78, 79, 0, 65, 48,
	60, 47, 46, 43, 44, 45, 0, 0, 0, 0,
	67, 0, 74, 0, 0, 75, 68, 0, 0, 0,
	63, 0, 0, 0, 0, 0, 0, 0, 0, 69,
	70, 0, 50, 0, 0, 76, 0, 0, 0, 0,
	0, 0, 71, 0, 62, 0, 0, 41, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 80, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 61, 0, 0, 0, 0, 0, 0, 0, 0,
	81, 0, 0, 0, 0, 0, 0, 0, 82, 0,
	0, 0, 0, 35, 36, 64, 0, 0, 0, 0,
	0, 37, 38, 25, 0, 0, 0, 0, 0, 0,
	0, 39, 66, 78, 79, 0, 65, 48, 60, 47,
	46, 43, 44, 45, 0, 0, 0, 0, 67, 0,
	74, 0, 0, 75, 68, 0, 0, 0, 63, 0,
	0, 0, 0, 0, 0, 0, 0, 69, 70, 0,
	50, 0, 0, 76, 0, 0, 0, 0, 0, 0,
	71, 0, 62, 0, 0, 41, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 80, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 61,
	0, 0, 0, 0, 0, 0, 0, 0, 81, 0,
	0, 0, 0, 0, 0, 0, 82, 0, 0, 0,
	0, 35, 36, 64, 0, 0, 0, 0, 0, 37,
	38, 0, 0, 0, 223, 0, 0, 0, 0, 39,
	66, 78, 79, 0, 65, 48, 60, 47, 46, 43,
	44, 45, 0, 0, 0, 0, 67, 0, 74, 0,
	0, 75, 68, 0, 0, 0, 63, 0, 0, 0,
	0, 0, 0, 0, 0, 69, 70, 0, 50, 0,
	0, 76, 0, 0, 0, 0, 0, 0, 71, 0,
	62, 0, 0, 41, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 80, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 61, 0, 0,
	0, 0, 0, 0, 0, 0, 81, 0, 0, 0,
	0, 0, 0, 0, 82, 0, 0, 0, 0, 35,
	36, 64, 0, 0, 0, 0, 0, 37, 38, 0,
	0, 0, 221, 0, 0, 0, 0, 39, 66, 78,
	79, 0, 65, 48, 60, 47, 46, 43, 44, 45,
	0, 0, 0, 0, 67, 0, 74, 0, 0, 75,
	68, 0, 0, 0, 63, 0, 0, 0, 0, 0,
	0, 0, 0, 69, 70, 0, 50, 0, 0, 76,
	0, 0, 0, 0, 0, 0, 71, 0, 62, 0,
	0, 41, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 80, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 61, 0, 0, 0, 0,
	0, 0, 0, 0, 81, 0, 0, 0, 0, 0,
	0, 0, 82, 0, 0, 0, 0, 35, 36, 64,
	0, 0, 0, 0, 0, 37, 38, 0, 0, 0,
	0, 0, 0, 0, 0, 39, 66,
}
var yyPact = [...]int{

	-11, -1000, -1000, -1000, -1000, -1000, 326, -11, 636, -1000,
	458, 362, 362, 362, 336, -23, -25, 283, 990, 197,
	-1000, -1000, -1000, -1000, 474, 990, 364, -1000, 34, 2,
	-1000, -1000, -1000, -1000, -1000, 1344, 1344, 1344, 1344, 754,
	-1000, -84, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -9, -1000, -16, -1000, -1000, -1000, -1000, -60,
	-1000, 990, -1000, -1000, -1000, -1000, 23, -62, -63, -66,
	-67, -68, -70, -71, -77, -79, -83, -1000, -1000, -1000,
	-1000, -1000, -1000, 296, 306, 222, 61, 260, 268, 990,
	990, -11, -1000, -1000, -11, -11, 990, -1000, -1000, 309,
	415, 227, 990, 990, 990, 990, 990, 990, -1000, -1000,
	659, -1000, -1000, -1000, -1000, -1000, -1000, 411, 1344, 518,
	3, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 131, 71,
	1344, 1344, 1344, 1344, 1226, 1108, 1344, 1344, 1344, 1344,
	1344, 1344, -1000, 378, 2, 2, 2, 2, -25, 284,
	-1000, 121, 367, 872, 89, 549, -1000, 367, 362, 400,
	362, 362, 362, 990, 990, 990, 990, 1344, -1000, -1000,
	-1000, -1000, -1000, 243, -1000, -1000, 178, 549, -1000, 549,
	311, -1000, -1000, -1000, 311, 176, -1000, 527, 219, 48,
	-1000, -1000, -1000, -1000, 309, -1000, 193, 193, 221, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -84, -1000, -1000, -1000,
	-1000, -87, 1344, 1344, 1344, 1344, 722, 839, 328, 328,
	18, 990, 18, 990, 8, 8, 8, 8, 8, -1000,
	-1000, -1000, -1000, -1000, -17, -1000, -1000, 30, 549, 130,
	-1000, 990, -1000, 990, 990, 990, 190, -30, 990, 990,
	990, 535, 431, 289, 504, 487, 269, 990, 990, -1000,
	-1000, -1000, 303, 990, 164, 567, -1000, -1000, 553, 553,
	40, -1000, -18, -1000, -1000, 219, -1000, -1000, 754, 369,
	2, 605, 605, 439, 439, 121, -1000, 990, -45, -1000,
	990, 275, 159, 29, 154, -1000, -1000, 126, 109, 78,
	-2, -4, 367, 367, 367, 990, -1000, -1000, -1000, 294,
	325, 549, 48, 48, 48, 195, 264, -1000, 161, 153,
	-1000, -1000, -1000, -1000, 70, 11, 567, 415, 303, 1,
	1344, -1000, -1000, -1000, -1000, -1000, 549, -1000, 549, 990,
	-1000, -1000, -1000, -1000, -1000, -1000, 990, 990, -33, -1000,
	-35, -36, -6, -1000, 990, 990, 567, 248, 125, 149,
	-1000, 195, -1000, -1000, -1000, 48, -1000, 294, -1000, -1000,
	549, 439, 439, -1000, -1000, -1000, -1000, 549, 160, -1000,
	549, -1000, 990, -89, 990, 48, 144, 567, -1000, -39,
	-43, 990, 549, 990, 549, 654, 48, -1000, -1000, -1000,
	0, -1000, -1000, -1000,
}
var yyPgo = [...]int{

	0, 33, 601, 600, 584, 576, 574, 30, 573, 8,
	569, 566, 122, 17, 15, 559, 191, 23, 16, 6,
	146, 556, 555, 547, 216, 199, 418, 545, 3, 206,
	542, 21, 24, 150, 25, 541, 539, 537, 536, 10,
	18, 525, 20, 520, 14, 519, 41, 502, 500, 499,
	497, 493, 491, 0, 488, 2, 1, 487, 484, 480,
	207, 12, 466, 27, 29, 28, 464, 5, 13, 463,
	152, 459, 448, 22, 7, 444, 443, 158, 421, 404,
	395, 398, 393, 384,
}
var yyR1 = [...]int{

	0, 83, 27, 27, 28, 28, 26, 26, 26, 26,
	24, 24, 29, 29, 29, 29, 29, 53, 53, 53,
	53, 53, 53, 53, 53, 47, 47, 47, 47, 47,
	61, 61, 61, 61, 61, 61, 61, 46, 46, 46,
	46, 46, 46, 46, 46, 46, 46, 46, 46, 46,
	46, 46, 64, 64, 64, 64, 64, 64, 64, 64,
	64, 64, 64, 64, 64, 64, 58, 58, 58, 58,
	58, 58, 58, 58, 39, 40, 40, 72, 72, 72,
	51, 51, 51, 54, 54, 54, 54, 54, 54, 45,
	45, 45, 45, 45, 45, 45, 45, 71, 71, 50,
	50, 50, 66, 74, 74, 59, 49, 48, 48, 82,
	82, 73, 52, 52, 78, 78, 81, 81, 70, 43,
	37, 38, 38, 42, 42, 32, 32, 33, 33, 33,
	33, 12, 12, 8, 8, 8, 35, 35, 35, 35,
	35, 35, 35, 10, 10, 11, 11, 9, 3, 41,
	41, 41, 34, 4, 5, 15, 1, 1, 1, 7,
	7, 20, 20, 44, 44, 44, 17, 17, 75, 75,
	68, 68, 67, 67, 67, 62, 62, 36, 36, 69,
	69, 57, 57, 57, 57, 57, 57, 6, 6, 6,
	13, 13, 14, 14, 56, 56, 31, 31, 18, 18,
	76, 76, 55, 19, 19, 23, 23, 30, 80, 80,
	65, 2, 2, 2, 21, 21, 16, 16, 22, 22,
	25, 25, 25, 25, 77, 77, 79, 60, 63,
}
var yyR2 = [...]int{

	0, 1, 1, 1, 3, 3, 1, 2, 4, 4,
	1, 1, 7, 8, 4, 4, 4, 3, 3, 3,
	3, 3, 2, 1, 3, 3, 3, 3, 4, 1,
	4, 6, 6, 4, 4, 4, 1, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 5,
	5, 1, 1, 1, 1, 1, 1, 3, 2, 2,
	2, 2, 3, 1, 2, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 3,
	1, 3, 5, 1, 1, 1, 1, 3, 4, 5,
	5, 5, 4, 4, 5, 5, 5, 8, 8, 6,
	6, 6, 6, 1, 3, 1, 5, 1, 0, 1,
	2, 4, 2, 0, 1, 0, 1, 3, 2, 2,
	3, 1, 1, 1, 0, 2, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 2, 1, 0, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 0, 1,
	1, 2, 0, 1, 1, 1, 2, 0, 1, 3,
	1, 1, 2, 2, 3, 1, 3, 1, 1, 1,
	3, 3, 4, 3, 5, 6, 6, 1, 2, 2,
	1, 1, 1, 0, 2, 4, 2, 0, 3, 0,
	1, 3, 1, 2, 0, 1, 0, 3, 1, 3,
	2, 1, 1, 0, 1, 0, 2, 4, 1, 0,
	2, 2, 1, 1, 2, 4, 1, 2, 1,
}
var yyChk = [...]int{

	-1000, -83, -27, -26, -28, -29, -24, 121, 47, -30,
	-23, 60, 62, 61, 44, -26, -28, -78, -15, -81,
	-7, -70, 28, 29, -53, 113, -47, -61, -46, -64,
	-58, -51, -72, -54, -59, 103, 104, 111, 112, 121,
	-28, 57, -49, 13, 14, 15, 12, 11, 9, -12,
	42, -40, -39, -43, -37, -45, -71, -50, -66, -38,
	10, 81, 54, 30, 105, 8, 122, 20, 26, 39,
	40, 52, -4, -5, 22, 25, 45, -41, 5, 6,
	67, 90, 98, -16, -25, 31, 38, -77, -60, 37,
	43, -1, 16, -7, -1, -1, 21, 119, 119, -20,
	36, -81, 63, 75, 76, 77, 78, 79, -42, -32,
	18, -33, -39, 12, 9, 11, -53, -10, 97, -35,
	98, 85, 91, 87, 92, 86, 94, 96, -11, 51,
	99, 100, 101, 102, 103, 104, 105, 106, 108, 109,
	107, 110, 113, 115, -64, -64, -64, -64, -28, -53,
	-28, 118, 118, 121, -48, -53, -39, 122, 121, 121,
	121, 121, 121, 121, 121, 121, 121, 121, -22, -25,
	-21, -16, 55, 90, -60, -77, -79, -53, -63, -53,
	-24, -29, -28, -24, -24, -80, -65, -53, -17, 32,
	-44, -39, 12, 9, -20, -70, -53, -53, -53, -53,
	-53, -33, -12, 42, -61, -61, -8, 50, 17, 16,
	113, 90, 80, 88, 89, 88, -46, -46, -46, -46,
	-46, 116, -46, 116, -46, -46, -46, -46, -46, -46,
	-3, 4, 119, -40, -39, -39, 119, -74, -53, -82,
	-73, 82, -39, -1, -7, 16, -53, 105, -1, -1,
	-1, -53, -53, -53, -53, -46, 49, 63, 63, -2,
	19, 27, -31, 56, -75, -68, -67, -57, -62, -28,
	121, -44, -36, -39, 12, -17, -28, -28, 121, -46,
	-64, -46, -46, -53, -53, 118, 119, 63, -52, -73,
	84, -53, -53, -74, -53, 119, 119, -53, -53, -53,
	63, 63, 18, 63, 74, 90, 41, -63, -65, -18,
	33, -53, 63, -6, 65, -13, 66, 64, 69, 70,
	67, 68, -42, -32, -28, -69, -68, 118, -31, -74,
	78, -9, 7, -9, -40, -39, -53, 120, -53, 83,
	119, 119, 119, 119, 119, 119, 116, 116, -34, -39,
	-34, -34, -53, -19, 34, 21, -68, -68, -67, -14,
	58, -13, 64, 64, 119, 63, -44, -18, 119, -61,
	-53, -53, -53, 119, 119, 119, 119, -53, -76, -55,
	-53, -56, 73, 74, 73, 64, -14, -68, -19, -9,
	-9, 63, -53, 121, -53, -68, 64, 119, 119, -55,
	-74, -56, -67, 119,
}
var yyDef = [...]int{

	0, -2, 1, 2, -2, -2, 206, 0, 115, -2,
	0, 158, 158, 158, 0, 0, 11, 162, 0, 114,
	155, 116, 159, 160, 124, 0, 23, 29, -2, 51,
	52, 53, 54, 55, 56, 0, 0, 0, 0, 0,
	63, 0, 65, 66, 67, 68, 69, 70, 71, 72,
	73, 80, -2, 77, 78, 83, 84, 85, 86, 0,
	105, 108, 131, 132, 76, 74, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 122, 153, 154,
	149, 150, 151, 219, 215, 0, 0, 222, 223, 0,
	0, 0, 156, 157, 0, 0, 0, 4, 5, 167,
	0, 162, 0, 0, 0, 0, 0, 0, 118, 123,
	0, 126, 127, 128, 129, 130, 22, 0, 0, 0,
	143, 136, 137, 138, 139, 140, 141, 142, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 145, 0, 58, 59, 60, 61, -2, 0,
	64, 0, 0, 0, 0, 107, 119, 0, 158, 0,
	158, 158, 158, 0, 0, 0, 0, 0, 8, 218,
	9, 214, 216, 0, 220, 221, 224, 226, 227, 228,
	14, 10, 11, 15, 16, 207, 208, 213, 197, 0,
	161, 163, 164, 165, 167, 117, 17, 18, 19, 20,
	21, 125, 24, 25, 26, 27, 0, 133, 134, 135,
	144, 0, 0, 0, 0, 0, 37, 38, 39, 40,
	41, 0, 42, 0, 43, 44, 45, 46, 47, 48,
	57, 148, 62, 81, 75, 79, 87, 0, 103, 113,
	109, 0, 120, 0, 0, 0, 0, 76, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 210,
	211, 212, 199, 0, 166, 168, 170, 171, 124, 0,
	0, 175, 0, -2, -2, 197, 28, 30, 0, 0,
	34, 35, 33, 0, 0, 0, 88, 0, 0, 110,
	0, 0, 0, 0, 0, 92, 93, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 217, 225, 209, 204,
	0, 196, 0, 0, 0, 193, 0, 187, 0, 0,
	190, 191, 172, 173, 11, 0, 179, 0, 199, 0,
	0, 49, 147, 50, 82, 75, 104, 106, 112, 0,
	89, 90, 91, 94, 95, 96, 0, 0, 0, 152,
	0, 0, 0, 12, 0, 0, 169, 181, 183, 0,
	192, 193, 188, 189, 174, 0, 176, 204, 31, 32,
	111, 0, 0, 99, 100, 101, 102, 203, 198, 200,
	202, 182, 0, 0, 0, 0, 0, 180, 13, 0,
	0, 0, 194, 0, 184, 0, 0, 97, 98, 201,
	0, 185, 186, 195,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 112, 3, 3, 3, 107, 100, 3,
	121, 119, 105, 103, 63, 104, 118, 106, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	86, 85, 87, 3, 122, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 110, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 99, 3, 111,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
	52, 53, 54, 55, 56, 57, 58, 59, 60, 61,
	62, 64, 65, 66, 67, 68, 69, 70, 71, 72,
	73, 74, 75, 76, 77, 78, 79, 80, 81, 82,
	83, 84, 88, 89, 90, 91, 92, 93, 94, 95,
	96, 97, 98, 101, 102, 108, 109, 113, 114, 115,
	116, 117, 120,
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
			expr := &ast.Expr{Operator: "OR", Left: yyDollar[1].expr, Right: yyDollar[3].expr}
			expr.SetTag(ast.Expr_Or)
			yyVAL.expr = expr
		}
	case 18:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:366
		{
			expr := &ast.Expr{Operator: "OR", Left: yyDollar[1].expr, Right: yyDollar[3].expr}
			expr.SetTag(ast.Expr_Or)
			yyVAL.expr = expr
		}
	case 19:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:371
		{
			yyVAL.expr = &ast.Expr{Operator: "XOR", Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 20:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:373
		{
			expr := &ast.Expr{Operator: "AND", Left: yyDollar[1].expr, Right: yyDollar[3].expr}
			expr.SetTag(ast.Expr_And)
			yyVAL.expr = expr
		}
	case 21:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:379
		{
			expr := &ast.Expr{Operator: "AND", Left: yyDollar[1].expr, Right: yyDollar[3].expr}
			expr.SetTag(ast.Expr_And)
			yyVAL.expr = expr
		}
	case 22:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:384
		{
			yyVAL.expr = &ast.NotExpr{Expr: yyDollar[2].expr}
		}
	case 23:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:385
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 24:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:387
		{
			yyVAL.expr = &ast.IsTrueExpr{HasNot: yyDollar[2].boolean, Left: yyDollar[1].expr, Right: &ast.True{IsTrue: yyDollar[3].boolean}}
		}
	case 25:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:393
		{
			yyVAL.expr = &ast.IsNullExpr{HasNot: yyDollar[2].boolean, Left: yyDollar[1].expr}
		}
	case 26:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:397
		{
			yyVAL.expr = &ast.Expr{Operator: "<=>", Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 27:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:401
		{
			yyVAL.expr = &ast.Expr{Operator: yyDollar[2].val, Left: yyDollar[1].expr, Right: yyDollar[3].expr}

		}
	case 28:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:406
		{
			expr := &ast.SubqueryExpr{Operator: yyDollar[2].val, Subtype: yyDollar[3].tag, Left: yyDollar[1].expr, Select: yyDollar[4].node}
			expr.SetTag(ast.Subquery_Operator)
			yyVAL.expr = expr
		}
	case 29:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:412
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 30:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:418
		{
			expr := &ast.SubqueryExpr{HasNot: yyDollar[2].boolean, Left: yyDollar[1].expr, Select: yyDollar[4].node}
			expr.SetTag(ast.Subquery_In)
			yyVAL.expr = expr
		}
	case 31:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.y:424
		{
			yyVAL.expr = &ast.InExpr{HasNot: yyDollar[2].boolean, Left: yyDollar[1].expr, Right: yyDollar[5].list}
		}
	case 32:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.y:428
		{
			yyVAL.expr = &ast.BetweenExpr{HasNot: yyDollar[2].boolean, Expr: yyDollar[1].expr, Left: yyDollar[4].expr, Right: yyDollar[6].expr}
		}
	case 33:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:432
		{
			yyVAL.expr = &ast.LikeExpr{HasNot: false, Left: yyDollar[1].expr, Right: yyDollar[4].expr}
		}
	case 34:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:436
		{
			yyVAL.expr = &ast.LikeExpr{HasNot: yyDollar[2].boolean, Left: yyDollar[1].expr, Right: yyDollar[4].expr}
		}
	case 35:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:440
		{
			yyVAL.expr = &ast.RegexpExpr{HasNot: yyDollar[2].boolean, Left: yyDollar[1].expr, Right: yyDollar[4].expr}
		}
	case 36:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:444
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 37:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:450
		{
			yyVAL.expr = &ast.Expr{Operator: "|", Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 38:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:451
		{
			yyVAL.expr = &ast.Expr{Operator: "&", Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 39:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:452
		{
			yyVAL.expr = &ast.Expr{Operator: "<<", Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 40:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:453
		{
			yyVAL.expr = &ast.Expr{Operator: ">>", Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 41:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:454
		{
			yyVAL.expr = &ast.Expr{Operator: "+", Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 42:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:455
		{
			yyVAL.expr = &ast.Expr{Operator: "-", Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 43:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:456
		{
			yyVAL.expr = &ast.Expr{Operator: "*", Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 44:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:457
		{
			yyVAL.expr = &ast.Expr{Operator: "/", Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 45:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:458
		{
			yyVAL.expr = &ast.Expr{Operator: "DIV", Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 46:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:459
		{
			yyVAL.expr = &ast.Expr{Operator: "MOD", Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 47:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:460
		{
			yyVAL.expr = &ast.Expr{Operator: "%", Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 48:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:461
		{
			yyVAL.expr = &ast.Expr{Operator: "^", Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 49:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:463
		{
			yyVAL.expr = &ast.IntervalExpr{Operator: ast.Op_Add, Left: yyDollar[1].expr, Right: yyDollar[4].expr, TimeUnit: yyDollar[5].tag}
		}
	case 50:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:467
		{
			yyVAL.expr = &ast.IntervalExpr{Operator: ast.Op_Minus, Left: yyDollar[1].expr, Right: yyDollar[4].expr, TimeUnit: yyDollar[5].tag}
		}
	case 51:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:471
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 52:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:476
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 53:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:477
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 54:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:478
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 55:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:479
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 56:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:480
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 57:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:482
		{
			yyVAL.expr = &ast.CollateExpr{Expr: yyDollar[1].expr, Collate: yyDollar[3].tag}
		}
	case 58:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:486
		{
			yyVAL.expr = yyDollar[2].expr
		}
	case 59:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:490
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "-", Expr: yyDollar[2].expr}
		}
	case 60:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:494
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "~", Expr: yyDollar[2].expr}
		}
	case 61:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:498
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "!", Expr: yyDollar[2].expr}
		}
	case 62:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:502
		{
			switch node := yyDollar[2].expr.(type) {
			case *ast.Expr:
				node.IsEnclosed = true
			}
			yyVAL.expr = yyDollar[2].expr
		}
	case 63:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:511
		{
			expr := &ast.Tuple{Ref: yyDollar[1].node.(ast.ExprNode)}
			expr.SetTag(ast.Tuple_subquery)
			yyVAL.expr = expr
		}
	case 64:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:517
		{
			expr := &ast.SubqueryExpr{Select: yyDollar[2].node}
			expr.SetTag(ast.Subquery_Exists)
			yyVAL.expr = expr
		}
	case 65:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:523
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 66:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:529
		{
			Expr := &ast.Number{Symbol: true, Val: yyDollar[1].val}
			Expr.SetTag(ast.Bit)
			yyVAL.expr = Expr
		}
	case 67:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:535
		{
			Expr := &ast.Number{Symbol: true, Val: yyDollar[1].val}
			Expr.SetTag(ast.Hex)
			yyVAL.expr = Expr
		}
	case 68:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:541
		{
			Expr := &ast.Number{Symbol: true, Val: yyDollar[1].val}
			Expr.SetTag(ast.Digit)
			yyVAL.expr = Expr
		}
	case 69:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:546
		{
			yyVAL.expr = &ast.String{Str: yyDollar[1].str}
		}
	case 70:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:547
		{
			yyVAL.expr = &ast.String{Str: yyDollar[1].str}
		}
	case 71:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:548
		{
			yyVAL.expr = &ast.String{Str: yyDollar[1].str}
		}
	case 72:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:549
		{
			yyVAL.expr = &ast.True{IsTrue: yyDollar[1].boolean}
		}
	case 73:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:550
		{
			yyVAL.expr = &ast.Null{}
		}
	case 74:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:554
		{
			yyVAL.val = yyDollar[1].ident
		}
	case 75:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:558
		{
			yyVAL.val = yyDollar[1].val
		}
	case 76:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:562
		{
			yyVAL.val = "*"
		}
	case 77:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:567
		{
			yyVAL.expr = &ast.UserVariable{Variable: yyDollar[1].val}
		}
	case 78:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:568
		{
			yyVAL.expr = &ast.SystemVariable{Schema: yyDollar[1].val}
		}
	case 79:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:570
		{
			yyVAL.expr = &ast.SystemVariable{Schema: yyDollar[1].val, Parameter: yyDollar[3].val}
		}
	case 80:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:576
		{
			if yyDollar[1].val == "*" {
				yyVAL.expr = &ast.Column{IsStar: true, Column: "*"}
			} else {
				yyVAL.expr = &ast.Column{Column: yyDollar[1].val}
			}
		}
	case 81:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:584
		{
			if yyDollar[3].val == "*" {
				yyVAL.expr = &ast.Column{IsStar: true, Table: yyDollar[1].val, Column: "*"}
			} else {
				yyVAL.expr = &ast.Column{Table: yyDollar[1].val, Column: yyDollar[3].val}
			}
		}
	case 82:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:592
		{
			if yyDollar[5].val == "*" {
				yyVAL.expr = &ast.Column{IsStar: true, DB: yyDollar[1].val, Table: yyDollar[3].val, Column: "*"}
			} else {
				yyVAL.expr = &ast.Column{DB: yyDollar[1].val, Table: yyDollar[3].val, Column: yyDollar[5].val}
			}
		}
	case 83:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:603
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 84:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:607
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 85:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:611
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 86:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:615
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 87:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:619
		{
			yyVAL.expr = &ast.FuncCall{Name: yyDollar[1].val}
		}
	case 88:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:623
		{
			yyVAL.expr = &ast.FuncCall{Name: yyDollar[1].val, Arg: yyDollar[3].list}
		}
	case 89:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:629
		{
			expr := &ast.AggregationFuncCall{DistinctType: yyDollar[3].tag, Arg: []ast.ExprNode{yyDollar[4].expr}}
			expr.SetTag(ast.Aggregation_Avg)
			yyVAL.expr = expr
		}
	case 90:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:635
		{
			expr := &ast.AggregationFuncCall{DistinctType: yyDollar[3].tag, Arg: yyDollar[4].list}
			expr.SetTag(ast.Aggregation_Count)
			yyVAL.expr = expr
		}
	case 91:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:641
		{
			expr := &ast.AggregationFuncCall{DistinctType: ast.AST_ALL, Arg: []ast.ExprNode{yyDollar[4].expr}}
			expr.SetTag(ast.Aggregation_Count)
			yyVAL.expr = expr
		}
	case 92:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:647
		{
			expr := &ast.AggregationFuncCall{DistinctType: ast.AST_EMPTY, Arg: []ast.ExprNode{yyDollar[3].expr}}
			expr.SetTag(ast.Aggregation_Count)
			yyVAL.expr = expr
		}
	case 93:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:653
		{
			expr := &ast.AggregationFuncCall{DistinctType: ast.AST_EMPTY, Arg: []ast.ExprNode{&ast.Column{IsStar: true, Column: "*"}}}
			expr.SetTag(ast.Aggregation_Count)
			yyVAL.expr = expr
		}
	case 94:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:659
		{
			expr := &ast.AggregationFuncCall{DistinctType: yyDollar[3].tag, Arg: []ast.ExprNode{yyDollar[4].expr}}
			expr.SetTag(ast.Aggregation_Max)
			yyVAL.expr = expr
		}
	case 95:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:665
		{
			expr := &ast.AggregationFuncCall{DistinctType: yyDollar[3].tag, Arg: []ast.ExprNode{yyDollar[4].expr}}
			expr.SetTag(ast.Aggregation_Min)
			yyVAL.expr = expr
		}
	case 96:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:671
		{
			expr := &ast.AggregationFuncCall{DistinctType: yyDollar[3].tag, Arg: []ast.ExprNode{yyDollar[4].expr}}
			expr.SetTag(ast.Aggregation_Sum)
			yyVAL.expr = expr
		}
	case 97:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line parser.y:679
		{
			Expr := &ast.CalTimeFuncCall{Interval: &ast.IntervalExpr{Operator: ast.Op_Add, Left: yyDollar[3].expr, Right: yyDollar[6].expr, TimeUnit: yyDollar[7].tag}}
			Expr.SetTag(yyDollar[1].tag)
			yyVAL.expr = Expr
		}
	case 98:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line parser.y:685
		{
			Expr := &ast.CalTimeFuncCall{Interval: &ast.IntervalExpr{Operator: ast.Op_Minus, Left: yyDollar[3].expr, Right: yyDollar[6].expr, TimeUnit: yyDollar[7].tag}}
			Expr.SetTag(yyDollar[1].tag)
			yyVAL.expr = Expr
		}
	case 99:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.y:693
		{
			yyVAL.expr = &ast.CastFuncCall{Expr: yyDollar[3].expr, CastType: yyDollar[5].val}
		}
	case 100:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.y:697
		{
			yyVAL.expr = &ast.CastFuncCall{Expr: yyDollar[3].expr, CastType: yyDollar[5].val}
		}
	case 101:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.y:701
		{
			yyVAL.expr = &ast.CastFuncCall{Expr: yyDollar[3].expr, CastType: yyDollar[5].val}
		}
	case 102:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.y:707
		{
			yyVAL.expr = &ast.StringFuncCall{Substr: yyDollar[3].expr, Str: yyDollar[5].expr}
		}
	case 103:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:712
		{
			yyVAL.list = []ast.ExprNode{yyDollar[1].expr}
		}
	case 104:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:713
		{
			yyVAL.list = append(yyDollar[1].list, yyDollar[3].expr)
		}
	case 105:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:716
		{
			yyVAL.expr = &ast.Marker{Str: yyDollar[1].str}
		}
	case 106:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:720
		{
			yyVAL.expr = &ast.CaseExpr{Case: yyDollar[2].expr, When: yyDollar[3].list, Else: yyDollar[4].expr}
		}
	case 107:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:725
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 108:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:726
		{
			yyVAL.expr = nil
		}
	case 109:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:729
		{
			yyVAL.list = []ast.ExprNode{yyDollar[1].expr}
		}
	case 110:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:730
		{
			yyVAL.list = append(yyDollar[1].list, yyDollar[2].expr)
		}
	case 111:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:733
		{
			yyVAL.expr = &ast.Expr{Left: yyDollar[2].expr, Right: yyDollar[4].expr}
		}
	case 112:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:736
		{
			yyVAL.expr = &ast.Expr{Left: yyDollar[2].expr}
		}
	case 113:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:737
		{
			yyVAL.expr = nil
		}
	case 114:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:740
		{
			yyVAL.list = yyDollar[1].list
		}
	case 115:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:741
		{
			yyVAL.list = nil
		}
	case 116:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:744
		{
			yyVAL.list = []ast.ExprNode{yyDollar[1].expr}
		}
	case 117:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:745
		{
			yyVAL.list = append(yyDollar[1].list, yyDollar[3].expr)
		}
	case 118:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:749
		{
			tb := &ast.Tuple{Ref: yyDollar[1].expr, Alias: yyDollar[2].val}
			switch yyDollar[1].expr.(type) {
			case *ast.AggregationFuncCall:
				tb.SetTag(ast.Tuple_agg)
			case *ast.FuncCall, *ast.CastFuncCall, *ast.CalTimeFuncCall, *ast.StringFuncCall:
				tb.SetTag(ast.Tuple_funcall)
			case *ast.Column:
				tb.SetTag(ast.Tuple_column)
			default:
				tb.SetTag(ast.Tuple_expr)
			}
			yyVAL.expr = tb
		}
	case 119:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:765
		{
			yyVAL.val = yyDollar[2].val
		}
	case 120:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:768
		{
			yyVAL.val = yyDollar[3].val
		}
	case 121:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:771
		{
			yyVAL.val = yyDollar[1].val
		}
	case 122:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:772
		{
			yyVAL.val = yyDollar[1].val
		}
	case 123:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:775
		{
			yyVAL.val = yyDollar[1].val
		}
	case 124:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:776
		{
			yyVAL.val = ""
		}
	case 125:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:779
		{
			yyVAL.val = yyDollar[2].val
		}
	case 126:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:780
		{
			yyVAL.val = yyDollar[1].val
		}
	case 127:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:783
		{
			yyVAL.val = yyDollar[1].val
		}
	case 128:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:784
		{
			yyVAL.val = yyDollar[1].str
		}
	case 129:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:785
		{
			yyVAL.val = yyDollar[1].str
		}
	case 130:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:786
		{
			yyVAL.val = yyDollar[1].str
		}
	case 131:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:790
		{
			yyVAL.boolean = true
		}
	case 132:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:791
		{
			yyVAL.boolean = false
		}
	case 133:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:796
		{
			yyVAL.tag = ast.Operator_Some
		}
	case 134:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:797
		{
			yyVAL.tag = ast.Operator_All
		}
	case 135:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:798
		{
			yyVAL.tag = ast.Operator_Any
		}
	case 136:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:801
		{
			yyVAL.val = "="
		}
	case 137:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:802
		{
			yyVAL.val = ">="
		}
	case 138:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:803
		{
			yyVAL.val = ">"
		}
	case 139:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:804
		{
			yyVAL.val = "<="
		}
	case 140:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:805
		{
			yyVAL.val = "<"
		}
	case 141:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:806
		{
			yyVAL.val = "<>"
		}
	case 142:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:807
		{
			yyVAL.val = "!="
		}
	case 143:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:810
		{
			yyVAL.boolean = false
		}
	case 144:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:811
		{
			yyVAL.boolean = true
		}
	case 145:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:815
		{
			yyVAL.boolean = true
		}
	case 146:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:816
		{
			yyVAL.boolean = false
		}
	case 147:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:819
		{
			yyVAL.tag = ast.TimeUnit[yyDollar[1].ident]
		}
	case 148:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:823
		{
			yyVAL.tag = ast.Character[yyDollar[1].ident]
		}
	case 149:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:827
		{
			yyVAL.val = "LEFT"
		}
	case 150:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:828
		{
			yyVAL.val = "IN"
		}
	case 151:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:829
		{
			yyVAL.val = "IS"
		}
	case 152:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:832
		{
			yyVAL.val = yyDollar[1].val
		}
	case 153:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:835
		{
			yyVAL.tag = ast.FuncTime[yyDollar[1].ident]
		}
	case 154:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:838
		{
			yyVAL.tag = ast.FuncTime[yyDollar[1].ident]
		}
	case 155:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:849
		{
			clause := &ast.DistinctClause{}
			clause.SetTag(yyDollar[1].tag)
			yyVAL.node = clause
		}
	case 156:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:856
		{
			yyVAL.tag = ast.AST_ALL
		}
	case 157:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:857
		{
			yyVAL.tag = yyDollar[1].tag
		}
	case 158:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:858
		{
			yyVAL.tag = ast.AST_EMPTY
		}
	case 159:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:861
		{
			yyVAL.tag = ast.AST_DISTINCT
		}
	case 160:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:862
		{
			yyVAL.tag = ast.AST_DISTINCTROW
		}
	case 161:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:872
		{
			yyVAL.node = &ast.IntoClause{Relation: &ast.Relation{Table: yyDollar[2].val}}
		}
	case 162:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:874
		{
			clause := &ast.IntoClause{}
			clause.SetTag(ast.AST_EMPTY)
			yyVAL.node = clause
		}
	case 163:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:881
		{
			yyVAL.val = yyDollar[1].val
		}
	case 164:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:882
		{
			yyVAL.val = yyDollar[1].str
		}
	case 165:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:883
		{
			yyVAL.val = yyDollar[1].str
		}
	case 166:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:893
		{
			yyVAL.node = &ast.FromClause{Table_ref: yyDollar[2].list}
		}
	case 167:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:895
		{
			clause := &ast.FromClause{}
			clause.SetTag(ast.AST_EMPTY)
			yyVAL.node = clause
		}
	case 168:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:903
		{
			yyVAL.list = []ast.ExprNode{yyDollar[1].expr}
		}
	case 169:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:904
		{
			yyVAL.list = append(yyDollar[1].list, yyDollar[3].expr)
		}
	case 170:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:907
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 171:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:908
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 172:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:912
		{
			yyDollar[1].expr.(*ast.Relation).Alias = yyDollar[2].val
			yyVAL.expr = yyDollar[1].expr
		}
	case 173:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:917
		{
			tb := &ast.Relation{Select: yyDollar[1].node, Alias: yyDollar[2].val}
			tb.SetTag(ast.Relation_Subquery)
			yyVAL.expr = tb
		}
	case 174:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:923
		{
			switch node := yyDollar[2].expr.(type) {
			case *ast.Join:
				node.IsEnclosed = true
			}
			yyVAL.expr = yyDollar[2].expr
		}
	case 175:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:933
		{
			tb := &ast.Relation{Table: yyDollar[1].val}
			tb.SetTag(ast.Relation_Table)
			yyVAL.expr = tb
		}
	case 176:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:939
		{
			tb := &ast.Relation{DB: yyDollar[1].val, Table: yyDollar[3].val}
			tb.SetTag(ast.Relation_Table)
			yyVAL.expr = tb
		}
	case 177:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:946
		{
			yyVAL.val = yyDollar[1].val
		}
	case 178:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:947
		{
			yyVAL.val = yyDollar[1].str
		}
	case 179:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:950
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 180:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:952
		{
			rel := &ast.Join{Left: yyDollar[1].expr, Right: yyDollar[3].expr}
			rel.SetTag(ast.Cross_join)
			yyVAL.expr = rel
		}
	case 181:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:963
		{
			rel := &ast.Join{Left: yyDollar[1].expr, Right: yyDollar[3].expr}
			rel.SetTag(ast.Cross_join)
			yyVAL.expr = rel
		}
	case 182:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:969
		{
			rel := &ast.Join{Left: yyDollar[1].expr, Right: yyDollar[3].expr, Condition: yyDollar[4].expr}
			rel.SetTag(ast.Inner_join)
			yyVAL.expr = rel
		}
	case 183:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:975
		{
			rel := &ast.Join{Left: yyDollar[1].expr, Right: yyDollar[3].expr}
			rel.SetTag(ast.Straight_join)
			yyVAL.expr = rel
		}
	case 184:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:981
		{
			rel := &ast.Join{Left: yyDollar[1].expr, Right: yyDollar[3].expr, Condition: yyDollar[5].expr}
			rel.SetTag(ast.Straight_join)
			yyVAL.expr = rel
		}
	case 185:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.y:987
		{
			rel := &ast.Join{HasOuter: yyDollar[3].boolean, Left: yyDollar[1].expr, Right: yyDollar[5].expr, Condition: yyDollar[6].expr}
			if yyDollar[2].boolean {
				rel.SetTag(ast.Left_join)
			} else {
				rel.SetTag(ast.Right_join)
			}
			yyVAL.expr = rel
		}
	case 186:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.y:997
		{
			rel := &ast.Join{HasNatural: true, HasOuter: yyDollar[4].boolean, Left: yyDollar[1].expr, Right: yyDollar[6].expr}
			if yyDollar[3].boolean {
				rel.SetTag(ast.Left_join)
			} else {
				rel.SetTag(ast.Right_join)
			}
			yyVAL.expr = rel
		}
	case 187:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1008
		{
			yyVAL.tag = 1
		}
	case 188:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:1009
		{
			yyVAL.tag = 2
		}
	case 189:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:1010
		{
			yyVAL.tag = 3
		}
	case 190:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1014
		{
			yyVAL.boolean = true
		}
	case 191:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1015
		{
			yyVAL.boolean = false
		}
	case 192:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1018
		{
			yyVAL.boolean = true
		}
	case 193:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:1019
		{
			yyVAL.boolean = false
		}
	case 194:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:1022
		{
			yyVAL.expr = yyDollar[2].expr
		}
	case 195:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:1023
		{
			yyVAL.expr = &ast.Using{Column: yyDollar[3].list}
		}
	case 196:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:1033
		{
			yyVAL.node = &ast.WhereClause{Where: yyDollar[2].expr}
		}
	case 197:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:1035
		{
			clause := &ast.WhereClause{}
			clause.SetTag(ast.AST_EMPTY)
			yyVAL.node = clause
		}
	case 198:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:1048
		{
			yyVAL.node = &ast.GroupClause{Target_ref: yyDollar[3].list}
		}
	case 199:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:1050
		{
			clause := &ast.GroupClause{}
			clause.SetTag(ast.AST_EMPTY)
			yyVAL.node = clause
		}
	case 200:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1058
		{
			yyVAL.list = []ast.ExprNode{yyDollar[1].expr}
		}
	case 201:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:1062
		{
			yyVAL.list = append(yyDollar[1].list, yyDollar[3].expr)
		}
	case 202:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1068
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 203:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:1077
		{
			yyVAL.node = &ast.HavingClause{Having: yyDollar[2].expr}
		}
	case 204:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:1079
		{
			clause := &ast.HavingClause{}
			clause.SetTag(ast.AST_EMPTY)
			yyVAL.node = clause
		}
	case 205:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1093
		{
			yyVAL.node = yyDollar[1].node
		}
	case 206:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:1095
		{
			clause := &ast.SortClause{}
			clause.SetTag(ast.AST_EMPTY)
			yyVAL.node = clause
		}
	case 207:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:1102
		{
			yyVAL.node = &ast.SortClause{Target_ref: yyDollar[3].list}
		}
	case 208:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1105
		{
			yyVAL.list = []ast.ExprNode{yyDollar[1].expr}
		}
	case 209:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:1106
		{
			yyVAL.list = append(yyDollar[1].list, yyDollar[3].expr)
		}
	case 210:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:1109
		{
			yyVAL.expr = &ast.Sortby{Expr: yyDollar[1].expr, AscType: yyDollar[2].tag}
		}
	case 211:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1112
		{
			yyVAL.tag = ast.Sort_Asc
		}
	case 212:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1113
		{
			yyVAL.tag = ast.Sort_Desc
		}
	case 213:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:1114
		{
			yyVAL.tag = ast.Sort_Empty
		}
	case 214:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1123
		{
			yyVAL.node = yyDollar[1].node
		}
	case 215:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:1125
		{
			clause := &ast.LockClause{}
			clause.SetTag(ast.AST_EMPTY)
			yyVAL.node = clause
		}
	case 216:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:1133
		{
			lock := &ast.LockClause{}
			lock.SetTag(ast.Lock_update)
			yyVAL.node = lock
		}
	case 217:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:1139
		{
			lock := &ast.LockClause{}
			lock.SetTag(ast.Lock_share)
			yyVAL.node = lock
		}
	case 218:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1153
		{
			yyVAL.node = yyDollar[1].node
		}
	case 219:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:1155
		{
			clause := &ast.LimitClause{}
			clause.SetTag(ast.AST_EMPTY)
			yyVAL.node = clause
		}
	case 220:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:1162
		{
			yyVAL.node = &ast.LimitClause{Limit: yyDollar[1].list, Offset: yyDollar[2].expr}
		}
	case 221:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:1163
		{
			yyVAL.node = &ast.LimitClause{Limit: yyDollar[2].list, Offset: yyDollar[1].expr}
		}
	case 222:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1164
		{
			yyVAL.node = &ast.LimitClause{Limit: yyDollar[1].list}
		}
	case 223:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1165
		{
			yyVAL.node = &ast.LimitClause{Offset: yyDollar[1].expr}
		}
	case 224:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:1171
		{
			yyVAL.list = yyDollar[2].list
		}
	case 225:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:1175
		{
			yyVAL.list = append(yyDollar[2].list, yyDollar[4].expr)
		}
	case 226:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1180
		{
			yyVAL.list = []ast.ExprNode{yyDollar[1].expr}
		}
	case 227:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:1183
		{
			yyVAL.expr = yyDollar[2].expr
		}
	case 228:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:1186
		{
			yyVAL.expr = yyDollar[1].expr
		}
	}
	goto yystack /* stack new state and value */
}
