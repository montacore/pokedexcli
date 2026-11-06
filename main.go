package main
/*
* Pokedexcli is a command line program that allows users to explore the
* world of pokemon through a number of different commands.
*
* Usage:
* 	pokedexcli
* 	enters the pokedex prompt
*
* available commands
* 		help
* 			prints help documentation for usage.
* 		map
* 			prints 20 results from the location-area pokeapi response.
*/ 
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
