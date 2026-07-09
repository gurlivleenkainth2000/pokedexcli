package command

import (
	"time"

	"github.com/gurlivleenkainth2000/pokedexcli/internal/pokeapi"
)

// Config holds shared state passed to every command callback.
type Config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
	pokedex          map[string]pokeapi.Pokemon
}

// NewConfig returns a Config with a ready-to-use PokeAPI client and an empty
// Pokedex.
func NewConfig() *Config {
	return &Config{
		pokeapiClient: pokeapi.NewClient(5*time.Second, 5*time.Minute),
		pokedex:       map[string]pokeapi.Pokemon{},
	}
}
