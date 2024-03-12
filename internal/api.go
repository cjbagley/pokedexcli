package internal

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

const API_BASE_URL = "https://pokeapi.co/api/v2/"
const LOCATION_AREA_ENDPOINT = "location-area"

type Client struct {
	httpClient http.Client
}

func NewClient(timeout time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}

func (c *Client) getApiResponse(endpoint string) ([]byte, error) {
	data := []byte{}
	url := fmt.Sprintf("%v%v", API_BASE_URL, endpoint)
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
	return data, err
}
