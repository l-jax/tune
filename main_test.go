package vtune

import "testing"

var thresholdTests = map[string]struct {
	baseThreshold, scaleFactor, tuples, want float64
}{
	"defaults":            {50, 0.2, 1000, 250.0},
	"zero base threshold": {0, 0.2, 1000, 200.0},
	"zero scale factor":   {50, 0, 1000, 50.0},
	"zero tuples":         {50, 0.2, 0, 50.0},
}

func TestGetAutovacuumThreshold(t *testing.T) {
	for name, test := range thresholdTests {
		t.Run(name, func(t *testing.T) {
			got := getAutovacuumThreshold(test.baseThreshold, test.scaleFactor, test.tuples)
			assertFloats(t, got, test.want)
		})
	}
}

var frequencyTests = map[string]struct {
	threshold, dailyUpdateOrDelete, want float64
}{
	"fewer than one per day":   {250, 50, 0.2},
	"greater than one per day": {250, 500, 2},
	"one per day":              {250, 250, 1},
}

func TestAutovacuumFrequency(t *testing.T) {
	for name, test := range frequencyTests {
		t.Run(name, func(t *testing.T) {
			got := getAutovacuumPerDay(test.threshold, test.dailyUpdateOrDelete)
			assertFloats(t, got, test.want)
		})
	}
}

func TestGetScaleFactorForDailyVacuum(t *testing.T) {
	tuples := 1000.0
	dailyUpdateOrDelete := 100.0
	want := 0.05

	got := getScaleFactorForDailyVacuum(tuples, dailyUpdateOrDelete, 50)

	assertFloats(t, got, want)
}

func TestGetThresholdForDailyVacuum(t *testing.T) {
	tuples := 1000.0
	dailyUpdateOrDelete := 100.0
	want := 50

	got := getThresholdForDailyVacuum(tuples, dailyUpdateOrDelete, 0.05)

	assertInts(t, got, want)
}

func TestGetStorageParamsForDailyVacuum(t *testing.T) {
	tuples := 1000.0
	dailyUpdateOrDelete := 100.0
	want := Params{50, 0.05}

	got := GetParamsForDailyAutovacuum(tuples, dailyUpdateOrDelete)

	assertParams(t, got, want)
}

func assertParams(t *testing.T, got Params, want Params) {
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func assertFloats(t *testing.T, got, want float64) {
	t.Helper()
	if got != want {
		t.Errorf("got %f, want %f", got, want)
	}
}

func assertInts(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
