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
	ax, ay     float64
	angle      float64
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
	gopherImage, _, err = ebitenutil.NewImageFromFile("asset/zera.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	world.x = 400
	world.y = 300
	world.acc = 0
}

func main() {
	if err := ebiten.Run(update, 800, 600, 1, "Hello, world!"); err != nil {
		log.Fatal(err)
	}
}

func update(screen *ebiten.Image) error {
	// Check for key presses.
	if ebiten.IsKeyPressed(ebiten.KeyEscape) || ebiten.IsKeyPressed(ebiten.KeyQ) {
		os.Exit(0)
	}

	// Check for gamepad movement.
	hv := ebiten.GamepadAxis(0, 0)
	vv := ebiten.GamepadAxis(0, 1)
	hs := ebiten.GamepadAxis(0, 2)
	vs := ebiten.GamepadAxis(0, 3)

	var offsetX, offsetY float64
	acc := 15.0
	//if math.Abs(hv) > 0.20 || math.Abs(vv) > 0.20 {
	//	offsetX = hv*math.Cos(world.angle) + vv*math.Sin(world.angle)
	//	offsetY = -hv*math.Sin(world.angle) + vv*math.Cos(world.angle)
	//	world.x = world.x + (offsetX * acc)
	//	world.y = world.y + (offsetY * acc)
	//}

	lineLen := 1000.0
	dx := world.x + lineLen*hs
	dy := world.y + lineLen*vs

	if math.Abs(hs) > 0.20 || math.Abs(vs) > 0.20 {
		world.angle = math.Atan2(dy-world.y, dx-world.x) + math.Pi/2
	}

	world.ax = hv * hs
	world.x += world.ax

	if math.Abs(hv) > 0.10 {
		world.x += hv*acc + hs
	}
	if math.Abs(vv) > 0.10 {
		world.y += vv*acc + vs
	}
	if math.Abs(hv) <= 0.10 && math.Abs(vv) <= 0.10 {
		world.accStarted = 0.0
	} else {
		world.accStarted += 0.1
	}

	world.frame++

	if ebiten.IsDrawingSkipped() {
		return nil
	}

	// Move after check for draw.
	if math.Abs(hs) > 0.20 || math.Abs(vs) > 0.20 {
		world.angle = math.Atan2(dy-world.y, dx-world.x) + math.Pi/2
		ebitenutil.DrawLine(screen, world.x, world.y, dx, dy, color.RGBA{255, 255, 0, 255})
	}
	//ebitenutil.DebugPrint(screen, fmt.Sprintf("hs=%.3g, vs=%.3g, angle=%.3g", hs, vs, world.angle))
	ebitenutil.DebugPrint(screen, fmt.Sprintf("ox=%.3g oy=%.3g", offsetX, offsetY))
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%.3f", world.angle*180/math.Pi), int(world.x+20), int(world.y+20))

	op := &ebiten.DrawImageOptions{}
	w, h := gopherImage.Size()
	op.GeoM.Translate(-float64(w)/2, -float64(h)/2)
	//op.GeoM.Rotate(world.angle)
	op.GeoM.Scale(0.1, 0.1)
	op.GeoM.Translate(world.x, world.y)
	screen.DrawImage(gopherImage, op)

	return nil
}
