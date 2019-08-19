package scenes

import (
	"strconv"

	"github.com/andersondalmina/flappy-bird-neural-network/components"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

type gameOverScene struct {
	points int64
}

// CreateGameOverScene create a scene to game over
func CreateGameOverScene(p int64) Scene {
	s := gameOverScene{
		points: p,
	}

	return &s
}

func (s *gameOverScene) Run(win *pixelgl.Window) Scene {
	win.Clear(colornames.Skyblue)

	drawBackground(win)
	drawFloor(win)

	var text []components.Text
	p := components.CreateTextLine("Game Over", colornames.White)
	text = append(text, p)

	t := strconv.FormatInt(s.points, 10)
	p = components.CreateTextLine(t, colornames.White)
	text = append(text, p)

	mat := pixel.IM.Scaled(win.Bounds().Center(), 5).Moved(pixel.V(0, 50))
	components.WriteText(text, colorMenu, win, mat)

	if win.JustPressed(pixelgl.KeyEnter) {
		return CreateMenuScene()
	}

	return s
}
