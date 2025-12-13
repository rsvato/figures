package figures

import (
	"fmt"
)

type Pawn struct {
	x int
	y int
}

func (p *Pawn) Threatens(cell Cell) bool {
	return AbsInt(cell.X-p.x) == 1 && cell.Y-p.y == 1
}

func (p *Pawn) useCell(cell *Cell) {
	p.x = cell.X
	p.y = cell.Y
}

func (p Pawn) String() string {
	return fmt.Sprintf("Pawn{X=%d, Y=%d}", p.x, p.y)
}
