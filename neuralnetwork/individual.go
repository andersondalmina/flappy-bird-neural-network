package neuralnetwork

import (
	"math/rand"
	"time"

	"github.com/andersondalmina/flappy-bird-neural-network/components"
)

// Individual is an interface
type Individual struct {
	bird   *components.Bird
	neuron *Neuron
	inputs []float64
}

// NewIndividual teste
func NewIndividual(b *components.Bird) *Individual {
	i := Individual{
		bird:   b,
		neuron: NewRandomNeuron(2, rand.NewSource(time.Now().UnixNano())),
		inputs: make([]float64, 2),
	}

	return &i
}

func (i *Individual) Bird() *components.Bird {
	return i.bird
}

func (i *Individual) Inputs() []float64 {
	return i.inputs
}

func (i *Individual) SetInputs(inputs []float64) {
	i.inputs = inputs
}

func (i *Individual) Neuron() *Neuron {
	return i.neuron
}
