package main

import (
	"fmt"
)

var ErrEmptyTable = fmt.Errorf("table must have one or more rows")
var ErrNoUpdates = fmt.Errorf("table must have at least one update per day")
var ErrNoDaysBetweenVacuums = fmt.Errorf("days between vacuums must be more than zero")
var ErrNegativeScaleFactor = fmt.Errorf("scale factor must be positive or 0")

type Table struct {
	numberOfRows  uint
	updatesPerDay uint
}

func NewTable(numberOfRows, updatesPerDay uint) (*Table, error) {
	if numberOfRows == 0 {
		return nil, ErrEmptyTable
	}

	if updatesPerDay == 0 {
		return nil, ErrNoUpdates
	}

	return &Table{numberOfRows, updatesPerDay}, nil
}

func calculateScaleFactor(table Table, baseThreshold uint, daysBetweenVacuums float64) (float64, error) {
	if daysBetweenVacuums <= 0 {
		return 0, fmt.Errorf("getting scale factor: %w", ErrNoDaysBetweenVacuums)
	}

	updatesBeforeVacuum := float64(table.updatesPerDay) * daysBetweenVacuums
	return (updatesBeforeVacuum - float64(baseThreshold)) / float64(table.numberOfRows), nil
}

func calculateThreshold(table Table, scaleFactor, daysBetweenVacuums float64) (uint, error) {
	if daysBetweenVacuums <= 0 {
		return 0, fmt.Errorf("getting threshold: %w", ErrNoDaysBetweenVacuums)
	}

	if scaleFactor < 0 {
		return 0, fmt.Errorf("getting threshold: %w", ErrNegativeScaleFactor)
	}

	updatesBeforeVacuum := float64(table.updatesPerDay) * daysBetweenVacuums
	return uint(updatesBeforeVacuum - (scaleFactor * float64(table.numberOfRows))), nil
}
