package components

import (
	"github.com/andersondalmina/flappy-bird-neural-network/neuralnetwork"
)

// Individual component
type Individual struct {
	bird   *Bird
	neural *neuralnetwork.NeuralNetwork
	inputs []float64
}

// NewIndividual create a new individual
func NewIndividual(b *Bird, n *neuralnetwork.NeuralNetwork) *Individual {
	i := Individual{
		bird:   b,
		neural: n,
		inputs: make([]float64, 2),
	}

	return &i
}

// Bird return the individual bird
func (i *Individual) Bird() *Bird {
	return i.bird
}

// Inputs return the inputs
func (i *Individual) Inputs() []float64 {
	return i.inputs
}

// SetInputs set the inputs
func (i *Individual) SetInputs(inputs []float64) {
	i.inputs = inputs
}

// Neural return the neural
func (i *Individual) Neural() *neuralnetwork.NeuralNetwork {
	return i.neural
}
