package main

import (
	"fmt"
	"functions/simplemath"
)

type MathExpr = string

const (
	AddExpr      = MathExpr("add")
	SubtractExpr = MathExpr("subtract")
	MultiplyExpr = MathExpr("divide")
)

func main() {
	// addExpr := mathExpression(AddExpr)
	// println(addExpr(2, 3))
	fmt.Printf("%f", double(3, 2, mathExpression(AddExpr)))
}

func mathExpression(expr MathExpr) func(float64, float64) float64 {
	switch expr {
	case AddExpr:
		return simplemath.Add
	case SubtractExpr:
		return simplemath.Subtract
	case MultiplyExpr:
		return simplemath.Multiply
	default:
		return func(f float64, f2 float64) float64 {
			return 0
		}
	}

}

func double(f1, f2 float64, MathExpr func(float64, float64) float64) float64 {
	return 2 * MathExpr(f1, f2)
}
