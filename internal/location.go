package internal

import "fmt"

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

func (l *LocationData) PrintLocations() {
	for _, location := range l.Locations {
		location.PrintLocation()
	}
}

func (l *Location) PrintLocation() {
	fmt.Printf("%v\n", l.Name)
}
