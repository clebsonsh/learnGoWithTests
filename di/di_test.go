package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Clebson")

	got := buffer.String()
	want := "Hello, Clebson"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
