package pokeapi

import (
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"
	"time"
)

const pokemonResponse = `{
	"name": "pikachu",
	"base_experience": 112,
	"height": 4,
	"weight": 60,
	"stats": [
		{"base_stat": 35, "stat": {"name": "hp"}},
		{"base_stat": 55, "stat": {"name": "attack"}}
	],
	"types": [
		{"type": {"name": "electric"}}
	]
}`

func TestGetPokemonParsesResponse(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(pokemonResponse))
	}))
	defer server.Close()

	client := NewClient(5*time.Second, 5*time.Minute)
	client.baseURL = server.URL

	pokemon, err := client.GetPokemon("pikachu")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if pokemon.Name != "pikachu" {
		t.Errorf("Name = %q, want pikachu", pokemon.Name)
	}
	if pokemon.BaseExperience != 112 {
		t.Errorf("BaseExperience = %d, want 112", pokemon.BaseExperience)
	}
	if pokemon.Height != 4 {
		t.Errorf("Height = %d, want 4", pokemon.Height)
	}
	if pokemon.Weight != 60 {
		t.Errorf("Weight = %d, want 60", pokemon.Weight)
	}
	if len(pokemon.Stats) != 2 {
		t.Fatalf("len(Stats) = %d, want 2", len(pokemon.Stats))
	}
	if pokemon.Stats[0].Stat.Name != "hp" || pokemon.Stats[0].BaseStat != 35 {
		t.Errorf("Stats[0] = %+v, want hp:35", pokemon.Stats[0])
	}
	if len(pokemon.Types) != 1 || pokemon.Types[0].Type.Name != "electric" {
		t.Errorf("Types = %+v, want [electric]", pokemon.Types)
	}
}

func TestGetPokemonUsesCache(t *testing.T) {
	var requestCount int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt32(&requestCount, 1)
		w.Write([]byte(pokemonResponse))
	}))
	defer server.Close()

	client := NewClient(5*time.Second, 5*time.Minute)
	client.baseURL = server.URL

	if _, err := client.GetPokemon("pikachu"); err != nil {
		t.Fatalf("first call error: %v", err)
	}
	if _, err := client.GetPokemon("pikachu"); err != nil {
		t.Fatalf("second call error: %v", err)
	}

	if got := atomic.LoadInt32(&requestCount); got != 1 {
		t.Errorf("server received %d requests, want 1 (second call should be cached)", got)
	}
}
