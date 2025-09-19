package models

// Defining the games
type Games int64

// StartGame and EndGame are only usefull to iterate over the enum values using a simple for loop
const (
	StartGame Games = iota
	Hortis
	Hive
	Star_Realms
	EndGame
)

func (s Games) String() string {
	switch s {
	case Hortis:
		return "Hortis"
	case Hive:
		return "Hive"
	case Star_Realms:
		return "Star Realms"
	}
	return "Unknown"
}

func IsValidGame(value int64) bool {
	g := Games(value)
	return g > StartGame && g < EndGame
}
