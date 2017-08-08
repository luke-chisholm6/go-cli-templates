package readers

import "errors"

type ErrorReader struct{}

func (reader *ErrorReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("I always error huehue")
}

func NewErrorReader() *ErrorReader {
	return &ErrorReader{}
}
