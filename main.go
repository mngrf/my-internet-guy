package main

import "github.com/mngrf/my-internet-guy/core"

func main() {
	n1 := core.NewNeuron()
	n2 := core.NewNeuron()

	n1.ConnectTo(n2)

	n1.Dendrits[0] = core.NewSynapse()
	n1.RecieveSignal(20, 0)
}
