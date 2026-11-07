package main

import (
	"fmt"
	"errors"
)


func commandMapb(cfg *config, args ...string) error{
	if cfg.previous == nil {
		return errors.New("you're on the first page")
		
	}
	resp, err := cfg.client.ListLocations(cfg.previous)
	if err != nil {
		return err
	}
	cfg.next = resp.Next
	cfg.previous = resp.Previous

	for _, val := range resp.Results{
		fmt.Println(val.Name)
	}
	return nil
	
}
