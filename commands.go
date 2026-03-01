package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
)

type errMsg error

type optionsExistMsg struct {
	exists bool
}

func (model *Model) checkIfOptionsExist() tea.Cmd {
	return func() tea.Msg {
		_, err := os.Stat(model.filepath)
		if err != nil {
			return optionsExistMsg{exists: false}
		}

		return optionsExistMsg{exists: true}
	}
}

type optionsMsg struct {
	options     map[string]Option
	optionNames []string
}

func (model *Model) createDefaultOptions() tea.Cmd {
	return func() tea.Msg {
		defaultOptions := make(map[string]Option)
		defaultOptions["Neovim"] = Option{Type: "program", CommandOrPath: "nvim"}
		defaultOptions["YOLO Options"] = Option{Type: "config", CommandOrPath: model.filepath}

		jsonOptions, _ := json.MarshalIndent(defaultOptions, "", "  ")

		dirPath := filepath.Dir(model.filepath)
		if _, err := os.Stat(dirPath); err != nil {
			err = os.Mkdir(dirPath, 0o755)
			if err != nil {
				return fmt.Errorf("failed to create .yolo dir:\n%v", err)
			}
		}

		file, err := os.Create(model.filepath)
		if err != nil {
			return fmt.Errorf("failed to create new file:\n%v", err)
		}
		defer file.Close()

		_, err = file.WriteString(string(jsonOptions))
		if err != nil {
			return fmt.Errorf("failed to write json to file\n%v", err)
		}
		file.Sync()

		var optionNames []string
		for key := range defaultOptions {
			optionNames = append(optionNames, key)
		}
		slices.Sort(optionNames)
		return optionsMsg{
			options:     defaultOptions,
			optionNames: optionNames,
		}
	}
}

func (model *Model) getOptions() tea.Cmd {
	return func() tea.Msg {
		optionFile, err := os.ReadFile(model.filepath)
		if err != nil {
			return errMsg(fmt.Errorf("failed to read option file:\n%v", err))
		}
		options := make(map[string]Option)
		if err := json.Unmarshal(optionFile, &options); err != nil {
			return errMsg(fmt.Errorf("failed to unmarshal options json:\n%v", err))
		}

		var optionNames []string

		for option := range options {
			optionNames = append(optionNames, option)
		}
		slices.Sort(optionNames)
		return optionsMsg{
			options:     options,
			optionNames: optionNames,
		}
	}
}

type delimitedOptionsMsg struct {
	optionNames []string
}

func (model *Model) delimitOptions() tea.Cmd {
	return func() tea.Msg {
		var delimiter string
		var optionNames []string
		switch model.displayMode {
		case combinedDisplay:
			for key := range model.options {
				optionNames = append(optionNames, key)
			}
			slices.Sort(optionNames)
			return delimitedOptionsMsg{optionNames: optionNames}
		case programDisplay:
			delimiter = "program"
		case configDisplay:
			delimiter = "config"
		}
		for key, value := range model.options {
			if value.Type == delimiter {
				optionNames = append(optionNames, key)
			}
		}
		slices.Sort(optionNames)
		return delimitedOptionsMsg{optionNames: optionNames}
	}
}

type searchMsg struct {
	searchResults []string
}

func (model *Model) searchOptions(searchValue string) tea.Cmd {
	return func() tea.Msg {
		var searchResults []string
		for _, option := range model.optionNames {
			if strings.Contains(strings.ToLower(option), strings.ToLower(searchValue)) {
				searchResults = append(searchResults, option)
			}
		}
		slices.Sort(searchResults)
		return searchMsg{searchResults: searchResults}
	}
}

type tableCreatedMsg struct {
	tableRows []table.Row
}

func (model *Model) createTable() tea.Cmd {
	return func() tea.Msg {
		var tableRows []table.Row
		for _, option := range model.selectedNames {
			tableRows = append(tableRows, table.Row{option})
		}
		return tableCreatedMsg{tableRows: tableRows}
	}
}

type processStartedMsg struct{}

func (model *Model) launchProgram(launchCommand string) tea.Cmd {
	return func() tea.Msg {
		// launch program command
		return processStartedMsg{}
	}
}

func (model *Model) launchConfig(filepath string) tea.Cmd {
	return func() tea.Msg {
		// launch Nvim with specified file
		return processStartedMsg{}
	}
}
