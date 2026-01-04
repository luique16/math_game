package game_view

import (
	"github.com/charmbracelet/lipgloss"
)

func (v *GameView) View() string {
	var print string

	print += lipgloss.JoinHorizontal(
		lipgloss.Top, 
		Hearts(v.Game.Lives),
		)

	print += "\n"

	print += lipgloss.JoinHorizontal(
		lipgloss.Top, 
		lipgloss.JoinVertical(
			lipgloss.Left,
			ColsSum(v.Game.ColsSums),
			lipgloss.JoinHorizontal(
				lipgloss.Top,
				RowsSum(v.Game.RowsSums),
				Board(v.Game, v.CursorX, v.CursorY),
			),
		),
	)

	print += "\n"

	print += ControlsBar()

	print += "\n\n"

	if v.Game.Lives == 0 {
		print += "Game over!\n"
	}

	if VerifyEndGame(v.Game) {
		print += "You win!\n"
	}

	return print
}
