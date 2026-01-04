package menu

import (
	"github.com/charmbracelet/lipgloss"
)

func (v *Menu) View() string {
	var print string

	welcomeStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#ff55ff")).
		Bold(true).
		Margin(1).
		Align(lipgloss.Center)

	print += welcomeStyle.Render("========= Welcome to the Math Game! =========")


	gameExplanationStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#666666")).
		Margin(1).
		Align(lipgloss.Center)

	print += gameExplanationStyle.Render("Each column and row has a specific sum,\n find and select the correct numbers that\n make up the sum and delete the\n incorrect ones.")

	print += "\n"

	if !v.GameModeSelected {
		print += "Select a game mode:\n\n"

		for i := range(v.Choices) {
			notSelectedStyle := lipgloss.NewStyle().
				Foreground(lipgloss.Color("#ffffff")).
				Margin(0).
				Align(lipgloss.Left)

			selectedStyle := lipgloss.NewStyle().
				Foreground(lipgloss.Color("#ff55ff")).
				Bold(true).
				Margin(0).
				Align(lipgloss.Left)

			if i == v.Cursor {
				print += selectedStyle.Render("→ " + v.Choices[i])
			} else {
				print += notSelectedStyle.Render(v.Choices[i])
			}

			print += "\n"
		}
	} else {
		print += "Gamemode: " + v.Choices[v.GameMode] + "\n"
		print += "\n"

		print += "Are you ready to play?\n"

		yesStyle := lipgloss.NewStyle().
			Foreground(lipgloss.Color("#ffffff")).
			Background(lipgloss.Color("#ff55ff")).
			Bold(true).
			Margin(1).
			Align(lipgloss.Center)

		print += yesStyle.Render("Yes")

		print += "\n"
	}

	return print
}

