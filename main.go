package main

import "github.com/mngrf/my-internet-guy/core"

func main() {
	organ := core.NewOrgan(3, 3)

	neuron1 := core.NewCell(10, 10)
	neuron2 := core.NewCell(10, 10)

	muscle := core.NewMuscle(4, 1)

	organ.Cells[0][0].ConnectTo(&neuron1)
	organ.Cells[0][1].ConnectTo(&neuron2)

	neuron1.ConnectTo(&neuron2)

	neuron1.ConnectTo(muscle.Cells[0][0])
	neuron2.ConnectTo(muscle.Cells[0][1])

	organ.SendSignals([][]float64{
		{1, 2, 3},
	})
}
