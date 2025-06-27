package main

import (
	"fmt"
)

const defaultThreshold = 50

var (
	ErrMustBeNumeric         = fmt.Errorf("must be a number")
	ErrMustBeGreaterThanZero = fmt.Errorf("must be greater than zero")
	ErrMustNotBeNegative     = fmt.Errorf("must not be negative")
)

type Table struct {
	numberOfRows  uint64
	updatesPerDay uint64
}

type Params struct {
	scaleFactor float64
	threshold   uint64
}

func suggestAutovacuumParameters(table Table, daysBetweenVacuums float64) (*Params, error) {
	var initialThreshold uint64 = defaultThreshold

	if table.updatesPerDay < initialThreshold {
		initialThreshold = table.updatesPerDay
	}

	scaleFactor, err := calculateScaleFactor(table, initialThreshold, daysBetweenVacuums)

	if err != nil {
		return nil, err
	}

	if scaleFactor < 0 {
		return nil, ErrMustNotBeNegative
	}

	if scaleFactor > 0.001 {
		return &Params{scaleFactor, initialThreshold}, nil
	}

	threshold, err := calculateThreshold(table, 0, daysBetweenVacuums)
	return &Params{0, threshold}, nil
}

func calculateScaleFactor(table Table, baseThreshold uint64, daysBetweenVacuums float64) (float64, error) {
	if daysBetweenVacuums <= 0 {
		return 0, fmt.Errorf("days between vacuums %w", ErrMustBeGreaterThanZero)
	}

	updatesBeforeVacuum := float64(table.updatesPerDay) * daysBetweenVacuums
	return (updatesBeforeVacuum - float64(baseThreshold)) / float64(table.numberOfRows), nil
}

func calculateThreshold(table Table, scaleFactor, daysBetweenVacuums float64) (uint64, error) {
	if daysBetweenVacuums <= 0 {
		return 0, fmt.Errorf("days between vacuums %w", ErrMustBeGreaterThanZero)
	}

	if scaleFactor < 0 {
		return 0, fmt.Errorf("scale factor %w", ErrMustNotBeNegative)
	}

	updatesBeforeVacuum := float64(table.updatesPerDay) * daysBetweenVacuums
	return uint64(updatesBeforeVacuum - (scaleFactor * float64(table.numberOfRows))), nil
}
