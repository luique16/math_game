package game_view

import (
	"github.com/luique16/math_game/internal/models"

	tea "github.com/charmbracelet/bubbletea"
)

type GameView struct {
	Game    *models.Game
	CursorX int
	CursorY int
}

func (v GameView) Init() tea.Cmd {
    return nil
}

func Create(game *models.Game) *GameView {
	return &GameView{
		Game: game,
		CursorX: 0,
		CursorY: 0,
	}
}
