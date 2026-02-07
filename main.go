package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


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
	newTile.print()
}
