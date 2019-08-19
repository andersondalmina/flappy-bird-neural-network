package components

import (
	"math"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

const PipeHeight = 40.0

// Pipe component
type Pipe struct {
	x          float64
	y          float64
	defeated   bool
	speed      float64
	spriteUp   *pixel.Sprite
	spriteDown *pixel.Sprite
}

// NewPipe creates a new pipe component
func NewPipe(x float64, y float64) *Pipe {
	return &Pipe{
		x:          x,
		y:          y,
		defeated:   false,
		speed:      0,
		spriteUp:   Sprites["pipeUp"],
		spriteDown: Sprites["pipeDown"],
	}
}

func (b *Pipe) GetX() float64 {
	return b.x
}

func (b *Pipe) GetY() float64 {
	return b.y
}

func (b *Pipe) IsDefeated() bool {
	return b.defeated
}

func (b *Pipe) Defeat() {
	b.defeated = true
}

func (p *Pipe) Draw(win *pixelgl.Window) error {
	// Draw down part of pipe
	p.spriteUp.Draw(win, pixel.IM.Moved(pixel.V(p.x, p.y-100)))
	for i := p.y - 100 - PipeHeight; i > 0; i -= PipeHeight {
		p.spriteDown.Draw(win, pixel.IM.Moved(pixel.V(p.x, i)))
	}

	// Draw up part of pipe
	p.spriteUp.Draw(win, pixel.IM.Rotated(pixel.V(0, 0), math.Pi).Moved(pixel.V(p.x, p.y+100)))
	for i := p.y + 100 + PipeHeight; i < WindowHeight+PipeHeight; i += PipeHeight {
		p.spriteDown.Draw(win, pixel.IM.Rotated(pixel.V(0, 0), math.Pi).Moved(pixel.V(p.x, i)))
	}

	p.x -= XSpeed * Delta

	return nil
}
