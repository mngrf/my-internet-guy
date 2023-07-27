package core

import "fmt"

type Brain struct {
	Organs  [][]float64
	Neurons []Neuron
	Muscles []func(float64)
}

func Move(signal float64) { fmt.Println("Moved: ", signal) }

func NewBrain(organShapes []int, neuronsCount int) *Brain {
	organs := make([][]float64, len(organShapes))
	for i := 0; i < len(organShapes); i++ {
		organs[i] = make([]float64, organShapes[i])
	}

	neurons := make([]Neuron, neuronsCount)
	for i := 0; i < neuronsCount; i++ {
		neurons[i] = NewNeuron()
	}

	brain := &Brain{
		Organs:  organs,
		Neurons: neurons,
		Muscles: []func(float64){Move},
	}

	return brain
}
