package main

import (
	"github.com/charmbracelet/lipgloss"
)

func (model *Model) renderOverlay() string {
	modalStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("238")).
		Padding(1).
		Width(model.width - 2)

	form := lipgloss.JoinVertical(
		lipgloss.Left,
		lipgloss.NewStyle().
			Bold(false).
			Italic(true).
			Width(model.width - modalStyle.GetHorizontalPadding()).
			Render(model.errMsg),
		"",
		lipgloss.NewStyle().Foreground(lipgloss.Color("238")).Render("[Esc] to return"),
	)

	return lipgloss.Place(
		model.width,
		model.height,
		lipgloss.Center,
		lipgloss.Center,
		modalStyle.Render(form),
		lipgloss.WithWhitespaceForeground(lipgloss.Color("238")),
	)
}

func (model *Model) renderMain() string {
	return lipgloss.Place(
		model.width,
		model.height,
		lipgloss.Center,
		lipgloss.Top,
		lipgloss.JoinVertical(
			lipgloss.Left,
			model.searchInput.View(),
			model.optionsTable.View(),
		),
	)
}

func (model *Model) View() string {
	if model.errorActive {
		return model.renderOverlay()
	}

	return model.renderMain()
}
