package vtune

import "testing"

var thresholdTests = map[string]struct {
	params       Params
	tuples, want float64
}{
	"defaults":            {params: Params{threshold: 50, scaleFactor: 0.2}, tuples: 1000, want: 250.0},
	"zero base threshold": {params: Params{threshold: 0, scaleFactor: 0.2}, tuples: 1000, want: 200.0},
	"zero scale factor":   {params: Params{threshold: 50, scaleFactor: 0}, tuples: 1000, want: 50.0},
	"zero tuples":         {params: Params{threshold: 50, scaleFactor: 0.2}, tuples: 0, want: 50.0},
}

func TestGetAutovacuumThreshold(t *testing.T) {
	for name, test := range thresholdTests {
		t.Run(name, func(t *testing.T) {
			got := test.params.getAutovacuumThreshold(test.tuples)
			assertFloats(t, got, test.want)
		})
	}
}

var frequencyTests = map[string]struct {
	params                            Params
	tuples, dailyUpdateOrDelete, want float64
}{
	"fewer than one per day":   {params: Params{threshold: 50, scaleFactor: 0.2}, tuples: 1000, dailyUpdateOrDelete: 50, want: 0.2},
	"greater than one per day": {params: Params{threshold: 50, scaleFactor: 0.2}, tuples: 1000, dailyUpdateOrDelete: 500, want: 2},
	"one per day":              {params: Params{threshold: 50, scaleFactor: 0.2}, tuples: 1000, dailyUpdateOrDelete: 250, want: 1},
}

func TestAutovacuumFrequency(t *testing.T) {
	for name, test := range frequencyTests {
		t.Run(name, func(t *testing.T) {
			got := test.params.getAutovacuumPerDay(test.tuples, test.dailyUpdateOrDelete)
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
