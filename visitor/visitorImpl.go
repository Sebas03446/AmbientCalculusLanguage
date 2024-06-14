package visitor

import (
	"ambiencalculus/parser"
	"fmt"
	"strings"
)

type TreeShapeListener struct {
	*parser.BaseAmbientCalculusListener
	Environments    []Environment
	Processes       map[string]interface{}
	CurrentAmbients []string
	Nodes           []string
	Edges           []string
	NodeId          int
	NodeStack       []int
	PostActions     []func()
}

type Environment struct {
	Name          string
	Processes     []string
	AmbientParent string
}

func NewAmbientCalculusVisitorImpl() *TreeShapeListener {
	return &TreeShapeListener{
		Environments: make([]Environment, 0),
		Processes:    make(map[string]interface{}),
		Nodes:        make([]string, 0),
		Edges:        make([]string, 0),
		NodeStack:    make([]int, 0),
		PostActions:  make([]func(), 0),
	}
}

func (v *TreeShapeListener) addNode(label string) int {
	id := v.NodeId
	v.Nodes = append(v.Nodes, fmt.Sprintf("  node%d [label=\"%s\"];", id, label))
	v.NodeId++
	return id
}

func (v *TreeShapeListener) addEdge(parent, child int) {
	v.Edges = append(v.Edges, fmt.Sprintf("  node%d -> node%d;", parent, child))
}

func (v *TreeShapeListener) EnterProgram(ctx *parser.ProgramContext) {
	rootId := v.addNode("Program")
	v.NodeStack = append(v.NodeStack, rootId)
}

func (v *TreeShapeListener) ExitProgram(ctx *parser.ProgramContext) {
	v.NodeStack = v.NodeStack[:len(v.NodeStack)-1]

	for _, action := range v.PostActions {
		action()
	}
	v.PostActions = nil
}

func (v *TreeShapeListener) ToDot() string {
	var sb strings.Builder
	sb.WriteString("digraph G {\n")
	for _, node := range v.Nodes {
		sb.WriteString(node + "\n")
	}
	for _, edge := range v.Edges {
		sb.WriteString(edge + "\n")
	}
	sb.WriteString("}\n")
	return sb.String()
}

func (v *TreeShapeListener) EnterAmbientDeclaration(ctx *parser.AmbientDeclarationContext) {
	label := "ambient " + ctx.AMBIENTID().GetText()
	nodeId := v.addNode(label)
	if len(v.NodeStack) > 0 {
		parentId := v.NodeStack[len(v.NodeStack)-1]
		v.addEdge(parentId, nodeId)
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

func (v *TreeShapeListener) EnterProcessDeclaration(ctx *parser.ProcessDeclarationContext) {
	label := "process " + ctx.ID().GetText()
	nodeId := v.addNode(label)
	if len(v.NodeStack) > 0 {
		parentId := v.NodeStack[len(v.NodeStack)-1]
		v.addEdge(parentId, nodeId)
	}
	v.NodeStack = append(v.NodeStack, nodeId)

	for key := range v.Processes {
		if key == ctx.ID().GetText() {
			panic("No se pueden agregar procesos con el mismo nombre: " + ctx.ID().GetText())
		}
	}
	if ctx.ID() != nil {
		processName := ctx.ID().GetText()
		v.Processes[processName] = nil
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
}

func (v *TreeShapeListener) EnterInStatement(ctx *parser.InStatementContext) {
	label := "in " + ctx.AMBIENTID().GetText()
	nodeId := v.addNode(label)
	if len(v.NodeStack) > 0 {
		parentId := v.NodeStack[len(v.NodeStack)-1]
		v.addEdge(parentId, nodeId)
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

func (v *TreeShapeListener) EnterOutStatement(ctx *parser.OutStatementContext) {
	label := "out"
	nodeId := v.addNode(label)
	if len(v.NodeStack) > 0 {
		parentId := v.NodeStack[len(v.NodeStack)-1]
		v.addEdge(parentId, nodeId)
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
				v.CurrentAmbients = append(v.CurrentAmbients, currentAmbient)
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

func (v *TreeShapeListener) EnterMovementStatement(ctx *parser.MovementStatementContext) {
	label := "move " + ctx.ID().GetText()
	nodeId := v.addNode(label)
	if len(v.NodeStack) > 0 {
		parentId := v.NodeStack[len(v.NodeStack)-1]
		v.addEdge(parentId, nodeId)
	}
	v.NodeStack = append(v.NodeStack, nodeId)

	processName := ctx.ID().GetText()
	ambient := ctx.AMBIENTID().GetText()
	currentAmbient := v.CurrentAmbients[len(v.CurrentAmbients)-1]

	v.PostActions = append(v.PostActions, func() {
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

				v.CurrentAmbients = append(v.CurrentAmbients, ambient)
				fmt.Printf("El proceso %s ha sido movido exitosamente al ambiente %s\n", processName, ambient)
				return
			}
		}
	})

}

func (v *TreeShapeListener) ExitMovementStatement(ctx *parser.MovementStatementContext) {
	v.NodeStack = v.NodeStack[:len(v.NodeStack)-1]
}

func (v *TreeShapeListener) EnterCommunicationStatement(ctx *parser.CommunicationStatementContext) {
	fmt.Println("Visiting communication statement")
}

func (v *TreeShapeListener) EnterExpression(ctx *parser.ExpressionContext) {
	fmt.Println("Visiting expression")
	if ctx.ID() != nil {
		fmt.Println(ctx.ID().GetText())
	} else if ctx.INT() != nil {
		fmt.Println(ctx.INT().GetText())
	} else if ctx.STRING() != nil {
		fmt.Println(ctx.STRING().GetText())
	} else if ctx.GetOp() != nil {
		left := ctx.Expression(0)
		right := ctx.Expression(1)
		switch ctx.GetOp().GetTokenType() {
		case '*':
			fmt.Println(left.GetText() + " * " + right.GetText())
		case '/':
			fmt.Println(left.GetText() + " / " + right.GetText())
		case '+':
			fmt.Println(left.GetText() + " + " + right.GetText())
		case '-':
			fmt.Println(left.GetText() + " - " + right.GetText())
		}
	}
}
