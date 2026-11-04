package main

import (
	"strings"
	"fmt"
	"bufio"
	"os"
	"pokedexcli/internal/pokeapi"
)
type cliCommand struct {
	name		string
	description	string
	callback	func(cfg *config) error
}

func getCommands() map[string]cliCommand{

	commands := map[string]cliCommand{
		"exit": {
			name:			"exit",
			description:	"Exit the Pokedex",
			callback:		commandExit,
		},
		"help": {
			name:			"help",
			description:	"Display a help message",
			callback:		commandHelp,
		},
		"map": {
			name:			"map",
			description:	"Display a list of 20 map locations",
			callback:		commandMap,

		},
		"mapd": {
			name:			"mapd",
			description:	"Display a list of the previous 20 map locations",
			callback:		commandMapd,
		},
	}
	return commands	
}
type config struct {
	commands	map[string]cliCommand
	client		pokeapi.Client
	next		*string
	previous	*string
}
func cleanInput(text string) []string{
	output := strings.Fields(strings.ToLower(text))
	return output
}
func startRepl(cfg *config){


	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		if scanner.Scan() {
			input := scanner.Text()

			cleanedInput := cleanInput(input)

			if len(cleanedInput) > 0 {
				firstWord := cleanedInput[0]
				commands := getCommands()
				if cmd, ok := commands[firstWord]; ok {

					err := cmd.callback(cfg)
					if err != nil {
						fmt.Println(err)
					}
				} else {
					fmt.Print("Unknown command")
				}
			} 
		}

	}
}
