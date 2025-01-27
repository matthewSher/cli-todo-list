package local

import (
	constants "cli-todo-list/config"
	"fmt"
)

var (
	ErrWriteData  = fmt.Errorf("failed to write data into %s", constants.TableFileName)
	ErrReadData   = fmt.Errorf("failed to read data from %s", constants.TableFileName)
	ErrAddElement = fmt.Errorf("failed to add new element")
)
