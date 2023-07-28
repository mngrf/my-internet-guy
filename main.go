package main

import (
	"github.com/mngrf/my-internet-guy/core"
)

func main() {
	o := core.NewOrgan()

	n1 := core.NewNeuron()
	n2 := core.NewNeuron()
	n3 := core.NewNeuron()

	m := core.NewMuscle()

	o.ConnectTo(&n1)
	o.ConnectTo(&n3)

	n3.ConnectTo(&n1)
	n1.ConnectTo(&n2)
	n2.ConnectTo(m)

	o.SendSignals([]float64{50})
}
