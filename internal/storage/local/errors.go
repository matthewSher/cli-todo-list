package local

import (
	"fmt"
)

var (
	ErrWriteData  = fmt.Errorf("failed to write data into %s", TableFileName)
	ErrReadData   = fmt.Errorf("failed to read data from %s", TableFileName)
	ErrAddElement = fmt.Errorf("failed to add new element")
)
