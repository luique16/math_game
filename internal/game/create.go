package game

import (
	"math/rand"
	"math"
	"errors"

	"github.com/luique16/math_game/internal/models"
)

func CreateGame(rows, cols, relation int) (error, *models.Game) {
	if rows < 1 || cols < 1 || relation < 0 || relation > 100 {
		return errors.New("Invalid configuration"), &models.Game{}
	}

	var legitCells = int(math.Floor(float64(rows)*float64(cols)*(float64(relation)/100)))

	if legitCells < cols {
		return errors.New("Invalid configuration"), &models.Game{}
	}


	var rowsBase int = int(math.Floor(float64(legitCells / rows)))

	var rowsQuantity []int = make([]int, rows)

	for i := range(rows) {
		rowsQuantity[i] = rowsBase
	}

	for range(legitCells - rowsBase*rows) {
		i := rand.Intn(rows)

		rowsQuantity[i]++
	}

	game := &models.Game{
		Rows:     rows,
		Cols:     cols,
		Board:    make([][]models.Cell, rows),
		Lives:    3,
		RowsSums: make([]int, rows),
		ColsSums: make([]int, cols),
	}

	for i := range(rows) {
		game.Board[i] = make([]models.Cell, cols)

		for j := range(cols) {
			v := rand.Intn(10)

			for v == 0 {
				v = rand.Intn(10)
			}

			game.Board[i][j] = models.Cell{
				Value: v,
				Legit: false,
				State: models.Active,
			}
		}
	}

	simulationMap := make([][]bool, rows)
	for i := range(rows) {
		simulationMap[i] = make([]bool, cols)
	}

	for {
		simCols := make([]int, cols)

		for i := range(rows) {
			for j := range(cols) {
				simulationMap[i][j] = false
			}
		}

		for i := range(rows) {
			for range(rowsQuantity[i]) {
				j := rand.Intn(cols)

				for simulationMap[i][j] == true {
					j = rand.Intn(cols)
				}

				simulationMap[i][j] = true

				simCols[j]++
			}
		}

		flag := false
		for i := range(cols) {
			if simCols[i] == 0 {
				flag = true
			}
		}

		if flag == true {
			continue
		}

		break
	}

	for i := range(rows) {
		for j := range(cols) {
			if simulationMap[i][j] == true {
				game.Board[i][j].Legit = true
				game.RowsSums[i] += game.Board[i][j].Value
				game.ColsSums[j] += game.Board[i][j].Value
			}
		}
	}

	return nil, game
}
