package model

type Borad struct {
	Cells [][]Cell
}

func (Borad) Create(x, y int) Borad {
	Borad := Borad{}
	Borad.Cells = [][]Cell{}
	return Borad
}
