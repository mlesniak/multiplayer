// TODO Compute new line length
// TODO Refactor into multiple files under state/
package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"math"
	"os"
	"time"
)

type State struct {
	// Input
	hv, vv, hs, vs float64

	// Objects
	players   []Player
	obstacles []Obstacle

	// Obstacles

	// Timing
	timer    time.Time
	msPassed int64
	frames   int64
	pause    bool
}

var state = State{
	timer: time.Now(),
}

type Player struct {
	x, y  float64 // center
	angle float64
}

// Currently, only rectangular obstacles (and players)
type Obstacle struct {
	x, y          float64 // center
	width, height float64
	hit           bool
}

func init() {
	state.players = make([]Player, 2)
	pl0 := &state.players[0]
	pl0.x = 400
	pl0.y = 200
	pl0.angle = 0

	state.obstacles = make([]Obstacle, 1)
	w, h := obstacleImage.Size()
	state.obstacles = []Obstacle{
		{
			x:      600,
			y:      200,
			width:  float64(w),
			height: float64(h),
		},
	}
}

func (o *Obstacle) inside(x, y float64) bool {
	return x >= o.x-o.width/2 &&
		x <= o.x+o.width/2 &&
		y >= o.y-o.height/2 &&
		y <= o.y+o.height/2
}

func updateState() {
	// Check for key presses.
	if ebiten.IsKeyPressed(ebiten.KeyEscape) || ebiten.IsKeyPressed(ebiten.KeyQ) {
		os.Exit(0)
	}
	if ebiten.IsKeyPressed(ebiten.KeyP) {
		state.pause = !state.pause
		show = true
	}
	if state.pause {
		return
	}

	// Update internal state dependent on pause.
	state.frames++

	// Check for gamepad movement.
	state.hv = ebiten.GamepadAxis(0, 0)
	state.vv = ebiten.GamepadAxis(0, 1)
	state.hs = ebiten.GamepadAxis(0, 2)
	state.vs = ebiten.GamepadAxis(0, 3)

	// - Simulate gamepad input ----------------------------------------------------------------------
	state.hs = 1.0
	state.vs = math.Sin(0.01 * float64(state.frames))
	//state.vs = 0.0
	// -----------------------------------------------------------------------------------------------

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

	updateHit()
}

var show bool

func updateHit() {
	// Reset hit state
	for i, _ := range state.obstacles {
		state.obstacles[i].hit = false
	}

	// Check if any obstacle was hit by the laser beam by raycasting. CURRENTLY, pixel by pixel.
	p := state.players[0]
	dx := math.Sin(p.angle)
	dy := math.Cos(p.angle)

loop:
	for ll := 1.0; ll < globalConfig.lineLen; ll++ {
		tx := p.x + ll*dx
		ty := p.y + ll*dy
		// Check all objects.
		for i, _ := range state.obstacles {
			if state.obstacles[i].inside(tx, ty) {
				if show {
					fmt.Printf("%f		%f, %f\n", ll, tx, ty)
					show = false
				}

				state.obstacles[i].hit = true
				break loop
			}
		}
	}
}
