package main

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "exit the Pockedex",
			callback:    exitCommand,
		},
		"help": {
			name:        "help",
			description: "displays a help message",
			callback:    helpCommand,
		},
		"map": {
			name:        "map",
			description: "get the next page of locations",
			callback:    mapCommand,
		},
		"mapb": {
			name:        "mapb",
			description: "get the previous page of locations",
			callback:    mapBackCommand,
		},
		"explore": {
			name:        "explore",
			description: "pokemons in the location: <explore location-name>",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "try to catch the pokemon: <catch pokemon-name>",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "show info of the pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "show all your catched pokemons",
			callback:    commandPokedex,
		},
	}
}

func exitCommand(cfg *config, args ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func helpCommand(cfg *config, args ...string) error {
	fmt.Println("\nWelcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	commands := getCommands()

	listCommands := make([]string, 0, len(commands))

	for k := range commands {
		listCommands = append(listCommands, k)
	}

	sort.Strings(listCommands)

	for _, name := range listCommands {
		cmd := commands[name]
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}

	fmt.Println()

	return nil
}

func mapCommand(cfg *config, args ...string) error {
	locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationsResp.Next
	cfg.prevLocationsURL = locationsResp.Previous

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}

	return nil
}

func mapBackCommand(cfg *config, args ...string) error {
	if cfg.prevLocationsURL == nil {
		return fmt.Errorf("you're on the first page")
	}

	locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationsResp.Next
	cfg.prevLocationsURL = locationsResp.Previous

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}

	return nil
}

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("you must provide a location name")
	}
	name := args[0]

	fmt.Printf("Exploring %s...\n", name)
	locationArea, err := cfg.pokeapiClient.GetLocationArea(name)
	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")
	for _, encounter := range locationArea.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}
	return nil
}

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("you must provide a pokemon name")
	}
	name := args[0]

	fmt.Printf("Throwing a Pokeball at %s...\n", name)

	pokemon, err := cfg.pokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}

	// Логика поимки:
	// Чем выше BaseExperience, тем меньше вероятность, что rand.Intn будет выше порога
	// Например, используем порог 40-50.
	threshold := 50
	randNum := rand.Intn(pokemon.BaseExperience)

	if randNum > threshold {
		fmt.Printf("%s escaped!\n", name)
		return nil
	}

	// Ура! Добавляем в мапу
	cfg.caughtPokemon[name] = pokemon
	fmt.Printf("%s was caught!\n", name)
	return nil
}

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("you must provide a pokemon name")
	}
	name := args[0]

	// Ищем покемона ТОЛЬКО в нашей мапе, а не в API
	pokemon, ok := cfg.caughtPokemon[name]
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	// Выводим данные
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)

	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %v\n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Println("Types:")
	for _, typeInfo := range pokemon.Types {
		fmt.Printf("  - %s\n", typeInfo.Type.Name)
	}

	return nil
}

func commandPokedex(cfg *config, args ...string) error {
	fmt.Println("Your Pokedex:")

	if len(cfg.caughtPokemon) == 0 {
		fmt.Println("Your Pokedex is empty. Go catch some!")
		return nil
	}

	for name := range cfg.caughtPokemon {
		fmt.Printf(" - %s\n", name)
	}

	return nil
}
