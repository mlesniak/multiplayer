package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"image/color"
	"log"
)

type World struct {
	a  uint8
	up bool
}

var world = World{
	// Empty for now.
}

func main() {
	if err := ebiten.Run(update, 320, 200, 2, "Hello, world!"); err != nil {
		log.Fatal(err)
	}
}

func update(screen *ebiten.Image) error {
	// Update state.
	if world.a == 255 || world.a == 0 {
		world.up = !world.up
	}
	if world.up {
		world.a++
	} else {
		world.a--
	}

	if ebiten.IsDrawingSkipped() {
		return nil
	}

	// Draw.
	screen.Fill(color.RGBA{
		R: world.a,
		G: 0,
		B: 0,
		A: 255,
	})
	ebitenutil.DebugPrint(screen, "Hello, world")

	return nil
}
