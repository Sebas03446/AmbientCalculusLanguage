// Code generated from AmbientCalculus.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // AmbientCalculus

import "github.com/antlr4-go/antlr/v4"

// BaseAmbientCalculusListener is a complete listener for a parse tree produced by AmbientCalculusParser.
type BaseAmbientCalculusListener struct{}

var _ AmbientCalculusListener = &BaseAmbientCalculusListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseAmbientCalculusListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseAmbientCalculusListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseAmbientCalculusListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseAmbientCalculusListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseAmbientCalculusListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseAmbientCalculusListener) ExitProgram(ctx *ProgramContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseAmbientCalculusListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseAmbientCalculusListener) ExitStatement(ctx *StatementContext) {}

// EnterProcessDeclaration is called when production processDeclaration is entered.
func (s *BaseAmbientCalculusListener) EnterProcessDeclaration(ctx *ProcessDeclarationContext) {}

// ExitProcessDeclaration is called when production processDeclaration is exited.
func (s *BaseAmbientCalculusListener) ExitProcessDeclaration(ctx *ProcessDeclarationContext) {}

// EnterProcessStatement is called when production processStatement is entered.
func (s *BaseAmbientCalculusListener) EnterProcessStatement(ctx *ProcessStatementContext) {}

// ExitProcessStatement is called when production processStatement is exited.
func (s *BaseAmbientCalculusListener) ExitProcessStatement(ctx *ProcessStatementContext) {}

// EnterMoveStatement is called when production moveStatement is entered.
func (s *BaseAmbientCalculusListener) EnterMoveStatement(ctx *MoveStatementContext) {}

// ExitMoveStatement is called when production moveStatement is exited.
func (s *BaseAmbientCalculusListener) ExitMoveStatement(ctx *MoveStatementContext) {}

// EnterAmbientDeclaration is called when production ambientDeclaration is entered.
func (s *BaseAmbientCalculusListener) EnterAmbientDeclaration(ctx *AmbientDeclarationContext) {}

// ExitAmbientDeclaration is called when production ambientDeclaration is exited.
func (s *BaseAmbientCalculusListener) ExitAmbientDeclaration(ctx *AmbientDeclarationContext) {}

// EnterConditions is called when production conditions is entered.
func (s *BaseAmbientCalculusListener) EnterConditions(ctx *ConditionsContext) {}

// ExitConditions is called when production conditions is exited.
func (s *BaseAmbientCalculusListener) ExitConditions(ctx *ConditionsContext) {}

// EnterMovementStatement is called when production movementStatement is entered.
func (s *BaseAmbientCalculusListener) EnterMovementStatement(ctx *MovementStatementContext) {}

// ExitMovementStatement is called when production movementStatement is exited.
func (s *BaseAmbientCalculusListener) ExitMovementStatement(ctx *MovementStatementContext) {}

// EnterSendStatement is called when production sendStatement is entered.
func (s *BaseAmbientCalculusListener) EnterSendStatement(ctx *SendStatementContext) {}

// ExitSendStatement is called when production sendStatement is exited.
func (s *BaseAmbientCalculusListener) ExitSendStatement(ctx *SendStatementContext) {}

// EnterReceiveStatement is called when production receiveStatement is entered.
func (s *BaseAmbientCalculusListener) EnterReceiveStatement(ctx *ReceiveStatementContext) {}

// ExitReceiveStatement is called when production receiveStatement is exited.
func (s *BaseAmbientCalculusListener) ExitReceiveStatement(ctx *ReceiveStatementContext) {}

// EnterInStatement is called when production inStatement is entered.
func (s *BaseAmbientCalculusListener) EnterInStatement(ctx *InStatementContext) {}

// ExitInStatement is called when production inStatement is exited.
func (s *BaseAmbientCalculusListener) ExitInStatement(ctx *InStatementContext) {}

// EnterOutStatement is called when production outStatement is entered.
func (s *BaseAmbientCalculusListener) EnterOutStatement(ctx *OutStatementContext) {}

// ExitOutStatement is called when production outStatement is exited.
func (s *BaseAmbientCalculusListener) ExitOutStatement(ctx *OutStatementContext) {}

// EnterOpenStatement is called when production openStatement is entered.
func (s *BaseAmbientCalculusListener) EnterOpenStatement(ctx *OpenStatementContext) {}

// ExitOpenStatement is called when production openStatement is exited.
func (s *BaseAmbientCalculusListener) ExitOpenStatement(ctx *OpenStatementContext) {}

// EnterCallStatement is called when production callStatement is entered.
func (s *BaseAmbientCalculusListener) EnterCallStatement(ctx *CallStatementContext) {}

// ExitCallStatement is called when production callStatement is exited.
func (s *BaseAmbientCalculusListener) ExitCallStatement(ctx *CallStatementContext) {}

// EnterPrintStatement is called when production printStatement is entered.
func (s *BaseAmbientCalculusListener) EnterPrintStatement(ctx *PrintStatementContext) {}

// ExitPrintStatement is called when production printStatement is exited.
func (s *BaseAmbientCalculusListener) ExitPrintStatement(ctx *PrintStatementContext) {}

// EnterPrintConditionStatement is called when production printConditionStatement is entered.
func (s *BaseAmbientCalculusListener) EnterPrintConditionStatement(ctx *PrintConditionStatementContext) {
}

// ExitPrintConditionStatement is called when production printConditionStatement is exited.
func (s *BaseAmbientCalculusListener) ExitPrintConditionStatement(ctx *PrintConditionStatementContext) {
}

// EnterModifyConditionStatement is called when production modifyConditionStatement is entered.
func (s *BaseAmbientCalculusListener) EnterModifyConditionStatement(ctx *ModifyConditionStatementContext) {
}

// ExitModifyConditionStatement is called when production modifyConditionStatement is exited.
func (s *BaseAmbientCalculusListener) ExitModifyConditionStatement(ctx *ModifyConditionStatementContext) {
}

// EnterConditionsBlock is called when production conditionsBlock is entered.
func (s *BaseAmbientCalculusListener) EnterConditionsBlock(ctx *ConditionsBlockContext) {}

// ExitConditionsBlock is called when production conditionsBlock is exited.
func (s *BaseAmbientCalculusListener) ExitConditionsBlock(ctx *ConditionsBlockContext) {}

// EnterVariableDeclaration is called when production variableDeclaration is entered.
func (s *BaseAmbientCalculusListener) EnterVariableDeclaration(ctx *VariableDeclarationContext) {}

// ExitVariableDeclaration is called when production variableDeclaration is exited.
func (s *BaseAmbientCalculusListener) ExitVariableDeclaration(ctx *VariableDeclarationContext) {}

// EnterConditionsVariableDeclaration is called when production conditionsVariableDeclaration is entered.
func (s *BaseAmbientCalculusListener) EnterConditionsVariableDeclaration(ctx *ConditionsVariableDeclarationContext) {
}

// ExitConditionsVariableDeclaration is called when production conditionsVariableDeclaration is exited.
func (s *BaseAmbientCalculusListener) ExitConditionsVariableDeclaration(ctx *ConditionsVariableDeclarationContext) {
}

// EnterRestriction is called when production restriction is entered.
func (s *BaseAmbientCalculusListener) EnterRestriction(ctx *RestrictionContext) {}

// ExitRestriction is called when production restriction is exited.
func (s *BaseAmbientCalculusListener) ExitRestriction(ctx *RestrictionContext) {}

// EnterAssignmentStatement is called when production assignmentStatement is entered.
func (s *BaseAmbientCalculusListener) EnterAssignmentStatement(ctx *AssignmentStatementContext) {}

// ExitAssignmentStatement is called when production assignmentStatement is exited.
func (s *BaseAmbientCalculusListener) ExitAssignmentStatement(ctx *AssignmentStatementContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseAmbientCalculusListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseAmbientCalculusListener) ExitExpression(ctx *ExpressionContext) {}

// EnterOperator is called when production operator is entered.
func (s *BaseAmbientCalculusListener) EnterOperator(ctx *OperatorContext) {}

// ExitOperator is called when production operator is exited.
func (s *BaseAmbientCalculusListener) ExitOperator(ctx *OperatorContext) {}

// EnterComparator is called when production comparator is entered.
func (s *BaseAmbientCalculusListener) EnterComparator(ctx *ComparatorContext) {}

// ExitComparator is called when production comparator is exited.
func (s *BaseAmbientCalculusListener) ExitComparator(ctx *ComparatorContext) {}
