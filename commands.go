package main

import (
	"fmt"
	"os"
)

func commandExit() error {
	fmt.Println("Closing the Pokedex...")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	message := `
	Welcome to the Pokedex!
	Usage:

	help: Display a help message
	exit: Exit the Pokedex
	`

	fmt.Println(message)
	return nil
}
