package main

import (
	"fmt"
	"os"

	"github.com/luique16/math_game/internal/game"
	"github.com/luique16/math_game/internal/view"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	err, game := game.CreateGame(10, 10, 50)

	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}

	app := view.Create(game)

	p := tea.NewProgram(app)

    if _, err := p.Run(); err != nil {
        fmt.Printf("Error: %s\n", err)
        os.Exit(1)
    }
}
