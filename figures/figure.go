package figures

type Figure interface {
	Threatens(cell Cell) bool
	useCell(cell *Cell)
}
