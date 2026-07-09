# Pokedex CLI

A REPL-based Pokédex you drive from the terminal. Explore the Pokémon world, catch Pokémon, and inspect the ones you've caught, all backed by the [PokeAPI](https://pokeapi.co/) with an in-memory response cache.

Built as part of the [Boot.dev](https://boot.dev) "Build a Pokedex" course.

## Features

- Interactive REPL prompt (`Pokedex >`)
- Paginated exploration of location areas (`map` / `mapb`)
- List the Pokémon found in a location area (`explore`)
- Catch Pokémon with odds weighted by their base experience (`catch`)
- Inspect the stats of Pokémon you've caught (`inspect`)
- Review your collection (`pokedex`)
- A thread-safe cache that reaps stale entries, so repeated lookups are instant

## Requirements

- [Go](https://go.dev/dl/) 1.26 or newer

## Install and run

```bash
git clone https://github.com/gurlivleenkainth2000/pokedexcli.git
cd pokedexcli
go run .
```

Or build a binary:

```bash
go build -o pokedexcli .
./pokedexcli
```

## Commands

| Command | Arguments | Description |
| --- | --- | --- |
| `help` | | Show the list of available commands |
| `map` | | Show the next 20 location areas |
| `mapb` | | Show the previous 20 location areas |
| `explore` | `<area_name>` | List the Pokémon in a location area |
| `catch` | `<pokemon_name>` | Attempt to catch a Pokémon |
| `inspect` | `<pokemon_name>` | Show name, height, weight, stats, and types of a caught Pokémon |
| `pokedex` | | List every Pokémon you've caught |
| `exit` | | Exit the program |

## Example session

```
Pokedex > map
canalave-city-area
eterna-city-area
...
Pokedex > explore pastoria-city-area
Exploring pastoria-city-area...
Found Pokemon:
 - tentacool
 - magikarp
 ...
Pokedex > catch pikachu
Throwing a Pokeball at pikachu...
pikachu escaped!
Pokedex > catch pikachu
Throwing a Pokeball at pikachu...
pikachu was caught!
You may now inspect it with the inspect command.
Pokedex > inspect pikachu
Name: pikachu
Height: 4
Weight: 60
Stats:
  -hp: 35
  -attack: 55
  ...
Types:
  - electric
Pokedex > pokedex
Your Pokedex:
 - pikachu
```

## Project layout

```
.
├── main.go                REPL loop and config setup
├── commands.go            command registry, config, help/exit
├── command_map.go         map / mapb
├── command_explore.go     explore
├── command_catch.go       catch
├── command_inspect.go     inspect
├── command_pokedex.go     pokedex
├── repl.go                input cleaning
├── repl_test.go
└── internal/
    ├── pokeapi/           PokeAPI client (location areas, Pokémon) + tests
    └── pokecache/         thread-safe, self-reaping cache + tests
```

## Tests

```bash
go test ./...
```

## License

Released under the [MIT License](LICENSE).
