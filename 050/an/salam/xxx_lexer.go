// Code generated from XXX.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser

import (
	"fmt"
	"sync"
	"unicode"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type XXXLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var XXXLexerLexerStaticData struct {
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

func xxxlexerLexerInit() {
	staticData := &XXXLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'{'", "','", "'}'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "INT", "WS",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "INT", "WS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 5, 29, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 4, 3, 19, 8, 3, 11,
		3, 12, 3, 20, 1, 4, 4, 4, 24, 8, 4, 11, 4, 12, 4, 25, 1, 4, 1, 4, 0, 0,
		5, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 1, 0, 2, 1, 0, 48, 57, 3, 0, 9, 10, 13,
		13, 32, 32, 30, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0,
		7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 1, 11, 1, 0, 0, 0, 3, 13, 1, 0, 0, 0,
		5, 15, 1, 0, 0, 0, 7, 18, 1, 0, 0, 0, 9, 23, 1, 0, 0, 0, 11, 12, 5, 123,
		0, 0, 12, 2, 1, 0, 0, 0, 13, 14, 5, 44, 0, 0, 14, 4, 1, 0, 0, 0, 15, 16,
		5, 125, 0, 0, 16, 6, 1, 0, 0, 0, 17, 19, 7, 0, 0, 0, 18, 17, 1, 0, 0, 0,
		19, 20, 1, 0, 0, 0, 20, 18, 1, 0, 0, 0, 20, 21, 1, 0, 0, 0, 21, 8, 1, 0,
		0, 0, 22, 24, 7, 1, 0, 0, 23, 22, 1, 0, 0, 0, 24, 25, 1, 0, 0, 0, 25, 23,
		1, 0, 0, 0, 25, 26, 1, 0, 0, 0, 26, 27, 1, 0, 0, 0, 27, 28, 6, 4, 0, 0,
		28, 10, 1, 0, 0, 0, 3, 0, 20, 25, 1, 6, 0, 0,
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

// XXXLexerInit initializes any static state used to implement XXXLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewXXXLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func XXXLexerInit() {
	staticData := &XXXLexerLexerStaticData
	staticData.once.Do(xxxlexerLexerInit)
}

// NewXXXLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewXXXLexer(input antlr.CharStream) *XXXLexer {
	XXXLexerInit()
	l := new(XXXLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &XXXLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "XXX.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// XXXLexer tokens.
const (
	XXXLexerT__0 = 1
	XXXLexerT__1 = 2
	XXXLexerT__2 = 3
	XXXLexerINT  = 4
	XXXLexerWS   = 5
)
