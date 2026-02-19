package main

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
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

type displayMode int

const (
	programDisplay displayMode = iota
	configDisplay
)

type Model struct {
	width        int
	height       int
	displayMode  displayMode
	errorActive  bool
	options      Options
	programNames []string
	configNames  []string
	programList  list.Model
	configList   list.Model
}

func NewModel() Model {
	return Model{}
}

func (model *Model) Init() tea.Cmd {
	return nil
}
