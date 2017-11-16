package db

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/nikhosagala/herobio/models"
)

// Seed ...
func Seed() {
	filePath := "./hero"

	files, err := ioutil.ReadDir(filePath)

	if err != nil {
		fmt.Printf("# Error when reading path %s\n", filePath)
		os.Exit(1)
	}

	var hero models.Hero

	for _, file := range files {
		fmt.Printf("# Reading file %s\n", file.Name())
		path := filePath + "/" + file.Name()
		content, err := ioutil.ReadFile(path)
		if err != nil {
			fmt.Printf("# Error: %v\n ", err)
			os.Exit(1)
		}
		err2 := json.Unmarshal(content, &hero)
		if err2 != nil {
			fmt.Printf("# Error: %v\n ", err)
			os.Exit(1)
		}
		DB.Create(&hero)
	}
}
