package core

import (
	"fmt"
)

type Neuron struct {
	Dendrites         map[int]Synapse
	MembranePotential float64
	Threshold         float64
	Axon              Axon
}

func (n *Neuron) Fire() {
	fmt.Println("Fire")
	for synapse, port := range n.Axon.Terminal {
		synapse.RecieveSignal(n.MembranePotential, port)
	}

	n.MembranePotential = 0
}

func (n *Neuron) AddInputConnection(port int) {
	n.Dendrites[port] = NewSynapse()
}

func (n *Neuron) AddOutputConnection(sr SignalReciever, port int) {
	n.Axon.Terminal[sr] = port
}

func (n *Neuron) ConnectTo(synapse SignalReciever) {
	connPort := n.GetFreePort()

	n.Dendrites[connPort] = NewSynapse()

	n.Axon.Terminal[synapse] = connPort
}

func (n *Neuron) RecieveSignal(signal float64, dendritePort int) {
	fmt.Println("Neuron recieved signal!")
	signal = signal + n.Dendrites[dendritePort].Bias

	n.MembranePotential += signal

	if n.MembranePotential > n.Threshold {
		n.Fire()
	}
}

func NewNeuron() Neuron {
	return Neuron{
		Dendrites:         make(map[int]Synapse),
		MembranePotential: 30,
		Threshold:         42,
		Axon:              NewAxon(),
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
	Terminal map[SignalReciever]int
}

func NewAxon() Axon {
	return Axon{
		Bias:     0,
		Terminal: make(map[SignalReciever]int),
	}
}

type SignalReciever interface {
	AddInputConnection(int)
	GetFreePort() int
	RecieveSignal(signal float64, dendritePort int)
}
