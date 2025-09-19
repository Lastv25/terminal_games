package main

import (
	"Coding/games/models"
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	choices  []models.Games
	cursor   int
	selected int
}

func initialModel() model {
	return model{
		choices:  []models.Games{models.Hive, models.Hortis, models.Star_Realms},
		cursor:   0,
		selected: -1, // none selected initially
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case "enter":
			m.selected = m.cursor
			return m, tea.Quit
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m model) View() string {
	s := "Choose an option:\n\n"
	for i, choice := range m.choices {
		cursor := " " // no cursor by default
		if m.cursor == i {
			cursor = ">" // cursor pointer
		}
		selected := " " // no selection by default
		if m.selected == i {
			selected = "x" // selected marker
		}
		s += fmt.Sprintf("%s [%s] %s\n", cursor, selected, choice)
	}
	if m.selected >= 0 {
		s += fmt.Sprintf("\nYou selected: %s\n", m.choices[m.selected])
	} else {
		s += "\nPress Enter to select, q to quit.\n"
	}
	return s
}

func main() {
	p := tea.NewProgram(initialModel())
	if err := p.Start(); err != nil {
		fmt.Printf("Error starting program: %v\n", err)
		return
	}
}
