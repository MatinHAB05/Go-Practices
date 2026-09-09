// Code generated from XXX.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // XXX

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

type XXXParser struct {
	*antlr.BaseParser
}

var XXXParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func xxxParserInit() {
	staticData := &XXXParserStaticData
	staticData.LiteralNames = []string{
		"", "'{'", "','", "'}'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "INT", "WS",
	}
	staticData.RuleNames = []string{
		"start_mamad", "value",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 5, 20, 2, 0, 7, 0, 2, 1, 7, 1, 1, 0, 1, 0, 1, 0, 1, 0, 5, 0, 9, 8,
		0, 10, 0, 12, 0, 12, 9, 0, 1, 0, 1, 0, 1, 1, 1, 1, 3, 1, 18, 8, 1, 1, 1,
		0, 0, 2, 0, 2, 0, 0, 19, 0, 4, 1, 0, 0, 0, 2, 17, 1, 0, 0, 0, 4, 5, 5,
		1, 0, 0, 5, 10, 3, 2, 1, 0, 6, 7, 5, 2, 0, 0, 7, 9, 3, 2, 1, 0, 8, 6, 1,
		0, 0, 0, 9, 12, 1, 0, 0, 0, 10, 8, 1, 0, 0, 0, 10, 11, 1, 0, 0, 0, 11,
		13, 1, 0, 0, 0, 12, 10, 1, 0, 0, 0, 13, 14, 5, 3, 0, 0, 14, 1, 1, 0, 0,
		0, 15, 18, 3, 0, 0, 0, 16, 18, 5, 4, 0, 0, 17, 15, 1, 0, 0, 0, 17, 16,
		1, 0, 0, 0, 18, 3, 1, 0, 0, 0, 2, 10, 17,
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

// XXXParserInit initializes any static state used to implement XXXParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewXXXParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func XXXParserInit() {
	staticData := &XXXParserStaticData
	staticData.once.Do(xxxParserInit)
}

// NewXXXParser produces a new parser instance for the optional input antlr.TokenStream.
func NewXXXParser(input antlr.TokenStream) *XXXParser {
	XXXParserInit()
	this := new(XXXParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &XXXParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "XXX.g4"

	return this
}

// XXXParser tokens.
const (
	XXXParserEOF  = antlr.TokenEOF
	XXXParserT__0 = 1
	XXXParserT__1 = 2
	XXXParserT__2 = 3
	XXXParserINT  = 4
	XXXParserWS   = 5
)

// XXXParser rules.
const (
	XXXParserRULE_start_mamad = 0
	XXXParserRULE_value       = 1
)

// IStart_mamadContext is an interface to support dynamic dispatch.
type IStart_mamadContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllValue() []IValueContext
	Value(i int) IValueContext

	// IsStart_mamadContext differentiates from other interfaces.
	IsStart_mamadContext()
}

type Start_mamadContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStart_mamadContext() *Start_mamadContext {
	var p = new(Start_mamadContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = XXXParserRULE_start_mamad
	return p
}

func InitEmptyStart_mamadContext(p *Start_mamadContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = XXXParserRULE_start_mamad
}

func (*Start_mamadContext) IsStart_mamadContext() {}

func NewStart_mamadContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Start_mamadContext {
	var p = new(Start_mamadContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = XXXParserRULE_start_mamad

	return p
}

func (s *Start_mamadContext) GetParser() antlr.Parser { return s.parser }

func (s *Start_mamadContext) AllValue() []IValueContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IValueContext); ok {
			len++
		}
	}

	tst := make([]IValueContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IValueContext); ok {
			tst[i] = t.(IValueContext)
			i++
		}
	}

	return tst
}

func (s *Start_mamadContext) Value(i int) IValueContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
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

	return t.(IValueContext)
}

func (s *Start_mamadContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Start_mamadContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Start_mamadContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(XXXListener); ok {
		listenerT.EnterStart_mamad(s)
	}
}

func (s *Start_mamadContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(XXXListener); ok {
		listenerT.ExitStart_mamad(s)
	}
}

func (p *XXXParser) Start_mamad() (localctx IStart_mamadContext) {
	localctx = NewStart_mamadContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, XXXParserRULE_start_mamad)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(4)
		p.Match(XXXParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(5)
		p.Value()
	}
	p.SetState(10)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == XXXParserT__1 {
		{
			p.SetState(6)
			p.Match(XXXParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(7)
			p.Value()
		}

		p.SetState(12)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(13)
		p.Match(XXXParserT__2)
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

// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Start_mamad() IStart_mamadContext
	INT() antlr.TerminalNode

	// IsValueContext differentiates from other interfaces.
	IsValueContext()
}

type ValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueContext() *ValueContext {
	var p = new(ValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = XXXParserRULE_value
	return p
}

func InitEmptyValueContext(p *ValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = XXXParserRULE_value
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = XXXParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) Start_mamad() IStart_mamadContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStart_mamadContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStart_mamadContext)
}

func (s *ValueContext) INT() antlr.TerminalNode {
	return s.GetToken(XXXParserINT, 0)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(XXXListener); ok {
		listenerT.EnterValue(s)
	}
}

func (s *ValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(XXXListener); ok {
		listenerT.ExitValue(s)
	}
}

func (p *XXXParser) Value() (localctx IValueContext) {
	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, XXXParserRULE_value)
	p.SetState(17)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case XXXParserT__0:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(15)
			p.Start_mamad()
		}

	case XXXParserINT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(16)
			p.Match(XXXParserINT)
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
