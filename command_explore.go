package main

import (
	"fmt"
	"errors"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("explore explicitly must receive only one argument, a location-area name")
	}
	resp, err := cfg.client.GetLocationAreaWithPokemonResponse(args[0])
	if err != nil {
		return err
	}
	fmt.Printf("Exploring area: %s\n", resp.Name)
	for _, val := range resp.PokemonEncounters {
		fmt.Printf("- %s\n",val.Pokemon.Name)
	}
	return nil

} 
