package visitor

import (
	"ambiencalculus/parser"
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
