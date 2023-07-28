package main

import (
	"fmt"

	"github.com/mngrf/my-internet-guy/core"
)

func main() {
	brain := core.NewBrain(
		[]int{128 * 128, 3},
		[]int{6 * 3},
		1000,
	)

	fmt.Println(brain)
}
