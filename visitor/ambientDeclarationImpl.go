package visitor

import "ambiencalculus/parser"

func (v *TreeShapeListener) EnterAmbientDeclaration(ctx *parser.AmbientDeclarationContext) {
	label := "ambient " + ctx.AMBIENTID().GetText()
	nodeId := v.AddNode(label)
	if len(v.NodeStack) > 0 {
		parentId := v.NodeStack[len(v.NodeStack)-1]
		v.AddEdge(parentId, nodeId)
	}
	v.NodeStack = append(v.NodeStack, nodeId)

	if ctx.AMBIENTID() != nil {
		envName := ctx.AMBIENTID().GetText()
		parentName := ""

		if len(v.CurrentAmbients) > 0 {
			parentName = v.CurrentAmbients[len(v.CurrentAmbients)-1]
		}

		for _, env := range v.Environments {
			if env.Name == envName {
				panic("No se pueden agregar ambientes con el mismo nombre: " + envName)
			}
		}

		v.Environments = append(v.Environments, Environment{Name: envName, AmbientParent: parentName})
		v.CurrentAmbients = append(v.CurrentAmbients, envName)
	}
}

func (v *TreeShapeListener) ExitAmbientDeclaration(ctx *parser.AmbientDeclarationContext) {
	v.NodeStack = v.NodeStack[:len(v.NodeStack)-1]
	if len(v.CurrentAmbients) > 0 {
		v.CurrentAmbients = v.CurrentAmbients[:len(v.CurrentAmbients)-1]
	}
}
