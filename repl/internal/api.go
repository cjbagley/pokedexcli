package internal

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

const API_ENDPOINT = "https://pokeapi.co/api/v2/"

type LocationData struct {
	Count    int        `json:"count"`
	Next     string     `json:"next"`
	Previous string     `json:"previous"`
	Location []Location `json:"results"`
}

type Location struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func GetLocations(url string) (LocationData, error) {
	var location LocationData
	data, err := getApiResponse(url)
	if err != nil {
		return location, err
	}

	err = json.Unmarshal(data, &location)
	if err != nil {
		return location, err
	}
	return location, nil
}

func getApiResponse(url string) ([]byte, error) {
	data := []byte{}
	res, err := http.Get(fmt.Sprintf("%v%v", API_ENDPOINT, url))
	if err != nil {
		return data, err
	}

	if res.StatusCode > 299 {
		msg := fmt.Sprintf("Response failed with status code: %d", res.StatusCode)
		return []byte{}, errors.New(msg)
	}

	data, err = io.ReadAll(res.Body)
	if err != nil {
		return data, err
	}
	res.Body.Close()

	return data, nil
}
