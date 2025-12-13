package figures

import (
	"fmt"
)

type Knight struct {
	x int
	y int
}

func (k *Knight) Threatens(cell Cell) bool {
	distX := AbsInt(cell.X - k.x)
	distY := AbsInt(cell.Y - k.y)
	return distX*distY == 2 && distX != distY
}

func (k *Knight) useCell(cell *Cell) {
	k.x = cell.X
	k.y = cell.Y
}

func (k Knight) String() string {
	return fmt.Sprintf("King{X=%d, Y=%d}", k.x, k.y)
}
