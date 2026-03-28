package dungeon

type Grid struct {
	Width  int
	Height int
	Cells  []Cell
}

func NewGrid(width, height int) Grid {

	cells := make([]Cell, width*height)

	x, y := 0, 0

	for i := range cells {

		cells[i] = Cell{
			X:    x,
			Y:    y,
			Tile: TileWall,
		}

		x++

		if width == x {
			y++
			x = 0
		}

	}

	return Grid{
		Width:  width,
		Height: height,
		Cells:  cells,
	}
}

func (g Grid) InBounds(x, y int) bool {
	if x > g.Width-1 || y > g.Height-1 || x < 0 || y < 0 {
		return false
	}

	return true
}

func (g Grid) Get(x, y int) (Tile, bool) {
	if g.InBounds(x, y) {
		return g.Cells[y*g.Width+x].Tile, true
	}

	return 0, false
}

func (g Grid) Set(x, y int, tile Tile) bool {
	if g.InBounds(x, y) {
		g.Cells[y*g.Width+x].Tile = tile
		return true
	}

	return false
}
