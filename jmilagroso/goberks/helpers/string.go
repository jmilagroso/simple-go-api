// string.go
// String manipulations
// Jay Milagroso <jmilagroso@quadx.xyz> / Jan 31 2019

package helpers

import (
	"bytes"
)

func Concat(buffer bytes.Buffer, str string) bytes.Buffer {
	_, err := buffer.WriteString(str)

	Error(err)

	return buffer
}
