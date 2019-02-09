package helpers

import "strconv"

// StrToInt64 converts string to int64
func StrToInt64(str string) int64 {

	i, err := strconv.ParseInt(str, 10, 64)

	Error(err)

	return i
}

// StrToInt converts string to int
func StrToInt(str string) int {
	i, err := strconv.Atoi(str)

	Error(err)

	return i
}
