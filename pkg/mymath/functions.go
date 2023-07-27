package mymath

import "math"

func Ssigmoid(x float64) float64 {
	return 1 / (1 + math.Exp(-x))
}

func Stanh(x float64) float64 {
	return (math.Exp(x) - math.Exp(-x)) / (math.Exp(x) + math.Exp(-x))
}

func Sigmoid(vector []float64) []float64 {
	for i := 0; i < len(vector); i++ {
		vector[i] = Ssigmoid(vector[i])
	}

	return vector
}

func Tanh(vector []float64) []float64 {
	for i := 0; i < len(vector); i++ {
		vector[i] = Stanh(vector[i])
	}

	return vector
}
