package game

import (
	"math/rand"
	"time"

	tl "github.com/JoelOtter/termloop"
)

func NewFood() *Target {
	food := new(Target)
	food.Entity = tl.NewEntity(1, 1, 2, 2)
	food.MoveFood()

	return food
}

// MoveFood moves the food into a new random position.
func (food *Target) MoveFood() {

	NewX := RandomInsideArena(gameWidth-5, 5)
	NewY := RandomInsideArena(gameHeight-5, 5)

	// Changes the X and Y coordinates of the food.
	food.Foodposition.X = NewX
	food.Foodposition.Y = NewY
	food.Emoji = RandomFood()

	food.SetPosition(food.Foodposition.X, food.Foodposition.Y)
}

// RandomFood will use the ASCII-charset to pick a random rune from the slice and print it out as food.
func RandomFood() rune {
	// This slice contains all of the possible food icons.
	emoji := []rune{
		'R', // Favourite dish, extra points!!!
		'🍒',
		'🍍',
		'🍑',
		'🍇',
		'🍏',
		'🍌',
		'🍫',
		'🍭',
		'🍕',
		'🍩',
		'🍗',
		'🍖',
		'🍬',
		'🍤',
		'🍪',
		'S', // You do not want to eat the skull
	}

	rand.Seed(time.Now().UnixNano())

	// Return a random rune picked from the slice
	return emoji[rand.Intn(len(emoji))]
}

// Draw will print out the food on the screen.
func (food *Target) Draw(screen *tl.Screen) {
	screen.RenderCell(food.Foodposition.X, food.Foodposition.Y, &tl.Cell{
		Ch: food.Emoji,
	})
}

// Contains checks if food contains the coordinates, if so this will return a bool.
func (food *Target) Contains(c Coordinates) bool {
	return c.X == food.Foodposition.X && c.Y == food.Foodposition.Y
}

// RandomInsideArena will the minimal, which is just inside the border and the maximal, being the arena width or height.
func RandomInsideArena(iMax int, iMin int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(iMax-iMin) + iMin
}
