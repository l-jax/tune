package internal

const (
	minScaleFactor = 0.0001
	maxScaleFactor = 1
)

var testThresholds = []int{0, 50, 1000, 5000, 10_000, 50_000, 100_000, 250_000}

func GetVacuumsPerDay(tuples, updates, scaleFactor float64, threshold int) float64 {
	params := Params{threshold, scaleFactor}
	return params.GetFrequency(tuples, updates)
}

func GetParamsForDailyVacuum(tuples, updatesPerDay float64) []Params {
	var params []Params
	for _, threshold := range testThresholds {
		scaleFactor := getScaleFactorForDailyVacuum(tuples, updatesPerDay, threshold)
		if scaleFactor < minScaleFactor || scaleFactor > maxScaleFactor {
			continue
		}
		params = append(params, Params{threshold, scaleFactor})
	}

	return params
}

func getScaleFactorForDailyVacuum(tuples, updatesPerDay float64, threshold int) float64 {
	return (updatesPerDay - float64(threshold)) / tuples
}

func getThresholdForDailyVacuum(tuples, updatesPerDay, scaleFactor float64) int {
	return int(updatesPerDay - (scaleFactor * tuples))
}
