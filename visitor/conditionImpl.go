package visitor

import "ambiencalculus/parser"

func (v *TreeShapeListener) EnterConditions(ctx *parser.ConditionsContext) {
	label := ctx.GetText()
	nodeId := v.AddNode(label)
	if len(v.NodeStack) > 0 {
		parentId := v.NodeStack[len(v.NodeStack)-1]
		v.AddEdge(parentId, nodeId)
	}
	v.NodeStack = append(v.NodeStack, nodeId)
}

func (v *TreeShapeListener) ExitConditions(ctx *parser.ConditionsContext) {
	v.NodeStack = v.NodeStack[:len(v.NodeStack)-1]
}
