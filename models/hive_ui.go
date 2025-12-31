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
	PanelStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#7D56F4")).
			Padding(1)
	
	PanelTitleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FF06B7")).
			Bold(true).
			Underline(true)
	
	PieceStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFFFFF")).
			PaddingLeft(1)
	
	BoardStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#7D56F4"))
)
// HiveModel handles the 4-panel Hive game interface
type HiveModel struct {
	game         Game
	textInput    textinput.Model
	messages     []string
	pieces       []string
	board        []string
	width        int
	height       int
}

// NewHiveModel creates a new Hive game model with 4-panel layout
func NewHiveModel(game Game) HiveModel {
	ti := textinput.New()
	ti.Placeholder = "Enter move (e.g., 'place Queen at 0,0')..."
	ti.Focus()
	ti.CharLimit = 100
	ti.Width = 40
	ti.PromptStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF06B7"))
	ti.TextStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFFFFF"))
	
	// Initialize available pieces
	pieces := []string{
		"🐝 Queen Bee (1)",
		"🐜 Ant (3)",
		"🦗 Grasshopper (3)",
		"🕷️  Spider (2)",
		"🪲 Beetle (2)",
	}
	
	// Initialize empty board representation
	board := []string{
		"     Board",
		"",
		"  [Empty Board]",
		"",
		"  Use commands to",
		"  place pieces",
	}
	
	return HiveModel{
		game:      game,
		textInput: ti,
		messages:  []string{},
		pieces:    pieces,
		board:     board,
		width:     120,
		height:    30,
	}
}

func (m HiveModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m HiveModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, nil
		
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			return m, tea.Quit
		case "enter":
			value := strings.TrimSpace(m.textInput.Value())
			if value != "" {
				m.messages = append(m.messages, value)
				m.textInput.SetValue("")
				
				// Simple command processing (you can expand this)
				if strings.Contains(strings.ToLower(value), "place") {
					m.board = append(m.board, fmt.Sprintf("  ● Piece placed: %s", value))
				}
			}
			return m, nil
		}
	}
	
	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m HiveModel) View() string {
	// Calculate dimensions for each panel
	leftPanelWidth := 25
	rightPanelWidth := m.width - leftPanelWidth - 6
	topHeight := m.height - 15
	bottomHeight := 10
	
	// Top section: Left panel (pieces) and Right panel (board)
	leftPanel := m.renderPiecesPanel(leftPanelWidth, topHeight)
	rightPanel := m.renderBoardPanel(rightPanelWidth, topHeight)
	
	topSection := lipgloss.JoinHorizontal(lipgloss.Top, leftPanel, rightPanel)
	
	// Bottom section: Command input and history (side by side)
	commandWidth := (m.width / 2) - 2
	historyWidth := (m.width / 2) - 2
	
	commandPanel := m.renderCommandPanel(commandWidth, bottomHeight)
	historyPanel := m.renderHistoryPanel(historyWidth, bottomHeight)
	
	bottomSection := lipgloss.JoinHorizontal(lipgloss.Top, commandPanel, historyPanel)
	
	// Combine top and bottom
	fullView := lipgloss.JoinVertical(lipgloss.Left, topSection, bottomSection)
	
	return fullView
}

func (m HiveModel) renderPiecesPanel(width, height int) string {
	var b strings.Builder
	
	b.WriteString(PanelTitleStyle.Render("Available Pieces"))
	b.WriteString("\n\n")
	
	for _, piece := range m.pieces {
		b.WriteString(PieceStyle.Render(piece))
		b.WriteString("\n")
	}
	
	content := b.String()
	return PanelStyle.Width(width).Height(height).Render(content)
}

func (m HiveModel) renderBoardPanel(width, height int) string {
	var b strings.Builder
	
	b.WriteString(PanelTitleStyle.Render(fmt.Sprintf("%s Game Board", m.game.Icon)))
	b.WriteString("\n\n")
	
	for _, line := range m.board {
		b.WriteString(BoardStyle.Render(line))
		b.WriteString("\n")
	}
	
	content := b.String()
	return PanelStyle.Width(width).Height(height).Render(content)
}

func (m HiveModel) renderCommandPanel(width, height int) string {
	var b strings.Builder
	
	b.WriteString(PanelTitleStyle.Render("Command Input"))
	b.WriteString("\n\n")
	
	b.WriteString(InputLabelStyle.Render("› "))
	b.WriteString(m.textInput.View())
	b.WriteString("\n\n")
	
	b.WriteString(HelpStyle.Render("enter: submit • esc: quit"))
	
	content := b.String()
	return PanelStyle.Width(width).Height(height).Render(content)
}

func (m HiveModel) renderHistoryPanel(width, height int) string {
	var b strings.Builder
	
	b.WriteString(PanelTitleStyle.Render("Command History"))
	b.WriteString("\n\n")
	
	if len(m.messages) == 0 {
		b.WriteString(DescriptionStyle.Render("  No commands yet..."))
	} else {
		// Show last 5 messages
		start := 0
		if len(m.messages) > 5 {
			start = len(m.messages) - 5
		}
		
		for i := start; i < len(m.messages); i++ {
			msg := MessageStyle.Render(fmt.Sprintf("[%d] %s", i+1, m.messages[i]))
			b.WriteString(msg)
			b.WriteString("\n")
		}
	}
	
	content := b.String()
	return PanelStyle.Width(width).Height(height).Render(content)
}

// Getters
func (m HiveModel) Messages() []string {
	return m.messages
}
