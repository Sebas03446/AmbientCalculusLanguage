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
	parentId := v.NodeStack[len(v.NodeStack)-1]
	v.addEdge(parentId, nodeId)
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
}

func (v *TreeShapeListener) EnterProcessDeclaration(ctx *parser.ProcessDeclarationContext) {
	label := "process " + ctx.ID().GetText()
	nodeId := v.addNode(label)
	parentId := v.NodeStack[len(v.NodeStack)-1]
	v.addEdge(parentId, nodeId)
	v.NodeStack = append(v.NodeStack, nodeId)

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
	parentId := v.NodeStack[len(v.NodeStack)-1]
	v.addEdge(parentId, nodeId)
	v.NodeStack = append(v.NodeStack, nodeId)

	if ctx.AMBIENTID() != nil {
		destAmbient := ctx.AMBIENTID().GetText()
		fmt.Println("Entrando al ambiente:", destAmbient)

		currentAmbient := v.CurrentAmbients[len(v.CurrentAmbients)-1]
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
	}
}

// EnterOutStatement sentencia out; se debe ver el ambiente actual y salir de él, tambien verificar si hay un ambiente padre en caso que no, imprimir que no se puede salir
func (v *TreeShapeListener) EnterOutStatement(ctx *parser.OutStatementContext) {
	label := "out"
	nodeId := v.addNode(label)
	parentId := v.NodeStack[len(v.NodeStack)-1]
	v.addEdge(parentId, nodeId)
	v.NodeStack = append(v.NodeStack, nodeId)

	currentAmbient := v.CurrentAmbients[len(v.CurrentAmbients)-1]

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
			v.CurrentAmbients = v.CurrentAmbients[:len(v.CurrentAmbients)-1]
			v.Environments[i] = env
			fmt.Printf("El ambiente %s ha salido exitosamente\n", currentAmbient)
			return
		}
	}
}

func (v *TreeShapeListener) ExitOutStatement(ctx *parser.OutStatementContext) {
	v.NodeStack = v.NodeStack[:len(v.NodeStack)-1]
}

func (v *TreeShapeListener) ExitInStatement(ctx *parser.InStatementContext) {
	v.NodeStack = v.NodeStack[:len(v.NodeStack)-1]
}

func (v *TreeShapeListener) EnterMovementStatement(ctx *parser.MovementStatementContext) {
	label := "move " + ctx.ID().GetText()
	nodeId := v.addNode(label)
	parentId := v.NodeStack[len(v.NodeStack)-1]
	v.addEdge(parentId, nodeId)
	v.NodeStack = append(v.NodeStack, nodeId)

	if ctx.ID() != nil {
		fmt.Println(ctx.ID().GetText())
	}
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
