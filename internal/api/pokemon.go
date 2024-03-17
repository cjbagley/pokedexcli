package api

import (
	"encoding/json"
	"errors"
	"strings"

	t "github.com/cjbagley/pokedexcli/internal/types"
)

func (c *Client) GetPokemon(name string) (t.Pokémon, error) {
	var pokémon t.Pokémon

	url := POKEMON_ENDPOINT + "/" + name

	data, err := c.getApiResponse(url)
	if err != nil {

		if strings.Contains(err.Error(), "404") {
			return pokémon, errors.New("pokémon given does not exist")
		}

		return pokémon, err
	}

	err = json.Unmarshal(data, &pokémon)
	if err != nil {
		return t.Pokémon{}, err
	}

	return pokémon, nil
}
