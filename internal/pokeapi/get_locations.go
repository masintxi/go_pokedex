package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetLocations(url string) (PokeMap, error) {
	res, err := http.Get(url)
	if err != nil {
		return PokeMap{}, fmt.Errorf("error fetching from URL (%s): %w", url, err)
	}
	defer res.Body.Close()

	if res.StatusCode > 299 {
		return PokeMap{}, fmt.Errorf("error: %s", res.Status)
	}

	var data PokeMap
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&data)
	if err != nil {
		return PokeMap{}, fmt.Errorf("error decoding the response: %s", err)
	}

	return data, nil
}
