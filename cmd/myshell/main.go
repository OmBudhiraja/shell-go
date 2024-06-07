package main

import (
	"bufio"
	"fmt"
	"os"
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

func evaluate(input string) {

	args := strings.Split(input, " ")
	command := args[0]
	args = args[1:]

	switch command {
	case "exit":
		os.Exit(0)
	case "echo":
		fmt.Println(strings.Join(args, " "))

	default:
		fmt.Printf("%s: command not found\n", command)
	}

}
