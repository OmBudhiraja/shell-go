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
		command, err := bufio.NewReader(os.Stdin).ReadString('\n')

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// strip the newline character
		command = strings.TrimRight(command, "\n")

		evaluate(command)

	}

}

func evaluate(command string) {
	fmt.Printf("%s: command not found\n", command)
}
