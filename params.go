package vtune

type Params struct {
	threshold   int
	scaleFactor float64
}

func (p Params) getAutovacuumThreshold(tuples float64) float64 {
	return (p.scaleFactor * tuples) + float64(p.threshold)
}

func (p Params) getAutovacuumPerDay(tuples, dailyUpdateOrDelete float64) float64 {
	return dailyUpdateOrDelete / p.getAutovacuumThreshold(tuples)
}
