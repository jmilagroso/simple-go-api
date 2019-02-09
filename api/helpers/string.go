package helpers

import (
	"bytes"
)

// Concat string
func Concat(buffer bytes.Buffer, str string) bytes.Buffer {
	_, err := buffer.WriteString(str)

	Error(err)

	return buffer
}
