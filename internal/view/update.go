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
					v.CursorY = max(0, v.CursorY - 1)
				case "down", "j":
					v.CursorY = min(v.Game.Rows - 1, v.CursorY + 1)
				case "left", "h":
					v.CursorX = max(0, v.CursorX - 1)
				case "right", "l":
					v.CursorX = min(v.Game.Cols - 1, v.CursorX + 1)
				case "backspace":
					// v.Game.Erase(v.CursorX, v.CursorY)
				case "enter":
					// v.Game.Select(v.CursorX, v.CursorY)
			}
	}

	return v, nil
}
