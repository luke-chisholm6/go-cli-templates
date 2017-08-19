package writers

import "errors"

type ErrorWriter struct {}

func (ErrorWriter) Write(p []byte) (n int, err error) {
	return 0, errors.New("I always error huehue")
}

func NewErrorWriter() *ErrorWriter {
	return &ErrorWriter{}
}

