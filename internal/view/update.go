package view

import (
	tea "github.com/charmbracelet/bubbletea"
)

func (v *View) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
				case "ctrl+c", "q":
					return v, tea.Quit
				case "up", "k":
					if v.CursorY < v.Game.Rows - 1 {
						v.CursorY++
					}
				case "down", "j":
					if v.CursorY > 0 {
						v.CursorY--
					}
				case "left", "h":
					if v.CursorX > 0 {
						v.CursorX--
					}
				case "right", "l":
					if v.CursorX < v.Game.Cols - 1 {
						v.CursorX++
					}
				case "backspace":
					// v.Game.Erase(v.CursorX, v.CursorY)
				case "enter":
					// v.Game.Select(v.CursorX, v.CursorY)
			}
	}

	return v, nil
}
