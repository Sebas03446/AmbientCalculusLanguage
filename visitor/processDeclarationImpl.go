package visitor

import "ambiencalculus/parser"

func (v *TreeShapeListener) EnterProcessDeclaration(ctx *parser.ProcessDeclarationContext) {
	label := "process " + ctx.ID().GetText()
	nodeId := v.AddNode(label)
	if len(v.NodeStack) > 0 {
		parentId := v.NodeStack[len(v.NodeStack)-1]
		v.AddEdge(parentId, nodeId)
	}
	v.NodeStack = append(v.NodeStack, nodeId)

	for _, process := range v.Processes {
		if process.Name == ctx.ID().GetText() {
			panic("No se pueden agregar procesos con el mismo nombre: " + ctx.ID().GetText())
		}
	}
	if ctx.ID() != nil {
		processName := ctx.ID().GetText()
		newProcess := Process{Name: processName, Variables: make(map[string]interface{}), Ambient: v.CurrentAmbients[len(v.CurrentAmbients)-1]}
		v.Processes = append(v.Processes, newProcess)
		v.CurrentProcess = processName
		for i, env := range v.Environments {
			if env.Name == v.CurrentAmbients[len(v.CurrentAmbients)-1] {
				v.Environments[i].Processes = append(v.Environments[i].Processes, processName)
				break
			}
		}
	}
}

func (v *TreeShapeListener) ExitProcessDeclaration(ctx *parser.ProcessDeclarationContext) {
	v.NodeStack = v.NodeStack[:len(v.NodeStack)-1]
	v.CurrentProcess = ""
}
