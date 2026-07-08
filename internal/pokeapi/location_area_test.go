package pokeapi

import (
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"
	"time"
)

const testResponse = `{
	"count": 1054,
	"next": "https://pokeapi.co/api/v2/location-area?offset=20&limit=20",
	"previous": null,
	"results": [
		{"name": "canalave-city-area", "url": "https://pokeapi.co/api/v2/location-area/1/"},
		{"name": "eterna-city-area", "url": "https://pokeapi.co/api/v2/location-area/2/"}
	]
}`

func TestListLocationAreasParsesResponse(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(testResponse))
	}))
	defer server.Close()

	client := NewClient(5*time.Second, 5*time.Minute)
	url := server.URL

	resp, err := client.ListLocationAreas(&url)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if resp.Count != 1054 {
		t.Errorf("Count = %d, want 1054", resp.Count)
	}
	if resp.Next == nil {
		t.Errorf("Next = nil, want a next-page URL")
	} else if *resp.Next != "https://pokeapi.co/api/v2/location-area?offset=20&limit=20" {
		t.Errorf("Next = %q, want the next-page URL", *resp.Next)
	}
	if resp.Previous != nil {
		t.Errorf("Previous = %q, want nil", *resp.Previous)
	}
	if len(resp.Results) != 2 {
		t.Fatalf("len(Results) = %d, want 2", len(resp.Results))
	}
	if resp.Results[0].Name != "canalave-city-area" {
		t.Errorf("Results[0].Name = %q, want canalave-city-area", resp.Results[0].Name)
	}
}

func TestListLocationAreasUsesCache(t *testing.T) {
	var requestCount int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt32(&requestCount, 1)
		w.Write([]byte(testResponse))
	}))
	defer server.Close()

	client := NewClient(5*time.Second, 5*time.Minute)
	url := server.URL

	first, err := client.ListLocationAreas(&url)
	if err != nil {
		t.Fatalf("first call error: %v", err)
	}
	second, err := client.ListLocationAreas(&url)
	if err != nil {
		t.Fatalf("second call error: %v", err)
	}

	if got := atomic.LoadInt32(&requestCount); got != 1 {
		t.Errorf("server received %d requests, want 1 (second call should be served from cache)", got)
	}
	if len(second.Results) != len(first.Results) {
		t.Errorf("cached result differs: first %d results, second %d", len(first.Results), len(second.Results))
	}
	if second.Results[0].Name != "canalave-city-area" {
		t.Errorf("cached Results[0].Name = %q, want canalave-city-area", second.Results[0].Name)
	}
}

func TestListLocationAreasInvalidJSON(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("not valid json"))
	}))
	defer server.Close()

	client := NewClient(5*time.Second, 5*time.Minute)
	url := server.URL

	if _, err := client.ListLocationAreas(&url); err == nil {
		t.Error("expected an error for malformed JSON, got nil")
	}
}
