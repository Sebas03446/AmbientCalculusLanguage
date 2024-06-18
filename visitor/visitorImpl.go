package visitor

import (
	"ambiencalculus/parser"
	"fmt"
	"strconv"
)

type TreeShapeListener struct {
	*parser.BaseAmbientCalculusListener
	Environments    []Environment
	Processes       []Process
	CurrentAmbients []string
	Nodes           []string
	Edges           []string
	NodeId          int
	NodeStack       []int
	PostActions     []func()
	CurrentProcess  string
}

type Environment struct {
	Name          string
	Processes     []string
	AmbientParent string
	Conditions    map[string]Condition
}

type Condition struct {
	Value            int
	Comparator       string
	ValueRestriction int
}

type Process struct {
	Name      string
	Variables map[string]interface{}
	Ambient   string
	Messages  []Message
}

type Message struct {
	Content string
	From    string
}

func NewAmbientCalculusVisitorImpl() *TreeShapeListener {
	return &TreeShapeListener{
		Environments:   make([]Environment, 0),
		Processes:      make([]Process, 0),
		Nodes:          make([]string, 0),
		Edges:          make([]string, 0),
		NodeStack:      make([]int, 0),
		PostActions:    make([]func(), 0),
		CurrentProcess: "",
	}
}

func (v *TreeShapeListener) evaluateExpression(ctx *parser.ExpressionContext) interface{} {
	if ctx.ID() != nil {
		varName := ctx.ID().GetText()
		for _, process := range v.Processes {
			if process.Name == v.CurrentProcess {
				if value, ok := process.Variables[varName]; ok {
					return value
				}
				panic("Variable no definida: " + varName)
			} else {
				panic("El proceso no existe")
			}
		}
	} else if ctx.INT() != nil {
		value, _ := strconv.Atoi(ctx.INT().GetText())
		return value
	} else if ctx.STRING() != nil {
		return ctx.STRING().GetText()
	} else if ctx.Operator() != nil {
		left := v.evaluateExpression(ctx.Expression(0).(*parser.ExpressionContext))
		right := v.evaluateExpression(ctx.Expression(1).(*parser.ExpressionContext))
		switch ctx.Operator().GetText() {
		case "+":
			return left.(int) + right.(int)
		case "-":
			return left.(int) - right.(int)
		case "*":
			return left.(int) * right.(int)
		case "/":
			return left.(int) / right.(int)
		}
	}
	return nil
}

func (v *TreeShapeListener) EnterSendStatement(ctx *parser.SendStatementContext) {
	message := ctx.STRING().GetText()
	label := "send " + message[1:len(message)-1] + " to " + ctx.ID().GetText()
	nodeId := v.AddNode(label)
	if len(v.NodeStack) > 0 {
		parentId := v.NodeStack[len(v.NodeStack)-1]
		v.AddEdge(parentId, nodeId)
	}
	v.NodeStack = append(v.NodeStack, nodeId)

	currentProcess := v.CurrentProcess
	destinationProcess := ctx.ID().GetText()
	existDestinationProcess := false

	v.PostActions = append(v.PostActions, func() {

		if currentProcess == destinationProcess {
			panic("No se puede enviar un mensaje al mismo proceso")
		}

		for i, process := range v.Processes {
			if process.Name == destinationProcess {
				existDestinationProcess = true
				process.Messages = append(process.Messages, Message{Content: message, From: currentProcess})
				v.Processes[i] = process
				break
			}
		}

		if !existDestinationProcess {
			panic("El proceso " + destinationProcess + " no existe")
		}

	})
}

func (v *TreeShapeListener) ExitSendStatement(ctx *parser.SendStatementContext) {
	v.NodeStack = v.NodeStack[:len(v.NodeStack)-1]
}

func (v *TreeShapeListener) EnterReceiveStatement(ctx *parser.ReceiveStatementContext) {
	label := "receive"
	nodeId := v.AddNode(label)
	if len(v.NodeStack) > 0 {
		parentId := v.NodeStack[len(v.NodeStack)-1]
		v.AddEdge(parentId, nodeId)
	}
	v.NodeStack = append(v.NodeStack, nodeId)

	currentProcess := v.CurrentProcess
	v.PostActions = append(v.PostActions, func() {
		for i, process := range v.Processes {
			if process.Name == currentProcess {
				if len(process.Messages) == 0 {
					fmt.Println("No hay mensajes para recibir")
				}
				fmt.Println("Mensaje recibido de:", process.Messages[0].From, "con contenido:", process.Messages[0].Content)
				process.Messages = process.Messages[1:]
				v.Processes[i] = process
			}
		}
	})
}

func (v *TreeShapeListener) ExitReceiveStatement(ctx *parser.ReceiveStatementContext) {
	v.NodeStack = v.NodeStack[:len(v.NodeStack)-1]
}
