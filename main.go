package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"log"
)

func main() {
	if err := ebiten.Run(update, 320, 200, 2, "Hello, world!"); err != nil {
		log.Fatal(err)
	}
}

func update(screen *ebiten.Image) error {
	if ebiten.IsDrawingSkipped() {
		return nil
	}

	ebitenutil.DebugPrint(screen, "Hello, world")
	return nil
}
