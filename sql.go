//line sql.y:6
package sqlparser

import __yyfmt__ "fmt"

//line sql.y:6
import "bytes"

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
	SHARE        = []byte("share")
	MODE         = []byte("mode")
	IF_BYTES     = []byte("if")
	VALUES_BYTES = []byte("values")
)

//line sql.y:31
type yySymType struct {
	yys         int
	empty       struct{}
	statement   Statement
	selStmt     SelectStatement
	byt         byte
	bytes       []byte
	bytes2      [][]byte
	runes       []rune
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

const yyNprod = 262
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 757

var yyAct = [...]int{

	95, 298, 160, 442, 92, 363, 249, 93, 63, 163,
	376, 253, 371, 290, 199, 210, 240, 81, 179, 452,
	86, 455, 162, 3, 137, 136, 51, 439, 263, 264,
	265, 266, 267, 104, 268, 269, 82, 452, 375, 66,
	417, 452, 71, 65, 370, 74, 64, 187, 440, 78,
	131, 54, 231, 52, 53, 301, 296, 77, 131, 87,
	29, 30, 31, 32, 29, 30, 31, 32, 394, 407,
	69, 121, 344, 346, 259, 414, 408, 44, 73, 45,
	129, 47, 48, 49, 131, 133, 231, 229, 125, 354,
	393, 392, 367, 164, 454, 413, 415, 165, 122, 42,
	241, 124, 348, 345, 70, 50, 46, 241, 274, 288,
	159, 161, 453, 173, 66, 406, 451, 66, 65, 183,
	182, 65, 177, 439, 120, 353, 135, 350, 118, 114,
	302, 295, 169, 284, 87, 205, 183, 72, 136, 285,
	39, 209, 41, 230, 217, 218, 219, 221, 222, 223,
	224, 225, 226, 227, 228, 193, 197, 207, 208, 282,
	389, 232, 291, 181, 203, 149, 150, 151, 204, 233,
	87, 87, 409, 212, 255, 191, 66, 66, 206, 194,
	65, 247, 137, 136, 245, 235, 237, 128, 256, 391,
	390, 238, 342, 341, 403, 404, 248, 244, 137, 136,
	351, 251, 144, 145, 146, 147, 148, 149, 150, 151,
	340, 116, 116, 356, 233, 257, 273, 260, 277, 278,
	231, 275, 291, 280, 147, 148, 149, 150, 151, 60,
	276, 190, 192, 189, 281, 203, 180, 14, 180, 87,
	144, 145, 146, 147, 148, 149, 150, 151, 212, 130,
	440, 293, 338, 98, 283, 297, 287, 339, 102, 294,
	336, 109, 289, 401, 358, 337, 405, 175, 103, 85,
	99, 100, 101, 117, 76, 202, 334, 335, 428, 90,
	176, 427, 352, 107, 261, 201, 116, 400, 14, 112,
	426, 355, 115, 166, 213, 203, 203, 66, 131, 360,
	211, 359, 361, 364, 89, 382, 377, 372, 105, 106,
	83, 184, 170, 365, 168, 110, 29, 30, 31, 32,
	202, 167, 373, 374, 79, 447, 111, 421, 399, 420,
	201, 108, 435, 436, 418, 272, 171, 378, 379, 380,
	383, 381, 384, 144, 145, 146, 147, 148, 149, 150,
	151, 134, 271, 72, 395, 67, 254, 385, 349, 396,
	347, 331, 236, 330, 98, 196, 398, 195, 72, 102,
	178, 126, 109, 123, 119, 61, 419, 80, 75, 103,
	85, 99, 100, 101, 113, 438, 437, 397, 357, 59,
	90, 233, 434, 422, 107, 457, 185, 449, 424, 14,
	127, 57, 430, 431, 364, 55, 433, 432, 423, 388,
	425, 14, 15, 16, 17, 89, 450, 299, 300, 105,
	106, 83, 250, 243, 441, 387, 110, 443, 443, 443,
	66, 444, 445, 214, 65, 215, 216, 333, 446, 180,
	18, 279, 108, 144, 145, 146, 147, 148, 149, 150,
	151, 62, 456, 458, 429, 14, 34, 366, 459, 412,
	460, 411, 368, 306, 308, 318, 319, 320, 321, 322,
	323, 324, 325, 326, 327, 307, 234, 328, 329, 313,
	314, 315, 316, 317, 311, 309, 310, 312, 410, 416,
	369, 304, 305, 19, 20, 21, 23, 22, 24, 252,
	303, 186, 40, 258, 188, 98, 43, 25, 26, 27,
	102, 33, 68, 109, 246, 174, 448, 14, 402, 362,
	103, 85, 99, 100, 101, 386, 332, 35, 36, 37,
	38, 90, 286, 98, 172, 107, 239, 97, 102, 94,
	96, 109, 292, 91, 242, 138, 88, 343, 103, 67,
	99, 100, 101, 200, 262, 198, 89, 84, 270, 90,
	105, 106, 83, 107, 132, 56, 28, 110, 144, 145,
	146, 147, 148, 149, 150, 151, 58, 13, 12, 11,
	10, 9, 98, 108, 89, 14, 8, 102, 105, 106,
	109, 7, 6, 5, 4, 110, 2, 103, 67, 99,
	100, 101, 1, 0, 0, 0, 102, 0, 90, 109,
	0, 108, 107, 0, 0, 0, 103, 67, 99, 100,
	101, 0, 263, 264, 265, 266, 267, 166, 268, 269,
	0, 107, 0, 89, 0, 0, 0, 105, 106, 0,
	102, 0, 0, 109, 110, 0, 0, 0, 0, 0,
	103, 67, 99, 100, 101, 0, 105, 106, 0, 102,
	108, 166, 109, 110, 0, 107, 0, 0, 0, 103,
	67, 99, 100, 101, 0, 0, 0, 0, 0, 108,
	166, 0, 0, 0, 107, 0, 220, 0, 0, 0,
	105, 106, 0, 0, 0, 0, 0, 110, 139, 143,
	141, 142, 0, 0, 0, 0, 0, 0, 0, 105,
	106, 0, 0, 108, 0, 0, 110, 0, 0, 155,
	156, 157, 158, 0, 152, 153, 154, 0, 0, 0,
	0, 0, 108, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 140, 144,
	145, 146, 147, 148, 149, 150, 151,
}
var yyPact = [...]int{

	406, -1000, -1000, 262, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	46, -19, 12, -13, 11, -1000, -1000, -1000, 450, 387,
	-1000, -1000, -1000, 382, -1000, 359, 338, 442, 318, -29,
	9, 316, -1000, -16, 316, -1000, 341, -42, 316, -42,
	340, -1000, -1000, -1000, -1000, -1000, 484, -1000, 284, 338,
	350, 48, 338, 154, -1000, 225, -1000, 47, 337, 52,
	316, -1000, -1000, 336, -1000, -9, 334, 379, 118, 316,
	-1000, 240, -1000, -1000, 331, 45, 112, 676, -1000, 561,
	512, -1000, -1000, -1000, 633, 274, 267, -1000, 265, -1000,
	-1000, -1000, -1000, 298, -1000, -1000, -1000, -1000, -1000, -1000,
	633, -1000, 233, 318, 333, 429, 318, 633, 316, 264,
	375, -53, -1000, 142, -1000, 330, -1000, -1000, 328, -1000,
	238, 484, -1000, -1000, 316, 100, 561, 561, 633, 253,
	411, 633, 633, 614, 633, 633, 633, 633, 633, 633,
	633, 633, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	676, -46, 10, 28, 676, -1000, 580, 343, 232, -1000,
	450, -1000, 16, 495, 394, 318, 318, 228, -1000, 409,
	561, -1000, 495, -1000, 319, -1000, 105, 316, -1000, -23,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 226, 563,
	315, 283, 27, -1000, -1000, -1000, -1000, -1000, 67, 495,
	-1000, 580, -1000, -1000, 253, 633, 633, 495, 370, 495,
	633, 148, 148, 148, 87, 87, -1000, -1000, -1000, -1000,
	-1000, 633, -1000, 495, -1000, 26, 484, 0, 6, 23,
	-1000, 561, 93, 246, 262, 153, -2, -1000, 409, 402,
	404, 112, -3, -1000, 361, 326, -1000, -1000, 324, -1000,
	426, 238, 238, -1000, -1000, 201, 193, 151, 134, 133,
	5, -1000, 323, -31, 321, -6, -1000, 495, 129, 633,
	495, 495, -1000, -8, -1000, -1000, 2, -1000, 633, 128,
	-1000, 357, 206, -1000, -1000, -1000, 318, 402, -1000, 633,
	633, 319, -1, -1000, -70, -1000, -1000, 260, -1000, 260,
	260, -1000, -89, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, 259, 259, 259, 258, 258,
	-1000, -1000, 413, 395, 563, 91, -1000, 131, -1000, 130,
	-1000, -1000, -1000, -1000, -4, -5, -27, -1000, -1000, -1000,
	-1000, 633, 495, -1000, -1000, 495, 633, 355, 246, -1000,
	-1000, 270, 205, -1000, 167, -1000, -1000, 218, 43, -75,
	-1000, -1000, 295, -1000, -1000, 316, -1000, 290, -1000, -1000,
	-1000, -1000, 288, -1000, -1000, -1000, 409, 561, 633, 561,
	-1000, -1000, 243, 234, 231, 495, 495, 447, -1000, 633,
	633, 633, -1000, -1000, -1000, 316, 366, -1000, 294, -1000,
	-1000, -1000, -1000, 354, -1000, 353, -1000, -1000, -106, -1000,
	192, -10, 402, 112, 162, 112, 316, 316, 316, 318,
	495, 495, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	286, 381, -17, -1000, -21, -39, 154, -112, -1000, 445,
	373, -1000, 316, -1000, -1000, -1000, -1000, 316, -1000, 316,
	-1000,
}
var yyPgo = [...]int{

	0, 602, 596, 22, 594, 593, 592, 591, 586, 581,
	580, 579, 578, 577, 511, 576, 566, 565, 17, 36,
	564, 558, 557, 555, 14, 554, 553, 229, 547, 3,
	18, 20, 546, 545, 544, 543, 2, 15, 9, 542,
	7, 540, 33, 539, 4, 537, 536, 16, 534, 532,
	526, 525, 6, 519, 5, 518, 1, 516, 515, 514,
	13, 8, 46, 274, 512, 506, 504, 503, 502, 501,
	0, 26, 500, 11, 499, 493, 12, 492, 491, 490,
	489, 488, 475, 464, 10, 463, 462, 461, 459, 457,
	456,
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
	49, 49, 44, 44, 35, 35, 35, 35, 35, 50,
	50, 51, 51, 52, 52, 53, 53, 54, 55, 55,
	55, 56, 56, 56, 56, 57, 57, 57, 58, 58,
	59, 59, 60, 60, 34, 34, 39, 39, 40, 40,
	61, 61, 62, 63, 63, 64, 64, 65, 65, 66,
	66, 66, 66, 66, 67, 67, 68, 68, 69, 69,
	70, 71,
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
	0, 2, 1, 3, 1, 1, 1, 1, 2, 0,
	3, 0, 2, 0, 3, 1, 3, 2, 0, 1,
	1, 0, 2, 4, 4, 0, 2, 4, 0, 3,
	1, 3, 0, 5, 2, 1, 1, 3, 3, 1,
	1, 3, 3, 0, 2, 0, 3, 0, 1, 1,
	1, 1, 1, 1, 0, 1, 0, 1, 0, 2,
	1, 0,
}
var yyChk = [...]int{

	-1000, -1, -2, -3, -4, -5, -6, -7, -8, -9,
	-10, -11, -12, -13, 5, 6, 7, 8, 34, -75,
	88, 89, 91, 90, 92, 101, 102, 103, -16, 54,
	55, 56, 57, -14, -90, -14, -14, -14, -14, 94,
	-68, 96, 53, -65, 96, 98, 94, 94, 95, 96,
	94, -71, -71, -71, -3, 18, -17, 19, -15, 30,
	-27, 37, 9, -61, -62, -44, -70, 37, -64, 99,
	95, -70, 37, 94, -70, 37, -63, 99, -70, -63,
	37, -18, -19, 78, -22, 37, -31, -36, -32, 72,
	47, -35, -44, -40, -43, -70, -41, -45, 21, 38,
	39, 40, 26, 36, -42, 76, 77, 51, 99, 29,
	83, 42, -27, 34, 81, -27, 58, 48, 81, 37,
	72, -70, -71, 37, -71, 97, 37, 21, 69, -70,
	9, 58, -20, -70, 20, 81, 71, 70, -33, 22,
	72, 24, 25, 23, 73, 74, 75, 76, 77, 78,
	79, 80, 48, 49, 50, 43, 44, 45, 46, -31,
	-36, -31, -3, -38, -36, -36, 47, 47, 47, -42,
	47, 38, -48, -36, -58, 34, 47, -61, 37, -30,
	10, -62, -36, -70, 47, 21, -69, 100, -66, 91,
	89, 33, 90, 13, 37, 37, 37, -71, -23, -24,
	-26, 47, 37, -42, -19, -70, 78, -31, -31, -36,
	-37, 47, -42, 41, 22, 24, 25, -36, -36, -36,
	72, -36, -36, -36, -36, -36, -36, -36, -36, 133,
	133, 58, 133, -36, 133, -18, 19, -18, -3, -46,
	-47, 84, -34, 29, -3, -61, -59, -44, -30, -52,
	13, -31, -74, -73, 37, 69, -70, -71, -67, 97,
	-30, 58, -25, 59, 60, 61, 62, 63, 65, 66,
	-21, 37, 20, -24, 81, -38, -37, -36, -36, 71,
	-36, -36, 133, -18, 133, 133, -49, -47, 86, -31,
	-60, 69, -39, -40, -60, 133, 58, -52, -56, 15,
	14, 58, 133, -72, -78, -77, -85, -82, -83, 124,
	125, 123, 126, 118, 119, 120, 121, 122, 104, 105,
	106, 107, 108, 109, 110, 111, 112, 113, 116, 117,
	37, 37, -50, 11, -24, -24, 59, 64, 59, 64,
	59, 59, 59, -28, 67, 98, 68, 37, 133, 37,
	133, 71, -36, 133, 87, -36, 85, 31, 58, -44,
	-56, -36, -53, -54, -36, -73, -89, 93, -86, -79,
	114, -76, 47, -76, -76, 127, -84, 47, -84, -84,
	-84, -76, 47, -84, -76, -71, -51, 12, 14, 69,
	59, 59, 95, 95, 95, -36, -36, 32, -40, 58,
	17, 58, -55, 27, 28, 48, 72, 26, 33, 129,
	-81, -87, -88, 52, 32, 53, -80, 115, 39, -70,
	39, 39, -52, -31, -38, -31, 47, 47, 47, 7,
	-36, -36, -54, -70, 26, 38, 39, 32, 32, 133,
	58, -56, -29, -70, -29, -29, -61, 39, -57, 16,
	35, 133, 58, 133, 133, 133, 7, 22, -70, -70,
	-70,
}
var yyDef = [...]int{

	0, -2, 1, 2, 3, 4, 5, 6, 7, 8,
	9, 10, 11, 12, 91, 91, 91, 91, 91, 77,
	256, 247, 0, 0, 0, 261, 261, 261, 0, 95,
	97, 98, 99, 100, 93, 0, 0, 0, 0, 245,
	0, 0, 257, 0, 0, 248, 0, 243, 0, 243,
	0, 88, 89, 90, 15, 96, 0, 101, 92, 0,
	0, 133, 0, 20, 240, 0, 202, 260, 0, 0,
	0, 261, 260, 0, 261, 0, 0, 0, 0, 0,
	87, 13, 102, 104, 109, 260, 107, 108, 143, 0,
	0, 171, 172, 173, 0, 202, 0, 188, 0, 204,
	205, 206, 207, 0, 239, 191, 192, 193, 189, 190,
	195, 94, 228, 0, 0, 141, 0, 0, 0, 0,
	0, 258, 79, 0, 82, 0, 84, 244, 0, 261,
	0, 0, 105, 110, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 158, 159, 160, 161, 162, 163, 164, 146,
	0, 0, 0, 0, 169, 182, 0, 0, 0, 157,
	0, 208, 0, 196, 0, 0, 0, 141, 134, 213,
	0, 241, 242, 203, 0, 246, 0, 0, 261, 254,
	249, 250, 251, 252, 253, 83, 85, 86, 141, 112,
	118, 0, 130, 132, 103, 111, 106, 144, 145, 148,
	149, 0, 166, 167, 0, 0, 0, 151, 0, 155,
	0, 174, 175, 176, 177, 178, 179, 180, 181, 147,
	168, 0, 238, 169, 183, 0, 0, 0, 0, 200,
	197, 0, 232, 0, 235, 232, 0, 230, 213, 221,
	0, 142, 0, 72, 0, 0, 259, 80, 0, 255,
	209, 0, 0, 121, 122, 0, 0, 0, 0, 0,
	135, 119, 0, 0, 0, 0, 150, 152, 0, 0,
	156, 170, 184, 0, 186, 187, 0, 198, 0, 0,
	16, 0, 234, 236, 17, 229, 0, 221, 19, 0,
	0, 0, 75, 58, 56, 24, 25, 54, 37, 54,
	54, 33, 34, 26, 27, 28, 29, 30, 38, 39,
	40, 41, 42, 43, 44, 52, 52, 52, 52, 52,
	261, 81, 211, 0, 113, 116, 123, 0, 125, 0,
	127, 128, 129, 114, 0, 0, 0, 120, 115, 131,
	165, 0, 153, 185, 194, 201, 0, 0, 0, 231,
	18, 222, 214, 215, 218, 73, 74, 0, 71, 21,
	57, 36, 0, 31, 32, 0, 45, 0, 46, 47,
	48, 49, 0, 50, 51, 78, 213, 0, 0, 0,
	124, 126, 0, 0, 0, 154, 199, 0, 237, 0,
	0, 0, 217, 219, 220, 0, 0, 60, 0, 63,
	64, 65, 66, 0, 68, 69, 23, 22, 0, 35,
	0, 0, 221, 212, 210, 117, 0, 0, 0, 0,
	223, 224, 216, 76, 59, 61, 62, 67, 70, 55,
	0, 225, 0, 139, 0, 0, 233, 0, 14, 0,
	0, 136, 0, 137, 138, 53, 226, 0, 140, 0,
	227,
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
		//line sql.y:182
		{
			SetParseTree(yylex, yyDollar[1].statement)
		}
	case 2:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:188
		{
			yyVAL.statement = yyDollar[1].selStmt
		}
	case 13:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:204
		{
			yyVAL.selStmt = &Select{Comments: Comments(yyDollar[2].bytes2), Distinct: yyDollar[3].str, SelectExprs: yyDollar[4].selectExprs}
		}
	case 14:
		yyDollar = yyS[yypt-12 : yypt+1]
		//line sql.y:208
		{
			yyVAL.selStmt = &Select{Comments: Comments(yyDollar[2].bytes2), Distinct: yyDollar[3].str, SelectExprs: yyDollar[4].selectExprs, From: yyDollar[6].tableExprs, Where: NewWhere(AST_WHERE, yyDollar[7].boolExpr), GroupBy: GroupBy(yyDollar[8].valExprs), Having: NewWhere(AST_HAVING, yyDollar[9].boolExpr), OrderBy: yyDollar[10].orderBy, Limit: yyDollar[11].limit, Lock: yyDollar[12].str}
		}
	case 15:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:212
		{
			yyVAL.selStmt = &Union{Type: yyDollar[2].str, Left: yyDollar[1].selStmt, Right: yyDollar[3].selStmt}
		}
	case 16:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line sql.y:218
		{
			yyVAL.statement = &Insert{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[4].tableName, Columns: yyDollar[5].columns, Rows: yyDollar[6].insRows, OnDup: OnDup(yyDollar[7].updateExprs)}
		}
	case 17:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line sql.y:222
		{
			cols := make(Columns, 0, len(yyDollar[6].updateExprs))
			vals := make(ValTuple, 0, len(yyDollar[6].updateExprs))
			for _, col := range yyDollar[6].updateExprs {
				cols = append(cols, &NonStarExpr{Expr: col.Name})
				vals = append(vals, col.Expr)
			}
			yyVAL.statement = &Insert{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[4].tableName, Columns: cols, Rows: Values{vals}, OnDup: OnDup(yyDollar[7].updateExprs)}
		}
	case 18:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line sql.y:234
		{
			yyVAL.statement = &Update{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[3].tableName, Exprs: yyDollar[5].updateExprs, Where: NewWhere(AST_WHERE, yyDollar[6].boolExpr), OrderBy: yyDollar[7].orderBy, Limit: yyDollar[8].limit}
		}
	case 19:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line sql.y:240
		{
			yyVAL.statement = &Delete{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[4].tableName, Where: NewWhere(AST_WHERE, yyDollar[5].boolExpr), OrderBy: yyDollar[6].orderBy, Limit: yyDollar[7].limit}
		}
	case 20:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:246
		{
			yyVAL.statement = &Set{Comments: Comments(yyDollar[2].bytes2), Exprs: yyDollar[3].updateExprs}
		}
	case 21:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:251
		{
			yyVAL.str = ""
		}
	case 22:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:255
		{
			yyVAL.str = AST_ZEROFILL
		}
	case 23:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:260
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
		//line sql.y:274
		{
			yyVAL.str = AST_DATE
		}
	case 27:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:278
		{
			yyVAL.str = AST_TIME
		}
	case 28:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:282
		{
			yyVAL.str = AST_TIMESTAMP
		}
	case 29:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:286
		{
			yyVAL.str = AST_DATETIME
		}
	case 30:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:290
		{
			yyVAL.str = AST_YEAR
		}
	case 31:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:296
		{
			if yyDollar[2].str == "" {
				yyVAL.str = AST_CHAR
			} else {
				yyVAL.str = AST_CHAR + yyDollar[2].str
			}
		}
	case 32:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:304
		{
			if yyDollar[2].str == "" {
				yyVAL.str = AST_VARCHAR
			} else {
				yyVAL.str = AST_VARCHAR + yyDollar[2].str
			}
		}
	case 33:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:312
		{
			yyVAL.str = AST_TEXT
		}
	case 34:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:316
		{
			yyVAL.str = AST_MEDIUMTEXT
		}
	case 35:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:320
		{
			yyVAL.str = AST_MEDIUMTEXT
			// do something with the charset id?
		}
	case 36:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:327
		{
			yyVAL.str = yyDollar[1].str + yyDollar[2].str
		}
	case 37:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:331
		{
			yyVAL.str = yyDollar[1].str
		}
	case 38:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:337
		{
			yyVAL.str = AST_BIT
		}
	case 39:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:341
		{
			yyVAL.str = AST_TINYINT
		}
	case 40:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:345
		{
			yyVAL.str = AST_SMALLINT
		}
	case 41:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:349
		{
			yyVAL.str = AST_MEDIUMINT
		}
	case 42:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:353
		{
			yyVAL.str = AST_INT
		}
	case 43:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:357
		{
			yyVAL.str = AST_INTEGER
		}
	case 44:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:361
		{
			yyVAL.str = AST_BIGINT
		}
	case 45:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:367
		{
			yyVAL.str = AST_REAL + yyDollar[2].str
		}
	case 46:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:371
		{
			yyVAL.str = AST_DOUBLE + yyDollar[2].str
		}
	case 47:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:375
		{
			yyVAL.str = AST_FLOAT + yyDollar[2].str
		}
	case 48:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:379
		{
			yyVAL.str = AST_DECIMAL + yyDollar[2].str
		}
	case 49:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:383
		{
			yyVAL.str = AST_DECIMAL + yyDollar[2].str
		}
	case 50:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:387
		{
			yyVAL.str = AST_NUMERIC + yyDollar[2].str
		}
	case 51:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:391
		{
			yyVAL.str = AST_NUMERIC + yyDollar[2].str
		}
	case 52:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:396
		{
			yyVAL.str = ""
		}
	case 53:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:400
		{
			yyVAL.str = "(" + string(yyDollar[2].bytes) + ", " + string(yyDollar[4].bytes) + ")"
		}
	case 54:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:405
		{
			yyVAL.str = ""
		}
	case 55:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:409
		{
			yyVAL.str = "(" + string(yyDollar[2].bytes) + ")"
		}
	case 56:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:414
		{
			yyVAL.str = ""
		}
	case 57:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:418
		{
			yyVAL.str = AST_UNSIGNED
		}
	case 58:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:423
		{
			yyVAL.columnAtts = ColumnAtts{}
		}
	case 59:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:427
		{
			yyVAL.columnAtts = append(yyVAL.columnAtts, AST_NOT_NULL)
		}
	case 61:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:433
		{
			node := StrVal(yyDollar[3].bytes)
			yyVAL.columnAtts = append(yyVAL.columnAtts, "default "+String(node))
		}
	case 62:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:438
		{
			node := NumVal(yyDollar[3].bytes)
			yyVAL.columnAtts = append(yyVAL.columnAtts, "default "+String(node))
		}
	case 63:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:443
		{
			yyVAL.columnAtts = append(yyVAL.columnAtts, AST_AUTO_INCREMENT)
		}
	case 64:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:447
		{
			yyVAL.columnAtts = append(yyVAL.columnAtts, yyDollar[2].str)
		}
	case 65:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:453
		{
			yyVAL.str = AST_PRIMARY_KEY
		}
	case 66:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:457
		{
			yyVAL.str = AST_UNIQUE_KEY
		}
	case 71:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:471
		{
			yyVAL.columnDefinition = &ColumnDefinition{ColName: string(yyDollar[1].bytes), ColType: yyDollar[2].str, ColumnAtts: yyDollar[3].columnAtts}
		}
	case 72:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:477
		{
			yyVAL.columnDefinitions = ColumnDefinitions{yyDollar[1].columnDefinition}
		}
	case 73:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:481
		{
			yyVAL.columnDefinitions = append(yyVAL.columnDefinitions, yyDollar[3].columnDefinition)
		}
	case 74:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line sql.y:487
		{
			yyVAL.statement = &CreateTable{Name: yyDollar[4].bytes, ColumnDefinitions: yyDollar[6].columnDefinitions}
		}
	case 77:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:497
		{
			yyVAL.statement = yyDollar[1].statement
		}
	case 78:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line sql.y:501
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[7].bytes, NewName: yyDollar[7].bytes}
		}
	case 79:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:506
		{
			yyVAL.statement = &DDL{Action: AST_CREATE, NewName: yyDollar[3].bytes}
		}
	case 80:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line sql.y:512
		{
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[4].bytes, NewName: yyDollar[4].bytes}
		}
	case 81:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line sql.y:516
		{
			// Change this to a rename statement
			yyVAL.statement = &DDL{Action: AST_RENAME, Table: yyDollar[4].bytes, NewName: yyDollar[7].bytes}
		}
	case 82:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:521
		{
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[3].bytes, NewName: yyDollar[3].bytes}
		}
	case 83:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:527
		{
			yyVAL.statement = &DDL{Action: AST_RENAME, Table: yyDollar[3].bytes, NewName: yyDollar[5].bytes}
		}
	case 84:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:533
		{
			yyVAL.statement = &DDL{Action: AST_DROP, Table: yyDollar[4].bytes}
		}
	case 85:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:537
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[5].bytes, NewName: yyDollar[5].bytes}
		}
	case 86:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:542
		{
			yyVAL.statement = &DDL{Action: AST_DROP, Table: yyDollar[4].bytes}
		}
	case 87:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:548
		{
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[3].bytes, NewName: yyDollar[3].bytes}
		}
	case 88:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:554
		{
			yyVAL.statement = &Other{}
		}
	case 89:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:558
		{
			yyVAL.statement = &Other{}
		}
	case 90:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:562
		{
			yyVAL.statement = &Other{}
		}
	case 91:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:567
		{
			SetAllowComments(yylex, true)
		}
	case 92:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:571
		{
			yyVAL.bytes2 = yyDollar[2].bytes2
			SetAllowComments(yylex, false)
		}
	case 93:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:577
		{
			yyVAL.bytes2 = nil
		}
	case 94:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:581
		{
			yyVAL.bytes2 = append(yyDollar[1].bytes2, yyDollar[2].bytes)
		}
	case 95:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:587
		{
			yyVAL.str = AST_UNION
		}
	case 96:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:591
		{
			yyVAL.str = AST_UNION_ALL
		}
	case 97:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:595
		{
			yyVAL.str = AST_SET_MINUS
		}
	case 98:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:599
		{
			yyVAL.str = AST_EXCEPT
		}
	case 99:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:603
		{
			yyVAL.str = AST_INTERSECT
		}
	case 100:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:608
		{
			yyVAL.str = ""
		}
	case 101:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:612
		{
			yyVAL.str = AST_DISTINCT
		}
	case 102:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:618
		{
			yyVAL.selectExprs = SelectExprs{yyDollar[1].selectExpr}
		}
	case 103:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:622
		{
			yyVAL.selectExprs = append(yyVAL.selectExprs, yyDollar[3].selectExpr)
		}
	case 104:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:628
		{
			yyVAL.selectExpr = &StarExpr{}
		}
	case 105:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:632
		{
			yyVAL.selectExpr = &NonStarExpr{Expr: yyDollar[1].expr, As: yyDollar[2].bytes}
		}
	case 106:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:636
		{
			yyVAL.selectExpr = &StarExpr{TableName: yyDollar[1].bytes}
		}
	case 107:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:642
		{
			yyVAL.expr = yyDollar[1].boolExpr
		}
	case 108:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:646
		{
			yyVAL.expr = yyDollar[1].valExpr
		}
	case 109:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:651
		{
			yyVAL.bytes = nil
		}
	case 110:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:655
		{
			yyVAL.bytes = yyDollar[1].bytes
		}
	case 111:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:659
		{
			yyVAL.bytes = yyDollar[2].bytes
		}
	case 112:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:665
		{
			yyVAL.tableExprs = TableExprs{yyDollar[1].tableExpr}
		}
	case 113:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:669
		{
			yyVAL.tableExprs = append(yyVAL.tableExprs, yyDollar[3].tableExpr)
		}
	case 114:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:675
		{
			yyVAL.tableExpr = &AliasedTableExpr{Expr: yyDollar[1].smTableExpr, As: yyDollar[2].bytes, Hints: yyDollar[3].indexHints}
		}
	case 115:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:679
		{
			yyVAL.tableExpr = &ParenTableExpr{Expr: yyDollar[2].tableExpr}
		}
	case 116:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:683
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr}
		}
	case 117:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:687
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr, On: yyDollar[5].boolExpr}
		}
	case 118:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:692
		{
			yyVAL.bytes = nil
		}
	case 119:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:696
		{
			yyVAL.bytes = yyDollar[1].bytes
		}
	case 120:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:700
		{
			yyVAL.bytes = yyDollar[2].bytes
		}
	case 121:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:706
		{
			yyVAL.str = AST_JOIN
		}
	case 122:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:710
		{
			yyVAL.str = AST_STRAIGHT_JOIN
		}
	case 123:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:714
		{
			yyVAL.str = AST_LEFT_JOIN
		}
	case 124:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:718
		{
			yyVAL.str = AST_LEFT_JOIN
		}
	case 125:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:722
		{
			yyVAL.str = AST_RIGHT_JOIN
		}
	case 126:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:726
		{
			yyVAL.str = AST_RIGHT_JOIN
		}
	case 127:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:730
		{
			yyVAL.str = AST_JOIN
		}
	case 128:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:734
		{
			yyVAL.str = AST_CROSS_JOIN
		}
	case 129:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:738
		{
			yyVAL.str = AST_NATURAL_JOIN
		}
	case 130:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:744
		{
			yyVAL.smTableExpr = &TableName{Name: yyDollar[1].bytes}
		}
	case 131:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:748
		{
			yyVAL.smTableExpr = &TableName{Qualifier: yyDollar[1].bytes, Name: yyDollar[3].bytes}
		}
	case 132:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:752
		{
			yyVAL.smTableExpr = yyDollar[1].subquery
		}
	case 133:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:758
		{
			yyVAL.tableName = &TableName{Name: yyDollar[1].bytes}
		}
	case 134:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:762
		{
			yyVAL.tableName = &TableName{Qualifier: yyDollar[1].bytes, Name: yyDollar[3].bytes}
		}
	case 135:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:767
		{
			yyVAL.indexHints = nil
		}
	case 136:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:771
		{
			yyVAL.indexHints = &IndexHints{Type: AST_USE, Indexes: yyDollar[4].bytes2}
		}
	case 137:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:775
		{
			yyVAL.indexHints = &IndexHints{Type: AST_IGNORE, Indexes: yyDollar[4].bytes2}
		}
	case 138:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:779
		{
			yyVAL.indexHints = &IndexHints{Type: AST_FORCE, Indexes: yyDollar[4].bytes2}
		}
	case 139:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:785
		{
			yyVAL.bytes2 = [][]byte{yyDollar[1].bytes}
		}
	case 140:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:789
		{
			yyVAL.bytes2 = append(yyDollar[1].bytes2, yyDollar[3].bytes)
		}
	case 141:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:794
		{
			yyVAL.boolExpr = nil
		}
	case 142:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:798
		{
			yyVAL.boolExpr = yyDollar[2].boolExpr
		}
	case 144:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:805
		{
			yyVAL.boolExpr = &AndExpr{Left: yyDollar[1].boolExpr, Right: yyDollar[3].boolExpr}
		}
	case 145:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:809
		{
			yyVAL.boolExpr = &OrExpr{Left: yyDollar[1].boolExpr, Right: yyDollar[3].boolExpr}
		}
	case 146:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:813
		{
			yyVAL.boolExpr = &NotExpr{Expr: yyDollar[2].boolExpr}
		}
	case 147:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:817
		{
			yyVAL.boolExpr = &ParenBoolExpr{Expr: yyDollar[2].boolExpr}
		}
	case 148:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:823
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: yyDollar[2].str, Right: yyDollar[3].valExpr}
		}
	case 149:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:827
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_IN, Right: yyDollar[3].colTuple}
		}
	case 150:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:831
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_NOT_IN, Right: yyDollar[4].colTuple}
		}
	case 151:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:835
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_LIKE, Right: yyDollar[3].valExpr}
		}
	case 152:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:839
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_NOT_LIKE, Right: yyDollar[4].valExpr}
		}
	case 153:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:843
		{
			yyVAL.boolExpr = &RangeCond{Left: yyDollar[1].valExpr, Operator: AST_BETWEEN, From: yyDollar[3].valExpr, To: yyDollar[5].valExpr}
		}
	case 154:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line sql.y:847
		{
			yyVAL.boolExpr = &RangeCond{Left: yyDollar[1].valExpr, Operator: AST_NOT_BETWEEN, From: yyDollar[4].valExpr, To: yyDollar[6].valExpr}
		}
	case 155:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:851
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_IS, Right: yyDollar[3].valExpr}
		}
	case 156:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:855
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_IS_NOT, Right: yyDollar[4].valExpr}
		}
	case 157:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:859
		{
			yyVAL.boolExpr = &ExistsExpr{Subquery: yyDollar[2].subquery}
		}
	case 158:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:865
		{
			yyVAL.str = AST_EQ
		}
	case 159:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:869
		{
			yyVAL.str = AST_LT
		}
	case 160:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:873
		{
			yyVAL.str = AST_GT
		}
	case 161:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:877
		{
			yyVAL.str = AST_LE
		}
	case 162:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:881
		{
			yyVAL.str = AST_GE
		}
	case 163:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:885
		{
			yyVAL.str = AST_NE
		}
	case 164:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:889
		{
			yyVAL.str = AST_NSE
		}
	case 165:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:895
		{
			yyVAL.colTuple = ValTuple(yyDollar[2].valExprs)
		}
	case 166:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:899
		{
			yyVAL.colTuple = yyDollar[1].subquery
		}
	case 167:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:903
		{
			yyVAL.colTuple = ListArg(yyDollar[1].bytes)
		}
	case 168:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:909
		{
			yyVAL.subquery = &Subquery{yyDollar[2].selStmt}
		}
	case 169:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:915
		{
			yyVAL.valExprs = ValExprs{yyDollar[1].valExpr}
		}
	case 170:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:919
		{
			yyVAL.valExprs = append(yyDollar[1].valExprs, yyDollar[3].valExpr)
		}
	case 171:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:925
		{
			yyVAL.valExpr = yyDollar[1].valExpr
		}
	case 172:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:929
		{
			yyVAL.valExpr = yyDollar[1].colName
		}
	case 173:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:933
		{
			yyVAL.valExpr = yyDollar[1].rowTuple
		}
	case 174:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:937
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_BITAND, Right: yyDollar[3].valExpr}
		}
	case 175:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:941
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_BITOR, Right: yyDollar[3].valExpr}
		}
	case 176:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:945
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_BITXOR, Right: yyDollar[3].valExpr}
		}
	case 177:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:949
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_PLUS, Right: yyDollar[3].valExpr}
		}
	case 178:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:953
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_MINUS, Right: yyDollar[3].valExpr}
		}
	case 179:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:957
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_MULT, Right: yyDollar[3].valExpr}
		}
	case 180:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:961
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_DIV, Right: yyDollar[3].valExpr}
		}
	case 181:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:965
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_MOD, Right: yyDollar[3].valExpr}
		}
	case 182:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:987
		{
			if num, ok := yyDollar[2].valExpr.(NumVal); ok {
				switch yyDollar[1].byt {
				case '-':
					yyVAL.valExpr = append(NumVal("-"), num...)
				case '+':
					yyVAL.valExpr = num
				default:
					yyVAL.valExpr = &UnaryExpr{Operator: yyDollar[1].byt, Expr: yyDollar[2].valExpr}
				}
			} else {
				yyVAL.valExpr = &UnaryExpr{Operator: yyDollar[1].byt, Expr: yyDollar[2].valExpr}
			}
		}
	case 183:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1002
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].bytes}
		}
	case 184:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:1006
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].bytes, Exprs: yyDollar[3].selectExprs}
		}
	case 185:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:1010
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].bytes, Distinct: true, Exprs: yyDollar[4].selectExprs}
		}
	case 186:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:1014
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].bytes, Exprs: yyDollar[3].selectExprs}
		}
	case 187:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:1018
		{
			// XXX(toshok) select_statement is dropped here
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].bytes}
		}
	case 188:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1023
		{
			yyVAL.valExpr = yyDollar[1].caseExpr
		}
	case 189:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1029
		{
			yyVAL.bytes = IF_BYTES
		}
	case 190:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1033
		{
			yyVAL.bytes = VALUES_BYTES
		}
	case 191:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1039
		{
			yyVAL.byt = AST_UPLUS
		}
	case 192:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1043
		{
			yyVAL.byt = AST_UMINUS
		}
	case 193:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1047
		{
			yyVAL.byt = AST_TILDA
		}
	case 194:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:1053
		{
			yyVAL.caseExpr = &CaseExpr{Expr: yyDollar[2].valExpr, Whens: yyDollar[3].whens, Else: yyDollar[4].valExpr}
		}
	case 195:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1058
		{
			yyVAL.valExpr = nil
		}
	case 196:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1062
		{
			yyVAL.valExpr = yyDollar[1].valExpr
		}
	case 197:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1068
		{
			yyVAL.whens = []*When{yyDollar[1].when}
		}
	case 198:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1072
		{
			yyVAL.whens = append(yyDollar[1].whens, yyDollar[2].when)
		}
	case 199:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:1078
		{
			yyVAL.when = &When{Cond: yyDollar[2].boolExpr, Val: yyDollar[4].valExpr}
		}
	case 200:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1083
		{
			yyVAL.valExpr = nil
		}
	case 201:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1087
		{
			yyVAL.valExpr = yyDollar[2].valExpr
		}
	case 202:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1093
		{
			yyVAL.colName = &ColName{Name: yyDollar[1].bytes}
		}
	case 203:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1097
		{
			yyVAL.colName = &ColName{Qualifier: yyDollar[1].bytes, Name: yyDollar[3].bytes}
		}
	case 204:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1103
		{
			yyVAL.valExpr = StrVal(yyDollar[1].bytes)
		}
	case 205:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1107
		{
			yyVAL.valExpr = NumVal(yyDollar[1].bytes)
		}
	case 206:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1111
		{
			yyVAL.valExpr = ValArg(yyDollar[1].bytes)
		}
	case 207:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1115
		{
			yyVAL.valExpr = &NullVal{}
		}
	case 208:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1119
		{
			yyVAL.valExpr = BinaryVal(yyDollar[2].bytes)
		}
	case 209:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1124
		{
			yyVAL.valExprs = nil
		}
	case 210:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1128
		{
			yyVAL.valExprs = yyDollar[3].valExprs
		}
	case 211:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1133
		{
			yyVAL.boolExpr = nil
		}
	case 212:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1137
		{
			yyVAL.boolExpr = yyDollar[2].boolExpr
		}
	case 213:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1142
		{
			yyVAL.orderBy = nil
		}
	case 214:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1146
		{
			yyVAL.orderBy = yyDollar[3].orderBy
		}
	case 215:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1152
		{
			yyVAL.orderBy = OrderBy{yyDollar[1].order}
		}
	case 216:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1156
		{
			yyVAL.orderBy = append(yyDollar[1].orderBy, yyDollar[3].order)
		}
	case 217:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1162
		{
			yyVAL.order = &Order{Expr: yyDollar[1].valExpr, Direction: yyDollar[2].str}
		}
	case 218:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1167
		{
			yyVAL.str = AST_ASC
		}
	case 219:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1171
		{
			yyVAL.str = AST_ASC
		}
	case 220:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1175
		{
			yyVAL.str = AST_DESC
		}
	case 221:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1180
		{
			yyVAL.limit = nil
		}
	case 222:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1184
		{
			yyVAL.limit = &Limit{Rowcount: yyDollar[2].valExpr}
		}
	case 223:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:1188
		{
			yyVAL.limit = &Limit{Offset: yyDollar[2].valExpr, Rowcount: yyDollar[4].valExpr}
		}
	case 224:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:1192
		{
			yyVAL.limit = &Limit{Offset: yyDollar[4].valExpr, Rowcount: yyDollar[2].valExpr}
		}
	case 225:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1197
		{
			yyVAL.str = ""
		}
	case 226:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1201
		{
			yyVAL.str = AST_FOR_UPDATE
		}
	case 227:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:1205
		{
			if !bytes.Equal(yyDollar[3].bytes, SHARE) {
				yylex.Error("expecting share")
				return 1
			}
			if !bytes.Equal(yyDollar[4].bytes, MODE) {
				yylex.Error("expecting mode")
				return 1
			}
			yyVAL.str = AST_SHARE_MODE
		}
	case 228:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1218
		{
			yyVAL.columns = nil
		}
	case 229:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1222
		{
			yyVAL.columns = yyDollar[2].columns
		}
	case 230:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1228
		{
			yyVAL.columns = Columns{&NonStarExpr{Expr: yyDollar[1].colName}}
		}
	case 231:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1232
		{
			yyVAL.columns = append(yyVAL.columns, &NonStarExpr{Expr: yyDollar[3].colName})
		}
	case 232:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1237
		{
			yyVAL.updateExprs = nil
		}
	case 233:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:1241
		{
			yyVAL.updateExprs = yyDollar[5].updateExprs
		}
	case 234:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1247
		{
			yyVAL.insRows = yyDollar[2].values
		}
	case 235:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1251
		{
			yyVAL.insRows = yyDollar[1].selStmt
		}
	case 236:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1257
		{
			yyVAL.values = Values{yyDollar[1].rowTuple}
		}
	case 237:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1261
		{
			yyVAL.values = append(yyDollar[1].values, yyDollar[3].rowTuple)
		}
	case 238:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1267
		{
			yyVAL.rowTuple = ValTuple(yyDollar[2].valExprs)
		}
	case 239:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1271
		{
			yyVAL.rowTuple = yyDollar[1].subquery
		}
	case 240:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1277
		{
			yyVAL.updateExprs = UpdateExprs{yyDollar[1].updateExpr}
		}
	case 241:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1281
		{
			yyVAL.updateExprs = append(yyDollar[1].updateExprs, yyDollar[3].updateExpr)
		}
	case 242:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1287
		{
			yyVAL.updateExpr = &UpdateExpr{Name: yyDollar[1].colName, Expr: yyDollar[3].valExpr}
		}
	case 243:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1292
		{
			yyVAL.empty = struct{}{}
		}
	case 244:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1294
		{
			yyVAL.empty = struct{}{}
		}
	case 245:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1297
		{
			yyVAL.empty = struct{}{}
		}
	case 246:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1299
		{
			yyVAL.empty = struct{}{}
		}
	case 247:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1302
		{
			yyVAL.empty = struct{}{}
		}
	case 248:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1304
		{
			yyVAL.empty = struct{}{}
		}
	case 249:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1308
		{
			yyVAL.empty = struct{}{}
		}
	case 250:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1310
		{
			yyVAL.empty = struct{}{}
		}
	case 251:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1312
		{
			yyVAL.empty = struct{}{}
		}
	case 252:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1314
		{
			yyVAL.empty = struct{}{}
		}
	case 253:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1316
		{
			yyVAL.empty = struct{}{}
		}
	case 254:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1319
		{
			yyVAL.empty = struct{}{}
		}
	case 255:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1321
		{
			yyVAL.empty = struct{}{}
		}
	case 256:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1324
		{
			yyVAL.empty = struct{}{}
		}
	case 257:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1326
		{
			yyVAL.empty = struct{}{}
		}
	case 258:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1329
		{
			yyVAL.empty = struct{}{}
		}
	case 259:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1331
		{
			yyVAL.empty = struct{}{}
		}
	case 260:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1335
		{
			yyVAL.bytes = bytes.ToLower(yyDollar[1].bytes)
		}
	case 261:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1340
		{
			ForceEOF(yylex)
		}
	}
	goto yystack /* stack new state and value */
}
