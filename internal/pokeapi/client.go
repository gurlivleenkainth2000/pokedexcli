package pokeapi

import (
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
