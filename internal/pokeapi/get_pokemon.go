// Deprecated. Use get_resource.go instead.
package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (client *Client) GetPokemon_old(url string) (PokeInLocation, error) {
	if val, ok := client.cache.Get(url); ok {
		var data PokeInLocation
		err := json.Unmarshal(val, &data)
		if err != nil {
			return PokeInLocation{}, fmt.Errorf("error unmarshalling the response: %s", err)
		}
		fmt.Println("___getting from cache____")
		fmt.Println(url)
		fmt.Println("_________________________")
		return data, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokeInLocation{}, fmt.Errorf("error creating the request: %s", err)
	}

	res, err := client.httpClient.Do(req)
	if err != nil {
		return PokeInLocation{}, fmt.Errorf("error executing the request: %s", err)
	}

	defer res.Body.Close()

	if res.StatusCode > 299 {
		return PokeInLocation{}, fmt.Errorf("error: %s", res.Status)
	}

	rData, err := io.ReadAll(res.Body)
	if err != nil {
		return PokeInLocation{}, fmt.Errorf("error reading the response: %s", err)
	}

	var data PokeInLocation
	err = json.Unmarshal(rData, &data)
	if err != nil {
		return PokeInLocation{}, fmt.Errorf("error decoding the response: %s", err)
	}

	client.cache.Add(url, rData)

	return data, nil
}
