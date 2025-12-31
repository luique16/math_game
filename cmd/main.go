package main

import (
	"fmt"

	"github.com/luique16/math_game/internal/game"
)

func main() {
	err, game := game.CreateGame(10, 10, 50)

	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}

	fmt.Printf("   ")
	for i := range(game.Cols) {
		fmt.Printf("%d ", game.ColsSums[i])
	}
	fmt.Printf("\n\n")

	for i := range(game.Rows) {
		fmt.Printf("%d  ", game.RowsSums[i])
		for j := range(game.Cols) {
			if game.Board[i][j].Legit {
				fmt.Printf("\033[31m%d\033[0m  ", game.Board[i][j].Value) // Red
			} else {
				fmt.Printf("%d  ", game.Board[i][j].Value)
			}
		}
		fmt.Printf("\n")
	}
}
