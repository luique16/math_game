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

func (g *Game) Erase(x, y int) {
	if g.Board[y][x].State != Active {
		return
	} else if g.Board[y][x].Legit == true {
		g.Lives--
		return
	}

	g.Board[y][x].State = Erased
}

func (g *Game) Select(x, y int) {
	if g.Board[y][x].State != Active {
		return
	} else if g.Board[y][x].Legit == false {
		g.Lives--
		return
	}

	g.Board[y][x].State = Selected
}
