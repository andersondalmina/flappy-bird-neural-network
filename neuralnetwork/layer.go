package neuralnetwork

// Layer is a Layer struct
type Layer struct {
	neurons   []Neuron
	nextLayer *Layer
}

func NewLayer() *Layer {
	return &Layer{}
}
