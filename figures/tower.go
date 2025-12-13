package figures

import "fmt"

type Tower struct {
	x int
	y int
}

func (t *Tower) Threatens(cell Cell) bool {
	return cell.X == t.x || cell.Y == t.y
}

func (t *Tower) useCell(cell *Cell) {
	t.x = cell.X
	t.y = cell.Y
}

func (t Tower) String() string {
	return fmt.Sprintf("Tower{X=%d, Y=%d}", t.x, t.y)
}
