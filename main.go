package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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
	var separatedTasks []string
	// Get title
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("----ADD A NEW TILE----\nInput the title: ")
	newTitle, _ := reader.ReadString('\n')

	// Get tasks
	fmt.Println("\n----ADD TASKS----\nInput tasks separated by commas: ")
	allTasks, _ := reader.ReadString('\n')

	//Separate tasks by comma
	allTasks = strings.TrimSpace(allTasks) //Remove \n from end
	commaTasks := strings.SplitSeq(allTasks, ",")
	for v := range commaTasks {
		separatedTasks = append(separatedTasks, strings.TrimSpace(v))
	}

	//Create and print new tile
	newTile := tileConstructor(newTitle, separatedTasks...)
	printTile(newTile)
}
