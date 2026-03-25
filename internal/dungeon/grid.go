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
	return 0, false
}

func (g Grid) Set(x, y int, tile Tile) bool {
	return false
}
