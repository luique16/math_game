package view

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

func (v *View) View() string {
	var print string

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

	for i := range(v.Game.Cols) {
		colsSum = append(colsSum, colsElementStyle.Render(fmt.Sprintf("%d", v.Game.ColsSums[i])))
	}

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

	for i := range(v.Game.Rows) {
		rowsSum = append(rowsSum, rowElementStyle.Render(fmt.Sprintf("%d", v.Game.RowsSums[i])))
	}

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

	for i := range(v.Game.Rows) {
		var row []string
		for j := range(v.Game.Cols) {
			if v.CursorY == i && v.CursorX == j {
				row = append(row, selectedElementStyle.Render(fmt.Sprintf("%d", v.Game.Board[i][j].Value)))
			} else if v.Game.Board[i][j].Legit {
				row = append(row, legitElementStyle.Render(fmt.Sprintf("%d", v.Game.Board[i][j].Value)))
			} else {
				row = append(row, boardElementStyle.Render(fmt.Sprintf("%d", v.Game.Board[i][j].Value)))
			}
		}
		board = append(board, row)
	}

	var boardRows []string

	for i := range(v.Game.Rows) {
		boardRows = append(boardRows, lipgloss.JoinHorizontal(lipgloss.Left, board[i]...))
	}

	print += lipgloss.JoinHorizontal(
		lipgloss.Top, 
		lipgloss.JoinVertical(
			lipgloss.Left,
			colsSumStyle.Render(
				lipgloss.JoinHorizontal(lipgloss.Top, colsSum...),
				),
			lipgloss.JoinHorizontal(
				lipgloss.Top,
				rowsSumStyle.Render(
					lipgloss.JoinVertical(lipgloss.Left, rowsSum...),
					),
				boardStyle.Render(
					lipgloss.JoinVertical(lipgloss.Right, boardRows...),
					),
				),
		),
	)

	print += "\n\n"

	return print
}
