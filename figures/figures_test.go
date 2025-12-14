package figures

import (
	"log"
	"testing"
)

func TestQueenThreatens(t *testing.T) {
	queen := &Queen{}
	log.Printf("Queen is %s", queen)
	cell := &Cell{X: 0, Y: 0}
	queen.useCell(&Cell{X: 0, Y: 0})
	if !queen.Threatens(*cell) {
		t.Fatal("Cell 0,0 must be under threat too")
	}
	if (!queen.Threatens(Cell{X: 8, Y: 8})) {
		t.Error("Cell on the same diagonal must be under threat")
	}
	if (!queen.Threatens(Cell{X: 0, Y: 8})) {
		t.Error("Cell on the same horizonal must be under threat")
	}
	if (!queen.Threatens(Cell{X: 0, Y: -8})) {
		t.Error("Cell on the same horizonal must be under threat")
	}
	if (!queen.Threatens(Cell{X: -8, Y: 0})) {
		t.Error("Cell on the same vertical must be under threat")
	}
	if (!queen.Threatens(Cell{X: 8, Y: 0})) {
		t.Error("Cell on the same vertical must be under threat")
	}
	if (!queen.Threatens(Cell{X: -8, Y: 0})) {
		t.Error("Cell on the same vertical must be under threat")
	}
}
