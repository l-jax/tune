package internal

import (
	"reflect"
	"testing"
)

func TestGetVacuumsPerDayWithDefaultParams(t *testing.T) {
	tuples := 1000
	updates := 100
	want := 0.4

	got := GetVacuumsPerDay(tuples, updates, 50, 0.2)
	assertFloats(t, got, want)
}

func TestGetScaleFactorForDailyVacuum(t *testing.T) {
	tuples := 1000
	updates := 100
	want := 0.05

	got := getScaleFactorForDailyVacuum(tuples, updates, 50)

	assertFloats(t, got, want)
}

func TestGetThresholdForDailyVacuum(t *testing.T) {
	tuples := 1000
	updates := 100
	want := 50

	got := getThresholdForDailyVacuum(tuples, updates, 0.05)

	assertInts(t, got, want)
}

func TestGetParamsForDailyVacuum(t *testing.T) {
	tuples := 1000
	updates := 100

	want := []Params{
		{0, 0.1},
		{5, 0.095},
		{10, 0.09},
		{20, 0.08},
		{50, 0.05},
		{100, 0},
	}

	got := GetParamsForDailyVacuum(tuples, updates)

	assertParams(t, got, want)
}

func TestGetTestThresholds(t *testing.T) {
	updates := 1000
	want := []int{0, 50, 100, 200, 500, 1000}

	got := getTestThresholds(updates)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func assertParams(t *testing.T, got []Params, want []Params) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func assertInts(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
