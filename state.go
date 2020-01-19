package main

import (
	"github.com/hajimehoshi/ebiten"
	"math"
	"os"
	"time"
)

type State struct {
	// Position of gopher.
	x, y       float64
	ax, ay     float64
	angle      float64
	acc        float64
	accStarted float64

	fullscreen bool
	timer      time.Time

	hv, vv, hs, vs float64
	msPassed       int64
}

var state = State{
	x:     400,
	y:     300,
	acc:   0,
	timer: time.Now(),
}

func updateState() {
	// Check for key presses.
	if ebiten.IsKeyPressed(ebiten.KeyEscape) || ebiten.IsKeyPressed(ebiten.KeyQ) {
		os.Exit(0)
	}

	// Check for gamepad movement.
	state.hv = ebiten.GamepadAxis(0, 0)
	state.vv = ebiten.GamepadAxis(0, 1)
	state.hs = ebiten.GamepadAxis(0, 2)
	state.vs = ebiten.GamepadAxis(0, 3)

	globalConfig.lineLen = 1000
	if math.Abs(state.hs) > 0.20 || math.Abs(state.vs) > 0.20 {
		dx := state.x + globalConfig.lineLen*state.hs
		dy := state.y + globalConfig.lineLen*state.vs
		state.angle = math.Atan2(dy-state.y, dx-state.x) + math.Pi/2
	}

	acc := 15.0
	if math.Abs(state.hv) > 0.10 {
		state.x += state.hv*acc + state.hs
	}
	if math.Abs(state.vv) > 0.10 {
		state.y += state.vv*acc + state.vs
	}
	if math.Abs(state.hv) <= 0.10 && math.Abs(state.vv) <= 0.10 {
		state.accStarted = 0.0
	} else {
		state.accStarted += 0.1
	}

	state.msPassed = time.Now().Sub(state.timer).Milliseconds()
	if state.msPassed > 5*1000 {
		state.timer = time.Now()
	}
}
