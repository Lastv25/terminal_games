package main

import (
	"Coding/games/models"
	"fmt"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"strings"
)

// model to display after the choice is made
type displayModel struct {
	selected      string
	inserted_text textinput.Model
}

func (m displayModel) Init() tea.Cmd { return textinput.Blink }

func (m displayModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			fmt.Println(strings.Repeat("*", 100))
			fmt.Printf("%s\n", m.inserted_text.Value())
			m.inserted_text.SetValue("")
			return m, nil
		case "ctrl+c":
			return m, tea.Quit
		default:
			var cmd tea.Cmd
			m.inserted_text, cmd = m.inserted_text.Update(msg)
			return m, cmd
		}
		var cmd tea.Cmd
		m.inserted_text, cmd = m.inserted_text.Update(msg)
		return m, cmd

	}
	return m, nil
}

func (m displayModel) View() string {
	return fmt.Sprintf("%s\n\n", m.inserted_text.View())
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
	fmt.Println(strings.Repeat("-", 100))
	fmt.Println("This a new TUI loop")
	fmt.Printf("You selected: %s\n", sel)

	// creating text input
	ti := textinput.New()
	ti.Placeholder = "Type here"
	ti.Focus()
	ti.CharLimit = 64
	ti.Width = 32

	displayProg := tea.NewProgram(displayModel{selected: sel, inserted_text: ti})
	if err := displayProg.Start(); err != nil {
		fmt.Printf("Error running display model: %v\n", err)
	}
}
