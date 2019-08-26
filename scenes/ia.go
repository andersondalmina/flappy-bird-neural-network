package scenes

import (
	"math/rand"
	"strconv"

	"github.com/andersondalmina/flappy-bird-neural-network/components"
	"github.com/andersondalmina/flappy-bird-neural-network/neuralnetwork"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

var bestWeights = make([][][]float64, 10)
var weights = make([][][]float64, 10)
var bestPoints int64

type ia struct {
	pop       *components.Population
	obstacles []components.Obstacle
	status    bool
}

// CreateIAScene create a scene when a machine plays
func CreateIAScene(gn int64) Scene {
	s := ia{
		status: true,
	}

	s.pop = components.CreateNewPopulation(gn)

	var n int64
	var t string
	var ind *components.Individual
	for i := 0; i < 100; i++ {
		n = rand.Int63n(4) + 1
		t = strconv.FormatInt(n, 10)
		neural := neuralnetwork.NewNeuralNetwork(neuralnetwork.Config{
			Inputs: 2,
			Layers: []int64{3, 1},
		})
		ind = components.NewIndividual(components.NewBird(components.BirdX-rand.Float64()*200, components.Sprites["bird1"+t]), neural)

		if gn > 1 {
			weights = ind.Neural().Weights()
			for z := range bestWeights {
				for zz := range bestWeights[z] {
					for zzz := range bestWeights[z][zz] {
						weights[z][zz][zzz] = bestWeights[z][zz][zzz]
						if rand.Intn(4) == 0 {
							weights[z][zz][zzz] += (rand.Float64()*2 - 1) * 100
						}
					}
				}
			}

			ind.Neural().SetWeights(weights)
		}

		s.pop.AddIndividual(ind)
	}

	for i := 0.0; i < 4; i++ {
		s.obstacles = append(s.obstacles, components.NewPipe(components.WindowWidth+320*i, (components.WindowHeight-components.PipeHeight*2)-rand.Float64()*10*components.PipeHeight))
	}

	return &s
}

func (s *ia) Run(win *pixelgl.Window) Scene {
	drawBackground(win)

	if win.JustPressed(pixelgl.KeyEnter) {
		for _, b := range s.pop.GetIndividuals() {
			b.Bird().Death()
		}
	}

	var bInputs []float64
	var np components.Obstacle
	for _, b := range s.pop.GetIndividuals() {
		go b.Bird().Update()
		b.Bird().Draw(win)

		if b.Bird().IsDeath() == true {
			continue
		}

		bInputs = make([]float64, 2)
		np = s.getBirdNextPipe(b.Bird())
		bInputs[0] = np.GetX() - b.Bird().GetX()
		bInputs[1] = np.GetY() - components.PipeHeight - b.Bird().GetY()
		b.SetInputs(bInputs)

		if Max(b.Neural().Predict(bInputs)[0], 0) > 0 {
			b.Bird().Jump()
		}

		b.Bird().IncreasePoint()
		if s.checkCrash(b.Bird()) {
			b.Bird().Death()
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
			if b.Bird().GetPoints() > best.Bird().GetPoints() {
				best = b
			}
		}

		if best.Bird().GetPoints() > bestPoints {
			bestWeights = best.Neural().Weights()
			bestPoints = best.Bird().GetPoints()
		}

		return CreateIAScene(s.pop.Generation() + 1)
	}

	return s
}

func (s *ia) checkPipes() {
	for _, b := range s.pop.GetIndividuals() {
		for i, p := range s.obstacles {
			if p.GetX() <= b.Bird().GetX()-50 && p.IsDefeated() == false {
				p.Defeat()
			}

			if s.countFollowingPipes() < 4 {
				s.obstacles = append(s.obstacles, components.NewPipe(components.WindowWidth+p.GetX(), components.WindowHeight-components.WindowHeight*0.1-rand.Float64()*250))
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
	if b.GetY() <= 80 || b.GetY() >= components.WindowHeight {
		return true
	}

	for _, o := range s.obstacles {
		if o.CheckCrash(*b) {
			return true
		}
	}

	return false
}

func (s *ia) getBirdNextPipe(b *components.Bird) components.Obstacle {
	for _, o := range s.obstacles {
		if o.GetX()+components.PipeWidth/2 > b.GetX() {
			return o
		}
	}

	return s.obstacles[0]
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

// Max returns the larger of x or y.
func Max(x, y float64) float64 {
	if x < y {
		return y
	}

	return x
}
