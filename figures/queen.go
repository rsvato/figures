package figures

import "fmt"

type Queen struct {
	x int
	y int
}

func (q *Queen) Threatens(cell Cell) bool {
	return AbsInt(cell.X-q.x) == (AbsInt(cell.Y-q.y)) ||
		cell.X == q.x || cell.Y == q.y
}

func (q *Queen) useCell(cell *Cell) {
	q.x = cell.X
	q.y = cell.Y
}

func (q *Queen) String() string {
	return fmt.Sprintf("Queen{X: %d, Y: %d}", q.x, q.y)
}
