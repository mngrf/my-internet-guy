package core

type Organ struct {
	Shape     int
	tempStore []float64
	Terminal  []BioAddr
}

func (o *Organ) LoadSignals(signals []float64) {
	if len(signals) != o.Shape {
		return
	}

	o.tempStore = signals
}

type BioAddr struct {
	Synapse SignalReciever
	port    int
}

func (ba *BioAddr) SendSignal(signal float64) {
	ba.Synapse.RecieveSignal(signal, ba.port)
}

func (o *Organ) ProcessSignals() {
	for i := 0; i < o.Shape; i++ {
		o.Terminal[i].Synapse.RecieveSignal(
			o.tempStore[i],
			o.Terminal[i].port,
		)
	}
}

func NewOrgan(shape int) Organ {
	return Organ{
		Shape:    shape,
		Terminal: []BioAddr{},
	}
}

func (o *Organ) ConnectTo(synapse SignalReciever) {
	connPort := synapse.GetFreePort()

	synapse.AddInputConnection(connPort)
	o.Terminal = append(o.Terminal, BioAddr{
		Synapse: synapse,
		port:    connPort,
	})
}
