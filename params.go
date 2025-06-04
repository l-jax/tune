package main

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

func (p *Params) getThreshold(tuples int) int {
	return int((p.scaleFactor * float64(tuples)) + float64(p.baseThreshold))
}

func (p *Params) GetFrequency(tuples, updates int) float64 {
	return float64(updates) / float64(p.getThreshold(tuples))
}
