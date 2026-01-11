package main

import (
	"fmt"
	"strings"

	"Coding/games/models"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	// Phase 1: Game Selection
	menuModel := models.NewMenuModel()
	p := tea.NewProgram(menuModel)

	finalModel, err := p.Run()
	if err != nil {
		fmt.Printf("Error running menu: %v\n", err)
		return
	}

	// Check if user made a selection
	menu := finalModel.(models.MenuModel)
	if menu.Selected() < 0 {
		fmt.Println("\nNo game selected. Goodbye!")
		return
	}

	selectedGame := menu.SelectedGame()

	// Transition screen
	fmt.Print("\033[H\033[2J") // Clear screen
	fmt.Println()
	fmt.Println(strings.Repeat("═", 60))
	fmt.Printf("Starting %s...\n", selectedGame.Name)
	fmt.Printf("%s\n", selectedGame.Description)
	fmt.Println(strings.Repeat("═", 60))
	fmt.Println()

	// Phase 2: Game Interface - Use special HiveModel for Hive game
	if selectedGame.Name == "Hive" {
		// Use 4-panel Hive interface
		hiveModel := models.NewHiveModel(selectedGame)
		p = tea.NewProgram(hiveModel, tea.WithAltScreen())
		if _, err := p.Run(); err != nil {
			fmt.Printf("Error running Hive interface: %v\n", err)
			return
		}
	} else {
		// Use standard input interface for other games
		// this is a placeholder for now
		inputModel := models.NewInputModel(selectedGame)
		p = tea.NewProgram(inputModel)
		if _, err := p.Run(); err != nil {
			fmt.Printf("Error running game interface: %v\n", err)
			return
		}
	}

	// Exit message
	fmt.Println("\nThanks for playing! See you next time!\n")
}

