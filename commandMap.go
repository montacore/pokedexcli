package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"io"
	"log"
)


func commandMap(cfg *config) string{
	url := "https://pokeapi.co/api/v2/location/"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Errorf("Error retrieving location data: %s", err)
	}
	defer resp.Body.Close()

	bodyBytes, byte_err := io.ReadAll(resp.Body)
	if byte_err != nil {
		log.Fatalf("Failed to read response body: %v", byte_err)
	}
	
	var loc Location
	json_err := json.Unmarshal(bodyBytes, &loc)
	if json_err != nil {
		log.Fatalf("Failed to Unmarshal JSON: %v", json_err)
	}

	return loc.Name 
}
