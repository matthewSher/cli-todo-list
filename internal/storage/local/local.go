package local

import (
	"cli-todo-list/internal/task"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func InitCSVFile(filepath string) error {
	_, err := os.Stat(filepath)

	if os.IsNotExist(err) {
		file, err := os.Create(filepath)

		if err != nil {
			return fmt.Errorf("%w: %v", ErrCreateFile, err)
		}
		defer file.Close()

		writer := csv.NewWriter(file)
		defer writer.Flush()

		if err := writer.Write(TableTitles); err != nil {
			return fmt.Errorf("%w: %v", ErrWriteData, err)
		}
	}

	return nil
}

func AddElement(filepath string, task task.Task) error {
	file, err := os.OpenFile(filepath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrOpenFile, err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	data := []string{task.Description, strconv.FormatBool(task.IsCompleted)}

	if err := writer.Write(data); err != nil {
		return fmt.Errorf("%w: %v", ErrWriteData, err)
	}

	return nil
}
