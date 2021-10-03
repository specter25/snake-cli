package game

import (
	tl "github.com/JoelOtter/termloop"
	tb "github.com/nsf/termbox-go"
)

var counterSnake = 10
var counterArena = 10

// Tick listens for a keypress and then returns a direction for the snake.
func (snake *Snake) Tick(event tl.Event) {
	// Checks if the event is a keyevent.
	if event.Type == tl.EventKey {
		switch event.Key {
		case tl.KeyArrowRight:
			if snake.Direction != left {
				snake.Direction = right
			}
		case tl.KeyArrowLeft:
			if snake.Direction != right {
				snake.Direction = left
			}
		case tl.KeyArrowUp:
			if snake.Direction != down {
				snake.Direction = up
			}
		case tl.KeyArrowDown:
			if snake.Direction != up {
				snake.Direction = down
			}
		}
	}
	UpdateText()
}

func (gos *Gameoverscreen) Tick(event tl.Event) {
	if event.Type == tl.EventKey {
		switch event.Key {
		case tl.KeyHome:
			RestartGame()
		case tl.KeyDelete:
			tb.Close()
		}
	}
}

// Tick will listen for a keypress to initiate the game.
func (ts *LandingScreen) Tick(event tl.Event) {
	if event.Type == tl.EventKey {
		if event.Key == tl.KeyEnter {
			gs = NewGamescreen()
			sg.Screen().SetLevel(gs)
		}
	}
}

func CheckSelectedColor(c int) tl.Attr {
	switch c {
	case 10:
		return tl.ColorWhite
	case 12:
		return tl.ColorRed
	case 14:
		return tl.ColorGreen
	case 16:
		return tl.ColorBlue
	case 18:
		return tl.ColorYellow
	case 20:
		return tl.ColorMagenta
	case 22:
		return tl.ColorCyan
	default:
		return tl.ColorDefault
	}
}
