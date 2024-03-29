// Code generated from ZitiQl.g4 by ANTLR 4.13.1. DO NOT EDIT.

package zitiql

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type ZitiQlLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var ZitiQlLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func zitiqllexerLexerInit() {
	staticData := &ZitiQlLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "','", "", "'('", "')'", "'['", "']'",
	}
	staticData.SymbolicNames = []string{
		"", "", "WS", "LPAREN", "RPAREN", "LBRACKET", "RBRACKET", "AND", "OR",
		"LT", "GT", "EQ", "CONTAINS", "IN", "BETWEEN", "BOOL", "DATETIME", "ALL_OF",
		"ANY_OF", "COUNT", "ISEMPTY", "STRING", "NUMBER", "NULL", "NOT", "ASC",
		"DESC", "SORT", "BY", "SKIP_ROWS", "LIMIT_ROWS", "NONE", "WHERE", "FROM",
		"IDENTIFIER", "RFC3339_DATE_TIME",
	}
	staticData.RuleNames = []string{
		"T__0", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L",
		"M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
		"WS", "LPAREN", "RPAREN", "LBRACKET", "RBRACKET", "AND", "OR", "LT",
		"GT", "EQ", "CONTAINS", "IN", "BETWEEN", "BOOL", "TRUE", "FALSE", "DATETIME",
		"FULL_DATE", "FULL_TIME", "YEAR", "MONTH", "DAY", "HOUR", "MINUTE",
		"SECOND", "SECOND_FRACTION", "TIME_NUM_OFFSET", "TIME_NUM_OFFSET_HOUR",
		"TIME_NUM_OFFSET_MINUTE", "ALL_OF", "ANY_OF", "COUNT", "ISEMPTY", "STRING",
		"NUMBER", "NULL", "NOT", "ASC", "DESC", "SORT", "BY", "SKIP_ROWS", "LIMIT_ROWS",
		"NONE", "WHERE", "FROM", "INT", "EXP", "ESC", "IDENTSEP", "IDENTIFIER",
		"RFC3339_DATE_TIME", "SAFECODEPOINT",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 35, 580, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7,
		41, 2, 42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46,
		2, 47, 7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 2, 50, 7, 50, 2, 51, 7, 51, 2,
		52, 7, 52, 2, 53, 7, 53, 2, 54, 7, 54, 2, 55, 7, 55, 2, 56, 7, 56, 2, 57,
		7, 57, 2, 58, 7, 58, 2, 59, 7, 59, 2, 60, 7, 60, 2, 61, 7, 61, 2, 62, 7,
		62, 2, 63, 7, 63, 2, 64, 7, 64, 2, 65, 7, 65, 2, 66, 7, 66, 2, 67, 7, 67,
		2, 68, 7, 68, 2, 69, 7, 69, 2, 70, 7, 70, 2, 71, 7, 71, 2, 72, 7, 72, 2,
		73, 7, 73, 2, 74, 7, 74, 2, 75, 7, 75, 2, 76, 7, 76, 2, 77, 7, 77, 2, 78,
		7, 78, 2, 79, 7, 79, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1,
		4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1,
		10, 1, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 15,
		1, 15, 1, 16, 1, 16, 1, 17, 1, 17, 1, 18, 1, 18, 1, 19, 1, 19, 1, 20, 1,
		20, 1, 21, 1, 21, 1, 22, 1, 22, 1, 23, 1, 23, 1, 24, 1, 24, 1, 25, 1, 25,
		1, 26, 1, 26, 1, 27, 1, 27, 1, 28, 1, 28, 1, 29, 1, 29, 1, 30, 1, 30, 1,
		31, 1, 31, 1, 32, 1, 32, 1, 32, 1, 32, 1, 33, 1, 33, 1, 33, 1, 34, 1, 34,
		3, 34, 235, 8, 34, 1, 35, 1, 35, 3, 35, 239, 8, 35, 1, 36, 3, 36, 242,
		8, 36, 1, 36, 1, 36, 1, 37, 1, 37, 1, 37, 1, 37, 4, 37, 250, 8, 37, 11,
		37, 12, 37, 251, 3, 37, 254, 8, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37,
		1, 37, 1, 37, 1, 37, 1, 37, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 3, 38, 270,
		8, 38, 1, 38, 1, 38, 1, 38, 1, 39, 1, 39, 1, 39, 1, 39, 4, 39, 279, 8,
		39, 11, 39, 12, 39, 280, 3, 39, 283, 8, 39, 1, 39, 1, 39, 1, 39, 1, 39,
		1, 39, 1, 39, 1, 39, 1, 39, 1, 40, 1, 40, 3, 40, 295, 8, 40, 1, 41, 1,
		41, 1, 41, 1, 41, 1, 41, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 43,
		1, 43, 1, 43, 1, 43, 1, 43, 1, 43, 1, 43, 1, 43, 1, 43, 1, 43, 1, 43, 5,
		43, 319, 8, 43, 10, 43, 12, 43, 322, 9, 43, 1, 43, 1, 43, 5, 43, 326, 8,
		43, 10, 43, 12, 43, 329, 9, 43, 1, 43, 1, 43, 1, 44, 1, 44, 1, 44, 1, 44,
		1, 44, 1, 44, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 3, 45, 346,
		8, 45, 1, 45, 1, 45, 1, 46, 4, 46, 351, 8, 46, 11, 46, 12, 46, 352, 1,
		47, 1, 47, 1, 47, 1, 47, 3, 47, 359, 8, 47, 1, 48, 1, 48, 1, 48, 1, 48,
		1, 48, 1, 48, 3, 48, 367, 8, 48, 1, 49, 1, 49, 1, 49, 1, 49, 3, 49, 373,
		8, 49, 1, 50, 1, 50, 1, 50, 1, 51, 1, 51, 1, 51, 1, 51, 3, 51, 382, 8,
		51, 1, 52, 4, 52, 385, 8, 52, 11, 52, 12, 52, 386, 1, 53, 1, 53, 1, 53,
		1, 53, 1, 53, 1, 53, 3, 53, 395, 8, 53, 1, 54, 1, 54, 1, 54, 1, 54, 3,
		54, 401, 8, 54, 1, 55, 1, 55, 1, 55, 1, 56, 1, 56, 1, 56, 1, 56, 1, 56,
		1, 56, 1, 57, 1, 57, 1, 57, 1, 57, 1, 57, 1, 57, 1, 58, 1, 58, 1, 58, 1,
		58, 1, 58, 1, 58, 1, 59, 1, 59, 1, 59, 1, 59, 1, 59, 1, 59, 1, 59, 1, 59,
		1, 60, 1, 60, 1, 60, 5, 60, 435, 8, 60, 10, 60, 12, 60, 438, 9, 60, 1,
		60, 1, 60, 1, 61, 3, 61, 443, 8, 61, 1, 61, 1, 61, 1, 61, 4, 61, 448, 8,
		61, 11, 61, 12, 61, 449, 3, 61, 452, 8, 61, 1, 61, 3, 61, 455, 8, 61, 1,
		62, 1, 62, 1, 62, 1, 62, 1, 62, 1, 63, 1, 63, 1, 63, 1, 63, 1, 64, 1, 64,
		1, 64, 1, 64, 1, 65, 1, 65, 1, 65, 1, 65, 1, 65, 1, 66, 1, 66, 1, 66, 1,
		66, 1, 66, 1, 67, 1, 67, 1, 67, 1, 68, 1, 68, 1, 68, 1, 68, 1, 68, 1, 69,
		1, 69, 1, 69, 1, 69, 1, 69, 1, 69, 1, 70, 1, 70, 1, 70, 1, 70, 1, 70, 1,
		71, 1, 71, 1, 71, 1, 71, 1, 71, 1, 71, 1, 72, 1, 72, 1, 72, 1, 72, 1, 72,
		1, 73, 1, 73, 1, 73, 5, 73, 513, 8, 73, 10, 73, 12, 73, 516, 9, 73, 3,
		73, 518, 8, 73, 1, 74, 1, 74, 3, 74, 522, 8, 74, 1, 74, 1, 74, 1, 75, 1,
		75, 1, 75, 1, 76, 1, 76, 1, 77, 1, 77, 5, 77, 533, 8, 77, 10, 77, 12, 77,
		536, 9, 77, 1, 77, 1, 77, 1, 77, 5, 77, 541, 8, 77, 10, 77, 12, 77, 544,
		9, 77, 5, 77, 546, 8, 77, 10, 77, 12, 77, 549, 9, 77, 1, 77, 1, 77, 1,
		77, 5, 77, 554, 8, 77, 10, 77, 12, 77, 557, 9, 77, 1, 77, 1, 77, 1, 77,
		5, 77, 562, 8, 77, 10, 77, 12, 77, 565, 9, 77, 5, 77, 567, 8, 77, 10, 77,
		12, 77, 570, 9, 77, 1, 77, 3, 77, 573, 8, 77, 1, 78, 1, 78, 1, 78, 1, 78,
		1, 79, 1, 79, 0, 0, 80, 1, 1, 3, 0, 5, 0, 7, 0, 9, 0, 11, 0, 13, 0, 15,
		0, 17, 0, 19, 0, 21, 0, 23, 0, 25, 0, 27, 0, 29, 0, 31, 0, 33, 0, 35, 0,
		37, 0, 39, 0, 41, 0, 43, 0, 45, 0, 47, 0, 49, 0, 51, 0, 53, 0, 55, 2, 57,
		3, 59, 4, 61, 5, 63, 6, 65, 7, 67, 8, 69, 9, 71, 10, 73, 11, 75, 12, 77,
		13, 79, 14, 81, 15, 83, 0, 85, 0, 87, 16, 89, 0, 91, 0, 93, 0, 95, 0, 97,
		0, 99, 0, 101, 0, 103, 0, 105, 0, 107, 0, 109, 0, 111, 0, 113, 17, 115,
		18, 117, 19, 119, 20, 121, 21, 123, 22, 125, 23, 127, 24, 129, 25, 131,
		26, 133, 27, 135, 28, 137, 29, 139, 30, 141, 31, 143, 32, 145, 33, 147,
		0, 149, 0, 151, 0, 153, 0, 155, 34, 157, 35, 159, 0, 1, 0, 39, 2, 0, 65,
		65, 97, 97, 2, 0, 66, 66, 98, 98, 2, 0, 67, 67, 99, 99, 2, 0, 68, 68, 100,
		100, 2, 0, 69, 69, 101, 101, 2, 0, 70, 70, 102, 102, 2, 0, 71, 71, 103,
		103, 2, 0, 72, 72, 104, 104, 2, 0, 73, 73, 105, 105, 2, 0, 74, 74, 106,
		106, 2, 0, 75, 75, 107, 107, 2, 0, 76, 76, 108, 108, 2, 0, 77, 77, 109,
		109, 2, 0, 78, 78, 110, 110, 2, 0, 79, 79, 111, 111, 2, 0, 80, 80, 112,
		112, 2, 0, 81, 81, 113, 113, 2, 0, 82, 82, 114, 114, 2, 0, 83, 83, 115,
		115, 2, 0, 84, 84, 116, 116, 2, 0, 85, 85, 117, 117, 2, 0, 86, 86, 118,
		118, 2, 0, 87, 87, 119, 119, 2, 0, 88, 88, 120, 120, 2, 0, 89, 89, 121,
		121, 2, 0, 90, 90, 122, 122, 3, 0, 9, 10, 13, 13, 32, 32, 1, 0, 49, 57,
		1, 0, 48, 50, 1, 0, 49, 50, 1, 0, 48, 57, 1, 0, 48, 49, 1, 0, 48, 51, 1,
		0, 48, 53, 2, 0, 43, 43, 45, 45, 6, 0, 34, 34, 92, 92, 102, 102, 110, 110,
		114, 114, 116, 116, 2, 0, 65, 90, 97, 122, 3, 0, 65, 90, 95, 95, 97, 122,
		3, 0, 0, 31, 34, 34, 92, 92, 571, 0, 1, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0,
		0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 63, 1, 0, 0,
		0, 0, 65, 1, 0, 0, 0, 0, 67, 1, 0, 0, 0, 0, 69, 1, 0, 0, 0, 0, 71, 1, 0,
		0, 0, 0, 73, 1, 0, 0, 0, 0, 75, 1, 0, 0, 0, 0, 77, 1, 0, 0, 0, 0, 79, 1,
		0, 0, 0, 0, 81, 1, 0, 0, 0, 0, 87, 1, 0, 0, 0, 0, 113, 1, 0, 0, 0, 0, 115,
		1, 0, 0, 0, 0, 117, 1, 0, 0, 0, 0, 119, 1, 0, 0, 0, 0, 121, 1, 0, 0, 0,
		0, 123, 1, 0, 0, 0, 0, 125, 1, 0, 0, 0, 0, 127, 1, 0, 0, 0, 0, 129, 1,
		0, 0, 0, 0, 131, 1, 0, 0, 0, 0, 133, 1, 0, 0, 0, 0, 135, 1, 0, 0, 0, 0,
		137, 1, 0, 0, 0, 0, 139, 1, 0, 0, 0, 0, 141, 1, 0, 0, 0, 0, 143, 1, 0,
		0, 0, 0, 145, 1, 0, 0, 0, 0, 155, 1, 0, 0, 0, 0, 157, 1, 0, 0, 0, 1, 161,
		1, 0, 0, 0, 3, 163, 1, 0, 0, 0, 5, 165, 1, 0, 0, 0, 7, 167, 1, 0, 0, 0,
		9, 169, 1, 0, 0, 0, 11, 171, 1, 0, 0, 0, 13, 173, 1, 0, 0, 0, 15, 175,
		1, 0, 0, 0, 17, 177, 1, 0, 0, 0, 19, 179, 1, 0, 0, 0, 21, 181, 1, 0, 0,
		0, 23, 183, 1, 0, 0, 0, 25, 185, 1, 0, 0, 0, 27, 187, 1, 0, 0, 0, 29, 189,
		1, 0, 0, 0, 31, 191, 1, 0, 0, 0, 33, 193, 1, 0, 0, 0, 35, 195, 1, 0, 0,
		0, 37, 197, 1, 0, 0, 0, 39, 199, 1, 0, 0, 0, 41, 201, 1, 0, 0, 0, 43, 203,
		1, 0, 0, 0, 45, 205, 1, 0, 0, 0, 47, 207, 1, 0, 0, 0, 49, 209, 1, 0, 0,
		0, 51, 211, 1, 0, 0, 0, 53, 213, 1, 0, 0, 0, 55, 215, 1, 0, 0, 0, 57, 217,
		1, 0, 0, 0, 59, 219, 1, 0, 0, 0, 61, 221, 1, 0, 0, 0, 63, 223, 1, 0, 0,
		0, 65, 225, 1, 0, 0, 0, 67, 229, 1, 0, 0, 0, 69, 232, 1, 0, 0, 0, 71, 236,
		1, 0, 0, 0, 73, 241, 1, 0, 0, 0, 75, 253, 1, 0, 0, 0, 77, 269, 1, 0, 0,
		0, 79, 282, 1, 0, 0, 0, 81, 294, 1, 0, 0, 0, 83, 296, 1, 0, 0, 0, 85, 301,
		1, 0, 0, 0, 87, 307, 1, 0, 0, 0, 89, 332, 1, 0, 0, 0, 91, 338, 1, 0, 0,
		0, 93, 350, 1, 0, 0, 0, 95, 358, 1, 0, 0, 0, 97, 366, 1, 0, 0, 0, 99, 372,
		1, 0, 0, 0, 101, 374, 1, 0, 0, 0, 103, 381, 1, 0, 0, 0, 105, 384, 1, 0,
		0, 0, 107, 394, 1, 0, 0, 0, 109, 400, 1, 0, 0, 0, 111, 402, 1, 0, 0, 0,
		113, 405, 1, 0, 0, 0, 115, 411, 1, 0, 0, 0, 117, 417, 1, 0, 0, 0, 119,
		423, 1, 0, 0, 0, 121, 431, 1, 0, 0, 0, 123, 442, 1, 0, 0, 0, 125, 456,
		1, 0, 0, 0, 127, 461, 1, 0, 0, 0, 129, 465, 1, 0, 0, 0, 131, 469, 1, 0,
		0, 0, 133, 474, 1, 0, 0, 0, 135, 479, 1, 0, 0, 0, 137, 482, 1, 0, 0, 0,
		139, 487, 1, 0, 0, 0, 141, 493, 1, 0, 0, 0, 143, 498, 1, 0, 0, 0, 145,
		504, 1, 0, 0, 0, 147, 517, 1, 0, 0, 0, 149, 519, 1, 0, 0, 0, 151, 525,
		1, 0, 0, 0, 153, 528, 1, 0, 0, 0, 155, 572, 1, 0, 0, 0, 157, 574, 1, 0,
		0, 0, 159, 578, 1, 0, 0, 0, 161, 162, 5, 44, 0, 0, 162, 2, 1, 0, 0, 0,
		163, 164, 7, 0, 0, 0, 164, 4, 1, 0, 0, 0, 165, 166, 7, 1, 0, 0, 166, 6,
		1, 0, 0, 0, 167, 168, 7, 2, 0, 0, 168, 8, 1, 0, 0, 0, 169, 170, 7, 3, 0,
		0, 170, 10, 1, 0, 0, 0, 171, 172, 7, 4, 0, 0, 172, 12, 1, 0, 0, 0, 173,
		174, 7, 5, 0, 0, 174, 14, 1, 0, 0, 0, 175, 176, 7, 6, 0, 0, 176, 16, 1,
		0, 0, 0, 177, 178, 7, 7, 0, 0, 178, 18, 1, 0, 0, 0, 179, 180, 7, 8, 0,
		0, 180, 20, 1, 0, 0, 0, 181, 182, 7, 9, 0, 0, 182, 22, 1, 0, 0, 0, 183,
		184, 7, 10, 0, 0, 184, 24, 1, 0, 0, 0, 185, 186, 7, 11, 0, 0, 186, 26,
		1, 0, 0, 0, 187, 188, 7, 12, 0, 0, 188, 28, 1, 0, 0, 0, 189, 190, 7, 13,
		0, 0, 190, 30, 1, 0, 0, 0, 191, 192, 7, 14, 0, 0, 192, 32, 1, 0, 0, 0,
		193, 194, 7, 15, 0, 0, 194, 34, 1, 0, 0, 0, 195, 196, 7, 16, 0, 0, 196,
		36, 1, 0, 0, 0, 197, 198, 7, 17, 0, 0, 198, 38, 1, 0, 0, 0, 199, 200, 7,
		18, 0, 0, 200, 40, 1, 0, 0, 0, 201, 202, 7, 19, 0, 0, 202, 42, 1, 0, 0,
		0, 203, 204, 7, 20, 0, 0, 204, 44, 1, 0, 0, 0, 205, 206, 7, 21, 0, 0, 206,
		46, 1, 0, 0, 0, 207, 208, 7, 22, 0, 0, 208, 48, 1, 0, 0, 0, 209, 210, 7,
		23, 0, 0, 210, 50, 1, 0, 0, 0, 211, 212, 7, 24, 0, 0, 212, 52, 1, 0, 0,
		0, 213, 214, 7, 25, 0, 0, 214, 54, 1, 0, 0, 0, 215, 216, 7, 26, 0, 0, 216,
		56, 1, 0, 0, 0, 217, 218, 5, 40, 0, 0, 218, 58, 1, 0, 0, 0, 219, 220, 5,
		41, 0, 0, 220, 60, 1, 0, 0, 0, 221, 222, 5, 91, 0, 0, 222, 62, 1, 0, 0,
		0, 223, 224, 5, 93, 0, 0, 224, 64, 1, 0, 0, 0, 225, 226, 3, 3, 1, 0, 226,
		227, 3, 29, 14, 0, 227, 228, 3, 9, 4, 0, 228, 66, 1, 0, 0, 0, 229, 230,
		3, 31, 15, 0, 230, 231, 3, 37, 18, 0, 231, 68, 1, 0, 0, 0, 232, 234, 5,
		60, 0, 0, 233, 235, 5, 61, 0, 0, 234, 233, 1, 0, 0, 0, 234, 235, 1, 0,
		0, 0, 235, 70, 1, 0, 0, 0, 236, 238, 5, 62, 0, 0, 237, 239, 5, 61, 0, 0,
		238, 237, 1, 0, 0, 0, 238, 239, 1, 0, 0, 0, 239, 72, 1, 0, 0, 0, 240, 242,
		5, 33, 0, 0, 241, 240, 1, 0, 0, 0, 241, 242, 1, 0, 0, 0, 242, 243, 1, 0,
		0, 0, 243, 244, 5, 61, 0, 0, 244, 74, 1, 0, 0, 0, 245, 246, 3, 29, 14,
		0, 246, 247, 3, 31, 15, 0, 247, 249, 3, 41, 20, 0, 248, 250, 3, 55, 27,
		0, 249, 248, 1, 0, 0, 0, 250, 251, 1, 0, 0, 0, 251, 249, 1, 0, 0, 0, 251,
		252, 1, 0, 0, 0, 252, 254, 1, 0, 0, 0, 253, 245, 1, 0, 0, 0, 253, 254,
		1, 0, 0, 0, 254, 255, 1, 0, 0, 0, 255, 256, 3, 7, 3, 0, 256, 257, 3, 31,
		15, 0, 257, 258, 3, 29, 14, 0, 258, 259, 3, 41, 20, 0, 259, 260, 3, 3,
		1, 0, 260, 261, 3, 19, 9, 0, 261, 262, 3, 29, 14, 0, 262, 263, 3, 39, 19,
		0, 263, 76, 1, 0, 0, 0, 264, 265, 3, 29, 14, 0, 265, 266, 3, 31, 15, 0,
		266, 267, 3, 41, 20, 0, 267, 268, 3, 55, 27, 0, 268, 270, 1, 0, 0, 0, 269,
		264, 1, 0, 0, 0, 269, 270, 1, 0, 0, 0, 270, 271, 1, 0, 0, 0, 271, 272,
		3, 19, 9, 0, 272, 273, 3, 29, 14, 0, 273, 78, 1, 0, 0, 0, 274, 275, 3,
		29, 14, 0, 275, 276, 3, 31, 15, 0, 276, 278, 3, 41, 20, 0, 277, 279, 3,
		55, 27, 0, 278, 277, 1, 0, 0, 0, 279, 280, 1, 0, 0, 0, 280, 278, 1, 0,
		0, 0, 280, 281, 1, 0, 0, 0, 281, 283, 1, 0, 0, 0, 282, 274, 1, 0, 0, 0,
		282, 283, 1, 0, 0, 0, 283, 284, 1, 0, 0, 0, 284, 285, 3, 5, 2, 0, 285,
		286, 3, 11, 5, 0, 286, 287, 3, 41, 20, 0, 287, 288, 3, 47, 23, 0, 288,
		289, 3, 11, 5, 0, 289, 290, 3, 11, 5, 0, 290, 291, 3, 29, 14, 0, 291, 80,
		1, 0, 0, 0, 292, 295, 3, 83, 41, 0, 293, 295, 3, 85, 42, 0, 294, 292, 1,
		0, 0, 0, 294, 293, 1, 0, 0, 0, 295, 82, 1, 0, 0, 0, 296, 297, 3, 41, 20,
		0, 297, 298, 3, 37, 18, 0, 298, 299, 3, 43, 21, 0, 299, 300, 3, 11, 5,
		0, 300, 84, 1, 0, 0, 0, 301, 302, 3, 13, 6, 0, 302, 303, 3, 3, 1, 0, 303,
		304, 3, 25, 12, 0, 304, 305, 3, 39, 19, 0, 305, 306, 3, 11, 5, 0, 306,
		86, 1, 0, 0, 0, 307, 308, 5, 100, 0, 0, 308, 309, 5, 97, 0, 0, 309, 310,
		5, 116, 0, 0, 310, 311, 5, 101, 0, 0, 311, 312, 5, 116, 0, 0, 312, 313,
		5, 105, 0, 0, 313, 314, 5, 109, 0, 0, 314, 315, 5, 101, 0, 0, 315, 316,
		5, 40, 0, 0, 316, 320, 1, 0, 0, 0, 317, 319, 3, 55, 27, 0, 318, 317, 1,
		0, 0, 0, 319, 322, 1, 0, 0, 0, 320, 318, 1, 0, 0, 0, 320, 321, 1, 0, 0,
		0, 321, 323, 1, 0, 0, 0, 322, 320, 1, 0, 0, 0, 323, 327, 3, 157, 78, 0,
		324, 326, 3, 55, 27, 0, 325, 324, 1, 0, 0, 0, 326, 329, 1, 0, 0, 0, 327,
		325, 1, 0, 0, 0, 327, 328, 1, 0, 0, 0, 328, 330, 1, 0, 0, 0, 329, 327,
		1, 0, 0, 0, 330, 331, 5, 41, 0, 0, 331, 88, 1, 0, 0, 0, 332, 333, 3, 93,
		46, 0, 333, 334, 5, 45, 0, 0, 334, 335, 3, 95, 47, 0, 335, 336, 5, 45,
		0, 0, 336, 337, 3, 97, 48, 0, 337, 90, 1, 0, 0, 0, 338, 339, 3, 99, 49,
		0, 339, 340, 5, 58, 0, 0, 340, 341, 3, 101, 50, 0, 341, 342, 5, 58, 0,
		0, 342, 345, 3, 103, 51, 0, 343, 344, 5, 46, 0, 0, 344, 346, 3, 105, 52,
		0, 345, 343, 1, 0, 0, 0, 345, 346, 1, 0, 0, 0, 346, 347, 1, 0, 0, 0, 347,
		348, 3, 107, 53, 0, 348, 92, 1, 0, 0, 0, 349, 351, 3, 147, 73, 0, 350,
		349, 1, 0, 0, 0, 351, 352, 1, 0, 0, 0, 352, 350, 1, 0, 0, 0, 352, 353,
		1, 0, 0, 0, 353, 94, 1, 0, 0, 0, 354, 355, 5, 48, 0, 0, 355, 359, 7, 27,
		0, 0, 356, 357, 5, 49, 0, 0, 357, 359, 7, 28, 0, 0, 358, 354, 1, 0, 0,
		0, 358, 356, 1, 0, 0, 0, 359, 96, 1, 0, 0, 0, 360, 361, 5, 48, 0, 0, 361,
		367, 7, 27, 0, 0, 362, 363, 7, 29, 0, 0, 363, 367, 7, 30, 0, 0, 364, 365,
		5, 51, 0, 0, 365, 367, 7, 31, 0, 0, 366, 360, 1, 0, 0, 0, 366, 362, 1,
		0, 0, 0, 366, 364, 1, 0, 0, 0, 367, 98, 1, 0, 0, 0, 368, 369, 7, 31, 0,
		0, 369, 373, 7, 30, 0, 0, 370, 371, 5, 50, 0, 0, 371, 373, 7, 32, 0, 0,
		372, 368, 1, 0, 0, 0, 372, 370, 1, 0, 0, 0, 373, 100, 1, 0, 0, 0, 374,
		375, 7, 33, 0, 0, 375, 376, 7, 30, 0, 0, 376, 102, 1, 0, 0, 0, 377, 378,
		7, 33, 0, 0, 378, 382, 7, 30, 0, 0, 379, 380, 5, 54, 0, 0, 380, 382, 5,
		48, 0, 0, 381, 377, 1, 0, 0, 0, 381, 379, 1, 0, 0, 0, 382, 104, 1, 0, 0,
		0, 383, 385, 7, 30, 0, 0, 384, 383, 1, 0, 0, 0, 385, 386, 1, 0, 0, 0, 386,
		384, 1, 0, 0, 0, 386, 387, 1, 0, 0, 0, 387, 106, 1, 0, 0, 0, 388, 395,
		3, 53, 26, 0, 389, 390, 7, 34, 0, 0, 390, 391, 3, 109, 54, 0, 391, 392,
		5, 58, 0, 0, 392, 393, 3, 111, 55, 0, 393, 395, 1, 0, 0, 0, 394, 388, 1,
		0, 0, 0, 394, 389, 1, 0, 0, 0, 395, 108, 1, 0, 0, 0, 396, 397, 7, 31, 0,
		0, 397, 401, 7, 30, 0, 0, 398, 399, 5, 50, 0, 0, 399, 401, 7, 32, 0, 0,
		400, 396, 1, 0, 0, 0, 400, 398, 1, 0, 0, 0, 401, 110, 1, 0, 0, 0, 402,
		403, 7, 33, 0, 0, 403, 404, 7, 30, 0, 0, 404, 112, 1, 0, 0, 0, 405, 406,
		3, 3, 1, 0, 406, 407, 3, 25, 12, 0, 407, 408, 3, 25, 12, 0, 408, 409, 3,
		31, 15, 0, 409, 410, 3, 13, 6, 0, 410, 114, 1, 0, 0, 0, 411, 412, 3, 3,
		1, 0, 412, 413, 3, 29, 14, 0, 413, 414, 3, 51, 25, 0, 414, 415, 3, 31,
		15, 0, 415, 416, 3, 13, 6, 0, 416, 116, 1, 0, 0, 0, 417, 418, 3, 7, 3,
		0, 418, 419, 3, 31, 15, 0, 419, 420, 3, 43, 21, 0, 420, 421, 3, 29, 14,
		0, 421, 422, 3, 41, 20, 0, 422, 118, 1, 0, 0, 0, 423, 424, 3, 19, 9, 0,
		424, 425, 3, 39, 19, 0, 425, 426, 3, 11, 5, 0, 426, 427, 3, 27, 13, 0,
		427, 428, 3, 33, 16, 0, 428, 429, 3, 41, 20, 0, 429, 430, 3, 51, 25, 0,
		430, 120, 1, 0, 0, 0, 431, 436, 5, 34, 0, 0, 432, 435, 3, 151, 75, 0, 433,
		435, 3, 159, 79, 0, 434, 432, 1, 0, 0, 0, 434, 433, 1, 0, 0, 0, 435, 438,
		1, 0, 0, 0, 436, 434, 1, 0, 0, 0, 436, 437, 1, 0, 0, 0, 437, 439, 1, 0,
		0, 0, 438, 436, 1, 0, 0, 0, 439, 440, 5, 34, 0, 0, 440, 122, 1, 0, 0, 0,
		441, 443, 5, 45, 0, 0, 442, 441, 1, 0, 0, 0, 442, 443, 1, 0, 0, 0, 443,
		444, 1, 0, 0, 0, 444, 451, 3, 147, 73, 0, 445, 447, 5, 46, 0, 0, 446, 448,
		7, 30, 0, 0, 447, 446, 1, 0, 0, 0, 448, 449, 1, 0, 0, 0, 449, 447, 1, 0,
		0, 0, 449, 450, 1, 0, 0, 0, 450, 452, 1, 0, 0, 0, 451, 445, 1, 0, 0, 0,
		451, 452, 1, 0, 0, 0, 452, 454, 1, 0, 0, 0, 453, 455, 3, 149, 74, 0, 454,
		453, 1, 0, 0, 0, 454, 455, 1, 0, 0, 0, 455, 124, 1, 0, 0, 0, 456, 457,
		3, 29, 14, 0, 457, 458, 3, 43, 21, 0, 458, 459, 3, 25, 12, 0, 459, 460,
		3, 25, 12, 0, 460, 126, 1, 0, 0, 0, 461, 462, 3, 29, 14, 0, 462, 463, 3,
		31, 15, 0, 463, 464, 3, 41, 20, 0, 464, 128, 1, 0, 0, 0, 465, 466, 3, 3,
		1, 0, 466, 467, 3, 39, 19, 0, 467, 468, 3, 7, 3, 0, 468, 130, 1, 0, 0,
		0, 469, 470, 3, 9, 4, 0, 470, 471, 3, 11, 5, 0, 471, 472, 3, 39, 19, 0,
		472, 473, 3, 7, 3, 0, 473, 132, 1, 0, 0, 0, 474, 475, 3, 39, 19, 0, 475,
		476, 3, 31, 15, 0, 476, 477, 3, 37, 18, 0, 477, 478, 3, 41, 20, 0, 478,
		134, 1, 0, 0, 0, 479, 480, 3, 5, 2, 0, 480, 481, 3, 51, 25, 0, 481, 136,
		1, 0, 0, 0, 482, 483, 3, 39, 19, 0, 483, 484, 3, 23, 11, 0, 484, 485, 3,
		19, 9, 0, 485, 486, 3, 33, 16, 0, 486, 138, 1, 0, 0, 0, 487, 488, 3, 25,
		12, 0, 488, 489, 3, 19, 9, 0, 489, 490, 3, 27, 13, 0, 490, 491, 3, 19,
		9, 0, 491, 492, 3, 41, 20, 0, 492, 140, 1, 0, 0, 0, 493, 494, 3, 29, 14,
		0, 494, 495, 3, 31, 15, 0, 495, 496, 3, 29, 14, 0, 496, 497, 3, 11, 5,
		0, 497, 142, 1, 0, 0, 0, 498, 499, 3, 47, 23, 0, 499, 500, 3, 17, 8, 0,
		500, 501, 3, 11, 5, 0, 501, 502, 3, 37, 18, 0, 502, 503, 3, 11, 5, 0, 503,
		144, 1, 0, 0, 0, 504, 505, 3, 13, 6, 0, 505, 506, 3, 37, 18, 0, 506, 507,
		3, 31, 15, 0, 507, 508, 3, 27, 13, 0, 508, 146, 1, 0, 0, 0, 509, 518, 5,
		48, 0, 0, 510, 514, 7, 27, 0, 0, 511, 513, 7, 30, 0, 0, 512, 511, 1, 0,
		0, 0, 513, 516, 1, 0, 0, 0, 514, 512, 1, 0, 0, 0, 514, 515, 1, 0, 0, 0,
		515, 518, 1, 0, 0, 0, 516, 514, 1, 0, 0, 0, 517, 509, 1, 0, 0, 0, 517,
		510, 1, 0, 0, 0, 518, 148, 1, 0, 0, 0, 519, 521, 7, 4, 0, 0, 520, 522,
		7, 34, 0, 0, 521, 520, 1, 0, 0, 0, 521, 522, 1, 0, 0, 0, 522, 523, 1, 0,
		0, 0, 523, 524, 3, 147, 73, 0, 524, 150, 1, 0, 0, 0, 525, 526, 5, 92, 0,
		0, 526, 527, 7, 35, 0, 0, 527, 152, 1, 0, 0, 0, 528, 529, 5, 46, 0, 0,
		529, 154, 1, 0, 0, 0, 530, 534, 7, 36, 0, 0, 531, 533, 7, 37, 0, 0, 532,
		531, 1, 0, 0, 0, 533, 536, 1, 0, 0, 0, 534, 532, 1, 0, 0, 0, 534, 535,
		1, 0, 0, 0, 535, 547, 1, 0, 0, 0, 536, 534, 1, 0, 0, 0, 537, 538, 3, 153,
		76, 0, 538, 542, 7, 36, 0, 0, 539, 541, 7, 37, 0, 0, 540, 539, 1, 0, 0,
		0, 541, 544, 1, 0, 0, 0, 542, 540, 1, 0, 0, 0, 542, 543, 1, 0, 0, 0, 543,
		546, 1, 0, 0, 0, 544, 542, 1, 0, 0, 0, 545, 537, 1, 0, 0, 0, 546, 549,
		1, 0, 0, 0, 547, 545, 1, 0, 0, 0, 547, 548, 1, 0, 0, 0, 548, 573, 1, 0,
		0, 0, 549, 547, 1, 0, 0, 0, 550, 551, 5, 39, 0, 0, 551, 555, 7, 36, 0,
		0, 552, 554, 7, 37, 0, 0, 553, 552, 1, 0, 0, 0, 554, 557, 1, 0, 0, 0, 555,
		553, 1, 0, 0, 0, 555, 556, 1, 0, 0, 0, 556, 568, 1, 0, 0, 0, 557, 555,
		1, 0, 0, 0, 558, 559, 3, 153, 76, 0, 559, 563, 7, 36, 0, 0, 560, 562, 7,
		37, 0, 0, 561, 560, 1, 0, 0, 0, 562, 565, 1, 0, 0, 0, 563, 561, 1, 0, 0,
		0, 563, 564, 1, 0, 0, 0, 564, 567, 1, 0, 0, 0, 565, 563, 1, 0, 0, 0, 566,
		558, 1, 0, 0, 0, 567, 570, 1, 0, 0, 0, 568, 566, 1, 0, 0, 0, 568, 569,
		1, 0, 0, 0, 569, 571, 1, 0, 0, 0, 570, 568, 1, 0, 0, 0, 571, 573, 5, 39,
		0, 0, 572, 530, 1, 0, 0, 0, 572, 550, 1, 0, 0, 0, 573, 156, 1, 0, 0, 0,
		574, 575, 3, 89, 44, 0, 575, 576, 3, 41, 20, 0, 576, 577, 3, 91, 45, 0,
		577, 158, 1, 0, 0, 0, 578, 579, 8, 38, 0, 0, 579, 160, 1, 0, 0, 0, 37,
		0, 234, 238, 241, 251, 253, 269, 280, 282, 294, 320, 327, 345, 352, 358,
		366, 372, 381, 386, 394, 400, 434, 436, 442, 449, 451, 454, 514, 517, 521,
		534, 542, 547, 555, 563, 568, 572, 0,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// ZitiQlLexerInit initializes any static state used to implement ZitiQlLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewZitiQlLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func ZitiQlLexerInit() {
	staticData := &ZitiQlLexerLexerStaticData
	staticData.once.Do(zitiqllexerLexerInit)
}

// NewZitiQlLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewZitiQlLexer(input antlr.CharStream) *ZitiQlLexer {
	ZitiQlLexerInit()
	l := new(ZitiQlLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &ZitiQlLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "ZitiQl.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// ZitiQlLexer tokens.
const (
	ZitiQlLexerT__0              = 1
	ZitiQlLexerWS                = 2
	ZitiQlLexerLPAREN            = 3
	ZitiQlLexerRPAREN            = 4
	ZitiQlLexerLBRACKET          = 5
	ZitiQlLexerRBRACKET          = 6
	ZitiQlLexerAND               = 7
	ZitiQlLexerOR                = 8
	ZitiQlLexerLT                = 9
	ZitiQlLexerGT                = 10
	ZitiQlLexerEQ                = 11
	ZitiQlLexerCONTAINS          = 12
	ZitiQlLexerIN                = 13
	ZitiQlLexerBETWEEN           = 14
	ZitiQlLexerBOOL              = 15
	ZitiQlLexerDATETIME          = 16
	ZitiQlLexerALL_OF            = 17
	ZitiQlLexerANY_OF            = 18
	ZitiQlLexerCOUNT             = 19
	ZitiQlLexerISEMPTY           = 20
	ZitiQlLexerSTRING            = 21
	ZitiQlLexerNUMBER            = 22
	ZitiQlLexerNULL              = 23
	ZitiQlLexerNOT               = 24
	ZitiQlLexerASC               = 25
	ZitiQlLexerDESC              = 26
	ZitiQlLexerSORT              = 27
	ZitiQlLexerBY                = 28
	ZitiQlLexerSKIP_ROWS         = 29
	ZitiQlLexerLIMIT_ROWS        = 30
	ZitiQlLexerNONE              = 31
	ZitiQlLexerWHERE             = 32
	ZitiQlLexerFROM              = 33
	ZitiQlLexerIDENTIFIER        = 34
	ZitiQlLexerRFC3339_DATE_TIME = 35
)
