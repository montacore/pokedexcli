package main


import "fmt"

func commandInspect(cfg *config, args ...string) error {
	/*TODO: Function should take the name of a Pokemon and print the name, height,
	  weight stats and type(s) of the Pokemon:
	   Example Usage:
	  		Pokedex > inspect pidgey
	  		you have not caught that pokemon
	  		Pokedex > catch pidgey
	  		Throwing a Pokeball at pidgey...
	  		pidgey was caught!
	  		Pokedex > inspect pidgey
	  		Name: pidgey
	  		Height: 3
	  		Weight: 18
	  		Stats
	  			-hp: 40
	  			-attack: 45
	  			-defense: 40
	 			-special-attack: 35
	 			-special-defense: 35
	 			speed: 56
	 		Types:
	 			- normal
	 			- flying
	*NOTE: You should not need to make an api call, because this information
	* should be stored when you caught the Pokemon.
	*/
	if len(args) != 1 {
		return fmt.Errorf("you must provide a target name")
	}
	if pokemon, ok := cfg.pokedex[args[0]]; ok{
		stats := pokemon.Stats
		fmt.Printf("Name: %s\n", pokemon.Name)
		fmt.Printf("Height: %d\n", pokemon.Height)
		fmt.Printf("Weight: %d\n", pokemon.Weight)
		fmt.Println("Stats:")
		for _, stat := range stats {
			fmt.Printf("	-%s: %d\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Println("Types:")
		for _, t := range pokemon.Types {
			fmt.Printf("	- %s\n", t.Type.Name)
		}
	} else {
		return fmt.Errorf("you have not caught that pokemon\n")
	}
	return nil
}
