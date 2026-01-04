package game_view

import (
	tea "github.com/charmbracelet/bubbletea"

	"github.com/luique16/math_game/internal/models"
)

func VerifyEndGame(game *models.Game) bool {
	for i := range(game.Rows) {
		for j := range(game.Cols) {
			if game.Board[i][j].State == models.Active {
				return false
			}
		}
	}

	return true
}

func (v *GameView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if v.Game.Lives == 0 || VerifyEndGame(v.Game) {
		return v, tea.Quit
	}

	switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
				case "ctrl+c", "q":
					return v, tea.Quit
				case "up", "k", "w":
					v.CursorY = max(0, v.CursorY - 1)
				case "down", "j", "s":
					v.CursorY = min(v.Game.Rows - 1, v.CursorY + 1)
				case "left", "h", "a":
					v.CursorX = max(0, v.CursorX - 1)
				case "right", "l", "d":
					v.CursorX = min(v.Game.Cols - 1, v.CursorX + 1)
				case "backspace":
					v.Game.Erase(v.CursorX, v.CursorY)
				case "enter":
					v.Game.Select(v.CursorX, v.CursorY)
			}
	}


	return v, nil
}
