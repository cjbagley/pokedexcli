package types

import (
	"sync"
)

var lock = &sync.Mutex{}

type Pokedex struct {
	dex map[string]Pokémon
}

func (pd Pokedex) Add(p Pokémon) {
	pd.dex[p.Name] = p
}

func (pd Pokedex) IsCaught(p Pokémon) bool {
	_, ok := pd.GetByName(p.Name)
	return ok
}

func (pd Pokedex) GetByName(name string) (pokémon Pokémon, ok bool) {
	pokémon, ok = pd.dex[name]
	return
}

var singlePokedex *Pokedex

func GetPokedex() *Pokedex {
	if singlePokedex == nil {
		lock.Lock()
		defer lock.Unlock()
		// Correct to be in twice, see
		// https://refactoring.guru/design-patterns/singleton/go/example
		if singlePokedex == nil {
			singlePokedex = &Pokedex{
				dex: make(map[string]Pokémon),
			}
		}
	}

	return singlePokedex
}
