package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageUrl *string) (ResponseLocations, error) {
	url := baseURL + "/location-area/"
	if pageUrl != nil {
		url = *pageUrl
	}

	if val, ok := c.cache.Get(url); ok {
		locationsResp := ResponseLocations{}
		err := json.Unmarshal(val, &locationsResp)
		if err != nil {
			return ResponseLocations{}, fmt.Errorf("error unmarshaling data: %w", err)
		}

		return locationsResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ResponseLocations{}, fmt.Errorf("error creating request: %w", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return ResponseLocations{}, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return ResponseLocations{}, fmt.Errorf("error reading data: %w", err)
	}

	locationsResp := ResponseLocations{}
	err = json.Unmarshal(data, &locationsResp)
	if err != nil {
		return ResponseLocations{}, fmt.Errorf("error unmarshaling data: %w", err)
	}

	c.cache.Add(url, data)

	return locationsResp, nil
}
