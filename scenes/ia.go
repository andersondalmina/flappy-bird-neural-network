package scenes

import (
	"math"
	"math/rand"
	"strconv"

	"github.com/andersondalmina/flappy-bird-neural-network/components"
	"github.com/andersondalmina/flappy-bird-neural-network/neuralnetwork"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

var bestWeights = make([]float64, 2)
var weights = make([]float64, 2)
var bestPoints int64

type ia struct {
	pop    *neuralnetwork.Population
	pipes  []*components.Pipe
	status bool
}

// CreateIAScene create a scene when a machine plays
func CreateIAScene(gn int64) Scene {
	s := ia{
		status: true,
	}

	s.pop = neuralnetwork.CreateNewPopulation(gn)

	var n int64
	var t string
	var ind *neuralnetwork.Individual
	for i := 0; i < 10; i++ {
		n = rand.Int63n(4) + 1
		t = strconv.FormatInt(n, 10)
		ind = neuralnetwork.NewIndividual(components.NewBird(components.BirdX-rand.Float64()*200, components.Sprites["bird1"+t]))

		if gn > 1 {
			weights = bestWeights
			weights[0] += (rand.Float64()*2 - 1) * 150
			weights[1] += (rand.Float64()*2 - 1) * 150
			ind.Neuron().SetWeights(weights)
		}

		s.pop.AddIndividual(ind)
	}

	for i := 0.0; i < 4; i++ {
		s.pipes = append(s.pipes, components.NewPipe(components.WindowWidth+320*i, (components.WindowHeight-components.PipeHeight*2)-rand.Float64()*10*components.PipeHeight))
	}

	return &s
}

func (s *ia) Run(win *pixelgl.Window) Scene {
	drawBackground(win)

	if win.JustPressed(pixelgl.KeySpace) || win.JustPressed(pixelgl.KeyUp) {
		s.pop.GetIndividuals()[0].Bird().Jump()
	}

	var bInputs []float64
	var np *components.Pipe
	for _, b := range s.pop.GetIndividuals() {
		b.Bird().Update()
		b.Bird().Draw(win)

		if b.Bird().IsDeath() == true {
			continue
		}

		bInputs = make([]float64, 2)
		np = s.getBirdNextPipe(b.Bird())
		bInputs[0] = np.GetX() - b.Bird().GetX()
		bInputs[1] = np.GetY() - b.Bird().GetY()
		b.SetInputs(bInputs)

		if math.Tanh(b.Neuron().Predict(bInputs)) > 0 {
			b.Bird().Jump()
		}

		b.Bird().IncreasePoint()
		if s.checkCrash(b.Bird()) {
			b.Bird().Death()
		}
	}

	for _, pipe := range s.pipes {
		pipe.Draw(win)
	}

	s.checkPipes()
	drawFloor(win)
	s.drawInterface(win)

	if s.checkBirdsAlive() == false {
		best := s.pop.GetIndividuals()[0]
		for _, b := range s.pop.GetIndividuals() {
			if b.Bird().GetPoints() > best.Bird().GetPoints() {
				best = b
			}
		}

		if best.Bird().GetPoints() >= bestPoints {
			bestWeights = best.Neuron().Weights()
			bestPoints = best.Bird().GetPoints()
		}
		return CreateIAScene(s.pop.Generation() + 1)
	}

	return s
}

func (s *ia) checkPipes() {
	for _, b := range s.pop.GetIndividuals() {
		for i, p := range s.pipes {
			if p.GetX() <= b.Bird().GetX()-50 && p.IsDefeated() == false {
				p.Defeat()
			}

			if s.countFollowingPipes() < 4 {
				s.pipes = append(s.pipes, components.NewPipe(components.WindowWidth+p.GetX(), components.WindowHeight-components.WindowHeight*0.1-rand.Float64()*250))
			}

			if p.GetX() <= -50 {
				s.pipes = append(s.pipes[:i], s.pipes[i+1:]...)
			}
		}
	}
}

func (s *ia) countFollowingPipes() int {
	n := 0
	for _, p := range s.pipes {
		if p.IsDefeated() == false {
			n++
		}
	}

	return n
}

func (s *ia) checkCrash(b *components.Bird) bool {
	if b.GetY() <= 80 || b.GetY() >= components.WindowHeight {
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

func (s *ia) getBirdNextPipe(b *components.Bird) *components.Pipe {
	for _, p := range s.pipes {
		if p.GetX()+50 > b.GetX() {
			return p
		}
	}

	return s.pipes[0]
}

func (s *ia) checkBirdsAlive() bool {
	for _, b := range s.pop.GetIndividuals() {
		if b.Bird().IsDeath() == false {
			return true
		}
	}

	return false
}

func (s *ia) drawInterface(win *pixelgl.Window) {
	gn := strconv.FormatInt(s.pop.Generation(), 10)
	p := components.CreateTextLine("Gen "+gn, colornames.White)

	var text []components.Text
	text = append(text, p)

	mat := pixel.IM.Scaled(win.Bounds().Center(), 5).Moved(pixel.V(-10, components.WindowHeight/3))
	components.WriteText(text, colorMenu, win, mat)
}
