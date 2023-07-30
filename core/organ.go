package core

type Organ struct {
	biotype   BioType
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

func (o *Organ) GetAllConnections() []SignalReciever {
	sr := make([]SignalReciever, len(o.Terminal))

	for i := 0; i < len(o.Terminal); i++ {
		sr[i] = o.Terminal[i].Synapse
	}

	return sr
}

func (o *Organ) Type() BioType {
	return o.biotype
}

func NewOrgan(shape int) Organ {
	return Organ{
		biotype:  NewBioTypeOrgan(),
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
