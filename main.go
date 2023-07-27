package main

import (
	"fmt"

	"github.com/mngrf/my-internet-guy/core"
)

func main() {
	b := core.NewBrain([]int{
		1, 20, 256 * 256,
	}, 1_000_000)

	fmt.Print(b)
}
