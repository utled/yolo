package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/charmbracelet/bubbles/list"
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
	options      Options
	programNames []string
	configNames  []string
}

func (model *Model) createDefaultOptions() tea.Cmd {
	return func() tea.Msg {
		defaultProgram := Program{Description: "standard editor", RunCommand: "nvim"}
		defaultConfig := Config{Description: "programs and configs to launch", FullPath: model.filepath}

		programsMap := make(map[string]Program)
		programsMap["Neovim"] = defaultProgram

		configsMap := make(map[string]Config)
		configsMap["YOLO Options"] = defaultConfig

		defaultOptions := Options{Programs: programsMap, Configs: configsMap}
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
		return optionsMsg{
			options:      defaultOptions,
			programNames: []string{"Neovim"},
			configNames:  []string{"YOLO Options"},
		}
	}
}

func (model *Model) getOptions() tea.Cmd {
	return func() tea.Msg {
		optionFile, err := os.ReadFile(model.filepath)
		if err != nil {
			return errMsg(fmt.Errorf("failed to read option file:\n%v", err))
		}
		var options Options
		if err := json.Unmarshal(optionFile, &options); err != nil {
			return errMsg(fmt.Errorf("failed to unmarshal options json:\n%v", err))
		}

		var programs []string
		var configs []string

		for program := range options.Programs {
			programs = append(programs, program)
		}
		for config := range options.Configs {
			configs = append(configs, config)
		}
		return optionsMsg{
			options:      options,
			programNames: programs,
			configNames:  configs,
		}
	}
}

type listItemsMsg struct {
	programList  []list.Item
	configList   []list.Item
	combinedList []list.Item
}

func (model *Model) createListItems() tea.Cmd {
	return func() tea.Msg {
		var programList []list.Item
		var configList []list.Item
		var combinedList []list.Item

		for _, program := range model.programNames {
			programList = append(
				programList,
				listItem{
					title:       program,
					description: model.options.Programs[program].Description,
				},
			)
		}
		for _, config := range model.configNames {
			configList = append(
				configList,
				listItem{
					title:       config,
					description: model.options.Configs[config].Description,
				},
			)
		}
		combinedList = append(combinedList, programList...)
		combinedList = append(combinedList, configList...)

		return listItemsMsg{
			programList:  programList,
			configList:   configList,
			combinedList: combinedList,
		}
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
