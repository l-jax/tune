package internal

const (
	minScaleFactor = 0.001
	maxScaleFactor = 0.1
)

var thresholds = []int{0, 50, 1000, 5000, 10_000, 50_000, 100_000, 250_000, 500_000, 1_000_000}

func GetAutovacuumParams(tuples, updatesPerDay float64) []Params {
	var params []Params
	for _, threshold := range thresholds {
		scaleFactor := getScaleFactor(tuples, updatesPerDay, threshold)
		if scaleFactor < minScaleFactor || scaleFactor > maxScaleFactor {
			continue
		}
		params = append(params, Params{threshold, scaleFactor})
	}

	return params
}

func getScaleFactor(tuples, updatesPerDay float64, threshold int) float64 {
	return (updatesPerDay - float64(threshold)) / tuples
}

func getThreshold(tuples, updatesPerDay, scaleFactor float64) int {
	return int(updatesPerDay - (scaleFactor * tuples))
}
