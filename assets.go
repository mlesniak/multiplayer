package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"log"
)

var gopherImage *ebiten.Image
var obstacleImage *ebiten.Image

func init() {
	var err error
	gopherImage, _, err = ebitenutil.NewImageFromFile("asset/red.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	obstacleImage, _, err = ebitenutil.NewImageFromFile("asset/blue.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
}
