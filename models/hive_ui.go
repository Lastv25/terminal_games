package models

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/lipgloss"
)

// Style definitions for Hive
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
	
	ErrorStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FF0000")).
			PaddingLeft(2)

)

// HiveModel handles the 4-panel Hive game interface
type HiveModel struct {
	game         Game
	textInput    textinput.Model
	messages     []string
	board        *HexBoard
	renderer     *HexRenderer
	width        int
	height       int
	lastError    string
}

// NewHiveModel creates a new Hive game model with 4-panel layout
func NewHiveModel(game Game) HiveModel {
	ti := textinput.New()
	ti.Placeholder = "place WQ 0 0  |  move WQ 0 0 1 0"
	ti.Focus()
	ti.CharLimit = 100
	ti.Width = 40
	ti.PromptStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF06B7"))
	ti.TextStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFFFFF"))
	
	board := NewHexBoard()
	renderer := NewHexRenderer(board)
	
	return HiveModel{
		game:      game,
		textInput: ti,
		messages:  []string{},
		board:     board,
		renderer:  renderer,
		lastError: "",
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
				m = m.handleCommand(value)
				m.textInput.SetValue("")
			}
			return m, nil
		}
	}
	
	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m HiveModel) handleCommand(input string) HiveModel {
	// Add to message history
	m.messages = append(m.messages, input)
	m.lastError = ""
	
	// Parse command
	command := ParseCommand(input)
	
	if command.Type == InvalidCommand {
		m.lastError = command.Error
		return m
	}
	
	switch command.Type {
	case PlaceCommand:
		m = m.handlePlaceCommand(command)
	case MoveCommand:
		m = m.handleMoveCommand(command)
	}
	
	return m
}

func (m HiveModel) handlePlaceCommand(cmd Command) HiveModel {
	// Parse the piece string
	piece, err := ParsePieceString(cmd.Piece)
	if err != nil {
		m.lastError = err.Error()
		return m
	}
	
	// Place the piece on the board
	m.board.PlacePiece(cmd.ToCoord, piece)
	
	return m
}

func (m HiveModel) handleMoveCommand(cmd Command) HiveModel {
	// Remove piece from source
	piece, ok := m.board.RemovePiece(cmd.FromCoord)
	if !ok {
		m.lastError = fmt.Sprintf("No piece at (%d, %d)", cmd.FromCoord.Q, cmd.FromCoord.R)
		return m
	}
	
	// Place piece at destination
	m.board.PlacePiece(cmd.ToCoord, piece)
	
	return m
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
	
	b.WriteString(PanelTitleStyle.Render("Piece Reference"))
	b.WriteString("\n\n")
	
	b.WriteString(PieceStyle.Render("White pieces:"))
	b.WriteString("\n")
	b.WriteString(PieceStyle.Render("  WQ - Queen"))
	b.WriteString("\n")
	b.WriteString(PieceStyle.Render("  WA1,2,3 - Ants"))
	b.WriteString("\n")
	b.WriteString(PieceStyle.Render("  WG1,2,3 - Hoppers"))
	b.WriteString("\n")
	b.WriteString(PieceStyle.Render("  WS1,2 - Spiders"))
	b.WriteString("\n")
	b.WriteString(PieceStyle.Render("  WB1,2 - Beetles"))
	b.WriteString("\n\n")
	
	b.WriteString(PieceStyle.Render("Black pieces:"))
	b.WriteString("\n")
	b.WriteString(PieceStyle.Render("  BQ, BA1-3, etc."))
	b.WriteString("\n\n")
	
	b.WriteString(DescriptionStyle.Render("Coordinates:"))
	b.WriteString("\n")
	b.WriteString(DescriptionStyle.Render("  (q, r) format"))
	b.WriteString("\n")
	b.WriteString(DescriptionStyle.Render("  Axial hex grid"))
	
	content := b.String()
	return PanelStyle.Width(width).Height(height).Render(content)
}

func (m HiveModel) renderBoardPanel(width, height int) string {
	var b strings.Builder
	
	b.WriteString(PanelTitleStyle.Render("Game Board"))
	b.WriteString("\n")
	
	// Render the hexagonal board
	boardLines := m.renderer.Render(width, height)
	for _, line := range boardLines {
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
	
	b.WriteString(InputLabelStyle.Render("> "))
	b.WriteString(m.textInput.View())
	b.WriteString("\n")
	
	// Show error if present
	if m.lastError != "" {
		b.WriteString("\n")
		b.WriteString(ErrorStyle.Render("Error: " + m.lastError))
	}
	
	b.WriteString("\n")
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
		b.WriteString("\n\n")
		b.WriteString(DescriptionStyle.Render("  Try:"))
		b.WriteString("\n")
		b.WriteString(DescriptionStyle.Render("  place WQ 0 0"))
		b.WriteString("\n")
		b.WriteString(DescriptionStyle.Render("  place BA1 1 0"))
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

