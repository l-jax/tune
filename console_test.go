package main

import (
	"bytes"
	"testing"
)

func TestRun(t *testing.T) {
	buffer := &bytes.Buffer{}
	Run(buffer)

	got := buffer.String()
	want := greeting

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
