package main

import (
	"fmt"
	"os"

	"github.com/luique16/math_game/internal/game"
	"github.com/luique16/math_game/internal/view"
	"github.com/luique16/math_game/internal/menu"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	menu := menu.Create()

	p := tea.NewProgram(menu)

	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}

	err, game := game.CreateGame(6, 6, 50)

	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}

	app := view.Create(game)

	p = tea.NewProgram(app)

    if _, err := p.Run(); err != nil {
        fmt.Printf("Error: %s\n", err)
        os.Exit(1)
    }
}
