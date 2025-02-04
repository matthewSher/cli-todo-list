package cmd

import (
	"cli-todo-list/internal/storage/local"
	"cli-todo-list/internal/table"
	"fmt"

	"github.com/spf13/cobra"
)

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show list",
	Long:  `Use this command to print list of your tasks.`,
	Run: func(cmd *cobra.Command, args []string) {
		// local.InitCSVFile(local.TableFilename)
		err := table.RenderFromCSV(local.TableFilename)
		if err != nil {
			fmt.Println("error:", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(showCmd)
}
