package local

import (
	"cli-todo-list/internal/task"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
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
		fmt.Printf("Failed to create %s file\n", TableFileName)
		return nil, err
	}

	csvWriter := csv.NewWriter(csvFile)
	csvReader := csv.NewReader(csvFile)

	// Writing table titles into the csvFile
	tableTitles := []string{"ID", "Description", "Completion"}

	if err := csvWriter.Write(tableTitles); err != nil {
		fmt.Println(ErrWriteData.Error())
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
	csvFile, err := os.Create(TableFileName)
	if err != nil {
		fmt.Println(ErrWriteData.Error())
		return nil, err
	}

	return csvFile, nil
}

func (t *CSVTable) AddElement(task task.Task) {
	data := []string{strconv.Itoa(task.Id), task.Description, strconv.FormatBool(task.IsCompleted)}

	if err := t.writer.Write(data); err != nil {
		fmt.Println(ErrWriteData.Error())
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
