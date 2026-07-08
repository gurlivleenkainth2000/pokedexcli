package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/gurlivleenkainth2000/pokedexcli/internal/pokeapi"
)

func main() {
	cfg := &config{
		pokeapiClient: pokeapi.NewClient(5*time.Second, 5*time.Minute),
	}

	scanner := bufio.NewScanner(os.Stdin)
	commands := getCommands()
	for {
		fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			break
		}
		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}
		commandName := words[0]
		command, ok := commands[commandName]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		if err := command.callback(cfg); err != nil {
			fmt.Println(err)
		}
	}
}
