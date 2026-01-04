package view

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"

	"github.com/luique16/math_game/internal/models"
)

func Hearts(lives int) string {
	heartsStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#ff6961")).
		Bold(true).
		Margin(1).
		Align(lipgloss.Center)

	aliveHeartElementStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#ff0000")).
		Bold(true).
		Width(2).
		Padding(0).
		Align(lipgloss.Center)

	deadHeartElementStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#666666")).
		Bold(true).
		Width(2).
		Padding(0).
		Align(lipgloss.Center)

	var hearts []string

	for range(lives) {
		hearts = append(hearts, aliveHeartElementStyle.Render("❤ "))
	}

	for range(3 - lives) {
		hearts = append(hearts, deadHeartElementStyle.Render("❤ "))
	}

	return heartsStyle.Render(lipgloss.JoinHorizontal(lipgloss.Top, hearts...))
}

func ColsSum(colsSums []int) string {
	colsSumStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#ffffff")).
		Bold(true).
		Border(lipgloss.RoundedBorder(), true).
		MarginLeft(7).
		PaddingRight(2).
		Align(lipgloss.Center)

	colsElementStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#ffffff")).
		PaddingLeft(2).
		Width(4).
		Align(lipgloss.Center)

	var colsSum []string

	for i := range(colsSums) {
		colsSum = append(colsSum, colsElementStyle.Render(fmt.Sprintf("%d", colsSums[i])))
	}

	return colsSumStyle.Render(lipgloss.JoinHorizontal(lipgloss.Top, colsSum...))
}

func RowsSum(rowsSums []int) string {
	rowsSumStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#ffffff")).
		Bold(true).
		Padding(0, 1, 1, 1).
		Border(lipgloss.RoundedBorder(), true).
		Align(lipgloss.Left)

	rowElementStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#ffffff")).
		PaddingTop(1).
		Align(lipgloss.Center)

	var rowsSum []string

	for i := range(len(rowsSums)) {
		rowsSum = append(rowsSum, rowElementStyle.Render(fmt.Sprintf("%d", rowsSums[i])))
	}

	return rowsSumStyle.Render(lipgloss.JoinVertical(lipgloss.Left, rowsSum...))
}

func Board(game *models.Game, cursorX, cursorY int) string {
	boardStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#ffffff")).
		Border(lipgloss.RoundedBorder(), true).
		MarginLeft(1).
		PaddingLeft(1).
		PaddingRight(1).
		PaddingTop(1).
		Align(lipgloss.Center)

	boardElementStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#ffffff")).
		Width(4).
		Padding(0).
		MarginBottom(1).
		Align(lipgloss.Center)

	selectedElementStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#ff6961")).
		Bold(true).
		Width(4).
		Padding(0).
		MarginBottom(1).
		Align(lipgloss.Center)

	legitElementStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#666666")).
		Width(4).
		Padding(0).
		MarginBottom(1).
		Align(lipgloss.Center)


	var board [][]string

	for i := range(game.Rows) {
		var row []string
		for j := range(game.Cols) {
			if cursorY == i && cursorX == j {
				if game.Board[i][j].State != models.Erased {
					row = append(row, selectedElementStyle.Render(fmt.Sprintf("%d", game.Board[i][j].Value)))
				} else {
					row = append(row, selectedElementStyle.Render("⚪ "))
				}
			} else if game.Board[i][j].Legit && game.Board[i][j].State == models.Selected {
				row = append(row, legitElementStyle.Render(fmt.Sprintf("%d", game.Board[i][j].Value)))
			} else if game.Board[i][j].State == models.Erased {
				row = append(row, boardElementStyle.Render(" "))
			} else {
				row = append(row, boardElementStyle.Render(fmt.Sprintf("%d", game.Board[i][j].Value)))
			}
		}
		board = append(board, row)
	}

	var boardRows []string

	for i := range(game.Rows) {
		boardRows = append(boardRows, lipgloss.JoinHorizontal(lipgloss.Left, board[i]...))
	}

	return boardStyle.Render(lipgloss.JoinVertical(lipgloss.Right, boardRows...))
}

func ControlsBar() string {
	style := lipgloss.NewStyle().
		Padding(0, 1).
		Foreground(lipgloss.Color("#666666")).
		MarginTop(1)

	return style.Render(
		"↑↓←→ Move  •  Enter Select  •  Backspace Erase  •  Q Quit",
	)
}
