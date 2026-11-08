package main

import "fmt"

func commandPokedex(cfg *config, args ...string) error {
	if len(args) > 0 {
		return fmt.Errorf("Pokedex command does not accept any arguments")
	}
	fmt.Print("Your Pokedex:\n")
	for pokemon := range cfg.pokedex {
		fmt.Printf(" - %v\n", pokemon)
	}

	return nil 
}
