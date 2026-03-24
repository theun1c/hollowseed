package dungeon

import (
	"testing"
)

func TestGrid_InitializeWalls(t *testing.T) {

	w, h := 100, 100

	g := NewGrid(w, h)

	if g.Width != w || g.Height != h {
		t.Fatalf("error: size error NewGrid(%d, %d) ", w, h)
	}

	for _, cell := range g.Cells {
		if cell.Tile != TileWall {
			t.Fatalf("error: Cell[%v] not a wall", cell.Tile)
		}
	}
}

func TestGrid_InBounds(t *testing.T) {

	w, h := 100, 100

	g := NewGrid(w, h)

	cases := []struct {
		name string
		x, y int
		want bool
	}{
		{"default", 0, 1, true},
		{"out of bounds (< w)", -1, 0, false},
		{"out of bounds (< h)", 0, -1, false},
		{"out of bounds (> h)", w, 0, false},
		{"out of bounds (> w)", 0, h, false},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			inBounds := g.InBounds(tc.x, tc.y)

			if tc.want != inBounds {
				t.Fatalf("error: InBounds(%d, %d) = %v , want %v", tc.x, tc.y, inBounds, tc.want)
			}
		})
	}

}
