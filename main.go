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

	frame int64
}

var world = World{
	// Empty for now.
}

var gopherImage *ebiten.Image

func init() {
	var err error
	gopherImage, _, err = ebitenutil.NewImageFromFile("asset/gopher.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	if err := ebiten.Run(update, 320, 200, 2, "Hello, world!"); err != nil {
		log.Fatal(err)
	}
}

func update(screen *ebiten.Image) error {
	// Update state.
	world.frame++
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

	op := &ebiten.DrawImageOptions{}
	w, h := gopherImage.Size()
	op.GeoM.Translate(-float64(w)/2, -float64(h)/2)
	op.GeoM.Rotate(float64(world.frame) * 0.01)
	op.GeoM.Scale(0.2, 0.2)
	op.GeoM.Translate(160, 100)
	screen.DrawImage(gopherImage, op)

	return nil
}
