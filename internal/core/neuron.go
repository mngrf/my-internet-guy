package core

type Neuron struct {
	Dendrits []*Neuron
	Axons    []*Neuron
}
