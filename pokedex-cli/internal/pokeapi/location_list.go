package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
)

func (c *Client) ListLocations(pageURL *string) (RespShallowLocations, error) {
	url := "https://pokeapi.co/api/v2/location-area"

	if pageURL != nil {
		url = *pageURL
	}

	if val, ok := c.cache.Get(url); ok {
		// Если нашли (ok == true), сразу "распаковываем" байты в структуру
		locationsResp := RespShallowLocations{}
		err := json.Unmarshal(val, &locationsResp)
		if err != nil {
			return RespShallowLocations{}, err
		}
		return locationsResp, nil // Возвращаем данные без запроса в сеть!
	}

	res, err := c.httpClient.Get(url)
	if err != nil {
		return RespShallowLocations{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return RespShallowLocations{}, err
	}

	locationsResp := RespShallowLocations{}
	err = json.Unmarshal(data, &locationsResp)
	if err != nil {
		return RespShallowLocations{}, err
	}

	return locationsResp, nil
}

func (c *Client) GetLocationArea(locationAreaName string) (RespLocationArea, error) {
	url := "https://pokeapi.co/api/v2/location-area/" + locationAreaName

	// Сначала проверяем кэш
	if val, ok := c.cache.Get(url); ok {
		locationAreaResp := RespLocationArea{}
		err := json.Unmarshal(val, &locationAreaResp)
		if err != nil {
			return RespLocationArea{}, err
		}
		return locationAreaResp, nil
	}

	// Если нет в кэше — в сеть
	res, err := c.httpClient.Get(url)
	if err != nil {
		return RespLocationArea{}, err
	}
	defer res.Body.Close()

	if res.StatusCode == 404 {
		return RespLocationArea{}, fmt.Errorf("location area not found")
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return RespLocationArea{}, err
	}

	c.cache.Add(url, data)

	locationAreaResp := RespLocationArea{}
	err = json.Unmarshal(data, &locationAreaResp)
	if err != nil {
		return RespLocationArea{}, err
	}

	return locationAreaResp, nil
}
