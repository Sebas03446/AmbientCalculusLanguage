// Code generated from AmbientCalculus.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // AmbientCalculus

import "github.com/antlr4-go/antlr/v4"

// AmbientCalculusListener is a complete listener for a parse tree produced by AmbientCalculusParser.
type AmbientCalculusListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterProcessDeclaration is called when entering the processDeclaration production.
	EnterProcessDeclaration(c *ProcessDeclarationContext)

	// EnterProcessStatement is called when entering the processStatement production.
	EnterProcessStatement(c *ProcessStatementContext)

	// EnterMoveStatement is called when entering the moveStatement production.
	EnterMoveStatement(c *MoveStatementContext)

	// EnterAmbientDeclaration is called when entering the ambientDeclaration production.
	EnterAmbientDeclaration(c *AmbientDeclarationContext)

	// EnterConditions is called when entering the conditions production.
	EnterConditions(c *ConditionsContext)

	// EnterMovementStatement is called when entering the movementStatement production.
	EnterMovementStatement(c *MovementStatementContext)

	// EnterCommunicationStatement is called when entering the communicationStatement production.
	EnterCommunicationStatement(c *CommunicationStatementContext)

	// EnterInStatement is called when entering the inStatement production.
	EnterInStatement(c *InStatementContext)

	// EnterOutStatement is called when entering the outStatement production.
	EnterOutStatement(c *OutStatementContext)

	// EnterOpenStatement is called when entering the openStatement production.
	EnterOpenStatement(c *OpenStatementContext)

	// EnterCallStatement is called when entering the callStatement production.
	EnterCallStatement(c *CallStatementContext)

	// EnterPrintStatement is called when entering the printStatement production.
	EnterPrintStatement(c *PrintStatementContext)

	// EnterIfStatement is called when entering the ifStatement production.
	EnterIfStatement(c *IfStatementContext)

	// EnterVariableDeclaration is called when entering the variableDeclaration production.
	EnterVariableDeclaration(c *VariableDeclarationContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitProcessDeclaration is called when exiting the processDeclaration production.
	ExitProcessDeclaration(c *ProcessDeclarationContext)

	// ExitProcessStatement is called when exiting the processStatement production.
	ExitProcessStatement(c *ProcessStatementContext)

	// ExitMoveStatement is called when exiting the moveStatement production.
	ExitMoveStatement(c *MoveStatementContext)

	// ExitAmbientDeclaration is called when exiting the ambientDeclaration production.
	ExitAmbientDeclaration(c *AmbientDeclarationContext)

	// ExitConditions is called when exiting the conditions production.
	ExitConditions(c *ConditionsContext)

	// ExitMovementStatement is called when exiting the movementStatement production.
	ExitMovementStatement(c *MovementStatementContext)

	// ExitCommunicationStatement is called when exiting the communicationStatement production.
	ExitCommunicationStatement(c *CommunicationStatementContext)

	// ExitInStatement is called when exiting the inStatement production.
	ExitInStatement(c *InStatementContext)

	// ExitOutStatement is called when exiting the outStatement production.
	ExitOutStatement(c *OutStatementContext)

	// ExitOpenStatement is called when exiting the openStatement production.
	ExitOpenStatement(c *OpenStatementContext)

	// ExitCallStatement is called when exiting the callStatement production.
	ExitCallStatement(c *CallStatementContext)

	// ExitPrintStatement is called when exiting the printStatement production.
	ExitPrintStatement(c *PrintStatementContext)

	// ExitIfStatement is called when exiting the ifStatement production.
	ExitIfStatement(c *IfStatementContext)

	// ExitVariableDeclaration is called when exiting the variableDeclaration production.
	ExitVariableDeclaration(c *VariableDeclarationContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)
}
