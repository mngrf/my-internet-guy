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
		500,
		4,
	)

	fmt.Println(brain.IsAllInputsToAllOutputs())
}
