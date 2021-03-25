package main

import (
	"errors"
	"fmt"
	"functions/simplemath"
	"io"
	"math"
)

type MathExpr = string

const (
	AddExpr      = MathExpr("add")
	SubtractExpr = MathExpr("subtract")
	MultiplyExpr = MathExpr("divide")
)

func main() {
	if err := ReadFullFile(); err != nil {
		println("somthing bad occurred")
	}
}

func ReadFullFile() error {
	var r io.ReadCloser = &SimpleReader{}
	defer func() {
		_ = r.Close()
	}()

	defer func() {
		println("before for-loop")
	}()
	for {
		value, err := r.Read([]byte("text that does nothing"))
		if err == io.EOF {
			println("finished reading file")
			break
		} else if err != nil {
			return err
		}
		println(value)
	}

	defer func() {
		println("after for-loop")
	}()
	return nil
}

func ReadSomethingBad() error {
	var r io.Reader = BadReader{errors.New("my nonsense reader")}
	if _, err := r.Read([]byte("test something")); err != nil {
		fmt.Printf("an error occurred %s", err)
		return err
	}
	return nil
}

type BadReader struct {
	err error
}

func (br BadReader) Read(p []byte) (n int, err error) {
	return -1, br.err
}

type SimpleReader struct {
	count int
}

func (br *SimpleReader) Read(p []byte) (n int, err error) {
	if br.count > 3 {
		return 0, io.EOF
	}
	br.count += 1
	return br.count, nil
}

func (br *SimpleReader) Close() error {
	println("closing reader")
	return nil
}

func powerOfTwo() func() int64 {
	x := 1.0
	return func() int64 {
		x += 1
		return int64(math.Pow(x, 2))
	}
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
