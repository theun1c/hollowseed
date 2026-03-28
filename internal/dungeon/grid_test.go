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
		{"in bounds (x <= w-1)", w - 1, 0, true},
		{"in bounds (y <= h-1)", 0, h - 1, true},
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

func TestGrid_Get(t *testing.T) {

	w, h := 100, 50

	g := NewGrid(w, h)

	cases := []struct {
		name string
		x, y int
		want bool
	}{
		{"in bounds", w - 1, h - 1, true},
		{"out of bounds", w, h, false},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if _, ok := g.Get(tc.x, tc.y); ok != tc.want {
				t.Fatalf("error: cannot get beyond bounds tile value (%d, %d)", tc.x, tc.y)
			}
		})
	}

}

func TestGrid_Set(t *testing.T) {
	w, h := 100, 50

	g := NewGrid(w, h)

	cases := []struct {
		name string
		x, y int
		tile Tile
		want bool
	}{
		{"in bounds tile floor", w - 1, h - 1, TileFloor, true},
		{"in bounds tile wall", w - 1, h - 1, TileWall, true},
		{"out of bounds tile floor", w, h, TileFloor, false},
		{"out of bounds tile wall", w, h, TileWall, false},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if ok := g.Set(tc.x, tc.y, tc.tile); ok != tc.want {
				t.Fatalf("error: cannot set beyond bounds tile value (%d, %d, %v)", tc.x, tc.y, tc.tile)
			} else {
				tile, ok := g.Get(tc.x, tc.y)
				if tc.want {
					if tile != tc.tile || !ok {
						t.Fatalf("error: cannot get and set beyond bounds tile value (%d, %d, %v)", tc.x, tc.y, tc.tile)
					}
				}
			}
		})
	}

}
