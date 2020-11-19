// Package calculator provides a library for simple calculations in Go.
package calculator

import (
	"fmt"
	"math"
)

// Add takes two numbers and returns the result of adding them together.
func Add(a, b float64) float64 {
	return a + b
}

// Subtract takes two numbers and returns the result of subtracting the second
// from the first.
func Subtract(a, b float64) float64 {
	return a - b
}

func Multiply(a, b float64) float64 {
	return a * b
}

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("Can't divide by zero")
	}
	return a / b, nil
}

func Sqrt(a float64) (float64, error) {
	if a < 0 {
		return 0, fmt.Errorf("Can't Sqrt with negative numbers")
	}
	return math.Sqrt(a), nil
}
