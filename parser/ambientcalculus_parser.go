// Code generated from AmbientCalculus.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // AmbientCalculus

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type AmbientCalculusParser struct {
	*antlr.BaseParser
}

var AmbientCalculusParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func ambientcalculusParserInit() {
	staticData := &AmbientCalculusParserStaticData
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
		"program", "statement", "processDeclaration", "processStatement", "moveStatement",
		"ambientDeclaration", "conditions", "movementStatement", "sendStatement",
		"receiveStatement", "inStatement", "outStatement", "openStatement",
		"callStatement", "printStatement", "printConditionStatement", "modifyConditionStatement",
		"conditionsBlock", "variableDeclaration", "conditionsVariableDeclaration",
		"restriction", "assignmentStatement", "expression", "operator", "comparator",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 41, 211, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 1, 0, 5, 0, 52, 8,
		0, 10, 0, 12, 0, 55, 9, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 72, 8, 1, 1, 2, 1, 2,
		1, 2, 1, 2, 5, 2, 78, 8, 2, 10, 2, 12, 2, 81, 9, 2, 1, 2, 1, 2, 1, 3, 1,
		3, 1, 3, 3, 3, 88, 8, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5,
		3, 5, 98, 8, 5, 1, 5, 5, 5, 101, 8, 5, 10, 5, 12, 5, 104, 9, 5, 1, 5, 1,
		5, 1, 6, 1, 6, 1, 6, 5, 6, 111, 8, 6, 10, 6, 12, 6, 114, 9, 6, 1, 6, 1,
		6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1,
		8, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1,
		12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14,
		1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1,
		16, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19,
		1, 19, 1, 19, 1, 19, 3, 19, 175, 8, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1,
		20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 22,
		1, 22, 1, 22, 1, 22, 1, 22, 3, 22, 196, 8, 22, 1, 22, 1, 22, 1, 22, 1,
		22, 5, 22, 202, 8, 22, 10, 22, 12, 22, 205, 9, 22, 1, 23, 1, 23, 1, 24,
		1, 24, 1, 24, 0, 1, 44, 25, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22,
		24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 0, 3, 1, 0, 17, 18,
		1, 0, 26, 29, 1, 0, 30, 35, 209, 0, 53, 1, 0, 0, 0, 2, 71, 1, 0, 0, 0,
		4, 73, 1, 0, 0, 0, 6, 87, 1, 0, 0, 0, 8, 89, 1, 0, 0, 0, 10, 93, 1, 0,
		0, 0, 12, 107, 1, 0, 0, 0, 14, 117, 1, 0, 0, 0, 16, 123, 1, 0, 0, 0, 18,
		129, 1, 0, 0, 0, 20, 132, 1, 0, 0, 0, 22, 136, 1, 0, 0, 0, 24, 139, 1,
		0, 0, 0, 26, 143, 1, 0, 0, 0, 28, 147, 1, 0, 0, 0, 30, 151, 1, 0, 0, 0,
		32, 155, 1, 0, 0, 0, 34, 161, 1, 0, 0, 0, 36, 163, 1, 0, 0, 0, 38, 169,
		1, 0, 0, 0, 40, 178, 1, 0, 0, 0, 42, 182, 1, 0, 0, 0, 44, 195, 1, 0, 0,
		0, 46, 206, 1, 0, 0, 0, 48, 208, 1, 0, 0, 0, 50, 52, 3, 2, 1, 0, 51, 50,
		1, 0, 0, 0, 52, 55, 1, 0, 0, 0, 53, 51, 1, 0, 0, 0, 53, 54, 1, 0, 0, 0,
		54, 56, 1, 0, 0, 0, 55, 53, 1, 0, 0, 0, 56, 57, 5, 0, 0, 1, 57, 1, 1, 0,
		0, 0, 58, 72, 3, 4, 2, 0, 59, 72, 3, 10, 5, 0, 60, 72, 3, 14, 7, 0, 61,
		72, 3, 16, 8, 0, 62, 72, 3, 18, 9, 0, 63, 72, 3, 20, 10, 0, 64, 72, 3,
		22, 11, 0, 65, 72, 3, 24, 12, 0, 66, 72, 3, 26, 13, 0, 67, 72, 3, 28, 14,
		0, 68, 72, 3, 36, 18, 0, 69, 72, 3, 42, 21, 0, 70, 72, 3, 32, 16, 0, 71,
		58, 1, 0, 0, 0, 71, 59, 1, 0, 0, 0, 71, 60, 1, 0, 0, 0, 71, 61, 1, 0, 0,
		0, 71, 62, 1, 0, 0, 0, 71, 63, 1, 0, 0, 0, 71, 64, 1, 0, 0, 0, 71, 65,
		1, 0, 0, 0, 71, 66, 1, 0, 0, 0, 71, 67, 1, 0, 0, 0, 71, 68, 1, 0, 0, 0,
		71, 69, 1, 0, 0, 0, 71, 70, 1, 0, 0, 0, 72, 3, 1, 0, 0, 0, 73, 74, 5, 1,
		0, 0, 74, 75, 5, 37, 0, 0, 75, 79, 5, 2, 0, 0, 76, 78, 3, 6, 3, 0, 77,
		76, 1, 0, 0, 0, 78, 81, 1, 0, 0, 0, 79, 77, 1, 0, 0, 0, 79, 80, 1, 0, 0,
		0, 80, 82, 1, 0, 0, 0, 81, 79, 1, 0, 0, 0, 82, 83, 5, 3, 0, 0, 83, 5, 1,
		0, 0, 0, 84, 88, 3, 2, 1, 0, 85, 88, 3, 8, 4, 0, 86, 88, 3, 30, 15, 0,
		87, 84, 1, 0, 0, 0, 87, 85, 1, 0, 0, 0, 87, 86, 1, 0, 0, 0, 88, 7, 1, 0,
		0, 0, 89, 90, 5, 4, 0, 0, 90, 91, 5, 36, 0, 0, 91, 92, 5, 5, 0, 0, 92,
		9, 1, 0, 0, 0, 93, 94, 5, 6, 0, 0, 94, 95, 5, 36, 0, 0, 95, 97, 5, 2, 0,
		0, 96, 98, 3, 12, 6, 0, 97, 96, 1, 0, 0, 0, 97, 98, 1, 0, 0, 0, 98, 102,
		1, 0, 0, 0, 99, 101, 3, 2, 1, 0, 100, 99, 1, 0, 0, 0, 101, 104, 1, 0, 0,
		0, 102, 100, 1, 0, 0, 0, 102, 103, 1, 0, 0, 0, 103, 105, 1, 0, 0, 0, 104,
		102, 1, 0, 0, 0, 105, 106, 5, 3, 0, 0, 106, 11, 1, 0, 0, 0, 107, 108, 3,
		34, 17, 0, 108, 112, 5, 2, 0, 0, 109, 111, 3, 38, 19, 0, 110, 109, 1, 0,
		0, 0, 111, 114, 1, 0, 0, 0, 112, 110, 1, 0, 0, 0, 112, 113, 1, 0, 0, 0,
		113, 115, 1, 0, 0, 0, 114, 112, 1, 0, 0, 0, 115, 116, 5, 3, 0, 0, 116,
		13, 1, 0, 0, 0, 117, 118, 5, 4, 0, 0, 118, 119, 5, 37, 0, 0, 119, 120,
		5, 7, 0, 0, 120, 121, 5, 36, 0, 0, 121, 122, 5, 5, 0, 0, 122, 15, 1, 0,
		0, 0, 123, 124, 5, 8, 0, 0, 124, 125, 5, 39, 0, 0, 125, 126, 5, 7, 0, 0,
		126, 127, 5, 37, 0, 0, 127, 128, 5, 5, 0, 0, 128, 17, 1, 0, 0, 0, 129,
		130, 5, 9, 0, 0, 130, 131, 5, 5, 0, 0, 131, 19, 1, 0, 0, 0, 132, 133, 5,
		10, 0, 0, 133, 134, 5, 36, 0, 0, 134, 135, 5, 5, 0, 0, 135, 21, 1, 0, 0,
		0, 136, 137, 5, 11, 0, 0, 137, 138, 5, 5, 0, 0, 138, 23, 1, 0, 0, 0, 139,
		140, 5, 12, 0, 0, 140, 141, 5, 36, 0, 0, 141, 142, 5, 5, 0, 0, 142, 25,
		1, 0, 0, 0, 143, 144, 5, 13, 0, 0, 144, 145, 5, 37, 0, 0, 145, 146, 5,
		5, 0, 0, 146, 27, 1, 0, 0, 0, 147, 148, 5, 14, 0, 0, 148, 149, 3, 44, 22,
		0, 149, 150, 5, 5, 0, 0, 150, 29, 1, 0, 0, 0, 151, 152, 5, 15, 0, 0, 152,
		153, 5, 37, 0, 0, 153, 154, 5, 5, 0, 0, 154, 31, 1, 0, 0, 0, 155, 156,
		5, 16, 0, 0, 156, 157, 5, 37, 0, 0, 157, 158, 7, 0, 0, 0, 158, 159, 3,
		44, 22, 0, 159, 160, 5, 5, 0, 0, 160, 33, 1, 0, 0, 0, 161, 162, 5, 19,
		0, 0, 162, 35, 1, 0, 0, 0, 163, 164, 5, 20, 0, 0, 164, 165, 5, 37, 0, 0,
		165, 166, 5, 21, 0, 0, 166, 167, 3, 44, 22, 0, 167, 168, 5, 5, 0, 0, 168,
		37, 1, 0, 0, 0, 169, 170, 5, 22, 0, 0, 170, 171, 5, 37, 0, 0, 171, 172,
		5, 21, 0, 0, 172, 174, 5, 38, 0, 0, 173, 175, 3, 40, 20, 0, 174, 173, 1,
		0, 0, 0, 174, 175, 1, 0, 0, 0, 175, 176, 1, 0, 0, 0, 176, 177, 5, 5, 0,
		0, 177, 39, 1, 0, 0, 0, 178, 179, 5, 23, 0, 0, 179, 180, 3, 48, 24, 0,
		180, 181, 5, 38, 0, 0, 181, 41, 1, 0, 0, 0, 182, 183, 5, 37, 0, 0, 183,
		184, 5, 21, 0, 0, 184, 185, 3, 44, 22, 0, 185, 186, 5, 5, 0, 0, 186, 43,
		1, 0, 0, 0, 187, 188, 6, 22, -1, 0, 188, 189, 5, 24, 0, 0, 189, 190, 3,
		44, 22, 0, 190, 191, 5, 25, 0, 0, 191, 196, 1, 0, 0, 0, 192, 196, 5, 37,
		0, 0, 193, 196, 5, 38, 0, 0, 194, 196, 5, 39, 0, 0, 195, 187, 1, 0, 0,
		0, 195, 192, 1, 0, 0, 0, 195, 193, 1, 0, 0, 0, 195, 194, 1, 0, 0, 0, 196,
		203, 1, 0, 0, 0, 197, 198, 10, 5, 0, 0, 198, 199, 3, 46, 23, 0, 199, 200,
		3, 44, 22, 6, 200, 202, 1, 0, 0, 0, 201, 197, 1, 0, 0, 0, 202, 205, 1,
		0, 0, 0, 203, 201, 1, 0, 0, 0, 203, 204, 1, 0, 0, 0, 204, 45, 1, 0, 0,
		0, 205, 203, 1, 0, 0, 0, 206, 207, 7, 1, 0, 0, 207, 47, 1, 0, 0, 0, 208,
		209, 7, 2, 0, 0, 209, 49, 1, 0, 0, 0, 10, 53, 71, 79, 87, 97, 102, 112,
		174, 195, 203,
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

// AmbientCalculusParserInit initializes any static state used to implement AmbientCalculusParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewAmbientCalculusParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func AmbientCalculusParserInit() {
	staticData := &AmbientCalculusParserStaticData
	staticData.once.Do(ambientcalculusParserInit)
}

// NewAmbientCalculusParser produces a new parser instance for the optional input antlr.TokenStream.
func NewAmbientCalculusParser(input antlr.TokenStream) *AmbientCalculusParser {
	AmbientCalculusParserInit()
	this := new(AmbientCalculusParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &AmbientCalculusParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "AmbientCalculus.g4"

	return this
}

// AmbientCalculusParser tokens.
const (
	AmbientCalculusParserEOF       = antlr.TokenEOF
	AmbientCalculusParserT__0      = 1
	AmbientCalculusParserT__1      = 2
	AmbientCalculusParserT__2      = 3
	AmbientCalculusParserT__3      = 4
	AmbientCalculusParserT__4      = 5
	AmbientCalculusParserT__5      = 6
	AmbientCalculusParserT__6      = 7
	AmbientCalculusParserT__7      = 8
	AmbientCalculusParserT__8      = 9
	AmbientCalculusParserT__9      = 10
	AmbientCalculusParserT__10     = 11
	AmbientCalculusParserT__11     = 12
	AmbientCalculusParserT__12     = 13
	AmbientCalculusParserT__13     = 14
	AmbientCalculusParserT__14     = 15
	AmbientCalculusParserT__15     = 16
	AmbientCalculusParserT__16     = 17
	AmbientCalculusParserT__17     = 18
	AmbientCalculusParserT__18     = 19
	AmbientCalculusParserT__19     = 20
	AmbientCalculusParserT__20     = 21
	AmbientCalculusParserT__21     = 22
	AmbientCalculusParserT__22     = 23
	AmbientCalculusParserT__23     = 24
	AmbientCalculusParserT__24     = 25
	AmbientCalculusParserT__25     = 26
	AmbientCalculusParserT__26     = 27
	AmbientCalculusParserT__27     = 28
	AmbientCalculusParserT__28     = 29
	AmbientCalculusParserT__29     = 30
	AmbientCalculusParserT__30     = 31
	AmbientCalculusParserT__31     = 32
	AmbientCalculusParserT__32     = 33
	AmbientCalculusParserT__33     = 34
	AmbientCalculusParserT__34     = 35
	AmbientCalculusParserAMBIENTID = 36
	AmbientCalculusParserID        = 37
	AmbientCalculusParserINT       = 38
	AmbientCalculusParserSTRING    = 39
	AmbientCalculusParserWS        = 40
	AmbientCalculusParserCOMMENT   = 41
)

// AmbientCalculusParser rules.
const (
	AmbientCalculusParserRULE_program                       = 0
	AmbientCalculusParserRULE_statement                     = 1
	AmbientCalculusParserRULE_processDeclaration            = 2
	AmbientCalculusParserRULE_processStatement              = 3
	AmbientCalculusParserRULE_moveStatement                 = 4
	AmbientCalculusParserRULE_ambientDeclaration            = 5
	AmbientCalculusParserRULE_conditions                    = 6
	AmbientCalculusParserRULE_movementStatement             = 7
	AmbientCalculusParserRULE_sendStatement                 = 8
	AmbientCalculusParserRULE_receiveStatement              = 9
	AmbientCalculusParserRULE_inStatement                   = 10
	AmbientCalculusParserRULE_outStatement                  = 11
	AmbientCalculusParserRULE_openStatement                 = 12
	AmbientCalculusParserRULE_callStatement                 = 13
	AmbientCalculusParserRULE_printStatement                = 14
	AmbientCalculusParserRULE_printConditionStatement       = 15
	AmbientCalculusParserRULE_modifyConditionStatement      = 16
	AmbientCalculusParserRULE_conditionsBlock               = 17
	AmbientCalculusParserRULE_variableDeclaration           = 18
	AmbientCalculusParserRULE_conditionsVariableDeclaration = 19
	AmbientCalculusParserRULE_restriction                   = 20
	AmbientCalculusParserRULE_assignmentStatement           = 21
	AmbientCalculusParserRULE_expression                    = 22
	AmbientCalculusParserRULE_operator                      = 23
	AmbientCalculusParserRULE_comparator                    = 24
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AmbientCalculusParserRULE_program
	return p
}

func InitEmptyProgramContext(p *ProgramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AmbientCalculusParserRULE_program
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AmbientCalculusParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) EOF() antlr.TerminalNode {
	return s.GetToken(AmbientCalculusParserEOF, 0)
}

func (s *ProgramContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AmbientCalculusListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AmbientCalculusListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (s *ProgramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AmbientCalculusVisitor:
		return t.VisitProgram(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AmbientCalculusParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, AmbientCalculusParserRULE_program)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(53)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&137440100178) != 0 {
		{
			p.SetState(50)
			p.Statement()
		}

		p.SetState(55)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(56)
		p.Match(AmbientCalculusParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ProcessDeclaration() IProcessDeclarationContext
	AmbientDeclaration() IAmbientDeclarationContext
	MovementStatement() IMovementStatementContext
	SendStatement() ISendStatementContext
	ReceiveStatement() IReceiveStatementContext
	InStatement() IInStatementContext
	OutStatement() IOutStatementContext
	OpenStatement() IOpenStatementContext
	CallStatement() ICallStatementContext
	PrintStatement() IPrintStatementContext
	VariableDeclaration() IVariableDeclarationContext
	AssignmentStatement() IAssignmentStatementContext
	ModifyConditionStatement() IModifyConditionStatementContext

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AmbientCalculusParserRULE_statement
	return p
}

func InitEmptyStatementContext(p *StatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AmbientCalculusParserRULE_statement
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AmbientCalculusParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) ProcessDeclaration() IProcessDeclarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IProcessDeclarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IProcessDeclarationContext)
}

func (s *StatementContext) AmbientDeclaration() IAmbientDeclarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAmbientDeclarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAmbientDeclarationContext)
}

func (s *StatementContext) MovementStatement() IMovementStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMovementStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMovementStatementContext)
}

func (s *StatementContext) SendStatement() ISendStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISendStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISendStatementContext)
}

func (s *StatementContext) ReceiveStatement() IReceiveStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReceiveStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReceiveStatementContext)
}

func (s *StatementContext) InStatement() IInStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInStatementContext)
}

func (s *StatementContext) OutStatement() IOutStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOutStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOutStatementContext)
}

func (s *StatementContext) OpenStatement() IOpenStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOpenStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOpenStatementContext)
}

func (s *StatementContext) CallStatement() ICallStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICallStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICallStatementContext)
}

func (s *StatementContext) PrintStatement() IPrintStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrintStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrintStatementContext)
}

func (s *StatementContext) VariableDeclaration() IVariableDeclarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableDeclarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableDeclarationContext)
}

func (s *StatementContext) AssignmentStatement() IAssignmentStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignmentStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignmentStatementContext)
}

func (s *StatementContext) ModifyConditionStatement() IModifyConditionStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IModifyConditionStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IModifyConditionStatementContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AmbientCalculusListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AmbientCalculusListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (s *StatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AmbientCalculusVisitor:
		return t.VisitStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AmbientCalculusParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, AmbientCalculusParserRULE_statement)
	p.SetState(71)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case AmbientCalculusParserT__0:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(58)
			p.ProcessDeclaration()
		}

	case AmbientCalculusParserT__5:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(59)
			p.AmbientDeclaration()
		}

	case AmbientCalculusParserT__3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(60)
			p.MovementStatement()
		}

	case AmbientCalculusParserT__7:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(61)
			p.SendStatement()
		}

	case AmbientCalculusParserT__8:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(62)
			p.ReceiveStatement()
		}

	case AmbientCalculusParserT__9:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(63)
			p.InStatement()
		}

	case AmbientCalculusParserT__10:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(64)
			p.OutStatement()
		}

	case AmbientCalculusParserT__11:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(65)
			p.OpenStatement()
		}

	case AmbientCalculusParserT__12:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(66)
			p.CallStatement()
		}

	case AmbientCalculusParserT__13:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(67)
			p.PrintStatement()
		}

	case AmbientCalculusParserT__19:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(68)
			p.VariableDeclaration()
		}

	case AmbientCalculusParserID:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(69)
			p.AssignmentStatement()
		}

	case AmbientCalculusParserT__15:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(70)
			p.ModifyConditionStatement()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IProcessDeclarationContext is an interface to support dynamic dispatch.
type IProcessDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	AllProcessStatement() []IProcessStatementContext
	ProcessStatement(i int) IProcessStatementContext

	// IsProcessDeclarationContext differentiates from other interfaces.
	IsProcessDeclarationContext()
}

type ProcessDeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProcessDeclarationContext() *ProcessDeclarationContext {
	var p = new(ProcessDeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AmbientCalculusParserRULE_processDeclaration
	return p
}

func InitEmptyProcessDeclarationContext(p *ProcessDeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AmbientCalculusParserRULE_processDeclaration
}

func (*ProcessDeclarationContext) IsProcessDeclarationContext() {}

func NewProcessDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProcessDeclarationContext {
	var p = new(ProcessDeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AmbientCalculusParserRULE_processDeclaration

	return p
}

func (s *ProcessDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *ProcessDeclarationContext) ID() antlr.TerminalNode {
	return s.GetToken(AmbientCalculusParserID, 0)
}

func (s *ProcessDeclarationContext) AllProcessStatement() []IProcessStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IProcessStatementContext); ok {
			len++
		}
	}

	tst := make([]IProcessStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IProcessStatementContext); ok {
			tst[i] = t.(IProcessStatementContext)
			i++
		}
	}

	return tst
}

func (s *ProcessDeclarationContext) ProcessStatement(i int) IProcessStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IProcessStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IProcessStatementContext)
}

func (s *ProcessDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProcessDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProcessDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AmbientCalculusListener); ok {
		listenerT.EnterProcessDeclaration(s)
	}
}

func (s *ProcessDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AmbientCalculusListener); ok {
		listenerT.ExitProcessDeclaration(s)
	}
}

func (s *ProcessDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AmbientCalculusVisitor:
		return t.VisitProcessDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AmbientCalculusParser) ProcessDeclaration() (localctx IProcessDeclarationContext) {
	localctx = NewProcessDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, AmbientCalculusParserRULE_processDeclaration)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(73)
		p.Match(AmbientCalculusParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(74)
		p.Match(AmbientCalculusParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(75)
		p.Match(AmbientCalculusParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(79)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&137440132946) != 0 {
		{
			p.SetState(76)
			p.ProcessStatement()
		}

		p.SetState(81)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(82)
		p.Match(AmbientCalculusParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IProcessStatementContext is an interface to support dynamic dispatch.
type IProcessStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Statement() IStatementContext
	MoveStatement() IMoveStatementContext
	PrintConditionStatement() IPrintConditionStatementContext

	// IsProcessStatementContext differentiates from other interfaces.
	IsProcessStatementContext()
}

type ProcessStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProcessStatementContext() *ProcessStatementContext {
	var p = new(ProcessStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AmbientCalculusParserRULE_processStatement
	return p
}

func InitEmptyProcessStatementContext(p *ProcessStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AmbientCalculusParserRULE_processStatement
}

func (*ProcessStatementContext) IsProcessStatementContext() {}

func NewProcessStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProcessStatementContext {
	var p = new(ProcessStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AmbientCalculusParserRULE_processStatement

	return p
}

func (s *ProcessStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ProcessStatementContext) Statement() IStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *ProcessStatementContext) MoveStatement() IMoveStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMoveStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMoveStatementContext)
}

func (s *ProcessStatementContext) PrintConditionStatement() IPrintConditionStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrintConditionStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrintConditionStatementContext)
}

func (s *ProcessStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProcessStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProcessStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AmbientCalculusListener); ok {
		listenerT.EnterProcessStatement(s)
	}
}

func (s *ProcessStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AmbientCalculusListener); ok {
		listenerT.ExitProcessStatement(s)
	}
}

func (s *ProcessStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AmbientCalculusVisitor:
		return t.VisitProcessStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AmbientCalculusParser) ProcessStatement() (localctx IProcessStatementContext) {
	localctx = NewProcessStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, AmbientCalculusParserRULE_processStatement)
	p.SetState(87)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(84)
			p.Statement()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(85)
			p.MoveStatement()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(86)
			p.PrintConditionStatement()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMoveStatementContext is an interface to support dynamic dispatch.
type IMoveStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AMBIENTID() antlr.TerminalNode

	// IsMoveStatementContext differentiates from other interfaces.
	IsMoveStatementContext()
}

type MoveStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMoveStatementContext() *MoveStatementContext {
	var p = new(MoveStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AmbientCalculusParserRULE_moveStatement
	return p
}

func InitEmptyMoveStatementContext(p *MoveStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AmbientCalculusParserRULE_moveStatement
}

func (*MoveStatementContext) IsMoveStatementContext() {}

func NewMoveStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MoveStatementContext {
	var p = new(MoveStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AmbientCalculusParserRULE_moveStatement

	return p
}

func (s *MoveStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *MoveStatementContext) AMBIENTID() antlr.TerminalNode {
	return s.GetToken(AmbientCalculusParserAMBIENTID, 0)
}

func (s *MoveStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MoveStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MoveStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AmbientCalculusListener); ok {
		listenerT.EnterMoveStatement(s)
	}
}

func (s *MoveStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AmbientCalculusListener); ok {
		listenerT.ExitMoveStatement(s)
	}
}

func (s *MoveStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AmbientCalculusVisitor:
		return t.VisitMoveStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AmbientCalculusParser) MoveStatement() (localctx IMoveStatementContext) {
	localctx = NewMoveStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, AmbientCalculusParserRULE_moveStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(89)
		p.Match(AmbientCalculusParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(90)
		p.Match(AmbientCalculusParserAMBIENTID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(91)
		p.Match(AmbientCalculusParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAmbientDeclarationContext is an interface to support dynamic dispatch.
type IAmbientDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AMBIENTID() antlr.TerminalNode
	Conditions() IConditionsContext
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

	// IsAmbientDeclarationContext differentiates from other interfaces.
	IsAmbientDeclarationContext()
}

type AmbientDeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAmbientDeclarationContext() *AmbientDeclarationContext {
	var p = new(AmbientDeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AmbientCalculusParserRULE_ambientDeclaration
	return p
}

func InitEmptyAmbientDeclarationContext(p *AmbientDeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AmbientCalculusParserRULE_ambientDeclaration
}

func (*AmbientDeclarationContext) IsAmbientDeclarationContext() {}

func NewAmbientDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AmbientDeclarationContext {
	var p = new(AmbientDeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AmbientCalculusParserRULE_ambientDeclaration

	return p
}

func (s *AmbientDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *AmbientDeclarationContext) AMBIENTID() antlr.TerminalNode {
	return s.GetToken(AmbientCalculusParserAMBIENTID, 0)
}

func (s *AmbientDeclarationContext) Conditions() IConditionsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConditionsContext)
}

func (s *AmbientDeclarationContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *AmbientDeclarationContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *AmbientDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AmbientDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AmbientDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AmbientCalculusListener); ok {
		listenerT.EnterAmbientDeclaration(s)
	}
}

func (s *AmbientDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AmbientCalculusListener); ok {
		listenerT.ExitAmbientDeclaration(s)
	}
}

func (s *AmbientDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AmbientCalculusVisitor:
		return t.VisitAmbientDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AmbientCalculusParser) AmbientDeclaration() (localctx IAmbientDeclarationContext) {
	localctx = NewAmbientDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, AmbientCalculusParserRULE_ambientDeclaration)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(93)
		p.Match(AmbientCalculusParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(94)
		p.Match(AmbientCalculusParserAMBIENTID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(95)
		p.Match(AmbientCalculusParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(97)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == AmbientCalculusParserT__18 {
		{
			p.SetState(96)
			p.Conditions()
		}

	}
	p.SetState(102)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&137440100178) != 0 {
		{
			p.SetState(99)
			p.Statement()
		}

		p.SetState(104)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(105)
		p.Match(AmbientCalculusParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IConditionsContext is an interface to support dynamic dispatch.
type IConditionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ConditionsBlock() IConditionsBlockContext
	AllConditionsVariableDeclaration() []IConditionsVariableDeclarationContext
	ConditionsVariableDeclaration(i int) IConditionsVariableDeclarationContext

	// IsConditionsContext differentiates from other interfaces.
	IsConditionsContext()
}

type ConditionsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConditionsContext() *ConditionsContext {
	var p = new(ConditionsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AmbientCalculusParserRULE_conditions
	return p
}

func InitEmptyConditionsContext(p *ConditionsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AmbientCalculusParserRULE_conditions
}

func (*ConditionsContext) IsConditionsContext() {}

func NewConditionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionsContext {
	var p = new(ConditionsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AmbientCalculusParserRULE_conditions

	return p
}

func (s *ConditionsContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionsContext) ConditionsBlock() IConditionsBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionsBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConditionsBlockContext)
}

func (s *ConditionsContext) AllConditionsVariableDeclaration() []IConditionsVariableDeclarationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IConditionsVariableDeclarationContext); ok {
			len++
		}
	}

	tst := make([]IConditionsVariableDeclarationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IConditionsVariableDeclarationContext); ok {
			tst[i] = t.(IConditionsVariableDeclarationContext)
			i++
		}
	}

	return tst
}

func (s *ConditionsContext) ConditionsVariableDeclaration(i int) IConditionsVariableDeclarationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionsVariableDeclarationContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConditionsVariableDeclarationContext)
}

func (s *ConditionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConditionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AmbientCalculusListener); ok {
		listenerT.EnterConditions(s)
	}
}

func (s *ConditionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AmbientCalculusListener); ok {
		listenerT.ExitConditions(s)
	}
}

func (s *ConditionsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AmbientCalculusVisitor:
		return t.VisitConditions(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AmbientCalculusParser) Conditions() (localctx IConditionsContext) {
	localctx = NewConditionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, AmbientCalculusParserRULE_conditions)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(107)
		p.ConditionsBlock()
	}
	{
		p.SetState(108)
		p.Match(AmbientCalculusParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(112)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == AmbientCalculusParserT__21 {
		{
			p.SetState(109)
			p.ConditionsVariableDeclaration()
		}

		p.SetState(114)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(115)
		p.Match(AmbientCalculusParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMovementStatementContext is an interface to support dynamic dispatch.
type IMovementStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	AMBIENTID() antlr.TerminalNode

	// IsMovementStatementContext differentiates from other interfaces.
	IsMovementStatementContext()
}

type MovementStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMovementStatementContext() *MovementStatementContext {
	var p = new(MovementStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AmbientCalculusParserRULE_movementStatement
	return p
}

func InitEmptyMovementStatementContext(p *MovementStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AmbientCalculusParserRULE_movementStatement
}

func (*MovementStatementContext) IsMovementStatementContext() {}

func NewMovementStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MovementStatementContext {
	var p = new(MovementStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AmbientCalculusParserRULE_movementStatement

	return p
}

func (s *MovementStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *MovementStatementContext) ID() antlr.TerminalNode {
	return s.GetToken(AmbientCalculusParserID, 0)
}

func (s *MovementStatementContext) AMBIENTID() antlr.TerminalNode {
	return s.GetToken(AmbientCalculusParserAMBIENTID, 0)
}

func (s *MovementStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MovementStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MovementStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AmbientCalculusListener); ok {
		listenerT.EnterMovementStatement(s)
	}
}

func (s *MovementStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AmbientCalculusListener); ok {
		listenerT.ExitMovementStatement(s)
	}
}

func (s *MovementStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AmbientCalculusVisitor:
		return t.VisitMovementStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AmbientCalculusParser) MovementStatement() (localctx IMovementStatementContext) {
	localctx = NewMovementStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, AmbientCalculusParserRULE_movementStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(117)
		p.Match(AmbientCalculusParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(118)
		p.Match(AmbientCalculusParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(119)
		p.Match(AmbientCalculusParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(120)
		p.Match(AmbientCalculusParserAMBIENTID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(121)
		p.Match(AmbientCalculusParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISendStatementContext is an interface to support dynamic dispatch.
type ISendStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STRING() antlr.TerminalNode
	ID() antlr.TerminalNode

	// IsSendStatementContext differentiates from other interfaces.
	IsSendStatementContext()
}

type SendStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySendStatementContext() *SendStatementContext {
	var p = new(SendStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AmbientCalculusParserRULE_sendStatement
	return p
}

func InitEmptySendStatementContext(p *SendStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AmbientCalculusParserRULE_sendStatement
}

func (*SendStatementContext) IsSendStatementContext() {}

func NewSendStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SendStatementContext {
	var p = new(SendStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AmbientCalculusParserRULE_sendStatement

	return p
}

func (s *SendStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *SendStatementContext) STRING() antlr.TerminalNode {
	return s.GetToken(AmbientCalculusParserSTRING, 0)
}

func (s *SendStatementContext) ID() antlr.TerminalNode {
	return s.GetToken(AmbientCalculusParserID, 0)
}

func (s *SendStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SendStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SendStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AmbientCalculusListener); ok {
		listenerT.EnterSendStatement(s)
	}
}

func (s *SendStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AmbientCalculusListener); ok {
		listenerT.ExitSendStatement(s)
	}
}

func (s *SendStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AmbientCalculusVisitor:
		return t.VisitSendStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AmbientCalculusParser) SendStatement() (localctx ISendStatementContext) {
	localctx = NewSendStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, AmbientCalculusParserRULE_sendStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(123)
		p.Match(AmbientCalculusParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(124)
		p.Match(AmbientCalculusParserSTRING)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(125)
		p.Match(AmbientCalculusParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(126)
		p.Match(AmbientCalculusParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(127)
		p.Match(AmbientCalculusParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IReceiveStatementContext is an interface to support dynamic dispatch.
type IReceiveStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsReceiveStatementContext differentiates from other interfaces.
	IsReceiveStatementContext()
}

type ReceiveStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReceiveStatementContext() *ReceiveStatementContext {
	var p = new(ReceiveStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AmbientCalculusParserRULE_receiveStatement
	return p
}

func InitEmptyReceiveStatementContext(p *ReceiveStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AmbientCalculusParserRULE_receiveStatement
}

func (*ReceiveStatementContext) IsReceiveStatementContext() {}

func NewReceiveStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReceiveStatementContext {
	var p = new(ReceiveStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AmbientCalculusParserRULE_receiveStatement

	return p
}

func (s *ReceiveStatementContext) GetParser() antlr.Parser { return s.parser }
func (s *ReceiveStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReceiveStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReceiveStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AmbientCalculusListener); ok {
		listenerT.EnterReceiveStatement(s)
	}
}

func (s *ReceiveStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AmbientCalculusListener); ok {
		listenerT.ExitReceiveStatement(s)
	}
}

func (s *ReceiveStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AmbientCalculusVisitor:
		return t.VisitReceiveStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AmbientCalculusParser) ReceiveStatement() (localctx IReceiveStatementContext) {
	localctx = NewReceiveStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, AmbientCalculusParserRULE_receiveStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(129)
		p.Match(AmbientCalculusParserT__8)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(130)
		p.Match(AmbientCalculusParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInStatementContext is an interface to support dynamic dispatch.
type IInStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AMBIENTID() antlr.TerminalNode

	// IsInStatementContext differentiates from other interfaces.
	IsInStatementContext()
}

type InStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInStatementContext() *InStatementContext {
	var p = new(InStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AmbientCalculusParserRULE_inStatement
	return p
}

func InitEmptyInStatementContext(p *InStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AmbientCalculusParserRULE_inStatement
}

func (*InStatementContext) IsInStatementContext() {}

func NewInStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InStatementContext {
	var p = new(InStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AmbientCalculusParserRULE_inStatement

	return p
}

func (s *InStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *InStatementContext) AMBIENTID() antlr.TerminalNode {
	return s.GetToken(AmbientCalculusParserAMBIENTID, 0)
}

func (s *InStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AmbientCalculusListener); ok {
		listenerT.EnterInStatement(s)
	}
}

func (s *InStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AmbientCalculusListener); ok {
		listenerT.ExitInStatement(s)
	}
}

func (s *InStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AmbientCalculusVisitor:
		return t.VisitInStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AmbientCalculusParser) InStatement() (localctx IInStatementContext) {
	localctx = NewInStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, AmbientCalculusParserRULE_inStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(132)
		p.Match(AmbientCalculusParserT__9)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(133)
		p.Match(AmbientCalculusParserAMBIENTID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(134)
		p.Match(AmbientCalculusParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOutStatementContext is an interface to support dynamic dispatch.
type IOutStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsOutStatementContext differentiates from other interfaces.
	IsOutStatementContext()
}

type OutStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOutStatementContext() *OutStatementContext {
	var p = new(OutStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AmbientCalculusParserRULE_outStatement
	return p
}

func InitEmptyOutStatementContext(p *OutStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AmbientCalculusParserRULE_outStatement
}

func (*OutStatementContext) IsOutStatementContext() {}

func NewOutStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OutStatementContext {
	var p = new(OutStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AmbientCalculusParserRULE_outStatement

	return p
}

func (s *OutStatementContext) GetParser() antlr.Parser { return s.parser }
func (s *OutStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OutStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OutStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AmbientCalculusListener); ok {
		listenerT.EnterOutStatement(s)
	}
}

func (s *OutStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AmbientCalculusListener); ok {
		listenerT.ExitOutStatement(s)
	}
}

func (s *OutStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AmbientCalculusVisitor:
		return t.VisitOutStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AmbientCalculusParser) OutStatement() (localctx IOutStatementContext) {
	localctx = NewOutStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, AmbientCalculusParserRULE_outStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(136)
		p.Match(AmbientCalculusParserT__10)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(137)
		p.Match(AmbientCalculusParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOpenStatementContext is an interface to support dynamic dispatch.
type IOpenStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AMBIENTID() antlr.TerminalNode

	// IsOpenStatementContext differentiates from other interfaces.
	IsOpenStatementContext()
}

type OpenStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOpenStatementContext() *OpenStatementContext {
	var p = new(OpenStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AmbientCalculusParserRULE_openStatement
	return p
}

func InitEmptyOpenStatementContext(p *OpenStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AmbientCalculusParserRULE_openStatement
}

func (*OpenStatementContext) IsOpenStatementContext() {}

func NewOpenStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OpenStatementContext {
	var p = new(OpenStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AmbientCalculusParserRULE_openStatement

	return p
}

func (s *OpenStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *OpenStatementContext) AMBIENTID() antlr.TerminalNode {
	return s.GetToken(AmbientCalculusParserAMBIENTID, 0)
}

func (s *OpenStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OpenStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OpenStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AmbientCalculusListener); ok {
		listenerT.EnterOpenStatement(s)
	}
}

func (s *OpenStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AmbientCalculusListener); ok {
		listenerT.ExitOpenStatement(s)
	}
}

func (s *OpenStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AmbientCalculusVisitor:
		return t.VisitOpenStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AmbientCalculusParser) OpenStatement() (localctx IOpenStatementContext) {
	localctx = NewOpenStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, AmbientCalculusParserRULE_openStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(139)
		p.Match(AmbientCalculusParserT__11)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(140)
		p.Match(AmbientCalculusParserAMBIENTID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(141)
		p.Match(AmbientCalculusParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICallStatementContext is an interface to support dynamic dispatch.
type ICallStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode

	// IsCallStatementContext differentiates from other interfaces.
	IsCallStatementContext()
}

type CallStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCallStatementContext() *CallStatementContext {
	var p = new(CallStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AmbientCalculusParserRULE_callStatement
	return p
}

func InitEmptyCallStatementContext(p *CallStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AmbientCalculusParserRULE_callStatement
}

func (*CallStatementContext) IsCallStatementContext() {}

func NewCallStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CallStatementContext {
	var p = new(CallStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AmbientCalculusParserRULE_callStatement

	return p
}

func (s *CallStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *CallStatementContext) ID() antlr.TerminalNode {
	return s.GetToken(AmbientCalculusParserID, 0)
}

func (s *CallStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CallStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CallStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AmbientCalculusListener); ok {
		listenerT.EnterCallStatement(s)
	}
}

func (s *CallStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AmbientCalculusListener); ok {
		listenerT.ExitCallStatement(s)
	}
}

func (s *CallStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AmbientCalculusVisitor:
		return t.VisitCallStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AmbientCalculusParser) CallStatement() (localctx ICallStatementContext) {
	localctx = NewCallStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, AmbientCalculusParserRULE_callStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(143)
		p.Match(AmbientCalculusParserT__12)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(144)
		p.Match(AmbientCalculusParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(145)
		p.Match(AmbientCalculusParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPrintStatementContext is an interface to support dynamic dispatch.
type IPrintStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext

	// IsPrintStatementContext differentiates from other interfaces.
	IsPrintStatementContext()
}

type PrintStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrintStatementContext() *PrintStatementContext {
	var p = new(PrintStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AmbientCalculusParserRULE_printStatement
	return p
}

func InitEmptyPrintStatementContext(p *PrintStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AmbientCalculusParserRULE_printStatement
}

func (*PrintStatementContext) IsPrintStatementContext() {}

func NewPrintStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrintStatementContext {
	var p = new(PrintStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AmbientCalculusParserRULE_printStatement

	return p
}

func (s *PrintStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *PrintStatementContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *PrintStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrintStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrintStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AmbientCalculusListener); ok {
		listenerT.EnterPrintStatement(s)
	}
}

func (s *PrintStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AmbientCalculusListener); ok {
		listenerT.ExitPrintStatement(s)
	}
}

func (s *PrintStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AmbientCalculusVisitor:
		return t.VisitPrintStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AmbientCalculusParser) PrintStatement() (localctx IPrintStatementContext) {
	localctx = NewPrintStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, AmbientCalculusParserRULE_printStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(147)
		p.Match(AmbientCalculusParserT__13)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(148)
		p.expression(0)
	}
	{
		p.SetState(149)
		p.Match(AmbientCalculusParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPrintConditionStatementContext is an interface to support dynamic dispatch.
type IPrintConditionStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode

	// IsPrintConditionStatementContext differentiates from other interfaces.
	IsPrintConditionStatementContext()
}

type PrintConditionStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrintConditionStatementContext() *PrintConditionStatementContext {
	var p = new(PrintConditionStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AmbientCalculusParserRULE_printConditionStatement
	return p
}

func InitEmptyPrintConditionStatementContext(p *PrintConditionStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AmbientCalculusParserRULE_printConditionStatement
}

func (*PrintConditionStatementContext) IsPrintConditionStatementContext() {}

func NewPrintConditionStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrintConditionStatementContext {
	var p = new(PrintConditionStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AmbientCalculusParserRULE_printConditionStatement

	return p
}

func (s *PrintConditionStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *PrintConditionStatementContext) ID() antlr.TerminalNode {
	return s.GetToken(AmbientCalculusParserID, 0)
}

func (s *PrintConditionStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrintConditionStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrintConditionStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AmbientCalculusListener); ok {
		listenerT.EnterPrintConditionStatement(s)
	}
}

func (s *PrintConditionStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AmbientCalculusListener); ok {
		listenerT.ExitPrintConditionStatement(s)
	}
}

func (s *PrintConditionStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AmbientCalculusVisitor:
		return t.VisitPrintConditionStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AmbientCalculusParser) PrintConditionStatement() (localctx IPrintConditionStatementContext) {
	localctx = NewPrintConditionStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, AmbientCalculusParserRULE_printConditionStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(151)
		p.Match(AmbientCalculusParserT__14)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(152)
		p.Match(AmbientCalculusParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(153)
		p.Match(AmbientCalculusParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IModifyConditionStatementContext is an interface to support dynamic dispatch.
type IModifyConditionStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// Getter signatures
	ID() antlr.TerminalNode
	Expression() IExpressionContext

	// IsModifyConditionStatementContext differentiates from other interfaces.
	IsModifyConditionStatementContext()
}

type ModifyConditionStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	op     antlr.Token
}

func NewEmptyModifyConditionStatementContext() *ModifyConditionStatementContext {
	var p = new(ModifyConditionStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AmbientCalculusParserRULE_modifyConditionStatement
	return p
}

func InitEmptyModifyConditionStatementContext(p *ModifyConditionStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AmbientCalculusParserRULE_modifyConditionStatement
}

func (*ModifyConditionStatementContext) IsModifyConditionStatementContext() {}

func NewModifyConditionStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ModifyConditionStatementContext {
	var p = new(ModifyConditionStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AmbientCalculusParserRULE_modifyConditionStatement

	return p
}

func (s *ModifyConditionStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ModifyConditionStatementContext) GetOp() antlr.Token { return s.op }

func (s *ModifyConditionStatementContext) SetOp(v antlr.Token) { s.op = v }

func (s *ModifyConditionStatementContext) ID() antlr.TerminalNode {
	return s.GetToken(AmbientCalculusParserID, 0)
}

func (s *ModifyConditionStatementContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ModifyConditionStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ModifyConditionStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ModifyConditionStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AmbientCalculusListener); ok {
		listenerT.EnterModifyConditionStatement(s)
	}
}

func (s *ModifyConditionStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AmbientCalculusListener); ok {
		listenerT.ExitModifyConditionStatement(s)
	}
}

func (s *ModifyConditionStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AmbientCalculusVisitor:
		return t.VisitModifyConditionStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AmbientCalculusParser) ModifyConditionStatement() (localctx IModifyConditionStatementContext) {
	localctx = NewModifyConditionStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, AmbientCalculusParserRULE_modifyConditionStatement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(155)
		p.Match(AmbientCalculusParserT__15)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(156)
		p.Match(AmbientCalculusParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(157)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*ModifyConditionStatementContext).op = _lt

		_la = p.GetTokenStream().LA(1)

		if !(_la == AmbientCalculusParserT__16 || _la == AmbientCalculusParserT__17) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*ModifyConditionStatementContext).op = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(158)
		p.expression(0)
	}
	{
		p.SetState(159)
		p.Match(AmbientCalculusParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IConditionsBlockContext is an interface to support dynamic dispatch.
type IConditionsBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsConditionsBlockContext differentiates from other interfaces.
	IsConditionsBlockContext()
}

type ConditionsBlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConditionsBlockContext() *ConditionsBlockContext {
	var p = new(ConditionsBlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AmbientCalculusParserRULE_conditionsBlock
	return p
}

func InitEmptyConditionsBlockContext(p *ConditionsBlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AmbientCalculusParserRULE_conditionsBlock
}

func (*ConditionsBlockContext) IsConditionsBlockContext() {}

func NewConditionsBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionsBlockContext {
	var p = new(ConditionsBlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AmbientCalculusParserRULE_conditionsBlock

	return p
}

func (s *ConditionsBlockContext) GetParser() antlr.Parser { return s.parser }
func (s *ConditionsBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionsBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConditionsBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AmbientCalculusListener); ok {
		listenerT.EnterConditionsBlock(s)
	}
}

func (s *ConditionsBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AmbientCalculusListener); ok {
		listenerT.ExitConditionsBlock(s)
	}
}

func (s *ConditionsBlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AmbientCalculusVisitor:
		return t.VisitConditionsBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AmbientCalculusParser) ConditionsBlock() (localctx IConditionsBlockContext) {
	localctx = NewConditionsBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, AmbientCalculusParserRULE_conditionsBlock)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(161)
		p.Match(AmbientCalculusParserT__18)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVariableDeclarationContext is an interface to support dynamic dispatch.
type IVariableDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	Expression() IExpressionContext

	// IsVariableDeclarationContext differentiates from other interfaces.
	IsVariableDeclarationContext()
}

type VariableDeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableDeclarationContext() *VariableDeclarationContext {
	var p = new(VariableDeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AmbientCalculusParserRULE_variableDeclaration
	return p
}

func InitEmptyVariableDeclarationContext(p *VariableDeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AmbientCalculusParserRULE_variableDeclaration
}

func (*VariableDeclarationContext) IsVariableDeclarationContext() {}

func NewVariableDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableDeclarationContext {
	var p = new(VariableDeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AmbientCalculusParserRULE_variableDeclaration

	return p
}

func (s *VariableDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableDeclarationContext) ID() antlr.TerminalNode {
	return s.GetToken(AmbientCalculusParserID, 0)
}

func (s *VariableDeclarationContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *VariableDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AmbientCalculusListener); ok {
		listenerT.EnterVariableDeclaration(s)
	}
}

func (s *VariableDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AmbientCalculusListener); ok {
		listenerT.ExitVariableDeclaration(s)
	}
}

func (s *VariableDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AmbientCalculusVisitor:
		return t.VisitVariableDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AmbientCalculusParser) VariableDeclaration() (localctx IVariableDeclarationContext) {
	localctx = NewVariableDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, AmbientCalculusParserRULE_variableDeclaration)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(163)
		p.Match(AmbientCalculusParserT__19)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(164)
		p.Match(AmbientCalculusParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(165)
		p.Match(AmbientCalculusParserT__20)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(166)
		p.expression(0)
	}
	{
		p.SetState(167)
		p.Match(AmbientCalculusParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IConditionsVariableDeclarationContext is an interface to support dynamic dispatch.
type IConditionsVariableDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	INT() antlr.TerminalNode
	Restriction() IRestrictionContext

	// IsConditionsVariableDeclarationContext differentiates from other interfaces.
	IsConditionsVariableDeclarationContext()
}

type ConditionsVariableDeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConditionsVariableDeclarationContext() *ConditionsVariableDeclarationContext {
	var p = new(ConditionsVariableDeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AmbientCalculusParserRULE_conditionsVariableDeclaration
	return p
}

func InitEmptyConditionsVariableDeclarationContext(p *ConditionsVariableDeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AmbientCalculusParserRULE_conditionsVariableDeclaration
}

func (*ConditionsVariableDeclarationContext) IsConditionsVariableDeclarationContext() {}

func NewConditionsVariableDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionsVariableDeclarationContext {
	var p = new(ConditionsVariableDeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AmbientCalculusParserRULE_conditionsVariableDeclaration

	return p
}

func (s *ConditionsVariableDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionsVariableDeclarationContext) ID() antlr.TerminalNode {
	return s.GetToken(AmbientCalculusParserID, 0)
}

func (s *ConditionsVariableDeclarationContext) INT() antlr.TerminalNode {
	return s.GetToken(AmbientCalculusParserINT, 0)
}

func (s *ConditionsVariableDeclarationContext) Restriction() IRestrictionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRestrictionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRestrictionContext)
}

func (s *ConditionsVariableDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionsVariableDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConditionsVariableDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AmbientCalculusListener); ok {
		listenerT.EnterConditionsVariableDeclaration(s)
	}
}

func (s *ConditionsVariableDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AmbientCalculusListener); ok {
		listenerT.ExitConditionsVariableDeclaration(s)
	}
}

func (s *ConditionsVariableDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AmbientCalculusVisitor:
		return t.VisitConditionsVariableDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AmbientCalculusParser) ConditionsVariableDeclaration() (localctx IConditionsVariableDeclarationContext) {
	localctx = NewConditionsVariableDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, AmbientCalculusParserRULE_conditionsVariableDeclaration)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(169)
		p.Match(AmbientCalculusParserT__21)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(170)
		p.Match(AmbientCalculusParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(171)
		p.Match(AmbientCalculusParserT__20)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(172)
		p.Match(AmbientCalculusParserINT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(174)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == AmbientCalculusParserT__22 {
		{
			p.SetState(173)
			p.Restriction()
		}

	}
	{
		p.SetState(176)
		p.Match(AmbientCalculusParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRestrictionContext is an interface to support dynamic dispatch.
type IRestrictionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Comparator() IComparatorContext
	INT() antlr.TerminalNode

	// IsRestrictionContext differentiates from other interfaces.
	IsRestrictionContext()
}

type RestrictionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRestrictionContext() *RestrictionContext {
	var p = new(RestrictionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AmbientCalculusParserRULE_restriction
	return p
}

func InitEmptyRestrictionContext(p *RestrictionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AmbientCalculusParserRULE_restriction
}

func (*RestrictionContext) IsRestrictionContext() {}

func NewRestrictionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RestrictionContext {
	var p = new(RestrictionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AmbientCalculusParserRULE_restriction

	return p
}

func (s *RestrictionContext) GetParser() antlr.Parser { return s.parser }

func (s *RestrictionContext) Comparator() IComparatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IComparatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IComparatorContext)
}

func (s *RestrictionContext) INT() antlr.TerminalNode {
	return s.GetToken(AmbientCalculusParserINT, 0)
}

func (s *RestrictionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RestrictionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RestrictionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AmbientCalculusListener); ok {
		listenerT.EnterRestriction(s)
	}
}

func (s *RestrictionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AmbientCalculusListener); ok {
		listenerT.ExitRestriction(s)
	}
}

func (s *RestrictionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AmbientCalculusVisitor:
		return t.VisitRestriction(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AmbientCalculusParser) Restriction() (localctx IRestrictionContext) {
	localctx = NewRestrictionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, AmbientCalculusParserRULE_restriction)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(178)
		p.Match(AmbientCalculusParserT__22)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(179)
		p.Comparator()
	}
	{
		p.SetState(180)
		p.Match(AmbientCalculusParserINT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAssignmentStatementContext is an interface to support dynamic dispatch.
type IAssignmentStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	Expression() IExpressionContext

	// IsAssignmentStatementContext differentiates from other interfaces.
	IsAssignmentStatementContext()
}

type AssignmentStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignmentStatementContext() *AssignmentStatementContext {
	var p = new(AssignmentStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AmbientCalculusParserRULE_assignmentStatement
	return p
}

func InitEmptyAssignmentStatementContext(p *AssignmentStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AmbientCalculusParserRULE_assignmentStatement
}

func (*AssignmentStatementContext) IsAssignmentStatementContext() {}

func NewAssignmentStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentStatementContext {
	var p = new(AssignmentStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AmbientCalculusParserRULE_assignmentStatement

	return p
}

func (s *AssignmentStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentStatementContext) ID() antlr.TerminalNode {
	return s.GetToken(AmbientCalculusParserID, 0)
}

func (s *AssignmentStatementContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AssignmentStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignmentStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AmbientCalculusListener); ok {
		listenerT.EnterAssignmentStatement(s)
	}
}

func (s *AssignmentStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AmbientCalculusListener); ok {
		listenerT.ExitAssignmentStatement(s)
	}
}

func (s *AssignmentStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AmbientCalculusVisitor:
		return t.VisitAssignmentStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AmbientCalculusParser) AssignmentStatement() (localctx IAssignmentStatementContext) {
	localctx = NewAssignmentStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, AmbientCalculusParserRULE_assignmentStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(182)
		p.Match(AmbientCalculusParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(183)
		p.Match(AmbientCalculusParserT__20)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(184)
		p.expression(0)
	}
	{
		p.SetState(185)
		p.Match(AmbientCalculusParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	ID() antlr.TerminalNode
	INT() antlr.TerminalNode
	STRING() antlr.TerminalNode
	Operator() IOperatorContext

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AmbientCalculusParserRULE_expression
	return p
}

func InitEmptyExpressionContext(p *ExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AmbientCalculusParserRULE_expression
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AmbientCalculusParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionContext) ID() antlr.TerminalNode {
	return s.GetToken(AmbientCalculusParserID, 0)
}

func (s *ExpressionContext) INT() antlr.TerminalNode {
	return s.GetToken(AmbientCalculusParserINT, 0)
}

func (s *ExpressionContext) STRING() antlr.TerminalNode {
	return s.GetToken(AmbientCalculusParserSTRING, 0)
}

func (s *ExpressionContext) Operator() IOperatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOperatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOperatorContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AmbientCalculusListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AmbientCalculusListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (s *ExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AmbientCalculusVisitor:
		return t.VisitExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AmbientCalculusParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *AmbientCalculusParser) expression(_p int) (localctx IExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 44
	p.EnterRecursionRule(localctx, 44, AmbientCalculusParserRULE_expression, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(195)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case AmbientCalculusParserT__23:
		{
			p.SetState(188)
			p.Match(AmbientCalculusParserT__23)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(189)
			p.expression(0)
		}
		{
			p.SetState(190)
			p.Match(AmbientCalculusParserT__24)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case AmbientCalculusParserID:
		{
			p.SetState(192)
			p.Match(AmbientCalculusParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case AmbientCalculusParserINT:
		{
			p.SetState(193)
			p.Match(AmbientCalculusParserINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case AmbientCalculusParserSTRING:
		{
			p.SetState(194)
			p.Match(AmbientCalculusParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(203)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewExpressionContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, AmbientCalculusParserRULE_expression)
			p.SetState(197)

			if !(p.Precpred(p.GetParserRuleContext(), 5)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				goto errorExit
			}
			{
				p.SetState(198)
				p.Operator()
			}
			{
				p.SetState(199)
				p.expression(6)
			}

		}
		p.SetState(205)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOperatorContext is an interface to support dynamic dispatch.
type IOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsOperatorContext differentiates from other interfaces.
	IsOperatorContext()
}

type OperatorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperatorContext() *OperatorContext {
	var p = new(OperatorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AmbientCalculusParserRULE_operator
	return p
}

func InitEmptyOperatorContext(p *OperatorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AmbientCalculusParserRULE_operator
}

func (*OperatorContext) IsOperatorContext() {}

func NewOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperatorContext {
	var p = new(OperatorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AmbientCalculusParserRULE_operator

	return p
}

func (s *OperatorContext) GetParser() antlr.Parser { return s.parser }
func (s *OperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AmbientCalculusListener); ok {
		listenerT.EnterOperator(s)
	}
}

func (s *OperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AmbientCalculusListener); ok {
		listenerT.ExitOperator(s)
	}
}

func (s *OperatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AmbientCalculusVisitor:
		return t.VisitOperator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AmbientCalculusParser) Operator() (localctx IOperatorContext) {
	localctx = NewOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, AmbientCalculusParserRULE_operator)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(206)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1006632960) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IComparatorContext is an interface to support dynamic dispatch.
type IComparatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsComparatorContext differentiates from other interfaces.
	IsComparatorContext()
}

type ComparatorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComparatorContext() *ComparatorContext {
	var p = new(ComparatorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AmbientCalculusParserRULE_comparator
	return p
}

func InitEmptyComparatorContext(p *ComparatorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AmbientCalculusParserRULE_comparator
}

func (*ComparatorContext) IsComparatorContext() {}

func NewComparatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComparatorContext {
	var p = new(ComparatorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AmbientCalculusParserRULE_comparator

	return p
}

func (s *ComparatorContext) GetParser() antlr.Parser { return s.parser }
func (s *ComparatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComparatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ComparatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AmbientCalculusListener); ok {
		listenerT.EnterComparator(s)
	}
}

func (s *ComparatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AmbientCalculusListener); ok {
		listenerT.ExitComparator(s)
	}
}

func (s *ComparatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AmbientCalculusVisitor:
		return t.VisitComparator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AmbientCalculusParser) Comparator() (localctx IComparatorContext) {
	localctx = NewComparatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, AmbientCalculusParserRULE_comparator)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(208)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&67645734912) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *AmbientCalculusParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 22:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *AmbientCalculusParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 5)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
