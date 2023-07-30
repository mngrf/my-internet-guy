package core

import (
	"fmt"
	"math/rand"
	"time"
)

type Brain struct {
	Organs       []Organ
	neuronsCount int
	Neurons      []Neuron
	Muscles      []Muscle
}

func (b *Brain) connectOrgansToNeurons() {
	neuronsCount := len(b.Neurons)

	for i := 0; i < len(b.Organs); i++ {
		for j := 0; j < b.Organs[i].Shape; j++ {
			b.Organs[i].ConnectTo(&b.Neurons[rand.Intn(neuronsCount)])
		}
	}
}

func (b *Brain) connectMusclesToNeurons() {
	neuronsCount := len(b.Neurons)

	for i := 0; i < len(b.Muscles); i++ {
		for j := 0; j < b.Muscles[i].Shape; j++ {
			b.Neurons[rand.Intn(neuronsCount)].ConnectTo(&b.Muscles[i])
		}
	}
}

func (b *Brain) allInputsToAllOutputs() bool {
	result := make([]bool, len(b.Organs))

	for i := 0; i < len(b.Organs); i++ {
		result[i] = b.organInputToOutputs(&b.Organs[i])
	}

	for i := 0; i < len(result); i++ {
		if !result[i] {
			return false
		}
	}

	return true
}

func (b *Brain) organInputToOutputs(organ *Organ) bool {
	cheked := make(map[*SignalReciever]bool)

	var results []bool = []bool{}

	for i := 0; i < len(organ.Terminal); i++ {
		b.neuronInputToOutputs(organ.Terminal[i].Synapse, &cheked, &results)
	}

	for _, result := range results {
		if result {
			return true
		}
	}

	return false
}

func (b *Brain) neuronInputToOutputs(sr SignalReciever, cheked *map[*SignalReciever]bool, results *[]bool) {
	if sr.Type().EqualTo(NewBioTypeMuscle()) {
		*results = append(*results, true)
	}

	(*cheked)[&sr] = true

	for _, conn := range sr.GetAllConnections() {
		if (*cheked)[&conn] {
			continue
		}

		b.neuronInputToOutputs(conn, cheked, results)
	}
}

func (b *Brain) Tick() {
	for i := 0; i < len(b.Organs); i++ {
		b.Organs[i].ProcessSignals()
	}

	for i := 0; i < b.neuronsCount; i++ {
		b.Neurons[i].Process()
	}

	for i := 0; i < len(b.Muscles); i++ {
		// fmt.Println("muscle buffer: ", b.Muscles[i].MuscleMemory) // TODO: handle muscles buffers
	}
}

func (b *Brain) Run() {
	startTime := time.Now() // Record the current time (start time)

	for i := 0; i < 60; i++ {
		b.Tick() // Call the 'Tick' method of the 'Brain' struct
	}

	endTime := time.Now() // Record the current time (end time)

	elapsedTime := endTime.Sub(startTime) // Calculate the time difference between start and end time

	fmt.Println(elapsedTime, b.Muscles[0].MuscleMemory) // Print the elapsed time
}

func (b *Brain) connectNeurons() {
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

func (b *Brain) LoadSignals(signals ...[]float64) {
	if len(signals) != len(b.Organs) {
		panic("Shapes do not match!")
	}

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

	brain.connectOrgansToNeurons()
	brain.connectMusclesToNeurons()

	for {
		if !brain.allInputsToAllOutputs() {
			brain.connectNeurons()
			continue
		}

		break
	}

	return &brain
}
