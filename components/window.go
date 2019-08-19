package components

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

// WindowWidth width of the window to be draw
const WindowWidth = 1280

// WindowHeight height of the window to be draw
const WindowHeight = 720

// Delta is to control FPS
var Delta float64

// XSpeed is the horizontal speed
var XSpeed = 180.0

// CreateWindow creates a window
func CreateWindow(t string) (*pixelgl.Window, error) {
	wConfig := pixelgl.WindowConfig{
		Title:  t,
		Bounds: pixel.R(0, 0, WindowWidth, WindowHeight),
		VSync:  true,
	}

	return pixelgl.NewWindow(wConfig)
}
