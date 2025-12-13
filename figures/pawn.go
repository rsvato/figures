package figures

type Pawn struct {
	FigureBase
}

func (p *Pawn) Threatens(cell Cell) bool {
	// this is wrong, but let's assume that pawn can beat in every direction like a King.
	// We don't have white and black here.
	return AbsInt(cell.X-p.x) == 1 && AbsInt(cell.Y-p.y) == 1
}

func (p *Pawn) String() string {
	return p.FigureBase.String("Pawn")
}
