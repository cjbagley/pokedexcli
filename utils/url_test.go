package utils

import "testing"

func TestGetEndpointFromUrl(t *testing.T) {
	url := "http://bing.com/search?q=dotnet"
	expected := "search?q=dotnet"
	got, err := GetEndpointFromUrl(url)
	if err != nil {
		t.Errorf("GetEndpointFromUrl errored: %v", err)
	}

	if got != expected {
		t.Errorf("GetEndpointFromUrl expected '%v', got '%v", expected, got)
	}
}
