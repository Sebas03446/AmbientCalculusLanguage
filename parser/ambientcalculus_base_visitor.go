// Code generated from AmbientCalculus.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // AmbientCalculus

import "github.com/antlr4-go/antlr/v4"

type BaseAmbientCalculusVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseAmbientCalculusVisitor) VisitProgram(ctx *ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAmbientCalculusVisitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAmbientCalculusVisitor) VisitProcessDeclaration(ctx *ProcessDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAmbientCalculusVisitor) VisitProcessStatement(ctx *ProcessStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAmbientCalculusVisitor) VisitMoveStatement(ctx *MoveStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAmbientCalculusVisitor) VisitAmbientDeclaration(ctx *AmbientDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAmbientCalculusVisitor) VisitConditions(ctx *ConditionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAmbientCalculusVisitor) VisitMovementStatement(ctx *MovementStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAmbientCalculusVisitor) VisitSendStatement(ctx *SendStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAmbientCalculusVisitor) VisitReceiveStatement(ctx *ReceiveStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAmbientCalculusVisitor) VisitInStatement(ctx *InStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAmbientCalculusVisitor) VisitOutStatement(ctx *OutStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAmbientCalculusVisitor) VisitOpenStatement(ctx *OpenStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAmbientCalculusVisitor) VisitCallStatement(ctx *CallStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAmbientCalculusVisitor) VisitPrintStatement(ctx *PrintStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAmbientCalculusVisitor) VisitPrintConditionStatement(ctx *PrintConditionStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAmbientCalculusVisitor) VisitModifyConditionStatement(ctx *ModifyConditionStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAmbientCalculusVisitor) VisitConditionsBlock(ctx *ConditionsBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAmbientCalculusVisitor) VisitVariableDeclaration(ctx *VariableDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAmbientCalculusVisitor) VisitConditionsVariableDeclaration(ctx *ConditionsVariableDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAmbientCalculusVisitor) VisitRestriction(ctx *RestrictionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAmbientCalculusVisitor) VisitAssignmentStatement(ctx *AssignmentStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAmbientCalculusVisitor) VisitExpression(ctx *ExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAmbientCalculusVisitor) VisitOperator(ctx *OperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAmbientCalculusVisitor) VisitComparator(ctx *ComparatorContext) interface{} {
	return v.VisitChildren(ctx)
}
