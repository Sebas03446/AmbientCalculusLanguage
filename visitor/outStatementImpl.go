package visitor

import (
	"ambiencalculus/parser"
	"fmt"
)

func (v *TreeShapeListener) EnterOutStatement(ctx *parser.OutStatementContext) {
	label := "out"
	nodeId := v.AddNode(label)
	if len(v.NodeStack) > 0 {
		parentId := v.NodeStack[len(v.NodeStack)-1]
		v.AddEdge(parentId, nodeId)
	}
	v.NodeStack = append(v.NodeStack, nodeId)

	currentAmbient := v.CurrentAmbients[len(v.CurrentAmbients)-1]

	v.PostActions = append(v.PostActions, func() {
		for i, env := range v.Environments {
			if env.Name == currentAmbient {
				if env.AmbientParent == "" {
					fmt.Println("No se puede salir del ambiente actual, ya que no tiene un ambiente padre")
					return
				}

				parentOfParentAmbient := ""
				for _, env2 := range v.Environments {
					if env2.Name == env.AmbientParent {
						parentOfParentAmbient = env2.AmbientParent
						break
					}
				}

				env.AmbientParent = parentOfParentAmbient
				v.CurrentAmbients = append(v.CurrentAmbients, env.AmbientParent)
				v.Environments[i] = env
				fmt.Printf("El ambiente %s ha salido exitosamente\n", currentAmbient)
				return
			}
		}
	})
}

func (v *TreeShapeListener) ExitOutStatement(ctx *parser.OutStatementContext) {
	v.NodeStack = v.NodeStack[:len(v.NodeStack)-1]
}
