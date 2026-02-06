package main

import "fmt"

type Tile struct {
	title string
	todo  []string
}

func printTile(t Tile) {
	fmt.Println(t.title)
	for _, v := range t.todo {
		fmt.Println("-", v)
	}
}

func main() {
	myTile := Tile{title: "First Things First", todo: []string{"test", "another task", "yet another task", "and another and another"}}
	printTile(myTile)
}
