package visitor

import "ambiencalculus/parser"

func (v *TreeShapeListener) EnterProgram(ctx *parser.ProgramContext) {
	rootId := v.AddNode("Program")
	v.NodeStack = append(v.NodeStack, rootId)
}

func (v *TreeShapeListener) ExitProgram(ctx *parser.ProgramContext) {
	v.NodeStack = v.NodeStack[:len(v.NodeStack)-1]

	for _, action := range v.PostActions {
		action()
	}
	v.PostActions = nil
}
