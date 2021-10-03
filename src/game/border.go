package game

import tl "github.com/JoelOtter/termloop"

// NewBorder will create a new border with the border with and border height given when this function was called in game.go
// This function will create a border using the border struct that can be found in the types.go file.
func NewBorder(w, h int) *Border {
	border := new(Border)
	border.Width = w - 1
	border.Height = h - 1
	border.Entity = tl.NewEntity(1, 1, 1, 1)
	border.ArenaBorder = make(map[Coordinates]int)

	for x := 0; x < border.Width; x++ {
		border.ArenaBorder[Coordinates{x, 0}] = 1
		border.ArenaBorder[Coordinates{x, border.Height}] = 1
	}

	for y := 0; y < border.Height+1; y++ {
		border.ArenaBorder[Coordinates{0, y}] = 1
		border.ArenaBorder[Coordinates{border.Width, y}] = 1
	}
	return border
}

func (border *Border) Contains(c Coordinates) bool {
	_, exists := border.ArenaBorder[c]
	return exists
}

func (border *Border) Draw(screen *tl.Screen) {
	for i := range border.ArenaBorder {
		screen.RenderCell(i.X, i.Y, &tl.Cell{
			Bg: CheckSelectedColor(counterArena),
		})
	}
}
