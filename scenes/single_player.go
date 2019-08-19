package scenes

import (
	"fmt"
	"math/rand"
	"strconv"

	"github.com/andersondalmina/flappy-bird-neural-network/components"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font/basicfont"
)

const floorWidth = 209
const floorHeight = 75

type singlePlayer struct {
	bird   *components.Bird
	pipes  []*components.Pipe
	status bool
}

// CreateSinglePlayerScene create a scene when a player play alone
func CreateSinglePlayerScene() Scene {
	s := singlePlayer{
		status: true,
	}

	s.bird = components.NewBird(components.BirdX, components.Sprites["bird10"])

	for i := 0.0; i < 4; i++ {
		s.pipes = append(s.pipes, components.NewPipe(components.WindowWidth+320*i, (components.WindowHeight-components.PipeHeight*2)-rand.Float64()*10*components.PipeHeight))
	}

	return &s
}

func (s *singlePlayer) Run(win *pixelgl.Window) Scene {
	drawBackground(win)

	if win.JustPressed(pixelgl.KeySpace) || win.JustPressed(pixelgl.KeyUp) {
		s.bird.Enable()
		s.bird.Jump()
	}

	s.bird.Update()
	s.bird.Draw(win)

	for _, pipe := range s.pipes {
		pipe.Draw(win)
	}

	s.checkPipes()
	s.drawPoints(win)
	drawFloor(win)

	if s.checkCrash() {
		return CreateGameOverScene(s.bird.GetPoints())
	}

	return s
}

func (s *singlePlayer) checkPipes() {
	for i, p := range s.pipes {
		if p.GetX() <= components.BirdX-50 && p.IsDefeated() == false {
			p.Defeat()
			s.bird.IncreasePoint()
		}

		if s.countFollowingPipes() < 4 {
			s.pipes = append(s.pipes, components.NewPipe(components.WindowWidth+p.GetX(), components.WindowHeight-components.WindowHeight*0.1-rand.Float64()*250))
		}

		if p.GetX() <= -50 {
			s.pipes = append(s.pipes[:i], s.pipes[i+1:]...)
		}
	}
}

func (s *singlePlayer) countFollowingPipes() int {
	n := 0
	for _, p := range s.pipes {
		if p.IsDefeated() == false {
			n++
		}
	}

	return n
}

func (s *singlePlayer) checkCrash() bool {
	b := s.bird

	if b.GetY() <= 80 {
		return true
	}

	for _, p := range s.pipes {
		if b.GetX() >= p.GetX()-50 && b.GetX() <= p.GetX()+50 {
			if b.GetY() <= (p.GetY()-55) || b.GetY() >= (p.GetY()+55) {
				return true
			}
		}
	}

	return false
}

func (s *singlePlayer) gameOver(win *pixelgl.Window) {
	s.status = false

	win.Clear(colornames.Black)

	basicAtlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)

	text := text.New(win.Bounds().Center(), basicAtlas)
	fmt.Fprintln(text, "Game Over")

	text.Draw(win, pixel.IM.Scaled(text.Bounds().Center(), 5))
}

func (s *singlePlayer) drawPoints(win *pixelgl.Window) {
	t := strconv.FormatInt(s.bird.GetPoints(), 10)
	p := components.CreateTextLine(t, colornames.White)

	var text []components.Text
	text = append(text, p)

	mat := pixel.IM.Scaled(win.Bounds().Center(), 5).Moved(pixel.V(-10, components.WindowHeight/3))
	components.WriteText(text, colorMenu, win, mat)
}
