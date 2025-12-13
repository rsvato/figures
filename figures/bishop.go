package figures


type Bishop struct {
	FigureBase
}

func (b *Bishop) Threatens(cell Cell) bool {
	return AbsInt(cell.X-b.x) == (AbsInt(cell.Y - b.y))
}

func (b *Bishop) String() string {
	return b.FigureBase.String("Bishop")
}
