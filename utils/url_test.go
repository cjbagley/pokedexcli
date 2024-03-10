package utils

import "testing"

func TestGetEndpointFromUrl(t *testing.T) {

	cases := []struct {
		input    string
		expected string
	}{
		{
			input:    "http://bing.com/search?q=dotnet",
			expected: "search?q=dotnet",
		},
		{
			input:    "https://pokeapi.co/api/v2/location",
			expected: "location",
		},
		{
			input:    "https://pokeapi.co/api/v2/location-area?offset=20&limit=20",
			expected: "location-area?limit=20&offset=20",
		},
	}

	for _, c := range cases {
		actual, err := GetEndpointFromUrl(c.input)
		if err != nil {
			t.Errorf("GetEndpointFromUrl(%v) errored: %v", c.input, err)
		}

		if actual != c.expected {
			t.Errorf("GetEndpointFromUrl(%v) == %v, expected %v", c.input, actual, c.expected)
		}
	}

}
