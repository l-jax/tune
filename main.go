package vtune

func getAutovacuumThreshold(baseThreshold, scaleFactor, tuples float64) float64 {
	return baseThreshold + (scaleFactor * tuples)
}
