package figures

import (
	"testing"
)

func TestPlacePawn(t *testing.T) {
	desk := NewDesk(2, 2)
	pawn := &Pawn{}
	_, ok := desk.PlaceFigure(pawn)
	if !ok {
		t.Errorf("Figure should be placed correctly on empty board")
	}
	safeCells := desk.safeCells()
	if len(safeCells) != 2 {
		t.Errorf("After one pawn desk must contain 2 safe cells, but result is %d", len(safeCells))
	}
}

func TestPlaceFigureOptimal(t *testing.T) {
	desk := NewDesk(3, 3)
	queen := &Queen{} // A new queen

	placedFigure, ok := desk.PlaceFigure(queen)
	if !ok {
		t.Fatalf("Failed to place queen")
	}

	expectedX := 0
	expectedY := 0

	placedQueen, isQueen := (*placedFigure).(*Queen)
	if !isQueen {
		t.Fatalf("Placed figure is not a Queen")
	}

	if placedQueen.x != expectedX || placedQueen.y != expectedY {
		t.Errorf("Expected Queen to be placed at (%d, %d), but got (%d, %d)", expectedX, expectedY, placedQueen.x, placedQueen.y)
	}

	cell00, _ := desk.cellAt(0, 0)
	if !cell00.occupied {
		t.Errorf("Cell (0,0) should be occupied")
	}

	// Check some threatened cells
	cell01, _ := desk.cellAt(0, 1) // Threatened by (0,0)
	if !cell01.underThreat {
		t.Errorf("Cell (0,1) should be under threat")
	}
	cell10, _ := desk.cellAt(1, 0) // Threatened by (0,0)
	if !cell10.underThreat {
		t.Errorf("Cell (1,0) should be under threat")
	}
	cell11, _ := desk.cellAt(1, 1) // Threatened by (0,0)
	if !cell11.underThreat {
		t.Errorf("Cell (1,1) should be under threat")
	}
	cell22, _ := desk.cellAt(2, 2) // Threatened by (0,0)
	if !cell22.underThreat {
		t.Errorf("Cell (2,2) should be under threat")
	}

	cell12, _ := desk.cellAt(1, 2)
	if cell12.occupied || cell12.underThreat {
		t.Errorf("Cell (1,2) should be safe")
	}
	cell21, _ := desk.cellAt(2, 1)
	if cell21.occupied || cell21.underThreat {
		t.Errorf("Cell (2,1) should be safe")
	}
}
