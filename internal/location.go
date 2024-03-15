package internal

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
