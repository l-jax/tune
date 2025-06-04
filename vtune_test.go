package main

import (
	"reflect"
	"testing"
	"testing/quick"
)

func TestGetVacuumsPerDayWithDefaultParams(t *testing.T) {
	var tuples uint = 1000
	var updates uint = 100
	want := 0.4

	got := GetVacuumsPerDay(tuples, updates, 50, 0.2)
	assertFloats(t, got, want)
}

func TestGetScaleFactorForDailyVacuum(t *testing.T) {
	var tuples uint = 1000
	var updates uint = 100
	want := 0.05

	got := getScaleFactorForDailyVacuum(tuples, updates, 50)

	assertFloats(t, got, want)
}

func TestGetThresholdForDailyVacuum(t *testing.T) {
	var tuples uint = 1000
	var updates uint = 100
	var want uint = 50

	got := getThresholdForDailyVacuum(tuples, updates, 0.05)

	assertInts(t, got, want)
}

func TestGetParamsForDailyVacuum(t *testing.T) {
	var tuples uint = 1000
	var updates uint = 100

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
	var updates uint = 1000
	want := []uint{0, 50, 100, 200, 500, 1000}

	got := getTestThresholds(updates)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestProperties(t *testing.T) {
	var tuples uint = 123456789
	var updates uint = 987654

	assertion := func(threshold uint) bool {
		if threshold > updates {
			return true
		}

		scaleFactor := getScaleFactorForDailyVacuum(tuples, updates, threshold)
		fromScaleFactor := getThresholdForDailyVacuum(tuples, updates, scaleFactor)
		return fromScaleFactor == threshold
	}

	if err := quick.Check(assertion, &quick.Config{
		MaxCount: 1000,
	}); err != nil {
		t.Error("Threshold derived from scale factor does not match threshold used to generate scale factor", err)
	}
}

func assertParams(t *testing.T, got []Params, want []Params) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func assertInts(t *testing.T, got, want uint) {
	t.Helper()
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
