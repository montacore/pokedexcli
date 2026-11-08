package main

import (
	"fmt"
	"math/rand"
)
func commandCatch(cfg *config, args ...string) error{
	if len(args) != 1 {
		return fmt.Errorf("Exactly one argument required, please provide a Pokemon name")
	}
	arg := args[0]
	pokemon, err := cfg.client.FetchPokemon(arg)
	if err != nil {
		return err
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
	
	res := catchCalculation(pokemon.BaseExperience)
	if res{
		fmt.Printf("%s was caught!\n", pokemon.Name)
		cfg.pokedex[pokemon.Name] = pokemon
	} else {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	}
	return nil
}

func catchCalculation(baseExperience int) bool {
	roll := rand.Intn(baseExperience)
	fmt.Printf("[DEBUG] roll: %d baseExperience: %d", roll, baseExperience)
	if roll > 70 {
		return false
	} else{
		return true
	}

}
