package mymath

type Vector3 struct {
	X, Y, Z float64
}

func Add(X, Y []float64) []float64 {
	if len(X) != len(Y) {
		panic("Shapes does not match!")
	}

	for i, val := range Y {
		X[i] = X[i] + val
	}

	return X
}
