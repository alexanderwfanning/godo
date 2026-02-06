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

func tileConstructor(tileTitle string, tasks ...string) Tile {
	newTile := Tile{title: tileTitle}
	newTile.addTask(tasks)
	return newTile
}

func (t *Tile) addTask(tasks []string) {
	for _, task := range tasks {
		t.todo = append(t.todo, task)
	}
}

func main() {
	myTile := tileConstructor("First Things First", "test", "another task", "yet another task", "and another and another")
	allMyTiles := []Tile{myTile}
	for _, tile := range allMyTiles {
		printTile(tile)
	}
}
