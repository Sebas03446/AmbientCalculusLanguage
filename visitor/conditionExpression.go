package visitor

import "ambiencalculus/parser"

func (v *TreeShapeListener) EnterConditionsExpression(ctx *parser.ConditionsContext) {
	label := ctx.GetText()
	nodeId := v.AddNode(label)
	if len(v.NodeStack) > 0 {
		parentId := v.NodeStack[len(v.NodeStack)-1]
		v.AddEdge(parentId, nodeId)
	}
	v.NodeStack = append(v.NodeStack, nodeId)
}

func (v *TreeShapeListener) ExitConditionsExpression(ctx *parser.ConditionsContext) {
	v.NodeStack = v.NodeStack[:len(v.NodeStack)-1]
}
