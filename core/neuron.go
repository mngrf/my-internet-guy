package core

type Neuron struct {
	Dendrites         map[int]Synapse
	MembranePotential float64
	Threshold         float64
	Axon              Axon
}

func (n *Neuron) Fire() {
	// fmt.Println("Fire")
	for i := 0; i < len(n.Axon.Terminal); i++ {
		n.Axon.Terminal[i].Synapse.RecieveSignal(
			n.MembranePotential,
			n.Axon.Terminal[i].port,
		)
	}

	n.MembranePotential = 0
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

func (n *Neuron) ConnectTo(synapse SignalReciever) {
	connPort := synapse.GetFreePort()

	n.Axon.Terminal = append(n.Axon.Terminal, BioAddr{
		Synapse: synapse,
		port:    connPort,
	})

	synapse.AddInputConnection(connPort)
}

func (n *Neuron) Process() {
	if n.MembranePotential > n.Threshold {
		n.Fire()
	}
}

func (n *Neuron) RecieveSignal(signal float64, dendritePort int) {
	// fmt.Println("Neuron recieved signal", signal)
	signal = signal + n.Dendrites[dendritePort].Bias

	n.MembranePotential += signal
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
}
