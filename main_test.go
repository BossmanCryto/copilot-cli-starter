package main

import (
	"testing"
)

func TestVersion(t *testing.T) {
	if VERSION == "" {
		t.Error("VERSION should not be empty")
	}
}
