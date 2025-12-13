package figures

import "fmt"

type FigureBase struct {
	x int
	y int
}

func (fb *FigureBase) useCell(cell *Cell) {
	fb.x = cell.X
	fb.y = cell.Y
}

func (fb *FigureBase) clearCell() {
	fb.x = -1
	fb.y = -1
}

func (fb *FigureBase) String(name string) string {
	return fmt.Sprintf("%s{X: %d, Y: %d}", name, fb.x, fb.y)
}
