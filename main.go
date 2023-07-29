package main

import (
	"github.com/mngrf/my-internet-guy/core"
)

func main() {
	data := []float64{
		100,
	}

	brain := core.NewBrain(
		[]int{len(data)},
		[]int{1},
		1,
	)

	brain.LoadSignals([][]float64{data})
	brain.Tick()
	brain.Tick()
}
