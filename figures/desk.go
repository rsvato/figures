package figures

import (
	"fmt"
	"log"
)

type Desk struct {
	X     int
	Y     int
	cells []Cell
}

type Cell struct {
	X           int
	Y           int
	occupied    bool
	underThreat bool
}

func (c *Cell) String() string {
	return fmt.Sprintf("Cell{X=%d, Y=%d}", c.X, c.Y)
}

type Figure interface {
	Threatens(cell Cell) bool
	useCell(cell *Cell)
}

func NewDesk(x int, y int) Desk {
	result := Desk{X: x, Y: y}
	result.cells = make([]Cell, x*y)
	for j := 0; j < y; j++ {
		for i := 0; i < x; i++ {
			index := j*x + i
			result.cells[index] = Cell{X: i, Y: j}
		}
	}
	return result
}

func (d *Desk) Print() {
	for y := range d.Y {
		for x := range d.X {
			cell, _ := d.cellAt(x, y)
			if cell.occupied {
				fmt.Print(" f")
			} else if cell.underThreat {
				fmt.Print(" !")
			} else {
				fmt.Print(" .")
			}
		}
		fmt.Print("\n")
	}
}

func (d *Desk) cellAt(x int, y int) (*Cell, error) {
	// Предполагаем, что Desk.X - это ширина
	if x >= d.X || y >= d.Y || x < 0 || y < 0 {
		return nil, fmt.Errorf("coordinates %d:%d out of bounds %d:%d", x, y, d.X, d.Y)
	}

	// Вычисляем индекс, как в NewDesk
	index := y*d.X + x

	// Возвращаем указатель на элемент в исходном массиве
	return &d.cells[index], nil
}

func (d *Desk) PlaceFigure(f Figure) (*Figure, bool) {
	safeCells := d.safeCells()
	if len(safeCells) == 0 {
		log.Println("No safe cells to place figure")
		return nil, false
	}

	//TODO: select with max possible cells left
	safeCell := safeCells[0]
	f.useCell(safeCell)
	safeCell.Use()
	for i := range d.cells {
		cell := &d.cells[i]
		if f.Threatens(*cell) {
			cell.underThreat = true
		}
	}

	return &f, true
}

func (c *Cell) Use() {
	c.occupied = true
}

func (d *Desk) safeCells() []*Cell {
	var result []*Cell
	for i := range d.cells {
		cell := &d.cells[i]
		if cell.occupied || cell.underThreat {
			continue
		} else {
			result = append(result, cell)
		}
	}
	return result
}
