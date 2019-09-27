package neuralnetwork

import (
	"log"
	"math/rand"

	"github.com/andersondalmina/flappy-bird-neural-network/persist"
)

// NeuralNetwork is a type
type NeuralNetwork struct {
	layers []*Layer
}

// Config is the info about neural network
type Config struct {
	Inputs int64
	Layers []int64
}

// NewNeuralNetwork create a neural network with the config passed
func NewNeuralNetwork(c Config) *NeuralNetwork {
	layers := make([]*Layer, len(c.Layers))
	nn := c.Inputs
	for i, n := range c.Layers {
		layers[i] = NewLayer(n, nn)
		nn = n
	}

	return &NeuralNetwork{
		layers: layers,
	}
}

// Predict computes a forward pass and returns a prediction
func (n *NeuralNetwork) Predict(input []float64) []float64 {
	for _, l := range n.layers {
		input = l.Predict(input)
	}

	result := make([]float64, len(input))

	for i, value := range input {
		result[i] = relu(value)
	}

	return result
}

// Weights return the weights of all neurons of the network
func (n *NeuralNetwork) Weights() [][][]float64 {
	res := make([][][]float64, len(n.layers))

	for il, l := range n.layers {
		res[il] = make([][]float64, len(l.neurons))
		for in, ne := range l.neurons {
			res[il][in] = ne.Weights()
		}
	}

	return res
}

// UpdateWeights set the weight of each neuron
func (n *NeuralNetwork) UpdateWeights(w [][][]float64) {
	weights := ajustWeight(w)

	for il, l := range n.layers {
		for in, ne := range l.neurons {
			ne.SetWeights(weights[il][in])
		}
	}
}

// Dump neural network weights to a file
func (n *NeuralNetwork) Dump(filepath string) error {
	err := persist.Save(filepath, n.Weights())

	if err != nil {
		log.Fatalln(err)
	}

	return nil
}

// ImportDump import the neural weights
func (n *NeuralNetwork) ImportDump(filepath string) error {
	var data [][][]float64
	err := persist.Load(filepath, &data)

	if err != nil {
		log.Fatalln(err)
	}

	n.UpdateWeights(data)

	return nil
}

// ajustWeight make changes in weights
func ajustWeight(bestWeights [][][]float64) [][][]float64 {
	w := make([][][]float64, len(bestWeights))

	for i := range bestWeights {
		w[i] = make([][]float64, len(bestWeights[i]))

		for ii := range bestWeights[i] {
			w[i][ii] = make([]float64, len(bestWeights[i][ii]))

			for iii := range bestWeights[i][ii] {
				w[i][ii][iii] = bestWeights[i][ii][iii]
				if rand.Intn(4) == 0 {
					w[i][ii][iii] += (rand.Float64()*2 - 1) * 100
				}
			}
		}
	}

	return w
}

func relu(x float64) float64 {
	if x <= 0 {
		return 0
	}

	return 1
}
