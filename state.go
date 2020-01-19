package main

import (
	"github.com/hajimehoshi/ebiten"
	"math"
	"os"
	"time"
)

type State struct {
	// Input
	hv, vv, hs, vs float64

	// Players
	players []Player

	// Timing
	timer    time.Time
	msPassed int64
}

var state = State{
	timer: time.Now(),
}

func init() {
	state.players = make([]Player, 2)
	pl0 := &state.players[0]
	pl0.x = 400
	pl0.y = 200
	pl0.angle = 0
}

type Player struct {
	x, y  float64
	angle float64
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
		dx := state.players[0].x + globalConfig.lineLen*state.hs
		dy := state.players[0].y + globalConfig.lineLen*state.vs
		state.players[0].angle = math.Atan2(dy-state.players[0].y, dx-state.players[0].x) + math.Pi/2
	}

	acc := 15.0
	if math.Abs(state.hv) > 0.10 {
		state.players[0].x += state.hv*acc + state.hs
	}
	if math.Abs(state.vv) > 0.10 {
		state.players[0].y += state.vv*acc + state.vs
	}

	state.msPassed = time.Now().Sub(state.timer).Milliseconds()
	if int(state.msPassed) > globalConfig.roundDuration*1000 {
		state.timer = time.Now()
	}
}
