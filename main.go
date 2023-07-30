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
		4,
	)

	brain.LoadSignals(data)

	brain.Run()
}
