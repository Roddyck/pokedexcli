package main

import (
	"errors"
	"fmt"
	"os"
)

func commandExit(cfg *config, args ...string) error {
	fmt.Println("Closing the Pokedex...")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config, args ...string) error {
	var message string

	message += "Welcome to the Pokedex!\nUsage:\n"

	for command, commandInfo := range getCommands() {
		message += fmt.Sprintf("  %s - %s\n", command, commandInfo.description)
	}


	fmt.Println(message)
	return nil
}

func commandMapf(cfg *config, args ...string) error {
	localionResp, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationUrl)
	if err != nil {
	    return err
	}

	cfg.nextLocationUrl = localionResp.Next
	cfg.prevLocationUrl = localionResp.Previous

	for _, loc := range localionResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandMapb(cfg *config, args ...string) error {
	if cfg.prevLocationUrl == nil {
		return errors.New("you are on the first page")
	}

	localionResp, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationUrl)
	if err != nil {
	    return err
	}

	cfg.nextLocationUrl = localionResp.Next
	cfg.prevLocationUrl = localionResp.Previous

	for _, loc := range localionResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a location name")
	}

	name := args[0]

	location, err := cfg.pokeapiClient.GetLocation(name)
	if err != nil {
	    return err
	}

	fmt.Printf("Exploring %s...\n", location.Name)
	fmt.Println("Found Pokemon: ")
	for _, enc := range location.PokemonEncounters {
		fmt.Printf(" - %s\n", enc.Pokemon.Name)
	}
	return nil
}
