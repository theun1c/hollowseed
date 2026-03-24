package dungeon

type Grid struct {
	Width  int
	Height int
	Cells  []Cell
}

func NewGrid(width, height int) Grid {
	return Grid{
		Width:  width,
		Height: height,
		Cells:  make([]Cell, width*height),
	}
}
