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

// GetWidth return the wall width
func (w *Wall) GetWidth() float64 {
	return WallWidth
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

	w.X -= GameXSpeed * Delta
}

// Update do nothing
func (w *Wall) Update() {

}

// CheckCrash check if a position crash on the wall
func (w *Wall) CheckCrash(x, y float64) bool {
	if x >= w.X-WallWidth/2 && x <= w.X+WallWidth/2 {
		return true
	}

	return false
}
