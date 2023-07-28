package core

import (
	"fmt"
)

type Muscle struct {
	Synapses     map[int]Synapse
	muscleMemory map[int]float64
}

func (m *Muscle) RecieveSignal(signal float64, synapsePort int) {
	fmt.Println("Muscle recieved signal!")
	signal = signal + m.Synapses[synapsePort].Bias

	m.muscleMemory[synapsePort] = signal
}

func (m *Muscle) AddInputConnection(port int) {
	m.Synapses[port] = NewSynapse()
}

func (m *Muscle) GetFreePort() int {
	for i := 0; ; i++ {
		if _, exist := m.Synapses[i]; !exist {
			return i
		}
	}
}

func NewMuscle() *Muscle {
	return &Muscle{
		Synapses:     map[int]Synapse{},
		muscleMemory: map[int]float64{},
	}
}
