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
		"", "'process'", "'{'", "'}'", "'move'", "';'", "'ambient'", "'conditions'",
		"'to'", "'send'", "'receive'", "'from'", "'in'", "'out'", "'open'",
		"'call'", "'print'", "'if'", "'('", "')'", "'let'", "'='", "'*'", "'/'",
		"'+'", "'-'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "AMBIENTID", "ID", "INT", "STRING",
		"WS", "COMMENT",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16",
		"T__17", "T__18", "T__19", "T__20", "T__21", "T__22", "T__23", "T__24",
		"AMBIENTID", "ID", "INT", "STRING", "WS", "COMMENT",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 31, 212, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 1,
		0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1,
		3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1,
		5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1,
		6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1,
		9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1,
		11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13,
		1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1,
		15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 18, 1, 18, 1, 19, 1, 19,
		1, 19, 1, 19, 1, 20, 1, 20, 1, 21, 1, 21, 1, 22, 1, 22, 1, 23, 1, 23, 1,
		24, 1, 24, 1, 25, 1, 25, 5, 25, 169, 8, 25, 10, 25, 12, 25, 172, 9, 25,
		1, 26, 1, 26, 5, 26, 176, 8, 26, 10, 26, 12, 26, 179, 9, 26, 1, 27, 4,
		27, 182, 8, 27, 11, 27, 12, 27, 183, 1, 28, 1, 28, 5, 28, 188, 8, 28, 10,
		28, 12, 28, 191, 9, 28, 1, 28, 1, 28, 1, 29, 4, 29, 196, 8, 29, 11, 29,
		12, 29, 197, 1, 29, 1, 29, 1, 30, 1, 30, 1, 30, 1, 30, 5, 30, 206, 8, 30,
		10, 30, 12, 30, 209, 9, 30, 1, 30, 1, 30, 0, 0, 31, 1, 1, 3, 2, 5, 3, 7,
		4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27,
		14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45,
		23, 47, 24, 49, 25, 51, 26, 53, 27, 55, 28, 57, 29, 59, 30, 61, 31, 1,
		0, 8, 1, 0, 97, 122, 3, 0, 48, 57, 95, 95, 97, 122, 1, 0, 65, 90, 3, 0,
		48, 57, 65, 90, 95, 95, 1, 0, 48, 57, 3, 0, 10, 10, 13, 13, 34, 34, 3,
		0, 9, 10, 13, 13, 32, 32, 2, 0, 10, 10, 13, 13, 217, 0, 1, 1, 0, 0, 0,
		0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0,
		0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0,
		0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0,
		0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1,
		0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41,
		1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0,
		49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0,
		0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 1, 63, 1, 0, 0,
		0, 3, 71, 1, 0, 0, 0, 5, 73, 1, 0, 0, 0, 7, 75, 1, 0, 0, 0, 9, 80, 1, 0,
		0, 0, 11, 82, 1, 0, 0, 0, 13, 90, 1, 0, 0, 0, 15, 101, 1, 0, 0, 0, 17,
		104, 1, 0, 0, 0, 19, 109, 1, 0, 0, 0, 21, 117, 1, 0, 0, 0, 23, 122, 1,
		0, 0, 0, 25, 125, 1, 0, 0, 0, 27, 129, 1, 0, 0, 0, 29, 134, 1, 0, 0, 0,
		31, 139, 1, 0, 0, 0, 33, 145, 1, 0, 0, 0, 35, 148, 1, 0, 0, 0, 37, 150,
		1, 0, 0, 0, 39, 152, 1, 0, 0, 0, 41, 156, 1, 0, 0, 0, 43, 158, 1, 0, 0,
		0, 45, 160, 1, 0, 0, 0, 47, 162, 1, 0, 0, 0, 49, 164, 1, 0, 0, 0, 51, 166,
		1, 0, 0, 0, 53, 173, 1, 0, 0, 0, 55, 181, 1, 0, 0, 0, 57, 185, 1, 0, 0,
		0, 59, 195, 1, 0, 0, 0, 61, 201, 1, 0, 0, 0, 63, 64, 5, 112, 0, 0, 64,
		65, 5, 114, 0, 0, 65, 66, 5, 111, 0, 0, 66, 67, 5, 99, 0, 0, 67, 68, 5,
		101, 0, 0, 68, 69, 5, 115, 0, 0, 69, 70, 5, 115, 0, 0, 70, 2, 1, 0, 0,
		0, 71, 72, 5, 123, 0, 0, 72, 4, 1, 0, 0, 0, 73, 74, 5, 125, 0, 0, 74, 6,
		1, 0, 0, 0, 75, 76, 5, 109, 0, 0, 76, 77, 5, 111, 0, 0, 77, 78, 5, 118,
		0, 0, 78, 79, 5, 101, 0, 0, 79, 8, 1, 0, 0, 0, 80, 81, 5, 59, 0, 0, 81,
		10, 1, 0, 0, 0, 82, 83, 5, 97, 0, 0, 83, 84, 5, 109, 0, 0, 84, 85, 5, 98,
		0, 0, 85, 86, 5, 105, 0, 0, 86, 87, 5, 101, 0, 0, 87, 88, 5, 110, 0, 0,
		88, 89, 5, 116, 0, 0, 89, 12, 1, 0, 0, 0, 90, 91, 5, 99, 0, 0, 91, 92,
		5, 111, 0, 0, 92, 93, 5, 110, 0, 0, 93, 94, 5, 100, 0, 0, 94, 95, 5, 105,
		0, 0, 95, 96, 5, 116, 0, 0, 96, 97, 5, 105, 0, 0, 97, 98, 5, 111, 0, 0,
		98, 99, 5, 110, 0, 0, 99, 100, 5, 115, 0, 0, 100, 14, 1, 0, 0, 0, 101,
		102, 5, 116, 0, 0, 102, 103, 5, 111, 0, 0, 103, 16, 1, 0, 0, 0, 104, 105,
		5, 115, 0, 0, 105, 106, 5, 101, 0, 0, 106, 107, 5, 110, 0, 0, 107, 108,
		5, 100, 0, 0, 108, 18, 1, 0, 0, 0, 109, 110, 5, 114, 0, 0, 110, 111, 5,
		101, 0, 0, 111, 112, 5, 99, 0, 0, 112, 113, 5, 101, 0, 0, 113, 114, 5,
		105, 0, 0, 114, 115, 5, 118, 0, 0, 115, 116, 5, 101, 0, 0, 116, 20, 1,
		0, 0, 0, 117, 118, 5, 102, 0, 0, 118, 119, 5, 114, 0, 0, 119, 120, 5, 111,
		0, 0, 120, 121, 5, 109, 0, 0, 121, 22, 1, 0, 0, 0, 122, 123, 5, 105, 0,
		0, 123, 124, 5, 110, 0, 0, 124, 24, 1, 0, 0, 0, 125, 126, 5, 111, 0, 0,
		126, 127, 5, 117, 0, 0, 127, 128, 5, 116, 0, 0, 128, 26, 1, 0, 0, 0, 129,
		130, 5, 111, 0, 0, 130, 131, 5, 112, 0, 0, 131, 132, 5, 101, 0, 0, 132,
		133, 5, 110, 0, 0, 133, 28, 1, 0, 0, 0, 134, 135, 5, 99, 0, 0, 135, 136,
		5, 97, 0, 0, 136, 137, 5, 108, 0, 0, 137, 138, 5, 108, 0, 0, 138, 30, 1,
		0, 0, 0, 139, 140, 5, 112, 0, 0, 140, 141, 5, 114, 0, 0, 141, 142, 5, 105,
		0, 0, 142, 143, 5, 110, 0, 0, 143, 144, 5, 116, 0, 0, 144, 32, 1, 0, 0,
		0, 145, 146, 5, 105, 0, 0, 146, 147, 5, 102, 0, 0, 147, 34, 1, 0, 0, 0,
		148, 149, 5, 40, 0, 0, 149, 36, 1, 0, 0, 0, 150, 151, 5, 41, 0, 0, 151,
		38, 1, 0, 0, 0, 152, 153, 5, 108, 0, 0, 153, 154, 5, 101, 0, 0, 154, 155,
		5, 116, 0, 0, 155, 40, 1, 0, 0, 0, 156, 157, 5, 61, 0, 0, 157, 42, 1, 0,
		0, 0, 158, 159, 5, 42, 0, 0, 159, 44, 1, 0, 0, 0, 160, 161, 5, 47, 0, 0,
		161, 46, 1, 0, 0, 0, 162, 163, 5, 43, 0, 0, 163, 48, 1, 0, 0, 0, 164, 165,
		5, 45, 0, 0, 165, 50, 1, 0, 0, 0, 166, 170, 7, 0, 0, 0, 167, 169, 7, 1,
		0, 0, 168, 167, 1, 0, 0, 0, 169, 172, 1, 0, 0, 0, 170, 168, 1, 0, 0, 0,
		170, 171, 1, 0, 0, 0, 171, 52, 1, 0, 0, 0, 172, 170, 1, 0, 0, 0, 173, 177,
		7, 2, 0, 0, 174, 176, 7, 3, 0, 0, 175, 174, 1, 0, 0, 0, 176, 179, 1, 0,
		0, 0, 177, 175, 1, 0, 0, 0, 177, 178, 1, 0, 0, 0, 178, 54, 1, 0, 0, 0,
		179, 177, 1, 0, 0, 0, 180, 182, 7, 4, 0, 0, 181, 180, 1, 0, 0, 0, 182,
		183, 1, 0, 0, 0, 183, 181, 1, 0, 0, 0, 183, 184, 1, 0, 0, 0, 184, 56, 1,
		0, 0, 0, 185, 189, 5, 34, 0, 0, 186, 188, 8, 5, 0, 0, 187, 186, 1, 0, 0,
		0, 188, 191, 1, 0, 0, 0, 189, 187, 1, 0, 0, 0, 189, 190, 1, 0, 0, 0, 190,
		192, 1, 0, 0, 0, 191, 189, 1, 0, 0, 0, 192, 193, 5, 34, 0, 0, 193, 58,
		1, 0, 0, 0, 194, 196, 7, 6, 0, 0, 195, 194, 1, 0, 0, 0, 196, 197, 1, 0,
		0, 0, 197, 195, 1, 0, 0, 0, 197, 198, 1, 0, 0, 0, 198, 199, 1, 0, 0, 0,
		199, 200, 6, 29, 0, 0, 200, 60, 1, 0, 0, 0, 201, 202, 5, 47, 0, 0, 202,
		203, 5, 47, 0, 0, 203, 207, 1, 0, 0, 0, 204, 206, 8, 7, 0, 0, 205, 204,
		1, 0, 0, 0, 206, 209, 1, 0, 0, 0, 207, 205, 1, 0, 0, 0, 207, 208, 1, 0,
		0, 0, 208, 210, 1, 0, 0, 0, 209, 207, 1, 0, 0, 0, 210, 211, 6, 30, 0, 0,
		211, 62, 1, 0, 0, 0, 7, 0, 170, 177, 183, 189, 197, 207, 1, 6, 0, 0,
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
	AmbientCalculusLexerAMBIENTID = 26
	AmbientCalculusLexerID        = 27
	AmbientCalculusLexerINT       = 28
	AmbientCalculusLexerSTRING    = 29
	AmbientCalculusLexerWS        = 30
	AmbientCalculusLexerCOMMENT   = 31
)
