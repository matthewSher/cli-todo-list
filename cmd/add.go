package cmd

import (
	"cli-todo-list/internal/storage/local"
	"cli-todo-list/internal/task"
	"fmt"

	"github.com/spf13/cobra"
)

var isCompleted bool

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add [task description]",
	Short: "Add new task",
	Long: `Use this command to add new task to your list.
			If you're adding task firstly, program will generate CSV file
			in current directory.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Create CSV table file if it isn't exist
		local.InitCSVFile(local.TableFilename)

		err := local.AddElement(local.TableFilename, task.Task{
			Description: args[0],
			IsCompleted: isCompleted,
		})
		if err != nil {
			fmt.Println("error:", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().BoolVarP(&isCompleted, "completed", "c", false, "Mark this task as completed")
}
