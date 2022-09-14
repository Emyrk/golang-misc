// Code generated from Wac.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // Wac

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 29, 72, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 3, 2, 3, 2, 3,
	2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 22, 10, 4, 3, 5, 3, 5, 3,
	5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 32, 10, 5, 3, 6, 3, 6, 3, 6, 3,
	6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3,
	6, 3, 6, 3, 6, 3, 6, 5, 6, 53, 10, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3,
	6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 7, 6, 67, 10, 6, 12, 6, 14, 6, 70,
	11, 6, 3, 6, 2, 3, 10, 7, 2, 4, 6, 8, 10, 2, 6, 3, 2, 14, 15, 3, 2, 16,
	17, 3, 2, 5, 10, 3, 2, 11, 12, 2, 79, 2, 12, 3, 2, 2, 2, 4, 15, 3, 2, 2,
	2, 6, 21, 3, 2, 2, 2, 8, 31, 3, 2, 2, 2, 10, 52, 3, 2, 2, 2, 12, 13, 5,
	8, 5, 2, 13, 14, 7, 2, 2, 3, 14, 3, 3, 2, 2, 2, 15, 16, 7, 29, 2, 2, 16,
	5, 3, 2, 2, 2, 17, 22, 7, 18, 2, 2, 18, 22, 7, 21, 2, 2, 19, 22, 7, 19,
	2, 2, 20, 22, 7, 20, 2, 2, 21, 17, 3, 2, 2, 2, 21, 18, 3, 2, 2, 2, 21,
	19, 3, 2, 2, 2, 21, 20, 3, 2, 2, 2, 22, 7, 3, 2, 2, 2, 23, 32, 5, 10, 6,
	2, 24, 25, 7, 23, 2, 2, 25, 26, 5, 10, 6, 2, 26, 27, 7, 25, 2, 2, 27, 28,
	5, 10, 6, 2, 28, 29, 7, 24, 2, 2, 29, 30, 5, 10, 6, 2, 30, 32, 3, 2, 2,
	2, 31, 23, 3, 2, 2, 2, 31, 24, 3, 2, 2, 2, 32, 9, 3, 2, 2, 2, 33, 34, 8,
	6, 1, 2, 34, 53, 5, 6, 4, 2, 35, 36, 7, 26, 2, 2, 36, 37, 7, 3, 2, 2, 37,
	38, 5, 10, 6, 2, 38, 39, 7, 4, 2, 2, 39, 53, 3, 2, 2, 2, 40, 41, 7, 27,
	2, 2, 41, 42, 7, 3, 2, 2, 42, 43, 5, 10, 6, 2, 43, 44, 7, 4, 2, 2, 44,
	53, 3, 2, 2, 2, 45, 46, 7, 3, 2, 2, 46, 47, 5, 10, 6, 2, 47, 48, 7, 4,
	2, 2, 48, 53, 3, 2, 2, 2, 49, 50, 7, 13, 2, 2, 50, 53, 5, 10, 6, 4, 51,
	53, 5, 4, 3, 2, 52, 33, 3, 2, 2, 2, 52, 35, 3, 2, 2, 2, 52, 40, 3, 2, 2,
	2, 52, 45, 3, 2, 2, 2, 52, 49, 3, 2, 2, 2, 52, 51, 3, 2, 2, 2, 53, 68,
	3, 2, 2, 2, 54, 55, 12, 8, 2, 2, 55, 56, 9, 2, 2, 2, 56, 67, 5, 10, 6,
	9, 57, 58, 12, 7, 2, 2, 58, 59, 9, 3, 2, 2, 59, 67, 5, 10, 6, 8, 60, 61,
	12, 6, 2, 2, 61, 62, 9, 4, 2, 2, 62, 67, 5, 10, 6, 7, 63, 64, 12, 5, 2,
	2, 64, 65, 9, 5, 2, 2, 65, 67, 5, 10, 6, 6, 66, 54, 3, 2, 2, 2, 66, 57,
	3, 2, 2, 2, 66, 60, 3, 2, 2, 2, 66, 63, 3, 2, 2, 2, 67, 70, 3, 2, 2, 2,
	68, 66, 3, 2, 2, 2, 68, 69, 3, 2, 2, 2, 69, 11, 3, 2, 2, 2, 70, 68, 3,
	2, 2, 2, 7, 21, 31, 52, 66, 68,
}
var literalNames = []string{
	"", "'('", "')'", "'=='", "'!='", "'<'", "'<='", "'>'", "'>='", "'and'",
	"'or'", "'!'", "'*'", "'/'", "'+'", "'-'", "", "", "", "", "", "'if'",
	"'else'", "'then'", "'int'", "'float'", "'.'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "MUL", "DIV", "ADD", "SUB",
	"INTEGER", "FLOAT_LIT", "INTERPRETED_STRING_LIT", "BOOLEAN", "WHITESPACE",
	"IF", "ELSE", "THEN", "INT_CAST", "FLOAT_CAST", "DELIMITER", "KEY_PATH",
}

var ruleNames = []string{
	"start", "key", "literal", "statement", "expression",
}

type WacParser struct {
	*antlr.BaseParser
}

// NewWacParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *WacParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewWacParser(input antlr.TokenStream) *WacParser {
	this := new(WacParser)
	deserializer := antlr.NewATNDeserializer(nil)
	deserializedATN := deserializer.DeserializeFromUInt16(parserATN)
	decisionToDFA := make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Wac.g4"

	return this
}

// WacParser tokens.
const (
	WacParserEOF                    = antlr.TokenEOF
	WacParserT__0                   = 1
	WacParserT__1                   = 2
	WacParserT__2                   = 3
	WacParserT__3                   = 4
	WacParserT__4                   = 5
	WacParserT__5                   = 6
	WacParserT__6                   = 7
	WacParserT__7                   = 8
	WacParserT__8                   = 9
	WacParserT__9                   = 10
	WacParserT__10                  = 11
	WacParserMUL                    = 12
	WacParserDIV                    = 13
	WacParserADD                    = 14
	WacParserSUB                    = 15
	WacParserINTEGER                = 16
	WacParserFLOAT_LIT              = 17
	WacParserINTERPRETED_STRING_LIT = 18
	WacParserBOOLEAN                = 19
	WacParserWHITESPACE             = 20
	WacParserIF                     = 21
	WacParserELSE                   = 22
	WacParserTHEN                   = 23
	WacParserINT_CAST               = 24
	WacParserFLOAT_CAST             = 25
	WacParserDELIMITER              = 26
	WacParserKEY_PATH               = 27
)

// WacParser rules.
const (
	WacParserRULE_start      = 0
	WacParserRULE_key        = 1
	WacParserRULE_literal    = 2
	WacParserRULE_statement  = 3
	WacParserRULE_expression = 4
)

// IStartContext is an interface to support dynamic dispatch.
type IStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStartContext differentiates from other interfaces.
	IsStartContext()
}

type StartContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStartContext() *StartContext {
	var p = new(StartContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = WacParserRULE_start
	return p
}

func (*StartContext) IsStartContext() {}

func NewStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartContext {
	var p = new(StartContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = WacParserRULE_start

	return p
}

func (s *StartContext) GetParser() antlr.Parser { return s.parser }

func (s *StartContext) Statement() IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *StartContext) EOF() antlr.TerminalNode {
	return s.GetToken(WacParserEOF, 0)
}

func (s *StartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StartContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WacListener); ok {
		listenerT.EnterStart(s)
	}
}

func (s *StartContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WacListener); ok {
		listenerT.ExitStart(s)
	}
}

func (p *WacParser) Start() (localctx IStartContext) {
	localctx = NewStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, WacParserRULE_start)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(10)
		p.Statement()
	}
	{
		p.SetState(11)
		p.Match(WacParserEOF)
	}

	return localctx
}

// IKeyContext is an interface to support dynamic dispatch.
type IKeyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsKeyContext differentiates from other interfaces.
	IsKeyContext()
}

type KeyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKeyContext() *KeyContext {
	var p = new(KeyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = WacParserRULE_key
	return p
}

func (*KeyContext) IsKeyContext() {}

func NewKeyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KeyContext {
	var p = new(KeyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = WacParserRULE_key

	return p
}

func (s *KeyContext) GetParser() antlr.Parser { return s.parser }

func (s *KeyContext) KEY_PATH() antlr.TerminalNode {
	return s.GetToken(WacParserKEY_PATH, 0)
}

func (s *KeyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KeyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KeyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WacListener); ok {
		listenerT.EnterKey(s)
	}
}

func (s *KeyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WacListener); ok {
		listenerT.ExitKey(s)
	}
}

func (p *WacParser) Key() (localctx IKeyContext) {
	localctx = NewKeyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, WacParserRULE_key)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(13)
		p.Match(WacParserKEY_PATH)
	}

	return localctx
}

// ILiteralContext is an interface to support dynamic dispatch.
type ILiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLiteralContext differentiates from other interfaces.
	IsLiteralContext()
}

type LiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralContext() *LiteralContext {
	var p = new(LiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = WacParserRULE_literal
	return p
}

func (*LiteralContext) IsLiteralContext() {}

func NewLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralContext {
	var p = new(LiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = WacParserRULE_literal

	return p
}

func (s *LiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralContext) CopyFrom(ctx *LiteralContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type FloatLitContext struct {
	*LiteralContext
}

func NewFloatLitContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FloatLitContext {
	var p = new(FloatLitContext)

	p.LiteralContext = NewEmptyLiteralContext()
	p.parser = parser
	p.CopyFrom(ctx.(*LiteralContext))

	return p
}

func (s *FloatLitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FloatLitContext) FLOAT_LIT() antlr.TerminalNode {
	return s.GetToken(WacParserFLOAT_LIT, 0)
}

func (s *FloatLitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WacListener); ok {
		listenerT.EnterFloatLit(s)
	}
}

func (s *FloatLitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WacListener); ok {
		listenerT.ExitFloatLit(s)
	}
}

type IntegerLitContext struct {
	*LiteralContext
}

func NewIntegerLitContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IntegerLitContext {
	var p = new(IntegerLitContext)

	p.LiteralContext = NewEmptyLiteralContext()
	p.parser = parser
	p.CopyFrom(ctx.(*LiteralContext))

	return p
}

func (s *IntegerLitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntegerLitContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(WacParserINTEGER, 0)
}

func (s *IntegerLitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WacListener); ok {
		listenerT.EnterIntegerLit(s)
	}
}

func (s *IntegerLitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WacListener); ok {
		listenerT.ExitIntegerLit(s)
	}
}

type BoolLitContext struct {
	*LiteralContext
}

func NewBoolLitContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BoolLitContext {
	var p = new(BoolLitContext)

	p.LiteralContext = NewEmptyLiteralContext()
	p.parser = parser
	p.CopyFrom(ctx.(*LiteralContext))

	return p
}

func (s *BoolLitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolLitContext) BOOLEAN() antlr.TerminalNode {
	return s.GetToken(WacParserBOOLEAN, 0)
}

func (s *BoolLitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WacListener); ok {
		listenerT.EnterBoolLit(s)
	}
}

func (s *BoolLitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WacListener); ok {
		listenerT.ExitBoolLit(s)
	}
}

type StringLitContext struct {
	*LiteralContext
}

func NewStringLitContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringLitContext {
	var p = new(StringLitContext)

	p.LiteralContext = NewEmptyLiteralContext()
	p.parser = parser
	p.CopyFrom(ctx.(*LiteralContext))

	return p
}

func (s *StringLitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringLitContext) INTERPRETED_STRING_LIT() antlr.TerminalNode {
	return s.GetToken(WacParserINTERPRETED_STRING_LIT, 0)
}

func (s *StringLitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WacListener); ok {
		listenerT.EnterStringLit(s)
	}
}

func (s *StringLitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WacListener); ok {
		listenerT.ExitStringLit(s)
	}
}

func (p *WacParser) Literal() (localctx ILiteralContext) {
	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, WacParserRULE_literal)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(19)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case WacParserINTEGER:
		localctx = NewIntegerLitContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(15)
			p.Match(WacParserINTEGER)
		}

	case WacParserBOOLEAN:
		localctx = NewBoolLitContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(16)
			p.Match(WacParserBOOLEAN)
		}

	case WacParserFLOAT_LIT:
		localctx = NewFloatLitContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(17)
			p.Match(WacParserFLOAT_LIT)
		}

	case WacParserINTERPRETED_STRING_LIT:
		localctx = NewStringLitContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(18)
			p.Match(WacParserINTERPRETED_STRING_LIT)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = WacParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = WacParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) CopyFrom(ctx *StatementContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ExpStmtContext struct {
	*StatementContext
}

func NewExpStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpStmtContext {
	var p = new(ExpStmtContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *ExpStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpStmtContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WacListener); ok {
		listenerT.EnterExpStmt(s)
	}
}

func (s *ExpStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WacListener); ok {
		listenerT.ExitExpStmt(s)
	}
}

type IfStmtContext struct {
	*StatementContext
}

func NewIfStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IfStmtContext {
	var p = new(IfStmtContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *IfStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStmtContext) IF() antlr.TerminalNode {
	return s.GetToken(WacParserIF, 0)
}

func (s *IfStmtContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *IfStmtContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *IfStmtContext) THEN() antlr.TerminalNode {
	return s.GetToken(WacParserTHEN, 0)
}

func (s *IfStmtContext) ELSE() antlr.TerminalNode {
	return s.GetToken(WacParserELSE, 0)
}

func (s *IfStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WacListener); ok {
		listenerT.EnterIfStmt(s)
	}
}

func (s *IfStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WacListener); ok {
		listenerT.ExitIfStmt(s)
	}
}

func (p *WacParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, WacParserRULE_statement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(29)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case WacParserT__0, WacParserT__10, WacParserINTEGER, WacParserFLOAT_LIT, WacParserINTERPRETED_STRING_LIT, WacParserBOOLEAN, WacParserINT_CAST, WacParserFLOAT_CAST, WacParserKEY_PATH:
		localctx = NewExpStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(21)
			p.expression(0)
		}

	case WacParserIF:
		localctx = NewIfStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(22)
			p.Match(WacParserIF)
		}
		{
			p.SetState(23)
			p.expression(0)
		}
		{
			p.SetState(24)
			p.Match(WacParserTHEN)
		}
		{
			p.SetState(25)
			p.expression(0)
		}
		{
			p.SetState(26)
			p.Match(WacParserELSE)
		}
		{
			p.SetState(27)
			p.expression(0)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = WacParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = WacParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) CopyFrom(ctx *ExpressionContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type LiteralExpContext struct {
	*ExpressionContext
}

func NewLiteralExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LiteralExpContext {
	var p = new(LiteralExpContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *LiteralExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralExpContext) Literal() ILiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *LiteralExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WacListener); ok {
		listenerT.EnterLiteralExp(s)
	}
}

func (s *LiteralExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WacListener); ok {
		listenerT.ExitLiteralExp(s)
	}
}

type IntCastContext struct {
	*ExpressionContext
}

func NewIntCastContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IntCastContext {
	var p = new(IntCastContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *IntCastContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntCastContext) INT_CAST() antlr.TerminalNode {
	return s.GetToken(WacParserINT_CAST, 0)
}

func (s *IntCastContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *IntCastContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WacListener); ok {
		listenerT.EnterIntCast(s)
	}
}

func (s *IntCastContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WacListener); ok {
		listenerT.ExitIntCast(s)
	}
}

type InvertLogicalContext struct {
	*ExpressionContext
}

func NewInvertLogicalContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *InvertLogicalContext {
	var p = new(InvertLogicalContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *InvertLogicalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InvertLogicalContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *InvertLogicalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WacListener); ok {
		listenerT.EnterInvertLogical(s)
	}
}

func (s *InvertLogicalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WacListener); ok {
		listenerT.ExitInvertLogical(s)
	}
}

type MulDivContext struct {
	*ExpressionContext
	op antlr.Token
}

func NewMulDivContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MulDivContext {
	var p = new(MulDivContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *MulDivContext) GetOp() antlr.Token { return s.op }

func (s *MulDivContext) SetOp(v antlr.Token) { s.op = v }

func (s *MulDivContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MulDivContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *MulDivContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *MulDivContext) MUL() antlr.TerminalNode {
	return s.GetToken(WacParserMUL, 0)
}

func (s *MulDivContext) DIV() antlr.TerminalNode {
	return s.GetToken(WacParserDIV, 0)
}

func (s *MulDivContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WacListener); ok {
		listenerT.EnterMulDiv(s)
	}
}

func (s *MulDivContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WacListener); ok {
		listenerT.ExitMulDiv(s)
	}
}

type AddSubContext struct {
	*ExpressionContext
	op antlr.Token
}

func NewAddSubContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AddSubContext {
	var p = new(AddSubContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *AddSubContext) GetOp() antlr.Token { return s.op }

func (s *AddSubContext) SetOp(v antlr.Token) { s.op = v }

func (s *AddSubContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddSubContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *AddSubContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AddSubContext) ADD() antlr.TerminalNode {
	return s.GetToken(WacParserADD, 0)
}

func (s *AddSubContext) SUB() antlr.TerminalNode {
	return s.GetToken(WacParserSUB, 0)
}

func (s *AddSubContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WacListener); ok {
		listenerT.EnterAddSub(s)
	}
}

func (s *AddSubContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WacListener); ok {
		listenerT.ExitAddSub(s)
	}
}

type LogicalContext struct {
	*ExpressionContext
	op antlr.Token
}

func NewLogicalContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LogicalContext {
	var p = new(LogicalContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *LogicalContext) GetOp() antlr.Token { return s.op }

func (s *LogicalContext) SetOp(v antlr.Token) { s.op = v }

func (s *LogicalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogicalContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *LogicalContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *LogicalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WacListener); ok {
		listenerT.EnterLogical(s)
	}
}

func (s *LogicalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WacListener); ok {
		listenerT.ExitLogical(s)
	}
}

type FloatCastContext struct {
	*ExpressionContext
}

func NewFloatCastContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FloatCastContext {
	var p = new(FloatCastContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *FloatCastContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FloatCastContext) FLOAT_CAST() antlr.TerminalNode {
	return s.GetToken(WacParserFLOAT_CAST, 0)
}

func (s *FloatCastContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *FloatCastContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WacListener); ok {
		listenerT.EnterFloatCast(s)
	}
}

func (s *FloatCastContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WacListener); ok {
		listenerT.ExitFloatCast(s)
	}
}

type ParenContext struct {
	*ExpressionContext
}

func NewParenContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParenContext {
	var p = new(ParenContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *ParenContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParenContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ParenContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WacListener); ok {
		listenerT.EnterParen(s)
	}
}

func (s *ParenContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WacListener); ok {
		listenerT.ExitParen(s)
	}
}

type KeyPathContext struct {
	*ExpressionContext
}

func NewKeyPathContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *KeyPathContext {
	var p = new(KeyPathContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *KeyPathContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KeyPathContext) Key() IKeyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKeyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IKeyContext)
}

func (s *KeyPathContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WacListener); ok {
		listenerT.EnterKeyPath(s)
	}
}

func (s *KeyPathContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WacListener); ok {
		listenerT.ExitKeyPath(s)
	}
}

type ComparatorContext struct {
	*ExpressionContext
	op antlr.Token
}

func NewComparatorContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ComparatorContext {
	var p = new(ComparatorContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *ComparatorContext) GetOp() antlr.Token { return s.op }

func (s *ComparatorContext) SetOp(v antlr.Token) { s.op = v }

func (s *ComparatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComparatorContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *ComparatorContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ComparatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WacListener); ok {
		listenerT.EnterComparator(s)
	}
}

func (s *ComparatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WacListener); ok {
		listenerT.ExitComparator(s)
	}
}

func (p *WacParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *WacParser) expression(_p int) (localctx IExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 8
	p.EnterRecursionRule(localctx, 8, WacParserRULE_expression, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(50)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case WacParserINTEGER, WacParserFLOAT_LIT, WacParserINTERPRETED_STRING_LIT, WacParserBOOLEAN:
		localctx = NewLiteralExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(32)
			p.Literal()
		}

	case WacParserINT_CAST:
		localctx = NewIntCastContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(33)
			p.Match(WacParserINT_CAST)
		}
		{
			p.SetState(34)
			p.Match(WacParserT__0)
		}
		{
			p.SetState(35)
			p.expression(0)
		}
		{
			p.SetState(36)
			p.Match(WacParserT__1)
		}

	case WacParserFLOAT_CAST:
		localctx = NewFloatCastContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(38)
			p.Match(WacParserFLOAT_CAST)
		}
		{
			p.SetState(39)
			p.Match(WacParserT__0)
		}
		{
			p.SetState(40)
			p.expression(0)
		}
		{
			p.SetState(41)
			p.Match(WacParserT__1)
		}

	case WacParserT__0:
		localctx = NewParenContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(43)
			p.Match(WacParserT__0)
		}
		{
			p.SetState(44)
			p.expression(0)
		}
		{
			p.SetState(45)
			p.Match(WacParserT__1)
		}

	case WacParserT__10:
		localctx = NewInvertLogicalContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(47)
			p.Match(WacParserT__10)
		}
		{
			p.SetState(48)
			p.expression(2)
		}

	case WacParserKEY_PATH:
		localctx = NewKeyPathContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(49)
			p.Key()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(66)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(64)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
			case 1:
				localctx = NewMulDivContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, WacParserRULE_expression)
				p.SetState(52)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(53)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*MulDivContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == WacParserMUL || _la == WacParserDIV) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*MulDivContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(54)
					p.expression(7)
				}

			case 2:
				localctx = NewAddSubContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, WacParserRULE_expression)
				p.SetState(55)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(56)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*AddSubContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == WacParserADD || _la == WacParserSUB) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*AddSubContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(57)
					p.expression(6)
				}

			case 3:
				localctx = NewComparatorContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, WacParserRULE_expression)
				p.SetState(58)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(59)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ComparatorContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<WacParserT__2)|(1<<WacParserT__3)|(1<<WacParserT__4)|(1<<WacParserT__5)|(1<<WacParserT__6)|(1<<WacParserT__7))) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ComparatorContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(60)
					p.expression(5)
				}

			case 4:
				localctx = NewLogicalContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, WacParserRULE_expression)
				p.SetState(61)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(62)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*LogicalContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == WacParserT__8 || _la == WacParserT__9) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*LogicalContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(63)
					p.expression(4)
				}

			}

		}
		p.SetState(68)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())
	}

	return localctx
}

func (p *WacParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 4:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *WacParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
