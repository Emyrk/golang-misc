package paramlang

import (
	"golang.org/x/xerrors"

	"coder.com/m/product/coder/pkg/env/envtemplate/paramlang/parser"
)

// ExitLogical handles `and`, `or`
func (s *wacListener) ExitLogical(ctx *parser.LogicalContext) {
	right, left := s.popBool(), s.popBool()
	op := ctx.GetOp().GetText()
	switch op {
	case "and":
		s.pushBool(left && right)
	case "or":
		s.pushBool(left || right)
	default:
		s.stopParsing(xerrors.Errorf("logical operator %q not supported", op))
	}
}

// ExitComparator handles math comparators like `<`, `>`, `==`, etc
func (s *wacListener) ExitComparator(ctx *parser.ComparatorContext) {
	right, left := s.pop(), s.pop()
	// conformNumbers will ensure both left and right are both the same numerical
	// type (int or float). We can only compare like types.
	ty, err := s.conformNumbers(left, right)
	if err != nil {
		s.stopParsing(xerrors.Errorf("math on different types: %w", err))
	}

	op := ctx.GetOp().GetText()
	switch ty {
	case stInt: // Compare ints with the integer function
		s.pushBool(s.compareInts(op, left, right))
	case stFloat: // Compare floats with the float function
		s.pushBool(s.compareFloats(op, left, right))
	default:
		s.stopParsing(xerrors.Errorf("cannot do binary comparator on type %q", ty))
	}
}

// compareFloats compares two float stack items.
func (s *wacListener) compareFloats(op string, ai, bi *stackItem) bool {
	as := func(si *stackItem) float64 {
		v, err := si.AsFloat()
		if err != nil {
			s.stopParsing(xerrors.Errorf("binary op=%q, v=%v: %w", op, si.Value, err))
		}

		return v
	}

	a, b := as(ai), as(bi)

	switch op {
	case "==":
		return a == b
	case "!=":
		return a != b
	case "<":
		return a < b
	case "<=":
		return a <= b
	case ">":
		return a > b
	case ">=":
		return a >= b
	default:
		s.stopParsing(xerrors.Errorf("binary operand not supported %q", op))
	}

	return false
}

// compareInts compares two float stack items.
func (s *wacListener) compareInts(op string, ai, bi *stackItem) bool {
	as := func(si *stackItem) int {
		v, err := si.AsInt()
		if err != nil {
			s.stopParsing(xerrors.Errorf("binary op=%q, v=%v: %w", op, si.Value, err))
		}

		return v
	}

	a, b := as(ai), as(bi)

	switch op {
	case "==":
		return a == b
	case "!=":
		return a != b
	case "<":
		return a < b
	case "<=":
		return a <= b
	case ">":
		return a > b
	case ">=":
		return a >= b
	default:
		s.stopParsing(xerrors.Errorf("binary operand not supported %q", op))
	}

	return false
}
