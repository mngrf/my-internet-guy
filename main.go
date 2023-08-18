package main

import (
	"fmt"
	"time"

	"github.com/mngrf/my-internet-guy/core"
)

func main() {
	data := []float64{60}

	brain := core.NewBrain(
		[]int{1},  // sensors/organs
		[]int{1},  // muscles
		1_000_000, // num of a neurons
		1000,      //neuron sections
		1,         // learning rate
	)

	brain.GenerateNeuronConnections()

	fmt.Println(brain.IsAllInputsToAllOutputs())

	start := time.Now()
	for j := 0; j < 60; j++ {
		brain.LoadSignals(data)
		for i := 0; i < 60; i++ {
			brain.ProcessSignals()
		}
	}
	end := time.Now()

	fmt.Println(end.Sub(start))
}
