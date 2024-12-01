package main

import tea "github.com/charmbracelet/bubbletea"

// the application's state.
type monthsModel struct {
	cursor int
	years  []string
	months []string
}

func initialMonthsModel() model {
	return model{
		years:  []string{"2024"}, // todo: use datetime to auto update this list
		months: []string{"November"}
	}
}

func (m monthsModel) Init() tea.Cmd {
	// No initialization command needed, so return nil
	return nil
}

// Update handles user input and updates the model state.
func (m monthsModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

}

// render the user interface
func (m monthsModel) View() string {

}
