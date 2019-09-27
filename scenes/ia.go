package scenes

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/andersondalmina/flappy-bird-neural-network/components"
	"github.com/andersondalmina/flappy-bird-neural-network/neuralnetwork"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

// IndNumber number of individuals of a population
const IndNumber = 200

var bestWeights = make([][][]float64, 10)
var weights = make([][][]float64, 10)
var bestPoints int64
var datafile string

type ia struct {
	generation int64
	pop        *components.Population
	obstacles  []components.Obstacle
}

// CreateIAScene create a scene when a machine plays
func CreateIAScene(gn int64) Scene {
	return &ia{
		generation: gn,
	}
}

func (s *ia) Load() Scene {
	datafile = "data/neuraldump_ia.json"

	rand.Seed(time.Now().UTC().UnixNano())

	s.pop = components.CreateNewPopulation(s.generation)
	s.obstacles = make([]components.Obstacle, 4)

	var n int64
	var t string
	var ind *components.Individual
	for i := 0; i < IndNumber; i++ {
		n = rand.Int63n(4) + 1
		t = strconv.FormatInt(n, 10)
		neural := neuralnetwork.NewNeuralNetwork(neuralnetwork.Config{
			Inputs: 2,
			Layers: []int64{10, 20, 1},
		})
		ind = components.NewIndividual(components.NewBird(components.BirdX-rand.Float64()*200, components.Sprites["bird1"+t]), neural)

		if s.generation > 1 {
			ind.Neural().UpdateWeights(bestWeights)
		}

		s.pop.AddIndividual(ind)
	}

	for i := 0; i < 4; i++ {
		s.obstacles[i] = components.NewPipe(components.WindowWidth+320*float64(i), (components.WindowHeight-components.PipeHeight*2)-rand.Float64()*10*components.PipeHeight)
	}

	return s
}

func (s *ia) Run(win *pixelgl.Window) Scene {
	drawBackground(win)

	pop := s.pop

	if win.JustPressed(pixelgl.KeyEnter) {
		for _, b := range pop.GetIndividuals() {
			b.Bird().Kill()
		}
	}

	var bInputs []float64
	var np components.Obstacle
	for i, ind := range pop.GetIndividuals() {
		bird := ind.Bird()
		go bird.Update()
		bird.Draw(win)

		if bird.Dead == true {
			if bird.X < 0 && len(pop.GetIndividuals()) > 1 && i > len(pop.GetIndividuals()) {
				pop.RemoveIndividual(i)
			}

			continue
		}

		bInputs = make([]float64, 2)
		np = s.getBirdNextPipe(bird)
		bInputs[0] = np.GetX() - bird.X
		bInputs[1] = np.GetY() - components.PipeHeight - bird.Y

		ind.SetInputs(bInputs)

		result := ind.Neural().Predict(bInputs)
		if result[0] > 0 {
			bird.Jump()
		}

		bird.IncreasePoint()
		if s.checkCrash(bird) {
			bird.Kill()
		}
	}

	for _, o := range s.obstacles {
		o.Draw(win)
	}

	go s.checkPipes()
	drawFloor(win)
	s.drawInterface(win)

	if s.checkBirdsAlive() == false {
		best := s.pop.GetIndividuals()[0]
		for _, b := range s.pop.GetIndividuals() {
			if b.Bird().Points > best.Bird().Points {
				best = b
			}
		}

		if best.Bird().Points > bestPoints {
			bestWeights = best.Neural().Weights()
			bestPoints = best.Bird().Points
		}

		return CreateIAScene(s.pop.Generation() + 1).Load()
	}

	return s
}

func (s *ia) checkPipes() {
	for _, b := range s.pop.GetIndividuals() {
		for i, p := range s.obstacles {
			if p.GetX() <= b.Bird().X-50 && p.IsDefeated() == false {
				p.Defeat()
			}

			if s.countFollowingPipes() < 4 {
				s.obstacles = append(s.obstacles, components.NewPipe(components.WindowWidth+p.GetX(), (components.WindowHeight-components.PipeHeight*2)-rand.Float64()*10*components.PipeHeight))
			}

			if p.GetX() <= -50 {
				s.obstacles = append(s.obstacles[:i], s.obstacles[i+1:]...)
			}
		}
	}
}

func (s *ia) countFollowingPipes() int {
	n := 0
	for _, o := range s.obstacles {
		if o.IsDefeated() == false {
			n++
		}
	}

	return n
}

func (s *ia) checkCrash(b *components.Bird) bool {
	if b.Y <= 80 || b.Y >= components.WindowHeight {
		return true
	}

	for _, o := range s.obstacles {
		if o.CheckCrash(b.X, b.Y) {
			return true
		}
	}

	return false
}

func (s *ia) getBirdNextPipe(b *components.Bird) components.Obstacle {
	for _, o := range s.obstacles {
		if o.GetX()+components.PipeWidth/2 > b.X {
			return o
		}
	}

	return s.obstacles[0]
}

func (s *ia) checkBirdsAlive() bool {
	for _, b := range s.pop.GetIndividuals() {
		if b.Bird().Dead == false {
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

// CheckIndividualsDead check if all individuals are dead
func CheckIndividualsDead(pop []*components.Individual) bool {
	for _, b := range pop {
		if b.Bird().Dead == false {
			return false
		}
	}

	return true
}
