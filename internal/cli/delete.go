package cli

import (
	"fmt"
	"os"
	"strconv"

	"github.com/Triiltz/ToDo-app/internal/task"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete [task-id]",
	Short: "Delete a task",
	Long:  `Delete a task permanently by providing its ID.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Invalid task ID: %s\n", args[0])
			os.Exit(1)
		}

		storage := task.NewStorage(GetDataFile())
		if err := storage.DeleteTask(id); err != nil {
			fmt.Fprintf(os.Stderr, "Error deleting task: %v\n", err)
			os.Exit(1)
		}

		fmt.Fprintf(os.Stdout, "Task %d deleted successfully!\n", id)
	},
}

func init() {
	RootCmd.AddCommand(deleteCmd)
}
