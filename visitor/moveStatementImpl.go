package visitor

import (
	"ambiencalculus/parser"
	"fmt"
)

func (v *TreeShapeListener) EnterMovementStatement(ctx *parser.MovementStatementContext) {
	label := "move " + ctx.ID().GetText()
	nodeId := v.AddNode(label)
	if len(v.NodeStack) > 0 {
		parentId := v.NodeStack[len(v.NodeStack)-1]
		v.AddEdge(parentId, nodeId)
	}
	v.NodeStack = append(v.NodeStack, nodeId)

	processName := ctx.ID().GetText()
	ambient := ctx.AMBIENTID().GetText()
	currentAmbient := v.CurrentAmbients[len(v.CurrentAmbients)-1]
	v.PostActions = append(v.PostActions, func() {
		isNewCurrentAmbient := len(v.CurrentAmbients) > 0 && v.CurrentAmbients[len(v.CurrentAmbients)-1] != ambient
		if isNewCurrentAmbient {
			currentAmbient = v.CurrentAmbients[len(v.CurrentAmbients)-1]
		}

		if currentAmbient == ambient {
			fmt.Println("Ambiente actual ya es el destino")
			return
		}

		for i, env := range v.Environments {
			if env.Name == ambient {
				for _, process := range env.Processes {
					if process == processName {
						fmt.Println("El proceso ya se encuentra en el ambiente actual")
						return
					}
				}

				env.Processes = append(env.Processes, processName)
				v.Environments[i] = env

				for j, env2 := range v.Environments {
					if env2.Name == currentAmbient {
						for k, process := range env2.Processes {
							if process == processName {
								env2.Processes = append(env2.Processes[:k], env2.Processes[k+1:]...)
								v.Environments[j] = env2
								break
							}
						}
						break
					}
				}

				for j, process := range v.Processes {
					if process.Name == processName {
						v.Processes[j].Ambient = ambient
						break
					}
				}

				v.CurrentAmbients = append(v.CurrentAmbients, ambient)
				fmt.Printf("El proceso %s ha sido movido exitosamente al ambiente %s\n", processName, ambient)
				return
			}
		}

		fmt.Println("El ambiente destino no existe")
	})

}

func (v *TreeShapeListener) ExitMovementStatement(ctx *parser.MovementStatementContext) {
	v.NodeStack = v.NodeStack[:len(v.NodeStack)-1]
}
