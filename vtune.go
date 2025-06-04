package main

const (
	minScaleFactor = 0
	maxScaleFactor = 1
)

func GetVacuumsPerDay(tuples, updates, threshold uint, scaleFactor float64) float64 {
	params := Params{threshold, scaleFactor}
	return params.GetFrequency(tuples, updates)
}

func GetParamsForDailyVacuum(tuples, updates uint) []Params {
	testThresholds := getTestThresholds(updates)
	var params []Params

	for _, threshold := range testThresholds {
		scaleFactor := getScaleFactorForDailyVacuum(tuples, updates, threshold)
		if scaleFactor < minScaleFactor || scaleFactor > maxScaleFactor {
			continue
		}
		params = append(params, Params{threshold, scaleFactor})
	}

	return params
}

func getScaleFactorForDailyVacuum(tuples, updates, threshold uint) float64 {
	return (float64(updates) - float64(threshold)) / float64(tuples)
}

func getThresholdForDailyVacuum(tuples, updates uint, scaleFactor float64) uint {
	return uint(float64(updates) - (scaleFactor * float64(tuples)))
}

func getTestThresholds(updates uint) []uint {
	return []uint{
		0,
		updates / 20,
		updates / 10,
		updates / 5,
		updates / 2,
		updates,
	}
}
