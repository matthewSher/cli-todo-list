package cmd

import (
	"cli-todo-list/pkg/task"
	"cli-todo-list/storage/local"
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
		local.InitTable()

		taskId, err := local.Table.DefineIdForTask()
		if err != nil {
			fmt.Println(local.ErrReadData)
			return
		}

		local.Table.AddElement(task.Task{
			Id:          taskId,
			Description: args[0],
			IsCompleted: isCompleted,
		})
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().BoolVarP(&isCompleted, "completed", "c", false, "Mark this task as completed")
}
