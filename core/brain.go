package core

import (
	"math/rand"
)

type Brain struct {
	Organs  []Organ
	Neurons []Neuron
	Muscles []Muscle
}

func (b *Brain) ConnectOrgansToNeurons() {
	neuronsCount := len(b.Neurons)

	for i, _ := range b.Organs {
		b.Organs[i].ConnectTo(&b.Neurons[rand.Intn(neuronsCount)])
	}
}

func (b *Brain) ConnectMusclesToNeurons() {

}

func NewBrain(organShapes, muscleShapes []int, neuronsCount int) *Brain {
	organs := make([]Organ, len(organShapes))
	for i := 0; i < len(organShapes); i++ {
		organs[i] = NewOrgan(organShapes[i])
	}

	neurons := make([]Neuron, neuronsCount)
	for i := 0; i < neuronsCount; i++ {
		neurons[i] = NewNeuron()
	}

	muscles := make([]Muscle, len(muscleShapes))
	for i := 0; i < len(muscleShapes); i++ {
		muscles[i] = NewMuscle(muscleShapes[i])
	}

	return &Brain{
		Organs:  organs,
		Neurons: neurons,
		Muscles: muscles,
	}
}
