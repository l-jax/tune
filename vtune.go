package main

import (
	"fmt"
)

const defaultThreshold = 50

var (
	ErrEmptyTable           = fmt.Errorf("table must have one or more rows")
	ErrNoUpdates            = fmt.Errorf("table must have at least one update per day")
	ErrNoDaysBetweenVacuums = fmt.Errorf("days between vacuums must be more than zero")
	ErrNegativeScaleFactor  = fmt.Errorf("scale factor must be positive or 0")
)

type Table struct {
	numberOfRows  uint
	updatesPerDay uint
}

type Params struct {
	scaleFactor float64
	threshold   uint
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

func suggestAutovacuumParameters(table Table, daysBetweenVacuums float64) (*Params, error) {
	scaleFactor, err := calculateScaleFactor(table, 50, daysBetweenVacuums)

	if err != nil {
		return nil, err
	}

	if scaleFactor < 0 {
		return nil, ErrNegativeScaleFactor
	}

	if scaleFactor > 0.001 {
		return &Params{scaleFactor, defaultThreshold}, nil
	}

	threshold, err := calculateThreshold(table, 0, daysBetweenVacuums)

	if err != nil {
		return nil, err
	}

	return &Params{0, threshold}, nil
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
