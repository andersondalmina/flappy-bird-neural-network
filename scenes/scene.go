package scenes

import (
	"github.com/andersondalmina/flappy-bird-neural-network/components"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

// Scene is a interface for scenes
type Scene interface {
	Run(win *pixelgl.Window) Scene
}

func drawBackground(win *pixelgl.Window) {
	components.Sprites["background"].Draw(win, pixel.IM.Moved(pixel.V(win.Bounds().Center().X, win.Bounds().Center().Y+60)))
}

func drawFloor(win *pixelgl.Window) {
	for i := 0.0; i <= 13; i++ {
		components.Sprites["floor"].Draw(win, pixel.IM.Moved(pixel.V(floorWidth/2+floorWidth*i, floorHeight/2)))
	}
}
