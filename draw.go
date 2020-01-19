package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"image/color"
	"math"
)

func draw(screen *ebiten.Image) {
	if math.Abs(state.hs) > 0.20 || math.Abs(state.vs) > 0.20 {
		dx := state.players[0].x + globalConfig.lineLen*state.hs
		dy := state.players[0].y + globalConfig.lineLen*state.vs
		ebitenutil.DrawLine(screen, state.players[0].x, state.players[0].y, dx, dy, color.RGBA{255, 255, 0, 255})
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
		op.GeoM.Scale(0.3, 0.3)
		op.GeoM.Translate(obstacle.x, obstacle.y)
		screen.DrawImage(obstacleImage, op)
	}

	op := &ebiten.DrawImageOptions{}
	w, h := gopherImage.Size()
	op.GeoM.Translate(-float64(w)/2, -float64(h)/2)
	//op.GeoM.Rotate(state.angle)
	op.GeoM.Scale(0.3, 0.3)
	op.GeoM.Translate(state.players[0].x, state.players[0].y)
	screen.DrawImage(gopherImage, op)
}
