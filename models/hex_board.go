package models

import (
	"fmt"
)

// HexBoard manages the game board state
type HexBoard struct {
	Pieces map[HexCoordinate][]Piece // Stack of pieces at each position (beetles can stack)
}

// NewHexBoard creates a new empty board
func NewHexBoard() *HexBoard {
	return &HexBoard{
		Pieces: make(map[HexCoordinate][]Piece),
	}
}

// PlacePiece places a piece at the given coordinate
func (b *HexBoard) PlacePiece(coord HexCoordinate, piece Piece) {
	if _, exists := b.Pieces[coord]; !exists {
		b.Pieces[coord] = []Piece{}
	}
	b.Pieces[coord] = append(b.Pieces[coord], piece)
}

// RemovePiece removes the top piece from the given coordinate
func (b *HexBoard) RemovePiece(coord HexCoordinate) (Piece, bool) {
	stack, exists := b.Pieces[coord]
	if !exists || len(stack) == 0 {
		return Piece{}, false
	}
	
	piece := stack[len(stack)-1]
	b.Pieces[coord] = stack[:len(stack)-1]
	
	if len(b.Pieces[coord]) == 0 {
		delete(b.Pieces, coord)
	}
	
	return piece, true
}

// GetTopPiece returns the top piece at the given coordinate
func (b *HexBoard) GetTopPiece(coord HexCoordinate) (Piece, bool) {
	stack, exists := b.Pieces[coord]
	if !exists || len(stack) == 0 {
		return Piece{}, false
	}
	return stack[len(stack)-1], true
}

// IsOccupied checks if a coordinate has any pieces
func (b *HexBoard) IsOccupied(coord HexCoordinate) bool {
	stack, exists := b.Pieces[coord]
	return exists && len(stack) > 0
}

// GetBounds returns the min and max coordinates of pieces on the board
func (b *HexBoard) GetBounds() (minQ, maxQ, minR, maxR int) {
	if len(b.Pieces) == 0 {
		return 0, 0, 0, 0
	}
	
	first := true
	for coord := range b.Pieces {
		if first {
			minQ, maxQ = coord.Q, coord.Q
			minR, maxR = coord.R, coord.R
			first = false
		} else {
			if coord.Q < minQ {
				minQ = coord.Q
			}
			if coord.Q > maxQ {
				maxQ = coord.Q
			}
			if coord.R < minR {
				minR = coord.R
			}
			if coord.R > maxR {
				maxR = coord.R
			}
		}
	}
	return
}

// GetAllCoordinates returns all coordinates that have pieces
func (b *HexBoard) GetAllCoordinates() []HexCoordinate {
	coords := make([]HexCoordinate, 0, len(b.Pieces))
	for coord := range b.Pieces {
		coords = append(coords, coord)
	}
	return coords
}

// PieceCount returns the total number of pieces on the board
func (b *HexBoard) PieceCount() int {
	count := 0
	for _, stack := range b.Pieces {
		count += len(stack)
	}
	return count
}

// String returns a debug string representation of the board
func (b *HexBoard) String() string {
	if len(b.Pieces) == 0 {
		return "Empty board"
	}
	
	result := fmt.Sprintf("Board (%d pieces):\n", b.PieceCount())
	for coord, stack := range b.Pieces {
		result += fmt.Sprintf("  (%d,%d): ", coord.Q, coord.R)
		for i, piece := range stack {
			if i > 0 {
				result += ", "
			}
			result += piece.String()
		}
		result += "\n"
	}
	return result
}
