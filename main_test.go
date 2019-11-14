package iel

import (
	"testing"
)

func TestReturnFalse(t *testing.T) {
	if returnFalse() != false {
		t.Fatal("expected false")
	}
}
