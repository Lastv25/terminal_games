package models

// HexCoordinate represents a position on the hexagonal board using axial coordinates
// Axial coordinates (q, r) are a common way to represent hex grids
// q represents the column, r represents the row
type HexCoordinate struct {
	Q int // Column coordinate
	R int // Row coordinate
}

// NewHexCoordinate creates a new hexagonal coordinate
func NewHexCoordinate(q, r int) HexCoordinate {
	return HexCoordinate{Q: q, R: r}
}

// Equals checks if two coordinates are the same
func (h HexCoordinate) Equals(other HexCoordinate) bool {
	return h.Q == other.Q && h.R == other.R
}

// Neighbors returns all 6 neighboring coordinates
func (h HexCoordinate) Neighbors() []HexCoordinate {
	return []HexCoordinate{
		{Q: h.Q + 1, R: h.R},     // East
		{Q: h.Q + 1, R: h.R - 1}, // Northeast
		{Q: h.Q, R: h.R - 1},     // Northwest
		{Q: h.Q - 1, R: h.R},     // West
		{Q: h.Q - 1, R: h.R + 1}, // Southwest
		{Q: h.Q, R: h.R + 1},     // Southeast
	}
}

// Distance calculates the hex distance between two coordinates
func (h HexCoordinate) Distance(other HexCoordinate) int {
	return (abs(h.Q-other.Q) + abs(h.Q+h.R-other.Q-other.R) + abs(h.R-other.R)) / 2
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
