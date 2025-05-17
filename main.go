package vtune

const (
	baseThreshold = 50
)

func getAutovacuumThreshold(baseThreshold, scaleFactor, tuples float64) float64 {
	return baseThreshold + (scaleFactor * tuples)
}

func getAutovacuumPerDay(threshold, dailyUpdateOrDelete float64) float64 {
	return dailyUpdateOrDelete / threshold
}

func getScaleFactorForDailyVacuum(tuples, dailyUpdateOrDelete float64) float64 {
	return (dailyUpdateOrDelete - baseThreshold) / tuples
}
