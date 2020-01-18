package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"image/color"
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

	screen.Fill(color.RGBA{
		R: 255,
		G: 0,
		B: 0,
		A: 255,
	})
	ebitenutil.DebugPrint(screen, "Hello, world")

	return nil
}
