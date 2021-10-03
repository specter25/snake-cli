package main

import (
	"fmt"

	"github.com/specter25/snake-cli/src/game"
)

func main() {
	var width int
	var height int
	fmt.Println("Enter arena width (recommended width 70)")
	fmt.Scanf("%d", &width)
	fmt.Println("Enter arena height (recommended height 25)")
	fmt.Scanf("%d", &height)
	game.NewGame(width, height)
}
