package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

		fmt.Println(command[:len(command)-1] + ": command not found")
	}

}
