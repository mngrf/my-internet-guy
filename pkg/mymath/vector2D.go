package mymath

import "fmt"

type Vector2D struct {
	Width  int
	Height int
	table  [][]float64
}

func NewVector2D(width, height int) *Vector2D {
	// Create a 2D slice with the given width and height
	table := make([][]float64, height)
	for i := range table {
		table[i] = make([]float64, width)
	}

	return &Vector2D{
		Width:  width,
		Height: height,
		table:  table,
	}
}

func (v2d *Vector2D) Mul(other *Vector2D) *Vector2D {
	if v2d.Width != other.Height {
		panic(fmt.Sprintf("Dimensions are not comparable! %v dim: %v and %v dim: %v", v2d, v2d.Width, other, other.Height))
	}

	result := NewVector2D(other.Width, v2d.Height)

	for i := 0; i < v2d.Height; i++ {
		for j := 0; j < other.Width; j++ {
			var sum float64
			for k := 0; k < v2d.Width; k++ {
				sum += v2d.table[i][k] * other.table[k][j]
			}
			result.table[i][j] = sum
		}
	}

	return result
}
