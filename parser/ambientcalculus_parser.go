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
		"program", "statement", "processDeclaration", "processStatement", "moveStatement",
		"ambientDeclaration", "conditions", "movementStatement", "communicationStatement",
		"inStatement", "outStatement", "openStatement", "callStatement", "printStatement",
		"ifStatement", "variableDeclaration", "expression",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 31, 177, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 1, 0, 5, 0, 36, 8, 0, 10, 0, 12, 0, 39, 9, 0, 1, 0, 1, 0,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1,
		54, 8, 1, 1, 2, 1, 2, 1, 2, 1, 2, 5, 2, 60, 8, 2, 10, 2, 12, 2, 63, 9,
		2, 1, 2, 1, 2, 1, 3, 1, 3, 3, 3, 69, 8, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5,
		1, 5, 1, 5, 1, 5, 3, 5, 79, 8, 5, 1, 5, 5, 5, 82, 8, 5, 10, 5, 12, 5, 85,
		9, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 5, 6, 92, 8, 6, 10, 6, 12, 6, 95, 9,
		6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1,
		8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 3, 8, 116, 8, 8, 1, 9, 1,
		9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12,
		1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1,
		14, 1, 14, 1, 14, 5, 14, 143, 8, 14, 10, 14, 12, 14, 146, 9, 14, 1, 14,
		1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1,
		16, 1, 16, 1, 16, 1, 16, 1, 16, 3, 16, 164, 8, 16, 1, 16, 1, 16, 1, 16,
		1, 16, 1, 16, 1, 16, 5, 16, 172, 8, 16, 10, 16, 12, 16, 175, 9, 16, 1,
		16, 0, 1, 32, 17, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28,
		30, 32, 0, 2, 1, 0, 22, 23, 1, 0, 24, 25, 182, 0, 37, 1, 0, 0, 0, 2, 53,
		1, 0, 0, 0, 4, 55, 1, 0, 0, 0, 6, 68, 1, 0, 0, 0, 8, 70, 1, 0, 0, 0, 10,
		74, 1, 0, 0, 0, 12, 88, 1, 0, 0, 0, 14, 98, 1, 0, 0, 0, 16, 115, 1, 0,
		0, 0, 18, 117, 1, 0, 0, 0, 20, 121, 1, 0, 0, 0, 22, 124, 1, 0, 0, 0, 24,
		128, 1, 0, 0, 0, 26, 132, 1, 0, 0, 0, 28, 136, 1, 0, 0, 0, 30, 149, 1,
		0, 0, 0, 32, 163, 1, 0, 0, 0, 34, 36, 3, 2, 1, 0, 35, 34, 1, 0, 0, 0, 36,
		39, 1, 0, 0, 0, 37, 35, 1, 0, 0, 0, 37, 38, 1, 0, 0, 0, 38, 40, 1, 0, 0,
		0, 39, 37, 1, 0, 0, 0, 40, 41, 5, 0, 0, 1, 41, 1, 1, 0, 0, 0, 42, 54, 3,
		4, 2, 0, 43, 54, 3, 10, 5, 0, 44, 54, 3, 14, 7, 0, 45, 54, 3, 16, 8, 0,
		46, 54, 3, 18, 9, 0, 47, 54, 3, 20, 10, 0, 48, 54, 3, 22, 11, 0, 49, 54,
		3, 24, 12, 0, 50, 54, 3, 26, 13, 0, 51, 54, 3, 28, 14, 0, 52, 54, 3, 30,
		15, 0, 53, 42, 1, 0, 0, 0, 53, 43, 1, 0, 0, 0, 53, 44, 1, 0, 0, 0, 53,
		45, 1, 0, 0, 0, 53, 46, 1, 0, 0, 0, 53, 47, 1, 0, 0, 0, 53, 48, 1, 0, 0,
		0, 53, 49, 1, 0, 0, 0, 53, 50, 1, 0, 0, 0, 53, 51, 1, 0, 0, 0, 53, 52,
		1, 0, 0, 0, 54, 3, 1, 0, 0, 0, 55, 56, 5, 1, 0, 0, 56, 57, 5, 27, 0, 0,
		57, 61, 5, 2, 0, 0, 58, 60, 3, 6, 3, 0, 59, 58, 1, 0, 0, 0, 60, 63, 1,
		0, 0, 0, 61, 59, 1, 0, 0, 0, 61, 62, 1, 0, 0, 0, 62, 64, 1, 0, 0, 0, 63,
		61, 1, 0, 0, 0, 64, 65, 5, 3, 0, 0, 65, 5, 1, 0, 0, 0, 66, 69, 3, 2, 1,
		0, 67, 69, 3, 8, 4, 0, 68, 66, 1, 0, 0, 0, 68, 67, 1, 0, 0, 0, 69, 7, 1,
		0, 0, 0, 70, 71, 5, 4, 0, 0, 71, 72, 5, 26, 0, 0, 72, 73, 5, 5, 0, 0, 73,
		9, 1, 0, 0, 0, 74, 75, 5, 6, 0, 0, 75, 76, 5, 26, 0, 0, 76, 78, 5, 2, 0,
		0, 77, 79, 3, 12, 6, 0, 78, 77, 1, 0, 0, 0, 78, 79, 1, 0, 0, 0, 79, 83,
		1, 0, 0, 0, 80, 82, 3, 2, 1, 0, 81, 80, 1, 0, 0, 0, 82, 85, 1, 0, 0, 0,
		83, 81, 1, 0, 0, 0, 83, 84, 1, 0, 0, 0, 84, 86, 1, 0, 0, 0, 85, 83, 1,
		0, 0, 0, 86, 87, 5, 3, 0, 0, 87, 11, 1, 0, 0, 0, 88, 89, 5, 7, 0, 0, 89,
		93, 5, 2, 0, 0, 90, 92, 3, 2, 1, 0, 91, 90, 1, 0, 0, 0, 92, 95, 1, 0, 0,
		0, 93, 91, 1, 0, 0, 0, 93, 94, 1, 0, 0, 0, 94, 96, 1, 0, 0, 0, 95, 93,
		1, 0, 0, 0, 96, 97, 5, 3, 0, 0, 97, 13, 1, 0, 0, 0, 98, 99, 5, 4, 0, 0,
		99, 100, 5, 27, 0, 0, 100, 101, 5, 8, 0, 0, 101, 102, 5, 26, 0, 0, 102,
		103, 5, 5, 0, 0, 103, 15, 1, 0, 0, 0, 104, 105, 5, 9, 0, 0, 105, 106, 3,
		32, 16, 0, 106, 107, 5, 8, 0, 0, 107, 108, 5, 27, 0, 0, 108, 109, 5, 5,
		0, 0, 109, 116, 1, 0, 0, 0, 110, 111, 5, 10, 0, 0, 111, 112, 5, 27, 0,
		0, 112, 113, 5, 11, 0, 0, 113, 114, 5, 27, 0, 0, 114, 116, 5, 5, 0, 0,
		115, 104, 1, 0, 0, 0, 115, 110, 1, 0, 0, 0, 116, 17, 1, 0, 0, 0, 117, 118,
		5, 12, 0, 0, 118, 119, 5, 26, 0, 0, 119, 120, 5, 5, 0, 0, 120, 19, 1, 0,
		0, 0, 121, 122, 5, 13, 0, 0, 122, 123, 5, 5, 0, 0, 123, 21, 1, 0, 0, 0,
		124, 125, 5, 14, 0, 0, 125, 126, 5, 26, 0, 0, 126, 127, 5, 5, 0, 0, 127,
		23, 1, 0, 0, 0, 128, 129, 5, 15, 0, 0, 129, 130, 5, 27, 0, 0, 130, 131,
		5, 5, 0, 0, 131, 25, 1, 0, 0, 0, 132, 133, 5, 16, 0, 0, 133, 134, 3, 32,
		16, 0, 134, 135, 5, 5, 0, 0, 135, 27, 1, 0, 0, 0, 136, 137, 5, 17, 0, 0,
		137, 138, 5, 18, 0, 0, 138, 139, 3, 32, 16, 0, 139, 140, 5, 19, 0, 0, 140,
		144, 5, 2, 0, 0, 141, 143, 3, 2, 1, 0, 142, 141, 1, 0, 0, 0, 143, 146,
		1, 0, 0, 0, 144, 142, 1, 0, 0, 0, 144, 145, 1, 0, 0, 0, 145, 147, 1, 0,
		0, 0, 146, 144, 1, 0, 0, 0, 147, 148, 5, 3, 0, 0, 148, 29, 1, 0, 0, 0,
		149, 150, 5, 20, 0, 0, 150, 151, 5, 27, 0, 0, 151, 152, 5, 21, 0, 0, 152,
		153, 3, 32, 16, 0, 153, 154, 5, 5, 0, 0, 154, 31, 1, 0, 0, 0, 155, 156,
		6, 16, -1, 0, 156, 157, 5, 18, 0, 0, 157, 158, 3, 32, 16, 0, 158, 159,
		5, 19, 0, 0, 159, 164, 1, 0, 0, 0, 160, 164, 5, 27, 0, 0, 161, 164, 5,
		28, 0, 0, 162, 164, 5, 29, 0, 0, 163, 155, 1, 0, 0, 0, 163, 160, 1, 0,
		0, 0, 163, 161, 1, 0, 0, 0, 163, 162, 1, 0, 0, 0, 164, 173, 1, 0, 0, 0,
		165, 166, 10, 6, 0, 0, 166, 167, 7, 0, 0, 0, 167, 172, 3, 32, 16, 7, 168,
		169, 10, 5, 0, 0, 169, 170, 7, 1, 0, 0, 170, 172, 3, 32, 16, 6, 171, 165,
		1, 0, 0, 0, 171, 168, 1, 0, 0, 0, 172, 175, 1, 0, 0, 0, 173, 171, 1, 0,
		0, 0, 173, 174, 1, 0, 0, 0, 174, 33, 1, 0, 0, 0, 175, 173, 1, 0, 0, 0,
		12, 37, 53, 61, 68, 78, 83, 93, 115, 144, 163, 171, 173,
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
	AmbientCalculusParserAMBIENTID = 26
	AmbientCalculusParserID        = 27
	AmbientCalculusParserINT       = 28
	AmbientCalculusParserSTRING    = 29
	AmbientCalculusParserWS        = 30
	AmbientCalculusParserCOMMENT   = 31
)

// AmbientCalculusParser rules.
const (
	AmbientCalculusParserRULE_program                = 0
	AmbientCalculusParserRULE_statement              = 1
	AmbientCalculusParserRULE_processDeclaration     = 2
	AmbientCalculusParserRULE_processStatement       = 3
	AmbientCalculusParserRULE_moveStatement          = 4
	AmbientCalculusParserRULE_ambientDeclaration     = 5
	AmbientCalculusParserRULE_conditions             = 6
	AmbientCalculusParserRULE_movementStatement      = 7
	AmbientCalculusParserRULE_communicationStatement = 8
	AmbientCalculusParserRULE_inStatement            = 9
	AmbientCalculusParserRULE_outStatement           = 10
	AmbientCalculusParserRULE_openStatement          = 11
	AmbientCalculusParserRULE_callStatement          = 12
	AmbientCalculusParserRULE_printStatement         = 13
	AmbientCalculusParserRULE_ifStatement            = 14
	AmbientCalculusParserRULE_variableDeclaration    = 15
	AmbientCalculusParserRULE_expression             = 16
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
	p.SetState(37)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1308242) != 0 {
		{
			p.SetState(34)
			p.Statement()
		}

		p.SetState(39)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(40)
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
	CommunicationStatement() ICommunicationStatementContext
	InStatement() IInStatementContext
	OutStatement() IOutStatementContext
	OpenStatement() IOpenStatementContext
	CallStatement() ICallStatementContext
	PrintStatement() IPrintStatementContext
	IfStatement() IIfStatementContext
	VariableDeclaration() IVariableDeclarationContext

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

func (s *StatementContext) CommunicationStatement() ICommunicationStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICommunicationStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICommunicationStatementContext)
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

func (s *StatementContext) IfStatement() IIfStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfStatementContext)
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
	p.SetState(53)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case AmbientCalculusParserT__0:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(42)
			p.ProcessDeclaration()
		}

	case AmbientCalculusParserT__5:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(43)
			p.AmbientDeclaration()
		}

	case AmbientCalculusParserT__3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(44)
			p.MovementStatement()
		}

	case AmbientCalculusParserT__8, AmbientCalculusParserT__9:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(45)
			p.CommunicationStatement()
		}

	case AmbientCalculusParserT__11:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(46)
			p.InStatement()
		}

	case AmbientCalculusParserT__12:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(47)
			p.OutStatement()
		}

	case AmbientCalculusParserT__13:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(48)
			p.OpenStatement()
		}

	case AmbientCalculusParserT__14:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(49)
			p.CallStatement()
		}

	case AmbientCalculusParserT__15:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(50)
			p.PrintStatement()
		}

	case AmbientCalculusParserT__16:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(51)
			p.IfStatement()
		}

	case AmbientCalculusParserT__19:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(52)
			p.VariableDeclaration()
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
		p.SetState(55)
		p.Match(AmbientCalculusParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(56)
		p.Match(AmbientCalculusParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(57)
		p.Match(AmbientCalculusParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(61)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1308242) != 0 {
		{
			p.SetState(58)
			p.ProcessStatement()
		}

		p.SetState(63)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(64)
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
	p.SetState(68)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(66)
			p.Statement()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(67)
			p.MoveStatement()
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
		p.SetState(70)
		p.Match(AmbientCalculusParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(71)
		p.Match(AmbientCalculusParserAMBIENTID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(72)
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
		p.SetState(74)
		p.Match(AmbientCalculusParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(75)
		p.Match(AmbientCalculusParserAMBIENTID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(76)
		p.Match(AmbientCalculusParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(78)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == AmbientCalculusParserT__6 {
		{
			p.SetState(77)
			p.Conditions()
		}

	}
	p.SetState(83)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1308242) != 0 {
		{
			p.SetState(80)
			p.Statement()
		}

		p.SetState(85)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(86)
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
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

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

func (s *ConditionsContext) AllStatement() []IStatementContext {
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

func (s *ConditionsContext) Statement(i int) IStatementContext {
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
		p.SetState(88)
		p.Match(AmbientCalculusParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(89)
		p.Match(AmbientCalculusParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(93)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1308242) != 0 {
		{
			p.SetState(90)
			p.Statement()
		}

		p.SetState(95)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(96)
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
		p.SetState(98)
		p.Match(AmbientCalculusParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(99)
		p.Match(AmbientCalculusParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(100)
		p.Match(AmbientCalculusParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(101)
		p.Match(AmbientCalculusParserAMBIENTID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(102)
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

// ICommunicationStatementContext is an interface to support dynamic dispatch.
type ICommunicationStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext
	AllID() []antlr.TerminalNode
	ID(i int) antlr.TerminalNode

	// IsCommunicationStatementContext differentiates from other interfaces.
	IsCommunicationStatementContext()
}

type CommunicationStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCommunicationStatementContext() *CommunicationStatementContext {
	var p = new(CommunicationStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AmbientCalculusParserRULE_communicationStatement
	return p
}

func InitEmptyCommunicationStatementContext(p *CommunicationStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AmbientCalculusParserRULE_communicationStatement
}

func (*CommunicationStatementContext) IsCommunicationStatementContext() {}

func NewCommunicationStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommunicationStatementContext {
	var p = new(CommunicationStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AmbientCalculusParserRULE_communicationStatement

	return p
}

func (s *CommunicationStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *CommunicationStatementContext) Expression() IExpressionContext {
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

func (s *CommunicationStatementContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(AmbientCalculusParserID)
}

func (s *CommunicationStatementContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(AmbientCalculusParserID, i)
}

func (s *CommunicationStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommunicationStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CommunicationStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AmbientCalculusListener); ok {
		listenerT.EnterCommunicationStatement(s)
	}
}

func (s *CommunicationStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AmbientCalculusListener); ok {
		listenerT.ExitCommunicationStatement(s)
	}
}

func (s *CommunicationStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AmbientCalculusVisitor:
		return t.VisitCommunicationStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AmbientCalculusParser) CommunicationStatement() (localctx ICommunicationStatementContext) {
	localctx = NewCommunicationStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, AmbientCalculusParserRULE_communicationStatement)
	p.SetState(115)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case AmbientCalculusParserT__8:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(104)
			p.Match(AmbientCalculusParserT__8)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(105)
			p.expression(0)
		}
		{
			p.SetState(106)
			p.Match(AmbientCalculusParserT__7)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(107)
			p.Match(AmbientCalculusParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(108)
			p.Match(AmbientCalculusParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case AmbientCalculusParserT__9:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(110)
			p.Match(AmbientCalculusParserT__9)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(111)
			p.Match(AmbientCalculusParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(112)
			p.Match(AmbientCalculusParserT__10)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(113)
			p.Match(AmbientCalculusParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(114)
			p.Match(AmbientCalculusParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
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
	p.EnterRule(localctx, 18, AmbientCalculusParserRULE_inStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(117)
		p.Match(AmbientCalculusParserT__11)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(118)
		p.Match(AmbientCalculusParserAMBIENTID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(119)
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
	p.EnterRule(localctx, 20, AmbientCalculusParserRULE_outStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(121)
		p.Match(AmbientCalculusParserT__12)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(122)
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
	p.EnterRule(localctx, 22, AmbientCalculusParserRULE_openStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(124)
		p.Match(AmbientCalculusParserT__13)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(125)
		p.Match(AmbientCalculusParserAMBIENTID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(126)
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
	p.EnterRule(localctx, 24, AmbientCalculusParserRULE_callStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(128)
		p.Match(AmbientCalculusParserT__14)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(129)
		p.Match(AmbientCalculusParserID)
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
	p.EnterRule(localctx, 26, AmbientCalculusParserRULE_printStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(132)
		p.Match(AmbientCalculusParserT__15)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(133)
		p.expression(0)
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

// IIfStatementContext is an interface to support dynamic dispatch.
type IIfStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

	// IsIfStatementContext differentiates from other interfaces.
	IsIfStatementContext()
}

type IfStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfStatementContext() *IfStatementContext {
	var p = new(IfStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AmbientCalculusParserRULE_ifStatement
	return p
}

func InitEmptyIfStatementContext(p *IfStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AmbientCalculusParserRULE_ifStatement
}

func (*IfStatementContext) IsIfStatementContext() {}

func NewIfStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfStatementContext {
	var p = new(IfStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AmbientCalculusParserRULE_ifStatement

	return p
}

func (s *IfStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *IfStatementContext) Expression() IExpressionContext {
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

func (s *IfStatementContext) AllStatement() []IStatementContext {
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

func (s *IfStatementContext) Statement(i int) IStatementContext {
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

func (s *IfStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AmbientCalculusListener); ok {
		listenerT.EnterIfStatement(s)
	}
}

func (s *IfStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AmbientCalculusListener); ok {
		listenerT.ExitIfStatement(s)
	}
}

func (s *IfStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AmbientCalculusVisitor:
		return t.VisitIfStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AmbientCalculusParser) IfStatement() (localctx IIfStatementContext) {
	localctx = NewIfStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, AmbientCalculusParserRULE_ifStatement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(136)
		p.Match(AmbientCalculusParserT__16)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(137)
		p.Match(AmbientCalculusParserT__17)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(138)
		p.expression(0)
	}
	{
		p.SetState(139)
		p.Match(AmbientCalculusParserT__18)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(140)
		p.Match(AmbientCalculusParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(144)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1308242) != 0 {
		{
			p.SetState(141)
			p.Statement()
		}

		p.SetState(146)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(147)
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
	p.EnterRule(localctx, 30, AmbientCalculusParserRULE_variableDeclaration)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(149)
		p.Match(AmbientCalculusParserT__19)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(150)
		p.Match(AmbientCalculusParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(151)
		p.Match(AmbientCalculusParserT__20)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(152)
		p.expression(0)
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

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// Getter signatures
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	ID() antlr.TerminalNode
	INT() antlr.TerminalNode
	STRING() antlr.TerminalNode

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	op     antlr.Token
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

func (s *ExpressionContext) GetOp() antlr.Token { return s.op }

func (s *ExpressionContext) SetOp(v antlr.Token) { s.op = v }

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
	_startState := 32
	p.EnterRecursionRule(localctx, 32, AmbientCalculusParserRULE_expression, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(163)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case AmbientCalculusParserT__17:
		{
			p.SetState(156)
			p.Match(AmbientCalculusParserT__17)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(157)
			p.expression(0)
		}
		{
			p.SetState(158)
			p.Match(AmbientCalculusParserT__18)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case AmbientCalculusParserID:
		{
			p.SetState(160)
			p.Match(AmbientCalculusParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case AmbientCalculusParserINT:
		{
			p.SetState(161)
			p.Match(AmbientCalculusParserINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case AmbientCalculusParserSTRING:
		{
			p.SetState(162)
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
	p.SetState(173)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(171)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, AmbientCalculusParserRULE_expression)
				p.SetState(165)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(166)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpressionContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == AmbientCalculusParserT__21 || _la == AmbientCalculusParserT__22) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpressionContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(167)
					p.expression(7)
				}

			case 2:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, AmbientCalculusParserRULE_expression)
				p.SetState(168)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(169)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpressionContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == AmbientCalculusParserT__23 || _la == AmbientCalculusParserT__24) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpressionContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(170)
					p.expression(6)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(175)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext())
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

func (p *AmbientCalculusParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 16:
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
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 5)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
