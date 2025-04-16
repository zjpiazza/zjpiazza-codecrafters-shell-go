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
			arg := strings.TrimSuffix(commandParts[1], "\n")
			if slices.Contains(commands, arg) {
				fmt.Printf("%s is a shell builtin\n", arg)
			} else {
				fmt.Printf("%s: not found\n", arg)
			}
		} else {
			fmt.Println(command[:len(command)-1] + ": command not found\n")
		}
	}

}
