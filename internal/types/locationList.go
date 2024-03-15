package types

import "fmt"

type LocationList struct {
	Count     int        `json:"count"`
	Next      string     `json:"next"`
	Previous  string     `json:"previous"`
	Locations []Location `json:"results"`
}

func (l *LocationList) Print() {
	fmt.Printf("========================\n")
	for _, location := range l.Locations {
		fmt.Printf("%v\n", location.Name)
	}
	fmt.Printf("========================\n")
}
