package main

import (
	"flag"
	"fmt"
	"github.com/eusebiosimango/samoralang/evaluator"
	"github.com/eusebiosimango/samoralang/lexer"
	"github.com/eusebiosimango/samoralang/object"
	"github.com/eusebiosimango/samoralang/parser"
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
