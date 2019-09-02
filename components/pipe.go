package components

import (
	"math"
	"math/rand"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

// PipeHeight is the width of the pipe sprite
const PipeHeight = 40.0

// PipeWidth is the width of the pipe sprite
const PipeWidth = 86.0

// Obstacle component interface
type Obstacle interface {
	GetX() float64
	GetY() float64
	Draw(win *pixelgl.Window)
	Update()
	IsDefeated() bool
	Defeat()
	CheckCrash(x, y float64) bool
	GetType() float64
	GetWidth() float64
}

// Pipe type of obstacle
type Pipe struct {
	x          float64
	y          float64
	defeated   bool
	speed      float64
	spriteUp   *pixel.Sprite
	spriteDown *pixel.Sprite
	direction  int
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
		direction:  rand.Intn(1),
	}
}

// GetWidth return the pipe width
func (p *Pipe) GetWidth() float64 {
	return PipeWidth
}

// GetType return the type of obstacle
func (p *Pipe) GetType() float64 {
	return 1
}

// GetX return X position
func (p *Pipe) GetX() float64 {
	return p.x
}

// GetY return Y position
func (p *Pipe) GetY() float64 {
	return p.y
}

// IsDefeated return if the wall is defeat
func (p *Pipe) IsDefeated() bool {
	return p.defeated
}

// Defeat mask this wall as defeat
func (p *Pipe) Defeat() {
	p.defeated = true
}

// Draw the wall on window
func (p *Pipe) Draw(win *pixelgl.Window) {
	// Draw down part of pipe
	p.spriteUp.Draw(win, pixel.IM.Moved(pixel.V(p.x, p.y-100)))
	for i := p.y - 100 - PipeHeight; i > 0; i -= PipeHeight - 1 {
		p.spriteDown.Draw(win, pixel.IM.Moved(pixel.V(p.x, i)))
	}

	// Draw up part of pipe
	p.spriteUp.Draw(win, pixel.IM.Rotated(pixel.V(0, 0), math.Pi).Moved(pixel.V(p.x, p.y+100)))
	for i := p.y + 100 + PipeHeight; i < WindowHeight+PipeHeight; i += PipeHeight - 1 {
		p.spriteDown.Draw(win, pixel.IM.Rotated(pixel.V(0, 0), math.Pi).Moved(pixel.V(p.x, i)))
	}

	p.x -= GameXSpeed * Delta
}

// Update the pipe
func (p *Pipe) Update() {
	if p.direction == 1 {
		p.y += 0.5
	} else {
		p.y -= 0.5
	}

	if p.y > WindowHeight-200 {
		p.direction = 0
	} else if p.y < 200 {
		p.direction = 1
	}
}

// CheckCrash check if a position crash on pipe
func (p *Pipe) CheckCrash(x, y float64) bool {
	if x >= p.GetX()-PipeWidth/2 && x <= p.GetX()+PipeWidth/2 {
		if y <= (p.GetY()-50) || y >= (p.GetY()+50) {
			return true
		}
	}

	return false
}
