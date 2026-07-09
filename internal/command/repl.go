package command

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// cleanInput lowercases text and splits it into whitespace-separated words.
func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

// StartREPL runs the read-eval-print loop, dispatching each command until the
// user exits or input ends.
func StartREPL(cfg *Config) {
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
		args := words[1:]
		cmd, ok := commands[commandName]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		if err := cmd.callback(cfg, args...); err != nil {
			fmt.Println(err)
		}
	}
}
