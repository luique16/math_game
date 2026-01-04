package menu

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Menu struct {
	Exit bool
	Cursor int
	Choices []string
	GameModeSelected bool
	GameMode int
}

func (v Menu) Init() tea.Cmd {
    return nil
}

func Create() *Menu {
	return &Menu{
		Exit: false,
		Cursor: 1,
		Choices: []string{"Easy", "Medium", "Hard", "Extreme"},
		GameModeSelected: false,
		GameMode: 0,
	}
}

