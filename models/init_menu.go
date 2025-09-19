package models

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
)

type InitMenu struct {
	choices  []Games
	cursor   int
	selected int
}

// Getters
func (m *InitMenu) Selected() int {
	return m.selected
}

func (m *InitMenu) Choices() []Games {
	return m.choices
}

// BubbleTea Logic
func InitialModel() InitMenu {
	return InitMenu{
		choices:  []Games{Hive, Hortis, Star_Realms},
		cursor:   0,
		selected: -1, // none selected initially
	}
}

func (m InitMenu) Init() tea.Cmd {
	return nil
}

func (m InitMenu) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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

func (m InitMenu) View() string {
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
		return s
	} else {
		s += "\nPress Enter to select, q to quit.\n"
	}
	return s
}
