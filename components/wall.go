package components

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

// WallWidth is the width of the wall sprite
const WallWidth = 160

// Wall type of obstacle
type Wall struct {
	X        float64
	Y        float64
	defeated bool
}

// NewWall creates a new wall component
func NewWall(x float64) *Wall {
	return &Wall{
		X: x,
		Y: WindowHeight / 2,
	}
}

// GetType return the type of obstacle
func (w *Wall) GetType() float64 {
	return 2
}

// GetX return X position
func (w *Wall) GetX() float64 {
	return w.X
}

// GetY return Y position
func (w *Wall) GetY() float64 {
	return w.Y
}

// IsDefeated return if the wall is defeat
func (w *Wall) IsDefeated() bool {
	return w.defeated
}

// Defeat mask this wall as defeat
func (w *Wall) Defeat() {
	w.defeated = true
}

// Draw the wall on window
func (w *Wall) Draw(win *pixelgl.Window) {
	x := WindowHeight / 160

	for i := 0; i <= x; i++ {
		Sprites["wall"].Draw(win, pixel.IM.Moved(pixel.V(w.X, float64(i)*160)))
	}

	w.X -= XSpeed * Delta
}

// Update do nothing
func (w *Wall) Update() {

}

// CheckCrash check if a bird crash on the wall
func (w *Wall) CheckCrash(b Bird) bool {
	if b.GetX() >= w.X-WallWidth/2 && b.GetX() <= w.X+WallWidth/2 {
		return true
	}

	return false
}
