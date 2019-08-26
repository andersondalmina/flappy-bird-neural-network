package neuralnetwork

import (
	"math/rand"
	"time"
)

// Layer is a Layer struct
type Layer struct {
	neurons []*Neuron
}

func NewLayer(numberNeurons, numberInputs int64) *Layer {
	neurons := make([]*Neuron, numberNeurons)
	for i := range neurons {
		neurons[i] = NewNeuron(numberInputs, rand.NewSource(time.Now().UTC().UnixNano()))
	}

	return &Layer{
		neurons: neurons,
	}
}

func (l *Layer) Neurons() []*Neuron {
	return l.neurons
}

func (l *Layer) Run(inputs []float64) []float64 {
	r := make([]float64, len(l.neurons))
	for i, n := range l.neurons {
		r[i] = n.Predict(inputs)
	}

	return r
}
