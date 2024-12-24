package main

import (
	"bytes"
	"fmt"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}

	name := "Bhuvansh"
	Greet(&buffer, name)

	got := buffer.String()
	want := fmt.Sprintf("Hello, %s", name)

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
