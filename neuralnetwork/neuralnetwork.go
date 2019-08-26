package neuralnetwork

// NeuralNetwork is a type
type NeuralNetwork struct {
	layers []*Layer
}

type Config struct {
	Inputs int64
	Layers []int64
}

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
		input = l.Run(input)
	}

	return input
}

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

func (n *NeuralNetwork) SetWeights(w [][][]float64) {
	for il, l := range n.layers {
		for in, ne := range l.neurons {
			ne.SetWeights(w[il][in])
		}
	}
}

// MeanSquaredErrorValue resolves the loss
func MeanSquaredErrorValue(p Predictor, data [][]float64) float64 {
	mse := 0.0
	for _, row := range data {
		inputs, output := row[:p.Inputs()], row[p.Inputs()]
		err := output - p.Predict(inputs)
		mse += err * err
	}

	return mse / float64(len(data))
}
