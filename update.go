package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

func (model *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	oldInput := model.searchInput.Value()
	var inputCmd tea.Cmd
	model.searchInput, inputCmd = model.searchInput.Update(msg)
	cmds = append(cmds, inputCmd)

	newInput := model.searchInput.Value()
	if newInput != oldInput && newInput != "" {
		return model, model.searchOptions(newInput)
	}
	if newInput != oldInput && newInput == "" {
		return model, model.delimitOptions()
	}

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		model.width = msg.Width
		model.height = msg.Height
		model.searchInput.Width = msg.Width
		model.optionsTable.SetWidth(msg.Width)
		model.optionsTable.SetHeight(msg.Height - 1)
		model.optionsTable.Columns()[0].Width = msg.Width
	case errMsg:
		model.errorActive = true
		model.errMsg = msg.Error()
		return model, nil
	case optionsExistMsg:
		if !msg.exists {
			return model, model.createDefaultOptions()
		}
		return model, model.getOptions()
	case optionsMsg:
		model.options = msg.options
		model.optionNames = msg.optionNames
		model.selectedNames = msg.optionNames
		return model, model.createTable()
	case delimitedOptionsMsg:
		model.selectedNames = msg.optionNames
		return model, model.createTable()
	case searchMsg:
		model.selectedNames = msg.searchResults
		return model, model.createTable()
	case tableCreatedMsg:
		model.optionsTable.SetRows(msg.tableRows)
		model.optionsTable.GotoTop()
		return model, nil
	case processStartedMsg:
		return model, tea.Quit
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "ctrl+q":
			return model, tea.Quit
		}
		switch msg.Type {
		case tea.KeyEsc:
			if !model.errorActive {
				return model, tea.Quit
			}
			model.errorActive = false
			return model, nil
		case tea.KeyUp, tea.KeyDown:
			var cmd tea.Cmd
			model.optionsTable, cmd = model.optionsTable.Update(msg)
			return model, cmd
		case tea.KeyTab:
			model.displayMode = (model.displayMode + 1) % 3
			return model, model.delimitOptions()
		case tea.KeyEnter:
			return model, model.launchProcess(model.optionsTable.SelectedRow()[0])
		}
	}
	return model, tea.Batch(cmds...)
}
