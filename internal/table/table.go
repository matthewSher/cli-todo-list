package table

import (
	"cli-todo-list/internal/storage/local"
	"encoding/csv"
	"fmt"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

func RenderFromCSV(filepath string) error {
	file, err := os.Open(filepath)
	if err != nil {
		return fmt.Errorf("%w: %v", local.ErrOpenFile, err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return fmt.Errorf("%w: %v", local.ErrReadData, err)
	}

	if len(records) <= 1 {
		fmt.Println("Todo list is empty")
		return nil
	}

	// Rendering the table
	generateTable(records)

	return nil
}

func generateTable(tasks [][]string) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)

	t.SetTitle("List of tasks")
	t.AppendHeader(convertToTableRow(tasks[0]))
	t.SetAutoIndex(true)
	t.SetStyle(setTableStyle())

	for _, task := range tasks[1:] {
		t.AppendRow(convertToTableRow(task))
	}

	t.Render()
}

func convertToTableRow(slice []string) table.Row {
	result := make(table.Row, len(slice))
	for i, v := range slice {
		result[i] = v
	}
	return result
}

func setTableStyle() table.Style {
	style := table.StyleColoredBright
	style.Title.Align = text.AlignCenter

	return style
}
