package main

import (
	"flag"
	"fmt"
	"io"
	"vtune/internal"
)

const (
	tuples                      = "tuples"
	updates                     = "updates"
	autovacuumVacuumThreshold   = "autovacuum_vacuum_threshold"
	autovacuumVacuumScaleFactor = "autovacuum_vacuum_scale_factor"
	vacuumsPerDay               = "vacuums per day"
)

func Run(out io.Writer) {
	userInput := getUserInput()
	vacuums := internal.GetVacuumsPerDay(userInput.tuples, userInput.updates, userInput.threshold, userInput.scaleFactor)

	fmt.Fprintf(out, "%s:%d\n%s:%d\n%s:%d\n%s:%.4f\n%s:%.4f\n\n",
		tuples, userInput.tuples,
		updates, userInput.updates,
		autovacuumVacuumThreshold, userInput.threshold,
		autovacuumVacuumScaleFactor, userInput.scaleFactor,
		vacuumsPerDay, vacuums,
	)
}

type UserInput struct {
	tuples, updates, threshold int
	scaleFactor                float64
}

func getUserInput() UserInput {
	var tuplesInput, updatesInput, thresholdInput int
	var scaleFactorInput float64
	flag.IntVar(&tuplesInput, "tuples", 1000, "table size in tuples")
	flag.IntVar(&updatesInput, "updates", 100, "table updates/deletes per day")
	flag.IntVar(&thresholdInput, "threshold", 50, "current autovacuum_vacuum_threshold")
	flag.Float64Var(&scaleFactorInput, "scaleFactor", 0.2, "current autovacuum_vacuum_scale_factor")
	flag.Parse()

	return UserInput{
		tuplesInput,
		updatesInput,
		thresholdInput,
		scaleFactorInput,
	}
}
