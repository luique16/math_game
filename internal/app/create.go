package app

import (
	tea "github.com/charmbracelet/bubbletea"

	"github.com/luique16/math_game/internal/models"
	"github.com/luique16/math_game/internal/menu"
	"github.com/luique16/math_game/internal/game_view"
)

type App struct {
	Game *models.Game
	Page int
	Menu *menu.Menu
	GameView *game_view.GameView
}

func (v App) Init() tea.Cmd {
    return nil
}

func Create() *App {
	return &App{
		Game: nil,
		Page: 0,
		Menu: menu.Create(),
		GameView: nil,
	}
}
