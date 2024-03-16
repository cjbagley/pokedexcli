package api

import (
	"testing"
	"time"
)

func TestCallApi(t *testing.T) {
	client := NewClient(5*time.Second, 5*time.Minute)
	result, err := client.getApiResponse(LOCATION_AREA_ENDPOINT)
	if err != nil {
		t.Errorf("API call error: %v", err)
	}
	if result == nil {
		t.Errorf("API call did not return any results")
	}
}
