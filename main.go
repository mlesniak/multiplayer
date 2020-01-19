package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"
	"image/color"
	"log"
	"os"
	"strconv"
	"strings"
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

var (
	gamepadIDs = map[int]struct{}{}
)

func init() {
	var err error
	gopherImage, _, err = ebitenutil.NewImageFromFile("asset/gopher.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	if err := ebiten.Run(update, 320, 200, 3, "Hello, world!"); err != nil {
		log.Fatal(err)
	}
}

func update(screen *ebiten.Image) error {
	// Check for gamepads.
	for _, id := range inpututil.JustConnectedGamepadIDs() {
		log.Printf("gamepad connected: id: %d", id)
		gamepadIDs[id] = struct{}{}
	}
	for id := range gamepadIDs {
		if inpututil.IsGamepadJustDisconnected(id) {
			log.Printf("gamepad disconnected: id: %d", id)
			delete(gamepadIDs, id)
		}
	}

	ids := ebiten.GamepadIDs()
	axes := map[int][]string{}
	pressedButtons := map[int][]string{}

	for _, id := range ids {
		maxAxis := ebiten.GamepadAxisNum(id)
		//log.Printf("%d/maxAxis: %d\n", id, maxAxis)
		for a := 0; a < maxAxis; a++ {
			v := ebiten.GamepadAxis(id, a)
			axes[id] = append(axes[id], fmt.Sprintf("%d:%0.2f", a, v))
		}
		maxButton := ebiten.GamepadButton(ebiten.GamepadButtonNum(id))
		for b := ebiten.GamepadButton(id); b < maxButton; b++ {
			if ebiten.IsGamepadButtonPressed(id, b) {
				pressedButtons[id] = append(pressedButtons[id], strconv.Itoa(int(b)))
			}

			// Log button events.
			if inpututil.IsGamepadButtonJustPressed(id, b) {
				log.Printf("button pressed: id: %d, button: %d", id, b)
			}
			if inpututil.IsGamepadButtonJustReleased(id, b) {
				log.Printf("button released: id: %d, button: %d", id, b)
			}
		}
	}

	// Check for key presses.
	if ebiten.IsKeyPressed(ebiten.KeyEscape) || ebiten.IsKeyPressed(ebiten.KeyQ) {
		os.Exit(0)
	}

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

	str := ""
	if len(ids) > 0 {
		for _, id := range ids {
			str += fmt.Sprintf("Gamepad (ID: %d):\n", id)
			str += fmt.Sprintf("  Axes:    %s\n", strings.Join(axes[id], ", "))
			str += fmt.Sprintf("  Buttons: %s\n", strings.Join(pressedButtons[id], ", "))
			str += "\n"
		}
	} else {
		str = "Please connect your gamepad."
	}
	ebitenutil.DebugPrintAt(screen, str, 0, 12)

	as := []float64{
		ebiten.GamepadAxis(0, 0), // 0：The horizontal value of the left axis
		ebiten.GamepadAxis(0, 1), // 1：The vertical value of the left axis
		ebiten.GamepadAxis(0, 2), // 2：The horizontal value of the right axis
		ebiten.GamepadAxis(0, 3)} // 3：The vertical value of the right axis
	ebitenutil.DebugPrint(screen, fmt.Sprintf("0: %0.6f, 1: %0.6f", as[0], as[1]))
	ebitenutil.DebugPrint(screen, fmt.Sprintf("\n2: %0.6f, 3: %0.6f", as[2], as[3]))

	return nil
}
