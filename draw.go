package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/text"
	"golang.org/x/image/font"
	"image/color"
	"math"
)

func draw(screen *ebiten.Image) {
	// Move after check for draw.
	if math.Abs(state.hs) > 0.20 || math.Abs(state.vs) > 0.20 {
		dx := state.x + globalConfig.lineLen*state.hs
		dy := state.y + globalConfig.lineLen*state.vs
		ebitenutil.DrawLine(screen, state.x, state.y, dx, dy, color.RGBA{255, 255, 0, 255})
	}

	msg := fmt.Sprintf("--- %.2f ---", float64(state.msPassed)/1000.0)
	b, _ := font.BoundString(arcadeFont, msg)
	a := b.Max.X.Ceil()
	text.Draw(screen, msg, arcadeFont, globalConfig.width/2-a/2, 30, color.White)

	op := &ebiten.DrawImageOptions{}
	w, h := gopherImage.Size()
	op.GeoM.Translate(-float64(w)/2, -float64(h)/2)
	//op.GeoM.Rotate(state.angle)
	op.GeoM.Scale(0.1, 0.1)
	op.GeoM.Translate(state.x, state.y)
	screen.DrawImage(gopherImage, op)
}
