package visitor

import (
	"ambiencalculus/parser"
	"fmt"
)

func (v *TreeShapeListener) EnterInStatement(ctx *parser.InStatementContext) {
	label := "in " + ctx.AMBIENTID().GetText()
	nodeId := v.AddNode(label)
	if len(v.NodeStack) > 0 {
		parentId := v.NodeStack[len(v.NodeStack)-1]
		v.AddEdge(parentId, nodeId)
	}
	v.NodeStack = append(v.NodeStack, nodeId)

	if ctx.AMBIENTID() != nil {
		destAmbient := ctx.AMBIENTID().GetText()
		fmt.Println("Entrando al ambiente:", destAmbient)
		currentAmbient := v.CurrentAmbients[len(v.CurrentAmbients)-1]

		v.PostActions = append(v.PostActions, func() {
			if currentAmbient == destAmbient {
				fmt.Println("Ambiente actual ya es el destino")
				return
			}
			for i, env := range v.Environments {
				if env.Name == currentAmbient {
					v.Environments[i].AmbientParent = destAmbient
					break
				}
			}

			fmt.Printf("El ambiente %s ha entrado exitosamente al ambiente %s\n", currentAmbient, destAmbient)
		})
	}
}

func (v *TreeShapeListener) ExitInStatement(ctx *parser.InStatementContext) {
	v.NodeStack = v.NodeStack[:len(v.NodeStack)-1]
}
