package api

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

const API_BASE_URL = "https://pokeapi.co/api/v2/"
const LOCATION_AREA_ENDPOINT = "location-area"
const POKEMON_ENDPOINT = "pokemon"

type Client struct {
	httpClient http.Client
	cache      Cache
}

func NewClient(httptimeout time.Duration, cachettl time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: httptimeout,
		},
		cache: NewCache(cachettl),
	}
}

func (c *Client) getApiResponse(endpoint string) ([]byte, error) {
	data := []byte{}
	url := fmt.Sprintf("%v%v", API_BASE_URL, endpoint)

	if cache, ok := c.cache.Get(url); ok {
		return cache, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return data, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return data, err
	}
	defer res.Body.Close()

	if res.StatusCode > 299 {
		msg := fmt.Sprintf("Response failed with status code: %d", res.StatusCode)
		return []byte{}, errors.New(msg)
	}

	data, err = io.ReadAll(res.Body)

	c.cache.Add(url, data)

	return data, err
}
