package models

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

// the application's state.
type model struct {
	cursor int
	months []string
}

func InitialModel() model {
	return model{
		months: []string{"November", "December"},
	}
}

func (m model) Init() tea.Cmd {
	// No initialization command needed, so return nil
	return nil
}

// Update handles user input and updates the model state.
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "up":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down":
			if m.cursor < len(m.months)-1 {
				m.cursor++
			}
		case "enter":
			fmt.Printf("You selected: %s\n", m.months[m.cursor])
			selectedMonth := m.months[m.cursor]
			//return m, tea.Quit
			return NewBudgetModel(2024, selectedMonth), nil
		}
	}
	return m, nil
}

// render the user interface
func (m model) View() string {
	s := "Select an option:\n\n"
	for i, item := range m.months {
		cursor := " " // no cursor
		if m.cursor == i {
			cursor = "👉" // cursor at the selected item
		}
		s += fmt.Sprintf("%s %s\n", cursor, item)
	}
	s += "\nPress q to quit."
	return s
}
