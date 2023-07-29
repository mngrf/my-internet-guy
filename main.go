package main

import (
	"github.com/mngrf/my-internet-guy/core"
)

func main() {
	data := []float64{
		42,
	}

	brain := core.NewBrain(
		[]int{len(data)},
		[]int{2},
		1_000_000,
	)

	brain.LoadSignals([][]float64{data})

	brain.Run()
}
