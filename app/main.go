package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Fprint

func main() {
	for {
		fmt.Fprint(os.Stdout, "$ ")
		command, err := bufio.NewReader(os.Stdin).ReadString('\n')

		if err != nil {
			log.Fatal("Error reading command")
		}

		fmt.Println(command)

		commandParts := strings.Split(command, " ")

		if commandParts[0] == "exit" {
			i, _ := strconv.Atoi(commandParts[1])
			os.Exit(i)
		} else if commandParts[0] == "echo" {
			echoString := strings.Join(commandParts[1:], " ")
			fmt.Println(strings.TrimSuffix(echoString, "\n"))
		} else {
			fmt.Println(command[:len(command)-1] + ": command not found")
		}
	}

}
