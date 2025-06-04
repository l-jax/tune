package main

import (
	"errors"
	"fmt"
	"github.com/charmbracelet/huh"
	"os"
	"strconv"
)

var (
	rows      string
	tuples    string
	frequency float64
	ready     bool
)

func main() {
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("How many rows does your table have?").
				Value(&rows).
				Validate(func(str string) error {
					_, err := strconv.Atoi(str)
					if err != nil {
						return errors.New("not a number")
					}
					return nil
				}),
			huh.NewInput().
				Title("How many dead tuples per day?").
				Value(&tuples).
				Validate(func(str string) error {
					_, err := strconv.Atoi(str)
					if err != nil {
						return errors.New("not a number")
					}
					return nil
				}),
			huh.NewSelect[float64]().
				Title("How often do you want to vacuum it?").
				Options(
					huh.NewOption("Every day", 1.0),
					huh.NewOption("Twice a day", 0.5),
				).
				Value(&frequency),
			huh.NewConfirm().
				Title("Ready to fix your autovacuum?").
				Value(&ready),
		),
	)

	err := form.Run()
	if err != nil {
		fmt.Println("Uh oh:", err)
		os.Exit(1)
	}
	calculateOutput()
}

func calculateOutput() {
	r, _ := strconv.ParseUint(rows, 10, 64)
	t, _ := strconv.ParseUint(tuples, 10, 64)
	table, _ := NewTable(uint(r), uint(t))
	params, err := suggestAutovacuumParameters(*table, 1.0)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Set autovacuum_vacuum_scale_factor to %.4f and autovacuum_vaccuum_threshold to %d for a daily vacuum\n", params.scaleFactor, params.threshold)
}
