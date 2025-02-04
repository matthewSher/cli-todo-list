package cmd

import (
	"cli-todo-list/internal/storage/local"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "todolist",
	Short: "Manage your tasks by this app",
	Long:  `Create, read, update and delete your tasks within cli-todo-list.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to simple todo-list!")
		fmt.Println("Type --help or -h to get help.")
		local.InitCSVFile(local.TableFilename)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

}
