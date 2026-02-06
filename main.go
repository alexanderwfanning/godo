package main

import "fmt"

type Tile struct {
	title string
	todo  []string
}

func printTile(t Tile) {
	fmt.Println("\n", t.title)
	for i, v := range t.todo {
		fmt.Printf("%d. %s\n", i+1, v)
	}
}

func main() {
	myTile := Tile{title: "First Things First", todo: []string{"test", "another task", "yet another task", "and another and another"}}
	anotherTile := Tile{title: "Secondary Things Second", todo: []string{"Make this work", "And make it good"}}
	tertiaryTile := Tile{title: "Tertiary Things Third", todo: []string{"Bubble Tea looks cool!", "Maybe I should use that?"}}
	allMyTiles := []Tile{myTile, anotherTile, tertiaryTile}
	for _, tile := range allMyTiles {
		printTile(tile)
	}
}
