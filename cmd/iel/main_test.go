package main

import (
	"testing"
)

func TestJustError(t *testing.T) {
	if justError() != nil {
		t.Fatal("expected error")
	}
}
