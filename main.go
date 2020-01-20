package main

import (
	"github.com/hajimehoshi/ebiten"
	"log"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	if err := ebiten.Run(update, globalConfig.width, globalConfig.height, 1, "Game"); err != nil {
		log.Fatal(err)
	}
}

func update(screen *ebiten.Image) error {
	updateState()
	if ebiten.IsDrawingSkipped() {
		return nil
	}
	draw(screen)
	return nil
}
