package main

import (
	"errors"
	"testing"
	"testing/quick"
)

const (
	defaultScaleFactor        = 0.2
	defaultDaysBetweenVacuums = 1
)

var testTable = Table{50000, 40000}

var scaleFactorTests = map[string]struct {
	daysBetweenVacuums, scaleFactor, vacuumsPerDay float64
}{
	"twice a day":    {daysBetweenVacuums: 0.5, vacuumsPerDay: 2.0},
	"daily":          {daysBetweenVacuums: 1.0, vacuumsPerDay: 1.0},
	"every two days": {daysBetweenVacuums: 2.0, vacuumsPerDay: 0.5},
}

func TestNewTableEmpty(t *testing.T) {
	_, err := NewTable(0, 1000)
	assertError(t, err, ErrEmptyTable)
}

func TestNewTableNoUpdates(t *testing.T) {
	_, err := NewTable(10000, 0)
	assertError(t, err, ErrNoUpdates)
}

func TestSuggestAutovacuumParametersTableWithManyActiveRows(t *testing.T) {
	table := Table{50000, 550}

	wantScaleFactor := 0.01
	var wantThreshold uint64 = 50

	params, _ := suggestAutovacuumParameters(table, defaultDaysBetweenVacuums)

	if params.scaleFactor != wantScaleFactor {
		t.Errorf("scaleFactor = %v, want %v", params.scaleFactor, wantScaleFactor)
	}

	if params.threshold != wantThreshold {
		t.Errorf("threshold = %v, want %v", params.threshold, wantThreshold)
	}
}

func TestSuggestAutovacuumParametersTableWithFewActiveRows(t *testing.T) {
	table := Table{50000, 100}

	wantScaleFactor := 0.0
	var wantThreshold uint64 = 100

	params, _ := suggestAutovacuumParameters(table, defaultDaysBetweenVacuums)

	if params.scaleFactor != wantScaleFactor {
		t.Errorf("scaleFactor = %v, want %v", params.scaleFactor, wantScaleFactor)
	}

	if params.threshold != wantThreshold {
		t.Errorf("threshold = %v, want %v", params.threshold, wantThreshold)
	}
}

func TestCalculateScaleFactor(t *testing.T) {
	for name, test := range scaleFactorTests {
		t.Run(name, func(t *testing.T) {
			scaleFactor, _ := calculateScaleFactor(testTable, defaultThreshold, test.daysBetweenVacuums)
			assertVacuumsPerDay(t, defaultThreshold, scaleFactor, test.vacuumsPerDay)
		})
	}
}

func TestCalculateScaleFactorNoDaysBetweenVacuums(t *testing.T) {
	_, err := calculateScaleFactor(testTable, defaultThreshold, 0)
	assertError(t, err, ErrNoDaysBetweenVacuums)
}

func TestCalculateThresholdForVacuum(t *testing.T) {
	for name, test := range scaleFactorTests {
		t.Run(name, func(t *testing.T) {
			threshold, _ := calculateThreshold(testTable, defaultScaleFactor, test.daysBetweenVacuums)
			assertVacuumsPerDay(t, threshold, defaultScaleFactor, test.vacuumsPerDay)
		})
	}
}

func TestCalculateThresholdNoDaysBetweenVacuums(t *testing.T) {
	_, err := calculateThreshold(testTable, defaultScaleFactor, 0)
	assertError(t, err, ErrNoDaysBetweenVacuums)
}

func TestCalculateThresholdNegativeScaleFactor(t *testing.T) {
	_, err := calculateThreshold(testTable, -1, defaultDaysBetweenVacuums)
	assertError(t, err, ErrNegativeScaleFactor)
}

func TestProperties(t *testing.T) {
	assertion := func(threshold uint64) bool {
		if threshold > testTable.updatesPerDay {
			return true
		}

		scaleFactor, _ := calculateScaleFactor(testTable, threshold, defaultDaysBetweenVacuums)
		fromScaleFactor, _ := calculateThreshold(testTable, scaleFactor, defaultDaysBetweenVacuums)
		return fromScaleFactor == threshold
	}

	if err := quick.Check(assertion, &quick.Config{
		MaxCount: 1000,
	}); err != nil {
		t.Error("Threshold derived from scale factor does not match threshold used to generate scale factor", err)
	}
}

func assertVacuumsPerDay(t *testing.T, baseThreshold uint64, scaleFactor, want float64) {
	t.Helper()
	autovacuumThreshold := (scaleFactor * float64(testTable.numberOfRows)) + float64(baseThreshold)
	vacuumsPerDay := float64(testTable.updatesPerDay) / autovacuumThreshold
	if vacuumsPerDay != want {
		t.Errorf("got %f, want %f", vacuumsPerDay, want)
	}
}

func assertError(t *testing.T, got error, want error) {
	t.Helper()
	if !errors.Is(got, want) {
		t.Errorf("got %s, want %s", got, want)
	}
}
