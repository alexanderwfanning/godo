package main

import "fmt"

type Tile struct {
	title string
	todo  []string
}

func (t Tile) print() {
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