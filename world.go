package main

import (
	"time"
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

	timer time.Time
}

var world = World{
	x:     400,
	y:     300,
	acc:   0,
	timer: time.Now(),
}
