package main

import (
	"fmt"
	"os"

	"github.com/luique16/math_game/internal/app"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	app := app.Create()

	p := tea.NewProgram(app)

	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}
}
