package core

import (
	"fmt"
)

type Muscle struct {
	Shape        int
	Synapses     map[int]Synapse
	MuscleMemory map[int]float64
}

func (m *Muscle) RecieveSignal(signal float64, synapsePort int) {
	fmt.Println("Muscle recieved signal!")
	signal = signal + m.Synapses[synapsePort].Bias

	m.MuscleMemory[synapsePort] = signal
}

func (m *Muscle) AddInputConnection(port int) {
	m.Synapses[port] = Synapse{
		Bias: -15, // TODO: delete and add neurotransmitters functional
	}
}

func (m *Muscle) GetFreePort() int {
	fmt.Println("free port")
	for i := 0; ; i++ {
		if _, exist := m.Synapses[i]; !exist {
			return i
		}
	}
}

func NewMuscle(shape int) Muscle {
	return Muscle{
		Shape:        shape,
		Synapses:     map[int]Synapse{},
		MuscleMemory: make(map[int]float64, shape),
	}
}
