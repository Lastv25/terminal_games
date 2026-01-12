package models

import (
	"fmt"
	"strings"
)

// HexRenderer handles rendering the hexagonal board to ASCII art
type HexRenderer struct {
	board *HexBoard
}

// NewHexRenderer creates a new renderer for the given board
func NewHexRenderer(board *HexBoard) *HexRenderer {
	return &HexRenderer{board: board}
}

// Render creates an ASCII representation of the hexagonal board
func (r *HexRenderer) Render(width, height int) []string {
	if r.board.PieceCount() == 0 {
		return []string{
			"",
			"  [Empty Board]",
			"",
			"  Place pieces using:",
			"  place <piece> <q> <r>",
			"",
			"  Example: place WQ 0 0",
		}
	}
	
	lines := []string{}
	
	// Get board bounds with some padding
	minQ, maxQ, minR, maxR := r.board.GetBounds()
	minQ -= 1
	maxQ += 1
	minR -= 1
	maxR += 1
	
	// Header
	lines = append(lines, fmt.Sprintf("  Pieces on board: %d", r.board.PieceCount()))
	lines = append(lines, "")
	
	// Render each row
	for rowIdx := minR; rowIdx <= maxR; rowIdx++ {
		line := r.renderRow(rowIdx, minQ, maxQ)
		lines = append(lines, line)
	}
	
	lines = append(lines, "")
	lines = append(lines, "  Coordinates: (q, r)")
	
	return lines
}

// renderRow renders a single row of hexagons
func (r *HexRenderer) renderRow(row, minQ, maxQ int) string {
	var sb strings.Builder
	
	// Calculate offset for this row (for hex staggering)
	offset := ""
	if row%2 != 0 {
		offset = "  " // Offset odd rows
	}
	
	sb.WriteString(offset)
	sb.WriteString("  ") // Left padding
	
	for q := minQ; q <= maxQ; q++ {
		coord := HexCoordinate{Q: q, R: row}
		
		if piece, exists := r.board.GetTopPiece(coord); exists {
			// Draw piece with border
			sb.WriteString(fmt.Sprintf("[%s]", piece.ShortString()))
		} else {
			// Empty space - show coordinate
			sb.WriteString(fmt.Sprintf(" %2d,%-2d ", q, row))
		}
		
		sb.WriteString(" ") // Spacing between hexes
	}
	
	return sb.String()
}

// RenderCompact creates a more compact ASCII representation
func (r *HexRenderer) RenderCompact(width, height int) []string {
	if r.board.PieceCount() == 0 {
		return []string{
			"",
			"  [Empty Board]",
			"",
			"  place <piece> <q> <r>",
		}
	}
	
	lines := []string{}
	lines = append(lines, fmt.Sprintf("  %d pieces", r.board.PieceCount()))
	lines = append(lines, "")
	
	minQ, maxQ, minR, maxR := r.board.GetBounds()
	minQ -= 1
	maxQ += 1
	minR -= 1
	maxR += 1
	
	for rowIdx := minR; rowIdx <= maxR; rowIdx++ {
		var sb strings.Builder
		
		// Offset for hex grid
		if rowIdx%2 != 0 {
			sb.WriteString(" ")
		}
		sb.WriteString(" ")
		
		for q := minQ; q <= maxQ; q++ {
			coord := HexCoordinate{Q: q, R: rowIdx}
			
			if piece, exists := r.board.GetTopPiece(coord); exists {
				sb.WriteString(fmt.Sprintf("[%s]", piece.ShortString()))
			} else {
				sb.WriteString("  . ")
			}
			sb.WriteString(" ")
		}
		
		lines = append(lines, sb.String())
	}
	
	return lines
}

// RenderWithHighlight renders the board with a specific coordinate highlighted
func (r *HexRenderer) RenderWithHighlight(highlight HexCoordinate) []string {
	if r.board.PieceCount() == 0 {
		return r.Render(0, 0)
	}
	
	lines := []string{}
	lines = append(lines, fmt.Sprintf("  %d pieces (highlighting %d,%d)", 
		r.board.PieceCount(), highlight.Q, highlight.R))
	lines = append(lines, "")
	
	minQ, maxQ, minR, maxR := r.board.GetBounds()
	minQ -= 1
	maxQ += 1
	minR -= 1
	maxR += 1
	
	for rowIdx := minR; rowIdx <= maxR; rowIdx++ {
		var sb strings.Builder
		
		if rowIdx%2 != 0 {
			sb.WriteString("  ")
		}
		sb.WriteString("  ")
		
		for q := minQ; q <= maxQ; q++ {
			coord := HexCoordinate{Q: q, R: rowIdx}
			
			if coord.Equals(highlight) {
				sb.WriteString(">>")
			}
			
			if piece, exists := r.board.GetTopPiece(coord); exists {
				sb.WriteString(fmt.Sprintf("[%s]", piece.ShortString()))
			} else {
				sb.WriteString(fmt.Sprintf(" %2d,%-2d ", q, rowIdx))
			}
			
			if coord.Equals(highlight) {
				sb.WriteString("<<")
			} else {
				sb.WriteString("  ")
			}
		}
		
		lines = append(lines, sb.String())
	}
	
	return lines
}

