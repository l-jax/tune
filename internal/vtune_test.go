package internal

import "testing"

func TestGetVacuumsPerDayWithDefaultParams(t *testing.T) {
	tuples := 1000.0
	updates := 100.0
	want := 0.4

	got := GetVacuumsPerDay(tuples, updates, 0.2, 50)
	assertFloats(t, got, want)
}

func TestGetScaleFactorForDailyVacuum(t *testing.T) {
	tuples := 1000.0
	updatesPerDay := 100.0
	want := 0.05

	got := getScaleFactorForDailyVacuum(tuples, updatesPerDay, 50)

	assertFloats(t, got, want)
}

func TestGetThresholdForDailyVacuum(t *testing.T) {
	tuples := 1000.0
	updatesPerDay := 100.0
	want := 50

	got := getThresholdForDailyVacuum(tuples, updatesPerDay, 0.05)

	assertInts(t, got, want)
}

func TestGetParamsForDailyVacuum(t *testing.T) {
	tuples := 1000.0
	updatesPerDay := 100.0

	want := []Params{
		{0, 0.1},
		{50, 0.05},
	}

	got := GetParamsForDailyVacuum(tuples, updatesPerDay)

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
