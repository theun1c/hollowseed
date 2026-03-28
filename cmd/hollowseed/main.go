package main

import (
	"fmt"
	"hollowseed/internal/dungeon"
)

func main() {
	g := dungeon.NewGrid(10, 10)

	fmt.Print(g.Render())
}
