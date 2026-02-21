package main

import (
	"log"
	"os"
	"path/filepath"

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
	allDisplay displayMode = iota
	programDisplay
	configDisplay
)

type Model struct {
	width        int
	height       int
	displayMode  displayMode
	errorActive  bool
	filepath     string
	options      Options
	programNames []string
	configNames  []string
	programList  list.Model
	configList   list.Model
}

func NewModel() Model {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("could not find user home dir:\n", err)
	}
	filepath := filepath.Join(homeDir, "/.yolo/optionFile.json")
	return Model{
		filepath: filepath,
	}
}

func (model *Model) Init() tea.Cmd {
	return model.checkIfOptionsExist()
}
