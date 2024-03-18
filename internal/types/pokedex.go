package types

import (
	"fmt"
	"sync"

	"github.com/cjbagley/pokedexcli/utils"
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

func (pd Pokedex) PrintCaughtPokemon() {
	if len(pd.dex) == 0 {
		fmt.Println("Pokédex is empty - go catch some Pokémon!")
		return
	}

	fmt.Println("Your Pokédex:")
	for _, pokémon := range pd.dex {
		fmt.Printf("\t- %v\n", utils.UcFirst(pokémon.Name))
	}
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
