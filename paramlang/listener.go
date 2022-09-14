package paramlang

import (
	"golang.org/x/xerrors"

	"coder.com/m/product/coder/pkg/env/envtemplate/paramlang/parser"
)

var _ parser.WacListener = &wacListener{}

// newEvalWacListener is the listener to use to evaluate a given
// paramlang expression.
func newEvalWacListener(params map[string]ParamValue) *wacListener {
	w := new(wacListener)
	w.lookup = params

	return w
}

type wacListener struct {
	parser.BaseWacListener

	stack  []*stackItem
	lookup map[string]ParamValue
}

// ExitStringLit handles string literals. All string literals should be quoted
func (s *wacListener) ExitStringLit(ctx *parser.StringLitContext) {
	val, err := toString(ctx.GetText())
	if err != nil {
		s.stopParsing(xerrors.Errorf("exit string: %w", err))
	}

	s.pushString(val)
}

// ExitFloatLit literals match the golang float literal syntax
func (s *wacListener) ExitFloatLit(ctx *parser.FloatLitContext) {
	val, err := toFloat(ctx.GetText())
	if err != nil {
		s.stopParsing(xerrors.Errorf("exit float: %w", err))
	}

	s.pushFloat(val)
}

// ExitIntegerLit literals are whole numbers
func (s *wacListener) ExitIntegerLit(ctx *parser.IntegerLitContext) {
	val, err := toInteger(ctx.GetText())
	if err != nil {
		s.stopParsing(xerrors.Errorf("exit int: %w", err))
	}

	s.pushInt(val)
}

// ExitBoolLit literals match 'true' and 'false'
func (s *wacListener) ExitBoolLit(ctx *parser.BoolLitContext) {
	val, err := toBool(ctx.GetText())
	if err != nil {
		s.stopParsing(xerrors.Errorf("exit bool: %w", err))
	}

	s.pushBool(val)
}

// ExitInvertLogical handles '!boolean'
func (s *wacListener) ExitInvertLogical(_ *parser.InvertLogicalContext) {
	b := s.popBool()
	s.pushBool(!b)
}

// ExitKeyPath handles expanding params
func (s *wacListener) ExitKeyPath(ctx *parser.KeyPathContext) {
	v, ok := s.lookup[ctx.GetText()]
	if !ok {
		s.stopParsing(xerrors.Errorf("no key %q", ctx.GetText()))
	}

	// TODO: Actually handle the types?
	s.pushDefer(v.Value)
}

// ExitIfStmt handles 'if ... then ... else ...'
func (s *wacListener) ExitIfStmt(_ *parser.IfStmtContext) {
	elseV, thenV, ifV := s.pop(), s.pop(), s.popBool()
	if ifV {
		s.pushStackItem(thenV)
	} else {
		s.pushStackItem(elseV)
	}
}

type parseError struct {
	err error
}

func (p parseError) Error() string {
	return p.err.Error()
}

func (p parseError) Unwrap() error {
	return p.err
}

func (s *wacListener) stopParsing(err error) {
	panic(&parseError{err: err})
}
