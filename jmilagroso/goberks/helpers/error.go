// error.go
// Error method
// Jay Milagroso <jmilagroso@quadx.xyz> / Jan 28 2019

package helpers

import "log"

func Error(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
