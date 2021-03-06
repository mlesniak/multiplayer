// TODO Refactor into multiple files under state/
// TODO Prevent right and bottom pixel-gap on collision detection
package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"math"
	"math/rand"
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
	x, y       float64 // center
	angle      float64
	lineLength float64

	width, height float64
}

// Currently, only rectangular obstacles (and players)
type Obstacle struct {
	x, y          float64 // center
	width, height float64
	hit           bool
}

func init() {
	//wp, hp := playerImage.Size()
	state.players = make([]Player, 2)
	pl0 := &state.players[0]
	pl0.x = 400
	pl0.y = 200
	pl0.width = 20
	pl0.height = 20
	pl0.angle = 0
	pl0.lineLength = globalConfig.lineLen

	rand.Seed(time.Now().Unix())
	numObstacles := 10
	state.obstacles = make([]Obstacle, numObstacles)
	w, h := obstacleImage.Size()
	for i := 0; i < numObstacles; i++ {
		obstacle := Obstacle{
			x:      float64(rand.Intn(globalConfig.width)),
			y:      float64(rand.Intn(globalConfig.height)),
			width:  float64(w),
			height: float64(h),
		}
		if obstacle.inside(pl0.x, pl0.y) {
			continue
		}
		state.obstacles = append(state.obstacles, obstacle)
	}
}

func (o *Obstacle) inside(x, y float64) bool {
	return x >= o.x-o.width/2 &&
		x <= o.x+o.width/2 &&
		y >= o.y-o.height/2 &&
		y <= o.y+o.height/2
}

func updateState() {
	if checkKeys() {
		return
	}

	updateInternalState()
	handleInput()

	updateShootingAngle()
	updatePlayerPosition()
	updateHit()
}

func updatePlayerPosition() {
	acc := 15.0
	steps := int(acc) * 10
	dx := (state.hv*acc + state.hs) / float64(steps)
	dy := (state.vv*acc + state.vs) / float64(steps)

	collision := false
	i := 0
loop:
	for ; i < steps; i++ {
		vx := state.players[0].x
		vy := state.players[0].y
		if math.Abs(state.hv) > 0.10 {
			vx += dx
		}
		if math.Abs(state.vv) > 0.10 {
			vy += dy
		}

		// Check if one of the corners collides with one of the obstacles. If yes, reset to previous position.
		c1x := vx - state.players[0].width/2
		c1y := vy - state.players[0].height/2
		c2x := vx + state.players[0].width/2
		c2y := vy + state.players[0].height/2
		c3x := vx + state.players[0].width/2
		c3y := vy - state.players[0].height/2
		c4x := vx - state.players[0].width/2
		c4y := vy + state.players[0].height/2
		for i, _ := range state.obstacles {
			if state.obstacles[i].inside(c1x, c1y) {
				collision = true
				break loop
			}
			if state.obstacles[i].inside(c2x, c2y) {
				collision = true
				break loop
			}
			if state.obstacles[i].inside(c3x, c3y) {
				collision = true
				break loop
			}
			if state.obstacles[i].inside(c4x, c4y) {
				collision = true
				break loop
			}
		}

		// No collision. Update
		state.players[0].x = vx
		state.players[0].y = vy
	}

	if collision {

	}
}

func updateInternalState() {
	state.frames++

	state.msPassed = time.Now().Sub(state.timer).Milliseconds()
	if int(state.msPassed) > globalConfig.roundDuration*1000 {
		state.timer = time.Now()
	}
}

func updateShootingAngle() {
	globalConfig.lineLen = 1000
	if math.Abs(state.hs) > 0.20 || math.Abs(state.vs) > 0.20 {
		dx := state.players[0].x + globalConfig.lineLen*state.hs
		dy := state.players[0].y + globalConfig.lineLen*state.vs
		state.players[0].angle = math.Atan2(dy-state.players[0].y, dx-state.players[0].x) //+ math.Pi/2
	}
}

func handleInput() {
	// Check for gamepad movement.
	// Commented out for debugging with keyboard.r
	state.hv = ebiten.GamepadAxis(0, 0)
	state.vv = ebiten.GamepadAxis(0, 1)
	state.hs = ebiten.GamepadAxis(0, 2)
	state.vs = ebiten.GamepadAxis(0, 3)

	// - Simulate gamepad input ----------------------------------------------------------------------
	//state.hs = 1.0
	//state.hs = math.Cos(0.05 * float64(state.frames))
	//state.vs = math.Sin(0.05 * float64(state.frames))
	//state.vs = 0.0
	// -----------------------------------------------------------------------------------------------
}

func checkKeys() bool {
	// Debug keypress for player 0.
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		state.vv = -0.5
	} else if ebiten.IsKeyPressed(ebiten.KeyS) {
		state.vv = 0.5
	} else {
		state.vv = 0.0
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		state.hv = -0.5
	} else if ebiten.IsKeyPressed(ebiten.KeyD) {
		state.hv = 0.5
	} else {
		state.hv = 0.0
	}

	// Check for key presses.
	if ebiten.IsKeyPressed(ebiten.KeyEscape) || ebiten.IsKeyPressed(ebiten.KeyQ) {
		os.Exit(0)
	}
	if ebiten.IsKeyPressed(ebiten.KeyP) {
		state.pause = !state.pause
		show = true
	}
	if state.pause {
		return true
	}
	return false
}

var show = false

func updateHit() {
	// Reset hit state
	for i, _ := range state.obstacles {
		state.obstacles[i].hit = false
	}

	// Check if any obstacle was hit by the laser beam by raycasting. CURRENTLY, pixel by pixel to find the exact collision point.
	p := state.players[0]
	state.players[0].lineLength = globalConfig.lineLen
loop:
	for ll := 1.0; ll < globalConfig.lineLen; ll += 1.00 {
		tx := p.x + ll*math.Cos(p.angle)
		ty := p.y + ll*math.Sin(p.angle)
		// Check all objects.
		for i, _ := range state.obstacles {
			if state.obstacles[i].inside(tx, ty) {
				if show {
					fmt.Printf("tx=%v ty=%v\n", tx, ty)
				}
				state.obstacles[i].hit = true
				state.players[0].lineLength = ll
				break loop
			}
		}
	}

	if show {
		show = false
	}
}
