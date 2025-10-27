package cli

import (
	"fmt"
	"os"

	"github.com/Triiltz/ToDo-app/internal/task"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [description]",
	Short: "Add a new task",
	Long:  `Add a new task to your todo list with the provided description.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		description := args[0]

		storage := task.NewStorage(GetDataFile())
		if err := storage.AddTask(description); err != nil {
			fmt.Fprintf(os.Stderr, "Error adding task: %v\n", err)
			os.Exit(1)
		}

		fmt.Fprintf(os.Stdout, "Task added successfully: %s\n", description)
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
