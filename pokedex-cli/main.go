package main

import (
	"time"

	"github.com/Serega-D/bootdev-learning/pokedex-cli/internal/pokeapi"
)

type config struct {
	pokeapiClient    pokeapi.Client
	caughtPokemon    map[string]pokeapi.Pokemon // Твой личный Покедекс
	nextLocationsURL *string
	prevLocationsURL *string
}

func main() {

	pokeClient := pokeapi.NewClient(20*time.Second, 5*time.Minute)

	caughtPokemonMap := make(map[string]pokeapi.Pokemon)

	cfg := &config{
		pokeapiClient: pokeClient,
		caughtPokemon: caughtPokemonMap,
	}

	startRepl(cfg)
}
