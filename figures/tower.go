package figures


type Tower struct {
	FigureBase
}

func (t *Tower) Threatens(cell Cell) bool {
	return cell.X == t.x || cell.Y == t.y
}

func (t *Tower) String() string {
	return t.FigureBase.String("Tower")
}
