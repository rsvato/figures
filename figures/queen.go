package figures


type Queen struct {
	FigureBase
}

func (q *Queen) Threatens(cell Cell) bool {
	return AbsInt(cell.X-q.x) == (AbsInt(cell.Y-q.y)) ||
		cell.X == q.x || cell.Y == q.y
}

func (q *Queen) String() string {
	return q.FigureBase.String("Queen")
}
