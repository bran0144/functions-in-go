package simplemath

import (
	"errors"
	"math"
)

func Sum(values ...float64) float64 {
	total := 0.0
	for _, value := range values {
		total += value
	}
	return total
}

func Divide(p1, p2 float64) (float64, error) {
	if p2 == 0 {
		return math.NaN(), errors.New("cannot divide by 0")
	}
	return p1 / p2, nil
}
func Add(p1, p2 float64) float64 {
	return p1 + p2
}
