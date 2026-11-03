package main

import (
	"strings"
	"fmt"
	"bufio"
	"os"
)
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
				fmt.Println("Your command was: ", firstWord)
			} else {
				fmt.Println("No command entered")
			}
		}

	}
}
