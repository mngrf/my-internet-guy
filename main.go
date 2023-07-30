package main

import (
	"github.com/mngrf/my-internet-guy/core"
)

func main() {
	data := []float64{
		4200000000,
	}

	brain := core.NewBrain(
		[]int{len(data)},
		[]int{2, 3},
		1_000_000,
	)

	brain.LoadSignals(data)

	brain.Run()
}
