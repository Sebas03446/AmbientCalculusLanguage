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

	// EnterPrintConditionStatement is called when entering the printConditionStatement production.
	EnterPrintConditionStatement(c *PrintConditionStatementContext)

	// EnterModifyConditionStatement is called when entering the modifyConditionStatement production.
	EnterModifyConditionStatement(c *ModifyConditionStatementContext)

	// EnterConditionsBlock is called when entering the conditionsBlock production.
	EnterConditionsBlock(c *ConditionsBlockContext)

	// EnterVariableDeclaration is called when entering the variableDeclaration production.
	EnterVariableDeclaration(c *VariableDeclarationContext)

	// EnterConditionsVariableDeclaration is called when entering the conditionsVariableDeclaration production.
	EnterConditionsVariableDeclaration(c *ConditionsVariableDeclarationContext)

	// EnterRestriction is called when entering the restriction production.
	EnterRestriction(c *RestrictionContext)

	// EnterAssignmentStatement is called when entering the assignmentStatement production.
	EnterAssignmentStatement(c *AssignmentStatementContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterOperator is called when entering the operator production.
	EnterOperator(c *OperatorContext)

	// EnterComparator is called when entering the comparator production.
	EnterComparator(c *ComparatorContext)

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

	// ExitPrintConditionStatement is called when exiting the printConditionStatement production.
	ExitPrintConditionStatement(c *PrintConditionStatementContext)

	// ExitModifyConditionStatement is called when exiting the modifyConditionStatement production.
	ExitModifyConditionStatement(c *ModifyConditionStatementContext)

	// ExitConditionsBlock is called when exiting the conditionsBlock production.
	ExitConditionsBlock(c *ConditionsBlockContext)

	// ExitVariableDeclaration is called when exiting the variableDeclaration production.
	ExitVariableDeclaration(c *VariableDeclarationContext)

	// ExitConditionsVariableDeclaration is called when exiting the conditionsVariableDeclaration production.
	ExitConditionsVariableDeclaration(c *ConditionsVariableDeclarationContext)

	// ExitRestriction is called when exiting the restriction production.
	ExitRestriction(c *RestrictionContext)

	// ExitAssignmentStatement is called when exiting the assignmentStatement production.
	ExitAssignmentStatement(c *AssignmentStatementContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitOperator is called when exiting the operator production.
	ExitOperator(c *OperatorContext)

	// ExitComparator is called when exiting the comparator production.
	ExitComparator(c *ComparatorContext)
}
