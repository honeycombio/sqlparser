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
const TABLE = 57420
const INDEX = 57421
const VIEW = 57422
const TO = 57423
const IGNORE = 57424
const IF = 57425
const USING = 57426
const SHOW = 57427
const DESCRIBE = 57428
const EXPLAIN = 57429
const BIT = 57430
const TINYINT = 57431
const SMALLINT = 57432
const MEDIUMINT = 57433
const INT = 57434
const INTEGER = 57435
const BIGINT = 57436
const REAL = 57437
const DOUBLE = 57438
const FLOAT = 57439
const UNSIGNED = 57440
const ZEROFILL = 57441
const DECIMAL = 57442
const NUMERIC = 57443
const DATE = 57444
const TIME = 57445
const TIMESTAMP = 57446
const DATETIME = 57447
const YEAR = 57448
const TEXT = 57449
const CHAR = 57450
const VARCHAR = 57451
const NULLX = 57452
const AUTO_INCREMENT = 57453
const BOOL = 57454
const APPROXNUM = 57455
const INTNUM = 57456

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

const yyNprod = 256
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 645

var yyAct = [...]int{

	95, 296, 160, 433, 92, 360, 248, 93, 63, 163,
	370, 252, 366, 199, 239, 288, 210, 443, 86, 162,
	3, 81, 443, 179, 137, 136, 82, 51, 446, 262,
	263, 264, 265, 266, 430, 267, 268, 410, 443, 66,
	431, 365, 71, 65, 131, 74, 64, 104, 54, 78,
	29, 30, 31, 32, 52, 53, 231, 299, 294, 87,
	400, 131, 341, 343, 187, 131, 407, 401, 231, 77,
	69, 121, 44, 240, 45, 258, 47, 48, 49, 351,
	129, 125, 388, 42, 229, 133, 406, 408, 387, 445,
	273, 386, 342, 164, 444, 70, 135, 165, 73, 122,
	345, 50, 124, 240, 46, 286, 399, 118, 159, 161,
	442, 120, 430, 173, 66, 219, 350, 66, 65, 183,
	182, 65, 177, 39, 114, 41, 230, 72, 347, 300,
	293, 137, 136, 283, 87, 205, 183, 281, 136, 383,
	232, 209, 289, 254, 217, 218, 169, 221, 222, 223,
	224, 225, 226, 227, 228, 207, 208, 197, 204, 193,
	402, 220, 116, 181, 149, 150, 151, 394, 206, 233,
	87, 87, 128, 289, 420, 385, 66, 66, 203, 191,
	65, 246, 384, 194, 244, 137, 136, 212, 255, 235,
	237, 339, 335, 333, 243, 397, 398, 336, 334, 250,
	353, 247, 147, 148, 149, 150, 151, 60, 393, 29,
	30, 31, 32, 338, 233, 272, 256, 337, 276, 277,
	130, 274, 259, 144, 145, 146, 147, 148, 149, 150,
	151, 275, 180, 116, 280, 190, 192, 189, 231, 87,
	180, 144, 145, 146, 147, 148, 149, 150, 151, 203,
	291, 431, 395, 285, 295, 355, 117, 419, 282, 287,
	292, 202, 212, 76, 418, 175, 236, 112, 98, 131,
	115, 201, 14, 102, 331, 332, 109, 438, 176, 166,
	260, 349, 376, 103, 85, 99, 100, 101, 116, 352,
	14, 371, 367, 102, 90, 66, 109, 357, 107, 356,
	358, 361, 184, 103, 67, 99, 100, 101, 203, 203,
	213, 362, 170, 79, 166, 168, 211, 167, 107, 89,
	368, 369, 202, 105, 106, 83, 111, 426, 427, 413,
	110, 412, 201, 411, 372, 373, 374, 377, 375, 378,
	171, 271, 134, 105, 106, 108, 72, 67, 253, 346,
	110, 389, 344, 328, 327, 379, 390, 102, 270, 72,
	109, 196, 195, 392, 178, 108, 126, 103, 67, 99,
	100, 101, 123, 119, 61, 80, 75, 234, 166, 113,
	440, 354, 107, 429, 428, 233, 391, 414, 14, 59,
	448, 214, 416, 215, 216, 185, 422, 423, 361, 441,
	415, 424, 417, 14, 15, 16, 17, 105, 106, 127,
	425, 279, 242, 57, 110, 55, 432, 297, 382, 434,
	434, 434, 66, 435, 436, 298, 65, 249, 381, 108,
	437, 348, 18, 144, 145, 146, 147, 148, 149, 150,
	151, 330, 180, 62, 449, 447, 421, 14, 34, 450,
	405, 451, 315, 316, 317, 318, 319, 320, 321, 322,
	323, 324, 404, 363, 325, 326, 310, 311, 312, 313,
	314, 309, 307, 308, 304, 306, 98, 305, 403, 409,
	14, 102, 364, 302, 109, 303, 20, 21, 23, 22,
	24, 103, 85, 99, 100, 101, 98, 19, 25, 26,
	27, 102, 90, 251, 109, 301, 107, 186, 40, 257,
	188, 103, 67, 99, 100, 101, 43, 262, 263, 264,
	265, 266, 90, 267, 268, 68, 107, 89, 245, 174,
	439, 105, 106, 83, 396, 359, 380, 278, 110, 144,
	145, 146, 147, 148, 149, 150, 151, 89, 329, 98,
	284, 105, 106, 108, 102, 172, 238, 109, 110, 97,
	94, 96, 290, 91, 103, 67, 99, 100, 101, 241,
	138, 88, 340, 108, 200, 90, 261, 198, 84, 107,
	269, 132, 56, 28, 58, 13, 139, 143, 141, 142,
	144, 145, 146, 147, 148, 149, 150, 151, 12, 33,
	89, 11, 10, 9, 105, 106, 8, 155, 156, 157,
	158, 110, 152, 153, 154, 35, 36, 37, 38, 7,
	6, 5, 4, 2, 1, 0, 108, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 140, 144, 145, 146,
	147, 148, 149, 150, 151,
}
var yyPact = [...]int{

	398, -1000, -1000, 155, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	30, -23, 11, -17, 8, -1000, -1000, -1000, 442, 397,
	-1000, -1000, -1000, 394, -1000, 359, 337, 434, 310, -28,
	1, 309, -1000, 5, 309, -1000, 339, -29, 309, -29,
	338, -1000, -1000, -1000, -1000, -1000, 455, -1000, 284, 337,
	345, 43, 337, 175, -1000, 208, -1000, 26, 336, 39,
	309, -1000, -1000, 335, -1000, -15, 329, 388, 103, 309,
	-1000, 211, -1000, -1000, 322, 15, 61, 564, -1000, 528,
	475, -1000, -1000, -1000, 331, 270, 268, -1000, 265, -1000,
	-1000, -1000, -1000, 302, -1000, -1000, -1000, -1000, -1000, -1000,
	331, -1000, 231, 310, 327, 432, 310, 331, 309, 255,
	374, -35, -1000, 146, -1000, 325, -1000, -1000, 324, -1000,
	224, 455, -1000, -1000, 309, 90, 528, 528, 331, 269,
	369, 331, 331, 89, 331, 331, 331, 331, 331, 331,
	331, 331, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	564, -46, -4, 10, 564, -1000, 267, 247, 455, -1000,
	442, -1000, -11, 517, 383, 310, 310, 230, -1000, 414,
	528, -1000, 517, -1000, 311, -1000, 74, 309, -1000, -21,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 222, 458,
	321, 285, 9, -1000, -1000, -1000, -1000, -1000, 67, 517,
	-1000, 267, -1000, -1000, 269, 331, 331, 517, 466, -1000,
	385, 126, 126, 126, 86, 86, -1000, -1000, -1000, -1000,
	-1000, 331, -1000, 517, -1000, 7, 455, 3, 19, -1000,
	528, 73, 232, 155, 104, 0, -1000, 414, 402, 411,
	61, -1, -1000, 349, 317, -1000, -1000, 316, -1000, 430,
	224, 224, -1000, -1000, 134, 133, 158, 154, 132, -5,
	-1000, 315, -30, 312, -2, -1000, 517, 360, 331, -1000,
	517, -1000, -14, -1000, -8, -1000, 331, 115, -1000, 350,
	197, -1000, -1000, -1000, 310, 402, -1000, 331, 331, 311,
	-1000, -1000, -72, -1000, -1000, 245, -1000, 245, 245, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, 244, 244, 244, 235, 235, -1000, -1000, 416,
	404, 458, 70, -1000, 123, -1000, 116, -1000, -1000, -1000,
	-1000, -3, -6, -12, -1000, -1000, -1000, -1000, 331, 517,
	-1000, -1000, 517, 331, 354, 232, -1000, -1000, 150, 194,
	-1000, 168, -1000, 34, -77, -1000, -1000, 294, -1000, -1000,
	-1000, 292, -1000, -1000, -1000, -1000, 290, -1000, -1000, -1000,
	414, 528, 331, 528, -1000, -1000, 217, 210, 127, 517,
	517, 439, -1000, 331, 331, 331, -1000, -1000, -1000, 384,
	-1000, 289, -1000, -1000, -1000, -1000, 352, -1000, 351, -1000,
	-1000, -96, 193, -18, 402, 61, 180, 61, 309, 309,
	309, 310, 517, 517, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 238, 364, -20, -1000, -36, -41, 175, -102, -1000,
	438, 368, -1000, 309, -1000, -1000, -1000, -1000, 309, -1000,
	309, -1000,
}
var yyPgo = [...]int{

	0, 624, 623, 19, 622, 621, 620, 619, 606, 603,
	602, 601, 598, 585, 599, 584, 583, 582, 21, 26,
	581, 580, 578, 577, 13, 576, 574, 207, 572, 3,
	23, 18, 571, 570, 569, 563, 2, 16, 9, 562,
	7, 561, 47, 560, 4, 559, 556, 14, 555, 550,
	548, 536, 6, 535, 5, 534, 1, 530, 529, 528,
	15, 8, 46, 263, 525, 516, 510, 509, 508, 507,
	0, 27, 505, 11, 503, 497, 12, 485, 483, 482,
	479, 478, 477, 475, 10, 474, 463, 462, 450, 448,
}
var yyR1 = [...]int{

	0, 1, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 3, 3, 4, 4, 5, 6, 7,
	80, 80, 72, 72, 72, 85, 85, 85, 85, 85,
	77, 77, 77, 78, 78, 82, 82, 82, 82, 82,
	82, 82, 83, 83, 83, 83, 83, 83, 83, 84,
	84, 76, 76, 79, 79, 86, 86, 86, 86, 86,
	86, 86, 81, 81, 87, 87, 88, 88, 73, 74,
	74, 75, 8, 8, 8, 9, 9, 9, 10, 11,
	11, 11, 12, 13, 13, 13, 89, 14, 15, 15,
	16, 16, 16, 16, 16, 17, 17, 18, 18, 19,
	19, 19, 22, 22, 20, 20, 20, 23, 23, 24,
	24, 24, 24, 21, 21, 21, 25, 25, 25, 25,
	25, 25, 25, 25, 25, 26, 26, 26, 27, 27,
	28, 28, 28, 28, 29, 29, 30, 30, 31, 31,
	31, 31, 31, 32, 32, 32, 32, 32, 32, 32,
	32, 32, 32, 33, 33, 33, 33, 33, 33, 33,
	37, 37, 37, 42, 38, 38, 36, 36, 36, 36,
	36, 36, 36, 36, 36, 36, 36, 36, 36, 36,
	36, 36, 36, 41, 41, 43, 43, 43, 45, 48,
	48, 46, 46, 47, 49, 49, 44, 44, 35, 35,
	35, 35, 35, 50, 50, 51, 51, 52, 52, 53,
	53, 54, 55, 55, 55, 56, 56, 56, 56, 57,
	57, 57, 58, 58, 59, 59, 60, 60, 34, 34,
	39, 39, 40, 40, 61, 61, 62, 63, 63, 64,
	64, 65, 65, 66, 66, 66, 66, 66, 67, 67,
	68, 68, 69, 69, 70, 71,
}
var yyR2 = [...]int{

	0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 12, 3, 7, 7, 8, 7, 3,
	0, 1, 3, 1, 1, 1, 1, 1, 1, 1,
	2, 2, 1, 2, 1, 1, 1, 1, 1, 1,
	1, 1, 2, 2, 2, 2, 2, 2, 2, 0,
	5, 0, 3, 0, 1, 0, 3, 2, 3, 3,
	2, 2, 1, 1, 2, 1, 1, 2, 3, 1,
	3, 7, 1, 8, 4, 6, 7, 4, 5, 4,
	5, 5, 3, 2, 2, 2, 0, 2, 0, 2,
	1, 2, 1, 1, 1, 0, 1, 1, 3, 1,
	2, 3, 1, 1, 0, 1, 2, 1, 3, 3,
	3, 3, 5, 0, 1, 2, 1, 1, 2, 3,
	2, 3, 2, 2, 2, 1, 3, 1, 1, 3,
	0, 5, 5, 5, 1, 3, 0, 2, 1, 3,
	3, 2, 3, 3, 3, 4, 3, 4, 5, 6,
	3, 4, 2, 1, 1, 1, 1, 1, 1, 1,
	3, 1, 1, 3, 1, 3, 1, 1, 1, 3,
	3, 3, 3, 3, 3, 3, 3, 2, 3, 4,
	5, 4, 1, 1, 1, 1, 1, 1, 5, 0,
	1, 1, 2, 4, 0, 2, 1, 3, 1, 1,
	1, 1, 2, 0, 3, 0, 2, 0, 3, 1,
	3, 2, 0, 1, 1, 0, 2, 4, 4, 0,
	2, 4, 0, 3, 1, 3, 0, 5, 2, 1,
	1, 3, 3, 1, 1, 3, 3, 0, 2, 0,
	3, 0, 1, 1, 1, 1, 1, 1, 0, 1,
	0, 1, 0, 2, 1, 0,
}
var yyChk = [...]int{

	-1000, -1, -2, -3, -4, -5, -6, -7, -8, -9,
	-10, -11, -12, -13, 5, 6, 7, 8, 34, -75,
	88, 89, 91, 90, 92, 100, 101, 102, -16, 54,
	55, 56, 57, -14, -89, -14, -14, -14, -14, 93,
	-68, 95, 53, -65, 95, 97, 93, 93, 94, 95,
	93, -71, -71, -71, -3, 18, -17, 19, -15, 30,
	-27, 37, 9, -61, -62, -44, -70, 37, -64, 98,
	94, -70, 37, 93, -70, 37, -63, 98, -70, -63,
	37, -18, -19, 78, -22, 37, -31, -36, -32, 72,
	47, -35, -44, -40, -43, -70, -41, -45, 21, 38,
	39, 40, 26, 36, -42, 76, 77, 51, 98, 29,
	83, 42, -27, 34, 81, -27, 58, 48, 81, 37,
	72, -70, -71, 37, -71, 96, 37, 21, 69, -70,
	9, 58, -20, -70, 20, 81, 71, 70, -33, 22,
	72, 24, 25, 23, 73, 74, 75, 76, 77, 78,
	79, 80, 48, 49, 50, 43, 44, 45, 46, -31,
	-36, -31, -3, -38, -36, -36, 47, 47, 47, -42,
	47, 38, -48, -36, -58, 34, 47, -61, 37, -30,
	10, -62, -36, -70, 47, 21, -69, 99, -66, 91,
	89, 33, 90, 13, 37, 37, 37, -71, -23, -24,
	-26, 47, 37, -42, -19, -70, 78, -31, -31, -36,
	-37, 47, -42, 41, 22, 24, 25, -36, -36, 26,
	72, -36, -36, -36, -36, -36, -36, -36, -36, 130,
	130, 58, 130, -36, 130, -18, 19, -18, -46, -47,
	84, -34, 29, -3, -61, -59, -44, -30, -52, 13,
	-31, -74, -73, 37, 69, -70, -71, -67, 96, -30,
	58, -25, 59, 60, 61, 62, 63, 65, 66, -21,
	37, 20, -24, 81, -38, -37, -36, -36, 71, 26,
	-36, 130, -18, 130, -49, -47, 86, -31, -60, 69,
	-39, -40, -60, 130, 58, -52, -56, 15, 14, 58,
	130, -72, -78, -77, -85, -82, -83, 123, 124, 122,
	117, 118, 119, 120, 121, 103, 104, 105, 106, 107,
	108, 109, 110, 111, 112, 115, 116, 37, 37, -50,
	11, -24, -24, 59, 64, 59, 64, 59, 59, 59,
	-28, 67, 97, 68, 37, 130, 37, 130, 71, -36,
	130, 87, -36, 85, 31, 58, -44, -56, -36, -53,
	-54, -36, -73, -86, -79, 113, -76, 47, -76, -76,
	-84, 47, -84, -84, -84, -76, 47, -84, -76, -71,
	-51, 12, 14, 69, 59, 59, 94, 94, 94, -36,
	-36, 32, -40, 58, 17, 58, -55, 27, 28, 72,
	26, 33, 126, -81, -87, -88, 52, 32, 53, -80,
	114, 39, 39, 39, -52, -31, -38, -31, 47, 47,
	47, 7, -36, -36, -54, 26, 38, 39, 32, 32,
	130, 58, -56, -29, -70, -29, -29, -61, 39, -57,
	16, 35, 130, 58, 130, 130, 130, 7, 22, -70,
	-70, -70,
}
var yyDef = [...]int{

	0, -2, 1, 2, 3, 4, 5, 6, 7, 8,
	9, 10, 11, 12, 86, 86, 86, 86, 86, 72,
	250, 241, 0, 0, 0, 255, 255, 255, 0, 90,
	92, 93, 94, 95, 88, 0, 0, 0, 0, 239,
	0, 0, 251, 0, 0, 242, 0, 237, 0, 237,
	0, 83, 84, 85, 14, 91, 0, 96, 87, 0,
	0, 128, 0, 19, 234, 0, 196, 254, 0, 0,
	0, 255, 254, 0, 255, 0, 0, 0, 0, 0,
	82, 0, 97, 99, 104, 254, 102, 103, 138, 0,
	0, 166, 167, 168, 0, 196, 0, 182, 0, 198,
	199, 200, 201, 0, 233, 185, 186, 187, 183, 184,
	189, 89, 222, 0, 0, 136, 0, 0, 0, 0,
	0, 252, 74, 0, 77, 0, 79, 238, 0, 255,
	0, 0, 100, 105, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 153, 154, 155, 156, 157, 158, 159, 141,
	0, 0, 0, 0, 164, 177, 0, 0, 0, 152,
	0, 202, 0, 190, 0, 0, 0, 136, 129, 207,
	0, 235, 236, 197, 0, 240, 0, 0, 255, 248,
	243, 244, 245, 246, 247, 78, 80, 81, 136, 107,
	113, 0, 125, 127, 98, 106, 101, 139, 140, 143,
	144, 0, 161, 162, 0, 0, 0, 146, 0, 150,
	0, 169, 170, 171, 172, 173, 174, 175, 176, 142,
	163, 0, 232, 164, 178, 0, 0, 0, 194, 191,
	0, 226, 0, 229, 226, 0, 224, 207, 215, 0,
	137, 0, 69, 0, 0, 253, 75, 0, 249, 203,
	0, 0, 116, 117, 0, 0, 0, 0, 0, 130,
	114, 0, 0, 0, 0, 145, 147, 0, 0, 151,
	165, 179, 0, 181, 0, 192, 0, 0, 15, 0,
	228, 230, 16, 223, 0, 215, 18, 0, 0, 0,
	71, 55, 53, 23, 24, 51, 34, 51, 51, 32,
	25, 26, 27, 28, 29, 35, 36, 37, 38, 39,
	40, 41, 49, 49, 49, 49, 49, 255, 76, 205,
	0, 108, 111, 118, 0, 120, 0, 122, 123, 124,
	109, 0, 0, 0, 115, 110, 126, 160, 0, 148,
	180, 188, 195, 0, 0, 0, 225, 17, 216, 208,
	209, 212, 70, 68, 20, 54, 33, 0, 30, 31,
	42, 0, 43, 44, 45, 46, 0, 47, 48, 73,
	207, 0, 0, 0, 119, 121, 0, 0, 0, 149,
	193, 0, 231, 0, 0, 0, 211, 213, 214, 0,
	57, 0, 60, 61, 62, 63, 0, 65, 66, 22,
	21, 0, 0, 0, 215, 206, 204, 112, 0, 0,
	0, 0, 217, 218, 210, 56, 58, 59, 64, 67,
	52, 0, 219, 0, 134, 0, 0, 227, 0, 13,
	0, 0, 131, 0, 132, 133, 50, 220, 0, 135,
	0, 221,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 80, 73, 3,
	47, 130, 78, 76, 58, 77, 81, 79, 3, 3,
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
	127, 128, 129,
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
		//line sql.y:181
		{
			SetParseTree(yylex, yyDollar[1].statement)
		}
	case 2:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:187
		{
			yyVAL.statement = yyDollar[1].selStmt
		}
	case 13:
		yyDollar = yyS[yypt-12 : yypt+1]
		//line sql.y:203
		{
			yyVAL.selStmt = &Select{Comments: Comments(yyDollar[2].bytes2), Distinct: yyDollar[3].str, SelectExprs: yyDollar[4].selectExprs, From: yyDollar[6].tableExprs, Where: NewWhere(AST_WHERE, yyDollar[7].boolExpr), GroupBy: GroupBy(yyDollar[8].valExprs), Having: NewWhere(AST_HAVING, yyDollar[9].boolExpr), OrderBy: yyDollar[10].orderBy, Limit: yyDollar[11].limit, Lock: yyDollar[12].str}
		}
	case 14:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:207
		{
			yyVAL.selStmt = &Union{Type: yyDollar[2].str, Left: yyDollar[1].selStmt, Right: yyDollar[3].selStmt}
		}
	case 15:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line sql.y:213
		{
			yyVAL.statement = &Insert{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[4].tableName, Columns: yyDollar[5].columns, Rows: yyDollar[6].insRows, OnDup: OnDup(yyDollar[7].updateExprs)}
		}
	case 16:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line sql.y:217
		{
			cols := make(Columns, 0, len(yyDollar[6].updateExprs))
			vals := make(ValTuple, 0, len(yyDollar[6].updateExprs))
			for _, col := range yyDollar[6].updateExprs {
				cols = append(cols, &NonStarExpr{Expr: col.Name})
				vals = append(vals, col.Expr)
			}
			yyVAL.statement = &Insert{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[4].tableName, Columns: cols, Rows: Values{vals}, OnDup: OnDup(yyDollar[7].updateExprs)}
		}
	case 17:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line sql.y:229
		{
			yyVAL.statement = &Update{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[3].tableName, Exprs: yyDollar[5].updateExprs, Where: NewWhere(AST_WHERE, yyDollar[6].boolExpr), OrderBy: yyDollar[7].orderBy, Limit: yyDollar[8].limit}
		}
	case 18:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line sql.y:235
		{
			yyVAL.statement = &Delete{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[4].tableName, Where: NewWhere(AST_WHERE, yyDollar[5].boolExpr), OrderBy: yyDollar[6].orderBy, Limit: yyDollar[7].limit}
		}
	case 19:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:241
		{
			yyVAL.statement = &Set{Comments: Comments(yyDollar[2].bytes2), Exprs: yyDollar[3].updateExprs}
		}
	case 20:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:246
		{
			yyVAL.str = ""
		}
	case 21:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:250
		{
			yyVAL.str = AST_ZEROFILL
		}
	case 22:
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
	case 25:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:269
		{
			yyVAL.str = AST_DATE
		}
	case 26:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:273
		{
			yyVAL.str = AST_TIME
		}
	case 27:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:277
		{
			yyVAL.str = AST_TIMESTAMP
		}
	case 28:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:281
		{
			yyVAL.str = AST_DATETIME
		}
	case 29:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:285
		{
			yyVAL.str = AST_YEAR
		}
	case 30:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:291
		{
			if yyDollar[2].str == "" {
				yyVAL.str = AST_CHAR
			} else {
				yyVAL.str = AST_CHAR + yyDollar[2].str
			}
		}
	case 31:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:299
		{
			if yyDollar[2].str == "" {
				yyVAL.str = AST_VARCHAR
			} else {
				yyVAL.str = AST_VARCHAR + yyDollar[2].str
			}
		}
	case 32:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:307
		{
			yyVAL.str = AST_TEXT
		}
	case 33:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:313
		{
			yyVAL.str = yyDollar[1].str + yyDollar[2].str
		}
	case 34:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:317
		{
			yyVAL.str = yyDollar[1].str
		}
	case 35:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:323
		{
			yyVAL.str = AST_BIT
		}
	case 36:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:327
		{
			yyVAL.str = AST_TINYINT
		}
	case 37:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:331
		{
			yyVAL.str = AST_SMALLINT
		}
	case 38:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:335
		{
			yyVAL.str = AST_MEDIUMINT
		}
	case 39:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:339
		{
			yyVAL.str = AST_INT
		}
	case 40:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:343
		{
			yyVAL.str = AST_INTEGER
		}
	case 41:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:347
		{
			yyVAL.str = AST_BIGINT
		}
	case 42:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:353
		{
			yyVAL.str = AST_REAL + yyDollar[2].str
		}
	case 43:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:357
		{
			yyVAL.str = AST_DOUBLE + yyDollar[2].str
		}
	case 44:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:361
		{
			yyVAL.str = AST_FLOAT + yyDollar[2].str
		}
	case 45:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:365
		{
			yyVAL.str = AST_DECIMAL + yyDollar[2].str
		}
	case 46:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:369
		{
			yyVAL.str = AST_DECIMAL + yyDollar[2].str
		}
	case 47:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:373
		{
			yyVAL.str = AST_NUMERIC + yyDollar[2].str
		}
	case 48:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:377
		{
			yyVAL.str = AST_NUMERIC + yyDollar[2].str
		}
	case 49:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:382
		{
			yyVAL.str = ""
		}
	case 50:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:386
		{
			yyVAL.str = "(" + string(yyDollar[2].bytes) + ", " + string(yyDollar[4].bytes) + ")"
		}
	case 51:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:391
		{
			yyVAL.str = ""
		}
	case 52:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:395
		{
			yyVAL.str = "(" + string(yyDollar[2].bytes) + ")"
		}
	case 53:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:400
		{
			yyVAL.str = ""
		}
	case 54:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:404
		{
			yyVAL.str = AST_UNSIGNED
		}
	case 55:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:409
		{
			yyVAL.columnAtts = ColumnAtts{}
		}
	case 56:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:413
		{
			yyVAL.columnAtts = append(yyVAL.columnAtts, AST_NOT_NULL)
		}
	case 58:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:419
		{
			node := StrVal(yyDollar[3].bytes)
			yyVAL.columnAtts = append(yyVAL.columnAtts, "default "+String(node))
		}
	case 59:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:424
		{
			node := NumVal(yyDollar[3].bytes)
			yyVAL.columnAtts = append(yyVAL.columnAtts, "default "+String(node))
		}
	case 60:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:429
		{
			yyVAL.columnAtts = append(yyVAL.columnAtts, AST_AUTO_INCREMENT)
		}
	case 61:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:433
		{
			yyVAL.columnAtts = append(yyVAL.columnAtts, yyDollar[2].str)
		}
	case 62:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:439
		{
			yyVAL.str = AST_PRIMARY_KEY
		}
	case 63:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:443
		{
			yyVAL.str = AST_UNIQUE_KEY
		}
	case 68:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:457
		{
			yyVAL.columnDefinition = &ColumnDefinition{ColName: string(yyDollar[1].bytes), ColType: yyDollar[2].str, ColumnAtts: yyDollar[3].columnAtts}
		}
	case 69:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:463
		{
			yyVAL.columnDefinitions = ColumnDefinitions{yyDollar[1].columnDefinition}
		}
	case 70:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:467
		{
			yyVAL.columnDefinitions = append(yyVAL.columnDefinitions, yyDollar[3].columnDefinition)
		}
	case 71:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line sql.y:473
		{
			yyVAL.statement = &CreateTable{Name: yyDollar[4].bytes, ColumnDefinitions: yyDollar[6].columnDefinitions}
		}
	case 72:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:479
		{
			yyVAL.statement = yyDollar[1].statement
		}
	case 73:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line sql.y:483
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[7].bytes, NewName: yyDollar[7].bytes}
		}
	case 74:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:488
		{
			yyVAL.statement = &DDL{Action: AST_CREATE, NewName: yyDollar[3].bytes}
		}
	case 75:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line sql.y:494
		{
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[4].bytes, NewName: yyDollar[4].bytes}
		}
	case 76:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line sql.y:498
		{
			// Change this to a rename statement
			yyVAL.statement = &DDL{Action: AST_RENAME, Table: yyDollar[4].bytes, NewName: yyDollar[7].bytes}
		}
	case 77:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:503
		{
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[3].bytes, NewName: yyDollar[3].bytes}
		}
	case 78:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:509
		{
			yyVAL.statement = &DDL{Action: AST_RENAME, Table: yyDollar[3].bytes, NewName: yyDollar[5].bytes}
		}
	case 79:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:515
		{
			yyVAL.statement = &DDL{Action: AST_DROP, Table: yyDollar[4].bytes}
		}
	case 80:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:519
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[5].bytes, NewName: yyDollar[5].bytes}
		}
	case 81:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:524
		{
			yyVAL.statement = &DDL{Action: AST_DROP, Table: yyDollar[4].bytes}
		}
	case 82:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:530
		{
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[3].bytes, NewName: yyDollar[3].bytes}
		}
	case 83:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:536
		{
			yyVAL.statement = &Other{}
		}
	case 84:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:540
		{
			yyVAL.statement = &Other{}
		}
	case 85:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:544
		{
			yyVAL.statement = &Other{}
		}
	case 86:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:549
		{
			SetAllowComments(yylex, true)
		}
	case 87:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:553
		{
			yyVAL.bytes2 = yyDollar[2].bytes2
			SetAllowComments(yylex, false)
		}
	case 88:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:559
		{
			yyVAL.bytes2 = nil
		}
	case 89:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:563
		{
			yyVAL.bytes2 = append(yyDollar[1].bytes2, yyDollar[2].bytes)
		}
	case 90:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:569
		{
			yyVAL.str = AST_UNION
		}
	case 91:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:573
		{
			yyVAL.str = AST_UNION_ALL
		}
	case 92:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:577
		{
			yyVAL.str = AST_SET_MINUS
		}
	case 93:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:581
		{
			yyVAL.str = AST_EXCEPT
		}
	case 94:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:585
		{
			yyVAL.str = AST_INTERSECT
		}
	case 95:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:590
		{
			yyVAL.str = ""
		}
	case 96:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:594
		{
			yyVAL.str = AST_DISTINCT
		}
	case 97:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:600
		{
			yyVAL.selectExprs = SelectExprs{yyDollar[1].selectExpr}
		}
	case 98:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:604
		{
			yyVAL.selectExprs = append(yyVAL.selectExprs, yyDollar[3].selectExpr)
		}
	case 99:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:610
		{
			yyVAL.selectExpr = &StarExpr{}
		}
	case 100:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:614
		{
			yyVAL.selectExpr = &NonStarExpr{Expr: yyDollar[1].expr, As: yyDollar[2].bytes}
		}
	case 101:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:618
		{
			yyVAL.selectExpr = &StarExpr{TableName: yyDollar[1].bytes}
		}
	case 102:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:624
		{
			yyVAL.expr = yyDollar[1].boolExpr
		}
	case 103:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:628
		{
			yyVAL.expr = yyDollar[1].valExpr
		}
	case 104:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:633
		{
			yyVAL.bytes = nil
		}
	case 105:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:637
		{
			yyVAL.bytes = yyDollar[1].bytes
		}
	case 106:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:641
		{
			yyVAL.bytes = yyDollar[2].bytes
		}
	case 107:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:647
		{
			yyVAL.tableExprs = TableExprs{yyDollar[1].tableExpr}
		}
	case 108:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:651
		{
			yyVAL.tableExprs = append(yyVAL.tableExprs, yyDollar[3].tableExpr)
		}
	case 109:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:657
		{
			yyVAL.tableExpr = &AliasedTableExpr{Expr: yyDollar[1].smTableExpr, As: yyDollar[2].bytes, Hints: yyDollar[3].indexHints}
		}
	case 110:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:661
		{
			yyVAL.tableExpr = &ParenTableExpr{Expr: yyDollar[2].tableExpr}
		}
	case 111:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:665
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr}
		}
	case 112:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:669
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr, On: yyDollar[5].boolExpr}
		}
	case 113:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:674
		{
			yyVAL.bytes = nil
		}
	case 114:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:678
		{
			yyVAL.bytes = yyDollar[1].bytes
		}
	case 115:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:682
		{
			yyVAL.bytes = yyDollar[2].bytes
		}
	case 116:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:688
		{
			yyVAL.str = AST_JOIN
		}
	case 117:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:692
		{
			yyVAL.str = AST_STRAIGHT_JOIN
		}
	case 118:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:696
		{
			yyVAL.str = AST_LEFT_JOIN
		}
	case 119:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:700
		{
			yyVAL.str = AST_LEFT_JOIN
		}
	case 120:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:704
		{
			yyVAL.str = AST_RIGHT_JOIN
		}
	case 121:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:708
		{
			yyVAL.str = AST_RIGHT_JOIN
		}
	case 122:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:712
		{
			yyVAL.str = AST_JOIN
		}
	case 123:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:716
		{
			yyVAL.str = AST_CROSS_JOIN
		}
	case 124:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:720
		{
			yyVAL.str = AST_NATURAL_JOIN
		}
	case 125:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:726
		{
			yyVAL.smTableExpr = &TableName{Name: yyDollar[1].bytes}
		}
	case 126:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:730
		{
			yyVAL.smTableExpr = &TableName{Qualifier: yyDollar[1].bytes, Name: yyDollar[3].bytes}
		}
	case 127:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:734
		{
			yyVAL.smTableExpr = yyDollar[1].subquery
		}
	case 128:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:740
		{
			yyVAL.tableName = &TableName{Name: yyDollar[1].bytes}
		}
	case 129:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:744
		{
			yyVAL.tableName = &TableName{Qualifier: yyDollar[1].bytes, Name: yyDollar[3].bytes}
		}
	case 130:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:749
		{
			yyVAL.indexHints = nil
		}
	case 131:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:753
		{
			yyVAL.indexHints = &IndexHints{Type: AST_USE, Indexes: yyDollar[4].bytes2}
		}
	case 132:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:757
		{
			yyVAL.indexHints = &IndexHints{Type: AST_IGNORE, Indexes: yyDollar[4].bytes2}
		}
	case 133:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:761
		{
			yyVAL.indexHints = &IndexHints{Type: AST_FORCE, Indexes: yyDollar[4].bytes2}
		}
	case 134:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:767
		{
			yyVAL.bytes2 = [][]byte{yyDollar[1].bytes}
		}
	case 135:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:771
		{
			yyVAL.bytes2 = append(yyDollar[1].bytes2, yyDollar[3].bytes)
		}
	case 136:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:776
		{
			yyVAL.boolExpr = nil
		}
	case 137:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:780
		{
			yyVAL.boolExpr = yyDollar[2].boolExpr
		}
	case 139:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:787
		{
			yyVAL.boolExpr = &AndExpr{Left: yyDollar[1].boolExpr, Right: yyDollar[3].boolExpr}
		}
	case 140:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:791
		{
			yyVAL.boolExpr = &OrExpr{Left: yyDollar[1].boolExpr, Right: yyDollar[3].boolExpr}
		}
	case 141:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:795
		{
			yyVAL.boolExpr = &NotExpr{Expr: yyDollar[2].boolExpr}
		}
	case 142:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:799
		{
			yyVAL.boolExpr = &ParenBoolExpr{Expr: yyDollar[2].boolExpr}
		}
	case 143:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:805
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: yyDollar[2].str, Right: yyDollar[3].valExpr}
		}
	case 144:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:809
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_IN, Right: yyDollar[3].colTuple}
		}
	case 145:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:813
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_NOT_IN, Right: yyDollar[4].colTuple}
		}
	case 146:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:817
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_LIKE, Right: yyDollar[3].valExpr}
		}
	case 147:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:821
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_NOT_LIKE, Right: yyDollar[4].valExpr}
		}
	case 148:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:825
		{
			yyVAL.boolExpr = &RangeCond{Left: yyDollar[1].valExpr, Operator: AST_BETWEEN, From: yyDollar[3].valExpr, To: yyDollar[5].valExpr}
		}
	case 149:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line sql.y:829
		{
			yyVAL.boolExpr = &RangeCond{Left: yyDollar[1].valExpr, Operator: AST_NOT_BETWEEN, From: yyDollar[4].valExpr, To: yyDollar[6].valExpr}
		}
	case 150:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:833
		{
			yyVAL.boolExpr = &NullCheck{Operator: AST_IS_NULL, Expr: yyDollar[1].valExpr}
		}
	case 151:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:837
		{
			yyVAL.boolExpr = &NullCheck{Operator: AST_IS_NOT_NULL, Expr: yyDollar[1].valExpr}
		}
	case 152:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:841
		{
			yyVAL.boolExpr = &ExistsExpr{Subquery: yyDollar[2].subquery}
		}
	case 153:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:847
		{
			yyVAL.str = AST_EQ
		}
	case 154:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:851
		{
			yyVAL.str = AST_LT
		}
	case 155:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:855
		{
			yyVAL.str = AST_GT
		}
	case 156:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:859
		{
			yyVAL.str = AST_LE
		}
	case 157:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:863
		{
			yyVAL.str = AST_GE
		}
	case 158:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:867
		{
			yyVAL.str = AST_NE
		}
	case 159:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:871
		{
			yyVAL.str = AST_NSE
		}
	case 160:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:877
		{
			yyVAL.colTuple = ValTuple(yyDollar[2].valExprs)
		}
	case 161:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:881
		{
			yyVAL.colTuple = yyDollar[1].subquery
		}
	case 162:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:885
		{
			yyVAL.colTuple = ListArg(yyDollar[1].bytes)
		}
	case 163:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:891
		{
			yyVAL.subquery = &Subquery{yyDollar[2].selStmt}
		}
	case 164:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:897
		{
			yyVAL.valExprs = ValExprs{yyDollar[1].valExpr}
		}
	case 165:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:901
		{
			yyVAL.valExprs = append(yyDollar[1].valExprs, yyDollar[3].valExpr)
		}
	case 166:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:907
		{
			yyVAL.valExpr = yyDollar[1].valExpr
		}
	case 167:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:911
		{
			yyVAL.valExpr = yyDollar[1].colName
		}
	case 168:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:915
		{
			yyVAL.valExpr = yyDollar[1].rowTuple
		}
	case 169:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:919
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_BITAND, Right: yyDollar[3].valExpr}
		}
	case 170:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:923
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_BITOR, Right: yyDollar[3].valExpr}
		}
	case 171:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:927
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_BITXOR, Right: yyDollar[3].valExpr}
		}
	case 172:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:931
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_PLUS, Right: yyDollar[3].valExpr}
		}
	case 173:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:935
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_MINUS, Right: yyDollar[3].valExpr}
		}
	case 174:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:939
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_MULT, Right: yyDollar[3].valExpr}
		}
	case 175:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:943
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_DIV, Right: yyDollar[3].valExpr}
		}
	case 176:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:947
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_MOD, Right: yyDollar[3].valExpr}
		}
	case 177:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:951
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
	case 178:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:966
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].bytes}
		}
	case 179:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:970
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].bytes, Exprs: yyDollar[3].selectExprs}
		}
	case 180:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:974
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].bytes, Distinct: true, Exprs: yyDollar[4].selectExprs}
		}
	case 181:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:978
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].bytes, Exprs: yyDollar[3].selectExprs}
		}
	case 182:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:982
		{
			yyVAL.valExpr = yyDollar[1].caseExpr
		}
	case 183:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:988
		{
			yyVAL.bytes = IF_BYTES
		}
	case 184:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:992
		{
			yyVAL.bytes = VALUES_BYTES
		}
	case 185:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:998
		{
			yyVAL.byt = AST_UPLUS
		}
	case 186:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1002
		{
			yyVAL.byt = AST_UMINUS
		}
	case 187:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1006
		{
			yyVAL.byt = AST_TILDA
		}
	case 188:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:1012
		{
			yyVAL.caseExpr = &CaseExpr{Expr: yyDollar[2].valExpr, Whens: yyDollar[3].whens, Else: yyDollar[4].valExpr}
		}
	case 189:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1017
		{
			yyVAL.valExpr = nil
		}
	case 190:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1021
		{
			yyVAL.valExpr = yyDollar[1].valExpr
		}
	case 191:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1027
		{
			yyVAL.whens = []*When{yyDollar[1].when}
		}
	case 192:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1031
		{
			yyVAL.whens = append(yyDollar[1].whens, yyDollar[2].when)
		}
	case 193:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:1037
		{
			yyVAL.when = &When{Cond: yyDollar[2].boolExpr, Val: yyDollar[4].valExpr}
		}
	case 194:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1042
		{
			yyVAL.valExpr = nil
		}
	case 195:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1046
		{
			yyVAL.valExpr = yyDollar[2].valExpr
		}
	case 196:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1052
		{
			yyVAL.colName = &ColName{Name: yyDollar[1].bytes}
		}
	case 197:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1056
		{
			yyVAL.colName = &ColName{Qualifier: yyDollar[1].bytes, Name: yyDollar[3].bytes}
		}
	case 198:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1062
		{
			yyVAL.valExpr = StrVal(yyDollar[1].bytes)
		}
	case 199:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1066
		{
			yyVAL.valExpr = NumVal(yyDollar[1].bytes)
		}
	case 200:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1070
		{
			yyVAL.valExpr = ValArg(yyDollar[1].bytes)
		}
	case 201:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1074
		{
			yyVAL.valExpr = &NullVal{}
		}
	case 202:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1078
		{
			yyVAL.valExpr = BinaryVal(yyDollar[2].bytes)
		}
	case 203:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1083
		{
			yyVAL.valExprs = nil
		}
	case 204:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1087
		{
			yyVAL.valExprs = yyDollar[3].valExprs
		}
	case 205:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1092
		{
			yyVAL.boolExpr = nil
		}
	case 206:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1096
		{
			yyVAL.boolExpr = yyDollar[2].boolExpr
		}
	case 207:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1101
		{
			yyVAL.orderBy = nil
		}
	case 208:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1105
		{
			yyVAL.orderBy = yyDollar[3].orderBy
		}
	case 209:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1111
		{
			yyVAL.orderBy = OrderBy{yyDollar[1].order}
		}
	case 210:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1115
		{
			yyVAL.orderBy = append(yyDollar[1].orderBy, yyDollar[3].order)
		}
	case 211:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1121
		{
			yyVAL.order = &Order{Expr: yyDollar[1].valExpr, Direction: yyDollar[2].str}
		}
	case 212:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1126
		{
			yyVAL.str = AST_ASC
		}
	case 213:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1130
		{
			yyVAL.str = AST_ASC
		}
	case 214:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1134
		{
			yyVAL.str = AST_DESC
		}
	case 215:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1139
		{
			yyVAL.limit = nil
		}
	case 216:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1143
		{
			yyVAL.limit = &Limit{Rowcount: yyDollar[2].valExpr}
		}
	case 217:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:1147
		{
			yyVAL.limit = &Limit{Offset: yyDollar[2].valExpr, Rowcount: yyDollar[4].valExpr}
		}
	case 218:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:1151
		{
			yyVAL.limit = &Limit{Offset: yyDollar[4].valExpr, Rowcount: yyDollar[2].valExpr}
		}
	case 219:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1156
		{
			yyVAL.str = ""
		}
	case 220:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1160
		{
			yyVAL.str = AST_FOR_UPDATE
		}
	case 221:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:1164
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
	case 222:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1177
		{
			yyVAL.columns = nil
		}
	case 223:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1181
		{
			yyVAL.columns = yyDollar[2].columns
		}
	case 224:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1187
		{
			yyVAL.columns = Columns{&NonStarExpr{Expr: yyDollar[1].colName}}
		}
	case 225:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1191
		{
			yyVAL.columns = append(yyVAL.columns, &NonStarExpr{Expr: yyDollar[3].colName})
		}
	case 226:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1196
		{
			yyVAL.updateExprs = nil
		}
	case 227:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:1200
		{
			yyVAL.updateExprs = yyDollar[5].updateExprs
		}
	case 228:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1206
		{
			yyVAL.insRows = yyDollar[2].values
		}
	case 229:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1210
		{
			yyVAL.insRows = yyDollar[1].selStmt
		}
	case 230:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1216
		{
			yyVAL.values = Values{yyDollar[1].rowTuple}
		}
	case 231:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1220
		{
			yyVAL.values = append(yyDollar[1].values, yyDollar[3].rowTuple)
		}
	case 232:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1226
		{
			yyVAL.rowTuple = ValTuple(yyDollar[2].valExprs)
		}
	case 233:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1230
		{
			yyVAL.rowTuple = yyDollar[1].subquery
		}
	case 234:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1236
		{
			yyVAL.updateExprs = UpdateExprs{yyDollar[1].updateExpr}
		}
	case 235:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1240
		{
			yyVAL.updateExprs = append(yyDollar[1].updateExprs, yyDollar[3].updateExpr)
		}
	case 236:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1246
		{
			yyVAL.updateExpr = &UpdateExpr{Name: yyDollar[1].colName, Expr: yyDollar[3].valExpr}
		}
	case 237:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1251
		{
			yyVAL.empty = struct{}{}
		}
	case 238:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1253
		{
			yyVAL.empty = struct{}{}
		}
	case 239:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1256
		{
			yyVAL.empty = struct{}{}
		}
	case 240:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1258
		{
			yyVAL.empty = struct{}{}
		}
	case 241:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1261
		{
			yyVAL.empty = struct{}{}
		}
	case 242:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1263
		{
			yyVAL.empty = struct{}{}
		}
	case 243:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1267
		{
			yyVAL.empty = struct{}{}
		}
	case 244:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1269
		{
			yyVAL.empty = struct{}{}
		}
	case 245:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1271
		{
			yyVAL.empty = struct{}{}
		}
	case 246:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1273
		{
			yyVAL.empty = struct{}{}
		}
	case 247:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1275
		{
			yyVAL.empty = struct{}{}
		}
	case 248:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1278
		{
			yyVAL.empty = struct{}{}
		}
	case 249:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1280
		{
			yyVAL.empty = struct{}{}
		}
	case 250:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1283
		{
			yyVAL.empty = struct{}{}
		}
	case 251:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1285
		{
			yyVAL.empty = struct{}{}
		}
	case 252:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1288
		{
			yyVAL.empty = struct{}{}
		}
	case 253:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1290
		{
			yyVAL.empty = struct{}{}
		}
	case 254:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1294
		{
			yyVAL.bytes = bytes.ToLower(yyDollar[1].bytes)
		}
	case 255:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1299
		{
			ForceEOF(yylex)
		}
	}
	goto yystack /* stack new state and value */
}
