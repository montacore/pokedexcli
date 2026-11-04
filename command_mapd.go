package main

import (
	"fmt"
	"errors"
)


func commandMapb(cfg *config) error{
	if cfg.previous == nil {
		error := new Error("you're on the first page")
		return error
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
