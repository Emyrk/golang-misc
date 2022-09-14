package paramlang

import (
	"golang.org/x/xerrors"

	"coder.com/m/product/coder/pkg/env/envtemplate/paramlang/parser"
)

var _ parser.WacListener = &wacTypeCheck{}

type wacTypeCheck struct {
	parser.BaseWacListener
	stack []stackItemType
}

func (w *wacTypeCheck) push(ty stackItemType) {
	w.stack = append(w.stack, ty)
}

func (w *wacTypeCheck) pop() stackItemType {
	if len(w.stack) < 1 {
		panic("stack is empty unable to pop")
	}

	// Get the last value from the stack.
	result := w.stack[len(w.stack)-1]

	// Remove the last element from the stack.
	w.stack = w.stack[:len(w.stack)-1]

	return result
}

// ExitIntegerLit is called when production IntegerLit is exited.
func (w *wacTypeCheck) ExitIntegerLit(_ *parser.IntegerLitContext) {
	w.push(stInt)
}

// ExitBoolLit is called when production BoolLit is exited.
func (w *wacTypeCheck) ExitBoolLit(_ *parser.BoolLitContext) {
	w.push(stBool)
}

// ExitFloatLit is called when production FloatLit is exited.
func (w *wacTypeCheck) ExitFloatLit(_ *parser.FloatLitContext) {
	w.push(stFloat)
}

// ExitStringLit is called when production StringLit is exited.
func (w *wacTypeCheck) ExitStringLit(_ *parser.StringLitContext) {
	w.push(stString)
}

// ExitExpStmt is called when production ExpStmt is exited.
func (w *wacTypeCheck) ExitExpStmt(_ *parser.ExpStmtContext) {}

// ExitIfStmt is called when production IfStmt is exited.
func (w *wacTypeCheck) ExitIfStmt(_ *parser.IfStmtContext) {
	e, t, i := w.pop(), w.pop(), w.pop()
	w.acceptable(i, stBool, stDefer)
	w.push(w.allSame(e, t))
}

// ExitLiteralExp is called when production LiteralExp is exited.
func (w *wacTypeCheck) ExitLiteralExp(_ *parser.LiteralExpContext) {}

// ExitIntCast is called when production IntCast is exited.
func (w *wacTypeCheck) ExitIntCast(_ *parser.IntCastContext) {
	w.acceptable(w.pop(), stDefer, stInt, stFloat)
	w.push(stInt)
}

// ExitInvertLogical is called when production InvertLogical is exited.
func (w *wacTypeCheck) ExitInvertLogical(_ *parser.InvertLogicalContext) {
	w.acceptable(w.pop(), stBool, stDefer)
	w.push(stBool)
}

// ExitMulDiv is called when production MulDiv is exited.
func (w *wacTypeCheck) ExitMulDiv(_ *parser.MulDivContext) {
	r, l := w.pop(), w.pop()
	acc := []stackItemType{stInt, stFloat}
	w.acceptable(r, acc...)
	w.acceptable(l, acc...)
	w.push(w.allSame(r, l))
}

// ExitAddSub is called when production AddSub is exited.
func (w *wacTypeCheck) ExitAddSub(ctx *parser.AddSubContext) {
	r, l := w.pop(), w.pop()
	acc := []stackItemType{stInt, stFloat}
	if ctx.GetOp().GetText() == "+" {
		acc = append(acc, stString)
	}
	w.acceptable(r, acc...)
	w.acceptable(l, acc...)
	w.push(w.allSame(r, l))
}

// ExitLogical is called when production Logical is exited.
func (w *wacTypeCheck) ExitLogical(_ *parser.LogicalContext) {
	r, l := w.pop(), w.pop()
	w.acceptable(r, stBool)
	w.acceptable(l, stBool)
	w.push(stBool)
}

// ExitFloatCast is called when production FloatCast is exited.
func (w *wacTypeCheck) ExitFloatCast(_ *parser.FloatCastContext) {
	w.acceptable(w.pop(), stDefer, stInt, stFloat)
	w.push(stFloat)
}

// ExitKeyPath is called when production KeyPath is exited.
func (w *wacTypeCheck) ExitKeyPath(_ *parser.KeyPathContext) {
	w.push(stDefer)
}

// ExitComparator is called when production Comparator is exited.
func (w *wacTypeCheck) ExitComparator(_ *parser.ComparatorContext) {
	r, l := w.pop(), w.pop()
	w.allSame(r, l)
	w.push(stBool)
}

func (w *wacTypeCheck) allSame(types ...stackItemType) stackItemType {
	if len(types) == 0 {
		return stUnknown
	}

	// Default to first type
	must := types[0]
	// If there is any non-defers, use that instead
	for _, ty := range types {
		if ty != stDefer {
			must = ty
			break
		}
	}

	for _, ty := range types {
		if ty == stDefer { // Defers are wild cards
			continue
		}
		if ty != must {
			w.stopParsing(xerrors.Errorf("found mixed types, %q and %q", ty, must))
		}
	}
	return must
}

func (w *wacTypeCheck) acceptable(found stackItemType, allowed ...stackItemType) {
	for _, allow := range allowed {
		if found == allow {
			return
		}
	}
	w.stopParsing(xerrors.Errorf("type %q not in allowed %+v", found, allowed))
}

func (w *wacTypeCheck) stopParsing(err error) {
	panic(&parseError{err: err})
}
