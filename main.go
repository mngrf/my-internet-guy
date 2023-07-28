package main

import (
	"github.com/mngrf/my-internet-guy/core"
)

func main() {
	data := []float64{
		0, 255, 0,
		0, 255, 0, // Image that this is the 3x3 image
		255, 255, 255,
	}

	brain := core.NewBrain(
		[]int{len(data)}, // 1 input organ with shape of an image
		[]int{3, 5 * 5},  // 2 muscle outputs with shapes 3 and 5*5=25.
		1,                // 1 neuron.
	)

	brain.ProcessSignals([][]float64{data})
}

// TODO: rewrite brain.ProcessSignals() to load data to neurons firstly and then process it with neurons.
