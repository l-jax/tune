package vtune

import "testing"

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

func TestGetParamsForDailyVacuum(t *testing.T) {
	tuples := 1000.0
	dailyUpdateOrDelete := 100.0

	want := []Params{
		{0, 0.1},
		{50, 0.05},
	}

	got := GetParamsForDailyAutovacuum(tuples, dailyUpdateOrDelete)

	assertParams(t, got, want)
}

func assertParams(t *testing.T, got []Params, want []Params) {
	t.Helper()
	for i := range got {
		if got[i] != want[i] {
			t.Errorf("got %v, want %v", got[i], want[i])
		}
	}
}

func assertInts(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
