package models

// PieceType represents the type of Hive piece
type PieceType string

const (
	QueenBee     PieceType = "Q"
	Ant          PieceType = "A"
	Grasshopper  PieceType = "G"
	Spider       PieceType = "S"
	Beetle       PieceType = "B"
)

// PieceColor represents which player owns the piece
type PieceColor string

const (
	White PieceColor = "W"
	Black PieceColor = "B"
)

// Piece represents a single game piece
type Piece struct {
	Type   PieceType
	Color  PieceColor
	Number int // For distinguishing multiple pieces of the same type (e.g., Ant #1, Ant #2)
}

// NewPiece creates a new piece
func NewPiece(pieceType PieceType, color PieceColor, number int) Piece {
	return Piece{
		Type:   pieceType,
		Color:  color,
		Number: number,
	}
}

// String returns a string representation of the piece (e.g., "WQ", "BA1", "WS2")
func (p Piece) String() string {
	if p.Type == QueenBee {
		return string(p.Color) + string(p.Type)
	}
	return string(p.Color) + string(p.Type) + string(rune('0'+p.Number))
}

// ShortString returns a shorter representation for display
func (p Piece) ShortString() string {
	return p.String()
}

// PieceInfo holds metadata about each piece type
type PieceInfo struct {
	Name     string
	Symbol   PieceType
	Quantity int
}

// GetAllPieceTypes returns information about all piece types
func GetAllPieceTypes() []PieceInfo {
	return []PieceInfo{
		{Name: "Queen Bee", Symbol: QueenBee, Quantity: 1},
		{Name: "Ant", Symbol: Ant, Quantity: 3},
		{Name: "Grasshopper", Symbol: Grasshopper, Quantity: 3},
		{Name: "Spider", Symbol: Spider, Quantity: 2},
		{Name: "Beetle", Symbol: Beetle, Quantity: 2},
	}
}
