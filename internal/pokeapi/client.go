package pokeapi

import (
	"io"
	"net/http"
	"time"

	"github.com/gurlivleenkainth2000/pokedexcli/internal/pokecache"
)

const baseURL = "https://pokeapi.co/api/v2"

// Client is a PokeAPI HTTP client with a response cache.
type Client struct {
	cache      *pokecache.Cache
	httpClient http.Client
}

// NewClient returns a new PokeAPI Client. timeout bounds each request;
// cacheInterval controls how long cached responses live before being reaped.
func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}

// getResource returns the raw response body for url, serving from the cache
// when available and caching fresh responses.
func (c *Client) getResource(url string) ([]byte, error) {
	if cached, ok := c.cache.Get(url); ok {
		return cached, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	c.cache.Add(url, body)
	return body, nil
}
