package dungeon

import "testing"

func TestNewGrid_InitializeWalls(t *testing.T) {

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
