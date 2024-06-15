package visitor

import (
	"ambiencalculus/parser"
	"fmt"
)

func (v *TreeShapeListener) EnterPrintConditionStatement(ctx *parser.PrintConditionStatementContext) {
	label := "printc " + ctx.ID().GetText()
	nodeId := v.AddNode(label)
	if len(v.NodeStack) > 0 {
		parentId := v.NodeStack[len(v.NodeStack)-1]
		v.AddEdge(parentId, nodeId)
	}
	v.NodeStack = append(v.NodeStack, nodeId)
	currentProcess := v.CurrentProcess

	v.PostActions = append(v.PostActions, func() {
		v.CurrentProcess = currentProcess
		condition := ctx.ID().GetText()
		currentAmbient := ""
		for _, process := range v.Processes {
			if process.Name == currentProcess {
				currentAmbient = process.Ambient
				break
			}
		}

		for _, env := range v.Environments {
			if env.Name == currentAmbient {
				if _, ok := env.Conditions[condition]; !ok {
					panic("La condición no existe en el ambiente: " + condition)
				}
				cond := env.Conditions[condition]
				fmt.Printf("La condición %s tiene el valor %d\n", condition, cond.Value)
				break
			}
		}

		v.CurrentProcess = ""
	})
}

func (v *TreeShapeListener) ExitPrintConditionStatement(ctx *parser.PrintConditionStatementContext) {
	v.NodeStack = v.NodeStack[:len(v.NodeStack)-1]
}
