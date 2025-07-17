package main

import (
	"time"

	"github.com/Roddyck/pokedexcli/internal/pokeapi"
)

func main() {
	pokeCLient := pokeapi.NewClient(5*time.Second, time.Minute*5)
	cfg := &config{
		pokeapiClient: *pokeCLient,
	}
	startRepl(cfg)
}
