package helpers

import (
	"crypto/sha256"
	"encoding/base64"
)

var hash256 = sha256.New()

// Hash256 hash
func Hash256(str string) string {

	_, err := hash256.Write(([]byte(str)))
	Error(err)

	return base64.URLEncoding.EncodeToString(hash256.Sum(nil))
}
