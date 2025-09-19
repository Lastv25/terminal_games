package main

import (
	"Coding/games/models"
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"strings"
)

// model to display after the choice is made
type displayModel struct {
	selected string
}

func (m displayModel) Init() tea.Cmd { return nil }

func (m displayModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "q" || msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m displayModel) View() string {
	fmt.Println(strings.Repeat("-", 100))
	fmt.Println("This a new TUI loop")
	return fmt.Sprintf("You selected: %s\n\nPress q to quit.\n", m.selected)
}
func main() {
	choiceProg := tea.NewProgram(models.InitialModel())
	finalModel, err := choiceProg.Run()
	if err != nil {
		fmt.Printf("Error starting program: %v\n", err)
		return
	}

	choiceM := finalModel.(models.InitMenu)
	if choiceM.Selected() < 0 {
		// No selection made, exit
		return
	}
	// After choice is made, run new Bubble Tea program displaying selection
	sel := choiceM.Choices()[choiceM.Selected()].String()
	displayProg := tea.NewProgram(displayModel{selected: sel})
	if err := displayProg.Start(); err != nil {
		fmt.Printf("Error running display model: %v\n", err)
	}
}
