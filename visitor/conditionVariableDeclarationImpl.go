package visitor

import (
	"ambiencalculus/parser"
	"strconv"
)

func (v *TreeShapeListener) EnterConditionsVariableDeclaration(ctx *parser.ConditionsVariableDeclarationContext) {
	label := " letc " + ctx.ID().GetText() + " = " + ctx.INT().GetText()
	nodeId := v.AddNode(label)
	if len(v.NodeStack) > 0 {
		parentId := v.NodeStack[len(v.NodeStack)-1]
		v.AddEdge(parentId, nodeId)
	}
	v.NodeStack = append(v.NodeStack, nodeId)

	if ctx.ID() != nil {
		varName := ctx.ID().GetText()
		result, _ := strconv.Atoi(ctx.INT().GetText())

		currentAmbient := v.CurrentAmbients[len(v.CurrentAmbients)-1]

		comparator := ""
		if ctx.Restriction() != nil && ctx.Restriction().Comparator() != nil {
			comparator = ctx.Restriction().Comparator().GetText()
		}

		valueRestriction := 0
		if ctx.Restriction() != nil && ctx.Restriction().INT() != nil {
			valueRestriction, _ = strconv.Atoi(ctx.Restriction().INT().GetText())
		}

		for i, env := range v.Environments {
			if env.Name == currentAmbient {
				if _, ok := env.Conditions[varName]; ok {
					panic("No se pueden agregar condiciones con el mismo nombre: " + varName)
				}

				condition := Condition{
					Value:            result,
					Comparator:       comparator,
					ValueRestriction: valueRestriction,
				}

				env.Conditions[varName] = condition
				v.Environments[i] = env
				break
			}
		}
	}
}

func (v *TreeShapeListener) ExitConditionsVariableDeclaration(ctx *parser.ConditionsVariableDeclarationContext) {
	v.NodeStack = v.NodeStack[:len(v.NodeStack)-1]
}
