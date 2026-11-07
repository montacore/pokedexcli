package pokeapi

import (
	"fmt"
	"encoding/json"
	"net/http"
	"io"
)

func (c *Client)FetchPokemon(pokemon string) (PokemonStatsResponse, error){
	finalURL := "https://pokeapi.co/api/v2/pokemon/" + pokemon 
	if data, ok := c.cache.Get(finalURL); ok {
		var responseVar PokemonStatsResponse
		if err := json.Unmarshal(data, &responseVar); err != nil {
			return PokemonStatsResponse{}, fmt.Errorf("failed to decode cached response: %w", err)
		}
		return responseVar, nil
	}
	resp, err := c.httpClient.Get(finalURL)
	if err != nil {
		return PokemonStatsResponse{}, fmt.Errorf("failed to make request %w", err)
	}
	
	defer func() {
		_ =	resp.Body.Close()
	}()

	var responseVar PokemonStatsResponse
	if resp.StatusCode != http.StatusOK {
		return PokemonStatsResponse{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	data, err2 := io.ReadAll(resp.Body)
	if err2 != nil {
		return PokemonStatsResponse{}, fmt.Errorf("failed to read body of the data into bytestream: %w", err2)
	}
	if err3 := json.Unmarshal(data, &responseVar); err3 != nil {
		return PokemonStatsResponse{}, fmt.Errorf("failed to unmarshal response body: %w", err3)
	}
	c.cache.Add(finalURL, data)
	return responseVar, nil


}
