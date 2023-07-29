package core

import (
	"fmt"
)

type Organ struct {
	Shape    int
	Terminal []BioAddr
}

type BioAddr struct {
	Synapse SignalReciever
	port    int
}

func (ba *BioAddr) SendSignal(signal float64) {
	ba.Synapse.RecieveSignal(signal, ba.port)
}

func (o *Organ) SendSignals(signals []float64) {
	fmt.Println("Organ has sent the signals")
	if len(signals) != o.Shape {
		panic("Shapes does not match!")
	}

	if len(signals) != len(o.Terminal) {
		fmt.Println(len(signals), len(o.Terminal))
		panic("Shapes does not match!")
	}

	for i := 0; i < len(o.Terminal); i++ {
		o.Terminal[i].SendSignal(signals[i])
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
