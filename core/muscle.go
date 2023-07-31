package core

type Muscle struct {
	biotype      BioType
	Shape        int
	Synapses     map[int]Synapse
	MuscleMemory map[int]float64
}

func (m *Muscle) RecieveSignal(signal float64, synapsePort int) {
	signal = signal + m.Synapses[synapsePort].Bias

	m.MuscleMemory[synapsePort] = signal
}

func (m *Muscle) AddInputConnection(port int) {
	m.Synapses[port] = NewSynapse() // TODO: delete and add neurotransmitters functional
}

func (m *Muscle) GetAllConnections() []SignalReciever {
	return []SignalReciever{}
}

func (m *Muscle) Type() BioType {
	return m.biotype
}

func (m *Muscle) GetFreePort() int {
	// fmt.Println("free port")
	for i := 0; ; i++ {
		if _, exist := m.Synapses[i]; !exist {
			return i
		}
	}
}

func NewMuscle(shape int) Muscle {
	return Muscle{
		biotype:      NewBioTypeMuscle(),
		Shape:        shape,
		Synapses:     map[int]Synapse{},
		MuscleMemory: map[int]float64{},
	}
}
