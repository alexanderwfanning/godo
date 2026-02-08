package main

import "fmt"

type Tile struct {
	Title string
	Todo  []string
	Uid   int
}

func (t Tile) print() {
	fmt.Println("\n", t.Title)
	for i, v := range t.Todo {
		fmt.Printf("%d. %s\n", i+1, v)
	}
}

func tileConstructor(tileTitle string, tasks ...string) Tile {
	newTile := Tile{Title: tileTitle}
	newTile.addTask(tasks)
	return newTile
}

func (t *Tile) addTask(tasks []string) {
	for _, task := range tasks {
		t.Todo = append(t.Todo, task)
	}
}
