package main

import (
	"bytes"
	"fmt"
	"testing"
)

func TestRun(t *testing.T) {
	buffer := &bytes.Buffer{}
	Run(buffer)

	got := buffer.String()
	want := fmt.Sprintf("%s:%d\n%s:%d\n%s:%d\n%s:%.4f\n%s:%.4f\n\n",
		tuples, 1000,
		updates, 100,
		autovacuumVacuumThreshold, 50,
		autovacuumVacuumScaleFactor, 0.2,
		vacuumsPerDay, 0.4,
	)

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
