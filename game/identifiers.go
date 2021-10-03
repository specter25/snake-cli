package game

import tl "github.com/JoelOtter/termloop"

// Own created types.
type direction int
type difficulty int

// Game options
var Difficulty = "Normal"
var ColorObject = "Snake"

type Coordinates struct {
	X int
	Y int
}

const (
	easy difficulty = iota
	normal
	hard
)

const (
	up direction = iota
	down
	left
	right
)

type LandingScreen struct {
	tl.Level
	Logo           *tl.Entity
	GameDifficulty difficulty
	OptionsText    []*tl.Text
}

type Gameoverscreen struct {
	tl.Level
	Logo              *tl.Entity
	Finalstats        []*tl.Text
	OptionsBackground *tl.Rectangle
	OptionsText       []*tl.Text
}

type Gamescreen struct {
	tl.Level
	FPS             float64
	Score           int
	SnakeEntity     *Snake
	FoodEntity      *Target
	ArenaEntity     *Border
	SidepanelObject *Panel
}

type Panel struct {
	Background     *tl.Rectangle
	Instructions   []string
	ScoreText      *tl.Text
	SpeedText      *tl.Text
	DifficultyText *tl.Text
	Width          *tl.Text
	Height         *tl.Text
	Direction      *tl.Text
	HeadX          *tl.Text
	HeadY          *tl.Text
	Round          *tl.Text
	Bodylength     *tl.Text
}

type Border struct {
	*tl.Entity
	Width       int
	Height      int
	ArenaBorder map[Coordinates]int
}

type Snake struct {
	*tl.Entity
	Direction  direction
	Length     int
	Bodylength []Coordinates
	Speed      int
}

type Target struct {
	*tl.Entity
	Foodposition Coordinates
	Emoji        rune
}

//Game Object Variables.
var sg *tl.Game
var sp *Panel
var gs *Gamescreen
var ts *LandingScreen
var gameWidth int
var gameHeight int
var round int = 0
