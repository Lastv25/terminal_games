package models

import (
	"fmt"
	"strconv"
	"strings"
)

// CommandType represents different command types
type CommandType int

const (
	PlaceCommand CommandType = iota
	MoveCommand
	InvalidCommand
)

// Command represents a parsed user command
type Command struct {
	Type      CommandType
	Piece     string
	FromCoord HexCoordinate
	ToCoord   HexCoordinate
	Error     string
}

// ParseCommand parses a user input string into a Command
func ParseCommand(input string) Command {
	input = strings.TrimSpace(input)
	parts := strings.Fields(input)
	
	if len(parts) == 0 {
		return Command{
			Type:  InvalidCommand,
			Error: "Empty command",
		}
	}
	
	cmdType := strings.ToLower(parts[0])
	
	switch cmdType {
	case "place":
		return parsePlaceCommand(parts)
	case "move":
		return parseMoveCommand(parts)
	default:
		return Command{
			Type:  InvalidCommand,
			Error: fmt.Sprintf("Unknown command: %s", cmdType),
		}
	}
}

// parsePlaceCommand parses a place command
// Format: place <piece> <q> <r>
// Example: place WQ 0 0
func parsePlaceCommand(parts []string) Command {
	if len(parts) < 4 {
		return Command{
			Type:  InvalidCommand,
			Error: "Place command format: place <piece> <q> <r>",
		}
	}
	
	piece := strings.ToUpper(parts[1])
	
	q, err := strconv.Atoi(parts[2])
	if err != nil {
		return Command{
			Type:  InvalidCommand,
			Error: fmt.Sprintf("Invalid q coordinate: %s", parts[2]),
		}
	}
	
	r, err := strconv.Atoi(parts[3])
	if err != nil {
		return Command{
			Type:  InvalidCommand,
			Error: fmt.Sprintf("Invalid r coordinate: %s", parts[3]),
		}
	}
	
	// Validate piece format
	if !isValidPiece(piece) {
		return Command{
			Type:  InvalidCommand,
			Error: fmt.Sprintf("Invalid piece: %s. Use format like WQ, BA1, WS2", piece),
		}
	}
	
	return Command{
		Type:    PlaceCommand,
		Piece:   piece,
		ToCoord: HexCoordinate{Q: q, R: r},
	}
}

// parseMoveCommand parses a move command
// Format: move <piece> <from_q> <from_r> <to_q> <to_r>
// Example: move WQ 0 0 1 0
func parseMoveCommand(parts []string) Command {
	if len(parts) < 6 {
		return Command{
			Type:  InvalidCommand,
			Error: "Move command format: move <piece> <from_q> <from_r> <to_q> <to_r>",
		}
	}
	
	piece := strings.ToUpper(parts[1])
	
	fromQ, err := strconv.Atoi(parts[2])
	if err != nil {
		return Command{
			Type:  InvalidCommand,
			Error: fmt.Sprintf("Invalid from_q coordinate: %s", parts[2]),
		}
	}
	
	fromR, err := strconv.Atoi(parts[3])
	if err != nil {
		return Command{
			Type:  InvalidCommand,
			Error: fmt.Sprintf("Invalid from_r coordinate: %s", parts[3]),
		}
	}
	
	toQ, err := strconv.Atoi(parts[4])
	if err != nil {
		return Command{
			Type:  InvalidCommand,
			Error: fmt.Sprintf("Invalid to_q coordinate: %s", parts[4]),
		}
	}
	
	toR, err := strconv.Atoi(parts[5])
	if err != nil {
		return Command{
			Type:  InvalidCommand,
			Error: fmt.Sprintf("Invalid to_r coordinate: %s", parts[5]),
		}
	}
	
	return Command{
		Type:      MoveCommand,
		Piece:     piece,
		FromCoord: HexCoordinate{Q: fromQ, R: fromR},
		ToCoord:   HexCoordinate{Q: toQ, R: toR},
	}
}

// isValidPiece checks if a piece string is valid
// Valid formats: WQ, BA1, WS2, etc.
func isValidPiece(piece string) bool {
	if len(piece) < 2 || len(piece) > 3 {
		return false
	}
	
	// First character must be W or B (color)
	color := piece[0]
	if color != 'W' && color != 'B' {
		return false
	}
	
	// Second character must be a valid piece type
	pieceType := piece[1]
	validTypes := map[byte]bool{
		'Q': true, // Queen
		'A': true, // Ant
		'G': true, // Grasshopper
		'S': true, // Spider
		'B': true, // Beetle
	}
	
	if !validTypes[pieceType] {
		return false
	}
	
	// If there's a third character, it must be a digit
	if len(piece) == 3 {
		number := piece[2]
		if number < '1' || number > '9' {
			return false
		}
	}
	
	return true
}

// ParsePieceString converts a piece string to a Piece struct
// Example: "WQ" -> Piece{Type: QueenBee, Color: White, Number: 0}
func ParsePieceString(pieceStr string) (Piece, error) {
	if !isValidPiece(pieceStr) {
		return Piece{}, fmt.Errorf("invalid piece format: %s", pieceStr)
	}
	
	color := White
	if pieceStr[0] == 'B' {
		color = Black
	}
	
	var pieceType PieceType
	switch pieceStr[1] {
	case 'Q':
		pieceType = QueenBee
	case 'A':
		pieceType = Ant
	case 'G':
		pieceType = Grasshopper
	case 'S':
		pieceType = Spider
	case 'B':
		pieceType = Beetle
	}
	
	number := 0
	if len(pieceStr) == 3 {
		number = int(pieceStr[2] - '0')
	}
	
	return Piece{
		Type:   pieceType,
		Color:  color,
		Number: number,
	}, nil
}

