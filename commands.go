package main

import tea "github.com/charmbracelet/bubbletea"

type errMsg error

type optionsMsg struct {}

func (model *Model) getOptions() tea.Cmd {
	return func() tea.Msg {
		return optionsMsg{}
	}
}

type listItemsMsg struct {
	listItems string
	listType  string
}

func (model *Model) createListItems() tea.Cmd {
	return func() tea.Msg {
		return listItemsMsg{}
	}
}

type processStartedMsg struct {}

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
