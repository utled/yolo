package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Options struct {
	Programs map[string]Program `json:"programs"`
	Configs  map[string]Config  `json:"configs"`
}

type Program struct {
	Description string `json:"description"`
	RunCommand  string `json:"runcommand"`
}

type Config struct {
	Description string `json:"description"`
	FullPath    string `json:"fullpath"`
}

type listItem struct {
	title       string
	description string
}

func (item listItem) Title() string       { return item.title }
func (item listItem) Description() string { return item.description }
func (item listItem) FilterValue() string { return item.title }

type displayMode int

const (
	combinedDisplay displayMode = iota
	programDisplay
	configDisplay
)

type Model struct {
	width             int
	height            int
	docStyle          lipgloss.Style
	displayMode       displayMode
	errorActive       bool
	filepath          string
	options           Options
	programNames      []string
	configNames       []string
	mainList          list.Model
	combinedListItems []list.Item
	programListItems  []list.Item
	configListItems   []list.Item
}

func NewModel() Model {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("could not find user home dir:\n", err)
	}
	filepath := filepath.Join(homeDir, "/.yolo/optionFile.json")
	mainList := list.New(
		[]list.Item{
			listItem{title: "default", description: "default"},
		},
		list.NewDefaultDelegate(), 0, 0,
	)
	return Model{
		filepath: filepath,
		docStyle: lipgloss.NewStyle().Margin(1, 2),
		mainList: mainList,
	}
}

func (model *Model) Init() tea.Cmd {
	return model.checkIfOptionsExist()
}
