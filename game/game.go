package game

import (
	"fmt"
	"io/ioutil"

	tl "github.com/JoelOtter/termloop"
)

// StartGame will start the game with the tilescreen.
func StartGame(w int, h int) {
	sg = tl.NewGame()

	gameWidth = w
	gameHeight = h

	ts := NewTitleScreen()

	ts.AddEntity(ts.Logo)

	for _, v := range ts.OptionsText {
		ts.AddEntity(v)
	}

	sg.Screen().SetFps(10)
	sg.Screen().SetLevel(ts)
	sg.Start()
}

// NewTitleScreen will create a new titlescreen and return it.
func NewTitleScreen() *LandingScreen {
	ts = new(LandingScreen)
	ts.Level = tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorBlack,
	})

	logofile, _ := ioutil.ReadFile("util/titlescreen-logo.txt")
	ts.Logo = tl.NewEntityFromCanvas(10, 3, tl.CanvasFromString(string(logofile)))

	ts.GameDifficulty = normal
	ts.OptionsText = []*tl.Text{
		tl.NewText(30, 15, "Press ENTER to start!", tl.ColorBlue, tl.ColorWhite),
	}

	return ts
}

func NewGamescreen() *Gamescreen {
	// Creates the gamescreen level and create the entities
	gs = new(Gamescreen)
	gs.Level = tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorBlack,
	})
	SetDiffiultyFPS()
	gs.Score = 0
	gs.SnakeEntity = NewSnake()
	gs.ArenaEntity = NewArena(gameWidth, gameHeight)
	gs.FoodEntity = NewFood()
	gs.SidepanelObject = NewSidepanel()

	// Add entities for the game level.
	gs.AddEntity(gs.FoodEntity)
	gs.AddEntity(gs.SidepanelObject.Background)
	gs.AddEntity(gs.SidepanelObject.ScoreText)
	gs.AddEntity(gs.SidepanelObject.SpeedText)
	gs.AddEntity(gs.SidepanelObject.DifficultyText)
	gs.AddEntity(gs.SidepanelObject.Width)
	gs.AddEntity(gs.SidepanelObject.Height)
	gs.AddEntity(gs.SidepanelObject.HeadX)
	gs.AddEntity(gs.SidepanelObject.Direction)
	gs.AddEntity(gs.SidepanelObject.HeadY)
	gs.AddEntity(gs.SnakeEntity)
	gs.AddEntity(gs.ArenaEntity)

	sg.Screen().SetFps(gs.FPS)

	return gs
}

// NewSidepanel will create a new sidepanel given the arena height and width.
func NewSidepanel() *Panel {
	sp = new(Panel)

	head := gs.SnakeEntity.Head()

	sp.Background = tl.NewRectangle(gameWidth+1, 0, 45, gameHeight, tl.ColorWhite)
	sp.ScoreText = tl.NewText(gameWidth+2, 1, fmt.Sprintf("Score: %d", gs.Score), tl.ColorBlack, tl.ColorWhite)
	sp.SpeedText = tl.NewText(gameWidth+2, 3, fmt.Sprintf("Speed: %.0f", gs.FPS), tl.ColorBlack, tl.ColorWhite)
	sp.DifficultyText = tl.NewText(gameWidth+2, 5, fmt.Sprintf("Difficulty: %s", Difficulty), tl.ColorBlack, tl.ColorWhite)
	sp.Width = tl.NewText(gameWidth+2, 6, fmt.Sprintf("Arena Width: %d", gameWidth), tl.ColorBlack, tl.ColorWhite)
	sp.Height = tl.NewText(gameWidth+2, 7, fmt.Sprintf("Arena Height: %d", gameHeight), tl.ColorBlack, tl.ColorWhite)
	sp.Height = tl.NewText(gameWidth+2, 8, fmt.Sprintf("Arena Height: %d", gameHeight), tl.ColorBlack, tl.ColorWhite)
	sp.Direction = tl.NewText(gameWidth+2, 9, fmt.Sprintf("Snake Direction: %s", GetDirection(gs.SnakeEntity.Direction)), tl.ColorBlack, tl.ColorWhite)
	sp.HeadX = tl.NewText(gameWidth+2, 10, fmt.Sprintf("X Coordiate(Snake's Head): %d", head.X), tl.ColorBlack, tl.ColorWhite)
	sp.HeadY = tl.NewText(gameWidth+2, 11, fmt.Sprintf("y Coordiate(Snake's Head) %d", head.Y), tl.ColorBlack, tl.ColorWhite)

	return sp
}

func Gameover() {
	gos := new(Gameoverscreen)
	gos.Level = tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorBlack,
	})
	logofile, _ := ioutil.ReadFile("util/gameover-logo.txt")
	gos.Logo = tl.NewEntityFromCanvas(10, 3, tl.CanvasFromString(string(logofile)))
	gos.Finalstats = []*tl.Text{
		tl.NewText(10, 13, fmt.Sprintf("Score: %d", gs.Score), tl.ColorWhite, tl.ColorBlack),
		tl.NewText(10, 15, fmt.Sprintf("Speed: %.0f", gs.FPS), tl.ColorWhite, tl.ColorBlack),
		tl.NewText(10, 17, fmt.Sprintf("Difficulty: %s", Difficulty), tl.ColorWhite, tl.ColorBlack),
	}
	gos.OptionsBackground = tl.NewRectangle(45, 12, 45, 7, tl.ColorWhite)
	gos.OptionsText = []*tl.Text{
		tl.NewText(47, 13, "Press \"Home\" to restart!", tl.ColorBlack, tl.ColorWhite),
		tl.NewText(47, 15, "Press \"Delete\" to quit!", tl.ColorBlack, tl.ColorWhite),
	}

	for _, v := range gos.Finalstats {
		gos.AddEntity(v)
	}
	gos.AddEntity(gos.Logo)
	gos.AddEntity(gos.OptionsBackground)

	for _, vv := range gos.OptionsText {
		gos.AddEntity(vv)
	}

	sg.Screen().SetLevel(gos)
}

// UpdateScore updates the score with the given amount of points.
func UpdateScore(amount int) {
	gs.Score += amount
	sp.ScoreText.SetText(fmt.Sprintf("Score: %d", gs.Score))
}

func UpdateText() {
	sp.Direction.SetText(fmt.Sprintf("Snake Direction: %s", GetDirection(gs.SnakeEntity.Direction)))
	sp.HeadX.SetText(fmt.Sprintf("X Coordiate(Snake's Head): %d", gs.SnakeEntity.Head().X))
	sp.HeadY.SetText(fmt.Sprintf("Y Coordiate(Snake's Head): %d", gs.SnakeEntity.Head().Y))
}

// UpdateFPS updates the fps text.
func UpdateFPS() {
	sg.Screen().SetFps(gs.FPS)
	sp.SpeedText.SetText(fmt.Sprintf("Speed: %.0f", gs.FPS))
}

// RestartGame will restart the game and reset the position of the food and the snake to prevent collision issues.
func RestartGame() {
	gs.RemoveEntity(gs.SnakeEntity)
	gs.RemoveEntity(gs.FoodEntity)

	gs.SnakeEntity = NewSnake()
	gs.FoodEntity = NewFood()

	SetDiffiultyFPS()
	gs.Score = 0

	sp.ScoreText.SetText(fmt.Sprintf("Score: %d", gs.Score))
	sp.SpeedText.SetText(fmt.Sprintf("Speed: %.0f", gs.FPS))

	gs.AddEntity(gs.SnakeEntity)
	gs.AddEntity(gs.FoodEntity)
	sg.Screen().SetFps(gs.FPS)
	sg.Screen().SetLevel(gs)
}

func SetDiffiultyFPS() {
	switch ts.GameDifficulty {
	case easy:
		gs.FPS = 8
	case normal:
		gs.FPS = 12
	case hard:
		gs.FPS = 25
	}
}

func GetDirection(dir direction) string {
	directionText := map[int]string{
		0: "up",
		1: "down",
		2: "left",
		3: "right",
	}
	return directionText[int(dir)]
}
