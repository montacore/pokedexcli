package main

import (
	"fmt"
)


func commandMap(cfg *config) error{
	resp, err := cfg.client.ListLocations(cfg.next)
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
