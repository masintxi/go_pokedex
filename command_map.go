package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type pokeMap struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func commandMap(cfg *Config) error {
	url := "https://pokeapi.co/api/v2/location-area/"
	if cfg.next != nil {
		url = *cfg.next
	}

	if url == "" {
		fmt.Println("No more results.")
		return nil
	}

	res, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error getting the url: %s\n", err)
		return fmt.Errorf("error getting the url: %s", err)
	}
	defer res.Body.Close()

	if res.StatusCode > 299 {
		fmt.Printf("Error: %s\n", res.Status)
		return fmt.Errorf("error: %s", res.Status)
	}

	var data pokeMap
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&data)
	if err != nil {
		fmt.Printf("Error decoding the response: %s\n", err)
		return fmt.Errorf("error decoding the response: %s", err)
	}

	cfg.next = data.Next
	cfg.previous = data.Previous

	for _, result := range data.Results {
		fmt.Println(result.Name)
	}

	return nil
}
