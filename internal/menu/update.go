package menu

import (
	tea "github.com/charmbracelet/bubbletea"
)

func (v *Menu) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if !v.GameModeSelected {
		switch msg := msg.(type) {
			case tea.KeyMsg:
				switch msg.String() {
					case "ctrl+c", "q":
						v.Exit = true
						return v, tea.Quit
					case "up", "k", "w":
						v.Cursor = max(0, v.Cursor - 1)
					case "down", "j", "s":
						v.Cursor = min(len(v.Choices) - 1, v.Cursor + 1)
					case "enter":
						switch v.Choices[v.Cursor] {
							case "Easy":
								v.GameMode = 0
							case "Medium":
								v.GameMode = 1
							case "Hard":
								v.GameMode = 2
							case "Extreme":
								v.GameMode = 3
						}
						v.GameModeSelected = true
				}
		}
	} else {
		switch msg := msg.(type) {
			case tea.KeyMsg:
				switch msg.String() {
					case "ctrl+c", "q":
						v.Exit = true
						return v, tea.Quit
					case "enter":
						return v, tea.Quit
				}
		}
	}

	return v, nil
}
