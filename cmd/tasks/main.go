package main

import (
	"fmt"
	"os"

	"github.com/Triiltz/ToDo-app/internal/cli"
)

func main() {
	if err := cli.RootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
