package main

import (
	"errors"
	"fmt"
	"os"
)

func commandExit(cfg *config) error {
	fmt.Println("Closing the Pokedex...")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config) error {
	var message string

	message += "Welcome to the Pokedex!\nUsage:\n"

	for command, commandInfo := range getCommands() {
		message += fmt.Sprintf("  %s - %s\n", command, commandInfo.description)
	}


	fmt.Println(message)
	return nil
}

func commandMapf(cfg *config) error {
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

func commandMapb(cfg *config) error {
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
