package commands

import (
	"os"
	"fmt"
	"github.com/grahms/samoralang/evaluator"
	"github.com/grahms/samoralang/lexer"
	"github.com/grahms/samoralang/object"
	"github.com/grahms/samoralang/parser"
)

var Commands = map[string]func([]byte)  {
	"run": func(input []byte) {
		Execute(string(input))
	},
}

func Execute(input string) int {
	env := object.NewEnvironment()
	l := lexer.New(input)

	p := parser.New(l)

	program := p.ParseProgram()
	if len(p.Errors()) != 0 {
		fmt.Printf("Error parsing: %s\n", p.Errors())
		os.Exit(1)
	}
	_ = evaluator.Eval(program, env)
	return 0
}