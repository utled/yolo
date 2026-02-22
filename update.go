package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

func (model *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		model.width = msg.Width
		model.height = msg.Height
		h, _ := model.docStyle.GetFrameSize()
		model.mainList.SetSize(msg.Width-h, 20)
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
		return model, model.createListItems()
	case listItemsMsg:
		model.programListItems = msg.programList
		model.configListItems = msg.configList
		model.combinedListItems = msg.combinedList

		var cmds []tea.Cmd
		var cmd tea.Cmd
		cmds = append(cmds, model.mainList.SetItems(model.combinedListItems))
		model.mainList, cmd = model.mainList.Update(msg)
		cmds = append(cmds, cmd)

		return model, tea.Batch(cmds...)
	case processStartedMsg:
		return nil, tea.Quit
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return nil, tea.Quit
		}
		switch msg.Type {
		case tea.KeyEsc:
			//return nil, tea.Quit
		case tea.KeyTab:
			model.displayMode = (model.displayMode + 1) % 3
			switch model.displayMode{
			case combinedDisplay:
				model.mainList.SetItems(model.combinedListItems)
			case programDisplay:
				model.mainList.SetItems(model.programListItems)
			case configDisplay:
				model.mainList.SetItems(model.configListItems)
			}
			var cmd tea.Cmd
			model.mainList, cmd = model.mainList.Update(msg)

			return model, cmd
			
		case tea.KeyEnter:
			switch model.displayMode {
			case programDisplay:
			case configDisplay:
			}
		case tea.KeyUp:
		case tea.KeyDown:
		}
	}
	var cmd tea.Cmd
	model.mainList, cmd = model.mainList.Update(msg)
	return model, cmd
}
