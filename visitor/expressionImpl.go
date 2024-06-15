package visitor

import "ambiencalculus/parser"

func (v *TreeShapeListener) EnterExpression(ctx *parser.ExpressionContext) {
	label := ctx.GetText()
	nodeId := v.AddNode(label)
	if len(v.NodeStack) > 0 {
		parentId := v.NodeStack[len(v.NodeStack)-1]
		v.AddEdge(parentId, nodeId)
	}
	v.NodeStack = append(v.NodeStack, nodeId)
}

func (v *TreeShapeListener) ExitExpression(ctx *parser.ExpressionContext) {
	v.NodeStack = v.NodeStack[:len(v.NodeStack)-1]
}
