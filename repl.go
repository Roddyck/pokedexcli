package main

import (
	"fmt"
	"os"
	"strings"
	"bufio"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		words := cleanInput(scanner.Text())

		command := words[0]
		if command, ok := getCommands()[command]; ok {
			err := command.callback()
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
		}
		
	}

}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Display a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}

func cleanInput(text string) []string {
	var result []string
	for word := range strings.SplitSeq(text, " ") {
		if word == "" {
			continue
		}
		word = strings.TrimSpace(word)
		result = append(result, strings.ToLower(word))
	}
	return result
}
