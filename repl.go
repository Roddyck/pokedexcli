package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Roddyck/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient   pokeapi.Client
	nextLocationUrl *string
	prevLocationUrl *string
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		words := cleanInput(scanner.Text())

		command := words[0]
		if command, ok := getCommands()[command]; ok {
			err := command.callback(cfg)
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
	callback    func(cfg *config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Display a help message",
			callback:    commandHelp,
		},
		"map": {
			name: "map",
			description: "Get next page of locations",
			callback: commandMapf,
		},
		"mapb": {
			name: "mapb",
			description: "Get previous page of locations",
			callback: commandMapb,
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
