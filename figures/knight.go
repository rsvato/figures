package figures


type Knight struct {
	FigureBase
}

func (k *Knight) Threatens(cell Cell) bool {
	distX := AbsInt(cell.X - k.x)
	distY := AbsInt(cell.Y - k.y)
	return distX*distY == 2 && distX != distY
}

func (k *Knight) String() string {
	return k.FigureBase.String("Knight")
}
