package core

import (
	"math/rand"

	"github.com/mngrf/my-internet-guy/pkg/mymath"
)

type LSTM struct {
	InputSize, Outputize int

	Memory, LastOutput []float64
}

func NewLSTM(InputSize, Outputize int) *LSTM {
	Memory := make([]float64, InputSize)

	for i := 0; i < InputSize; i++ {
		Memory[i] = rand.NormFloat64()
	}

	LastOutput := make([]float64, Outputize)

	for i := 0; i < Outputize; i++ {
		LastOutput[i] = rand.NormFloat64()
	}

	return &LSTM{
		InputSize:  InputSize,
		Outputize:  Outputize,
		Memory:     Memory,
		LastOutput: LastOutput,
	}
}

func Mul(X, Y []float64) []float64 {
	for i := 0; i < len(Y); i++ {
		X[i] = X[i] * Y[i]
	}

	return X
}

func (lstm *LSTM) Forward(X []float64) []float64 {
	// input gate
	lastOutNewIn := mymath.Add(X, lstm.LastOutput)
	sigmloni := mymath.Sigmoid(lastOutNewIn)

	up := Mul(sigmloni, lstm.Memory)

	// Forgotten gate
	tanhloni := mymath.Tanh(mymath.Add(X, lstm.LastOutput))

	ss := Mul(sigmloni, tanhloni)

	up = mymath.Add(up, ss)

	lstm.Memory = up

	// Output gate
	up = mymath.Tanh(up)

	lp := Mul(up, sigmloni)

	lstm.LastOutput = lp

	return lp
}
