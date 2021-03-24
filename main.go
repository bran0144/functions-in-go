package main

import "fmt"

func main() {
	// answer, err := simplemath.Divide(6, 3)
	// if err != nil {
	// 	fmt.Printf("An error occurred %s\n", err.Error())
	// } else {
	// 	fmt.Printf("%f\n", answer)
	// }
	// total := simplemath.Sum(12.2, 14, 16, 22.4)
	// fmt.Printf("total of sum: %f\n", total)

	// sv := simplemath.NewSemanticVerion(1, 2, 3)
	// sv.IncrementMajor()
	// sv.IncrementMajor()
	// sv.IncrementMajor()
	// println(sv.String())

	// var tripper http.RoundTripper = &RoundTripCounter{}
	// r, _ := http.NewRequest(http.MethodGet, "http://pluralsight.com", strings.NewReader("test call"))
	// _, _ = tripper.RoundTrip(r)

	a := func(name string) string {
		fmt.Printf("my first %s function\n", name)
		return name
	}
	value := a("function1")
	println(value)
}

// type RoundTripCounter struct {
// 	count int
// }

// func (rt *RoundTripCounter) RoundTrip(*http.Request) (*http.Response, error) {
// 	rt.count += 1
// 	return nil, nil
// }
