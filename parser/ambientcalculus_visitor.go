// Code generated from AmbientCalculus.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // AmbientCalculus

import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by AmbientCalculusParser.
type AmbientCalculusVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by AmbientCalculusParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by AmbientCalculusParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by AmbientCalculusParser#processDeclaration.
	VisitProcessDeclaration(ctx *ProcessDeclarationContext) interface{}

	// Visit a parse tree produced by AmbientCalculusParser#processStatement.
	VisitProcessStatement(ctx *ProcessStatementContext) interface{}

	// Visit a parse tree produced by AmbientCalculusParser#moveStatement.
	VisitMoveStatement(ctx *MoveStatementContext) interface{}

	// Visit a parse tree produced by AmbientCalculusParser#ambientDeclaration.
	VisitAmbientDeclaration(ctx *AmbientDeclarationContext) interface{}

	// Visit a parse tree produced by AmbientCalculusParser#conditions.
	VisitConditions(ctx *ConditionsContext) interface{}

	// Visit a parse tree produced by AmbientCalculusParser#movementStatement.
	VisitMovementStatement(ctx *MovementStatementContext) interface{}

	// Visit a parse tree produced by AmbientCalculusParser#communicationStatement.
	VisitCommunicationStatement(ctx *CommunicationStatementContext) interface{}

	// Visit a parse tree produced by AmbientCalculusParser#inStatement.
	VisitInStatement(ctx *InStatementContext) interface{}

	// Visit a parse tree produced by AmbientCalculusParser#outStatement.
	VisitOutStatement(ctx *OutStatementContext) interface{}

	// Visit a parse tree produced by AmbientCalculusParser#openStatement.
	VisitOpenStatement(ctx *OpenStatementContext) interface{}

	// Visit a parse tree produced by AmbientCalculusParser#callStatement.
	VisitCallStatement(ctx *CallStatementContext) interface{}

	// Visit a parse tree produced by AmbientCalculusParser#printStatement.
	VisitPrintStatement(ctx *PrintStatementContext) interface{}

	// Visit a parse tree produced by AmbientCalculusParser#ifStatement.
	VisitIfStatement(ctx *IfStatementContext) interface{}

	// Visit a parse tree produced by AmbientCalculusParser#variableDeclaration.
	VisitVariableDeclaration(ctx *VariableDeclarationContext) interface{}

	// Visit a parse tree produced by AmbientCalculusParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}
}
