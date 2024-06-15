package visitor

import "ambiencalculus/parser"

func (v *TreeShapeListener) EnterConditionsBlock(ctx *parser.ConditionsBlockContext) {
	label := "conditions"
	nodeId := v.AddNode(label)
	if len(v.NodeStack) > 0 {
		parentId := v.NodeStack[len(v.NodeStack)-1]
		v.AddEdge(parentId, nodeId)
	}
	v.NodeStack = append(v.NodeStack, nodeId)

	currentAmbient := v.CurrentAmbients[len(v.CurrentAmbients)-1]
	for _, env := range v.Environments {
		if env.Name == currentAmbient {
			env.Conditions = make(map[string]Condition)
			v.Environments = append(v.Environments[:len(v.Environments)-1], env)
		}
	}
}

func (v *TreeShapeListener) ExitConditionsBlock(ctx *parser.ConditionsBlockContext) {
	v.NodeStack = v.NodeStack[:len(v.NodeStack)-1]
}
