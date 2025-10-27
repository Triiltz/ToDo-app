package cli

import (
	"github.com/spf13/cobra"
)

var dataFile string

// RootCmd is the root command of the application
var RootCmd = &cobra.Command{
	Use:   "tasks",
	Short: "A CLI todo list manager",
	Long:  `Manage your tasks directly from the command line with a simple CSV-based storage.`,
}

func init() {
	// Global flag to specify custom data file
	RootCmd.PersistentFlags().StringVarP(&dataFile, "file", "f", "", "Custom data file path (default: ~/Code/go-projects/ToDo-app/.tasks/tasks.csv)")
}

// GetDataFile returns the data file path
func GetDataFile() string {
	return dataFile
}
