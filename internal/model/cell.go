package model

type Cell struct {
	Name         string
	Xaxis        int
	Yaxis        int
	IsFill       bool
	FilledSquare Square
}

func (Cell) Create(name string, x, y int) Cell {
	Cell := Cell{}
	Cell.Name = name
	Cell.Xaxis = x
	Cell.Yaxis = y
	Cell.IsFill = false
	Cell.FilledSquare = nil
	return Cell
}
