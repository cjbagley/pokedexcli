package internal

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

const API_BASE_URL = "https://pokeapi.co/api/v2/"
const LOCATION_AREA_ENDPOINT = "location-area"

func GetLocations(url string) (LocationData, error) {
	var location LocationData

	if url == "" {
		url = LOCATION_AREA_ENDPOINT
	}

	data, err := getApiResponse(url)
	if err != nil {
		return location, err
	}

	err = json.Unmarshal(data, &location)
	if err != nil {
		return LocationData{}, err
	}

	return location, nil
}

func getApiResponse(endpoint string) ([]byte, error) {
	data := []byte{}
	url := fmt.Sprintf("%v%v", API_BASE_URL, endpoint)
	res, err := http.Get(url)
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
