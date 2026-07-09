package main

import "github.com/gurlivleenkainth2000/pokedexcli/internal/command"

func main() {
	cfg := command.NewConfig()
	command.StartREPL(cfg)
}
