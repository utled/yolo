package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	model := NewModel()
	program := tea.NewProgram(&model, tea.WithAltScreen())
	_, err := program.Run()
	if err != nil {
		log.Fatalf("failed to run bubble tea program: %v", err)
	}
}
