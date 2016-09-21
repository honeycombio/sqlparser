//line sql.y:6
package sqlparser

import __yyfmt__ "fmt"

//line sql.y:6
func SetParseTree(yylex interface{}, stmt Statement) {
	yylex.(*Tokenizer).ParseTree = stmt
}

func SetAllowComments(yylex interface{}, allow bool) {
	yylex.(*Tokenizer).AllowComments = allow
}

func ForceEOF(yylex interface{}) {
	yylex.(*Tokenizer).ForceEOF = true
}

var (
	SHARE = "share"
	MODE  = "mode"
)

//line sql.y:27
type yySymType struct {
	yys         int
	empty       struct{}
	statement   Statement
	selStmt     SelectStatement
	runes2      [][]rune
	runes       []rune
	run         rune
	str         string
	selectExprs SelectExprs
	selectExpr  SelectExpr
	columns     Columns
	colName     *ColName
	tableExprs  TableExprs
	tableExpr   TableExpr
	smTableExpr SimpleTableExpr
	tableName   *TableName
	indexHints  *IndexHints
	expr        Expr
	boolExpr    BoolExpr
	valExpr     ValExpr
	colTuple    ColTuple
	valExprs    ValExprs
	values      Values
	rowTuple    RowTuple
	subquery    *Subquery
	caseExpr    *CaseExpr
	whens       []*When
	when        *When
	orderBy     OrderBy
	order       *Order
	limit       *Limit
	insRows     InsertRows
	updateExprs UpdateExprs
	updateExpr  *UpdateExpr

	/*
	   for CreateTable
	*/
	createTableStmt   CreateTable
	columnDefinition  *ColumnDefinition
	columnDefinitions ColumnDefinitions
	columnAtts        ColumnAtts
}

const LEX_ERROR = 57346
const SELECT = 57347
const INSERT = 57348
const UPDATE = 57349
const DELETE = 57350
const FROM = 57351
const WHERE = 57352
const GROUP = 57353
const HAVING = 57354
const ORDER = 57355
const BY = 57356
const LIMIT = 57357
const FOR = 57358
const OFFSET = 57359
const ALL = 57360
const DISTINCT = 57361
const AS = 57362
const EXISTS = 57363
const IN = 57364
const IS = 57365
const LIKE = 57366
const BETWEEN = 57367
const NULL = 57368
const ASC = 57369
const DESC = 57370
const VALUES = 57371
const INTO = 57372
const DUPLICATE = 57373
const KEY = 57374
const DEFAULT = 57375
const SET = 57376
const LOCK = 57377
const BINARY = 57378
const ID = 57379
const STRING = 57380
const NUMBER = 57381
const VALUE_ARG = 57382
const LIST_ARG = 57383
const COMMENT = 57384
const LE = 57385
const GE = 57386
const NE = 57387
const NULL_SAFE_EQUAL = 57388
const PRIMARY = 57389
const UNIQUE = 57390
const UNION = 57391
const MINUS = 57392
const EXCEPT = 57393
const INTERSECT = 57394
const JOIN = 57395
const STRAIGHT_JOIN = 57396
const LEFT = 57397
const RIGHT = 57398
const INNER = 57399
const OUTER = 57400
const CROSS = 57401
const NATURAL = 57402
const USE = 57403
const FORCE = 57404
const ON = 57405
const OR = 57406
const AND = 57407
const NOT = 57408
const UNARY = 57409
const CASE = 57410
const WHEN = 57411
const THEN = 57412
const ELSE = 57413
const END = 57414
const CREATE = 57415
const ALTER = 57416
const DROP = 57417
const RENAME = 57418
const ANALYZE = 57419
const ENGINE = 57420
const TABLE = 57421
const INDEX = 57422
const VIEW = 57423
const TO = 57424
const IGNORE = 57425
const IF = 57426
const USING = 57427
const SHOW = 57428
const DESCRIBE = 57429
const EXPLAIN = 57430
const BIT = 57431
const TINYINT = 57432
const SMALLINT = 57433
const MEDIUMINT = 57434
const INT = 57435
const INTEGER = 57436
const BIGINT = 57437
const REAL = 57438
const DOUBLE = 57439
const FLOAT = 57440
const UNSIGNED = 57441
const ZEROFILL = 57442
const DECIMAL = 57443
const NUMERIC = 57444
const DATE = 57445
const TIME = 57446
const TIMESTAMP = 57447
const DATETIME = 57448
const YEAR = 57449
const TEXT = 57450
const CHAR = 57451
const VARCHAR = 57452
const MEDIUMTEXT = 57453
const CHARSET = 57454
const NULLX = 57455
const AUTO_INCREMENT = 57456
const BOOL = 57457
const APPROXNUM = 57458
const INTNUM = 57459

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"LEX_ERROR",
	"SELECT",
	"INSERT",
	"UPDATE",
	"DELETE",
	"FROM",
	"WHERE",
	"GROUP",
	"HAVING",
	"ORDER",
	"BY",
	"LIMIT",
	"FOR",
	"OFFSET",
	"ALL",
	"DISTINCT",
	"AS",
	"EXISTS",
	"IN",
	"IS",
	"LIKE",
	"BETWEEN",
	"NULL",
	"ASC",
	"DESC",
	"VALUES",
	"INTO",
	"DUPLICATE",
	"KEY",
	"DEFAULT",
	"SET",
	"LOCK",
	"BINARY",
	"ID",
	"STRING",
	"NUMBER",
	"VALUE_ARG",
	"LIST_ARG",
	"COMMENT",
	"LE",
	"GE",
	"NE",
	"NULL_SAFE_EQUAL",
	"'('",
	"'='",
	"'<'",
	"'>'",
	"'~'",
	"PRIMARY",
	"UNIQUE",
	"UNION",
	"MINUS",
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
	"FORCE",
	"ON",
	"OR",
	"AND",
	"NOT",
	"'&'",
	"'|'",
	"'^'",
	"'+'",
	"'-'",
	"'*'",
	"'/'",
	"'%'",
	"'.'",
	"UNARY",
	"CASE",
	"WHEN",
	"THEN",
	"ELSE",
	"END",
	"CREATE",
	"ALTER",
	"DROP",
	"RENAME",
	"ANALYZE",
	"ENGINE",
	"TABLE",
	"INDEX",
	"VIEW",
	"TO",
	"IGNORE",
	"IF",
	"USING",
	"SHOW",
	"DESCRIBE",
	"EXPLAIN",
	"BIT",
	"TINYINT",
	"SMALLINT",
	"MEDIUMINT",
	"INT",
	"INTEGER",
	"BIGINT",
	"REAL",
	"DOUBLE",
	"FLOAT",
	"UNSIGNED",
	"ZEROFILL",
	"DECIMAL",
	"NUMERIC",
	"DATE",
	"TIME",
	"TIMESTAMP",
	"DATETIME",
	"YEAR",
	"TEXT",
	"CHAR",
	"VARCHAR",
	"MEDIUMTEXT",
	"CHARSET",
	"NULLX",
	"AUTO_INCREMENT",
	"BOOL",
	"APPROXNUM",
	"INTNUM",
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
}

const yyNprod = 336
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 1022

var yyAct = [...]int{

	159, 63, 90, 437, 372, 516, 163, 323, 91, 51,
	327, 257, 273, 364, 450, 314, 84, 445, 287, 253,
	79, 162, 3, 136, 135, 102, 337, 338, 339, 340,
	341, 526, 342, 343, 529, 526, 52, 53, 513, 80,
	449, 65, 491, 29, 30, 31, 32, 64, 526, 444,
	54, 514, 29, 30, 31, 32, 261, 85, 308, 129,
	75, 375, 68, 333, 370, 129, 418, 420, 44, 481,
	45, 123, 47, 48, 49, 488, 482, 42, 468, 467,
	120, 129, 122, 308, 135, 71, 306, 466, 50, 164,
	69, 46, 315, 165, 362, 487, 489, 419, 441, 428,
	422, 315, 348, 116, 158, 161, 528, 112, 118, 173,
	527, 136, 135, 177, 65, 480, 256, 65, 39, 463,
	41, 168, 359, 525, 178, 365, 513, 148, 149, 150,
	85, 307, 134, 425, 424, 85, 376, 271, 286, 369,
	358, 294, 295, 296, 298, 299, 300, 301, 302, 303,
	304, 305, 284, 285, 277, 282, 349, 329, 309, 136,
	135, 126, 255, 283, 289, 95, 116, 310, 85, 278,
	99, 465, 483, 107, 430, 114, 134, 319, 65, 321,
	100, 83, 96, 97, 98, 267, 365, 479, 311, 312,
	464, 88, 416, 415, 414, 105, 318, 322, 146, 147,
	148, 149, 150, 114, 308, 265, 60, 254, 514, 268,
	133, 337, 338, 339, 340, 341, 87, 342, 343, 412,
	103, 104, 81, 410, 413, 475, 432, 108, 411, 74,
	254, 128, 392, 393, 394, 395, 396, 397, 398, 399,
	400, 401, 115, 106, 402, 403, 387, 388, 389, 390,
	391, 385, 383, 384, 386, 335, 14, 15, 16, 17,
	276, 264, 266, 263, 101, 14, 110, 109, 502, 113,
	275, 325, 331, 29, 30, 31, 32, 281, 114, 77,
	129, 290, 501, 175, 85, 18, 500, 288, 347, 310,
	166, 456, 334, 353, 354, 351, 176, 276, 356, 451,
	446, 277, 258, 169, 350, 167, 521, 275, 495, 357,
	352, 509, 510, 494, 492, 171, 170, 289, 477, 478,
	346, 132, 534, 533, 532, 66, 367, 517, 512, 361,
	371, 507, 363, 368, 493, 328, 421, 345, 131, 20,
	21, 23, 22, 24, 405, 404, 330, 279, 408, 409,
	270, 269, 25, 26, 27, 127, 427, 14, 124, 121,
	423, 277, 277, 429, 143, 144, 145, 146, 147, 148,
	149, 150, 119, 433, 435, 438, 434, 117, 99, 61,
	78, 107, 76, 73, 72, 70, 439, 111, 100, 160,
	96, 97, 98, 511, 471, 431, 474, 59, 14, 166,
	523, 447, 448, 105, 426, 531, 143, 144, 145, 146,
	147, 148, 149, 150, 459, 452, 453, 454, 457, 524,
	455, 458, 317, 291, 508, 292, 293, 469, 103, 104,
	259, 470, 125, 57, 55, 108, 373, 473, 462, 374,
	355, 472, 143, 144, 145, 146, 147, 148, 149, 150,
	324, 106, 143, 144, 145, 146, 147, 148, 149, 150,
	461, 407, 254, 310, 62, 530, 503, 14, 496, 498,
	34, 440, 101, 33, 504, 505, 438, 486, 497, 506,
	499, 143, 144, 145, 146, 147, 148, 149, 150, 35,
	36, 37, 38, 485, 442, 380, 382, 381, 484, 490,
	443, 515, 378, 379, 19, 520, 65, 518, 519, 180,
	181, 182, 183, 184, 185, 186, 187, 188, 189, 190,
	191, 192, 193, 194, 195, 196, 197, 198, 199, 200,
	201, 202, 203, 204, 205, 206, 207, 208, 209, 210,
	211, 179, 326, 377, 260, 40, 332, 262, 43, 67,
	320, 174, 522, 476, 436, 460, 212, 213, 214, 215,
	216, 217, 406, 218, 219, 220, 221, 222, 223, 224,
	225, 226, 227, 228, 229, 230, 231, 360, 172, 313,
	94, 92, 280, 93, 366, 89, 316, 232, 233, 234,
	235, 236, 237, 238, 239, 240, 241, 242, 243, 244,
	245, 246, 247, 248, 249, 250, 251, 252, 180, 181,
	182, 183, 184, 185, 186, 187, 188, 189, 190, 191,
	192, 193, 194, 195, 196, 197, 198, 199, 200, 201,
	202, 203, 204, 205, 206, 207, 208, 209, 210, 211,
	179, 137, 86, 417, 274, 336, 272, 82, 344, 130,
	56, 28, 58, 13, 12, 212, 213, 214, 215, 216,
	217, 11, 218, 219, 220, 221, 222, 223, 224, 225,
	226, 227, 228, 229, 230, 231, 10, 9, 8, 7,
	6, 5, 4, 2, 1, 0, 232, 233, 234, 235,
	236, 237, 238, 239, 240, 241, 242, 243, 244, 245,
	246, 247, 248, 249, 250, 251, 252, 14, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 95, 0, 0, 0, 0, 99, 0,
	0, 107, 0, 0, 0, 0, 0, 0, 100, 83,
	96, 97, 98, 95, 0, 0, 0, 0, 99, 88,
	0, 107, 0, 105, 0, 0, 0, 0, 100, 83,
	96, 97, 98, 0, 0, 0, 0, 0, 0, 88,
	0, 0, 0, 105, 87, 0, 0, 0, 103, 104,
	81, 0, 0, 0, 0, 108, 0, 0, 0, 0,
	0, 0, 14, 0, 87, 0, 0, 0, 103, 104,
	81, 106, 0, 0, 0, 108, 0, 0, 95, 0,
	0, 0, 0, 99, 0, 0, 107, 0, 0, 0,
	0, 106, 101, 100, 160, 96, 97, 98, 95, 0,
	0, 0, 0, 99, 88, 0, 107, 0, 105, 0,
	0, 0, 101, 100, 160, 96, 97, 98, 0, 0,
	0, 0, 0, 0, 88, 0, 0, 0, 105, 87,
	0, 0, 0, 103, 104, 0, 0, 0, 0, 0,
	108, 0, 0, 0, 0, 0, 0, 0, 0, 87,
	0, 0, 0, 103, 104, 0, 106, 0, 99, 0,
	108, 107, 0, 0, 0, 0, 0, 0, 100, 160,
	96, 97, 98, 0, 0, 99, 106, 101, 107, 166,
	0, 0, 0, 105, 0, 100, 160, 96, 97, 98,
	0, 0, 0, 0, 0, 0, 166, 101, 0, 0,
	105, 0, 0, 0, 297, 0, 0, 0, 103, 104,
	0, 0, 0, 0, 0, 108, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 103, 104, 0, 0, 0,
	0, 106, 108, 138, 142, 140, 141, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 106, 0,
	0, 0, 101, 0, 154, 155, 156, 157, 0, 151,
	152, 153, 0, 0, 0, 0, 0, 0, 0, 101,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 139, 143, 144, 145, 146, 147, 148,
	149, 150,
}
var yyPact = [...]int{

	251, -1000, -1000, 219, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	24, -28, -3, -22, -6, -1000, -1000, -1000, 462, 416,
	-1000, -1000, -1000, 414, -1000, 367, 342, 455, 288, -37,
	-5, 348, -1000, -9, 347, -1000, 346, -39, 345, -39,
	343, -1000, -1000, -1000, -1000, -1000, 722, -1000, 225, 342,
	353, 26, 342, 145, -1000, 194, 22, 340, 36, 335,
	-1000, 322, -1000, -26, 321, 411, 92, 318, -1000, 222,
	-1000, -1000, 301, 129, 41, 941, -1000, 807, 787, -1000,
	-1000, -1000, 879, 258, -1000, 256, -1000, -1000, -1000, -1000,
	278, 277, -1000, -1000, -1000, -1000, -1000, -1000, 879, -1000,
	249, 288, 603, 452, 288, 879, 603, 255, 409, -44,
	-1000, 172, -1000, 314, -1000, -1000, 313, -1000, 223, 722,
	-1000, -1000, 310, 504, 144, 807, 807, 879, 240, 401,
	879, 879, 862, 879, 879, 879, 879, 879, 879, 879,
	879, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 941,
	85, -47, -2, 25, 941, -1000, 352, 702, -1000, 462,
	-1000, -1000, 17, 408, 393, 288, 288, 220, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 437, 807, -1000, 408, -1000, 298, -1000,
	88, 309, -1000, -34, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, 197, 152, 300, 260, 21, -1000, -1000, -1000,
	-1000, -1000, 23, 722, -1000, 13, 408, -1000, 352, -1000,
	-1000, 240, 879, 879, 408, 369, 408, 879, 122, 122,
	122, 49, 49, -1000, -1000, -1000, -1000, -1000, 879, -1000,
	408, 7, -11, 8, -1000, 807, 56, 243, 219, 117,
	6, -1000, 437, 421, 425, 41, 3, -1000, 128, 308,
	-1000, -1000, 307, -1000, 450, 223, 223, -1000, -1000, 164,
	160, 135, 134, 133, -1, -1000, 299, -33, 603, -1000,
	1, 0, -1000, 408, 333, 879, 408, 408, -1000, -1000,
	12, -1000, 879, 89, -1000, 364, 168, -1000, -1000, -1000,
	288, 421, -1000, 879, 879, 298, 5, -1000, -65, -1000,
	-1000, 253, -1000, 253, 253, -1000, -87, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 252,
	252, 252, 244, 244, -1000, -1000, 448, 424, 152, 50,
	-1000, 131, -1000, 112, -1000, -1000, -1000, -1000, -8, -16,
	-17, -1000, -1000, -1000, -1000, -1000, 879, 408, -1000, 408,
	879, 362, 243, -1000, -1000, 379, 167, -1000, 291, -1000,
	-1000, 139, 43, -73, -1000, -1000, 275, -1000, -1000, 297,
	-1000, 274, -1000, -1000, -1000, -1000, 269, -1000, -1000, -1000,
	437, 807, 879, 807, -1000, -1000, 239, 235, 221, 408,
	408, 459, -1000, 879, 879, 879, -1000, -1000, -1000, 294,
	398, -1000, 273, -1000, -1000, -1000, -1000, 361, -1000, 296,
	-1000, -1000, -95, -1000, 150, -7, 421, 41, 146, 41,
	290, 290, 290, 288, 408, 408, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 267, 384, -10, -1000, -23, -27,
	145, -99, -1000, 458, 383, -1000, 287, -1000, -1000, -1000,
	-1000, 286, -1000, 285, -1000,
}
var yyPgo = [...]int{

	0, 684, 683, 21, 682, 681, 680, 679, 678, 677,
	676, 661, 654, 653, 473, 652, 651, 650, 20, 39,
	649, 648, 647, 646, 12, 645, 644, 206, 643, 5,
	19, 16, 642, 641, 586, 585, 0, 18, 6, 584,
	8, 583, 25, 581, 2, 580, 579, 15, 578, 577,
	562, 555, 7, 554, 3, 553, 4, 552, 551, 550,
	13, 1, 47, 229, 549, 548, 547, 546, 545, 544,
	11, 9, 543, 10, 542, 504, 17, 503, 502, 500,
	499, 498, 497, 496, 14, 495, 494, 493, 477, 471,
	470,
}
var yyR1 = [...]int{

	0, 1, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 3, 3, 3, 4, 4, 5, 6,
	7, 80, 80, 72, 72, 72, 85, 85, 85, 85,
	85, 77, 77, 77, 77, 77, 78, 78, 82, 82,
	82, 82, 82, 82, 82, 83, 83, 83, 83, 83,
	83, 83, 84, 84, 76, 76, 79, 79, 86, 86,
	86, 86, 86, 86, 86, 81, 81, 87, 87, 88,
	88, 73, 74, 74, 75, 89, 89, 8, 8, 8,
	9, 9, 9, 10, 11, 11, 11, 12, 13, 13,
	13, 90, 14, 15, 15, 16, 16, 16, 16, 16,
	17, 17, 18, 18, 19, 19, 19, 22, 22, 20,
	20, 20, 23, 23, 24, 24, 24, 24, 21, 21,
	21, 25, 25, 25, 25, 25, 25, 25, 25, 25,
	26, 26, 26, 27, 27, 28, 28, 28, 28, 29,
	29, 30, 30, 31, 31, 31, 31, 31, 32, 32,
	32, 32, 32, 32, 32, 32, 32, 32, 33, 33,
	33, 33, 33, 33, 33, 37, 37, 37, 42, 38,
	38, 36, 36, 36, 36, 36, 36, 36, 36, 36,
	36, 36, 36, 36, 36, 36, 36, 36, 36, 41,
	41, 43, 43, 43, 45, 48, 48, 46, 46, 47,
	49, 49, 44, 44, 35, 35, 35, 35, 35, 35,
	50, 50, 51, 51, 52, 52, 53, 53, 54, 55,
	55, 55, 56, 56, 56, 56, 57, 57, 57, 58,
	58, 59, 59, 60, 60, 34, 34, 39, 39, 40,
	40, 61, 61, 62, 63, 63, 64, 64, 65, 65,
	66, 66, 66, 66, 66, 67, 67, 68, 68, 69,
	69, 70, 70, 70, 70, 70, 70, 70, 70, 70,
	70, 70, 70, 70, 70, 70, 70, 70, 70, 70,
	70, 70, 70, 70, 70, 70, 70, 70, 70, 70,
	70, 70, 70, 70, 70, 70, 70, 70, 70, 70,
	70, 70, 70, 70, 70, 70, 70, 70, 70, 70,
	70, 70, 70, 70, 70, 70, 70, 70, 70, 70,
	70, 70, 70, 70, 70, 70, 70, 70, 70, 70,
	70, 70, 70, 70, 70, 71,
}
var yyR2 = [...]int{

	0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 4, 12, 3, 7, 7, 8, 7,
	3, 0, 1, 3, 1, 1, 1, 1, 1, 1,
	1, 2, 2, 1, 1, 3, 2, 1, 1, 1,
	1, 1, 1, 1, 1, 2, 2, 2, 2, 2,
	2, 2, 0, 5, 0, 3, 0, 1, 0, 3,
	2, 3, 3, 2, 2, 1, 1, 2, 1, 1,
	2, 3, 1, 3, 8, 0, 3, 1, 8, 4,
	6, 7, 4, 5, 4, 5, 5, 3, 2, 2,
	2, 0, 2, 0, 2, 1, 2, 1, 1, 1,
	0, 1, 1, 3, 1, 2, 3, 1, 1, 0,
	1, 2, 1, 3, 3, 3, 3, 5, 0, 1,
	2, 1, 1, 2, 3, 2, 3, 2, 2, 2,
	1, 3, 1, 1, 3, 0, 5, 5, 5, 1,
	3, 0, 2, 1, 3, 3, 2, 3, 3, 3,
	4, 3, 4, 5, 6, 3, 4, 2, 1, 1,
	1, 1, 1, 1, 1, 3, 1, 1, 3, 1,
	3, 1, 1, 1, 3, 3, 3, 3, 3, 3,
	3, 3, 2, 3, 4, 5, 4, 4, 1, 1,
	1, 1, 1, 1, 5, 0, 1, 1, 2, 4,
	0, 2, 1, 3, 1, 1, 1, 1, 2, 2,
	0, 3, 0, 2, 0, 3, 1, 3, 2, 0,
	1, 1, 0, 2, 4, 4, 0, 2, 4, 0,
	3, 1, 3, 0, 5, 2, 1, 1, 3, 3,
	1, 1, 3, 3, 0, 2, 0, 3, 0, 1,
	1, 1, 1, 1, 1, 0, 1, 0, 1, 0,
	2, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 0,
}
var yyChk = [...]int{

	-1000, -1, -2, -3, -4, -5, -6, -7, -8, -9,
	-10, -11, -12, -13, 5, 6, 7, 8, 34, -75,
	88, 89, 91, 90, 92, 101, 102, 103, -16, 54,
	55, 56, 57, -14, -90, -14, -14, -14, -14, 94,
	-68, 96, 53, -65, 96, 98, 94, 94, 95, 96,
	94, -71, -71, -71, -3, 18, -17, 19, -15, 30,
	-27, 37, 9, -61, -62, -44, 37, -64, 99, 95,
	37, 94, 37, 37, -63, 99, 37, -63, 37, -18,
	-19, 78, -22, 37, -31, -36, -32, 72, 47, -35,
	-44, -40, -43, -41, -45, 21, 38, 39, 40, 26,
	36, 120, -42, 76, 77, 51, 99, 29, 83, 42,
	-27, 34, 81, -27, 58, 48, 81, 37, 72, 37,
	-71, 37, -71, 97, 37, 21, 69, 37, 9, 58,
	-20, 37, 20, 81, 47, 71, 70, -33, 22, 72,
	24, 25, 23, 73, 74, 75, 76, 77, 78, 79,
	80, 48, 49, 50, 43, 44, 45, 46, -31, -36,
	37, -31, -3, -38, -36, -36, 47, 47, -42, 47,
	38, 38, -48, -36, -58, 34, 47, -61, -70, 37,
	5, 6, 7, 8, 9, 10, 11, 12, 13, 14,
	15, 16, 17, 18, 19, 20, 21, 22, 23, 24,
	25, 26, 27, 28, 29, 30, 31, 32, 33, 34,
	35, 36, 52, 53, 54, 55, 56, 57, 59, 60,
	61, 62, 63, 64, 65, 66, 67, 68, 69, 70,
	71, 72, 83, 84, 85, 86, 87, 88, 89, 90,
	91, 92, 93, 94, 95, 96, 97, 98, 99, 100,
	101, 102, 103, -30, 10, -62, -36, -70, 47, 21,
	-69, 100, -66, 91, 89, 33, 90, 13, 37, 37,
	37, -71, -23, -24, -26, 47, 37, -42, -19, 37,
	78, 133, -18, 19, -31, -31, -36, -37, 47, -42,
	41, 22, 24, 25, -36, -36, -36, 72, -36, -36,
	-36, -36, -36, -36, -36, -36, 133, 133, 58, 133,
	-36, -18, -3, -46, -47, 84, -34, 29, -3, -61,
	-59, -44, -30, -52, 13, -31, -74, -73, 37, 69,
	37, -71, -67, 97, -30, 58, -25, 59, 60, 61,
	62, 63, 65, 66, -21, 37, 20, -24, 81, 133,
	-18, -38, -37, -36, -36, 71, -36, -36, 133, 133,
	-49, -47, 86, -31, -60, 69, -39, -40, -60, 133,
	58, -52, -56, 15, 14, 58, 133, -72, -78, -77,
	-85, -82, -83, 124, 125, 123, 126, 118, 119, 120,
	121, 122, 104, 105, 106, 107, 108, 109, 110, 111,
	112, 113, 116, 117, 37, 37, -50, 11, -24, -24,
	59, 64, 59, 64, 59, 59, 59, -28, 67, 98,
	68, 37, 133, -70, 133, 133, 71, -36, 87, -36,
	85, 31, 58, -44, -56, -36, -53, -54, -36, -73,
	-89, 93, -86, -79, 114, -76, 47, -76, -76, 127,
	-84, 47, -84, -84, -84, -76, 47, -84, -76, -71,
	-51, 12, 14, 69, 59, 59, 95, 95, 95, -36,
	-36, 32, -40, 58, 17, 58, -55, 27, 28, 48,
	72, 26, 33, 129, -81, -87, -88, 52, 32, 53,
	-80, 115, 39, 37, 39, 39, -52, -31, -38, -31,
	47, 47, 47, 7, -36, -36, -54, 37, 26, 38,
	39, 32, 32, 133, 58, -56, -29, 37, -29, -29,
	-61, 39, -57, 16, 35, 133, 58, 133, 133, 133,
	7, 22, 37, 37, 37,
}
var yyDef = [...]int{

	0, -2, 1, 2, 3, 4, 5, 6, 7, 8,
	9, 10, 11, 12, 91, 91, 91, 91, 91, 77,
	257, 248, 0, 0, 0, 335, 335, 335, 0, 95,
	97, 98, 99, 100, 93, 0, 0, 0, 0, 246,
	0, 0, 258, 0, 0, 249, 0, 244, 0, 244,
	0, 88, 89, 90, 15, 96, 0, 101, 92, 0,
	0, 133, 0, 20, 241, 0, 202, 0, 0, 0,
	335, 0, 335, 0, 0, 0, 0, 0, 87, 13,
	102, 104, 109, 202, 107, 108, 143, 0, 0, 171,
	172, 173, 0, 0, 188, 0, 204, 205, 206, 207,
	0, 0, 240, 191, 192, 193, 189, 190, 195, 94,
	229, 0, 0, 141, 0, 0, 0, 0, 0, 259,
	79, 0, 82, 0, 84, 245, 0, 335, 0, 0,
	105, 110, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 158, 159, 160, 161, 162, 163, 164, 146, 0,
	202, 0, 0, 0, 169, 182, 0, 0, 157, 0,
	208, 209, 0, 196, 0, 0, 0, 141, 134, 261,
	262, 263, 264, 265, 266, 267, 268, 269, 270, 271,
	272, 273, 274, 275, 276, 277, 278, 279, 280, 281,
	282, 283, 284, 285, 286, 287, 288, 289, 290, 291,
	292, 293, 294, 295, 296, 297, 298, 299, 300, 301,
	302, 303, 304, 305, 306, 307, 308, 309, 310, 311,
	312, 313, 314, 315, 316, 317, 318, 319, 320, 321,
	322, 323, 324, 325, 326, 327, 328, 329, 330, 331,
	332, 333, 334, 214, 0, 242, 243, 203, 0, 247,
	0, 0, 335, 255, 250, 251, 252, 253, 254, 83,
	85, 86, 141, 112, 118, 0, 130, 132, 103, 111,
	106, 183, 0, 0, 144, 145, 148, 149, 0, 166,
	167, 0, 0, 0, 151, 0, 155, 0, 174, 175,
	176, 177, 178, 179, 180, 181, 147, 168, 0, 239,
	169, 0, 0, 200, 197, 0, 233, 0, 236, 233,
	0, 231, 214, 222, 0, 142, 0, 72, 0, 0,
	260, 80, 0, 256, 210, 0, 0, 121, 122, 0,
	0, 0, 0, 0, 135, 119, 0, 0, 0, 184,
	0, 0, 150, 152, 0, 0, 156, 170, 186, 187,
	0, 198, 0, 0, 16, 0, 235, 237, 17, 230,
	0, 222, 19, 0, 0, 0, 75, 58, 56, 24,
	25, 54, 37, 54, 54, 33, 34, 26, 27, 28,
	29, 30, 38, 39, 40, 41, 42, 43, 44, 52,
	52, 52, 52, 52, 335, 81, 212, 0, 113, 116,
	123, 0, 125, 0, 127, 128, 129, 114, 0, 0,
	0, 120, 115, 131, 185, 165, 0, 153, 194, 201,
	0, 0, 0, 232, 18, 223, 215, 216, 219, 73,
	74, 0, 71, 21, 57, 36, 0, 31, 32, 0,
	45, 0, 46, 47, 48, 49, 0, 50, 51, 78,
	214, 0, 0, 0, 124, 126, 0, 0, 0, 154,
	199, 0, 238, 0, 0, 0, 218, 220, 221, 0,
	0, 60, 0, 63, 64, 65, 66, 0, 68, 69,
	23, 22, 0, 35, 0, 0, 222, 213, 211, 117,
	0, 0, 0, 0, 224, 225, 217, 76, 59, 61,
	62, 67, 70, 55, 0, 226, 0, 139, 0, 0,
	234, 0, 14, 0, 0, 136, 0, 137, 138, 53,
	227, 0, 140, 0, 228,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 80, 73, 3,
	47, 133, 78, 76, 58, 77, 81, 79, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	49, 48, 50, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 75, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 74, 3, 51,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 52, 53, 54, 55, 56,
	57, 59, 60, 61, 62, 63, 64, 65, 66, 67,
	68, 69, 70, 71, 72, 82, 83, 84, 85, 86,
	87, 88, 89, 90, 91, 92, 93, 94, 95, 96,
	97, 98, 99, 100, 101, 102, 103, 104, 105, 106,
	107, 108, 109, 110, 111, 112, 113, 114, 115, 116,
	117, 118, 119, 120, 121, 122, 123, 124, 125, 126,
	127, 128, 129, 130, 131, 132,
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
		//line sql.y:177
		{
			SetParseTree(yylex, yyDollar[1].statement)
		}
	case 2:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:183
		{
			yyVAL.statement = yyDollar[1].selStmt
		}
	case 13:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:199
		{
			yyVAL.selStmt = &Select{Comments: Comments(yyDollar[2].runes2), Distinct: yyDollar[3].str, SelectExprs: yyDollar[4].selectExprs}
		}
	case 14:
		yyDollar = yyS[yypt-12 : yypt+1]
		//line sql.y:203
		{
			yyVAL.selStmt = &Select{Comments: Comments(yyDollar[2].runes2), Distinct: yyDollar[3].str, SelectExprs: yyDollar[4].selectExprs, From: yyDollar[6].tableExprs, Where: NewWhere(AST_WHERE, yyDollar[7].boolExpr), GroupBy: GroupBy(yyDollar[8].valExprs), Having: NewWhere(AST_HAVING, yyDollar[9].boolExpr), OrderBy: yyDollar[10].orderBy, Limit: yyDollar[11].limit, Lock: yyDollar[12].str}
		}
	case 15:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:207
		{
			yyVAL.selStmt = &Union{Type: yyDollar[2].str, Left: yyDollar[1].selStmt, Right: yyDollar[3].selStmt}
		}
	case 16:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line sql.y:213
		{
			yyVAL.statement = &Insert{Comments: Comments(yyDollar[2].runes2), Table: yyDollar[4].tableName, Columns: yyDollar[5].columns, Rows: yyDollar[6].insRows, OnDup: OnDup(yyDollar[7].updateExprs)}
		}
	case 17:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line sql.y:217
		{
			cols := make(Columns, 0, len(yyDollar[6].updateExprs))
			vals := make(ValTuple, 0, len(yyDollar[6].updateExprs))
			for _, col := range yyDollar[6].updateExprs {
				cols = append(cols, &NonStarExpr{Expr: col.Name})
				vals = append(vals, col.Expr)
			}
			yyVAL.statement = &Insert{Comments: Comments(yyDollar[2].runes2), Table: yyDollar[4].tableName, Columns: cols, Rows: Values{vals}, OnDup: OnDup(yyDollar[7].updateExprs)}
		}
	case 18:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line sql.y:229
		{
			yyVAL.statement = &Update{Comments: Comments(yyDollar[2].runes2), Table: yyDollar[3].tableName, Exprs: yyDollar[5].updateExprs, Where: NewWhere(AST_WHERE, yyDollar[6].boolExpr), OrderBy: yyDollar[7].orderBy, Limit: yyDollar[8].limit}
		}
	case 19:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line sql.y:235
		{
			yyVAL.statement = &Delete{Comments: Comments(yyDollar[2].runes2), Table: yyDollar[4].tableName, Where: NewWhere(AST_WHERE, yyDollar[5].boolExpr), OrderBy: yyDollar[6].orderBy, Limit: yyDollar[7].limit}
		}
	case 20:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:241
		{
			yyVAL.statement = &Set{Comments: Comments(yyDollar[2].runes2), Exprs: yyDollar[3].updateExprs}
		}
	case 21:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:246
		{
			yyVAL.str = ""
		}
	case 22:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:250
		{
			yyVAL.str = AST_ZEROFILL
		}
	case 23:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:255
		{
			yyVAL.str = yyDollar[1].str
			if yyDollar[2].str != "" {
				yyVAL.str += " " + yyDollar[2].str
			}
			if yyDollar[3].str != "" {
				yyVAL.str += " " + yyDollar[3].str
			}
		}
	case 26:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:269
		{
			yyVAL.str = AST_DATE
		}
	case 27:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:273
		{
			yyVAL.str = AST_TIME
		}
	case 28:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:277
		{
			yyVAL.str = AST_TIMESTAMP
		}
	case 29:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:281
		{
			yyVAL.str = AST_DATETIME
		}
	case 30:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:285
		{
			yyVAL.str = AST_YEAR
		}
	case 31:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:291
		{
			if yyDollar[2].str == "" {
				yyVAL.str = AST_CHAR
			} else {
				yyVAL.str = AST_CHAR + yyDollar[2].str
			}
		}
	case 32:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:299
		{
			if yyDollar[2].str == "" {
				yyVAL.str = AST_VARCHAR
			} else {
				yyVAL.str = AST_VARCHAR + yyDollar[2].str
			}
		}
	case 33:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:307
		{
			yyVAL.str = AST_TEXT
		}
	case 34:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:311
		{
			yyVAL.str = AST_MEDIUMTEXT
		}
	case 35:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:315
		{
			yyVAL.str = AST_MEDIUMTEXT
			// do something with the charset id?
		}
	case 36:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:322
		{
			yyVAL.str = yyDollar[1].str + yyDollar[2].str
		}
	case 37:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:326
		{
			yyVAL.str = yyDollar[1].str
		}
	case 38:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:332
		{
			yyVAL.str = AST_BIT
		}
	case 39:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:336
		{
			yyVAL.str = AST_TINYINT
		}
	case 40:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:340
		{
			yyVAL.str = AST_SMALLINT
		}
	case 41:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:344
		{
			yyVAL.str = AST_MEDIUMINT
		}
	case 42:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:348
		{
			yyVAL.str = AST_INT
		}
	case 43:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:352
		{
			yyVAL.str = AST_INTEGER
		}
	case 44:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:356
		{
			yyVAL.str = AST_BIGINT
		}
	case 45:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:362
		{
			yyVAL.str = AST_REAL + yyDollar[2].str
		}
	case 46:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:366
		{
			yyVAL.str = AST_DOUBLE + yyDollar[2].str
		}
	case 47:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:370
		{
			yyVAL.str = AST_FLOAT + yyDollar[2].str
		}
	case 48:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:374
		{
			yyVAL.str = AST_DECIMAL + yyDollar[2].str
		}
	case 49:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:378
		{
			yyVAL.str = AST_DECIMAL + yyDollar[2].str
		}
	case 50:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:382
		{
			yyVAL.str = AST_NUMERIC + yyDollar[2].str
		}
	case 51:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:386
		{
			yyVAL.str = AST_NUMERIC + yyDollar[2].str
		}
	case 52:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:391
		{
			yyVAL.str = ""
		}
	case 53:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:395
		{
			yyVAL.str = "(" + string(yyDollar[2].runes) + ", " + string(yyDollar[4].runes) + ")"
		}
	case 54:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:400
		{
			yyVAL.str = ""
		}
	case 55:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:404
		{
			yyVAL.str = "(" + string(yyDollar[2].runes) + ")"
		}
	case 56:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:409
		{
			yyVAL.str = ""
		}
	case 57:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:413
		{
			yyVAL.str = AST_UNSIGNED
		}
	case 58:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:418
		{
			yyVAL.columnAtts = ColumnAtts{}
		}
	case 59:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:422
		{
			yyVAL.columnAtts = append(yyVAL.columnAtts, AST_NOT_NULL)
		}
	case 61:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:428
		{
			node := StrVal(yyDollar[3].runes)
			yyVAL.columnAtts = append(yyVAL.columnAtts, "default "+String(node))
		}
	case 62:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:433
		{
			node := NumVal(yyDollar[3].runes)
			yyVAL.columnAtts = append(yyVAL.columnAtts, "default "+String(node))
		}
	case 63:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:438
		{
			yyVAL.columnAtts = append(yyVAL.columnAtts, AST_AUTO_INCREMENT)
		}
	case 64:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:442
		{
			yyVAL.columnAtts = append(yyVAL.columnAtts, yyDollar[2].str)
		}
	case 65:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:448
		{
			yyVAL.str = AST_PRIMARY_KEY
		}
	case 66:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:452
		{
			yyVAL.str = AST_UNIQUE_KEY
		}
	case 71:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:466
		{
			yyVAL.columnDefinition = &ColumnDefinition{ColName: string(yyDollar[1].runes), ColType: yyDollar[2].str, ColumnAtts: yyDollar[3].columnAtts}
		}
	case 72:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:472
		{
			yyVAL.columnDefinitions = ColumnDefinitions{yyDollar[1].columnDefinition}
		}
	case 73:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:476
		{
			yyVAL.columnDefinitions = append(yyVAL.columnDefinitions, yyDollar[3].columnDefinition)
		}
	case 74:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line sql.y:482
		{
			yyVAL.statement = &CreateTable{Name: yyDollar[4].runes, ColumnDefinitions: yyDollar[6].columnDefinitions}
		}
	case 77:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:492
		{
			yyVAL.statement = yyDollar[1].statement
		}
	case 78:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line sql.y:496
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[7].runes, NewName: yyDollar[7].runes}
		}
	case 79:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:501
		{
			yyVAL.statement = &DDL{Action: AST_CREATE, NewName: yyDollar[3].runes}
		}
	case 80:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line sql.y:507
		{
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[4].runes, NewName: yyDollar[4].runes}
		}
	case 81:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line sql.y:511
		{
			// Change this to a rename statement
			yyVAL.statement = &DDL{Action: AST_RENAME, Table: yyDollar[4].runes, NewName: yyDollar[7].runes}
		}
	case 82:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:516
		{
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[3].runes, NewName: yyDollar[3].runes}
		}
	case 83:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:522
		{
			yyVAL.statement = &DDL{Action: AST_RENAME, Table: yyDollar[3].runes, NewName: yyDollar[5].runes}
		}
	case 84:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:528
		{
			yyVAL.statement = &DDL{Action: AST_DROP, Table: yyDollar[4].runes}
		}
	case 85:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:532
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[5].runes, NewName: yyDollar[5].runes}
		}
	case 86:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:537
		{
			yyVAL.statement = &DDL{Action: AST_DROP, Table: yyDollar[4].runes}
		}
	case 87:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:543
		{
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[3].runes, NewName: yyDollar[3].runes}
		}
	case 88:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:549
		{
			yyVAL.statement = &Other{}
		}
	case 89:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:553
		{
			yyVAL.statement = &Other{}
		}
	case 90:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:557
		{
			yyVAL.statement = &Other{}
		}
	case 91:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:562
		{
			SetAllowComments(yylex, true)
		}
	case 92:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:566
		{
			yyVAL.runes2 = yyDollar[2].runes2
			SetAllowComments(yylex, false)
		}
	case 93:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:572
		{
			yyVAL.runes2 = nil
		}
	case 94:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:576
		{
			yyVAL.runes2 = append(yyDollar[1].runes2, yyDollar[2].runes)
		}
	case 95:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:582
		{
			yyVAL.str = AST_UNION
		}
	case 96:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:586
		{
			yyVAL.str = AST_UNION_ALL
		}
	case 97:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:590
		{
			yyVAL.str = AST_SET_MINUS
		}
	case 98:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:594
		{
			yyVAL.str = AST_EXCEPT
		}
	case 99:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:598
		{
			yyVAL.str = AST_INTERSECT
		}
	case 100:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:603
		{
			yyVAL.str = ""
		}
	case 101:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:607
		{
			yyVAL.str = AST_DISTINCT
		}
	case 102:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:613
		{
			yyVAL.selectExprs = SelectExprs{yyDollar[1].selectExpr}
		}
	case 103:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:617
		{
			yyVAL.selectExprs = append(yyVAL.selectExprs, yyDollar[3].selectExpr)
		}
	case 104:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:623
		{
			yyVAL.selectExpr = &StarExpr{}
		}
	case 105:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:627
		{
			yyVAL.selectExpr = &NonStarExpr{Expr: yyDollar[1].expr, As: yyDollar[2].runes}
		}
	case 106:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:631
		{
			yyVAL.selectExpr = &StarExpr{TableName: yyDollar[1].runes}
		}
	case 107:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:637
		{
			yyVAL.expr = yyDollar[1].boolExpr
		}
	case 108:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:641
		{
			yyVAL.expr = yyDollar[1].valExpr
		}
	case 109:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:646
		{
			yyVAL.runes = nil
		}
	case 110:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:650
		{
			yyVAL.runes = yyDollar[1].runes
		}
	case 111:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:654
		{
			yyVAL.runes = yyDollar[2].runes
		}
	case 112:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:660
		{
			yyVAL.tableExprs = TableExprs{yyDollar[1].tableExpr}
		}
	case 113:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:664
		{
			yyVAL.tableExprs = append(yyVAL.tableExprs, yyDollar[3].tableExpr)
		}
	case 114:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:670
		{
			yyVAL.tableExpr = &AliasedTableExpr{Expr: yyDollar[1].smTableExpr, As: yyDollar[2].runes, Hints: yyDollar[3].indexHints}
		}
	case 115:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:674
		{
			yyVAL.tableExpr = &ParenTableExpr{Expr: yyDollar[2].tableExpr}
		}
	case 116:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:678
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr}
		}
	case 117:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:682
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr, On: yyDollar[5].boolExpr}
		}
	case 118:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:687
		{
			yyVAL.runes = nil
		}
	case 119:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:691
		{
			yyVAL.runes = yyDollar[1].runes
		}
	case 120:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:695
		{
			yyVAL.runes = yyDollar[2].runes
		}
	case 121:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:701
		{
			yyVAL.str = AST_JOIN
		}
	case 122:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:705
		{
			yyVAL.str = AST_STRAIGHT_JOIN
		}
	case 123:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:709
		{
			yyVAL.str = AST_LEFT_JOIN
		}
	case 124:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:713
		{
			yyVAL.str = AST_LEFT_JOIN
		}
	case 125:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:717
		{
			yyVAL.str = AST_RIGHT_JOIN
		}
	case 126:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:721
		{
			yyVAL.str = AST_RIGHT_JOIN
		}
	case 127:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:725
		{
			yyVAL.str = AST_JOIN
		}
	case 128:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:729
		{
			yyVAL.str = AST_CROSS_JOIN
		}
	case 129:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:733
		{
			yyVAL.str = AST_NATURAL_JOIN
		}
	case 130:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:739
		{
			yyVAL.smTableExpr = &TableName{Name: yyDollar[1].runes}
		}
	case 131:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:743
		{
			yyVAL.smTableExpr = &TableName{Qualifier: yyDollar[1].runes, Name: yyDollar[3].runes}
		}
	case 132:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:747
		{
			yyVAL.smTableExpr = yyDollar[1].subquery
		}
	case 133:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:753
		{
			yyVAL.tableName = &TableName{Name: yyDollar[1].runes}
		}
	case 134:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:757
		{
			yyVAL.tableName = &TableName{Qualifier: yyDollar[1].runes, Name: yyDollar[3].runes}
		}
	case 135:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:762
		{
			yyVAL.indexHints = nil
		}
	case 136:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:766
		{
			yyVAL.indexHints = &IndexHints{Type: AST_USE, Indexes: yyDollar[4].runes2}
		}
	case 137:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:770
		{
			yyVAL.indexHints = &IndexHints{Type: AST_IGNORE, Indexes: yyDollar[4].runes2}
		}
	case 138:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:774
		{
			yyVAL.indexHints = &IndexHints{Type: AST_FORCE, Indexes: yyDollar[4].runes2}
		}
	case 139:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:780
		{
			yyVAL.runes2 = [][]rune{yyDollar[1].runes}
		}
	case 140:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:784
		{
			yyVAL.runes2 = append(yyDollar[1].runes2, yyDollar[3].runes)
		}
	case 141:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:789
		{
			yyVAL.boolExpr = nil
		}
	case 142:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:793
		{
			yyVAL.boolExpr = yyDollar[2].boolExpr
		}
	case 144:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:800
		{
			yyVAL.boolExpr = &AndExpr{Left: yyDollar[1].boolExpr, Right: yyDollar[3].boolExpr}
		}
	case 145:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:804
		{
			yyVAL.boolExpr = &OrExpr{Left: yyDollar[1].boolExpr, Right: yyDollar[3].boolExpr}
		}
	case 146:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:808
		{
			yyVAL.boolExpr = &NotExpr{Expr: yyDollar[2].boolExpr}
		}
	case 147:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:812
		{
			yyVAL.boolExpr = &ParenBoolExpr{Expr: yyDollar[2].boolExpr}
		}
	case 148:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:818
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: yyDollar[2].str, Right: yyDollar[3].valExpr}
		}
	case 149:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:822
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_IN, Right: yyDollar[3].colTuple}
		}
	case 150:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:826
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_NOT_IN, Right: yyDollar[4].colTuple}
		}
	case 151:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:830
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_LIKE, Right: yyDollar[3].valExpr}
		}
	case 152:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:834
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_NOT_LIKE, Right: yyDollar[4].valExpr}
		}
	case 153:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:838
		{
			yyVAL.boolExpr = &RangeCond{Left: yyDollar[1].valExpr, Operator: AST_BETWEEN, From: yyDollar[3].valExpr, To: yyDollar[5].valExpr}
		}
	case 154:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line sql.y:842
		{
			yyVAL.boolExpr = &RangeCond{Left: yyDollar[1].valExpr, Operator: AST_NOT_BETWEEN, From: yyDollar[4].valExpr, To: yyDollar[6].valExpr}
		}
	case 155:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:846
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_IS, Right: yyDollar[3].valExpr}
		}
	case 156:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:850
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_IS_NOT, Right: yyDollar[4].valExpr}
		}
	case 157:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:854
		{
			yyVAL.boolExpr = &ExistsExpr{Subquery: yyDollar[2].subquery}
		}
	case 158:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:860
		{
			yyVAL.str = AST_EQ
		}
	case 159:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:864
		{
			yyVAL.str = AST_LT
		}
	case 160:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:868
		{
			yyVAL.str = AST_GT
		}
	case 161:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:872
		{
			yyVAL.str = AST_LE
		}
	case 162:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:876
		{
			yyVAL.str = AST_GE
		}
	case 163:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:880
		{
			yyVAL.str = AST_NE
		}
	case 164:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:884
		{
			yyVAL.str = AST_NSE
		}
	case 165:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:890
		{
			yyVAL.colTuple = ValTuple(yyDollar[2].valExprs)
		}
	case 166:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:894
		{
			yyVAL.colTuple = yyDollar[1].subquery
		}
	case 167:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:898
		{
			yyVAL.colTuple = ListArg(yyDollar[1].runes)
		}
	case 168:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:904
		{
			yyVAL.subquery = &Subquery{yyDollar[2].selStmt}
		}
	case 169:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:910
		{
			yyVAL.valExprs = ValExprs{yyDollar[1].valExpr}
		}
	case 170:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:914
		{
			yyVAL.valExprs = append(yyDollar[1].valExprs, yyDollar[3].valExpr)
		}
	case 171:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:920
		{
			yyVAL.valExpr = yyDollar[1].valExpr
		}
	case 172:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:924
		{
			yyVAL.valExpr = yyDollar[1].colName
		}
	case 173:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:928
		{
			yyVAL.valExpr = yyDollar[1].rowTuple
		}
	case 174:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:932
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_BITAND, Right: yyDollar[3].valExpr}
		}
	case 175:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:936
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_BITOR, Right: yyDollar[3].valExpr}
		}
	case 176:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:940
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_BITXOR, Right: yyDollar[3].valExpr}
		}
	case 177:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:944
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_PLUS, Right: yyDollar[3].valExpr}
		}
	case 178:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:948
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_MINUS, Right: yyDollar[3].valExpr}
		}
	case 179:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:952
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_MULT, Right: yyDollar[3].valExpr}
		}
	case 180:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:956
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_DIV, Right: yyDollar[3].valExpr}
		}
	case 181:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:960
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_MOD, Right: yyDollar[3].valExpr}
		}
	case 182:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:982
		{
			if num, ok := yyDollar[2].valExpr.(NumVal); ok {
				switch yyDollar[1].run {
				case '-':
					yyVAL.valExpr = append(NumVal("-"), num...)
				case '+':
					yyVAL.valExpr = num
				default:
					yyVAL.valExpr = &UnaryExpr{Operator: yyDollar[1].run, Expr: yyDollar[2].valExpr}
				}
			} else {
				yyVAL.valExpr = &UnaryExpr{Operator: yyDollar[1].run, Expr: yyDollar[2].valExpr}
			}
		}
	case 183:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:997
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].runes}
		}
	case 184:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:1001
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].runes, Exprs: yyDollar[3].selectExprs}
		}
	case 185:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:1005
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].runes, Distinct: true, Exprs: yyDollar[4].selectExprs}
		}
	case 186:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:1009
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].runes, Exprs: yyDollar[3].selectExprs}
		}
	case 187:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:1013
		{
			// XXX(toshok) select_statement is dropped here
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].runes}
		}
	case 188:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1018
		{
			yyVAL.valExpr = yyDollar[1].caseExpr
		}
	case 189:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1024
		{
			yyVAL.runes = yyDollar[1].runes
		}
	case 190:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1028
		{
			yyVAL.runes = yyDollar[1].runes
		}
	case 191:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1034
		{
			yyVAL.run = AST_UPLUS
		}
	case 192:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1038
		{
			yyVAL.run = AST_UMINUS
		}
	case 193:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1042
		{
			yyVAL.run = AST_TILDA
		}
	case 194:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:1048
		{
			yyVAL.caseExpr = &CaseExpr{Expr: yyDollar[2].valExpr, Whens: yyDollar[3].whens, Else: yyDollar[4].valExpr}
		}
	case 195:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1053
		{
			yyVAL.valExpr = nil
		}
	case 196:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1057
		{
			yyVAL.valExpr = yyDollar[1].valExpr
		}
	case 197:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1063
		{
			yyVAL.whens = []*When{yyDollar[1].when}
		}
	case 198:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1067
		{
			yyVAL.whens = append(yyDollar[1].whens, yyDollar[2].when)
		}
	case 199:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:1073
		{
			yyVAL.when = &When{Cond: yyDollar[2].boolExpr, Val: yyDollar[4].valExpr}
		}
	case 200:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1078
		{
			yyVAL.valExpr = nil
		}
	case 201:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1082
		{
			yyVAL.valExpr = yyDollar[2].valExpr
		}
	case 202:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1088
		{
			yyVAL.colName = &ColName{Name: yyDollar[1].runes}
		}
	case 203:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1092
		{
			yyVAL.colName = &ColName{Qualifier: yyDollar[1].runes, Name: yyDollar[3].runes}
		}
	case 204:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1098
		{
			yyVAL.valExpr = StrVal(yyDollar[1].runes)
		}
	case 205:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1102
		{
			yyVAL.valExpr = NumVal(yyDollar[1].runes)
		}
	case 206:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1106
		{
			yyVAL.valExpr = ValArg(yyDollar[1].runes)
		}
	case 207:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1110
		{
			yyVAL.valExpr = &NullVal{}
		}
	case 208:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1114
		{
			yyVAL.valExpr = BinaryVal(yyDollar[2].runes)
		}
	case 209:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1118
		{
			yyVAL.valExpr = TimestampVal(yyDollar[2].runes)
		}
	case 210:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1123
		{
			yyVAL.valExprs = nil
		}
	case 211:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1127
		{
			yyVAL.valExprs = yyDollar[3].valExprs
		}
	case 212:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1132
		{
			yyVAL.boolExpr = nil
		}
	case 213:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1136
		{
			yyVAL.boolExpr = yyDollar[2].boolExpr
		}
	case 214:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1141
		{
			yyVAL.orderBy = nil
		}
	case 215:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1145
		{
			yyVAL.orderBy = yyDollar[3].orderBy
		}
	case 216:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1151
		{
			yyVAL.orderBy = OrderBy{yyDollar[1].order}
		}
	case 217:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1155
		{
			yyVAL.orderBy = append(yyDollar[1].orderBy, yyDollar[3].order)
		}
	case 218:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1161
		{
			yyVAL.order = &Order{Expr: yyDollar[1].valExpr, Direction: yyDollar[2].str}
		}
	case 219:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1166
		{
			yyVAL.str = AST_ASC
		}
	case 220:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1170
		{
			yyVAL.str = AST_ASC
		}
	case 221:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1174
		{
			yyVAL.str = AST_DESC
		}
	case 222:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1179
		{
			yyVAL.limit = nil
		}
	case 223:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1183
		{
			yyVAL.limit = &Limit{Rowcount: yyDollar[2].valExpr}
		}
	case 224:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:1187
		{
			yyVAL.limit = &Limit{Offset: yyDollar[2].valExpr, Rowcount: yyDollar[4].valExpr}
		}
	case 225:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:1191
		{
			yyVAL.limit = &Limit{Offset: yyDollar[4].valExpr, Rowcount: yyDollar[2].valExpr}
		}
	case 226:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1196
		{
			yyVAL.str = ""
		}
	case 227:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1200
		{
			yyVAL.str = AST_FOR_UPDATE
		}
	case 228:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:1204
		{
			if string(yyDollar[3].runes) != SHARE {
				yylex.Error("expecting share")
				return 1
			}
			if string(yyDollar[4].runes) != MODE {
				yylex.Error("expecting mode")
				return 1
			}
			yyVAL.str = AST_SHARE_MODE
		}
	case 229:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1217
		{
			yyVAL.columns = nil
		}
	case 230:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1221
		{
			yyVAL.columns = yyDollar[2].columns
		}
	case 231:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1227
		{
			yyVAL.columns = Columns{&NonStarExpr{Expr: yyDollar[1].colName}}
		}
	case 232:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1231
		{
			yyVAL.columns = append(yyVAL.columns, &NonStarExpr{Expr: yyDollar[3].colName})
		}
	case 233:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1236
		{
			yyVAL.updateExprs = nil
		}
	case 234:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:1240
		{
			yyVAL.updateExprs = yyDollar[5].updateExprs
		}
	case 235:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1246
		{
			yyVAL.insRows = yyDollar[2].values
		}
	case 236:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1250
		{
			yyVAL.insRows = yyDollar[1].selStmt
		}
	case 237:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1256
		{
			yyVAL.values = Values{yyDollar[1].rowTuple}
		}
	case 238:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1260
		{
			yyVAL.values = append(yyDollar[1].values, yyDollar[3].rowTuple)
		}
	case 239:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1266
		{
			yyVAL.rowTuple = ValTuple(yyDollar[2].valExprs)
		}
	case 240:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1270
		{
			yyVAL.rowTuple = yyDollar[1].subquery
		}
	case 241:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1276
		{
			yyVAL.updateExprs = UpdateExprs{yyDollar[1].updateExpr}
		}
	case 242:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1280
		{
			yyVAL.updateExprs = append(yyDollar[1].updateExprs, yyDollar[3].updateExpr)
		}
	case 243:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1286
		{
			yyVAL.updateExpr = &UpdateExpr{Name: yyDollar[1].colName, Expr: yyDollar[3].valExpr}
		}
	case 244:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1291
		{
			yyVAL.empty = struct{}{}
		}
	case 245:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1293
		{
			yyVAL.empty = struct{}{}
		}
	case 246:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1296
		{
			yyVAL.empty = struct{}{}
		}
	case 247:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1298
		{
			yyVAL.empty = struct{}{}
		}
	case 248:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1301
		{
			yyVAL.empty = struct{}{}
		}
	case 249:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1303
		{
			yyVAL.empty = struct{}{}
		}
	case 250:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1307
		{
			yyVAL.empty = struct{}{}
		}
	case 251:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1309
		{
			yyVAL.empty = struct{}{}
		}
	case 252:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1311
		{
			yyVAL.empty = struct{}{}
		}
	case 253:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1313
		{
			yyVAL.empty = struct{}{}
		}
	case 254:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1315
		{
			yyVAL.empty = struct{}{}
		}
	case 255:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1318
		{
			yyVAL.empty = struct{}{}
		}
	case 256:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1320
		{
			yyVAL.empty = struct{}{}
		}
	case 257:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1323
		{
			yyVAL.empty = struct{}{}
		}
	case 258:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1325
		{
			yyVAL.empty = struct{}{}
		}
	case 259:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1328
		{
			yyVAL.empty = struct{}{}
		}
	case 260:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1330
		{
			yyVAL.empty = struct{}{}
		}
	case 261:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1333
		{
			yyVAL.runes = yyDollar[1].runes
		}
	case 262:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1334
		{
			yyVAL.runes = yyDollar[1].runes
		}
	case 263:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1335
		{
			yyVAL.runes = yyDollar[1].runes
		}
	case 264:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1336
		{
			yyVAL.runes = yyDollar[1].runes
		}
	case 265:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1337
		{
			yyVAL.runes = yyDollar[1].runes
		}
	case 266:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1338
		{
			yyVAL.runes = yyDollar[1].runes
		}
	case 267:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1339
		{
			yyVAL.runes = yyDollar[1].runes
		}
	case 268:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1340
		{
			yyVAL.runes = yyDollar[1].runes
		}
	case 269:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1341
		{
			yyVAL.runes = yyDollar[1].runes
		}
	case 270:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1342
		{
			yyVAL.runes = yyDollar[1].runes
		}
	case 271:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1343
		{
			yyVAL.runes = yyDollar[1].runes
		}
	case 272:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1344
		{
			yyVAL.runes = yyDollar[1].runes
		}
	case 273:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1345
		{
			yyVAL.runes = yyDollar[1].runes
		}
	case 274:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1346
		{
			yyVAL.runes = yyDollar[1].runes
		}
	case 275:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1347
		{
			yyVAL.runes = yyDollar[1].runes
		}
	case 276:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1348
		{
			yyVAL.runes = yyDollar[1].runes
		}
	case 277:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1349
		{
			yyVAL.runes = yyDollar[1].runes
		}
	case 278:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1350
		{
			yyVAL.runes = yyDollar[1].runes
		}
	case 279:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1351
		{
			yyVAL.runes = yyDollar[1].runes
		}
	case 280:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1352
		{
			yyVAL.runes = yyDollar[1].runes
		}
	case 281:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1353
		{
			yyVAL.runes = yyDollar[1].runes
		}
	case 282:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1354
		{
			yyVAL.runes = yyDollar[1].runes
		}
	case 283:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1355
		{
			yyVAL.runes = yyDollar[1].runes
		}
	case 284:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1356
		{
			yyVAL.runes = yyDollar[1].runes
		}
	case 285:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1357
		{
			yyVAL.runes = yyDollar[1].runes
		}
	case 286:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1358
		{
			yyVAL.runes = yyDollar[1].runes
		}
	case 287:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1359
		{
			yyVAL.runes = yyDollar[1].runes
		}
	case 288:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1360
		{
			yyVAL.runes = yyDollar[1].runes
		}
	case 289:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1361
		{
			yyVAL.runes = yyDollar[1].runes
		}
	case 290:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1362
		{
			yyVAL.runes = yyDollar[1].runes
		}
	case 291:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1363
		{
			yyVAL.runes = yyDollar[1].runes
		}
	case 292:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1364
		{
			yyVAL.runes = yyDollar[1].runes
		}
	case 293:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1365
		{
			yyVAL.runes = yyDollar[1].runes
		}
	case 294:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1366
		{
			yyVAL.runes = yyDollar[1].runes
		}
	case 295:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1367
		{
			yyVAL.runes = yyDollar[1].runes
		}
	case 296:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1368
		{
			yyVAL.runes = yyDollar[1].runes
		}
	case 297:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1369
		{
			yyVAL.runes = yyDollar[1].runes
		}
	case 298:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1370
		{
			yyVAL.runes = yyDollar[1].runes
		}
	case 299:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1371
		{
			yyVAL.runes = yyDollar[1].runes
		}
	case 300:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1372
		{
			yyVAL.runes = yyDollar[1].runes
		}
	case 301:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1373
		{
			yyVAL.runes = yyDollar[1].runes
		}
	case 302:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1374
		{
			yyVAL.runes = yyDollar[1].runes
		}
	case 303:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1375
		{
			yyVAL.runes = yyDollar[1].runes
		}
	case 304:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1376
		{
			yyVAL.runes = yyDollar[1].runes
		}
	case 305:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1377
		{
			yyVAL.runes = yyDollar[1].runes
		}
	case 306:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1378
		{
			yyVAL.runes = yyDollar[1].runes
		}
	case 307:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1379
		{
			yyVAL.runes = yyDollar[1].runes
		}
	case 308:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1380
		{
			yyVAL.runes = yyDollar[1].runes
		}
	case 309:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1381
		{
			yyVAL.runes = yyDollar[1].runes
		}
	case 310:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1382
		{
			yyVAL.runes = yyDollar[1].runes
		}
	case 311:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1383
		{
			yyVAL.runes = yyDollar[1].runes
		}
	case 312:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1384
		{
			yyVAL.runes = yyDollar[1].runes
		}
	case 313:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1385
		{
			yyVAL.runes = yyDollar[1].runes
		}
	case 314:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1386
		{
			yyVAL.runes = yyDollar[1].runes
		}
	case 315:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1387
		{
			yyVAL.runes = yyDollar[1].runes
		}
	case 316:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1388
		{
			yyVAL.runes = yyDollar[1].runes
		}
	case 317:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1389
		{
			yyVAL.runes = yyDollar[1].runes
		}
	case 318:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1390
		{
			yyVAL.runes = yyDollar[1].runes
		}
	case 319:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1391
		{
			yyVAL.runes = yyDollar[1].runes
		}
	case 320:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1392
		{
			yyVAL.runes = yyDollar[1].runes
		}
	case 321:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1393
		{
			yyVAL.runes = yyDollar[1].runes
		}
	case 322:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1394
		{
			yyVAL.runes = yyDollar[1].runes
		}
	case 323:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1395
		{
			yyVAL.runes = yyDollar[1].runes
		}
	case 324:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1396
		{
			yyVAL.runes = yyDollar[1].runes
		}
	case 325:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1397
		{
			yyVAL.runes = yyDollar[1].runes
		}
	case 326:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1398
		{
			yyVAL.runes = yyDollar[1].runes
		}
	case 327:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1399
		{
			yyVAL.runes = yyDollar[1].runes
		}
	case 328:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1400
		{
			yyVAL.runes = yyDollar[1].runes
		}
	case 329:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1401
		{
			yyVAL.runes = yyDollar[1].runes
		}
	case 330:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1402
		{
			yyVAL.runes = yyDollar[1].runes
		}
	case 331:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1403
		{
			yyVAL.runes = yyDollar[1].runes
		}
	case 332:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1404
		{
			yyVAL.runes = yyDollar[1].runes
		}
	case 333:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1405
		{
			yyVAL.runes = yyDollar[1].runes
		}
	case 334:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1406
		{
			yyVAL.runes = yyDollar[1].runes
		}
	case 335:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1409
		{
			ForceEOF(yylex)
		}
	}
	goto yystack /* stack new state and value */
}
