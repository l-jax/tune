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

	if m, ok := m.(model); ok && m.scaleFactor != 0 {
		fmt.Printf("Set your scale factor to %f for a daily vacuum\n", m.scaleFactor)
	}
}
