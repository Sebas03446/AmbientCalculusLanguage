package visitor

import (
	"ambiencalculus/parser"
)

func (v *TreeShapeListener) EnterOperator(ctx *parser.OperatorContext) {
	label := ctx.GetText()
	nodeId := v.AddNode(label)
	if len(v.NodeStack) > 0 {
		parentId := v.NodeStack[len(v.NodeStack)-1]
		v.AddEdge(parentId, nodeId)
	}
	v.NodeStack = append(v.NodeStack, nodeId)
}

func (v *TreeShapeListener) ExitOperator(ctx *parser.OperatorContext) {
	v.NodeStack = v.NodeStack[:len(v.NodeStack)-1]
}
