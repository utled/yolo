package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Option struct {
	Type          string `json:"type"`
	CommandOrPath string `json:"commandorpath"`
}

type displayMode int

const (
	combinedDisplay displayMode = iota
	programDisplay
	configDisplay
)

type Model struct {
	width         int
	height        int
	displayMode   displayMode
	errorActive   bool
	errMsg        string
	filepath      string
	options       map[string]Option
	optionNames   []string
	optionsTable  table.Model
	selectedNames []string
	searchInput   textinput.Model
}

func NewModel() Model {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("could not find user home dir:\n", err)
	}
	filepath := filepath.Join(homeDir, "/.yolo/optionFile.json")

	searchInputField := textinput.New()
	searchInputField.Width = 20
	searchInputField.Placeholder = ""
	searchInputField.Prompt = ""
	searchInputField.Focus()

	optionsTable := table.New(
		table.WithColumns([]table.Column{{Title: "", Width: 20}}),
		table.WithFocused(true),
		table.WithHeight(10),
		table.WithRows([]table.Row{}),
	)
	tableStyle := table.DefaultStyles()
	tableStyle.Selected = tableStyle.Header.
		BorderForeground(lipgloss.Color("238")).
		Background(lipgloss.Color("234")).
		PaddingLeft(0).
		Bold(true)
	optionsTable.SetStyles(tableStyle)

	return Model{
		filepath:     filepath,
		searchInput:  searchInputField,
		optionsTable: optionsTable,
	}
}

func (model *Model) Init() tea.Cmd {
	return model.checkIfOptionsExist()
}
