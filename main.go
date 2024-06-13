package main

import (
	"ambiencalculus/parser"
	"ambiencalculus/visitor"
	"fmt"
	"os"
	"os/exec"

	"github.com/antlr4-go/antlr/v4"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <input-file>")
		return
	}

	input, err := antlr.NewFileStream(os.Args[1])
	if err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}
	lexer := parser.NewAmbientCalculusLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewAmbientCalculusParser(stream)

	tree := p.Program()

	ambiencalculusVisitorImpl := visitor.NewAmbientCalculusVisitorImpl()
	walker := antlr.NewParseTreeWalker()
	walker.Walk(ambiencalculusVisitorImpl, tree)

	dot := ambiencalculusVisitorImpl.ToDot()

	os.WriteFile("ast.dot", []byte(dot), 0644)

	cmd := exec.Command("dot", "-Tpng", "ast.dot", "-o", "ast.png")
	err = cmd.Run()
	if err != nil {
		fmt.Println("Error generating graph:", err)
		os.Exit(1)
	}

	fmt.Println("AST Grafo generado como ast.png")
	fmt.Println("Ambientes:", ambiencalculusVisitorImpl.Environments)
}
