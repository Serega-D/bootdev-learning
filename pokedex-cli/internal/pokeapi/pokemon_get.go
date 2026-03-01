package pokeapi

import (
	"encoding/json"
	"io"
)

func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	url := "https://pokeapi.co/api/v2/pokemon/" + pokemonName

	// 1. Проверяем кэш
	if val, ok := c.cache.Get(url); ok {
		pokemon := Pokemon{}
		if err := json.Unmarshal(val, &pokemon); err != nil {
			return Pokemon{}, err
		}
		return pokemon, nil
	}

	// 2. Идем в сеть
	res, err := c.httpClient.Get(url)
	if err != nil {
		return Pokemon{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, err
	}

	// 3. Сохраняем в кэш
	c.cache.Add(url, data)

	pokemon := Pokemon{}
	if err := json.Unmarshal(data, &pokemon); err != nil {
		return Pokemon{}, err
	}

	return pokemon, nil
}
