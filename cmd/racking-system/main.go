package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/arifinoid/warehouse-rack-cli/pkg/command"
	"github.com/arifinoid/warehouse-rack-cli/pkg/rack"
)

func main() {
	if len(os.Args) > 1 {
		// file input mode
		fileName := os.Args[1]
		file, err := os.Open(fileName)
		if err != nil {
			fmt.Println("Error while opening the file:", err)
			return
		}
		defer file.Close()

		var rackSystem *rack.RackSystem
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			if line == "" {
				continue
			}
			rackSystem = command.ProcessCommand(rackSystem, line)
		}

		if err := scanner.Err(); err != nil {
			fmt.Println("Error while reading the file:", err)
		}

	} else {
		// interactive mode
		fmt.Println("Welcome to interactive mode of warehouse-rack-cli.")
		fmt.Println("Type commands, or 'exit' to quit.")

		reader := bufio.NewReader(os.Stdin)
		var rackSystem *rack.RackSystem
		for {
			fmt.Print("> ")
			line, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Error reading input:", err)
				continue
			}

			line = line[:len(line)-1]

			if line == "exit" {
				fmt.Println("Goodbye!")
				break
			}

			if line == "" {
				continue
			}
			rackSystem = command.ProcessCommand(rackSystem, line)
		}
	}
}
