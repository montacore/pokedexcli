package pokeapi 

import (
	"fmt"
	"encoding/json"
	"net/http"
	"time"
	"pokedexcli/internal/pokecache"
	"io"
	"log"
)

const	BASEURL = "https://pokeapi.co/api/v2/location-area/"
type Client struct {
	httpClient		http.Client
	cache			*pokecache.Cache
}

func NewClient(timeout time.Duration, interval time.Duration) Client {
	client := http.Client{
		Timeout: timeout,
	}
	cache := pokecache.NewCache(interval)
	return Client{httpClient: client, cache: cache}
}

func (c Client) ListLocations(pageURL *string) (LocationAreaResponse, error){
	finalURL := ""
	if pageURL != nil {
		finalURL = *pageURL
	} else{
		finalURL = BASEURL
	}
	if data, ok := c.cache.Get(finalURL); ok {
		var locResp LocationAreaResponse
		log.Println("Hit cache branch")
		if err:= json.Unmarshal(data, &locResp); err != nil {
			return LocationAreaResponse{}, fmt.Errorf("failed to decode cached response: %w", err)
		}
		return locResp, nil
	}
	resp, err := c.httpClient.Get(finalURL)
	if err != nil {
		return LocationAreaResponse{}, fmt.Errorf("failed to make request %w", err)
	}
	
	defer resp.Body.Close()

	var locResp LocationAreaResponse
	if resp.StatusCode != http.StatusOK {
		return LocationAreaResponse{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	data, err2 := io.ReadAll(resp.Body)
	log.Println("Missed cache branch")
	if err2 != nil {
		return LocationAreaResponse{}, fmt.Errorf("failed to read body of the data into bytestream: %v", err2)
	}
	if err3 := json.Unmarshal(data, &locResp); err3 != nil {
		return LocationAreaResponse{}, fmt.Errorf("failed to unmarshal response body: %w", err3)
	}
	c.cache.Add(finalURL, data)
	return locResp, nil


}

