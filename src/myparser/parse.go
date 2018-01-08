//line parse.y:2
package myparser

import __yyfmt__ "fmt"

//line parse.y:4
func SetParseTree(yylex interface{}, stmt string) {
	yylex.(*Tokener).AST = stmt
}

//line parse.y:13
type yySymType struct {
	yys     int
	keyword string
	val     string
	node    string
	list    string
}

const IDENT = 57346
const STRING = 57347
const NUMBER = 57348
const AS = 57349
const ASC = 57350
const ALL = 57351
const ANY = 57352
const AVG = 57353
const BY = 57354
const COMMENT = 57355
const COUNT = 57356
const DISTINCT = 57357
const DESC = 57358
const EXISTS = 57359
const FROM = 57360
const FOR = 57361
const FALSE = 57362
const GROUP = 57363
const HAVING = 57364
const INTO = 57365
const IF = 57366
const LIMIT = 57367
const LOCK = 57368
const NULL = 57369
const ORDER = 57370
const OFFSET = 57371
const SELECT = 57372
const SOME = 57373
const SHARE = 57374
const SUM = 57375
const TO = 57376
const TRUE = 57377
const UPDATE = 57378
const USING = 57379
const WHERE = 57380
const UNION = 57381
const EXCEPT = 57382
const INTERSECT = 57383
const JOIN = 57384
const STRAIGHT_JOIN = 57385
const LEFT = 57386
const RIGHT = 57387
const INNER = 57388
const OUTER = 57389
const CROSS = 57390
const NATURAL = 57391
const USE = 57392
const ON = 57393
const OR = 57394
const AND = 57395
const NOT = 57396
const BETWEEN = 57397
const CASE = 57398
const WHEN = 57399
const THEN = 57400
const ELSE = 57401
const IS = 57402
const LIKE = 57403
const IN = 57404
const GE = 57405
const LE = 57406
const NE = 57407
const IE = 57408
const LNE = 57409
const OP = 57410
const UMINUS = 57411
const END = 57412

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
	"AVG",
	"BY",
	"COMMENT",
	"COUNT",
	"DISTINCT",
	"DESC",
	"EXISTS",
	"FROM",
	"FOR",
	"FALSE",
	"GROUP",
	"HAVING",
	"INTO",
	"IF",
	"LIMIT",
	"LOCK",
	"NULL",
	"ORDER",
	"OFFSET",
	"SELECT",
	"SOME",
	"SHARE",
	"SUM",
	"TO",
	"TRUE",
	"UPDATE",
	"USING",
	"WHERE",
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
	86, 6,
	-2, 10,
	-1, 9,
	1, 7,
	86, 7,
	-2, 51,
	-1, 34,
	39, 173,
	-2, 31,
}

const yyPrivate = 57344

const yyLast = 732

var yyAct = [...]int{

	22, 33, 291, 278, 31, 274, 168, 255, 215, 154,
	165, 109, 171, 58, 101, 143, 111, 294, 210, 264,
	84, 53, 81, 82, 83, 90, 85, 52, 91, 86,
	181, 182, 183, 184, 185, 186, 186, 251, 59, 60,
	61, 62, 63, 64, 206, 100, 102, 24, 170, 4,
	64, 281, 110, 260, 84, 259, 16, 113, 284, 293,
	116, 117, 118, 119, 120, 121, 122, 123, 124, 125,
	126, 127, 128, 129, 130, 131, 132, 213, 87, 250,
	183, 184, 185, 186, 150, 21, 205, 245, 244, 155,
	134, 133, 135, 157, 12, 105, 54, 161, 105, 105,
	160, 57, 53, 55, 67, 65, 66, 77, 75, 79,
	72, 71, 69, 70, 68, 156, 140, 172, 59, 60,
	61, 62, 63, 64, 57, 42, 141, 39, 144, 158,
	287, 174, 204, 43, 175, 34, 26, 27, 61, 62,
	63, 64, 37, 115, 84, 38, 202, 145, 29, 8,
	249, 178, 263, 114, 218, 219, 220, 246, 7, 139,
	88, 102, 110, 7, 36, 248, 136, 212, 95, 98,
	28, 83, 83, 209, 138, 208, 97, 6, 172, 173,
	225, 227, 166, 221, 144, 226, 224, 247, 25, 280,
	141, 96, 254, 222, 16, 203, 137, 216, 268, 218,
	219, 220, 86, 214, 217, 252, 279, 155, 216, 40,
	218, 219, 220, 267, 214, 217, 253, 152, 172, 45,
	258, 172, 223, 261, 103, 83, 262, 106, 107, 141,
	141, 141, 141, 141, 141, 141, 141, 141, 141, 141,
	141, 141, 141, 141, 141, 265, 222, 14, 84, 93,
	16, 269, 270, 16, 177, 149, 148, 272, 275, 199,
	11, 13, 12, 179, 172, 172, 276, 277, 172, 176,
	282, 55, 41, 283, 8, 44, 44, 147, 44, 45,
	286, 285, 45, 173, 256, 275, 172, 211, 288, 292,
	289, 207, 141, 41, 112, 257, 292, 295, 228, 229,
	230, 231, 232, 233, 234, 235, 236, 237, 238, 239,
	240, 241, 242, 243, 74, 73, 76, 78, 51, 84,
	159, 46, 67, 65, 66, 77, 75, 79, 72, 71,
	69, 70, 68, 19, 49, 50, 59, 60, 61, 62,
	63, 64, 163, 84, 1, 271, 82, 47, 15, 3,
	164, 99, 56, 48, 74, 73, 76, 78, 92, 94,
	162, 266, 67, 65, 66, 77, 75, 79, 72, 71,
	69, 70, 68, 5, 108, 9, 59, 60, 61, 62,
	63, 64, 10, 84, 273, 151, 82, 290, 169, 74,
	73, 76, 78, 167, 18, 146, 80, 67, 65, 66,
	77, 75, 79, 72, 71, 69, 70, 68, 89, 35,
	153, 59, 60, 61, 62, 63, 64, 32, 30, 200,
	104, 17, 2, 104, 104, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 74, 73, 76, 78, 0, 0,
	0, 0, 67, 65, 66, 77, 75, 79, 72, 71,
	69, 70, 68, 0, 0, 0, 59, 60, 61, 62,
	63, 64, 74, 73, 76, 78, 0, 0, 0, 0,
	67, 65, 66, 77, 75, 79, 72, 71, 69, 70,
	68, 0, 0, 0, 59, 60, 61, 62, 63, 64,
	73, 76, 78, 0, 0, 0, 0, 67, 65, 66,
	77, 75, 79, 72, 71, 69, 70, 68, 0, 0,
	0, 59, 60, 61, 62, 63, 64, 195, 180, 197,
	0, 0, 0, 0, 0, 189, 187, 188, 198, 196,
	0, 194, 193, 191, 192, 190, 0, 0, 0, 181,
	182, 183, 184, 185, 186, 76, 78, 0, 0, 0,
	0, 67, 65, 66, 77, 75, 79, 72, 71, 69,
	70, 68, 0, 0, 0, 59, 60, 61, 62, 63,
	64, 197, 0, 0, 0, 0, 0, 189, 187, 188,
	198, 196, 0, 194, 193, 191, 192, 190, 0, 0,
	0, 181, 182, 183, 184, 185, 186, 34, 26, 27,
	0, 0, 0, 0, 37, 0, 0, 38, 20, 0,
	29, 34, 26, 27, 0, 0, 0, 0, 37, 0,
	0, 38, 0, 0, 29, 0, 36, 0, 34, 26,
	27, 0, 28, 0, 0, 37, 0, 0, 38, 0,
	36, 29, 34, 26, 27, 0, 28, 0, 0, 37,
	25, 0, 38, 0, 8, 29, 0, 36, 0, 0,
	0, 0, 0, 201, 25, 0, 0, 0, 0, 0,
	0, 36, 23, 0, 0, 0, 0, 28, 34, 26,
	27, 25, 0, 0, 0, 37, 23, 0, 38, 0,
	0, 29, 0, 0, 0, 25, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 36, 0, 0,
	0, 0, 0, 28, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 142,
}
var yyPact = [...]int{

	119, -1000, -1000, -1000, -1000, -1000, 219, 119, 593, -1000,
	253, 338, 338, 338, 306, -59, -65, 248, 607, 57,
	-1000, -1000, 379, -1000, -1000, 638, -1000, -1000, 638, 124,
	-1000, -1000, 121, -56, -1000, -1000, -1000, -1000, -1000, 250,
	274, 155, 190, 251, 638, 638, 119, -1000, -1000, 119,
	119, 638, -1000, -1000, 276, 315, 80, 607, -1000, 638,
	638, 638, 638, 638, 638, 638, 638, 638, 638, 638,
	638, 638, 638, 638, 638, 638, 24, 139, 674, 108,
	246, -1000, 315, -1000, -1000, 488, 299, -1000, 131, -56,
	-1000, 50, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 56,
	407, -1000, 407, 51, -1000, -1000, -1000, 51, 53, -1000,
	334, 144, 140, -1000, 276, -1000, 59, 59, -32, -32,
	-32, -1000, -39, -39, -39, -39, -39, -39, -39, -39,
	488, 434, -39, 638, 234, 108, -1000, 236, -1000, -1000,
	462, -1000, 674, -1000, -1000, 624, 93, -1000, -1000, -1000,
	-1000, -1000, -1000, 0, -1000, 407, -1000, -1000, -1000, 272,
	638, 638, -1000, -1000, -1000, 266, 638, 33, 163, 339,
	339, -1000, -1000, 244, 144, -39, -1000, -1000, -1000, -1000,
	638, 674, 674, 674, 674, 674, 674, 674, 674, 674,
	674, 674, 674, 674, 674, 674, 674, 21, 130, 514,
	-7, 624, 407, -1000, 624, -1000, 638, 160, -1000, -1000,
	262, 283, 407, 140, 10, 8, 140, 107, -1000, -1000,
	-1000, -1000, -1000, 16, -67, 163, 266, 41, 1, 1,
	-46, -46, -46, -1000, -47, -47, -47, -47, -47, -47,
	-47, -47, 514, -47, 674, 178, -1000, 224, -1000, -1000,
	-1000, 638, 259, -1000, -1000, -1000, 638, 638, 163, 140,
	140, 152, 6, 140, -1000, 262, -47, -1000, -1000, -1000,
	407, -1000, 407, 14, -1000, 407, -1000, 152, -1000, 638,
	91, 140, -1000, -1000, 638, -1000, 407, 315, -1000, -1000,
	-27, -1000, -1000, -1000, 315, -1000,
}
var yyPgo = [...]int{

	0, 422, 348, 48, 373, 177, 421, 333, 85, 13,
	22, 1, 0, 116, 47, 15, 419, 418, 417, 410,
	9, 409, 4, 408, 25, 396, 395, 394, 321, 96,
	16, 393, 6, 388, 12, 8, 3, 387, 2, 10,
	18, 384, 5, 7, 382, 375, 374, 11, 360, 359,
	127, 358, 209, 125, 133, 351, 14, 344, 320,
}
var yyR1 = [...]int{

	0, 57, 1, 1, 3, 3, 2, 2, 2, 2,
	5, 5, 4, 4, 4, 4, 4, 27, 28, 28,
	28, 6, 6, 7, 7, 8, 8, 9, 9, 10,
	10, 11, 29, 29, 30, 30, 31, 31, 32, 32,
	32, 33, 39, 39, 40, 40, 41, 41, 42, 43,
	43, 44, 44, 45, 46, 46, 47, 48, 48, 48,
	34, 34, 34, 34, 34, 34, 35, 35, 35, 36,
	36, 37, 37, 38, 49, 49, 58, 50, 51, 51,
	52, 52, 52, 52, 53, 53, 55, 56, 54, 12,
	12, 12, 12, 12, 12, 12, 12, 12, 12, 12,
	12, 12, 12, 12, 12, 12, 12, 12, 12, 12,
	12, 12, 12, 12, 12, 12, 12, 12, 12, 12,
	13, 13, 13, 13, 13, 13, 13, 13, 13, 13,
	13, 13, 13, 13, 13, 13, 13, 13, 13, 13,
	13, 13, 13, 13, 13, 14, 14, 14, 14, 14,
	14, 22, 22, 23, 23, 24, 24, 15, 15, 16,
	16, 25, 25, 25, 25, 25, 25, 25, 26, 26,
	26, 17, 17, 18, 18, 19, 19, 20, 21, 21,
	21,
}
var yyR2 = [...]int{

	0, 1, 1, 1, 3, 3, 1, 2, 4, 4,
	1, 1, 7, 8, 4, 4, 4, 1, 1, 1,
	0, 1, 0, 1, 3, 2, 1, 1, 0, 2,
	1, 1, 2, 0, 2, 0, 1, 3, 2, 2,
	1, 1, 2, 0, 3, 0, 1, 3, 1, 2,
	0, 1, 0, 3, 1, 3, 2, 1, 1, 0,
	3, 4, 5, 4, 5, 4, 1, 1, 1, 2,
	4, 1, 3, 1, 1, 0, 0, 5, 1, 0,
	2, 2, 1, 1, 2, 4, 1, 1, 2, 1,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 2, 3, 4, 3,
	4, 3, 4, 3, 4, 5, 3, 4, 4, 6,
	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 2, 3, 4, 3,
	4, 3, 4, 3, 4, 1, 1, 3, 2, 1,
	1, 1, 2, 1, 2, 2, 2, 1, 3, 1,
	3, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 3, 4, 1, 1, 1, 3, 1, 1, 1,
	1,
}
var yyChk = [...]int{

	-1000, -57, -1, -2, -3, -4, -5, 39, 30, -45,
	-44, 41, 43, 42, 28, -2, -3, -6, -27, -7,
	15, -8, -12, 79, -14, 57, 5, 6, 39, 17,
	-17, -22, -18, -11, 4, -21, 33, 11, 14, -50,
	-52, 19, -53, -54, 25, 29, -28, 9, 15, -28,
	-28, 12, 86, 86, -29, 23, -7, 44, -9, 77,
	78, 79, 80, 81, 82, 64, 65, 63, 73, 71,
	72, 70, 69, 56, 55, 67, 57, 66, 58, 68,
	-25, -10, 7, -11, 4, -12, -12, -3, 39, -23,
	-24, 84, -51, -52, -49, -50, 36, -54, -53, -55,
	-12, -56, -12, -5, -4, -3, -5, -5, -46, -47,
	-12, -30, 18, -11, -29, -8, -12, -12, -12, -12,
	-12, -12, -12, -12, -12, -12, -12, -12, -12, -12,
	-12, -12, -12, 67, 66, 68, 27, 57, 35, 20,
	-13, -14, 57, -15, -3, 39, -26, 31, 10, 9,
	-11, 86, 86, -19, -20, -12, -24, -11, 79, -58,
	44, 44, -48, 8, 16, -39, 38, -31, -32, -33,
	-3, -34, -22, 39, -30, -12, 35, 20, -15, 27,
	56, 77, 78, 79, 80, 81, 82, 64, 65, 63,
	73, 71, 72, 70, 69, 55, 67, 57, 66, -13,
	-16, 39, -12, -3, 39, 86, 44, 19, -56, -47,
	-40, 21, -12, 44, 51, -35, 45, 52, 47, 48,
	49, -9, -10, -3, -34, -32, -39, -12, -13, -13,
	-13, -13, -13, -13, -13, -13, -13, -13, -13, -13,
	-13, -13, -13, -13, 67, 66, 27, 57, 35, 20,
	86, 44, -12, -20, 32, -43, 22, 12, -32, 45,
	45, -32, -35, 45, 86, -40, -13, 35, 20, 27,
	-12, 86, -12, -41, -42, -12, -32, -32, -36, 54,
	37, 45, -32, -43, 44, -36, -12, 39, -32, -42,
	-37, -38, -11, 86, 44, -38,
}
var yyDef = [...]int{

	0, -2, 1, 2, -2, -2, 52, 0, 22, -2,
	0, 20, 20, 20, 0, 0, 11, 33, 0, 21,
	17, 23, 28, 26, 89, 0, 145, 146, 0, 0,
	149, 150, 0, 151, -2, 174, 178, 179, 180, 79,
	75, 0, 82, 83, 0, 0, 0, 18, 19, 0,
	0, 0, 4, 5, 35, 0, 33, 0, 25, 0,
	0, 0, 0, 0, 0, 161, 162, 163, 0, 167,
	164, 166, 165, 0, 0, 0, 0, 0, 0, 0,
	0, 27, 0, 30, 31, 106, 0, 148, 0, 152,
	153, 0, 8, 78, 9, 74, 76, 80, 81, 84,
	86, 88, 87, 14, 10, 11, 15, 16, 53, 54,
	59, 43, 0, 32, 35, 24, 90, 91, 92, 93,
	94, 95, 96, 97, 98, 99, 100, 101, 102, 103,
	104, 105, 107, 0, 0, 0, 109, 0, 111, 113,
	0, 120, 0, 116, 157, 0, 0, 168, 169, 170,
	29, 147, 171, 0, 175, 177, 154, 155, 156, 0,
	0, 0, 56, 57, 58, 45, 0, 34, 36, 28,
	0, 40, 41, 0, 43, 108, 112, 114, 117, 110,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 136,
	0, 0, 159, 118, 0, 172, 0, 0, 85, 55,
	50, 0, 42, 0, 0, 0, 0, 0, 66, 67,
	68, 38, 39, 11, 40, 0, 45, 115, 121, 122,
	123, 124, 125, 126, 127, 128, 129, 130, 131, 132,
	133, 134, 135, 137, 0, 0, 139, 0, 141, 143,
	158, 0, 0, 176, 77, 12, 0, 0, 37, 0,
	0, 0, 0, 0, 60, 50, 138, 142, 144, 140,
	160, 119, 49, 44, 46, 48, 61, 0, 63, 0,
	0, 0, 65, 13, 0, 62, 69, 0, 64, 47,
	0, 71, 73, 70, 0, 72,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 81, 76, 3,
	39, 86, 79, 77, 44, 78, 84, 80, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	64, 63, 65, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 82, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 75, 3, 40,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 41, 42, 43,
	45, 46, 47, 48, 49, 50, 51, 52, 53, 54,
	55, 56, 57, 58, 59, 60, 61, 62, 66, 67,
	68, 69, 70, 71, 72, 73, 74, 83, 85,
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
		//line parse.y:107
		{
			SetParseTree(yylex, yyDollar[1].node)
		}
	case 4:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:119
		{
			yyVAL.node = yyDollar[2].node
		}
	case 5:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:120
		{
			yyVAL.node = yyDollar[2].node
		}
	case 6:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:124
		{
			yyVAL.node = yyDollar[1].node
		}
	case 7:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parse.y:125
		{
			yyVAL.node = yyDollar[1].node + yyDollar[2].node
		}
	case 8:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parse.y:127
		{
			yyVAL.node = yyDollar[1].node + yyDollar[2].node + yyDollar[3].node + yyDollar[4].node
		}
	case 9:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parse.y:131
		{
			yyVAL.node = yyDollar[1].node + yyDollar[2].node + yyDollar[3].node + yyDollar[4].node
		}
	case 10:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:136
		{
			yyVAL.node = yyDollar[1].node
		}
	case 11:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:137
		{
			yyVAL.node = yyDollar[1].node
		}
	case 12:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parse.y:142
		{
			yyVAL.node = "SELECT " + yyDollar[2].node + yyDollar[3].node + yyDollar[4].node + yyDollar[5].node + yyDollar[6].node
		}
	case 13:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line parse.y:146
		{
			yyVAL.node = "SELECT " + yyDollar[2].node + yyDollar[3].node + yyDollar[4].node + yyDollar[5].node + yyDollar[6].node + yyDollar[7].node
		}
	case 14:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parse.y:150
		{
			yyVAL.node = yyDollar[1].node + " UNION " + yyDollar[3].node + yyDollar[4].node
		}
	case 15:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parse.y:154
		{
			yyVAL.node = yyDollar[1].node + " INTERSECT " + yyDollar[3].node + yyDollar[4].node
		}
	case 16:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parse.y:158
		{
			yyVAL.node = yyDollar[1].node + " EXCEPT " + yyDollar[3].node + yyDollar[4].node
		}
	case 17:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:169
		{
			yyVAL.node = " DISTINCT "
		}
	case 18:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:173
		{
			yyVAL.node = " ALL "
		}
	case 19:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:174
		{
			yyVAL.node = " DISTINCT "
		}
	case 20:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parse.y:175
		{
			yyVAL.node = ""
		}
	case 21:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:186
		{
			yyVAL.node = yyDollar[1].node
		}
	case 22:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parse.y:187
		{
			yyVAL.node = ""
		}
	case 23:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:190
		{
			yyVAL.node = yyDollar[1].node
		}
	case 24:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:191
		{
			yyVAL.node = yyVAL.node + "," + yyDollar[3].node
		}
	case 25:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parse.y:195
		{
			yyVAL.node = yyDollar[1].node + " " + yyDollar[2].node
		}
	case 26:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:196
		{
			yyVAL.node = "*"
		}
	case 27:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:207
		{
			yyVAL.node = yyDollar[1].node
		}
	case 28:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parse.y:208
		{
			yyVAL.node = ""
		}
	case 29:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parse.y:211
		{
			yyVAL.node = " AS " + yyDollar[2].node
		}
	case 30:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:212
		{
			yyVAL.node = " " + yyDollar[1].node
		}
	case 31:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:215
		{
			yyVAL.node = yyDollar[1].val
		}
	case 32:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parse.y:225
		{
			yyVAL.node = " INTO " + yyDollar[2].node
		}
	case 33:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parse.y:226
		{
			yyVAL.node = ""
		}
	case 34:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parse.y:235
		{
			yyVAL.node = " FROM " + yyDollar[2].node
		}
	case 35:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parse.y:236
		{
			yyVAL.node = ""
		}
	case 36:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:239
		{
			yyVAL.node = yyDollar[1].node
		}
	case 37:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:240
		{
			yyVAL.node = yyVAL.node + "," + yyDollar[3].node
		}
	case 38:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parse.y:243
		{
			yyVAL.node = yyDollar[1].node + " " + yyDollar[2].node
		}
	case 39:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parse.y:244
		{
			yyVAL.node = "(" + yyDollar[1].node + ") " + yyDollar[2].node
		}
	case 41:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:248
		{
			yyVAL.node = yyDollar[1].node
		}
	case 42:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parse.y:257
		{
			yyVAL.node = " WHERE " + yyDollar[2].node
		}
	case 43:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parse.y:258
		{
			yyVAL.node = ""
		}
	case 44:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:268
		{
			yyVAL.node = " GROUP BY " + yyDollar[3].node
		}
	case 45:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parse.y:269
		{
			yyVAL.node = ""
		}
	case 46:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:273
		{
			yyVAL.node = yyDollar[1].node
		}
	case 47:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:274
		{
			yyVAL.node = yyVAL.node + "," + yyDollar[3].node
		}
	case 48:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:278
		{
			yyVAL.node = yyDollar[1].node
		}
	case 49:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parse.y:288
		{
			yyVAL.node = "  HAVING " + yyDollar[2].node
		}
	case 50:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parse.y:289
		{
			yyVAL.node = ""
		}
	case 51:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:299
		{
			yyVAL.node = yyDollar[1].node
		}
	case 52:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parse.y:300
		{
			yyVAL.node = ""
		}
	case 53:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:303
		{
			yyVAL.node = " ORDER BY " + yyDollar[3].node
		}
	case 54:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:307
		{
			yyVAL.node = yyDollar[1].node
		}
	case 55:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:308
		{
			yyVAL.node = yyVAL.node + "," + yyDollar[3].node
		}
	case 56:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parse.y:311
		{
			yyVAL.node = yyDollar[1].node + yyDollar[2].node
		}
	case 57:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:314
		{
			yyVAL.node = " ASC "
		}
	case 58:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:315
		{
			yyVAL.node = " DESC "
		}
	case 59:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parse.y:316
		{
			yyVAL.node = ""
		}
	case 60:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:326
		{
			yyVAL.node = yyDollar[2].node
		}
	case 61:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parse.y:330
		{
			yyVAL.node = yyDollar[1].node + " CROSS JOIN " + yyDollar[2].keyword
		}
	case 62:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parse.y:334
		{
			yyVAL.node = yyDollar[1].node + " " + yyDollar[2].node + " JOIN " + yyDollar[4].node + " " + yyDollar[5].node
		}
	case 63:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parse.y:338
		{
			yyVAL.node = yyDollar[1].node + " JOIN " + yyDollar[3].node + " " + yyDollar[4].node
		}
	case 64:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parse.y:342
		{
			yyVAL.node = yyDollar[1].node + " NATURAL " + yyDollar[3].node + " JOIN " + yyDollar[4].keyword
		}
	case 65:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parse.y:346
		{
			yyVAL.node = yyDollar[1].node + " NATURAL JOIN " + yyDollar[4].node
		}
	case 66:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:351
		{
			yyVAL.node = " LEFT "
		}
	case 67:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:352
		{
			yyVAL.node = " RIGHT "
		}
	case 68:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:353
		{
			yyVAL.node = " INNER "
		}
	case 69:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parse.y:356
		{
			yyVAL.node = " ON " + yyDollar[2].node
		}
	case 70:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parse.y:357
		{
			yyVAL.node = " USING(" + yyDollar[3].node + ")"
		}
	case 71:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:361
		{
			yyVAL.node = yyDollar[1].node
		}
	case 72:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:362
		{
			yyVAL.node = yyVAL.node + "," + yyDollar[3].node
		}
	case 73:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:365
		{
			yyVAL.node = yyDollar[1].node
		}
	case 74:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:378
		{
			yyVAL.node = yyDollar[1].node
		}
	case 75:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parse.y:379
		{
			yyVAL.node = ""
		}
	case 76:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parse.y:382
		{
			yyVAL.node = " FOR UPDATE "
		}
	case 77:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parse.y:383
		{
			yyVAL.node = " FOR SHARE "
		}
	case 78:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:393
		{
			yyVAL.node = yyDollar[1].node
		}
	case 79:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parse.y:394
		{
			yyVAL.node = ""
		}
	case 80:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parse.y:398
		{
			yyVAL.node = yyDollar[1].node + yyDollar[2].node
		}
	case 81:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parse.y:399
		{
			yyVAL.node = yyDollar[1].node + yyDollar[2].node
		}
	case 82:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:400
		{
			yyVAL.node = yyDollar[1].node
		}
	case 83:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:401
		{
			yyVAL.node = yyDollar[1].node
		}
	case 84:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parse.y:406
		{
			yyVAL.node = " LIMIT " + yyDollar[2].node
		}
	case 85:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parse.y:408
		{
			yyVAL.node = " LIMIT " + yyDollar[2].node + "," + yyDollar[4].node
		}
	case 86:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:413
		{
			yyVAL.node = yyDollar[1].node
		}
	case 87:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:416
		{
			yyVAL.node = yyDollar[1].node
		}
	case 88:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parse.y:421
		{
			yyVAL.node = " OFFSET " + yyDollar[2].node
		}
	case 89:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:430
		{
			yyVAL.node = yyDollar[1].node
		}
	case 90:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:431
		{
			yyVAL.node = yyDollar[1].node + " + " + yyDollar[3].node
		}
	case 91:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:432
		{
			yyVAL.node = yyDollar[1].node + " - " + yyDollar[3].node
		}
	case 92:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:433
		{
			yyVAL.node = yyDollar[1].node + " * " + yyDollar[3].node
		}
	case 93:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:434
		{
			yyVAL.node = yyDollar[1].node + " / " + yyDollar[3].node
		}
	case 94:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:435
		{
			yyVAL.node = yyDollar[1].node + " % " + yyDollar[3].node
		}
	case 95:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:436
		{
			yyVAL.node = yyDollar[1].node + " ^ " + yyDollar[3].node
		}
	case 96:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:437
		{
			yyVAL.node = yyDollar[1].node + " < " + yyDollar[3].node
		}
	case 97:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:438
		{
			yyVAL.node = yyDollar[1].node + " > " + yyDollar[3].node
		}
	case 98:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:439
		{
			yyVAL.node = yyDollar[1].node + " = " + yyDollar[3].node
		}
	case 99:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:440
		{
			yyVAL.node = yyDollar[1].node + " != " + yyDollar[3].node
		}
	case 100:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:441
		{
			yyVAL.node = yyDollar[1].node + " <> " + yyDollar[3].node
		}
	case 101:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:442
		{
			yyVAL.node = yyDollar[1].node + " == " + yyDollar[3].node
		}
	case 102:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:443
		{
			yyVAL.node = yyDollar[1].node + " <= " + yyDollar[3].node
		}
	case 103:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:444
		{
			yyVAL.node = yyDollar[1].node + " >= " + yyDollar[3].node
		}
	case 104:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:445
		{
			yyVAL.node = yyDollar[1].node + " AND " + yyDollar[3].node
		}
	case 105:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:446
		{
			yyVAL.node = yyDollar[1].node + " OR " + yyDollar[3].node
		}
	case 106:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parse.y:447
		{
			yyVAL.node = " NOT " + yyDollar[2].node
		}
	case 107:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:448
		{
			yyVAL.node = yyDollar[1].node + yyDollar[2].keyword + " LIKE " + yyDollar[3].node
		}
	case 108:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parse.y:449
		{
			yyVAL.node = yyDollar[1].node + yyDollar[2].keyword + " NOT LIKE " + yyDollar[3].keyword
		}
	case 109:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:450
		{
			yyVAL.node = yyDollar[1].node + " IS NULL "
		}
	case 110:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parse.y:451
		{
			yyVAL.node = yyDollar[1].node + " IS NOT NULL "
		}
	case 111:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:452
		{
			yyVAL.node = yyDollar[1].node + " IS TRUE "
		}
	case 112:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parse.y:453
		{
			yyVAL.node = yyDollar[1].node + "  NOT IS TRUE "
		}
	case 113:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:454
		{
			yyVAL.node = yyDollar[1].node + " IS FALSE "
		}
	case 114:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parse.y:455
		{
			yyVAL.node = yyDollar[1].node + " NOT IS FALSE "
		}
	case 115:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parse.y:456
		{
			yyVAL.node = yyDollar[1].node + " BETWEEN " + yyDollar[3].node + " AND " + yyDollar[5].node
		}
	case 116:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:457
		{
			yyVAL.node = yyDollar[1].node + " IN " + yyDollar[3].node
		}
	case 117:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parse.y:458
		{
			yyVAL.node = yyDollar[1].node + " NOT IN " + yyDollar[3].keyword
		}
	case 118:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parse.y:459
		{
			yyVAL.node = yyDollar[1].node + yyDollar[2].node + yyDollar[3].node + "(" + yyDollar[4].node + ")"
		}
	case 119:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parse.y:460
		{
			yyVAL.node = yyDollar[1].node + yyDollar[2].node + yyDollar[3].node + "(" + yyDollar[5].node + ")"
		}
	case 121:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:464
		{
			yyVAL.node = yyDollar[1].node + " + " + yyDollar[3].node
		}
	case 122:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:465
		{
			yyVAL.node = yyDollar[1].node + " - " + yyDollar[3].node
		}
	case 123:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:466
		{
			yyVAL.node = yyDollar[1].node + " * " + yyDollar[3].node
		}
	case 124:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:467
		{
			yyVAL.node = yyDollar[1].node + " / " + yyDollar[3].node
		}
	case 125:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:468
		{
			yyVAL.node = yyDollar[1].node + " % " + yyDollar[3].node
		}
	case 126:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:469
		{
			yyVAL.node = yyDollar[1].node + " ^ " + yyDollar[3].node
		}
	case 127:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:470
		{
			yyVAL.node = yyDollar[1].node + " < " + yyDollar[3].node
		}
	case 128:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:471
		{
			yyVAL.node = yyDollar[1].node + " > " + yyDollar[3].node
		}
	case 129:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:472
		{
			yyVAL.node = yyDollar[1].node + " = " + yyDollar[3].node
		}
	case 130:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:473
		{
			yyVAL.node = yyDollar[1].node + " != " + yyDollar[3].node
		}
	case 131:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:474
		{
			yyVAL.node = yyDollar[1].node + " <> " + yyDollar[3].node
		}
	case 132:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:475
		{
			yyVAL.node = yyDollar[1].node + " == " + yyDollar[3].node
		}
	case 133:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:476
		{
			yyVAL.node = yyDollar[1].node + " <= " + yyDollar[3].node
		}
	case 134:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:477
		{
			yyVAL.node = yyDollar[1].node + " >= " + yyDollar[3].node
		}
	case 135:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:478
		{
			yyVAL.node = yyDollar[1].node + " OR " + yyDollar[3].node
		}
	case 136:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parse.y:479
		{
			yyVAL.node = " NOT " + yyDollar[2].node
		}
	case 137:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:480
		{
			yyVAL.node = yyDollar[1].node + yyDollar[2].keyword + " LIKE " + yyDollar[3].node
		}
	case 138:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parse.y:481
		{
			yyVAL.node = yyDollar[1].node + yyDollar[2].keyword + " NOT LIKE " + yyDollar[3].keyword
		}
	case 139:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:482
		{
			yyVAL.node = yyDollar[1].node + " IS NULL "
		}
	case 140:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parse.y:483
		{
			yyVAL.node = yyDollar[1].node + " IS NOT NULL "
		}
	case 141:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:484
		{
			yyVAL.node = yyDollar[1].node + " IS TRUE "
		}
	case 142:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parse.y:485
		{
			yyVAL.node = yyDollar[1].node + "  NOT IS TRUE "
		}
	case 143:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:486
		{
			yyVAL.node = yyDollar[1].node + " IS FALSE "
		}
	case 144:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parse.y:487
		{
			yyVAL.node = yyDollar[1].node + " NOT IS FALSE "
		}
	case 145:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:490
		{
			yyVAL.node = yyDollar[1].val
		}
	case 146:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:491
		{
			yyVAL.node = yyDollar[1].val
		}
	case 147:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:492
		{
			yyVAL.node = "( " + yyDollar[2].node + " )"
		}
	case 148:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parse.y:493
		{
			yyVAL.node = " EXISTS (" + yyDollar[2].node + ")"
		}
	case 149:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:494
		{
			yyVAL.node = yyDollar[1].node
		}
	case 150:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:495
		{
			yyVAL.node = yyDollar[1].node
		}
	case 151:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:506
		{
			yyVAL.node = yyDollar[1].node
		}
	case 152:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parse.y:510
		{
			yyVAL.node = yyDollar[1].node + yyDollar[2].node
		}
	case 153:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:515
		{
			yyVAL.node = yyDollar[1].node
		}
	case 154:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parse.y:516
		{
			yyVAL.node = yyDollar[1].node + yyVAL.node
		}
	case 155:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parse.y:521
		{
			yyVAL.node = "." + yyDollar[2].node
		}
	case 156:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parse.y:525
		{
			yyVAL.node = ".*"
		}
	case 157:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:535
		{
			yyVAL.node = yyDollar[1].node
		}
	case 158:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:536
		{
			yyVAL.node = "(" + yyDollar[2].node + ")"
		}
	case 159:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:540
		{
			yyVAL.node = yyDollar[1].node
		}
	case 160:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:541
		{
			yyVAL.node = yyVAL.node + "," + yyDollar[3].node
		}
	case 161:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:545
		{
			yyVAL.node = "<"
		}
	case 162:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:546
		{
			yyVAL.node = ">"
		}
	case 163:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:547
		{
			yyVAL.node = "="
		}
	case 164:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:548
		{
			yyVAL.node = "=="
		}
	case 165:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:549
		{
			yyVAL.node = ">="
		}
	case 166:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:550
		{
			yyVAL.node = "<="
		}
	case 167:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:551
		{
			yyVAL.node = "<>"
		}
	case 168:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:555
		{
			yyVAL.node = " SOME "
		}
	case 169:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:556
		{
			yyVAL.node = " ANY "
		}
	case 170:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:557
		{
			yyVAL.node = " ALL "
		}
	case 171:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:566
		{
			yyVAL.node = yyDollar[1].node + "()"
		}
	case 172:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parse.y:567
		{
			yyVAL.node = yyDollar[1].node + "(" + yyDollar[3].node + ")"
		}
	case 173:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:570
		{
			yyVAL.node = yyDollar[1].val
		}
	case 174:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:571
		{
			yyVAL.node = yyDollar[1].node
		}
	case 175:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:575
		{
			yyVAL.node = yyDollar[1].node
		}
	case 176:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:577
		{
			yyVAL.node = yyDollar[1].node + "," + yyDollar[3].node
		}
	case 177:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:582
		{
			yyVAL.node = yyDollar[1].node
		}
	case 178:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:587
		{
			yyVAL.node = " SUM "
		}
	case 179:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:588
		{
			yyVAL.node = " AVG "
		}
	case 180:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:589
		{
			yyVAL.node = " COUNT "
		}
	}
	goto yystack /* stack new state and value */
}
