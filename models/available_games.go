package models


// Game represents a playable game
type Game struct {
	Name        string
	Description string
}

func (g Game) String() string {
	return g.Name
}

// Predefined games
var (
	Hive = Game{
		Name:        "Hive",
		Description: "Strategic insect placement game",
	}
	Hortis = Game{
		Name:        "Hortis",
		Description: "Garden building game",
	}
	StarRealms = Game{
		Name:        "Star Realms",
		Description: "Space combat deck builder",
	}
)
