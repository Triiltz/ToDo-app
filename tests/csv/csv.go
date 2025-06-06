package main

import (
	"encoding/csv"
	"os"
)

func main() {
	writer := csv.NewWriter(os.Stdout) // Create a new CSV writer
	writer.Write([]string{
		"1", "My task", "today",
	})

	writer.Write([]string{
		"2", "Next task", "tomorrow",
	})
	writer.Flush() // Ensure all data is written to the output
}
