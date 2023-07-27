package core

import "fmt"

type Neuron struct {
	Dendrits map[int]Synapse

	MembranePotential float64
	Treshold          float64

	Axon Axon
}

func (n *Neuron) Fire() {
	for neuron, port := range n.Axon.Terminal {
		neuron.RecieveSignal(n.MembranePotential, port)
	}

	fmt.Println("Fire!")

	n.MembranePotential = 0
}

func (n *Neuron) ConnectTo(neuron *Neuron) {
	connPort := neuron.getFreePort()

	neuron.Dendrits[connPort] = NewSynapse()

	n.Axon.Terminal[neuron] = connPort
}

func (n *Neuron) RecieveSignal(signal float64, dendritPort int) {
	signal = signal + n.Dendrits[dendritPort].Bias

	n.MembranePotential += signal

	if n.MembranePotential > n.Treshold {
		n.Fire()
	}
}

func NewNeuron() *Neuron {
	return &Neuron{
		Dendrits:          map[int]Synapse{},
		MembranePotential: 30,
		Treshold:          42,
		Axon:              NewAxon(),
	}
}

func (n *Neuron) getFreePort() int {
	for i := 0; ; i++ {
		if _, exist := n.Dendrits[i]; !exist {
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
	Terminal map[*Neuron]int
}

func NewAxon() Axon {
	return Axon{
		Bias:     0,
		Terminal: map[*Neuron]int{},
	}
}
