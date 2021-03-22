package main

import (
	"fmt"
	"functions/simplemath"
)

func main() {
	answer, err := simplemath.Divide(6, 3)
	if err != nil {
		fmt.Printf("An error occurred %s\n", err.Error())
	} else {
		fmt.Printf("%f\n", answer)
	}
	total := simplemath.Sum(12.2, 14, 16, 22.4)
	fmt.Printf("total of sum: %f\n", total)
}
