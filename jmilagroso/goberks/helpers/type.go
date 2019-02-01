// type.go
// Type conversions
// Jay Milagroso <jmilagroso@quadx.xyz> / Jan 31 2019

package helpers

import "strconv"

func StrToInt64(str string) int64 {

	i, err := strconv.ParseInt(str, 10, 64)

	Error(err)

	return i
}

func StrToInt(str string) int {
	i, err := strconv.Atoi(str)

	Error(err)

	return i
}
