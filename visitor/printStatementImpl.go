package visitor

import (
	"ambiencalculus/parser"
	"fmt"
)

func (v *TreeShapeListener) EnterPrintStatement(ctx *parser.PrintStatementContext) {
	expression := ctx.Expression().GetText()
	label := "print " + expression
	nodeId := v.AddNode(label)
	if len(v.NodeStack) > 0 {
		parentId := v.NodeStack[len(v.NodeStack)-1]
		v.AddEdge(parentId, nodeId)
	}
	v.NodeStack = append(v.NodeStack, nodeId)

	result := v.evaluateExpression(ctx.Expression().(*parser.ExpressionContext))
	fmt.Println("Imprimiendo:", result)
}

func (v *TreeShapeListener) ExitPrintStatement(ctx *parser.PrintStatementContext) {
	v.NodeStack = v.NodeStack[:len(v.NodeStack)-1]
}
