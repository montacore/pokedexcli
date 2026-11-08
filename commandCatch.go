package main

import (
	"fmt"
	"math/rand"
	"math"
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
	
	res := catchCalculation(float64(pokemon.BaseExperience))
	if res{
		fmt.Printf("%s was caught!\n", pokemon.Name)
	} else {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	}
	cfg.pokedex[pokemon.Name] = pokemon
	return nil
}

func catchCalculation(baseExperience float64) bool {
	roll := rand.Intn(65535)
	maxRoll := 65535
	calculation := 65535 * math.Pow(baseExperience/255,.25)
	threshold := float64(maxRoll) - calculation
	if threshold >= float64(roll) {
		return true
	} else{
		return false
	}

}
