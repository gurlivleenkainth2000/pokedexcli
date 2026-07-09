package pokeapi

import (
	"encoding/json"
)

// LocationAreasResp is the paginated response from the location-area endpoint.
type LocationAreasResp struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

// LocationArea is the detailed response for a single location area.
type LocationArea struct {
	Name              string `json:"name"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

// ListLocationAreas fetches a page of location areas. If pageURL is nil, it
// fetches the first page from the default endpoint.
func (c *Client) ListLocationAreas(pageURL *string) (LocationAreasResp, error) {
	url := c.baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	body, err := c.getResource(url)
	if err != nil {
		return LocationAreasResp{}, err
	}

	var locationAreas LocationAreasResp
	if err := json.Unmarshal(body, &locationAreas); err != nil {
		return LocationAreasResp{}, err
	}

	return locationAreas, nil
}

// GetLocationArea fetches the details (including Pokemon encounters) for a
// single location area by name.
func (c *Client) GetLocationArea(name string) (LocationArea, error) {
	url := c.baseURL + "/location-area/" + name

	body, err := c.getResource(url)
	if err != nil {
		return LocationArea{}, err
	}

	var locationArea LocationArea
	if err := json.Unmarshal(body, &locationArea); err != nil {
		return LocationArea{}, err
	}

	return locationArea, nil
}
