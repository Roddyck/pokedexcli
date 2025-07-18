package main

import (
	"time"

	"github.com/Roddyck/pokedexcli/internal/pokeapi"
)

func main() {
	pokeCLient := pokeapi.NewClient(5*time.Second, time.Minute*5)
	cfg := &config{
		pokeapiClient: *pokeCLient,
		caughtPokemon: make(map[string]pokeapi.Pokemon),
	}
	startRepl(cfg)
}
