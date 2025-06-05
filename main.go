package main

import (
	"fmt"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
	"os"
	"strconv"
	"strings"
)

var (
	rows, updates, frequency string
)

func main() {
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("How many rows does your table have?").
				Value(&rows).
				Validate(func(str string) error {
					i, err := strconv.ParseUint(str, 10, 64)
					if err != nil {
						return ErrNotANumber
					}
					if i < 1 {
						return ErrEmptyTable
					}
					return nil
				}),
			huh.NewInput().
				Title("How many dead updates per day?").
				Value(&updates).
				Validate(func(str string) error {
					i, err := strconv.ParseUint(str, 10, 64)
					if err != nil {
						return ErrNotANumber
					}
					if i < 1 {
						return ErrNoUpdates
					}
					return nil
				}),
			huh.NewInput().
				Title("How many days do you want between vacuums?").
				Value(&frequency).
				Validate(func(str string) error {
					i, err := strconv.ParseFloat(str, 64)
					if err != nil {
						return ErrNotANumber
					}
					if i < 1 {
						return ErrNoDaysBetweenVacuums
					}
					return nil
				}),
		),
	)

	err := form.Run()
	if err != nil {
		fmt.Println("Uh oh:", err)
		os.Exit(1)
	}

	params, err := calculateOutput()

	if err != nil {
		fmt.Println("Uh oh:", err)
		os.Exit(1)
	}

	{
		var sb strings.Builder
		keyword := func(s string) string {
			return lipgloss.NewStyle().Foreground(lipgloss.Color("212")).Render(s)
		}
		fmt.Fprintf(&sb,
			"%s\n\nYour table has %s rows and %s updates a day\n\nYou want autovacuum to run every %s days\n\n%s\n\nautovacuum_vacuum_scale_factor %s\n\nautovacuum_vacuum_threshold %s",
			lipgloss.NewStyle().Bold(true).Render("FIX MY AUTOVACUUM"),
			keyword(rows),
			keyword(updates),
			keyword(frequency),
			lipgloss.NewStyle().Bold(true).Render("TRY THIS"),
			keyword(strconv.FormatFloat(params.scaleFactor, 'f', 3, 64)),
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
	u, _ := strconv.ParseUint(updates, 10, 64)
	f, _ := strconv.ParseFloat(frequency, 64)
	table, err := NewTable(r, u)

	if err != nil {
		return nil, err
	}

	return suggestAutovacuumParameters(*table, f)
}
