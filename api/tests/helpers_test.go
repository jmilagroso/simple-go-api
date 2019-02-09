package tests

import (
	"testing"

	h "github.com/jmilagroso/api/helpers"
)

func TestError(t *testing.T) {
	v := h.Error()
	if v != nil {
		t.Error("Expected some value, got ", v)
	}
}
