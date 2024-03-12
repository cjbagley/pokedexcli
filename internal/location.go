package internal

import (
	"encoding/json"
	"fmt"
)

type LocationData struct {
	Count     int        `json:"count"`
	Next      string     `json:"next"`
	Previous  string     `json:"previous"`
	Locations []Location `json:"results"`
}

type Location struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func (c *Client) GetLocations(url string) (LocationData, error) {
	var location LocationData

	if url == "" {
		url = LOCATION_AREA_ENDPOINT
	}

	data, err := c.getApiResponse(url)
	if err != nil {
		return location, err
	}

	err = json.Unmarshal(data, &location)
	if err != nil {
		return LocationData{}, err
	}

	return location, nil
}

func (l *LocationData) PrintLocations() {
	fmt.Printf("========================\n")
	for _, location := range l.Locations {
		location.PrintLocation()
	}
	fmt.Printf("========================\n")
}

func (l *Location) PrintLocation() {
	fmt.Printf("%v\n", l.Name)
}
