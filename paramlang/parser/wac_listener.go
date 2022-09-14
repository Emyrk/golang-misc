// Code generated from Wac.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // Wac

import "github.com/antlr/antlr4/runtime/Go/antlr"

// WacListener is a complete listener for a parse tree produced by WacParser.
type WacListener interface {
	antlr.ParseTreeListener

	// EnterStart is called when entering the start production.
	EnterStart(c *StartContext)

	// EnterKey is called when entering the key production.
	EnterKey(c *KeyContext)

	// EnterIntegerLit is called when entering the IntegerLit production.
	EnterIntegerLit(c *IntegerLitContext)

	// EnterBoolLit is called when entering the BoolLit production.
	EnterBoolLit(c *BoolLitContext)

	// EnterFloatLit is called when entering the FloatLit production.
	EnterFloatLit(c *FloatLitContext)

	// EnterStringLit is called when entering the StringLit production.
	EnterStringLit(c *StringLitContext)

	// EnterExpStmt is called when entering the ExpStmt production.
	EnterExpStmt(c *ExpStmtContext)

	// EnterIfStmt is called when entering the IfStmt production.
	EnterIfStmt(c *IfStmtContext)

	// EnterLiteralExp is called when entering the LiteralExp production.
	EnterLiteralExp(c *LiteralExpContext)

	// EnterIntCast is called when entering the IntCast production.
	EnterIntCast(c *IntCastContext)

	// EnterInvertLogical is called when entering the InvertLogical production.
	EnterInvertLogical(c *InvertLogicalContext)

	// EnterMulDiv is called when entering the MulDiv production.
	EnterMulDiv(c *MulDivContext)

	// EnterAddSub is called when entering the AddSub production.
	EnterAddSub(c *AddSubContext)

	// EnterLogical is called when entering the Logical production.
	EnterLogical(c *LogicalContext)

	// EnterFloatCast is called when entering the FloatCast production.
	EnterFloatCast(c *FloatCastContext)

	// EnterParen is called when entering the Paren production.
	EnterParen(c *ParenContext)

	// EnterKeyPath is called when entering the KeyPath production.
	EnterKeyPath(c *KeyPathContext)

	// EnterComparator is called when entering the Comparator production.
	EnterComparator(c *ComparatorContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitKey is called when exiting the key production.
	ExitKey(c *KeyContext)

	// ExitIntegerLit is called when exiting the IntegerLit production.
	ExitIntegerLit(c *IntegerLitContext)

	// ExitBoolLit is called when exiting the BoolLit production.
	ExitBoolLit(c *BoolLitContext)

	// ExitFloatLit is called when exiting the FloatLit production.
	ExitFloatLit(c *FloatLitContext)

	// ExitStringLit is called when exiting the StringLit production.
	ExitStringLit(c *StringLitContext)

	// ExitExpStmt is called when exiting the ExpStmt production.
	ExitExpStmt(c *ExpStmtContext)

	// ExitIfStmt is called when exiting the IfStmt production.
	ExitIfStmt(c *IfStmtContext)

	// ExitLiteralExp is called when exiting the LiteralExp production.
	ExitLiteralExp(c *LiteralExpContext)

	// ExitIntCast is called when exiting the IntCast production.
	ExitIntCast(c *IntCastContext)

	// ExitInvertLogical is called when exiting the InvertLogical production.
	ExitInvertLogical(c *InvertLogicalContext)

	// ExitMulDiv is called when exiting the MulDiv production.
	ExitMulDiv(c *MulDivContext)

	// ExitAddSub is called when exiting the AddSub production.
	ExitAddSub(c *AddSubContext)

	// ExitLogical is called when exiting the Logical production.
	ExitLogical(c *LogicalContext)

	// ExitFloatCast is called when exiting the FloatCast production.
	ExitFloatCast(c *FloatCastContext)

	// ExitParen is called when exiting the Paren production.
	ExitParen(c *ParenContext)

	// ExitKeyPath is called when exiting the KeyPath production.
	ExitKeyPath(c *KeyPathContext)

	// ExitComparator is called when exiting the Comparator production.
	ExitComparator(c *ComparatorContext)
}
