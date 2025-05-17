package internal

import "testing"

var thresholdTests = map[string]struct {
	params       Params
	tuples, want float64
}{
	"defaults":                {params: Params{baseThreshold: 50, scaleFactor: 0.2}, tuples: 1000, want: 250.0},
	"zero base baseThreshold": {params: Params{baseThreshold: 0, scaleFactor: 0.2}, tuples: 1000, want: 200.0},
	"zero scale factor":       {params: Params{baseThreshold: 50, scaleFactor: 0}, tuples: 1000, want: 50.0},
	"zero tuples":             {params: Params{baseThreshold: 50, scaleFactor: 0.2}, tuples: 0, want: 50.0},
}

func TestGetThreshold(t *testing.T) {
	for name, test := range thresholdTests {
		t.Run(name, func(t *testing.T) {
			got := test.params.getThreshold(test.tuples)
			assertFloats(t, got, test.want)
		})
	}
}

var frequencyTests = map[string]struct {
	params                      Params
	tuples, updatesPerDay, want float64
}{
	"fewer than one per day":   {params: Params{baseThreshold: 50, scaleFactor: 0.2}, tuples: 1000, updatesPerDay: 50, want: 0.2},
	"greater than one per day": {params: Params{baseThreshold: 50, scaleFactor: 0.2}, tuples: 1000, updatesPerDay: 500, want: 2},
	"one per day":              {params: Params{baseThreshold: 50, scaleFactor: 0.2}, tuples: 1000, updatesPerDay: 250, want: 1},
}

func TestGetFrequency(t *testing.T) {
	for name, test := range frequencyTests {
		t.Run(name, func(t *testing.T) {
			got := test.params.getFrequency(test.tuples, test.updatesPerDay)
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
