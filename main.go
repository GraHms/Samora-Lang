package main

import (
	"flag"
	"fmt"
	"github.com/EusebioSimango/samoralang/evaluator"
	"github.com/EusebioSimango/samoralang/lexer"
	"github.com/EusebioSimango/samoralang/object"
	"github.com/EusebioSimango/samoralang/parser"
	"io"
	"os"
)

func main() {

	flag.Parse()

	var input []byte
	var err error

	if len(flag.Args()) > 0 {
		input, err = os.ReadFile(os.Args[1])
	} else {
		input, err = io.ReadAll(os.Stdin)
	}

	if err != nil {
		fmt.Printf("Error reading: %s\n", err.Error())
	}

	Execute(string(input))
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
