package visitor

import "ambiencalculus/parser"

func (v *TreeShapeListener) EnterModifyConditionStatement(ctx *parser.ModifyConditionStatementContext) {
	label := "modifyc " + ctx.ID().GetText() + ctx.GetOp().GetText() + ctx.Expression().GetText()
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
		result := v.evaluateExpression(ctx.Expression().(*parser.ExpressionContext))
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
				if ctx.GetOp().GetText() == "+=" {
					c := env.Conditions[condition]
					comparator := c.Comparator
					switch comparator {
					case "<":
						if c.Value+result.(int) >= c.ValueRestriction {
							panic("No se puede modificar la condición: " + condition)
						}
					case ">":
						if c.Value+result.(int) <= c.ValueRestriction {
							panic("No se puede modificar la condición: " + condition)
						}
					case "<=":
						if c.Value+result.(int) > c.ValueRestriction {
							panic("No se puede modificar la condición: " + condition)
						}
					case ">=":
						if c.Value+result.(int) < c.ValueRestriction {
							panic("No se puede modificar la condición: " + condition)
						}
					case "==":
						if c.Value+result.(int) != c.ValueRestriction {
							panic("No se puede modificar la condición: " + condition)
						}
					case "!=":
						if c.Value+result.(int) == c.ValueRestriction {
							panic("No se puede modificar la condición: " + condition)
						}
					}
					c.Value = c.Value + result.(int)
					env.Conditions[condition] = c
				} else if ctx.GetOp().GetText() == "-=" {
					c := env.Conditions[condition]
					comparator := c.Comparator
					switch comparator {
					case "<":
						if c.Value-result.(int) >= c.ValueRestriction {
							panic("No se puede modificar la condición: " + condition)
						}
					case ">":
						if c.Value-result.(int) <= c.ValueRestriction {
							panic("No se puede modificar la condición: " + condition)
						}
					case "<=":
						if c.Value-result.(int) > c.ValueRestriction {
							panic("No se puede modificar la condición: " + condition)
						}
					case ">=":
						if c.Value-result.(int) < c.ValueRestriction {
							panic("No se puede modificar la condición: " + condition)
						}
					case "==":
						if c.Value-result.(int) != c.ValueRestriction {
							panic("No se puede modificar la condición: " + condition)
						}
					case "!=":
						if c.Value-result.(int) == c.ValueRestriction {
							panic("No se puede modificar la condición: " + condition)
						}
					}
					c.Value = c.Value - result.(int)
					env.Conditions[condition] = c
				}
				break
			}
		}

		v.CurrentProcess = ""
	})

}

func (v *TreeShapeListener) ExitModifyConditionStatement(ctx *parser.ModifyConditionStatementContext) {
	v.NodeStack = v.NodeStack[:len(v.NodeStack)-1]
}
