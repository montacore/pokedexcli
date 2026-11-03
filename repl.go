package main

import (
	"strings"
	"fmt"
	"bufio"
	"os"
)
type cliCommand struct {
	name		string
	description	string
	callback	func() error
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
	}
	return commands	
}

func cleanInput(text string) []string{
	output := strings.Fields(strings.ToLower(text))
	return output
}
func startRepl(){


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

					err := cmd.callback()
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
