package helpers

import (
	"crypto/sha256"
	"encoding/hex"
)

// Hash256 hash
func Hash256(str string) string {

	h := sha256.New()
	_, err := h.Write([]byte(str))
	Error(err)

	return hex.EncodeToString(h.Sum(nil))
}
