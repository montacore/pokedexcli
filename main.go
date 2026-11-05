package main

import (
	"pokedexcli/internal/pokeapi"
	"time"
)
func main() {
	cfg := &config{
		commands: getCommands(),
		client: pokeapi.NewClient(5 * time.Second, 5 * time.Second),
	}
	startRepl(cfg)
}
