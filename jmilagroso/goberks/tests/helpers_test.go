package tests

import (
	"testing"

	h "quadx.xyz/jmilagroso/goberks/helpers"
)

func TestGetEnvValue(t *testing.T) {
	v := h.GetEnvValue("HTTP_SERVER_ADDRESS")
	if v == "" {
		t.Error("Expected some value, got ", v)
	}
}
