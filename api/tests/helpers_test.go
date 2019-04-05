package tests

import (
	"testing"

	h "github.com/jmilagroso/api/helpers"
)

func TestHash256(t *testing.T) {
	v := h.Hash256("hello")
	if v == "" {
		t.Error("Expected some value, got ", v)
	}
}
