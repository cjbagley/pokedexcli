package internal

import "testing"

func TestCallApi(t *testing.T) {
	result, err := getApiResponse("location")
	if err != nil {
		t.Errorf("API call error: %v", err)
	}
	if result == nil {
		t.Errorf("API call did not return any results")
	}
}
