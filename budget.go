package main

import tea "github.com/charmbracelet/bubbletea"

// the application's state.
type budgetModel struct {
	cursor int
	items  []string
}

func initialBudgetModel() model {
	return model{
		items: []string{"Budget", "Task", "Book", "Idea"},
	}
}

func (m budgetModel) Init() tea.Cmd {
	// No initialization command needed, so return nil
	return nil
}

// Update handles user input and updates the model state.
func (m budgetModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

}

// render the user interface
func (m budgetModel) View() string {

}
