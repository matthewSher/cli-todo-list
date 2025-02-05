package local

import "errors"

var (
	ErrWriteData  = errors.New("failed to write data into file")
	ErrReadData   = errors.New("failed to read data from file")
	ErrAddElement = errors.New("failed to add new element")
	ErrCreateFile = errors.New("failed to create file")
	ErrOpenFile   = errors.New("failed to open file")
	ErrOutOfRange = errors.New("index value out of range")
)
