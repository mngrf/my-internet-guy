package core

import (
	"fmt"
	"math/rand"
)

type Brain struct {
	Organs       []Organ
	neuronsCount int
	Neurons      []Neuron
	Muscles      []Muscle
}

func (b *Brain) ConnectOrgansToNeurons() {
	neuronsCount := len(b.Neurons)

	for i := 0; i < len(b.Organs); i++ {
		for j := 0; j < b.Organs[i].Shape; j++ {
			b.Organs[i].ConnectTo(&b.Neurons[rand.Intn(neuronsCount)])
		}
	}
}

func (b *Brain) ConnectMusclesToNeurons() {
	neuronsCount := len(b.Neurons)

	for i := 0; i < len(b.Muscles); i++ {
		for j := 0; j < b.Muscles[i].Shape; j++ {
			b.Neurons[rand.Intn(neuronsCount)].ConnectTo(&b.Muscles[i])
		}
	}
}

func (b *Brain) Tick() {
	for i := 0; i < b.neuronsCount; i++ {
		b.Neurons[i].Process()
	}

	for i := 0; i < len(b.Organs); i++ {
		b.Organs[i].ProcessSignals()
	}

	for i := 0; i < len(b.Muscles); i++ {
		fmt.Println(b.Muscles[i].MuscleMemory) // TODO: handle muscles buffers
	}
}

func (b *Brain) ConnectNeurons() {
	neuronsCount := len(b.Neurons)

	if neuronsCount < 2 {
		return
	}

	for i := 0; i < neuronsCount; i++ {
		b.Neurons[i].ConnectTo(
			&b.Neurons[rand.Intn(neuronsCount)],
		)
	}
}

func (b *Brain) ProcessSignals(signals [][]float64) {
	if len(signals) != len(b.Organs) {
		panic("Shapes does not match")
	}

	for i := 0; i < len(signals); i++ {
		b.Organs[i].ProcessSignals()
	}
}

func (b *Brain) LoadSignals(signals [][]float64) {
	for i := 0; i < len(b.Organs); i++ {
		b.Organs[i].LoadSignals(signals[i])
	}

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

	brain := Brain{
		Organs:       organs,
		neuronsCount: neuronsCount,
		Neurons:      neurons,
		Muscles:      muscles,
	}

	brain.ConnectOrgansToNeurons()
	brain.ConnectNeurons()
	brain.ConnectMusclesToNeurons()

	return &brain
}
