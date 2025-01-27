package local

import (
	constants "cli-todo-list/config"
	"cli-todo-list/pkg/task"
	"encoding/csv"
	"fmt"
	"os"
)

type CSVTable struct {
	file   *os.File
	writer *csv.Writer
	reader *csv.Reader
}

var Table *CSVTable

func InitTable() {
	if Table == nil {
		table, err := createCSVTable()
		if err != nil {
			fmt.Println("Failed to initialize file of CSV table")
			return
		}
		Table = table
	}
}

func createCSVTable() (*CSVTable, error) {
	csvFile, err := createCSVFile()
	if err != nil {
		fmt.Println("Failed to create todolist.csv file")
		return nil, err
	}

	csvWriter := csv.NewWriter(csvFile)
	csvReader := csv.NewReader(csvFile)

	// Writing table titles into the csvFile
	tableTitles := []string{"ID", "Description", "Completion"}

	if err := csvWriter.Write(tableTitles); err != nil {
		fmt.Println("Failed to write data into todolist.csv")
		return nil, err
	}

	csvTable := CSVTable{
		file:   csvFile,
		writer: csvWriter,
		reader: csvReader,
	}

	return &csvTable, nil
}

func createCSVFile() (*os.File, error) {
	csvFile, err := os.Create(constants.TableFileName)
	if err != nil {
		fmt.Println(ErrWriteData.Error())
		return nil, err
	}

	return csvFile, nil
}

func (t *CSVTable) AddElement(task task.Task) {
	// Representation bool type as string
	completion := "X"
	if task.IsCompleted {
		completion = "V"
	}

	data := []string{fmt.Sprint(task.Id), task.Description, completion}

	if err := t.writer.Write(data); err != nil {
		fmt.Println("Failed to write data into todolist.csv")
		return
	}
}

func (t *CSVTable) DefineIdForTask() (int, error) {
	elements, err := t.reader.ReadAll()
	if err != nil {
		fmt.Println(ErrReadData.Error())
		return 0, err
	}

	tableLen := len(elements)
	if tableLen > 1 {
		return tableLen + 1, nil
	}
	return 1, nil
}
