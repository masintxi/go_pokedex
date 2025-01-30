package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func commandMapb(cfg *Config) error {
	if cfg.previous == nil {
		fmt.Println("You are at the first page.")
		return nil
	}
	url := *cfg.previous

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
