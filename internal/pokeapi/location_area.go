package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// LocationAreasResp is the response from the location-area endpoint.
type LocationAreasResp struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

// ListLocationAreas fetches a page of location areas. If pageURL is nil, it
// fetches the first page from the default endpoint.
func (c *Client) ListLocationAreas(pageURL *string) (LocationAreasResp, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	if cached, ok := c.cache.Get(url); ok {
		var locationAreas LocationAreasResp
		if err := json.Unmarshal(cached, &locationAreas); err != nil {
			return LocationAreasResp{}, err
		}
		return locationAreas, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationAreasResp{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreasResp{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreasResp{}, err
	}

	c.cache.Add(url, body)

	var locationAreas LocationAreasResp
	if err := json.Unmarshal(body, &locationAreas); err != nil {
		return LocationAreasResp{}, err
	}

	return locationAreas, nil
}
