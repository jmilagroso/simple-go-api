package helpers

import "log"

// Error custom logger
func Error(err error) {
	if err != nil {
		log.Println(err)
	}
}
