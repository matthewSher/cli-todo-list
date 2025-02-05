package cmd

import (
	"cli-todo-list/internal/storage/local"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete [task number]",
	Short: "Delete existing task",
	Long:  `Use this command to delete existing task.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		taskId, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("incorrect input type")
			return
		}

		err = local.DeleteElement(local.TableFilename, taskId)
		if err == local.ErrOutOfRange {
			fmt.Println("current index out of range of the list")
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
