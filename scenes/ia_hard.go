package scenes

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/andersondalmina/flappy-bird-neural-network/components"
	"github.com/andersondalmina/flappy-bird-neural-network/neuralnetwork"
	"github.com/andersondalmina/flappy-bird-neural-network/persist"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

type iaHard struct {
	generation int64
	pop        *components.Population
	obstacles  []components.Obstacle
}

// CreateIAHardScene create a scene when a machine plays
func CreateIAHardScene(gn int64) Scene {
	return &iaHard{
		generation: gn,
	}
}

func (s *iaHard) Load() Scene {
	datafile = "neuraldump_iahard.json"

	var data [][][]float64
	var err error

	rand.Seed(time.Now().UTC().UnixNano())

	resetWallTime()

	s.pop = components.CreateNewPopulation(s.generation)
	s.obstacles = make([]components.Obstacle, 4)

	err = persist.CheckFileExists(datafile)
	if err == nil {
		err = persist.Load(datafile, &data)
		if err != nil {
			panic(err)
		}
	}

	var n int64
	var t string
	var ind *components.Individual
	for i := 0; i < IndNumber; i++ {
		n = rand.Int63n(4) + 1
		t = strconv.FormatInt(n, 10)
		neural := neuralnetwork.NewNeuralNetwork(neuralnetwork.Config{
			Inputs: 3,
			Layers: []int64{10, 20, 2},
		})
		ind = components.NewIndividual(components.NewBird(components.BirdX-rand.Float64()*200, components.Sprites["bird1"+t]), neural)

		if s.generation == 1 && len(data) > 0 {
			ind.Neural().UpdateWeights(data)

		} else if s.generation > 1 {
			ind.Neural().UpdateWeights(ind.Neural().Weights())
		}

		s.pop.AddIndividual(ind)
	}

	for i := 0; i < 3; i++ {
		s.obstacles[i] = components.NewPipe(components.WindowWidth+320*float64(i), (components.WindowHeight-components.PipeHeight*2)-rand.Float64()*10*components.PipeHeight)
	}

	s.obstacles[3] = components.NewWall(components.WindowWidth + 320*3)

	return s
}

func (s *iaHard) Run(win *pixelgl.Window) Scene {
	wallTime--

	drawBackground(win)

	pop := s.pop

	if win.JustPressed(pixelgl.KeyEnter) {
		for _, b := range pop.GetIndividuals() {
			b.Bird().Kill()
		}
	}

	var bInputs []float64
	var np components.Obstacle
	var result []float64
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

		bInputs = make([]float64, 3)
		np = s.getBirdNextObstacle(bird)
		bInputs[0] = np.GetX() - bird.X
		bInputs[1] = np.GetY() - components.PipeHeight - bird.Y
		bInputs[2] = np.GetType()

		ind.SetInputs(bInputs)

		result = ind.Neural().Predict(bInputs)

		if result[0] > 0 {
			bird.Jump()
		}

		if result[1] > 0 {
			bird.UseGhost()
		}

		bird.IncreasePoint()
		if s.checkCrash(bird) {
			bird.Kill()
		}
	}

	for _, o := range s.obstacles {
		go o.Update()
		o.Draw(win)
	}

	go s.checkPipes()
	drawFloor(win)
	s.drawInterface(win)

	if CheckIndividualsDead(s.pop.GetIndividuals()) {
		best := s.pop.GetIndividuals()[0]
		for _, b := range s.pop.GetIndividuals() {
			if b.Bird().Points > best.Bird().Points {
				best = b
			}
		}

		bestWeights = best.Neural().Weights()
		bestPoints = best.Bird().Points

		err := best.Neural().Dump(datafile)
		if err != nil {
			panic(err)
		}

		return CreateIAHardScene(s.pop.Generation() + 1).Load()
	}

	return s
}

func (s *iaHard) checkPipes() {
	for _, b := range s.pop.GetIndividuals() {
		for i, o := range s.obstacles {
			if o.GetX() <= b.Bird().X-50 && o.IsDefeated() == false {
				o.Defeat()
			}

			if s.countFollowingPipes() < 4 {
				if wallTime <= 0 {
					resetWallTime()
					s.obstacles = append(s.obstacles, components.NewWall(components.WindowWidth+o.GetX()))
					return
				}

				s.obstacles = append(s.obstacles, components.NewPipe(components.WindowWidth+o.GetX(), components.WindowHeight-components.WindowHeight*0.1-rand.Float64()*250))
			}

			if o.GetX() <= -50 {
				s.obstacles = append(s.obstacles[:i], s.obstacles[i+1:]...)
			}
		}
	}
}

func (s *iaHard) countFollowingPipes() int {
	n := 0
	for _, p := range s.obstacles {
		if p.IsDefeated() == false {
			n++
		}
	}

	return n
}

func (s *iaHard) checkCrash(b *components.Bird) bool {
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

func (s *iaHard) getBirdNextObstacle(b *components.Bird) components.Obstacle {
	for _, o := range s.obstacles {
		if o.GetX()+o.GetWidth()/2 > b.X {
			return o
		}
	}

	return s.obstacles[0]
}

func (s *iaHard) drawInterface(win *pixelgl.Window) {
	gn := strconv.FormatInt(s.pop.Generation(), 10)
	p := components.CreateTextLine("Gen "+gn, colornames.White)

	var text []components.Text
	text = append(text, p)

	mat := pixel.IM.Scaled(win.Bounds().Center(), 5).Moved(pixel.V(-10, components.WindowHeight/3))
	components.WriteText(text, colorMenu, win, mat)
}
