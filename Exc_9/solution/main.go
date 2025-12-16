package main

import (
	"fmt"
	"os"
	"strings"

	"exc9/mapred"
)

// Main function
func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <input-file>")
		os.Exit(1)
	}

	// Read file
	filename := os.Args[1]
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	// Split file into lines (input for MapReduce)
	text := strings.Split(string(data), "\n")

	// Run MapReduce
	var mr mapred.MapReduce
	results := mr.Run(text)

	// Print results to stdout
	for word, count := range results {
		fmt.Printf("%s: %d\n", word, count)
	}
}
