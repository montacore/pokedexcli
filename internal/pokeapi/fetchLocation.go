package pokeapi 

import (
	"fmt"
	"io"
	"encoding/json"
	"net/http"
)
type Client struct {
	httpClient		http.Client
}

func NewClient(timeout time.Duration) Client {
	client := http.Client{
		Timeout: timeout,
	}
	return Client{httpClient: client}
}

func (c Client) ListLocations(pageURL *string) (LocationAreaResponse, error){
	baseURL := "https://pokeapi.co/api/v2/location-area/"
	if pageURL != nil {
		c.httpClient.Get(pageURL)
	}
}

func fetchLocations(url string) (*Location, error){
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("Failed to make request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var loc Location
	if json_err := json.NewDecoder(resp.Body).Decode(&loc); json_err != nil {
		return nil, fmt.Errorf("failed to decode JSON response: %w", json_err)
	}

	return &loc, nil 
}
