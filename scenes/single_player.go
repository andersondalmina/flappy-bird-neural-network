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

var wallTime int

type singlePlayer struct {
	bird      *components.Bird
	obstacles []components.Obstacle
	status    bool
}

// CreateSinglePlayerScene create a scene when a player play alone
func CreateSinglePlayerScene() Scene {
	resetWallTime()

	s := singlePlayer{
		status: true,
	}

	return &s
}

func (s *singlePlayer) Load() Scene {
	s.bird = components.NewBird(components.BirdX, components.Sprites["bird10"])

	for i := 0.0; i < 4; i++ {
		s.obstacles = append(s.obstacles, components.NewPipe(components.WindowWidth+320*i, (components.WindowHeight-components.PipeHeight*2)-rand.Float64()*10*components.PipeHeight))
	}

	return s
}

func (s *singlePlayer) Run(win *pixelgl.Window) Scene {
	wallTime--

	drawBackground(win)

	if win.JustPressed(pixelgl.KeySpace) || win.JustPressed(pixelgl.KeyUp) {
		s.bird.Jump()
	}

	if win.JustPressed(pixelgl.KeyLeftControl) {
		s.bird.UseGhost()
	}

	go s.bird.Update()
	s.bird.Draw(win)

	for _, o := range s.obstacles {
		o.Draw(win)
	}

	go s.checkObstacles()
	s.drawPoints(win)
	drawFloor(win)

	if s.checkCrash() {
		return CreateGameOverScene(s.bird.Points).Load()
	}

	return s
}

func (s *singlePlayer) checkObstacles() {
	for i, o := range s.obstacles {
		if o.GetX() <= components.BirdX-50 && o.IsDefeated() == false {
			o.Defeat()
			s.bird.IncreasePoint()
		}

		if s.countFollowingObstacles() < 4 {
			if wallTime <= 0 {
				resetWallTime()
				s.obstacles = append(s.obstacles, components.NewWall(components.WindowWidth+o.GetX()))
				return
			}

			s.obstacles = append(s.obstacles, components.NewPipe(components.WindowWidth+o.GetX(), components.WindowHeight-components.WindowHeight*0.1-rand.Float64()*250))
		}

		if o.GetX() <= -o.GetWidth()/2 {
			s.obstacles = append(s.obstacles[:i], s.obstacles[i+1:]...)
		}
	}
}

func (s *singlePlayer) countFollowingObstacles() int {
	n := 0
	for _, o := range s.obstacles {
		if o.IsDefeated() == false {
			n++
		}
	}

	return n
}

func (s *singlePlayer) checkCrash() bool {
	b := s.bird

	if b.Y <= 80 || b.Y >= components.WindowHeight {
		return true
	}

	if b.Ghost {
		return false
	}

	for _, o := range s.obstacles {
		if o.CheckCrash(b.X, b.Y) {
			return true
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
	t := strconv.FormatInt(s.bird.Points, 10)
	p := components.CreateTextLine(t, colornames.White)

	var text []components.Text
	text = append(text, p)

	mat := pixel.IM.Scaled(win.Bounds().Center(), 5).Moved(pixel.V(-10, components.WindowHeight/3))
	components.WriteText(text, colorMenu, win, mat)
}

func resetWallTime() {
	wallTime = 1500
}
