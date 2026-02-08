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

func addTile() int {
	return 0
}

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
	pushDB(db, newTile)
	tile, luerr := pullDB(db, "0")
	if luerr != nil {
		fmt.Println("Look up error! Tile ID {ID} not found") // Placeholder
	}
	tile.print()
}

func pushDB(db *buntdb.DB, t Tile) error {
	// Convert ID to str to store as key in DB
	key := fmt.Sprintf("%d", t.Uid)

	// Convert Tile to JSON
	// Tile -> []bytes -> JSON
	tileBytes, err := json.Marshal(t)
	if err != nil {
		return err
	}
	value := string(tileBytes)

	// Write to DB
	err = db.Update(func(tx *buntdb.Tx) error {
		_, _, err = tx.Set(key, value, nil)
		return err
	})
	return err
}

func pullDB(db *buntdb.DB, key string) (Tile, error) {
	// Initialize Tile
	tile := Tile{}

	// Read from DB
	err := db.View(func(tx *buntdb.Tx) error {
		// Tile is stored as JSON
		tileJSON, err := tx.Get(key)
		if err != nil {
			return err
		}

		// Tile is converted from JSON to Bytes
		tileBytes := []byte(tileJSON)

		// Tile is converted from Bytes to Tile
		err = json.Unmarshal(tileBytes, &tile)
		if err != nil {
			return err
		}
		return nil
	})
	return tile, err
}
