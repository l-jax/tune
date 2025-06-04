package main

type Table struct {
	numberOfRows  uint
	updatesPerDay uint
}

func getScaleFactorForVacuum(table Table, baseThreshold uint, daysBetweenVacuums float64) float64 {
	updatesBeforeVacuum := float64(table.updatesPerDay) * daysBetweenVacuums
	return (updatesBeforeVacuum - float64(baseThreshold)) / float64(table.numberOfRows)
}

func getThresholdForVacuum(table Table, scaleFactor, daysBetweenVacuums float64) uint {
	updatesBeforeVacuum := float64(table.updatesPerDay) * daysBetweenVacuums
	return uint(updatesBeforeVacuum - (scaleFactor * float64(table.numberOfRows)))
}
