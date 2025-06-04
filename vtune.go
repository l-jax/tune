package main

func GetVacuumsPerDay(tuples, updates, threshold uint, scaleFactor float64) float64 {
	params := Params{threshold, scaleFactor}
	return params.GetFrequency(tuples, updates)
}

func getScaleFactorForVacuum(rowsInTable, updatesPerDay, baseThreshold uint, daysBetweenVacuums float64) float64 {
	updatesBeforeVacuum := float64(updatesPerDay) * daysBetweenVacuums
	return (updatesBeforeVacuum - float64(baseThreshold)) / float64(rowsInTable)
}

func getThresholdForVacuum(rowsInTable, updatesPerDay uint, scaleFactor, daysBetweenVacuums float64) uint {
	updatesBeforeVacuum := float64(updatesPerDay) * daysBetweenVacuums
	return uint(updatesBeforeVacuum - (scaleFactor * float64(rowsInTable)))
}
