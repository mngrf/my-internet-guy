package main

import (
	"fmt"
	"time"

	"github.com/mngrf/my-internet-guy/core"
)

func main() {
	data := []float64{42}

	brain := core.NewBrain(
		[]int{1},
		[]int{1},
		100_000,
		1000,
	)
	brain.GenerateNeuronConnections()
	brain.GenerateNeuronConnections()
	brain.GenerateNeuronConnections()

	fmt.Println(brain.IsAllInputsToAllOutputs())

	minuteStart := time.Now()
	for i := 0; i < 60; i++ {
		brain.LoadSignals(data)
		for i := 0; i < 60; i++ {
			brain.ProcessSignals()
		}
	}
	minuteEnd := time.Since(minuteStart)
	fmt.Println("minute:", minuteEnd)

	fmt.Println(brain.Muscles[0].MuscleMemory[0])
}
