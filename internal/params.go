package internal

type Params struct {
	baseThreshold int
	scaleFactor   float64
}

func (p *Params) BaseThreshold() int {
	return p.baseThreshold
}

func (p *Params) ScaleFactor() float64 {
	return p.scaleFactor
}

func (p *Params) getThreshold(tuples float64) float64 {
	return (p.scaleFactor * tuples) + float64(p.baseThreshold)
}

func (p *Params) GetFrequency(tuples, updatesPerDay float64) float64 {
	return updatesPerDay / p.getThreshold(tuples)
}
