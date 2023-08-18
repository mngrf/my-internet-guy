package core

type Neuron struct {
	biotype           BioType
	Dendrites         map[int]Synapse
	MembranePotential float64
	Threshold         float64
	Axon              Axon

	// Experimental
	LearningRate *float64
	Serotonin    *float64
	Activity     float64
}

func (n *Neuron) Fire() {
	n.Activity += 2

	const signal float64 = 25

	for i := 0; i < len(n.Axon.Terminal); i++ {
		n.Axon.Terminal[i].Synapse.RecieveSignal(
			signal+(n.Activity*(*n.LearningRate)),
			n.Axon.Terminal[i].port,
		)
	}

	n.MembranePotential = -10
}

func (n *Neuron) AddInputConnection(port int) {
	n.Dendrites[port] = NewSynapse()
}

func (n *Neuron) AddOutputConnection(sr SignalReciever, port int) {
	n.Axon.Terminal = append(n.Axon.Terminal, BioAddr{
		Synapse: nil,
		port:    port,
	})
}

func (n *Neuron) GetAllConnections() []SignalReciever {
	conns := make([]SignalReciever, len(n.Axon.Terminal))
	for i := 0; i < len(conns); i++ {
		conns[i] = n.Axon.Terminal[i].Synapse
	}

	return conns
}

func (n *Neuron) ConnectTo(synapse SignalReciever) {
	connPort := synapse.GetFreePort()

	n.Axon.Terminal = append(n.Axon.Terminal, BioAddr{
		Synapse: synapse,
		port:    connPort,
	})

	synapse.AddInputConnection(connPort)
}

func (n *Neuron) Process() {
	n.Activity -= 1

	if n.MembranePotential > n.Threshold+*n.Serotonin {
		n.Fire()
	}
}

func (n *Neuron) RecieveSignal(signal float64, dendritePort int) {
	signal = signal + n.Dendrites[dendritePort].Bias

	n.MembranePotential += signal
}

func (n *Neuron) Type() BioType {
	return n.biotype
}

func NewNeuron(learningRate *float64, serotonin *float64) Neuron {
	return Neuron{
		biotype:           NewBioTypeNeuron(),
		Dendrites:         make(map[int]Synapse),
		MembranePotential: 0,
		Threshold:         42,
		Axon:              NewAxon(),
		LearningRate:      learningRate,
		Serotonin:         serotonin,
	}
}

func (n *Neuron) GetFreePort() int {
	for i := 0; ; i++ {
		if _, exist := n.Dendrites[i]; !exist {
			return i
		}
	}
}

type Synapse struct {
	Bias float64
}

func NewSynapse() Synapse {
	return Synapse{
		Bias: 0,
	}
}

type Axon struct {
	Bias     float64
	Terminal []BioAddr
}

func NewAxon() Axon {
	return Axon{
		Bias:     0,
		Terminal: []BioAddr{},
	}
}

type SignalReciever interface {
	AddInputConnection(int)
	GetFreePort() int
	RecieveSignal(signal float64, dendritePort int)
	Type() BioType
	GetAllConnections() []SignalReciever
}

type BioType [2]bool // can store up to 4 types, we need only 3 - organ, neuron, muscle

func NewBioTypeOrgan() BioType {
	return BioType{false, false}
}

func NewBioTypeNeuron() BioType {
	return BioType{false, true}
}

func NewBioTypeMuscle() BioType {
	return BioType{true, false}
}

func (bt BioType) EqualTo(other BioType) bool {
	if bt[0] == other[0] && bt[1] == other[1] {
		return true
	}

	return false
}
