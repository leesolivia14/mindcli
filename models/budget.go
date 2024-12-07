package models

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type budgetModel struct {
	year  int
	month string
	data  []BudgetRecord
}

func NewBudgetModel(year int, month string) budgetModel {
	records := fetchBudgetRecordsFromDB(year, month)
	return budgetModel{
		year:  year,
		month: month,
		data:  records,
	}
}

func (m budgetModel) Init() tea.Cmd {
	return nil
}

func (m budgetModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "q" {
			// Return to the previous model (e.g., the month selector)
			return main.initialModel(), nil
		}
	}
	return m, nil
}

func (m budgetModel) View() string {
	var output strings.Builder
	output.WriteString(fmt.Sprintf("Budget Records for %s %d\n\n", m.month, m.year))

	for _, record := range m.data {
		output.WriteString(fmt.Sprintf("- %s: $%.2f\n", record.Description, record.Amount))
	}

	output.WriteString("\nPress 'q' to go back.\n")
	return output.String()
}
