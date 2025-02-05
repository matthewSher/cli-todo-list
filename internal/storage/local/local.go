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

func DeleteElement(filepath string, taskId int) error {
	file, err := os.Open(filepath)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrOpenFile, err)
	}

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()

	// Explicitly close file before reopening it
	file.Close()

	if err != nil {
		return fmt.Errorf("%w: %v", ErrReadData, err)
	}

	if taskId < 1 || taskId > len(records) {
		return ErrOutOfRange
	}

	// Deleting element with ID taskId
	records = append(records[:taskId], records[taskId+1:]...)

	// Reopening file for writing
	file, err = os.OpenFile(filepath, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrOpenFile, err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	if err := writer.WriteAll(records); err != nil {
		return fmt.Errorf("%w: %v", ErrWriteData, err)
	}
	writer.Flush()

	return nil
}
