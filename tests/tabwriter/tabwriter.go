package main

import (
	"os"
	"text/tabwriter"
)

func main() {
	writer := tabwriter.NewWriter(
		os.Stdout, 0, 2, 4, ' ', 0,
	)
	writer.Write(
		[]byte("ID\tName\tDate\n"),
	)
	writer.Write(
		[]byte("1000\tFootbar\t2025-06-05\n"),
	)
	writer.Flush() // Ensure all data is written to the output

}
