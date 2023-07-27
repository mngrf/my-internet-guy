package core

import "fmt"

type Cell struct {
	Input  []*Cell
	Output []*Cell
	Charge float64
}

func (c *Cell) ConnectTo(cell *Cell) {
	c.Output[0] = cell

	cell.Input[0] = c
}

func (c *Cell) SendSignal(signal float64) {
	for i := 0; i < len(c.Output); i++ {
		c.Output[i].RecieveSignal(signal)
	}
}

func (c *Cell) RecieveSignal(signal float64) {
	c.Charge += signal
}

func NewSourceCell(outputSize int) Cell {
	return Cell{
		Output: make([]*Cell, outputSize),
		Charge: 0,
	}
}

func NewDestinationCell(inputSize int) Cell {
	return Cell{
		Input:  make([]*Cell, inputSize),
		Charge: 0,
	}
}

func NewCell(inSize, outSize int) Cell {
	return Cell{
		Input:  make([]*Cell, inSize),
		Output: make([]*Cell, outSize),
		Charge: 0,
	}
}

// ///////////////////////////////////////////////////////////////////////////////////////////////////////////////
type Organ struct {
	Width, Height int
	Cells         [][]*Cell
}

func (o *Organ) SendSignals(signals [][]float64) {
	if len(signals) != o.Height {
		return
	}

	if len(signals[0]) != o.Width {
		return
	}

	for i := 0; i < o.Height; i++ {
		for j := 0; j < o.Width; j++ {
			if o.Cells[i][j] == nil {
				fmt.Println("nil cells[i][j]")
				continue
			}

			for c := 0; c < len(o.Cells[i][j].Output); c++ {
				if o.Cells[i][j].Output[c] == nil {
					fmt.Println("tse ban!")
					continue
				}

				o.Cells[i][j].Output[c].RecieveSignal(signals[i][j])
			}
		}
	}
}

func NewOrgan(width, height int) Organ {
	cells := make([][]*Cell, height)

	for i := 0; i < width; i++ {
		cells[i] = make([]*Cell, width)
	}

	return Organ{
		Width:  width,
		Height: height,
		Cells:  cells,
	}
}

// ////////////////////////////////////////////////////////////////////////////////////////////////////////////////
type Muscle struct {
	Cells [][]*Cell
}

func NewMuscle(width, height int) Muscle {
	cells := make([][]*Cell, height)

	for i := 0; i < height; i++ {
		cells[i] = make([]*Cell, width)
	}

	return Muscle{
		Cells: cells,
	}
}
