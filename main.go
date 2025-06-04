package main

import (
	"errors"
	"fmt"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
	"os"
	"strconv"
	"strings"
)

var (
	rows      string
	tuples    string
	frequency float64
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
		),
	)

	err := form.Run()
	if err != nil {
		fmt.Println("Uh oh:", err)
		os.Exit(1)
	}

	params, _ := calculateOutput()
	{
		var sb strings.Builder
		keyword := func(s string) string {
			return lipgloss.NewStyle().Foreground(lipgloss.Color("212")).Render(s)
		}
		fmt.Fprintf(&sb,
			"%s\n\nYour table has %s rows with %s updates/day\n\n%s\n\nautovacuum_vacuum_scale_factor %s\n\nautovacuum_vacuum_threshold %s",
			lipgloss.NewStyle().Bold(true).Render("FIX YOUR AUTOVACUUM"),
			keyword(rows),
			keyword(tuples),
			lipgloss.NewStyle().Bold(true).Render("Try this:"),
			keyword(strconv.FormatFloat(params.scaleFactor, 'f', 5, 64)),
			keyword(strconv.FormatUint(params.threshold, 10)),
		)

		fmt.Println(
			lipgloss.NewStyle().
				Width(60).
				BorderStyle(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color("63")).
				Padding(1, 2).
				Render(sb.String()),
		)
	}
}

func calculateOutput() (*Params, error) {
	r, _ := strconv.ParseUint(rows, 10, 64)
	t, _ := strconv.ParseUint(tuples, 10, 64)
	table, _ := NewTable(r, t)
	return suggestAutovacuumParameters(*table, frequency)
}
