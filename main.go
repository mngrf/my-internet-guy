package main

import (
	"fmt"

	"github.com/mngrf/my-internet-guy/core"
)

func main() {
	brain := core.NewBrain(
		[]int{1},
		[]int{1},
		1,
	)

	brain.Organs[0].ConnectTo(&brain.Neurons[0])
	brain.Neurons[0].ConnectTo(&brain.Muscles[0])

	brain.Organs[0].SendSignals([]float64{42})

	fmt.Println(brain)
}
