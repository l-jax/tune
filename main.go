package vtune

const (
	minScaleFactor = 0.001
	maxScaleFactor = 0.1
)

var thresholds = []int{0, 50, 1000, 5000, 10_000, 50_000, 100_000, 250_000, 500_000, 1_000_000}

type Params struct {
	threshold   int
	scaleFactor float64
}

func GetParamsForDailyAutovacuum(tuples, dailyUpdateOrDelete float64) []Params {
	var params []Params
	for _, threshold := range thresholds {
		scaleFactor := getScaleFactorForDailyVacuum(tuples, dailyUpdateOrDelete, threshold)
		if scaleFactor < minScaleFactor || scaleFactor > maxScaleFactor {
			continue
		}
		params = append(params, Params{threshold, scaleFactor})
	}

	return params
}

func getAutovacuumThreshold(baseThreshold, scaleFactor, tuples float64) float64 {
	return baseThreshold + (scaleFactor * tuples)
}

func getAutovacuumPerDay(threshold, dailyUpdateOrDelete float64) float64 {
	return dailyUpdateOrDelete / threshold
}

func getScaleFactorForDailyVacuum(tuples, dailyUpdateOrDelete float64, threshold int) float64 {
	return (dailyUpdateOrDelete - float64(threshold)) / tuples
}

func getThresholdForDailyVacuum(tuples, dailyUpdateOrDelete, scaleFactor float64) int {
	return int(dailyUpdateOrDelete - (scaleFactor * tuples))
}
