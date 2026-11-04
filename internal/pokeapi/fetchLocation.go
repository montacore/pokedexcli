package pokeapi 

import (
	"fmt"
	"encoding/json"
	"net/http"
	"time"
)

const	BASEURL = "https://pokeapi.co/api/v2/location-area/"
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
	finalURL := ""
	if pageURL != nil {
		finalURL = *pageURL
	} else{
		finalURL = BASEURL
	}
	resp, err := c.httpClient.Get(finalURL)
	if err != nil {
		return LocationAreaResponse{}, fmt.Errorf("failed to make request %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return LocationAreaResponse{}, fmt.Errorf("Unexpected status code: %d", resp.StatusCode)
	}
	var locResp LocationAreaResponse
	if jsonErr := json.NewDecoder(resp.Body).Decode(&locResp); jsonErr != nil {
		return LocationAreaResponse{}, fmt.Errorf("failed to decode JSON response: %w", jsonErr)
	}
	return locResp, nil


}

