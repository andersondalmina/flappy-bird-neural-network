package neuralnetwork

import "math/rand"

// Predictor interface
type Predictor interface {
	Inputs() int
	Weights() []float64
	Predict([]float64) float64
}

// Neuron is a model of a neuron
type Neuron struct {
	weights []float64
	bias    float64
}

// NewRandomNeuron creates a new neuron with random weights and bias
func NewRandomNeuron(inputs int, src rand.Source) *Neuron {
	r := rand.New(src)
	n := &Neuron{
		weights: make([]float64, inputs),
		bias:    r.Float64(),
	}

	for i := range n.weights {
		n.weights[i] = (r.Float64()*2 - 1) * 1000
	}

	return n
}

// Weights neuron weights
func (n *Neuron) Weights() []float64 {
	return n.weights
}

// Inputs return how many weights a neuron has
func (n *Neuron) Inputs() int {
	return len(n.weights)
}

// SetWeights set new weights values to neuron
func (n *Neuron) SetWeights(w []float64) {
	n.weights = w
}

// Predict is a function that resolves
func (n *Neuron) Predict(inputs []float64) float64 {
	v := 0.0
	for i, input := range inputs {
		v += input * n.weights[i]
	}

	return v + n.bias
}
