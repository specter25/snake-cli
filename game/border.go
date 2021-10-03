package game

import tl "github.com/JoelOtter/termloop"

// NewArena will create a new arena with the arena with and arena height given when this function was called in game.go
// This function will create a arena using the arena struct that can be found in the types.go file.
func NewArena(w, h int) *Border {
	arena := new(Border)
	arena.Width = w - 1
	arena.Height = h - 1
	arena.Entity = tl.NewEntity(1, 1, 1, 1)
	arena.ArenaBorder = make(map[Coordinates]int)

	for x := 0; x < arena.Width; x++ {
		arena.ArenaBorder[Coordinates{x, 0}] = 1
		arena.ArenaBorder[Coordinates{x, arena.Height}] = 1
	}

	for y := 0; y < arena.Height+1; y++ {
		arena.ArenaBorder[Coordinates{0, y}] = 1
		arena.ArenaBorder[Coordinates{arena.Width, y}] = 1
	}
	return arena
}

func (arena *Border) Contains(c Coordinates) bool {
	_, exists := arena.ArenaBorder[c]
	return exists
}

func (arena *Border) Draw(screen *tl.Screen) {
	for i := range arena.ArenaBorder {
		screen.RenderCell(i.X, i.Y, &tl.Cell{
			Bg: CheckSelectedColor(counterArena),
		})
	}
}
