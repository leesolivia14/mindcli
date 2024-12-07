package main

import (
	"fmt"
	"os"
	"github.com/leesolivia14/mindcli/models"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(models.InitialModel())

	if err := p.Start(); err != nil {
		fmt.Fprintln(os.Stderr, "Error starting program:", err)
		os.Exit(1)
	}
}
