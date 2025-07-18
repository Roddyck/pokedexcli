package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func commandExit(cfg *config, args ...string) error {
	fmt.Println("Closing the Pokedex...")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config, args ...string) error {
	var message string

	message += "Welcome to the Pokedex!\nUsage:\n"

	for _, commandInfo := range getCommands() {
		message += fmt.Sprintf("  %s - %s\n", commandInfo.name, commandInfo.description)
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

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	name := args[0]

	pokemon, err := cfg.pokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", name)
	time.Sleep(time.Millisecond * 500)
	catchChange := calculateCatchChance(pokemon.BaseExperience)

	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)

	if rng.Float64() < catchChange {
		cfg.caughtPokemon[name] = pokemon
		fmt.Printf("%s was caught!\n", name)
	} else {
		fmt.Printf("%s escaped!\n", name)
	}

	return nil
}

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	name := args[0]

	pokemon, ok := cfg.caughtPokemon[name]

	if !ok {
		return errors.New("you have not caught that pokemon")
	}

	fmt.Println("Name:", pokemon.Name)
	fmt.Println("Height:", pokemon.Height)
	fmt.Println("Weight:", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf(" -%s: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, typeInfo := range pokemon.Types {
		fmt.Println(" -", typeInfo.Type.Name)
	}
	
	return nil
}

func commandPokedex(cfg *config, args ...string) error {
	fmt.Println("Your Pokedex:")
	for name := range cfg.caughtPokemon {
		fmt.Println(" -", name)
	}

	return nil
}

func calculateCatchChance(baseExperience int) float64 {
baseChance := 0.8
	scalingFactor := 0.005
	probability := baseChance - (float64(baseExperience) * scalingFactor)

	minProb := 0.05

	if probability < minProb {
		probability = minProb
	}
	if probability > baseChance {
		probability = baseChance
	}

	return probability
}
