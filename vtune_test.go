package main

import (
	"testing"
	"testing/quick"
)

const (
	defaultScaleFactor = 0.2
	defaultThreshold   = 50
)

var testTable = Table{50000, 40000}

var scaleFactorTests = map[string]struct {
	daysBetweenVacuums, scaleFactor, vacuumsPerDay float64
}{
	"twice a day":    {daysBetweenVacuums: 0.5, vacuumsPerDay: 2.0},
	"daily":          {daysBetweenVacuums: 1.0, vacuumsPerDay: 1.0},
	"every two days": {daysBetweenVacuums: 2.0, vacuumsPerDay: 0.5},
}

func TestGetScaleFactorForVacuum(t *testing.T) {
	for name, test := range scaleFactorTests {
		t.Run(name, func(t *testing.T) {
			scaleFactor := getScaleFactorForVacuum(testTable, defaultThreshold, test.daysBetweenVacuums)
			assertVacuumsPerDay(t, defaultThreshold, scaleFactor, test.vacuumsPerDay)
		})
	}
}

func TestGetThresholdForVacuum(t *testing.T) {
	for name, test := range scaleFactorTests {
		t.Run(name, func(t *testing.T) {
			threshold := getThresholdForVacuum(testTable, defaultScaleFactor, test.daysBetweenVacuums)
			assertVacuumsPerDay(t, threshold, defaultScaleFactor, test.vacuumsPerDay)
		})
	}
}

func TestProperties(t *testing.T) {
	assertion := func(threshold uint) bool {
		if threshold > testTable.updatesPerDay {
			return true
		}

		scaleFactor := getScaleFactorForVacuum(testTable, threshold, 0)
		fromScaleFactor := getThresholdForVacuum(testTable, scaleFactor, 1.0)
		return fromScaleFactor == threshold
	}

	if err := quick.Check(assertion, &quick.Config{
		MaxCount: 1000,
	}); err != nil {
		t.Error("Threshold derived from scale factor does not match threshold used to generate scale factor", err)
	}
}

func assertVacuumsPerDay(t *testing.T, baseThreshold uint, scaleFactor, want float64) {
	t.Helper()
	autovacuumThreshold := (scaleFactor * float64(testTable.numberOfRows)) + float64(baseThreshold)
	vacuumsPerDay := float64(testTable.updatesPerDay) / autovacuumThreshold
	if vacuumsPerDay != want {
		t.Errorf("got %f, want %f", vacuumsPerDay, want)
	}
}
