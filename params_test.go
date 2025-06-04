package main

import "testing"

var thresholdTests = map[string]struct {
	params       Params
	tuples, want uint
}{
	"defaults":                {params: Params{baseThreshold: 50, scaleFactor: 0.2}, tuples: 1000, want: 250},
	"zero base baseThreshold": {params: Params{baseThreshold: 0, scaleFactor: 0.2}, tuples: 1000, want: 200},
	"zero scale factor":       {params: Params{baseThreshold: 50, scaleFactor: 0}, tuples: 1000, want: 50},
	"zero tuples":             {params: Params{baseThreshold: 50, scaleFactor: 0.2}, tuples: 0, want: 50},
}

func TestGetThreshold(t *testing.T) {
	for name, test := range thresholdTests {
		t.Run(name, func(t *testing.T) {
			got := test.params.getThreshold(test.tuples)
			if got != test.want {
				t.Errorf("got %d, want %d", got, test.want)
			}
		})
	}
}

var frequencyTests = map[string]struct {
	params          Params
	tuples, updates uint
	want            float64
}{
	"fewer than one per day":   {params: Params{baseThreshold: 50, scaleFactor: 0.2}, tuples: 1000, updates: 50, want: 0.2},
	"greater than one per day": {params: Params{baseThreshold: 50, scaleFactor: 0.2}, tuples: 1000, updates: 500, want: 2},
	"one per day":              {params: Params{baseThreshold: 50, scaleFactor: 0.2}, tuples: 1000, updates: 250, want: 1},
}

func TestGetFrequency(t *testing.T) {
	for name, test := range frequencyTests {
		t.Run(name, func(t *testing.T) {
			got := test.params.GetFrequency(test.tuples, test.updates)
			assertFloats(t, got, test.want)
		})
	}
}

func assertFloats(t *testing.T, got, want float64) {
	t.Helper()
	if got != want {
		t.Errorf("got %f, want %f", got, want)
	}
}
