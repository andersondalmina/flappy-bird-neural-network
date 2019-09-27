package neuralnetwork

import (
	"math/rand"
	"testing"
	"time"
)

func TestNewNeuron(t *testing.T) {
	src := rand.NewSource(time.Now().UTC().UnixNano())

	n := NewNeuron(3, src)

	if n.Inputs() != 3 {
		t.Errorf("Inputs Len must be %d, has %d", 3, n.Inputs())
	}
}
