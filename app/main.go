package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Fprint

func main() {
	commands := []string{"exit", "echo", "type"}
	for {
		fmt.Fprint(os.Stdout, "$ ")
		command, err := bufio.NewReader(os.Stdin).ReadString('\n')

		if err != nil {
			log.Fatal("Error reading command")
		}

		commandParts := strings.Split(command, " ")

		if commandParts[0] == "exit" {
			i, _ := strconv.Atoi(commandParts[1])
			os.Exit(i)
		} else if commandParts[0] == "echo" {
			echoString := strings.Join(commandParts[1:], " ")
			fmt.Println(strings.TrimSuffix(echoString, "\n"))
		} else if commandParts[0] == "type" {
			// Check for builtin commands
			var builtin bool
			var executable string
			arg := strings.TrimSuffix(commandParts[1], "\n")
			if slices.Contains(commands, arg) {
				builtin = true
			}

			// Check for executables in path
			// 1. Get PATH environment variable
			path := os.Getenv("PATH")
			// 2. Break string into directories
			directories := strings.Split(path, ":")
			// 3. Iterate over each directory in the path and search for argument
			// fmt.Print(path)
			for _, directory := range directories {
				commandPath := fmt.Sprintf("%s/%s", directory, arg)
				if _, err := os.Stat(commandPath); !os.IsNotExist(err) {
					executable = commandPath
					break
				}
			}

			// Fallthrough
			if builtin {
				fmt.Printf("%s is a shell builtin\n", arg)
			} else if executable != "" {
				fmt.Printf("%s is %s\n", arg, executable)
			} else {
				fmt.Printf("%s: not found\n", arg)
			}

		} else {
			fmt.Println(command[:len(command)-1] + ": command not found")
		}
	}

}
