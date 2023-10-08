package commands

import (
	"os"
	"bufio"
	"github.com/grahms/samoralang/repl"
)

func StartREPL() {
		in := bufio.NewReader(os.Stdin)
  	out := os.Stdout
		repl.Start(in, out)
}