package vtune

type Params struct {
	threshold   int
	scaleFactor float64
}

func GetParamsForDailyAutovacuum(tuples, dailyUpdateOrDelete float64) Params {
	return Params{}
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
