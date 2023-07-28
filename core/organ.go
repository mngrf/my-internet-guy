package core

import (
	"fmt"
)

type Organ struct {
	Shape    int
	Terminal map[int]SignalReciever
}

func (o *Organ) SendSignals(signals []float64) {
	fmt.Println("Organ sended signal!")
	if len(signals) != o.Shape {
		panic("Shapes does not match!")
	}

	var i int = 0

	for port, synapse := range o.Terminal {
		synapse.RecieveSignal(signals[i], port)

		i++
	}
}

func NewOrgan(shape int) Organ {
	return Organ{
		Shape:    shape,
		Terminal: make(map[int]SignalReciever, shape),
	}
}

func (o *Organ) ConnectTo(synapse SignalReciever) {
	connPort := synapse.GetFreePort()

	o.Terminal[connPort] = synapse
	synapse.AddInputConnection(connPort)
}
