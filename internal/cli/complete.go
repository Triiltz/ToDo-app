package cli

import (
	"fmt"
	"os"
	"strconv"

	"github.com/Triiltz/ToDo-app/internal/task"
	"github.com/spf13/cobra"
)

var completeCmd = &cobra.Command{
	Use:   "complete [task-id]",
	Short: "Mark a task as complete",
	Long:  `Mark a task as completed by providing its ID.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Invalid task ID: %s\n", args[0])
			os.Exit(1)
		}

		storage := task.NewStorage(GetDataFile())
		if err := storage.CompleteTask(id); err != nil {
			fmt.Fprintf(os.Stderr, "Error completing task: %v\n", err)
			os.Exit(1)
		}

		fmt.Fprintf(os.Stdout, "Task %d marked as complete!\n", id)
	},
}

func init() {
	RootCmd.AddCommand(completeCmd)
}
