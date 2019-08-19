package neuralnetwork

// NeuralNetwork is a type
type NeuralNetwork struct {
	layers []Layer
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
