// Code generated from Wac.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // Wac

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseWacListener is a complete listener for a parse tree produced by WacParser.
type BaseWacListener struct{}

var _ WacListener = &BaseWacListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseWacListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseWacListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseWacListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseWacListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStart is called when production start is entered.
func (s *BaseWacListener) EnterStart(ctx *StartContext) {}

// ExitStart is called when production start is exited.
func (s *BaseWacListener) ExitStart(ctx *StartContext) {}

// EnterKey is called when production key is entered.
func (s *BaseWacListener) EnterKey(ctx *KeyContext) {}

// ExitKey is called when production key is exited.
func (s *BaseWacListener) ExitKey(ctx *KeyContext) {}

// EnterIntegerLit is called when production IntegerLit is entered.
func (s *BaseWacListener) EnterIntegerLit(ctx *IntegerLitContext) {}

// ExitIntegerLit is called when production IntegerLit is exited.
func (s *BaseWacListener) ExitIntegerLit(ctx *IntegerLitContext) {}

// EnterBoolLit is called when production BoolLit is entered.
func (s *BaseWacListener) EnterBoolLit(ctx *BoolLitContext) {}

// ExitBoolLit is called when production BoolLit is exited.
func (s *BaseWacListener) ExitBoolLit(ctx *BoolLitContext) {}

// EnterFloatLit is called when production FloatLit is entered.
func (s *BaseWacListener) EnterFloatLit(ctx *FloatLitContext) {}

// ExitFloatLit is called when production FloatLit is exited.
func (s *BaseWacListener) ExitFloatLit(ctx *FloatLitContext) {}

// EnterStringLit is called when production StringLit is entered.
func (s *BaseWacListener) EnterStringLit(ctx *StringLitContext) {}

// ExitStringLit is called when production StringLit is exited.
func (s *BaseWacListener) ExitStringLit(ctx *StringLitContext) {}

// EnterExpStmt is called when production ExpStmt is entered.
func (s *BaseWacListener) EnterExpStmt(ctx *ExpStmtContext) {}

// ExitExpStmt is called when production ExpStmt is exited.
func (s *BaseWacListener) ExitExpStmt(ctx *ExpStmtContext) {}

// EnterIfStmt is called when production IfStmt is entered.
func (s *BaseWacListener) EnterIfStmt(ctx *IfStmtContext) {}

// ExitIfStmt is called when production IfStmt is exited.
func (s *BaseWacListener) ExitIfStmt(ctx *IfStmtContext) {}

// EnterLiteralExp is called when production LiteralExp is entered.
func (s *BaseWacListener) EnterLiteralExp(ctx *LiteralExpContext) {}

// ExitLiteralExp is called when production LiteralExp is exited.
func (s *BaseWacListener) ExitLiteralExp(ctx *LiteralExpContext) {}

// EnterIntCast is called when production IntCast is entered.
func (s *BaseWacListener) EnterIntCast(ctx *IntCastContext) {}

// ExitIntCast is called when production IntCast is exited.
func (s *BaseWacListener) ExitIntCast(ctx *IntCastContext) {}

// EnterInvertLogical is called when production InvertLogical is entered.
func (s *BaseWacListener) EnterInvertLogical(ctx *InvertLogicalContext) {}

// ExitInvertLogical is called when production InvertLogical is exited.
func (s *BaseWacListener) ExitInvertLogical(ctx *InvertLogicalContext) {}

// EnterMulDiv is called when production MulDiv is entered.
func (s *BaseWacListener) EnterMulDiv(ctx *MulDivContext) {}

// ExitMulDiv is called when production MulDiv is exited.
func (s *BaseWacListener) ExitMulDiv(ctx *MulDivContext) {}

// EnterAddSub is called when production AddSub is entered.
func (s *BaseWacListener) EnterAddSub(ctx *AddSubContext) {}

// ExitAddSub is called when production AddSub is exited.
func (s *BaseWacListener) ExitAddSub(ctx *AddSubContext) {}

// EnterLogical is called when production Logical is entered.
func (s *BaseWacListener) EnterLogical(ctx *LogicalContext) {}

// ExitLogical is called when production Logical is exited.
func (s *BaseWacListener) ExitLogical(ctx *LogicalContext) {}

// EnterFloatCast is called when production FloatCast is entered.
func (s *BaseWacListener) EnterFloatCast(ctx *FloatCastContext) {}

// ExitFloatCast is called when production FloatCast is exited.
func (s *BaseWacListener) ExitFloatCast(ctx *FloatCastContext) {}

// EnterParen is called when production Paren is entered.
func (s *BaseWacListener) EnterParen(ctx *ParenContext) {}

// ExitParen is called when production Paren is exited.
func (s *BaseWacListener) ExitParen(ctx *ParenContext) {}

// EnterKeyPath is called when production KeyPath is entered.
func (s *BaseWacListener) EnterKeyPath(ctx *KeyPathContext) {}

// ExitKeyPath is called when production KeyPath is exited.
func (s *BaseWacListener) ExitKeyPath(ctx *KeyPathContext) {}

// EnterComparator is called when production Comparator is entered.
func (s *BaseWacListener) EnterComparator(ctx *ComparatorContext) {}

// ExitComparator is called when production Comparator is exited.
func (s *BaseWacListener) ExitComparator(ctx *ComparatorContext) {}
