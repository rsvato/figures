package figures

import "fmt"

type Bishop struct {
	x int
	y int
}

func (b *Bishop) Threatens(cell Cell) bool {
	return AbsInt(cell.X-b.x) == (AbsInt(cell.Y - b.y))
}

func (b *Bishop) useCell(cell *Cell) {
	b.x = cell.X
	b.y = cell.Y
}

func (b *Bishop) String() string {
	return fmt.Sprintf("Bishop{X: %d, Y: %d}", b.x, b.y)
}
