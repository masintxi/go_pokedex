package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (client *Client) GetPokemon(url string) (PokeInLocation, error) {
	var data PokeInLocation
	err := client.GetResource(url, &data)
	if err != nil {
		return PokeInLocation{}, err
	}
	return data, nil
}

func (client *Client) GetLocations(url string) (PokeMap, error) {
	var data PokeMap
	err := client.GetResource(url, &data)
	if err != nil {
		return PokeMap{}, err
	}
	return data, nil
}

func (client *Client) CatchPokemon(url string) (PokeInfo, error) {
	var data PokeInfo
	err := client.GetResource(url, &data)
	if err != nil {
		return PokeInfo{}, err
	}
	return data, nil
}

func (client *Client) GetResource(url string, target interface{}) error {
	if val, ok := client.cache.Get(url); ok {
		err := json.Unmarshal(val, target)
		if err != nil {
			return fmt.Errorf("error unmarshalling the response: %s", err)
		}
		// fmt.Println("___getting from cache____")
		// fmt.Println(url)
		// fmt.Println("_________________________")
		return nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fmt.Errorf("error creating the request: %s", err)
	}

	res, err := client.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("error executing the request: %s", err)
	}

	defer res.Body.Close()

	if res.StatusCode > 299 {
		return fmt.Errorf("%v", res.StatusCode)
	}

	rData, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("error reading the response: %s", err)
	}

	err = json.Unmarshal(rData, target)
	if err != nil {
		return fmt.Errorf("error decoding the response: %s", err)
	}

	client.cache.Add(url, rData)

	return nil
}
