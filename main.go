package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/text"
	"golang.org/x/image/font"
	"image/color"
	"log"
	"math"
	"os"
	"time"
)

func main() {
	if err := ebiten.Run(update, globalConfig.width, globalConfig.height, 1, "Hello, world!"); err != nil {
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

	acc := 15.0

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

	msPassed := time.Now().Sub(world.timer).Milliseconds()
	if msPassed > 5*1000 {
		world.timer = time.Now()
	}

	// --------------------------------------------------------------------------------------------------------------------
	if ebiten.IsDrawingSkipped() {
		return nil
	}

	// Move after check for draw.
	if math.Abs(hs) > 0.20 || math.Abs(vs) > 0.20 {
		world.angle = math.Atan2(dy-world.y, dx-world.x) + math.Pi/2
		ebitenutil.DrawLine(screen, world.x, world.y, dx, dy, color.RGBA{255, 255, 0, 255})
	}

	msg := fmt.Sprintf("--- %.2f ---", float64(msPassed)/1000.0)
	b, _ := font.BoundString(arcadeFont, msg)
	a := b.Max.X.Ceil()
	text.Draw(screen, msg, arcadeFont, globalConfig.width/2-a/2, 20, color.White)

	op := &ebiten.DrawImageOptions{}
	w, h := gopherImage.Size()
	op.GeoM.Translate(-float64(w)/2, -float64(h)/2)
	//op.GeoM.Rotate(world.angle)
	op.GeoM.Scale(0.1, 0.1)
	op.GeoM.Translate(world.x, world.y)
	screen.DrawImage(gopherImage, op)

	return nil
}
