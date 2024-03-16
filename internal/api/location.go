package api

import (
	"encoding/json"

	t "github.com/cjbagley/pokedexcli/internal/types"
)

func (c *Client) GetLocationList(url string) (t.LocationList, error) {
	var locations t.LocationList

	if url == "" {
		url = LOCATION_AREA_ENDPOINT
	}

	data, err := c.getApiResponse(url)
	if err != nil {
		return locations, err
	}

	err = json.Unmarshal(data, &locations)
	if err != nil {
		return t.LocationList{}, err
	}

	return locations, nil
}

func (c *Client) GetLocation(area string) (t.Location, error) {
	var location t.Location

	url := LOCATION_AREA_ENDPOINT + "/" + area

	data, err := c.getApiResponse(url)
	if err != nil {
		return location, err
	}

	err = json.Unmarshal(data, &location)
	if err != nil {
		return t.Location{}, err
	}

	return location, nil
}
