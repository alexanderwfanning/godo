package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/tidwall/buntdb"
)

func main() {
	db, err := buntdb.Open("data.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

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
	saveToDB(db, newTile)
}

func saveToDB(db *buntdb.DB, t Tile) error {
	// Convert ID to str to store as key in DB
	key := fmt.Sprintf("%d", t.Uid)

	// Convert tasks to JSON
	// []todo -> []bytes -> json string
	todoBytes, err := json.Marshal(t)
	if err != nil {
		return err
	}
	value := string(todoBytes)
	fmt.Println(value)
	// Write to DB
	db.Update(func(tx *buntdb.Tx) error {
		_, _, err := tx.Set(key, value, nil)
		return err
	})
	return err
}
