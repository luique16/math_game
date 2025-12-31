package view

import (
	"fmt"
)

func (v *View) View() string {
	var print string

	print += "   "

	for i := range(v.Game.Cols) {
		print += fmt.Sprintf("%d ", v.Game.ColsSums[i])
	}
	print += "\n\n"

	for i := range(v.Game.Rows) {
		print += fmt.Sprintf("%d  ", v.Game.RowsSums[i])
		for j := range(v.Game.Cols) {
			if v.CursorY == i && v.CursorX == j {
				print += fmt.Sprintf("\033[34m%d\033[0m  ", v.Game.Board[i][j].Value) // Blue
			} else if v.Game.Board[i][j].Legit {
				print += fmt.Sprintf("\033[31m%d\033[0m  ", v.Game.Board[i][j].Value) // Red
			} else {
				print += fmt.Sprintf("%d  ", v.Game.Board[i][j].Value)
			}
		}
		print += "\n"
	}

	return print
}
