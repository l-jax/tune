package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"os"
)

func main() {
	p := tea.NewProgram(initialModel())

	m, err := p.Run()
	if err != nil {
		fmt.Println("Oh no:", err)
		os.Exit(1)
	}

	if m, ok := m.(model); ok {
		fmt.Printf("Set autovacuum_vacuum_scale_factor to %.4f and autovacuum_vaccuum_threshold to %d for a daily vacuum\n", m.params.scaleFactor, m.params.threshold)
	}
}
