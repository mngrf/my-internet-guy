package main

import (
	"fmt"
	"time"

	"github.com/mngrf/my-internet-guy/core"
)

func main() {
	data := []float64{420000}

	brain := core.NewBrain(
		[]int{1},
		[]int{1},
		1_000_000,
		10,
	)

	brain.GenerateNeuronConnections()

	fmt.Println(brain.IsAllInputsToAllOutputs())

	start := time.Now()
	brain.LoadSignals(data)
	for i := 0; i < 60; i++ {
		brain.ProcessSignals()
	}
	end := time.Now()

	fmt.Println(end.Sub(start))
}
