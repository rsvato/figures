package figures


type King struct {
	FigureBase
}

func (k *King) Threatens(cell Cell) bool {
	return AbsInt(cell.X-k.x) == 1 || AbsInt(cell.Y-k.y) == 1
}

func (k *King) String() string {
	return k.FigureBase.String("King")
}
