package pokeapi

import (
	"net/http"
	"time"

	"github.com/masintxi/go_pokedex/internal/pokecache"
)

type Client struct {
	httpClient http.Client
	cache      pokecache.Cache
}

func NewClient(interval time.Duration) Client {
	return Client{
		httpClient: http.Client{},
		cache:      pokecache.NewCache(interval),
	}
}
