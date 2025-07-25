package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type contactInfo struct {
	Name  string
	Email string
}

type purchaseInfo struct {
	Name   string
	Price  float32
	Amount int
}

func main() {
	var contacts []contactInfo = loadJSON[contactInfo]("./contactInfo.json")
	fmt.Printf("\n%+v", contacts)

	var purchases []purchaseInfo = loadJSON[purchaseInfo]("./purchaseInfo.json")
	fmt.Printf("\n%+v", purchases)
}

func loadJSON[T any](filePath string) []T {
	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error reading file %s: %v\n", filePath, err)
		return nil
	}

	var loaded []T
	err = json.Unmarshal(data, &loaded)
	if err != nil {
		fmt.Printf("Error unmarshaling JSON from %s: %v\n", filePath, err)
		return nil
	}

	return loaded
}
