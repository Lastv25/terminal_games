package main

import (
	"Coding/games/models"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Prompting user choices and validating against enum
func WhichGame(label string) string {
	var str string
	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Fprint(os.Stderr, label+" ")
		str, _ = r.ReadString('\n')
		if str != "" {
			break
		}
	}
	return strings.TrimSpace(str)
}

// Main Loop
func main() {
	fmt.Println("To start a game select between the available ones:")
	idx := 1
	for i := models.StartGame + 1; i < models.EndGame; i++ {
		fmt.Println("For", models.Games(i), "Enter", idx)
		idx += 1
	}
	game_idx_str := WhichGame("Enter the game idx:")
	game_idx, err := strconv.ParseInt(game_idx_str, 10, 64)
	if err != nil {
		// handle error
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Input was:", game_idx)
	fmt.Println("Game chosen was:", models.Games(game_idx))
}
