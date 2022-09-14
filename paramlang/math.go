package paramlang

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"golang.org/x/xerrors"

	"coder.com/m/product/coder/pkg/env/envtemplate/paramlang/parser"
)

type mathCtx interface {
	GetOp() antlr.Token
}

func (s *wacListener) ExitMulDiv(ctx *parser.MulDivContext) {
	s.exitMath(ctx)
}

func (s *wacListener) ExitAddSub(ctx *parser.AddSubContext) {
	s.exitMath(ctx)
}

// ExitFloatCast handles `float(1)`
func (s *wacListener) ExitFloatCast(_ *parser.FloatCastContext) {
	exp := s.pop()
	switch exp.Type {
	case stFloat:
		s.pushStackItem(exp)
	case stInt:
		i, err := exp.AsInt()
		if err != nil {
			s.stopParsing(xerrors.Errorf("float cast (from int): %w", err))
		}
		s.pushFloat(float64(i))
	case stDefer:
		f, err := exp.AsFloat()
		if err != nil {
			s.stopParsing(xerrors.Errorf("float cast: %w", err))
		}
		s.pushFloat(f)
	default:
		s.stopParsing(xerrors.Errorf("cannot cast type %q to float", exp.Type))
	}
}

// ExitIntCast handles `int(1.0)`
func (s *wacListener) ExitIntCast(_ *parser.IntCastContext) {
	exp := s.pop()
	switch exp.Type {
	case stInt:
		s.pushStackItem(exp)
	case stFloat:
		f, err := exp.AsFloat()
		if err != nil {
			s.stopParsing(xerrors.Errorf("int cast (from float): %w", err))
		}
		s.pushInt(int(f))
	case stDefer:
		i, err := exp.AsInt()
		if err != nil {
			s.stopParsing(xerrors.Errorf("int cast: %w", err))
		}
		s.pushInt(i)
	default:
		s.stopParsing(xerrors.Errorf("cannot cast type %q to int", exp.Type))
	}
}

// exitMath handles math operations on numbers.
func (s *wacListener) exitMath(ctx mathCtx) {
	right, left := s.pop(), s.pop()
	op := ctx.GetOp().GetText()

	// Math operations expect like numerical types.
	ty, err := s.conformNumbers(left, right)
	if err != nil {
		// The 1 exception is string concatenation. Try that first before failing
		if op == "+" && s.convert(false, stString, right, left) == nil {
			ls, err := left.AsString()
			if err != nil {
				s.stopParsing(xerrors.Errorf("concat string (left): %w", err))
			}

			rs, err := right.AsString()
			if err != nil {
				s.stopParsing(xerrors.Errorf("concat string (right): %w", err))
			}
			s.pushString(ls + rs)
			return
		}

		s.stopParsing(xerrors.Errorf("math on different types: %w", err))
	}

	switch ty {
	case stInt:
		s.pushInt(s.mathInts(op, left, right))
	case stFloat:
		s.pushFloat(s.mathFloats(op, left, right))
	default:
		s.stopParsing(xerrors.Errorf("cannot do math op %q on type %q", op, ty))
	}
}

func (s *wacListener) mathFloats(op string, ai, bi *stackItem) float64 {
	as := func(si *stackItem) float64 {
		v, err := si.AsFloat()
		if err != nil {
			s.stopParsing(xerrors.Errorf("math op=%q, v=%v: %w", op, si.Value, err))
		}

		return v
	}

	a, b := as(ai), as(bi)

	switch op {
	case "*":
		return a * b
	case "/":
		return a / b
	case "+":
		return a + b
	case "-":
		return a - b
	}

	s.stopParsing(xerrors.Errorf("math operand not supported %q", op))
	return -1
}

func (s *wacListener) mathInts(op string, ai, bi *stackItem) int {
	as := func(si *stackItem) int {
		v, err := si.AsInt()
		if err != nil {
			s.stopParsing(xerrors.Errorf("math op=%q, v=%v: %w", op, si.Value, err))
		}

		return v
	}

	a, b := as(ai), as(bi)

	switch op {
	case "*":
		return a * b
	case "/":
		return a / b
	case "+":
		return a + b
	case "-":
		return a - b
	}

	s.stopParsing(xerrors.Errorf("math operand not supported %q", op))
	return -1
}

// conformNumbers conforms all the items to the same numerical type
// float or int. If there is any mismatching, then no stack items
// are changed, and an error is returned
func (s *wacListener) conformNumbers(items ...*stackItem) (stackItemType, error) {
	var ty stackItemType
	for _, it := range items {
		// Only defer, int, and float can be numbers
		switch it.Type {
		case stDefer:
		case stFloat, stInt:
			if ty == "" {
				ty = it.Type
			} else if ty != it.Type {
				return stUnknown, xerrors.Errorf("cannot conform numbers of type %q and %q", ty, it.Type)
			}
		default:
			return stUnknown, xerrors.Errorf("stack item of type %q can't be number, val=%+v", it.Type, it.Value)
		}
	}

	if ty == "" {
		// All were deferred. This is awkward. Let's try all as Int's, then all as floats
		err := s.convert(false, stInt, items...)
		if err == nil {
			ty = stInt
		} else {
			err := s.convert(false, stFloat, items...)
			if err != nil {
				return stUnknown, xerrors.Errorf("can not determine type to conform to")
			}
			ty = stFloat
		}
	}

	err := s.convert(true, ty, items...)
	if err != nil {
		return stUnknown, err
	}

	return ty, nil
}

// convert converts all the items to the specified type (if possible).
func (s *wacListener) convert(write bool, ty stackItemType, items ...*stackItem) error {
	for _, it := range items {
		v, err := it.As(ty)
		if err != nil {
			return xerrors.Errorf("conforming to %q: %w", ty, err)
		}

		if write {
			it.Type = ty
			it.Value = v
		}
	}
	return nil
}
