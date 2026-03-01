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
		switch model.displayMode{
		case combinedDisplay:
			model.selectedNames = model.programNames
			model.selectedNames = append(model.selectedNames, model.configNames...)
		case programDisplay:
			model.selectedNames = model.programNames
		case configDisplay:
			model.selectedNames = model.configNames
		}
		return model, model.createTable()
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
		return model, nil
	case optionsExistMsg:
		if !msg.exists {
			return model, model.createDefaultOptions()
		}
		return model, model.getOptions()
	case optionsMsg:
		model.options = msg.options
		model.programNames = msg.programNames
		model.configNames = msg.configNames
		model.selectedNames = msg.programNames
		model.selectedNames = append(model.selectedNames, msg.configNames...)
		return model, model.createTable()
	case searchMsg:
		model.selectedNames = msg.searchResults
		return model, model.createTable()
	case tableCreatedMsg:
		model.optionsTable.SetRows(msg.tableRows)
		model.optionsTable.GotoTop()
		return model, nil	
	case processStartedMsg:
		return nil, tea.Quit
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return nil, tea.Quit
		}
		switch msg.Type {
		case tea.KeyEsc:
			return nil, tea.Quit
		case tea.KeyUp, tea.KeyDown:
			var cmd tea.Cmd
			model.optionsTable, cmd = model.optionsTable.Update(msg)
			return model, cmd	
		case tea.KeyTab:
			model.displayMode = (model.displayMode + 1) % 3
			switch model.displayMode{
			case combinedDisplay:
				model.selectedNames = model.programNames
				model.selectedNames = append(model.selectedNames, model.configNames...)
			case programDisplay:
				model.selectedNames = model.programNames
			case configDisplay:
				model.selectedNames = model.configNames
			}
			return model, model.createTable()
			
		case tea.KeyEnter:
			switch model.displayMode {
			case programDisplay:
			case configDisplay:
			}
		}
	}
	return model, tea.Batch(cmds...)
}
