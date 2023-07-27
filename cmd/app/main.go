package main

import (
	"fmt"

	"github.com/mngrf/my-internet-guy/internal/core"
)

func main() {
	X := []float64{1, 2, 3}

	lstm := core.NewLSTM(len(X), len(X))

	fmt.Print(lstm.Forward(X), "\n")
	fmt.Print(lstm.Forward(X), "\n")
	fmt.Print(lstm.Forward(X), "\n")
	fmt.Print(lstm.Forward(X), "\n")
	fmt.Print(lstm.Forward(X), "\n")
}
