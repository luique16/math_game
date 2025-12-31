package view

import (
	"github.com/luique16/math_game/internal/models"

	tea "github.com/charmbracelet/bubbletea"
)

type View struct {
	Game    *models.Game
	CursorX int
	CursorY int
}

func (v View) Init() tea.Cmd {
    return nil
}

func Create(game *models.Game) *View {
	return &View{
		Game: game,
		CursorX: 0,
		CursorY: game.Rows - 1,
	}
}
