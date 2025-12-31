package models

type Game struct {
	Rows     int
	Cols     int
	RowsSums []int
	ColsSums []int
	Board    [][]Cell
	Lives    int
}

type CellState int

const (
	Active CellState = iota
	Erased
	Selected
)

type Cell struct {
	Value int
	Legit bool
	State CellState
}
