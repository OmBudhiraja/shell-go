package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {

	for {
		fmt.Fprint(os.Stdout, "$ ")
		input, err := bufio.NewReader(os.Stdin).ReadString('\n')

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// strip the newline character
		input = strings.TrimRight(input, "\n")

		evaluate(input)

	}

}

var (
	builtinCommands = []string{"echo", "exit", "type"}
)

func evaluate(input string) {

	args := strings.Split(input, " ")
	command := args[0]
	args = args[1:]

	switch command {
	case "exit":
		os.Exit(0)
	case "echo":
		fmt.Println(strings.Join(args, " "))
	case "type":
		if len(args) == 0 {
			fmt.Println("type: missing argument")
			return
		}

		if slices.Contains(builtinCommands, args[0]) {
			fmt.Printf("%s is a shell builtin\n", args[0])
		} else {
			fmt.Printf("%s not found\n", args[0])
		}

	default:
		fmt.Printf("%s: command not found\n", command)
	}

}
