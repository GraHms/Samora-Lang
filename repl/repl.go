package repl

import (
	"bufio"
	"fmt"
	"io"
	"latin/evaluator"
	"latin/lexer"
	"latin/object"
	"latin/parser"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		env := object.NewEnvironment()
		if !scanned {
			return
		}
		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)
		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}
		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			_, _ = io.WriteString(out, evaluated.Inspect())
			_, _ = io.WriteString(out, "\n")
		}

	}
}

func printParserErrors(out io.Writer, errors []string) {
	for _, msg := range errors {

		_, _ = io.WriteString(out, "\t"+msg+"\n")
	}
}
