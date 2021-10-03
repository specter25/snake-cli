package game

import (
	tl "github.com/JoelOtter/termloop"
)

// NewSnake will create a new snake and is called when the game is initialized.
func NewSnake() *Snake {
	snake := new(Snake)
	snake.Entity = tl.NewEntity(5, 5, 1, 1)
	snake.Direction = right
	snake.Bodylength = []Coordinates{
		{1, 6}, // Tail (The tail of the snake will stay the same unless the snake is not colliding with food)
		{2, 6}, // Body (The body will grow taller when a new head is created, the last piece of the body will become the tail if there is no collision with food)
		{3, 6}, // Head (Will become a piece of the body when a new head is created)
	}

	return snake
}

func (snake *Snake) Head() *Coordinates {
	return &snake.Bodylength[len(snake.Bodylength)-1]
}

func (snake *Snake) GetBodyLength() int {
	head := snake.Head()
	tail := snake.Bodylength[0]
	return (head.X - tail.X) + (head.Y - tail.Y)
}

// CollideBorder checks if the arena border contains the snakes head, if so it will return true.
func (snake *Snake) CollideBorder() bool {
	return gs.ArenaEntity.Contains(*snake.Head())
}

// CollideFood checks if the food contains the snakes head, if so it will return true.
func (snake *Snake) CollideFood() bool {
	return gs.FoodEntity.Contains(*snake.Head())
}

// SnakeCollision checks if the snakes body contains its head, if so it will return true.
func (snake *Snake) SnakeCollision() bool {
	return snake.Contains()
}

// Draw will check every tick and draw the snake on the screen, it also checks if the snake has any collisions
// using the funtions above.
func (snake *Snake) Draw(screen *tl.Screen) {
	newHead := *snake.Head()
	switch snake.Direction {
	case up:
		newHead.Y--
	case down:
		newHead.Y++
	case left:
		newHead.X--
	case right:
		newHead.X++
	}

	// Checks for a food collision using the collision function.
	if snake.CollideFood() {
		switch gs.FoodEntity.Emoji {
		case 'R':
			switch ts.GameDifficulty {
			case easy:
				if gs.FPS-3 <= 8 {
					UpdateScore(5)
				} else {
					gs.FPS -= 3
					UpdateScore(5)
					UpdateFPS()
				}
			case normal:
				if gs.FPS-2 <= 12 {
					UpdateScore(5)
				} else {
					gs.FPS -= 2
					UpdateScore(5)
					UpdateFPS()
				}
			case hard:
				if gs.FPS-1 <= 20 {
					UpdateScore(5)
				} else {
					gs.FPS--
					UpdateScore(5)
					UpdateFPS()
				}
			}
			snake.Bodylength = append(snake.Bodylength, newHead)
		case 'S':
			switch ts.GameDifficulty {
			case easy:
				gs.FPS++
			case normal:
				gs.FPS += 3
			case hard:
				gs.FPS += 5
			}
			UpdateFPS()
		default:
			UpdateScore(1)
			snake.Bodylength = append(snake.Bodylength, newHead)
		}
		gs.FoodEntity.MoveFood()
	} else {
		snake.Bodylength = append(snake.Bodylength[1:], newHead)
	}

	snake.SetPosition(newHead.X, newHead.Y)

	if snake.CollideBorder() || snake.SnakeCollision() {
		Gameover()
	}

	for _, c := range snake.Bodylength {
		screen.RenderCell(c.X, c.Y, &tl.Cell{
			Fg: CheckSelectedColor(counterSnake),
			Ch: '░',
		})
	}
}

// Contains checks if the snake contains the head of the snake, if so it will return true.
func (snake *Snake) Contains() bool {
	for i := 0; i < len(snake.Bodylength)-1; i++ {
		if *snake.Head() == snake.Bodylength[i] {
			return true
		}
	}
	return false
}
