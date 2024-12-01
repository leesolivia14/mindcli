package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

// the application's state.
type model struct {
	cursor int
	items  []string
}

func initialModel() model {
	return model{
		items: []string{"Budget", "Task", "Book", "Idea"},
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
			if m.cursor < len(m.items)-1 {
				m.cursor++
			}
		case "enter":
			fmt.Printf("You selected: %s\n", m.items[m.cursor])
			return m, tea.Quit
		}
	}
	return m, nil
}

// render the user interface
func (m model) View() string {
	s := "Select an option:\n\n"
	for i, item := range m.items {
		cursor := " " // no cursor
		if m.cursor == i {
			cursor = "👉" // cursor at the selected item
		}
		s += fmt.Sprintf("%s %s\n", cursor, item)
	}
	s += "\nPress q to quit."
	return s
}

func main() {
	p := tea.NewProgram(initialModel())

	if err := p.Start(); err != nil {
		fmt.Fprintln(os.Stderr, "Error starting program:", err)
		os.Exit(1)
	}
}
