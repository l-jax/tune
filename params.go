package main

type Params struct {
	baseThreshold uint
	scaleFactor   float64
}

func (p *Params) BaseThreshold() uint {
	return p.baseThreshold
}

func (p *Params) ScaleFactor() float64 {
	return p.scaleFactor
}

func (p *Params) getThreshold(tuples uint) uint {
	return uint((p.scaleFactor * float64(tuples)) + float64(p.baseThreshold))
}

func (p *Params) GetFrequency(tuples, updates uint) float64 {
	return float64(updates) / float64(p.getThreshold(tuples))
}
