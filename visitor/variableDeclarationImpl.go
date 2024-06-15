package visitor

import "ambiencalculus/parser"

func (v *TreeShapeListener) EnterVariableDeclaration(ctx *parser.VariableDeclarationContext) {
	label := "let " + ctx.ID().GetText() + " = " + ctx.Expression().GetText()
	nodeId := v.AddNode(label)
	if len(v.NodeStack) > 0 {
		parentId := v.NodeStack[len(v.NodeStack)-1]
		v.AddEdge(parentId, nodeId)
	}
	v.NodeStack = append(v.NodeStack, nodeId)

	if ctx.ID() != nil {
		varName := ctx.ID().GetText()
		result := v.evaluateExpression(ctx.Expression().(*parser.ExpressionContext))

		for i, process := range v.Processes {
			if process.Name == v.CurrentProcess {
				if _, ok := process.Variables[varName]; ok {
					panic("No se pueden agregar variables con el mismo nombre: " + varName)
				}
				v.Processes[i].Variables[varName] = result
				break
			}
		}
	}
}

func (v *TreeShapeListener) ExitVariableDeclaration(ctx *parser.VariableDeclarationContext) {
	v.NodeStack = v.NodeStack[:len(v.NodeStack)-1]
}
