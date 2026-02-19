package main

import tea "github.com/charmbracelet/bubbletea"

func (model *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		model.width = msg.Width
		model.height = msg.Height
	case errMsg:
	case optionsMsg:
	case listItemsMsg:
	case processStartedMsg:
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return nil, tea.Quit
		}
		switch msg.Type {
		case tea.KeyEsc:
			return nil, tea.Quit
		case tea.KeyEnter:
			switch model.displayMode {
			case programDisplay:
			case configDisplay:
			}
		case tea.KeyUp:
		case tea.KeyDown:
		}
	}
	return nil, nil
}
