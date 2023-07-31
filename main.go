package main

import (
	"fmt"

	"github.com/mngrf/my-internet-guy/core"
)

func main() {
	// data := []float64{42}

	brain := core.NewBrain(
		[]int{1},
		[]int{4},
		10_000_000,
		4,
	)

	fmt.Println(brain.IsAllInputsToAllOutputs())
}
