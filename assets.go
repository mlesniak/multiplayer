package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"log"
)

var playerImage *ebiten.Image
var obstacleImage *ebiten.Image
var obstacleHitImage *ebiten.Image

func init() {
	var err error
	playerImage, _, err = ebitenutil.NewImageFromFile("asset/red.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	obstacleImage, _, err = ebitenutil.NewImageFromFile("asset/blue.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	obstacleHitImage, _, err = ebitenutil.NewImageFromFile("asset/blue-hit.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
}
