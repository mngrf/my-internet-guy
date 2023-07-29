package main

import (
	"fmt"

	"github.com/mngrf/my-internet-guy/core"
)

func main() {
	data := []float64{
		100,
	}

	brain := core.NewBrain(
		[]int{len(data)},
		[]int{3, 2},
		1,
	)

	brain.ProcessSignals([][]float64{data})

	fmt.Println(brain.Muscles)
}
