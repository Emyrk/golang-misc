package paramlang

import (
	"regexp"
	"strings"

	"golang.org/x/xerrors"

	"github.com/antlr/antlr4/runtime/Go/antlr"

	"coder.com/m/product/coder/pkg/env/envtemplate/paramlang/parser"
)

const (
	StartTemplate   = "{{"
	EndTemplate     = "}}"
	exampleTemplate = StartTemplate + " ... " + EndTemplate
)

var TemplateRegex = regexp.MustCompile(`(?P<expression>{{[^}]*}})`)

// ValidateXXX will validate the given input will evaluate to the expected type XXX
// EvalXXX actually evaluates the given input expression

func ValidateInt(input string) error {
	return validate(input, stInt)
}

func ValidateFloat(input string) error {
	return validate(input, stFloat)
}

func ValidateBool(input string) error {
	return validate(input, stBool)
}

func ValidateString(input string) error {
	return validate(input, stString)
}

func EvalInt(input string, params map[string]ParamValue) (int, error) {
	v, err := eval(input, params)
	if err != nil {
		return -1, err
	}

	return v.AsInt()
}

func EvalFloat(input string, params map[string]ParamValue) (float64, error) {
	v, err := eval(input, params)
	if err != nil {
		return -1, err
	}

	return v.AsFloat()
}

func EvalBool(input string, params map[string]ParamValue) (bool, error) {
	v, err := eval(input, params)
	if err != nil {
		return false, err
	}

	return v.AsBool()
}

func EvalString(input string, params map[string]ParamValue) (string, error) {
	v, err := eval(input, params)
	if err != nil {
		return "", err
	}

	return v.AsString()
}

func setupInput(input string) (*parser.WacParser, *syntaxErrorListener) {
	is := antlr.NewInputStream(input)

	lexer := parser.NewWacLexer(is)
	s := newSyntaxErrorListener()
	// Default error listener prints to stdout
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(s)

	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewWacParser(stream)
	p.RemoveErrorListeners()
	p.AddErrorListener(s)
	return p, s
}

func eval(input string, params map[string]ParamValue) (*stackItem, error) {
	l := newEvalWacListener(params)
	hasExp, err := executeExpression(l, input)
	if err != nil {
		return nil, err
	}
	if !hasExp {
		return &stackItem{
			Value: input,
			Type:  stNoTemplate,
		}, nil
	}

	v := l.pop()
	if len(l.stack) > 0 {
		return nil, xerrors.Errorf("%d items remain on the stack, expression failed", len(l.stack))
	}
	return v, nil
}

func validate(input string, exp stackItemType) (err error) {
	l := &wacTypeCheck{}
	hasExp, err := executeExpression(l, input)
	if err != nil {
		return err
	}

	if !hasExp {
		return nil
	}

	v := l.pop()
	if v != exp && v != stDefer {
		return xerrors.Errorf("expected type %q, got %q for expression", exp, v)
	}

	return nil
}

// executeExpression will execute the given input with a given listener.
// If the input has no expression, it will return `false` for the expression return
func executeExpression(l parser.WacListener, input string) (expression bool, err error) {
	vsI := TemplateRegex.FindAllStringIndex(input, -1)
	if len(vsI) == 0 { // No parameterization
		return false, nil
	}

	// TODO: @emyrk is this an alright rule?
	if len(vsI) > 1 {
		return false, xerrors.Errorf("only allowed 1 expression %q per value", exampleTemplate)
	}

	inExp := input[vsI[0][0]:vsI[0][1]]
	input = strings.TrimSpace(input)
	if len(input) != len(inExp) {
		return false, xerrors.Errorf("if using %q, nothing can be outside of it", exampleTemplate)
	}

	inExp = strings.TrimPrefix(inExp, StartTemplate)
	inExp = strings.TrimSuffix(inExp, EndTemplate)

	p, errListener := setupInput(inExp)
	var _ = errListener

	defer func() {
		perr := recover()
		if len(errListener.Errors) > 0 {
			err = xerrors.Errorf("expression %q has syntax errors:\n%s", inExp, errListener.Error())
		} else if perr != nil {
			if pe, ok := perr.(*parseError); ok {
				err = pe
			} else {
				err = xerrors.Errorf("panic: %v", perr)
			}
		}
	}()
	// Walk executes the full lexing, parsing, and ast traverse.
	// Doing them all at once is the most optimal, but the process
	// can fail from a syntax error, or during parsing, or during the ast
	// traverse. The panic recover handles deciphering what went wrong if
	// something failed
	antlr.ParseTreeWalkerDefault.Walk(l, p.Start())
	if err != nil {
		return false, err
	}
	return true, nil
}
