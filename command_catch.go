package main

import (
	"errors"
	"fmt"
	"math/rand"
)

// catchThreshold tunes the catch difficulty. A Pokemon is caught when a random
// roll in [0, baseExperience) lands below this value, so higher base experience
// means a lower catch chance.
const catchThreshold = 50

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("catch requires a pokemon name: catch <pokemon_name>")
	}
	name := args[0]

	fmt.Printf("Throwing a Pokeball at %s...\n", name)

	pokemon, err := cfg.pokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}

	if rand.Intn(pokemon.BaseExperience) >= catchThreshold {
		fmt.Printf("%s escaped!\n", name)
		return nil
	}

	cfg.pokedex[pokemon.Name] = pokemon
	fmt.Printf("%s was caught!\n", name)
	return nil
}
