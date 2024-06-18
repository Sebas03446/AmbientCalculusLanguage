// Code generated from AmbientCalculus.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser

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

type AmbientCalculusLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var AmbientCalculusLexerLexerStaticData struct {
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

func ambientcalculuslexerLexerInit() {
	staticData := &AmbientCalculusLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'process'", "'{'", "'}'", "'move'", "';'", "'ambient'", "'to'",
		"'send'", "'receive'", "'in'", "'out'", "'open'", "'call'", "'print'",
		"'printc'", "'modifyc'", "'+='", "'-='", "'conditions'", "'let'", "'='",
		"'letc'", "'restriction'", "'('", "')'", "'+'", "'-'", "'*'", "'/'",
		"'>'", "'<'", "'>='", "'<='", "'=='", "'!='",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "AMBIENTID", "ID", "INT", "STRING", "WS", "COMMENT",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16",
		"T__17", "T__18", "T__19", "T__20", "T__21", "T__22", "T__23", "T__24",
		"T__25", "T__26", "T__27", "T__28", "T__29", "T__30", "T__31", "T__32",
		"T__33", "T__34", "AMBIENTID", "ID", "INT", "STRING", "WS", "COMMENT",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 41, 278, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 1, 0, 1,
		0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1,
		3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1,
		5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1,
		8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10,
		1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1,
		12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14,
		1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1,
		15, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18,
		1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1,
		19, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22,
		1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 23, 1,
		23, 1, 24, 1, 24, 1, 25, 1, 25, 1, 26, 1, 26, 1, 27, 1, 27, 1, 28, 1, 28,
		1, 29, 1, 29, 1, 30, 1, 30, 1, 31, 1, 31, 1, 31, 1, 32, 1, 32, 1, 32, 1,
		33, 1, 33, 1, 33, 1, 34, 1, 34, 1, 34, 1, 35, 1, 35, 5, 35, 235, 8, 35,
		10, 35, 12, 35, 238, 9, 35, 1, 36, 1, 36, 5, 36, 242, 8, 36, 10, 36, 12,
		36, 245, 9, 36, 1, 37, 4, 37, 248, 8, 37, 11, 37, 12, 37, 249, 1, 38, 1,
		38, 5, 38, 254, 8, 38, 10, 38, 12, 38, 257, 9, 38, 1, 38, 1, 38, 1, 39,
		4, 39, 262, 8, 39, 11, 39, 12, 39, 263, 1, 39, 1, 39, 1, 40, 1, 40, 1,
		40, 1, 40, 5, 40, 272, 8, 40, 10, 40, 12, 40, 275, 9, 40, 1, 40, 1, 40,
		0, 0, 41, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19,
		10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37,
		19, 39, 20, 41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51, 26, 53, 27, 55,
		28, 57, 29, 59, 30, 61, 31, 63, 32, 65, 33, 67, 34, 69, 35, 71, 36, 73,
		37, 75, 38, 77, 39, 79, 40, 81, 41, 1, 0, 8, 1, 0, 97, 122, 3, 0, 48, 57,
		95, 95, 97, 122, 1, 0, 65, 90, 3, 0, 48, 57, 65, 90, 95, 95, 1, 0, 48,
		57, 3, 0, 10, 10, 13, 13, 34, 34, 3, 0, 9, 10, 13, 13, 32, 32, 2, 0, 10,
		10, 13, 13, 283, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0,
		0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0,
		0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0,
		0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1,
		0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37,
		1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0,
		45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0,
		0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0,
		0, 0, 61, 1, 0, 0, 0, 0, 63, 1, 0, 0, 0, 0, 65, 1, 0, 0, 0, 0, 67, 1, 0,
		0, 0, 0, 69, 1, 0, 0, 0, 0, 71, 1, 0, 0, 0, 0, 73, 1, 0, 0, 0, 0, 75, 1,
		0, 0, 0, 0, 77, 1, 0, 0, 0, 0, 79, 1, 0, 0, 0, 0, 81, 1, 0, 0, 0, 1, 83,
		1, 0, 0, 0, 3, 91, 1, 0, 0, 0, 5, 93, 1, 0, 0, 0, 7, 95, 1, 0, 0, 0, 9,
		100, 1, 0, 0, 0, 11, 102, 1, 0, 0, 0, 13, 110, 1, 0, 0, 0, 15, 113, 1,
		0, 0, 0, 17, 118, 1, 0, 0, 0, 19, 126, 1, 0, 0, 0, 21, 129, 1, 0, 0, 0,
		23, 133, 1, 0, 0, 0, 25, 138, 1, 0, 0, 0, 27, 143, 1, 0, 0, 0, 29, 149,
		1, 0, 0, 0, 31, 156, 1, 0, 0, 0, 33, 164, 1, 0, 0, 0, 35, 167, 1, 0, 0,
		0, 37, 170, 1, 0, 0, 0, 39, 181, 1, 0, 0, 0, 41, 185, 1, 0, 0, 0, 43, 187,
		1, 0, 0, 0, 45, 192, 1, 0, 0, 0, 47, 204, 1, 0, 0, 0, 49, 206, 1, 0, 0,
		0, 51, 208, 1, 0, 0, 0, 53, 210, 1, 0, 0, 0, 55, 212, 1, 0, 0, 0, 57, 214,
		1, 0, 0, 0, 59, 216, 1, 0, 0, 0, 61, 218, 1, 0, 0, 0, 63, 220, 1, 0, 0,
		0, 65, 223, 1, 0, 0, 0, 67, 226, 1, 0, 0, 0, 69, 229, 1, 0, 0, 0, 71, 232,
		1, 0, 0, 0, 73, 239, 1, 0, 0, 0, 75, 247, 1, 0, 0, 0, 77, 251, 1, 0, 0,
		0, 79, 261, 1, 0, 0, 0, 81, 267, 1, 0, 0, 0, 83, 84, 5, 112, 0, 0, 84,
		85, 5, 114, 0, 0, 85, 86, 5, 111, 0, 0, 86, 87, 5, 99, 0, 0, 87, 88, 5,
		101, 0, 0, 88, 89, 5, 115, 0, 0, 89, 90, 5, 115, 0, 0, 90, 2, 1, 0, 0,
		0, 91, 92, 5, 123, 0, 0, 92, 4, 1, 0, 0, 0, 93, 94, 5, 125, 0, 0, 94, 6,
		1, 0, 0, 0, 95, 96, 5, 109, 0, 0, 96, 97, 5, 111, 0, 0, 97, 98, 5, 118,
		0, 0, 98, 99, 5, 101, 0, 0, 99, 8, 1, 0, 0, 0, 100, 101, 5, 59, 0, 0, 101,
		10, 1, 0, 0, 0, 102, 103, 5, 97, 0, 0, 103, 104, 5, 109, 0, 0, 104, 105,
		5, 98, 0, 0, 105, 106, 5, 105, 0, 0, 106, 107, 5, 101, 0, 0, 107, 108,
		5, 110, 0, 0, 108, 109, 5, 116, 0, 0, 109, 12, 1, 0, 0, 0, 110, 111, 5,
		116, 0, 0, 111, 112, 5, 111, 0, 0, 112, 14, 1, 0, 0, 0, 113, 114, 5, 115,
		0, 0, 114, 115, 5, 101, 0, 0, 115, 116, 5, 110, 0, 0, 116, 117, 5, 100,
		0, 0, 117, 16, 1, 0, 0, 0, 118, 119, 5, 114, 0, 0, 119, 120, 5, 101, 0,
		0, 120, 121, 5, 99, 0, 0, 121, 122, 5, 101, 0, 0, 122, 123, 5, 105, 0,
		0, 123, 124, 5, 118, 0, 0, 124, 125, 5, 101, 0, 0, 125, 18, 1, 0, 0, 0,
		126, 127, 5, 105, 0, 0, 127, 128, 5, 110, 0, 0, 128, 20, 1, 0, 0, 0, 129,
		130, 5, 111, 0, 0, 130, 131, 5, 117, 0, 0, 131, 132, 5, 116, 0, 0, 132,
		22, 1, 0, 0, 0, 133, 134, 5, 111, 0, 0, 134, 135, 5, 112, 0, 0, 135, 136,
		5, 101, 0, 0, 136, 137, 5, 110, 0, 0, 137, 24, 1, 0, 0, 0, 138, 139, 5,
		99, 0, 0, 139, 140, 5, 97, 0, 0, 140, 141, 5, 108, 0, 0, 141, 142, 5, 108,
		0, 0, 142, 26, 1, 0, 0, 0, 143, 144, 5, 112, 0, 0, 144, 145, 5, 114, 0,
		0, 145, 146, 5, 105, 0, 0, 146, 147, 5, 110, 0, 0, 147, 148, 5, 116, 0,
		0, 148, 28, 1, 0, 0, 0, 149, 150, 5, 112, 0, 0, 150, 151, 5, 114, 0, 0,
		151, 152, 5, 105, 0, 0, 152, 153, 5, 110, 0, 0, 153, 154, 5, 116, 0, 0,
		154, 155, 5, 99, 0, 0, 155, 30, 1, 0, 0, 0, 156, 157, 5, 109, 0, 0, 157,
		158, 5, 111, 0, 0, 158, 159, 5, 100, 0, 0, 159, 160, 5, 105, 0, 0, 160,
		161, 5, 102, 0, 0, 161, 162, 5, 121, 0, 0, 162, 163, 5, 99, 0, 0, 163,
		32, 1, 0, 0, 0, 164, 165, 5, 43, 0, 0, 165, 166, 5, 61, 0, 0, 166, 34,
		1, 0, 0, 0, 167, 168, 5, 45, 0, 0, 168, 169, 5, 61, 0, 0, 169, 36, 1, 0,
		0, 0, 170, 171, 5, 99, 0, 0, 171, 172, 5, 111, 0, 0, 172, 173, 5, 110,
		0, 0, 173, 174, 5, 100, 0, 0, 174, 175, 5, 105, 0, 0, 175, 176, 5, 116,
		0, 0, 176, 177, 5, 105, 0, 0, 177, 178, 5, 111, 0, 0, 178, 179, 5, 110,
		0, 0, 179, 180, 5, 115, 0, 0, 180, 38, 1, 0, 0, 0, 181, 182, 5, 108, 0,
		0, 182, 183, 5, 101, 0, 0, 183, 184, 5, 116, 0, 0, 184, 40, 1, 0, 0, 0,
		185, 186, 5, 61, 0, 0, 186, 42, 1, 0, 0, 0, 187, 188, 5, 108, 0, 0, 188,
		189, 5, 101, 0, 0, 189, 190, 5, 116, 0, 0, 190, 191, 5, 99, 0, 0, 191,
		44, 1, 0, 0, 0, 192, 193, 5, 114, 0, 0, 193, 194, 5, 101, 0, 0, 194, 195,
		5, 115, 0, 0, 195, 196, 5, 116, 0, 0, 196, 197, 5, 114, 0, 0, 197, 198,
		5, 105, 0, 0, 198, 199, 5, 99, 0, 0, 199, 200, 5, 116, 0, 0, 200, 201,
		5, 105, 0, 0, 201, 202, 5, 111, 0, 0, 202, 203, 5, 110, 0, 0, 203, 46,
		1, 0, 0, 0, 204, 205, 5, 40, 0, 0, 205, 48, 1, 0, 0, 0, 206, 207, 5, 41,
		0, 0, 207, 50, 1, 0, 0, 0, 208, 209, 5, 43, 0, 0, 209, 52, 1, 0, 0, 0,
		210, 211, 5, 45, 0, 0, 211, 54, 1, 0, 0, 0, 212, 213, 5, 42, 0, 0, 213,
		56, 1, 0, 0, 0, 214, 215, 5, 47, 0, 0, 215, 58, 1, 0, 0, 0, 216, 217, 5,
		62, 0, 0, 217, 60, 1, 0, 0, 0, 218, 219, 5, 60, 0, 0, 219, 62, 1, 0, 0,
		0, 220, 221, 5, 62, 0, 0, 221, 222, 5, 61, 0, 0, 222, 64, 1, 0, 0, 0, 223,
		224, 5, 60, 0, 0, 224, 225, 5, 61, 0, 0, 225, 66, 1, 0, 0, 0, 226, 227,
		5, 61, 0, 0, 227, 228, 5, 61, 0, 0, 228, 68, 1, 0, 0, 0, 229, 230, 5, 33,
		0, 0, 230, 231, 5, 61, 0, 0, 231, 70, 1, 0, 0, 0, 232, 236, 7, 0, 0, 0,
		233, 235, 7, 1, 0, 0, 234, 233, 1, 0, 0, 0, 235, 238, 1, 0, 0, 0, 236,
		234, 1, 0, 0, 0, 236, 237, 1, 0, 0, 0, 237, 72, 1, 0, 0, 0, 238, 236, 1,
		0, 0, 0, 239, 243, 7, 2, 0, 0, 240, 242, 7, 3, 0, 0, 241, 240, 1, 0, 0,
		0, 242, 245, 1, 0, 0, 0, 243, 241, 1, 0, 0, 0, 243, 244, 1, 0, 0, 0, 244,
		74, 1, 0, 0, 0, 245, 243, 1, 0, 0, 0, 246, 248, 7, 4, 0, 0, 247, 246, 1,
		0, 0, 0, 248, 249, 1, 0, 0, 0, 249, 247, 1, 0, 0, 0, 249, 250, 1, 0, 0,
		0, 250, 76, 1, 0, 0, 0, 251, 255, 5, 34, 0, 0, 252, 254, 8, 5, 0, 0, 253,
		252, 1, 0, 0, 0, 254, 257, 1, 0, 0, 0, 255, 253, 1, 0, 0, 0, 255, 256,
		1, 0, 0, 0, 256, 258, 1, 0, 0, 0, 257, 255, 1, 0, 0, 0, 258, 259, 5, 34,
		0, 0, 259, 78, 1, 0, 0, 0, 260, 262, 7, 6, 0, 0, 261, 260, 1, 0, 0, 0,
		262, 263, 1, 0, 0, 0, 263, 261, 1, 0, 0, 0, 263, 264, 1, 0, 0, 0, 264,
		265, 1, 0, 0, 0, 265, 266, 6, 39, 0, 0, 266, 80, 1, 0, 0, 0, 267, 268,
		5, 47, 0, 0, 268, 269, 5, 47, 0, 0, 269, 273, 1, 0, 0, 0, 270, 272, 8,
		7, 0, 0, 271, 270, 1, 0, 0, 0, 272, 275, 1, 0, 0, 0, 273, 271, 1, 0, 0,
		0, 273, 274, 1, 0, 0, 0, 274, 276, 1, 0, 0, 0, 275, 273, 1, 0, 0, 0, 276,
		277, 6, 40, 0, 0, 277, 82, 1, 0, 0, 0, 7, 0, 236, 243, 249, 255, 263, 273,
		1, 6, 0, 0,
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

// AmbientCalculusLexerInit initializes any static state used to implement AmbientCalculusLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewAmbientCalculusLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func AmbientCalculusLexerInit() {
	staticData := &AmbientCalculusLexerLexerStaticData
	staticData.once.Do(ambientcalculuslexerLexerInit)
}

// NewAmbientCalculusLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewAmbientCalculusLexer(input antlr.CharStream) *AmbientCalculusLexer {
	AmbientCalculusLexerInit()
	l := new(AmbientCalculusLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &AmbientCalculusLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "AmbientCalculus.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// AmbientCalculusLexer tokens.
const (
	AmbientCalculusLexerT__0      = 1
	AmbientCalculusLexerT__1      = 2
	AmbientCalculusLexerT__2      = 3
	AmbientCalculusLexerT__3      = 4
	AmbientCalculusLexerT__4      = 5
	AmbientCalculusLexerT__5      = 6
	AmbientCalculusLexerT__6      = 7
	AmbientCalculusLexerT__7      = 8
	AmbientCalculusLexerT__8      = 9
	AmbientCalculusLexerT__9      = 10
	AmbientCalculusLexerT__10     = 11
	AmbientCalculusLexerT__11     = 12
	AmbientCalculusLexerT__12     = 13
	AmbientCalculusLexerT__13     = 14
	AmbientCalculusLexerT__14     = 15
	AmbientCalculusLexerT__15     = 16
	AmbientCalculusLexerT__16     = 17
	AmbientCalculusLexerT__17     = 18
	AmbientCalculusLexerT__18     = 19
	AmbientCalculusLexerT__19     = 20
	AmbientCalculusLexerT__20     = 21
	AmbientCalculusLexerT__21     = 22
	AmbientCalculusLexerT__22     = 23
	AmbientCalculusLexerT__23     = 24
	AmbientCalculusLexerT__24     = 25
	AmbientCalculusLexerT__25     = 26
	AmbientCalculusLexerT__26     = 27
	AmbientCalculusLexerT__27     = 28
	AmbientCalculusLexerT__28     = 29
	AmbientCalculusLexerT__29     = 30
	AmbientCalculusLexerT__30     = 31
	AmbientCalculusLexerT__31     = 32
	AmbientCalculusLexerT__32     = 33
	AmbientCalculusLexerT__33     = 34
	AmbientCalculusLexerT__34     = 35
	AmbientCalculusLexerAMBIENTID = 36
	AmbientCalculusLexerID        = 37
	AmbientCalculusLexerINT       = 38
	AmbientCalculusLexerSTRING    = 39
	AmbientCalculusLexerWS        = 40
	AmbientCalculusLexerCOMMENT   = 41
)
