package main

import (
	"fmt"

	"github.com/mngrf/my-internet-guy/core"
)

func main() {

	o := core.NewBioTypeOrgan()
	n := core.NewBioTypeNeuron()
	x := core.NewBioTypeNeuron()

	fmt.Println(o.EqualTo(n), o.EqualTo(x), n.EqualTo(x))

	data := []float64{
		42,
	}

	brain := core.NewBrain(
		[]int{len(data)},
		[]int{2},
		100,
	)

	brain.LoadSignals(data)

	brain.Run()
}
