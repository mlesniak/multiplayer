package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"image/color"
	"log"
	"math"
	"os"
)

type World struct {
	a  uint8
	up bool

	// Position of gopher.
	x, y       float64
	acc        float64
	accStarted float64

	frame      int64
	fullscreen bool
}

var world = World{
	// Empty for now.
}

var gopherImage *ebiten.Image

var (
	gamepadIDs = map[int]struct{}{}
)

func init() {
	var err error
	gopherImage, _, err = ebitenutil.NewImageFromFile("asset/ship.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	world.x = 160
	world.y = 100
	world.acc = 0
}

func main() {
	if err := ebiten.Run(update, 320, 200, 2, "Hello, world!"); err != nil {
		log.Fatal(err)
	}
}

func update(screen *ebiten.Image) error {
	// Check for key presses.
	if ebiten.IsKeyPressed(ebiten.KeyEscape) || ebiten.IsKeyPressed(ebiten.KeyQ) {
		os.Exit(0)
	}

	//if ebiten.IsKeyPressed(ebiten.KeyF) {
	//	world.fullscreen = !world.fullscreen
	//	ebiten.SetFullscreen(world.fullscreen)
	//}

	// Check for gamepad movement.
	hv := ebiten.GamepadAxis(0, 0)
	vv := ebiten.GamepadAxis(0, 1)
	acc := 10.0
	if math.Abs(hv) > 0.10 {
		world.x += hv * acc
	}
	if math.Abs(vv) > 0.10 {
		world.y += vv * acc
	}
	if math.Abs(hv) <= 0.10 && math.Abs(vv) <= 0.10 {
		world.accStarted = 0.0
	} else {
		world.accStarted += 0.1
	}

	hs := ebiten.GamepadAxis(0, 2)
	vs := ebiten.GamepadAxis(0, 3)

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

	op := &ebiten.DrawImageOptions{}
	w, h := gopherImage.Size()
	op.GeoM.Translate(-float64(w)/2, -float64(h)/2)
	//op.GeoM.Rotate(float64(world.frame) * 0.01)
	op.GeoM.Scale(0.2, 0.2)
	op.GeoM.Translate(world.x, world.y)
	screen.DrawImage(gopherImage, op)

	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("0: %0.6f, 1: %0.6f", hv, vv), 0, 12)

	// Move after check for draw.
	if math.Abs(hs) > 0.20 || math.Abs(vs) > 0.20 {
		fmt.Println(hs, vs)
		lineLen := 1000.0
		dx := world.x + lineLen*hs
		dy := world.y + lineLen*vs
		ebitenutil.DrawLine(screen, world.x, world.y, dx, dy, color.RGBA{255, 255, 0, 255})
	}

	return nil
}
