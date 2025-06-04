package main

import (
	"testing"
	"testing/quick"
)

const (
	defaultScaleFactor = 0.2
	defaultThreshold   = 50
	rowsInTable        = 50000
	updatesPerDay      = 40000
)

var scaleFactorTests = map[string]struct {
	daysBetweenVacuums, scaleFactor, vacuumsPerDay float64
}{
	"twice a day":    {daysBetweenVacuums: 0.5, vacuumsPerDay: 2.0},
	"daily":          {daysBetweenVacuums: 1.0, vacuumsPerDay: 1.0},
	"every two days": {daysBetweenVacuums: 2.0, vacuumsPerDay: 0.5},
}

func TestGetVacuumsPerDayWithDefaultParams(t *testing.T) {
	want := 0.4
	got := GetVacuumsPerDay(1000, 100, defaultThreshold, defaultScaleFactor)
	assertFloats(t, got, want)
}

func TestGetScaleFactorForVacuum(t *testing.T) {
	for name, test := range scaleFactorTests {
		t.Run(name, func(t *testing.T) {
			scaleFactor := getScaleFactorForVacuum(rowsInTable, updatesPerDay, defaultThreshold, test.daysBetweenVacuums)
			assertVacuumsPerDay(t, rowsInTable, updatesPerDay, defaultThreshold, scaleFactor, test.vacuumsPerDay)
		})
	}
}

func TestGetThresholdForVacuum(t *testing.T) {
	for name, test := range scaleFactorTests {
		t.Run(name, func(t *testing.T) {
			threshold := getThresholdForVacuum(rowsInTable, updatesPerDay, defaultScaleFactor, test.daysBetweenVacuums)
			assertVacuumsPerDay(t, rowsInTable, updatesPerDay, threshold, defaultScaleFactor, test.vacuumsPerDay)
		})
	}
}

func TestProperties(t *testing.T) {
	var tuples uint = 123456789
	var updates uint = 987654

	assertion := func(threshold uint) bool {
		if threshold > updates {
			return true
		}

		scaleFactor := getScaleFactorForVacuum(tuples, updates, threshold, 0)
		fromScaleFactor := getThresholdForVacuum(tuples, updates, scaleFactor, 1.0)
		return fromScaleFactor == threshold
	}

	if err := quick.Check(assertion, &quick.Config{
		MaxCount: 1000,
	}); err != nil {
		t.Error("Threshold derived from scale factor does not match threshold used to generate scale factor", err)
	}
}

func assertVacuumsPerDay(t *testing.T, tuples, updates, baseThreshold uint, scaleFactor, want float64) {
	t.Helper()
	vacuumsPerDay := GetVacuumsPerDay(tuples, updates, baseThreshold, scaleFactor)
	if vacuumsPerDay != want {
		t.Errorf("got %f, want %f", vacuumsPerDay, want)
	}
}
