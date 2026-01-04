package app

import (
	tea "github.com/charmbracelet/bubbletea"

	"github.com/luique16/math_game/internal/game"
	"github.com/luique16/math_game/internal/game_view"
	"github.com/luique16/math_game/internal/menu"
)

func (v *App) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var m tea.Model

	switch v.Page {
		case 0:
			m, cmd = v.Menu.Update(msg)

			v.Menu = m.(*menu.Menu)

			if v.Menu.Exit {
				return v, tea.Quit
			}

			var err error

			if v.Menu.Ready {
				switch v.Menu.GameMode {
					case 0:
						err, v.Game = game.CreateGame(4, 4, 50)
					case 1:
						err, v.Game = game.CreateGame(5, 5, 50)
					case 2:
						err, v.Game = game.CreateGame(6, 6, 50)
					case 3:
						err, v.Game = game.CreateGame(7, 7, 50)
				}

				if err != nil {
					return v, tea.Quit
				}

				v.Page = 1
				v.GameView = game_view.Create(v.Game)
			}
		case 1:
			m, cmd = v.GameView.Update(msg)

			v.GameView = m.(*game_view.GameView)
	}

	return v, cmd
}
