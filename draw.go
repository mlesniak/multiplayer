package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"image/color"
	"math"
)

func draw(screen *ebiten.Image) {
	if state.pause {
		ebitenutil.DebugPrint(screen, "<PAUSED>")
	}

	pl0 := state.players[0]
	if math.Abs(state.hs) > 0.20 || math.Abs(state.vs) > 0.20 {
		dx := pl0.x + pl0.lineLength*state.hs
		dy := pl0.y + pl0.lineLength*state.vs
		ebitenutil.DrawLine(screen, pl0.x, pl0.y, dx, dy, color.RGBA{255, 255, 0, 255})
	}

	//ebitenutil.DrawRect(screen, 0, 0, float64(globalConfig.width), 40, color.RGBA{255, 255, 0, 255})
	//msg := fmt.Sprintf("--- %.2f ---", float64(state.msPassed)/1000.0)
	//b, _ := font.BoundString(arcadeFont, msg)
	//a := b.Max.X.Ceil()
	//text.Draw(screen, msg, arcadeFont, globalConfig.width/2-a/2, 30, color.Black)

	for _, obstacle := range state.obstacles {
		op := &ebiten.DrawImageOptions{}
		w, h := obstacleImage.Size()
		op.GeoM.Translate(-float64(w)/2, -float64(h)/2)
		//op.GeoM.Scale(0.3, 0.3)
		op.GeoM.Translate(obstacle.x, obstacle.y)
		if !obstacle.hit {
			screen.DrawImage(obstacleImage, op)
		} else {
			screen.DrawImage(obstacleHitImage, op)
		}
	}

	op := &ebiten.DrawImageOptions{}
	w, h := gopherImage.Size()
	op.GeoM.Translate(-float64(w)/2, -float64(h)/2)
	//op.GeoM.Rotate(state.angle)
	op.GeoM.Scale(0.3, 0.3)
	op.GeoM.Translate(state.players[0].x, state.players[0].y)
	screen.DrawImage(gopherImage, op)

	px, py := ebiten.CursorPosition()
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%v/%v", px, py), 0, 20)
	ebitenutil.DrawLine(screen, float64(px), 0, float64(px), float64(globalConfig.height), color.RGBA{40, 40, 40, 255})
	ebitenutil.DrawLine(screen, 0, float64(py), float64(globalConfig.width), float64(py), color.RGBA{40, 40, 40, 255})
}
