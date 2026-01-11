package models

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/lipgloss"
)

// Style definitions
var (
	TitleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#7D56F4")).
			Background(lipgloss.Color("#1a1a1a")).
			Padding(0, 2).
			MarginTop(1).
			MarginBottom(1)

	SelectedItemStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#FFFFFF")).
				Background(lipgloss.Color("#7D56F4")).
				Bold(true).
				Padding(0, 1).
				MarginLeft(2)

	ItemStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#CCCCCC")).
			PaddingLeft(4)

	DescriptionStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#888888")).
				Italic(true).
				PaddingLeft(4)

	HelpStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#626262")).
			MarginTop(1).
			Italic(true)

	BorderStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#7D56F4")).
			Padding(1, 2)

	InputLabelStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FF06B7")).
			Bold(true)

	MessageStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#04B575")).
			PaddingLeft(2)
)

// MenuModel handles game selection with improved UI
type MenuModel struct {
	choices  []Game
	cursor   int
	selected int
}

// NewMenuModel creates a new menu model
func NewMenuModel() MenuModel {
	return MenuModel{
		choices:  []Game{Hive, Hortis, StarRealms},
		cursor:   0,
		selected: -1,
	}
}

func (m MenuModel) Init() tea.Cmd {
	return nil
}

func (m MenuModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
		case "enter", " ":
			m.selected = m.cursor
			return m, tea.Quit
		case "ctrl+c", "q", "esc":
			m.selected = -1
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m MenuModel) View() string {
	var b strings.Builder

	// Title with border
	title := TitleStyle.Render("🎮  GAME SELECTOR  🎮")
	b.WriteString(title)
	b.WriteString("\n\n")

	// Game options
	for i, choice := range m.choices {
		cursor := "   "
		line := ""

		if m.cursor == i {
			cursor = " ▶ "
			line = SelectedItemStyle.Render(fmt.Sprintf("%s", choice.Name))
		} else {
			line = ItemStyle.Render(fmt.Sprintf("%s", choice.Name))
		}

		b.WriteString(cursor + line)
		b.WriteString("\n")

		// Show description for current cursor position
		if m.cursor == i {
			desc := DescriptionStyle.Render(fmt.Sprintf("     %s", choice.Description))
			b.WriteString(desc)
			b.WriteString("\n")
		}
		b.WriteString("\n")
	}

	// Help text
	b.WriteString(HelpStyle.Render("  ↑/↓ or j/k: navigate  •  enter/space: select  •  q/esc: quit"))

	return BorderStyle.Render(b.String())
}

// Getters
func (m MenuModel) Selected() int {
	return m.selected
}

func (m MenuModel) SelectedGame() Game {
	if m.selected >= 0 && m.selected < len(m.choices) {
		return m.choices[m.selected]
	}
	return Game{}
}

// InputModel handles text input for the selected game
type InputModel struct {
	game      Game
	textInput textinput.Model
	messages  []string
}

// NewInputModel creates a new input model for a game
func NewInputModel(game Game) InputModel {
	ti := textinput.New()
	ti.Placeholder = "Enter your command..."
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 60
	ti.PromptStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF06B7"))
	ti.TextStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFFFFF"))

	return InputModel{
		game:      game,
		textInput: ti,
		messages:  []string{},
	}
}

func (m InputModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m InputModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			return m, tea.Quit
		case "enter":
			value := strings.TrimSpace(m.textInput.Value())
			if value != "" {
				m.messages = append(m.messages, value)
				m.textInput.SetValue("")
			}
			return m, nil
		}
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m InputModel) View() string {
	var b strings.Builder

	// Game header
	header := TitleStyle.Render(fmt.Sprintf("%s", m.game.Name))
	b.WriteString(header)
	b.WriteString("\n\n")

	// Message history section
	if len(m.messages) > 0 {
		b.WriteString(InputLabelStyle.Render("Command History:"))
		b.WriteString("\n\n")

		// Display last 8 messages
		start := 0
		if len(m.messages) > 8 {
			start = len(m.messages) - 8
		}

		for i := start; i < len(m.messages); i++ {
			msg := MessageStyle.Render(fmt.Sprintf("  [%d] %s", i+1, m.messages[i]))
			b.WriteString(msg)
			b.WriteString("\n")
		}
		b.WriteString("\n")
	}

	// Input field
	b.WriteString(InputLabelStyle.Render("Command: "))
	b.WriteString(m.textInput.View())
	b.WriteString("\n\n")

	// Help text
	b.WriteString(HelpStyle.Render("  enter: submit  •  esc/ctrl+c: quit"))

	return BorderStyle.Render(b.String())
}

// Getter for messages (useful for game logic)
func (m InputModel) Messages() []string {
	return m.messages
}
