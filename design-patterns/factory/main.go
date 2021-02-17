package main

import "fmt"

type iPokemon interface {
	setName(name string)
	setLevel(power int)
	getName() string
	getLevel() int
}

type pokemon struct {
	name  string
	level int
}

func (g *pokemon) setName(name string) {
	g.name = name
}

func (g *pokemon) getName() string {
	return g.name
}

func (g *pokemon) setLevel(level int) {
	g.level = level
}

func (g *pokemon) getLevel() int {
	return g.level
}

type bulbasaur struct {
	pokemon
}

func newBulbasaur() iPokemon {
	return &bulbasaur{
		pokemon: pokemon{
			name:  "Bulbasaur",
			level: 5,
		},
	}
}

type charmander struct {
	pokemon
}

func newCharmander() iPokemon {
	return &charmander{
		pokemon: pokemon{
			name:  "Charmander",
			level: 5,
		},
	}
}

func getPokemon(pokemon string) (iPokemon, error) {
	if pokemon == "bulbasaur" {
		return newBulbasaur(), nil
	}
	if pokemon == "charmander" {
		return newCharmander(), nil
	}
	return nil, fmt.Errorf("Wrong pokemon type passed")
}

func main() {
	bulbasaur, _ := getPokemon("bulbasaur")
	charmander, _ := getPokemon("charmander")
	printDetails(bulbasaur)
	printDetails(charmander)
}

func printDetails(p iPokemon) {
	fmt.Printf("Pokemon: %s", p.getName())
	fmt.Println()
	fmt.Printf("Level: %d", p.getLevel())
	fmt.Println()
}