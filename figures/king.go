package figures

import (
	"fmt"
)

type King struct {
	x int
	y int
}

func (k *King) Threatens(cell Cell) bool {
	return AbsInt(cell.X-k.x) == 1 || AbsInt(cell.Y-k.y) == 1
}

func (k *King) useCell(cell *Cell) {
	k.x = cell.X
	k.y = cell.Y
}

func (k King) String() string {
	return fmt.Sprintf("King{X=%d, Y=%d}", k.x, k.y)
}
