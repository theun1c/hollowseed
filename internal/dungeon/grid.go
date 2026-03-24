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

