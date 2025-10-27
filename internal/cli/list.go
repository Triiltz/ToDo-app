package cli

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/Triiltz/ToDo-app/internal/task"
	"github.com/mergestat/timediff"
	"github.com/spf13/cobra"
)

var showAll bool

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",
	Long:  `List all uncompleted tasks, or all tasks with the --all flag.`,
	Run: func(cmd *cobra.Command, args []string) {
		storage := task.NewStorage(GetDataFile())
		tasks, err := storage.LoadTasks()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error loading tasks: %v\n", err)
			os.Exit(1)
		}

		// Filter tasks if not showing all
		var displayTasks []*task.Task
		if showAll {
			displayTasks = tasks
		} else {
			for _, t := range tasks {
				if !t.IsComplete {
					displayTasks = append(displayTasks, t)
				}
			}
		}

		if len(displayTasks) == 0 {
			fmt.Fprintln(os.Stdout, "No tasks found.")
			return
		}

		// Create tabwriter for formatted output
		w := tabwriter.NewWriter(os.Stdout, 0, 0, 4, ' ', 0)

		if showAll {
			fmt.Fprintln(w, "ID\tTask\tCreated\tDone")
			for _, t := range displayTasks {
				timeDiff := timediff.TimeDiff(t.CreatedAt)
				fmt.Fprintf(w, "%d\t%s\t%s\t%t\n",
					t.ID,
					t.Description,
					timeDiff,
					t.IsComplete,
				)
			}
		} else {
			fmt.Fprintln(w, "ID\tTask\tCreated")
			for _, t := range displayTasks {
				timeDiff := timediff.TimeDiff(t.CreatedAt)
				fmt.Fprintf(w, "%d\t%s\t%s\n",
					t.ID,
					t.Description,
					timeDiff,
				)
			}
		}

		w.Flush()
	},
}

func init() {
	listCmd.Flags().BoolVarP(&showAll, "all", "a", false, "Show all tasks including completed ones")
	RootCmd.AddCommand(listCmd)
}
