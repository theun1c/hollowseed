package dungeon

import (
	"testing"
)

func TestGrid_InitializeWalls(t *testing.T) {

	w, h := 100, 100

	g := NewGrid(w, h)

	if g.Width != w || g.Height != h {
		t.Fatalf("error: size init error")
	}

	for _, cell := range g.Cells {
		if cell.Tile != TileWall {
			t.Fatalf("error: not wall after init")
		}
	}
}

func TestGrid_InBounds(t *testing.T) {
	w, h := 100, 100

	g := NewGrid(w, h)

	if !(g.InBounds(0, 1)) {
		t.Fatalf("error: bounds error")
	}

}
