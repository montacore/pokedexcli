package internal

import (
	"fmt"
	"io"
	"encoding/json"
	"net/http"
)

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
