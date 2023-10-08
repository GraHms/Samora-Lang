package main

import (
	"flag"
	"fmt"
	"github.com/grahms/samoralang/commands"
	"os"
)

func main() {

	flag.Parse()

	switch len(flag.Args()) {
	case 0:
		commands.StartREPL()	
	case 1:
		handleCommand("run", 1)
	default:
		handleCommand(os.Args[1], 2)
	}
}

func handleCommand(command string, arg_n int) {
	input_file := os.Args[arg_n]
	input, err := os.ReadFile(input_file)

	if err != nil {
		pwd, _ := os.Getwd()
		fmt.Printf("Error: can't open %s/%s\n", pwd, input_file)
		return
	}

	commands.Commands[command](input);
}
