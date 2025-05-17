package vtune

type Params struct {
	baseThreshold int
	scaleFactor   float64
}

func (p Params) getThreshold(tuples float64) float64 {
	return (p.scaleFactor * tuples) + float64(p.baseThreshold)
}

func (p Params) getFrequency(tuples, updatesPerDay float64) float64 {
	return updatesPerDay / p.getThreshold(tuples)
}
