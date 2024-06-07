package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type CommandType int

const (
	Builtin CommandType = iota
	Executable
)

type Command struct {
	Type    CommandType
	Handler func([]string)
}

var commands = map[string]Command{}

func main() {

	initCommands()

	for {
		fmt.Fprint(os.Stdout, "$ ")
		input, err := bufio.NewReader(os.Stdin).ReadString('\n')

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// strip the newline character
		input = strings.TrimRight(input, "\n")

		execute(input)

	}

}

func initCommands() {
	type ExternalCmd struct {
		Name string
		Path string
	}
	var externalCommand []ExternalCmd

	path := os.Getenv("PATH")
	paths := strings.Split(path, ":")

	for _, p := range paths {
		files, err := os.ReadDir(p)

		if err != nil {
			continue
		}

		for _, f := range files {
			if f.IsDir() {
				continue
			}

			externalCommand = append(externalCommand, ExternalCmd{
				Name: f.Name(),
				Path: filepath.Join(p, f.Name()),
			})
		}

	}

	registerCommand("echo", Builtin, echoCmdHandler)
	registerCommand("exit", Builtin, exitCmdHandler)
	registerCommand("type", Builtin, typeCmdHandler)

	for _, cmd := range externalCommand {
		registerCommand(cmd.Name, Executable, func(args []string) {
			executeExternalCommand(cmd.Path, args)
		})

	}

}

func executeExternalCommand(cmdPath string, args []string) {
	cmd := exec.Command(cmdPath, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()

	if err != nil {
		fmt.Println(err)
	}

}

func registerCommand(name string, cmdType CommandType, handler func([]string)) {
	commands[name] = Command{
		Type:    cmdType,
		Handler: handler,
	}
}

func execute(input string) {

	args := strings.Split(input, " ")
	command := args[0]
	args = args[1:]

	cmd, ok := commands[command]

	if !ok {
		fmt.Printf("%s: command not found\n", command)
		return
	}

	cmd.Handler(args)

}

func echoCmdHandler(args []string) {
	fmt.Println(strings.Join(args, " "))
}

func exitCmdHandler(args []string) {
	os.Exit(0)
}

func typeCmdHandler(args []string) {
	if len(args) == 0 {
		fmt.Println("type: missing argument")
		return
	}

	cmd, ok := commands[args[0]]

	if !ok {
		fmt.Printf("%s not found\n", args[0])
		return
	}

	if cmd.Type == Builtin {
		fmt.Printf("%s is a shell builtin\n", args[0])
	} else {
		fmt.Printf("%s not found\n", args[0])
	}
}
